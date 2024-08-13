// * Requests initiated by logged in users
package requests

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/qdrant"
	"github.com/stablecog/sc-go/discobot/aspectratio"
	"github.com/stablecog/sc-go/utils"
)

// For filtering user's generations
type SortOrder string

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

type OrderBy string

const (
	OrderByCreatedAt         OrderBy = "created_at"
	OrderByUpdatedAt         OrderBy = "updated_at"
	OrderByLikeCount         OrderBy = "like_count"
	OrderByLikeCountTrending OrderBy = "like_count_trending"
)

type UpscaleStatus string

const (
	// Include upscaled and not upscaled
	UpscaleStatusAny UpscaleStatus = "any"
	// Only upscaled
	UpscaleStatusOnly UpscaleStatus = "only"
	// Not upscaled
	UpscaleStatusNot UpscaleStatus = "not"
)

type QueryGenerationFilters struct {
	ForHistory                bool                             `json:"-"`
	ForProfile                bool                             `json:"-"`
	ForAdmin                  bool                             `json:"-"`
	HideNsfw                  bool                             `json:"-"`
	ModelIDs                  []uuid.UUID                      `json:"model_ids"`
	SchedulerIDs              []uuid.UUID                      `json:"scheduler_ids"`
	MinHeight                 int32                            `json:"min_height"`
	MaxHeight                 int32                            `json:"max_height"`
	MinWidth                  int32                            `json:"min_width"`
	MaxWidth                  int32                            `json:"max_width"`
	Widths                    []int32                          `json:"widths"`
	Heights                   []int32                          `json:"heights"`
	MaxInferenceSteps         int32                            `json:"max_inference_steps"`
	MinInferenceSteps         int32                            `json:"min_inference_steps"`
	InferenceSteps            []int32                          `json:"inference_steps"`
	MaxGuidanceScale          float32                          `json:"max_guidance_scale"`
	MinGuidanceScale          float32                          `json:"min_guidance_scale"`
	GuidanceScales            []float32                        `json:"guidance_scales"`
	UpscaleStatus             UpscaleStatus                    `json:"upscale_status"`
	GalleryStatus             []generationoutput.GalleryStatus `json:"gallery_status"`
	Order                     SortOrder                        `json:"order"`
	StartDt                   *time.Time                       `json:"start_dt"`
	EndDt                     *time.Time                       `json:"end_dt"`
	UserID                    *uuid.UUID                       `json:"user_id"`
	OrderBy                   OrderBy                          `json:"order_by"`
	ScoreThreshold            *float32                         `json:"score_threshold,omitempty"`
	IsFavorited               *bool                            `json:"is_favorited,omitempty"`
	WasAutoSubmitted          *bool                            `json:"was_auto_submitted,omitempty"`
	PromptID                  *uuid.UUID                       `json:"prompt,omitempty"`
	IsPublic                  *bool                            `json:"is_public,omitempty"`
	IsLiked                   *bool                            `json:"is_liked,omitempty"`
	AestheticArtifactScoreLTE *float32                         `json:"aesthetic_artifact_score_lte,omitempty"`
	AestheticArtifactScoreGTE *float32                         `json:"aesthetic_artifact_score_gte,omitempty"`
	AestheticRatingScoreLTE   *float32                         `json:"aesthetic_rating_score_lte,omitempty"`
	AestheticRatingScoreGTE   *float32                         `json:"aesthetic_rating_score_gte,omitempty"`
	Username                  []string                         `json:"username,omitempty"`
	AspectRatio               []aspectratio.AspectRatio        `json:"aspect_ratio,omitempty"`
	Oversampling              *float32                         `json:"oversampling,omitempty"`
	AdminMode                 *bool                            `json:"admin_mode,omitempty"`
}

