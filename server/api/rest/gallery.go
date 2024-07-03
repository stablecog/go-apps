package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/qdrant"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/utils"
	"golang.org/x/exp/slices"
)

const GALLERY_PER_PAGE = 50

func (c *RestAPI) HandleSemanticSearchGallery(w http.ResponseWriter, r *http.Request) {
	// Get user for like data, if authenticated
	callingUser, err := c.GetUserIfAuthenticatedOnly(w, r)
	if err != nil {
		log.Error("Error getting user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}
	var callingUserId *uuid.UUID
	if callingUser != nil {
		callingUserId = utils.ToPtr(callingUser.ID)
	}

	// Get output_id param
	outputId := r.URL.Query().Get("output_id")
	if outputId != "" {
		// Validate output_id
		uid, err := uuid.Parse(outputId)
		if err != nil {
			responses.ErrBadRequest(w, r, "invalid_output_id", "")
			return
		}

		galleryData, err := c.Repo.RetrieveGalleryDataByID(uid, nil, callingUserId, false)
		if err != nil {
			if ent.IsNotFound(err) {
				responses.ErrNotFound(w, r, "generation_not_found")
				return
			}
			log.Error("Error retrieving gallery data", "err", err)
			responses.ErrInternalServerError(w, r, "Error retrieving gallery data")
			return
		}

		// Sanitize
		galleryData.UserID = nil

		render.Status(r, http.StatusOK)
		render.JSON(w, r, GalleryResponse[int]{
			Page: 1,
			Hits: []repository.GalleryData{*galleryData},
		})
		return
	}

	search := r.URL.Query().Get("search")
	cursor := r.URL.Query().Get("cursor")
	galleryData := []repository.GalleryData{}
	var nextCursorQdrant *uint
	var nextCursorPostgres *time.Time
	var nextOffsetPostgres *int

	// Parse filters
	filters := &requests.QueryGenerationFilters{}
	err = filters.ParseURLQueryParameters(r.URL.Query())
	if err != nil {
		responses.ErrBadRequest(w, r, err.Error(), "")
		return
	}

	// Validate query parameters
	perPage := GALLERY_PER_PAGE
	if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil {
			responses.ErrBadRequest(w, r, "per_page must be an integer", "")
			return
		} else if perPage < 1 || perPage > MAX_PER_PAGE {
			responses.ErrBadRequest(w, r, fmt.Sprintf("per_page must be between 1 and %d", MAX_PER_PAGE), "")
			return
		}
	}

	// Parse as qdrant filters
	qdrantFilters, scoreThreshold := filters.ToQdrantFilters(true)
	// Append gallery status requirement
	qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
		Key:   "gallery_status",
		Match: &qdrant.SCValue{Value: generationoutput.GalleryStatusApproved},
	})
	// Append public requirement
	qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
		Key:   "is_public",
		Match: &qdrant.SCValue{Value: true},
	})

	// Leverage qdrant for semantic search
	if search != "" {
		// Check username filter
		if filters != nil && len(filters.Username) > 0 {
			userIDs, err := c.Repo.GetUserIDsByUsernames(filters.Username)
			if err != nil {
				log.Error("Error getting user ids by usernames", "err", err)
				responses.ErrInternalServerError(w, r, "An unknown error occurred")
				return
			}
			if len(userIDs) == 0 {
				render.Status(r, http.StatusOK)
				render.JSON(w, r, GalleryResponse[*uint]{
					Next: nextCursorQdrant,
					Hits: galleryData,
				})
				return
			}
			shouldFilter := []qdrant.SCMatchCondition{}
			for _, userID := range userIDs {
				shouldFilter = append(shouldFilter, qdrant.SCMatchCondition{
					Key:   "user_id",
					Match: &qdrant.SCValue{Value: userID.String()},
				})
			}
			if len(shouldFilter) > 0 {
				qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
					Should: shouldFilter,
				})
			}
		}

		var offset *uint
		if cursor != "" {
			cursoru64, err := strconv.ParseUint(cursor, 10, 64)
			if err != nil {
				responses.ErrBadRequest(w, r, "cursor must be a valid uint", "")
				return
			}
			cursorU := uint(cursoru64)
			offset = &cursorU
		}
		// See if search is a uuid
		uid, err := uuid.Parse(search)
		var embeddings []float32
		if err == nil {
			// Get embeddings from qdrant
			getPointRes, err := c.Qdrant.GetPoint(uid, false)
			if err != nil {
				log.Error("Error getting point from qdrant", "err", err)
				if strings.Contains(err.Error(), "Error querying collection 404") {
					responses.ErrNotFound(w, r, "generation_not_found")
					return
				}
				responses.ErrInternalServerError(w, r, "An unknown error occurred")
				return
			}
			embeddings = getPointRes.Result.Vector.Image
		} else {
			embeddings, err = c.Clip.GetEmbeddingFromText(search, 3, true)
			if err != nil {
				log.Error("Error getting embeddings from clip service", "err", err)
				responses.ErrInternalServerError(w, r, "An unknown error occurred")
				return
			}
		}

		res, err := c.Qdrant.QueryGenerations(embeddings, perPage, offset, scoreThreshold, filters.Oversampling, qdrantFilters, false, false)
		if err != nil {
			log.Error("Error querying qdrant", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error occurred")
			return
		}

		// Get generation output ids
		var outputIds []uuid.UUID
		for _, hit := range res.Result {
			outputId, err := uuid.Parse(hit.Id)
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			outputIds = append(outputIds, outputId)
		}

		// Get gallery data
		galleryDataUnsorted, err := c.Repo.RetrieveGalleryDataWithOutputIDs(outputIds, callingUserId, false)
		if err != nil {
			log.Error("Error querying gallery data", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error occurred")
			return
		}
		gDataMap := make(map[uuid.UUID]repository.GalleryData)
		for _, gData := range galleryDataUnsorted {
			gDataMap[gData.ID] = gData
		}

		for _, hit := range res.Result {
			outputId, err := uuid.Parse(hit.Id)
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			item, ok := gDataMap[outputId]
			if !ok {
				log.Error("Error retrieving gallery data", "output_id", outputId)
				continue
			}
			galleryData = append(galleryData, item)
		}

		// Set next cursor
		nextCursorQdrant = res.Next
	} else {
		// Get most recent gallery data
		var qCursor *time.Time
		var qOffset *int
		if cursor != "" && filters != nil && (filters.OrderBy == requests.OrderByLikeCountTrending || filters.OrderBy == requests.OrderByLikeCount) {
			offset, err := strconv.Atoi(cursor)
			if err != nil {
				responses.ErrBadRequest(w, r, "cursor must be a valid uint", "")
				return
			}
			qOffset = &offset
		} else if cursor != "" {
			cursorTime, err := utils.ParseIsoTime(cursor)
			if err != nil {
				responses.ErrBadRequest(w, r, "cursor must be a valid iso time string", "")
				return
			}
			qCursor = &cursorTime
		}

		// Retrieve from postgres
		filters.GalleryStatus = []generationoutput.GalleryStatus{generationoutput.GalleryStatusApproved}
		filters.IsPublic = utils.ToPtr(true)
		galleryData, nextCursorPostgres, nextOffsetPostgres, err = c.Repo.RetrieveMostRecentGalleryDataV3(filters, callingUserId, perPage, qCursor, qOffset)
		if err != nil {
			log.Error("Error querying gallery data from postgres", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error occurred")
			return
		}
	}

	// Shuffle results if no search was specified and no seed was specified
	seed := r.URL.Query().Get("seed")
	if search == "" && seed != "" {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			log.Error("Error parsing seed", "err", err)
		} else {
			seededRand := rand.New(rand.NewSource(int64(seedInt)))
			seededRand.Shuffle(
				len(galleryData),
				func(i, j int) { galleryData[i], galleryData[j] = galleryData[j], galleryData[i] },
			)
		}
	}

	// We don't want to leak primary keys, so set to nil
	for i := range galleryData {
		galleryData[i].UserID = nil
	}

	if search == "" {
		render.Status(r, http.StatusOK)
		if filters.OrderBy == requests.OrderByLikeCount || filters.OrderBy == requests.OrderByLikeCountTrending {
			render.JSON(w, r, GalleryResponse[*int]{
				Next: nextOffsetPostgres,
				Hits: galleryData,
			})
			return
		}
		render.JSON(w, r, GalleryResponse[*time.Time]{
			Next: nextCursorPostgres,
			Hits: galleryData,
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, GalleryResponse[*uint]{
		Next: nextCursorQdrant,
		Hits: galleryData,
	})
}

type GalleryResponseCursor interface {
	*uint | *time.Time | int | *int
}

type GalleryResponse[T GalleryResponseCursor] struct {
	Total    *uint                    `json:"total,omitempty"`
	Next     T                        `json:"next,omitempty"`
	Page     int                      `json:"page"`
	Hits     []repository.GalleryData `json:"hits"`
	Metadata *UserProfileMetadata     `json:"metadata,omitempty"`
}

type GalleryResponseV3[T GalleryResponseCursor] struct {
	Total    *uint                                 `json:"total,omitempty"`
	Next     T                                     `json:"next,omitempty"`
	Page     int                                   `json:"page"`
	Outputs  []repository.V3GenerationOutputResult `json:"outputs"`
	Metadata *UserProfileMetadata                  `json:"metadata,omitempty"`
}

// HTTP PUT submit a generation to gallery - for user
// Only allow submitting user's own gallery items.
func (c *RestAPI) HandleSubmitGenerationToGallery(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var submitToGalleryReq requests.SubmitGalleryRequest
	err := json.Unmarshal(reqBody, &submitToGalleryReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	submitted, err := c.Repo.SubmitGenerationOutputsToGalleryForUser(submitToGalleryReq.GenerationOutputIDs, user.ID)
	if err != nil {
		responses.ErrInternalServerError(w, r, "Error submitting generation outputs to gallery")
		return
	}

	res := responses.SubmitGalleryResponse{
		Submitted: submitted,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP PUT make generation output public - for user
// Only allow submitting user's own gallery items.
func (c *RestAPI) HandleMakeGenerationOutputsPublic(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var submitToGalleryReq requests.SubmitGalleryRequest
	err := json.Unmarshal(reqBody, &submitToGalleryReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	updated, err := c.Repo.MakeGenerationOutputsPublicForUser(submitToGalleryReq.GenerationOutputIDs, user.ID)
	if err != nil {
		responses.ErrInternalServerError(w, r, "Error submitting generation outputs to gallery")
		return
	}

	res := responses.UpdatedResponse{
		Updated: updated,
	}

	if updated > 0 {
		render.Status(r, http.StatusOK)
	} else {
		render.Status(r, http.StatusBadRequest)
	}
	render.JSON(w, r, res)
}

// HTTP PUT make generation output public - for user
// Only allow submitting user's own gallery items.
func (c *RestAPI) HandleMakeGenerationOutputsPrivate(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	roles, err := c.Repo.GetRoles(user.ID)
	if err != nil {
		log.Error("Error getting roles for user", "err", err)
		responses.ErrInternalServerError(w, r, "Error getting roles for user")
		return
	}
	isSuperAdmin := slices.Contains(roles, "SUPER_ADMIN")

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var submitToGalleryReq requests.SubmitGalleryRequest
	err = json.Unmarshal(reqBody, &submitToGalleryReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	updated, err := c.Repo.MakeGenerationOutputsPrivateForUser(submitToGalleryReq.GenerationOutputIDs, user.ID, user.ActiveProductID != nil || isSuperAdmin)
	if err != nil {
		responses.ErrInternalServerError(w, r, "Error submitting generation outputs to gallery")
		return
	}

	res := responses.UpdatedResponse{
		Updated: updated,
	}

	if updated > 0 {
		render.Status(r, http.StatusOK)
	} else {
		render.Status(r, http.StatusBadRequest)
	}
	render.JSON(w, r, res)
}