// Parse all filters into a QueryGenerationFilters struct
func (filters *QueryGenerationFilters) ParseURLQueryParameters(urlValues url.Values) error {
	for key, value := range urlValues {
		// admin_mode
		if key == "admin_mode" {
			if strings.ToLower(value[0]) == "true" {
				t := true
				filters.AdminMode = &t
			} else if strings.ToLower(value[0]) == "false" {
				f := false
				filters.AdminMode = &f
			} else {
				return fmt.Errorf("invalid admin_mode: '%s' expected 'true' or 'false'", value[0])
			}
		}
		// model_ids
		if key == "model_ids" {
			if strings.Contains(value[0], ",") {
				for _, modelId := range strings.Split(value[0], ",") {
					parsed, err := uuid.Parse(modelId)
					if err != nil {
						return fmt.Errorf("invalid model id: %s", modelId)
					}
					filters.ModelIDs = append(filters.ModelIDs, parsed)
				}
			} else {
				parsed, err := uuid.Parse(value[0])
				if err != nil {
					return fmt.Errorf("invalid model id: %s", value[0])
				}
				filters.ModelIDs = []uuid.UUID{parsed}
			}
		}
		// scheduler_ids
		if key == "scheduler_ids" {
			if strings.Contains(value[0], ",") {
				for _, schedulerId := range strings.Split(value[0], ",") {
					parsed, err := uuid.Parse(schedulerId)
					if err != nil {
						return fmt.Errorf("invalid scheduler id: %s", schedulerId)
					}
					filters.SchedulerIDs = append(filters.SchedulerIDs, parsed)
				}
			} else {
				parsed, err := uuid.Parse(value[0])
				if err != nil {
					return fmt.Errorf("invalid scheduler id: %s", value[0])
				}
				filters.SchedulerIDs = []uuid.UUID{parsed}
			}
		}
		// Min and max height
		if key == "min_height" {
			minHeight, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid min height: %s", value[0])
			}
			if minHeight > math.MaxInt32 {
				return fmt.Errorf("min height too large: %d", minHeight)
			}
			filters.MinHeight = int32(minHeight)
		}
		if key == "max_height" {
			maxHeight, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid max height: %s", value[0])
			}
			if maxHeight > math.MaxInt32 {
				return fmt.Errorf("max height too large: %d", maxHeight)
			}
			filters.MaxHeight = int32(maxHeight)
		}
		// Min and max width
		if key == "min_width" {
			minWidth, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid min width: %s", value[0])
			}
			if minWidth > math.MaxInt32 {
				return fmt.Errorf("min width too large: %d", minWidth)
			}
			filters.MinWidth = int32(minWidth)
		}
		if key == "max_width" {
			maxWidth, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid max width: %s", value[0])
			}
			if maxWidth > math.MaxInt32 {
				return fmt.Errorf("max width too large: %d", maxWidth)
			}
			filters.MaxWidth = int32(maxWidth)
		}
		// Min and max inference steps
		if key == "min_inference_steps" {
			minInferenceSteps, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid min inference steps: %s", value[0])
			}
			if minInferenceSteps > math.MaxInt32 {
				return fmt.Errorf("min inference steps too large: %d", minInferenceSteps)
			}
			filters.MinInferenceSteps = int32(minInferenceSteps)
		}
		if key == "max_inference_steps" {
			maxInferenceSteps, err := strconv.Atoi(value[0])
			if err != nil {
				return fmt.Errorf("invalid max inference steps: %s", value[0])
			}
			if maxInferenceSteps > math.MaxInt32 {
				return fmt.Errorf("max inference steps too large: %d", maxInferenceSteps)
			}
			filters.MaxInferenceSteps = int32(maxInferenceSteps)
		}
		// Min and max guidance scale, the same but float32 not int32
		if key == "min_guidance_scale" {
			minGuidanceScale, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid min guidance scale: %s", value[0])
			}
			filters.MinGuidanceScale = float32(minGuidanceScale)
		}
		if key == "max_guidance_scale" {
			maxGuidanceScale, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid max guidance scale: %s", value[0])
			}
			filters.MaxGuidanceScale = float32(maxGuidanceScale)
		}
		// Widths
		if key == "widths" {
			if strings.Contains(value[0], ",") {
				for _, width := range strings.Split(value[0], ",") {
					parsed, err := strconv.Atoi(width)
					if err != nil {
						return fmt.Errorf("invalid width: %s", width)
					}
					if parsed > math.MaxInt32 {
						return fmt.Errorf("width too large: %d", parsed)
					}
					filters.Widths = append(filters.Widths, int32(parsed))
				}
			} else {
				parsed, err := strconv.Atoi(value[0])
				if err != nil {
					return fmt.Errorf("invalid width: %s", value[0])
				}
				if parsed > math.MaxInt32 {
					return fmt.Errorf("width too large: %d", parsed)
				}
				filters.Widths = []int32{int32(parsed)}
			}
		}
		// Heights
		if key == "heights" {
			if strings.Contains(value[0], ",") {
				for _, height := range strings.Split(value[0], ",") {
					parsed, err := strconv.Atoi(height)
					if err != nil {
						return fmt.Errorf("invalid height: %s", height)
					}
					if parsed > math.MaxInt32 {
						return fmt.Errorf("height too large: %d", parsed)
					}
					filters.Heights = append(filters.Heights, int32(parsed))
				}
			} else {
				parsed, err := strconv.Atoi(value[0])
				if err != nil {
					return fmt.Errorf("invalid height: %s", value[0])
				}
				if parsed > math.MaxInt32 {
					return fmt.Errorf("height too large: %d", parsed)
				}
				filters.Heights = []int32{int32(parsed)}
			}
		}
		// Inference Steps
		if key == "inference_steps" {
			if strings.Contains(value[0], ",") {
				for _, inferenceStep := range strings.Split(value[0], ",") {
					parsed, err := strconv.Atoi(inferenceStep)
					if err != nil {
						return fmt.Errorf("invalid inference step: %s", inferenceStep)
					}
					if parsed > math.MaxInt32 {
						return fmt.Errorf("inference step too large: %d", parsed)
					}
					filters.InferenceSteps = append(filters.InferenceSteps, int32(parsed))
				}
			} else {
				parsed, err := strconv.Atoi(value[0])
				if err != nil {
					return fmt.Errorf("invalid inference step: %s", value[0])
				}
				if parsed > math.MaxInt32 {
					return fmt.Errorf("inference step too large: %d", parsed)
				}
				filters.InferenceSteps = []int32{int32(parsed)}
			}
		}
		// Guidance Scales
		if key == "guidance_scales" {
			if strings.Contains(value[0], ",") {
				for _, guidanceScale := range strings.Split(value[0], ",") {
					parsed, err := strconv.ParseFloat(guidanceScale, 32)
					if err != nil {
						return fmt.Errorf("invalid guidance scale: %s", guidanceScale)
					}
					filters.GuidanceScales = append(filters.GuidanceScales, float32(parsed))
				}
			} else {
				parsed, err := strconv.ParseFloat(value[0], 32)
				if err != nil {
					return fmt.Errorf("invalid guidance scale: %s", value[0])
				}
				filters.GuidanceScales = []float32{float32(parsed)}
			}
		}

		// Order
		if key == "order" {
			if strings.ToLower(value[0]) == string(SortOrderAscending) {
				filters.Order = SortOrderAscending
			} else if strings.ToLower(value[0]) == string(SortOrderDescending) {
				filters.Order = SortOrderDescending
			} else {
				return fmt.Errorf("invalid order: '%s' expected '%s' or '%s'", value[0], SortOrderAscending, SortOrderDescending)
			}
		}

		// Upscale status
		if key == "upscaled" {
			if strings.ToLower(value[0]) == string(UpscaleStatusAny) {
				filters.UpscaleStatus = UpscaleStatusAny
			} else if strings.ToLower(value[0]) == string(UpscaleStatusNot) {
				filters.UpscaleStatus = UpscaleStatusNot
			} else if strings.ToLower(value[0]) == string(UpscaleStatusOnly) {
				filters.UpscaleStatus = UpscaleStatusOnly
			} else {
				return fmt.Errorf("invalid upscaled: '%s' expected '%s', '%s', or '%s'", value[0], UpscaleStatusAny, UpscaleStatusNot, UpscaleStatusOnly)
			}
		}

		if key == "oversampling" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid oversampling: %s", value[0])
			}
			filters.Oversampling = utils.ToPtr(float32(parsed))
		}

		// Gallery status
		if key == "gallery_status" {
			var statuses []string
			if strings.Contains(value[0], ",") {
				statuses = strings.Split(value[0], ",")
			} else {
				statuses = []string{value[0]}
			}
			for _, status := range statuses {
				if strings.ToLower(status) == string(generationoutput.GalleryStatusApproved) {
					filters.GalleryStatus = append(filters.GalleryStatus, generationoutput.GalleryStatusApproved)
				} else if strings.ToLower(status) == string(generationoutput.GalleryStatusRejected) {
					filters.GalleryStatus = append(filters.GalleryStatus, generationoutput.GalleryStatusRejected)
				} else if strings.ToLower(status) == string(generationoutput.GalleryStatusSubmitted) {
					filters.GalleryStatus = append(filters.GalleryStatus, generationoutput.GalleryStatusSubmitted)
				} else if strings.ToLower(status) == string(generationoutput.GalleryStatusNotSubmitted) {
					filters.GalleryStatus = append(filters.GalleryStatus, generationoutput.GalleryStatusNotSubmitted)
				} else if strings.ToLower(status) == string(generationoutput.GalleryStatusWaitingForApproval) {
					filters.GalleryStatus = append(filters.GalleryStatus, generationoutput.GalleryStatusWaitingForApproval)
				} else {
					return fmt.Errorf("invalid gallery_status: '%s' expected '%s', '%s', '%s', %s, or '%s'", value[0], generationoutput.GalleryStatusApproved, generationoutput.GalleryStatusRejected, generationoutput.GalleryStatusSubmitted, generationoutput.GalleryStatusNotSubmitted, generationoutput.GalleryStatusWaitingForApproval)
				}
			}
		}

		// Start and end date
		if key == "start_dt" {
			startDt, err := utils.ParseIsoTime(value[0])
			if err != nil {
				return fmt.Errorf("invalid start_dt: %s", value[0])
			}
			filters.StartDt = &startDt
		}
		if key == "end_dt" {
			endDt, err := utils.ParseIsoTime(value[0])
			if err != nil {
				return fmt.Errorf("invalid end_dt: %s", value[0])
			}
			filters.EndDt = &endDt
		}
		if key == "order_by" {
			if strings.ToLower(value[0]) == string(OrderByUpdatedAt) {
				filters.OrderBy = OrderByUpdatedAt
			} else if strings.ToLower(value[0]) == string(OrderByCreatedAt) {
				filters.OrderBy = OrderByCreatedAt
			} else {
				return fmt.Errorf("invalid order_by: '%s' expected '%s' or '%s'", value[0], OrderByUpdatedAt, OrderByCreatedAt)
			}
		}

		if key == "sort" {
			if strings.ToLower(value[0]) == "trending" {
				filters.OrderBy = OrderByLikeCountTrending
			} else if strings.ToLower(value[0]) == "new" {
				filters.OrderBy = OrderByCreatedAt
			} else if strings.ToLower(value[0]) == "top" {
				filters.OrderBy = OrderByLikeCount
			} else {
				return fmt.Errorf("invalid sort: '%s' expected 'trending', 'new', or 'top'", value[0])
			}
		}

		// Favorited
		if key == "is_favorited" {
			if strings.ToLower(value[0]) == "true" {
				t := true
				filters.IsFavorited = &t
			} else if strings.ToLower(value[0]) == "false" {
				f := false
				filters.IsFavorited = &f
			} else {
				return fmt.Errorf("invalid is_favorited: '%s' expected 'true' or 'false'", value[0])
			}
		}

		// Was auto submitted
		if key == "was_auto_submitted" {
			if strings.ToLower(value[0]) == "true" {
				t := true
				filters.WasAutoSubmitted = &t
			} else if strings.ToLower(value[0]) == "false" {
				f := false
				filters.WasAutoSubmitted = &f
			} else {
				return fmt.Errorf("invalid was_auto_submitted: '%s' expected 'true' or 'false'", value[0])
			}
		}

		// Liked
		if key == "is_liked" {
			if strings.ToLower(value[0]) == "true" {
				t := true
				filters.IsLiked = &t
				filters.IsPublic = &t
			} else {
				return fmt.Errorf("invalid is_liked: '%s' expected 'true'", value[0])
			}
		}

		// Score threshold
		if key == "score_threshold" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid score_threshold: %s", value[0])
			}
			filters.ScoreThreshold = utils.ToPtr(float32(parsed))
		}
		// Prompt id
		if key == "prompt_id" {
			parsed, err := uuid.Parse(value[0])
			if err != nil {
				return fmt.Errorf("invalid prompt_id: %s", value[0])
			}
			filters.PromptID = &parsed
		}

		// Aesthetic scores
		if key == "aesthetic_artifact_score_lte" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid aesthetic_artifact_score_lte: %s", value[0])
			}
			filters.AestheticArtifactScoreLTE = utils.ToPtr(float32(parsed))
		}
		if key == "aesthetic_artifact_score_gte" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid aesthetic_artifact_score_gte: %s", value[0])
			}
			filters.AestheticArtifactScoreGTE = utils.ToPtr(float32(parsed))
		}
		if key == "aesthetic_rating_score_lte" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid aesthetic_rating_score_lte: %s", value[0])
			}
			filters.AestheticRatingScoreLTE = utils.ToPtr(float32(parsed))
		}
		if key == "aesthetic_rating_score_gte" {
			parsed, err := strconv.ParseFloat(value[0], 32)
			if err != nil {
				return fmt.Errorf("invalid aesthetic_rating_score_gte: %s", value[0])
			}
			filters.AestheticRatingScoreGTE = utils.ToPtr(float32(parsed))
		}

		// username
		if key == "username" {
			filters.Username = strings.Split(value[0], ",")
			for i, str := range filters.Username {
				filters.Username[i] = strings.ToLower(str)
			}
		}

		// aspect ratio
		if key == "aspect_ratio" {
			ratioStrings := strings.Split(value[0], ",")
			filters.AspectRatio = make([]aspectratio.AspectRatio, len(ratioStrings))
			for i, ratioString := range ratioStrings {
				ratio, err := aspectratio.GetAspectRatioBySimpleString(strings.ToLower(ratioString))
				if err != nil {
					return fmt.Errorf("invalid aspect_ratio: %s", ratioString)
				}
				filters.AspectRatio[i] = ratio
			}
		}
	}
	// Descending default
	if filters.Order == "" {
		filters.Order = SortOrderDescending
	}
	// Upscale status any default
	if filters.UpscaleStatus == "" {
		filters.UpscaleStatus = UpscaleStatusAny
	}
	// Sort by created_at by default
	if filters.OrderBy == "" {
		filters.OrderBy = OrderByCreatedAt
	}
	return nil
}

func (filters *QueryGenerationFilters) ToQdrantFilters(ignoreGalleryStatus bool) (f *qdrant.SearchRequest_Filter, scoreThreshold *float32) {
	f = &qdrant.SearchRequest_Filter{}

	if filters.UserID != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key:   "user_id",
			Match: &qdrant.SCValue{Value: filters.UserID.String()},
		})
	}

	if len(filters.ModelIDs) > 0 {
		for _, modelID := range filters.ModelIDs {
			f.Should = append(f.Should, qdrant.SCMatchCondition{
				Key:   "model",
				Match: &qdrant.SCValue{Value: modelID.String()},
			})
		}
	}

	if len(filters.SchedulerIDs) > 0 {
		for _, schedulerID := range filters.SchedulerIDs {
			f.Should = append(f.Should, qdrant.SCMatchCondition{
				Key:   "scheduler",
				Match: &qdrant.SCValue{Value: schedulerID.String()},
			})
		}
	}

	// Width/height filters
	if filters.MinHeight > 0 {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "height",
			Range: qdrant.SCRange[int32]{
				Gte: utils.ToPtr(filters.MinHeight),
			},
		})
	}
	if filters.MaxHeight > 0 {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "height",
			Range: qdrant.SCRange[int32]{
				Lte: utils.ToPtr(filters.MaxHeight),
			},
		})
	}
	if filters.MinWidth > 0 {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "width",
			Range: qdrant.SCRange[int32]{
				Gte: utils.ToPtr(filters.MinWidth),
			},
		})
	}
	if filters.MaxWidth > 0 {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "width",
			Range: qdrant.SCRange[int32]{
				Lte: utils.ToPtr(filters.MaxWidth),
			},
		})
	}

	// Aestheti cscores
	if filters.AestheticArtifactScoreLTE != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "aesthetic_artifact_score",
			Range: qdrant.SCRange[float32]{
				Lte: filters.AestheticArtifactScoreLTE,
			},
		})
	}
	if filters.AestheticArtifactScoreGTE != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "aesthetic_artifact_score",
			Range: qdrant.SCRange[float32]{
				Gte: filters.AestheticArtifactScoreGTE,
			},
		})
	}
	if filters.AestheticRatingScoreLTE != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "aesthetic_rating_score",
			Range: qdrant.SCRange[float32]{
				Lte: filters.AestheticRatingScoreLTE,
			},
		})
	}
	if filters.AestheticRatingScoreGTE != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "aesthetic_rating_score",
			Range: qdrant.SCRange[float32]{
				Gte: filters.AestheticRatingScoreGTE,
			},
		})
	}

	// Widths/heights
	if len(filters.Heights) > 0 {
		for _, height := range filters.Heights {
			f.Should = append(f.Should, qdrant.SCMatchCondition{
				Key:   "height",
				Match: &qdrant.SCValue{Value: height},
			})
		}
	}
	if len(filters.Widths) > 0 {
		for _, width := range filters.Widths {
			f.Should = append(f.Should, qdrant.SCMatchCondition{
				Key:   "width",
				Match: &qdrant.SCValue{Value: width},
			})
		}
	}

	// Aspect ratio
	if len(filters.AspectRatio) > 0 {
		// With aspect ratio must be like
		// (width=width_1 and height=height_1) or (width=width_2 and height=height_2) or ...
		// In qdrant this is represented as
		// should = [ must = [ width=width_1, height=height_1 ], must = [ width=width_2, height=height_2 ], ...
		shouldConditions := []qdrant.SCMatchCondition{}

		for _, ratio := range filters.AspectRatio {
			widths, heights := ratio.GetAllWidthHeightCombos()
			for i := 0; i < len(widths); i++ {
				// Create a new must condition for each width-height pair
				widthHeightMust := []qdrant.SCMatchCondition{
					{
						Key: "width",
						Match: &qdrant.SCValue{
							Value: widths[i],
						},
					},
					{
						Key: "height",
						Match: &qdrant.SCValue{
							Value: heights[i],
						},
					},
				}

				// Append the must condition to should conditions
				shouldConditions = append(shouldConditions, qdrant.SCMatchCondition{
					Must: widthHeightMust,
				})
			}
		}

		// Create the final must condition that includes the should conditions
		finalMust := qdrant.SCMatchCondition{
			Should: shouldConditions,
		}

		f.Must = append(f.Must, finalMust)
	}

	// Gallery status
	if !ignoreGalleryStatus && len(filters.GalleryStatus) > 0 {
		for _, galleryStatus := range filters.GalleryStatus {
			f.Must = append(f.Must, qdrant.SCMatchCondition{
				Key:   "gallery_status",
				Match: &qdrant.SCValue{Value: galleryStatus},
			})
		}
	}

	// Date range
	if filters.StartDt != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "created_at",
			Range: qdrant.SCRange[int64]{
				Gte: utils.ToPtr(filters.StartDt.Unix()),
			},
		})
	}
	if filters.EndDt != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "created_at",
			Range: qdrant.SCRange[int64]{
				Lte: utils.ToPtr(filters.EndDt.Unix()),
			},
		})
	}

	// Is favorited
	if filters.IsFavorited != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "is_favorited",
			Match: &qdrant.SCValue{
				Value: *filters.IsFavorited,
			},
		})
	}

	// Was auto submitted
	if filters.WasAutoSubmitted != nil {
		f.Must = append(f.Must, qdrant.SCMatchCondition{
			Key: "was_auto_submitted",
			Match: &qdrant.SCValue{
				Value: *filters.WasAutoSubmitted,
			},
		})
	}

	if filters.ScoreThreshold == nil {
		filters.ScoreThreshold = utils.ToPtr[float32](50)
	} else if *filters.ScoreThreshold < 50 {
		filters.ScoreThreshold = utils.ToPtr[float32](50)
	}

	return f, filters.ScoreThreshold
}
