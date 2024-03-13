package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/generationoutputlike"
	"github.com/stablecog/sc-go/database/ent/negativeprompt"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/prompt"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/utils"
)

// Get generations by IDs
func (r *Repository) GetGenerationsByIDList(ids []uuid.UUID) ([]*ent.Generation, error) {
	return r.DB.Generation.Query().Where(generation.IDIn(ids...)).All(r.Ctx)
}

// Get generation by ID
func (r *Repository) GetGeneration(id uuid.UUID) (*ent.Generation, error) {
	return r.DB.Generation.Query().Where(generation.IDEQ(id)).First(r.Ctx)
}

// Get generation output by ID
func (r *Repository) GetGenerationOutput(id uuid.UUID) (*ent.GenerationOutput, error) {
	return r.DB.GenerationOutput.Query().Where(generationoutput.IDEQ(id)).First(r.Ctx)
}

// Get public generation output by ID
// Get generation output by ID
func (r *Repository) GetPublicGenerationOutput(id uuid.UUID) (*ent.GenerationOutput, error) {
	return r.DB.GenerationOutput.Query().Where(generationoutput.IDEQ(id), generationoutput.GalleryStatusIn(generationoutput.GalleryStatusApproved, generationoutput.GalleryStatusSubmitted, generationoutput.GalleryStatusWaitingForApproval)).First(r.Ctx)
}

// Get generation output for user
func (r *Repository) GetGenerationOutputForUser(id uuid.UUID, userID uuid.UUID) (*ent.GenerationOutput, error) {
	return r.DB.Generation.Query().Where(generation.UserIDEQ(userID)).QueryGenerationOutputs().Where(generationoutput.IDEQ(id)).First(r.Ctx)
}

// Get width/height for generation output
func (r *Repository) GetGenerationOutputWidthHeight(outputID uuid.UUID) (width, height int32, err error) {
	gen, err := r.DB.GenerationOutput.Query().Where(generationoutput.IDEQ(outputID)).QueryGenerations().Select(generation.FieldWidth, generation.FieldHeight).First(r.Ctx)
	if err != nil {
		return 0, 0, err
	}
	return gen.Width, gen.Height, nil
}

// Get last N generations, a basic view
func (r *Repository) GetGenerations(limit int) ([]*ent.Generation, error) {
	return r.DB.Generation.Query().Order(ent.Desc(generation.FieldCreatedAt)).Limit(limit).All(r.Ctx)
}

// Get avg(started_at - created_at) over a time period with limit
func (r *Repository) GetAvgGenerationQueueTime(since time.Time, limit int) (float64, error) {
	var rawQ []struct {
		CreatedAt time.Time `json:"created_at" sql:"created_at"`
		StartedAt time.Time `json:"started_at" sql:"started_at"`
		QueueS    float64   `json:"queue_s" sql:"queue_s"`
	}
	q := r.DB.Generation.Query().Where(generation.StatusEQ(generation.StatusSucceeded), generation.CreatedAtGT(since)).
		Order(ent.Desc(generation.FieldCreatedAt))
	if limit > 0 {
		q = q.Limit(limit)
	}
	err := q.GroupBy(generation.FieldCreatedAt, generation.FieldStartedAt).
		Aggregate(func(s *sql.Selector) string {
			var raw string
			switch r.ConnInfo.Dialect() {
			case "pgx":
				raw = fmt.Sprintf("EXTRACT(EPOCH FROM (%s - %s))", s.C(generation.FieldStartedAt), s.C(generation.FieldCreatedAt))
			default:
				raw = fmt.Sprintf("ROUND((JULIANDAY(%s) - JULIANDAY(%s)) * 86400)", s.C(generation.FieldStartedAt), s.C(generation.FieldCreatedAt))
			}
			return sql.As(raw, "queue_s")
		}).Scan(r.Ctx, &rawQ)
	if err != nil {
		return 0, err
	}

	if len(rawQ) == 0 {
		return 0, nil
	}
	var avg float64
	for _, q := range rawQ {
		avg += q.QueueS
	}
	if len(rawQ) == 0 {
		return 0, nil
	}
	return avg / float64(len(rawQ)), nil
}

// Apply all filters to root ent query
func (r *Repository) ApplyUserGenerationsFilters(query *ent.GenerationQuery, filters *requests.QueryGenerationFilters, omitEdges bool) *ent.GenerationQuery {
	resQuery := query
	if filters != nil {
		// Apply filters
		if len(filters.ModelIDs) > 0 {
			resQuery = resQuery.Where(generation.ModelIDIn(filters.ModelIDs...))
		}
		if len(filters.SchedulerIDs) > 0 {
			resQuery = resQuery.Where(generation.SchedulerIDIn(filters.SchedulerIDs...))
		}
		// Apply OR if both are present
		// Confusing, but example of what we want to do:
		// If min_width=100, max_width=200, widths=[300,400]
		// We want to query like; WHERE (width >= 100 AND width <= 200) OR width IN (300,400)
		if (filters.MinWidth != 0 || filters.MaxWidth != 0) && len(filters.Widths) > 0 {
			if filters.MinWidth != 0 && filters.MaxWidth != 0 {
				resQuery = resQuery.Where(generation.Or(generation.And(generation.WidthGTE(filters.MinWidth), generation.WidthLTE(filters.MaxWidth)), generation.WidthIn(filters.Widths...)))
			} else if filters.MinWidth != 0 {
				resQuery = resQuery.Where(generation.Or(generation.WidthGTE(filters.MinWidth), generation.WidthIn(filters.Widths...)))
			} else {
				resQuery = resQuery.Where(generation.Or(generation.WidthLTE(filters.MaxWidth), generation.WidthIn(filters.Widths...)))
			}
		} else {
			if filters.MinWidth != 0 {
				resQuery = resQuery.Where(generation.WidthGTE(filters.MinWidth))
			}
			if filters.MaxWidth != 0 {
				resQuery = resQuery.Where(generation.WidthLTE(filters.MaxWidth))
			}
			if len(filters.Widths) > 0 {
				resQuery = resQuery.Where(generation.WidthIn(filters.Widths...))
			}
		}

		// Height
		if (filters.MinHeight != 0 || filters.MaxHeight != 0) && len(filters.Heights) > 0 {
			if filters.MinHeight != 0 && filters.MaxHeight != 0 {
				resQuery = resQuery.Where(generation.Or(generation.And(generation.HeightGTE(filters.MinHeight), generation.HeightLTE(filters.MaxHeight)), generation.HeightIn(filters.Heights...)))
			} else if filters.MinHeight != 0 {
				resQuery = resQuery.Where(generation.Or(generation.HeightGTE(filters.MinHeight), generation.HeightIn(filters.Heights...)))
			} else {
				resQuery = resQuery.Where(generation.Or(generation.HeightLTE(filters.MaxHeight), generation.HeightIn(filters.Heights...)))
			}
		} else {
			if len(filters.Heights) > 0 {
				resQuery = resQuery.Where(generation.HeightIn(filters.Heights...))
			}
			if filters.MaxHeight != 0 {
				resQuery = resQuery.Where(generation.HeightLTE(filters.MaxHeight))
			}
			if filters.MinHeight != 0 {
				resQuery = resQuery.Where(generation.HeightGTE(filters.MinHeight))
			}
		}

		// Aspect ratios
		if len(filters.AspectRatio) > 0 {
			var widthHeightConditions []predicate.Generation

			for _, aspectRatio := range filters.AspectRatio {
				widths, heights := aspectRatio.GetAllWidthHeightCombos()
				for i := 0; i < len(widths); i++ {
					// Ensure that there is a corresponding height for each width
					if i < len(heights) {
						// Create an AND condition for each width-height pair
						widthHeightCondition := generation.And(
							generation.WidthEQ(widths[i]),
							generation.HeightEQ(heights[i]),
						)
						widthHeightConditions = append(widthHeightConditions, widthHeightCondition)
					}
				}
			}

			// Combine all the width-height conditions using OR
			if len(widthHeightConditions) > 0 {
				combinedCondition := widthHeightConditions[0]
				for _, cond := range widthHeightConditions[1:] {
					combinedCondition = generation.Or(combinedCondition, cond)
				}
				resQuery = resQuery.Where(combinedCondition)
			}
		}

		// Inference steps
		if (filters.MinInferenceSteps != 0 || filters.MaxInferenceSteps != 0) && len(filters.InferenceSteps) > 0 {
			if filters.MinInferenceSteps != 0 && filters.MaxInferenceSteps != 0 {
				resQuery = resQuery.Where(generation.Or(generation.And(generation.InferenceStepsGTE(filters.MinInferenceSteps), generation.InferenceStepsLTE(filters.MaxInferenceSteps)), generation.InferenceStepsIn(filters.InferenceSteps...)))
			} else if filters.MinInferenceSteps != 0 {
				resQuery = resQuery.Where(generation.Or(generation.InferenceStepsGTE(filters.MinInferenceSteps), generation.InferenceStepsIn(filters.InferenceSteps...)))
			} else {
				resQuery = resQuery.Where(generation.Or(generation.InferenceStepsLTE(filters.MaxInferenceSteps), generation.InferenceStepsIn(filters.InferenceSteps...)))
			}
		} else {
			if len(filters.InferenceSteps) > 0 {
				resQuery = resQuery.Where(generation.InferenceStepsIn(filters.InferenceSteps...))
			}
			if filters.MaxInferenceSteps != 0 {
				resQuery = resQuery.Where(generation.InferenceStepsLTE(filters.MaxInferenceSteps))
			}
			if filters.MinInferenceSteps != 0 {
				resQuery = resQuery.Where(generation.InferenceStepsGTE(filters.MinInferenceSteps))
			}
		}

		// Guidance Scales
		if (filters.MinGuidanceScale != 0 || filters.MaxGuidanceScale != 0) && len(filters.GuidanceScales) > 0 {
			if filters.MinGuidanceScale != 0 && filters.MaxGuidanceScale != 0 {
				resQuery = resQuery.Where(generation.Or(generation.And(generation.GuidanceScaleGTE(filters.MinGuidanceScale), generation.GuidanceScaleLTE(filters.MaxGuidanceScale)), generation.GuidanceScaleIn(filters.GuidanceScales...)))
			} else if filters.MinGuidanceScale != 0 {
				resQuery = resQuery.Where(generation.Or(generation.GuidanceScaleGTE(filters.MinGuidanceScale), generation.GuidanceScaleIn(filters.GuidanceScales...)))
			} else {
				resQuery = resQuery.Where(generation.Or(generation.GuidanceScaleLTE(filters.MaxGuidanceScale), generation.GuidanceScaleIn(filters.GuidanceScales...)))
			}
		} else {
			if len(filters.GuidanceScales) > 0 {
				resQuery = resQuery.Where(generation.GuidanceScaleIn(filters.GuidanceScales...))
			}
			if filters.MaxGuidanceScale != 0 {
				resQuery = resQuery.Where(generation.GuidanceScaleLTE(filters.MaxGuidanceScale))
			}
			if filters.MinGuidanceScale != 0 {
				resQuery = resQuery.Where(generation.GuidanceScaleGTE(filters.MinGuidanceScale))
			}
		}

		if !omitEdges {
			// Upscaled
			if filters.UpscaleStatus == requests.UpscaleStatusNot {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.IsNull("upscaled_image_path"))
				})
			} else if filters.UpscaleStatus == requests.UpscaleStatusOnly {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.NotNull("upscaled_image_path"))
				})
			}

			if len(filters.GalleryStatus) > 0 {
				v := make([]any, len(filters.GalleryStatus))
				for i := range v {
					v[i] = filters.GalleryStatus[i]
				}
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.In(generationoutput.FieldGalleryStatus, v...))
				})
			}

			if filters.IsFavorited != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.EQ(generationoutput.FieldIsFavorited, *filters.IsFavorited))
				})
			}

			if filters.IsPublic != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.EQ(generationoutput.FieldIsPublic, *filters.IsPublic))
				})
			}

			if filters.AestheticArtifactScoreGTE != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.GTE(generationoutput.FieldAestheticArtifactScore, *filters.AestheticArtifactScoreGTE))
				})
			}

			if filters.AestheticArtifactScoreLTE != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.LTE(generationoutput.FieldAestheticArtifactScore, *filters.AestheticArtifactScoreLTE))
				})
			}

			if filters.AestheticRatingScoreGTE != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.GTE(generationoutput.FieldAestheticRatingScore, *filters.AestheticRatingScoreGTE))
				})
			}

			if filters.AestheticRatingScoreLTE != nil {
				resQuery = resQuery.Where(func(s *sql.Selector) {
					s.Where(sql.LTE(generationoutput.FieldAestheticRatingScore, *filters.AestheticRatingScoreLTE))
				})
			}
		}

		// prompt ID
		if filters.PromptID != nil {
			resQuery = resQuery.Where(generation.PromptIDEQ(*filters.PromptID))
		}

		// Start dt
		if filters.StartDt != nil {
			resQuery = resQuery.Where(generation.CreatedAtGTE(*filters.StartDt))
		}

		// End dt
		if filters.EndDt != nil {
			resQuery = resQuery.Where(generation.CreatedAtLTE(*filters.EndDt))
		}

		if filters.WasAutoSubmitted != nil {
			resQuery = resQuery.Where(generation.WasAutoSubmittedEQ(*filters.WasAutoSubmitted))
		}

		if filters.UserID != nil && filters.IsLiked == nil {
			resQuery = resQuery.Where(generation.UserIDEQ(*filters.UserID))
		}
	}
	return resQuery
}

// Gets the count of generations with outputs user has with filters
func (r *Repository) GetGenerationCount(filters *requests.QueryGenerationFilters) (int, error) {
	var query *ent.GenerationQuery

	query = r.DB.Generation.Query().
		Where(generation.StatusEQ(generation.StatusSucceeded))

	// Apply filters
	query = r.ApplyUserGenerationsFilters(query, filters, false)

	// Join other data
	var res []UserGenCount
	err := query.Modify(func(s *sql.Selector) {
		got := sql.Table(generationoutput.Table)
		golikesT := sql.Table(generationoutputlike.Table)
		ut := sql.Table(user.Table)
		ltj := s.Join(got).OnP(
			sql.And(
				sql.ColumnsEQ(s.C(generation.FieldID), got.C(generationoutput.FieldGenerationID)),
				sql.IsNull(got.C(generationoutput.FieldDeletedAt)),
			),
		)
		if filters != nil && filters.UserID != nil && filters.IsLiked == nil {
			ltj.Join(ut).OnP(
				sql.And(
					sql.ColumnsEQ(s.C(generation.FieldUserID), ut.C(user.FieldID)),
					sql.EQ(ut.C(user.FieldID), *filters.UserID),
				),
			)
		} else {
			ltj.LeftJoin(ut).On(
				s.C(generation.FieldUserID), ut.C(user.FieldID),
			)
		}
		if filters != nil && filters.UserID != nil && filters.IsLiked != nil {
			ltj.Join(golikesT).OnP(
				sql.And(
					sql.ColumnsEQ(got.C(generationoutput.FieldID), golikesT.C(generationoutputlike.FieldOutputID)),
					sql.EQ(golikesT.C(generationoutputlike.FieldLikedByUserID), *filters.UserID),
				),
			)
		}
		ltj.Select(sql.As(sql.Count("*"), "total"))
	}).Scan(r.Ctx, &res)
	if err != nil {
		return 0, err
	} else if len(res) == 0 {
		return 0, nil
	}
	return res[0].Total, nil
}

type UserGenCount struct {
	Total int `json:"total" sql:"total"`
}

// Get user generations in the *GenerationQueryWithOutputsMeta format
// Using a list of generation_output ids
// For when we semantic search against our vector db
func (r *Repository) RetrieveGenerationsWithOutputIDs(outputIDs []uuid.UUID, callingUserId *uuid.UUID, admin bool) (*GenerationQueryWithOutputsMeta[*uint], error) {
	query := r.DB.GenerationOutput.Query().Where(generationoutput.IDIn(outputIDs...))
	if callingUserId != nil {
		query = query.WithGenerationOutputLikes(func(gql *ent.GenerationOutputLikeQuery) {
			gql.Where(generationoutputlike.LikedByUserID(*callingUserId))
		})
	}
	gQueryResult, err := query.WithGenerations(func(gq *ent.GenerationQuery) {
		gq.WithPrompt()
		gq.WithNegativePrompt()
		gq.WithUser()
	}).All(r.Ctx)
	if err != nil {
		log.Errorf("Error retrieving generation outputs %v", err)
		return nil, err
	}

	if len(gQueryResult) == 0 {
		meta := &GenerationQueryWithOutputsMeta[*uint]{
			Outputs: []GenerationQueryWithOutputsResultFormatted{},
		}
		return meta, nil
	}

	meta := &GenerationQueryWithOutputsMeta[*uint]{}

	// Get real image URLs for each
	for i, g := range gQueryResult {
		if g.ImagePath != "" {
			parsed := utils.GetEnv().GetURLFromImagePath(g.ImagePath)
			gQueryResult[i].ImagePath = parsed
		}
		if g.UpscaledImagePath != nil {
			parsed := utils.GetEnv().GetURLFromImagePath(*g.UpscaledImagePath)
			gQueryResult[i].UpscaledImagePath = &parsed
		}
	}

	// Format to GenerationQueryWithOutputsResultFormatted
	generationOutputMap := make(map[uuid.UUID][]GenerationUpscaleOutput)
	for _, g := range gQueryResult {
		gOutput := GenerationUpscaleOutput{
			ID:                     g.ID,
			ImageUrl:               g.ImagePath,
			GalleryStatus:          g.GalleryStatus,
			WasAutoSubmitted:       g.Edges.Generations.WasAutoSubmitted,
			IsFavorited:            g.IsFavorited,
			IsPublic:               g.IsPublic,
			LikeCount:              utils.ToPtr(g.LikeCount),
			IsLiked:                utils.ToPtr(len(g.Edges.GenerationOutputLikes) > 0),
			AestheticRatingScore:   utils.ToPtr(g.AestheticRatingScore),
			AestheticArtifactScore: utils.ToPtr(g.AestheticArtifactScore),
		}
		if g.UpscaledImagePath != nil {
			gOutput.UpscaledImageUrl = *g.UpscaledImagePath
		}
		output := GenerationQueryWithOutputsResultFormatted{
			GenerationUpscaleOutput: gOutput,
			Generation: GenerationQueryWithOutputsData{
				ID:             g.ID,
				Height:         g.Edges.Generations.Height,
				Width:          g.Edges.Generations.Width,
				InferenceSteps: g.Edges.Generations.InferenceSteps,
				Seed:           g.Edges.Generations.Seed,
				Status:         g.Edges.Generations.Status.String(),
				GuidanceScale:  g.Edges.Generations.GuidanceScale,
				SchedulerID:    g.Edges.Generations.SchedulerID,
				ModelID:        g.Edges.Generations.ModelID,
				CreatedAt:      g.CreatedAt,
				UpdatedAt:      g.UpdatedAt,
				StartedAt:      g.Edges.Generations.StartedAt,
				CompletedAt:    g.Edges.Generations.CompletedAt,
				Prompt: PromptType{
					Text: g.Edges.Generations.Edges.Prompt.Text,
					ID:   *g.Edges.Generations.PromptID,
				},
				IsFavorited: gOutput.IsFavorited,
				IsPublic:    gOutput.IsPublic,
			},
		}
		if g.Edges.Generations.InitImageURL != nil {
			output.Generation.InitImageURL = *g.Edges.Generations.InitImageURL
			output.Generation.PromptStrength = g.Edges.Generations.PromptStrength
		}
		if g.Edges.Generations.Edges.NegativePrompt != nil {
			output.Generation.NegativePrompt = &PromptType{
				Text: g.Edges.Generations.Edges.NegativePrompt.Text,
				ID:   *g.Edges.Generations.NegativePromptID,
			}
		}
		if g.Edges.Generations.Edges.User != nil {
			if admin {
				output.Generation.User = &UserType{
					Email:      g.Edges.Generations.Edges.User.Email,
					Username:   g.Edges.Generations.Edges.User.Username,
					Identifier: utils.Sha256(g.Edges.Generations.Edges.User.ID.String()),
				}
			} else {
				output.Generation.User = &UserType{
					Username:   g.Edges.Generations.Edges.User.Username,
					Identifier: utils.Sha256(g.Edges.Generations.Edges.User.ID.String()),
				}
			}
		}
		generationOutputMap[g.ID] = append(generationOutputMap[g.ID], gOutput)
		meta.Outputs = append(meta.Outputs, output)
	}
	// Now loop through and add outputs to each generation
	for i, g := range meta.Outputs {
		meta.Outputs[i].Generation.Outputs = generationOutputMap[g.Generation.ID]
	}

	return meta, nil
}

// Get user generations from the database using page options
// Cursor actually represents created_at, we paginate using this for performance reasons
// If present, we will get results after the cursor (anything before, represents previous pages)
// ! using ent .With... doesn't use joins, so we construct our own query to make it more efficient
func (r *Repository) QueryGenerations(per_page int, cursor *time.Time, filters *requests.QueryGenerationFilters) (*GenerationQueryWithOutputsMeta[*time.Time], error) {
	// Base fields to select in our query
	selectFields := []string{
		generation.FieldID,
		generation.FieldWidth,
		generation.FieldHeight,
		generation.FieldInferenceSteps,
		generation.FieldSeed,
		generation.FieldStatus,
		generation.FieldGuidanceScale,
		generation.FieldSchedulerID,
		generation.FieldModelID,
		generation.FieldPromptID,
		generation.FieldNegativePromptID,
		generation.FieldCreatedAt,
		generation.FieldUpdatedAt,
		generation.FieldStartedAt,
		generation.FieldCompletedAt,
		generation.FieldWasAutoSubmitted,
		generation.FieldInitImageURL,
		generation.FieldPromptStrength,
	}
	var query *ent.GenerationQuery
	var gQueryResult []GenerationQueryWithOutputsResult

	// Figure out order bys
	var orderByGeneration []string
	var orderByOutput []string
	if filters == nil || (filters != nil && filters.OrderBy == requests.OrderByCreatedAt) {
		orderByGeneration = []string{generation.FieldCreatedAt}
		orderByOutput = []string{generationoutput.FieldCreatedAt}
	} else {
		orderByGeneration = []string{generation.FieldCreatedAt, generation.FieldUpdatedAt}
		orderByOutput = []string{generationoutput.FieldCreatedAt, generationoutput.FieldUpdatedAt}
	}

	query = r.DB.Generation.Query().Select(selectFields...).
		Where(generation.StatusEQ(generation.StatusSucceeded))
	if cursor != nil {
		query = query.Where(generation.CreatedAtLT(*cursor))
	}

	// Exclude deleted at always
	// query = query.Where(func(s *sql.Selector) {
	// 	s.Where(sql.IsNull("deleted_at"))
	// })

	// Apply filters
	query = r.ApplyUserGenerationsFilters(query, filters, false)

	// Limits is + 1 so we can check if there are more pages
	query = query.Limit(per_page + 1)

	// Join other data
	err := query.Modify(func(s *sql.Selector) {
		gt := sql.Table(generation.Table)
		got := sql.Table(generationoutput.Table)
		golikesT := sql.Table(generationoutputlike.Table)
		ut := sql.Table(user.Table)
		ltj := s.Join(got).OnP(
			sql.And(
				sql.ColumnsEQ(gt.C(generation.FieldID), got.C(generationoutput.FieldGenerationID)),
				sql.IsNull(got.C(generationoutput.FieldDeletedAt)),
			),
		)
		if filters != nil && filters.UserID != nil && filters.IsLiked == nil {
			ltj.Join(ut).OnP(
				sql.And(
					sql.ColumnsEQ(s.C(generation.FieldUserID), ut.C(user.FieldID)),
					sql.EQ(ut.C(user.FieldID), *filters.UserID),
				),
			)
		} else {
			ltj.LeftJoin(ut).On(
				s.C(generation.FieldUserID), ut.C(user.FieldID),
			)
		}
		groupByFields := []string{}
		if filters != nil && filters.UserID != nil && filters.IsLiked != nil {
			ltj.Join(golikesT).OnP(
				sql.And(
					sql.ColumnsEQ(got.C(generationoutput.FieldID), golikesT.C(generationoutputlike.FieldOutputID)),
					sql.EQ(golikesT.C(generationoutputlike.FieldLikedByUserID), *filters.UserID),
				),
			)
			ltj.AppendSelect(sql.As(golikesT.C(generationoutputlike.FieldCreatedAt), "liked_at"))
			groupByFields = append(groupByFields, golikesT.C(generationoutputlike.FieldCreatedAt))
		}
		groupByFields = append(groupByFields, []string{s.C(generation.FieldID),
			got.C(generationoutput.FieldID), got.C(generationoutput.FieldGalleryStatus),
			got.C(generationoutput.FieldImagePath), got.C(generationoutput.FieldUpscaledImagePath),
			ut.C(user.FieldUsername), ut.C(user.FieldID)}...)
		ltj.AppendSelect(sql.As(got.C(generationoutput.FieldID), "output_id"), sql.As(got.C(generationoutput.FieldLikeCount), "like_count"), sql.As(got.C(generationoutput.FieldGalleryStatus), "output_gallery_status"), sql.As(got.C(generationoutput.FieldImagePath), "image_path"), sql.As(got.C(generationoutput.FieldUpscaledImagePath), "upscaled_image_path"), sql.As(got.C(generationoutput.FieldDeletedAt), "deleted_at"), sql.As(got.C(generationoutput.FieldIsFavorited), "is_favorited"), sql.As(ut.C(user.FieldUsername), "username"), sql.As(ut.C(user.FieldID), "user_id"), sql.As(got.C(generationoutput.FieldIsPublic), "is_public")).
			GroupBy(groupByFields...)
		orderDir := "asc"
		if filters == nil || (filters != nil && filters.Order == requests.SortOrderDescending) {
			orderDir = "desc"
		}

		var orderByCombined []string

		// Check if the IsLiked filter is set and prepend the order by from golikesT.
		if filters.IsLiked != nil {
			// Assuming "liked" is the column you want to order by in the golikesT table.
			// You may need to adjust this to fit the actual column you're interested in.
			if *filters.IsLiked {
				orderByCombined = append(orderByCombined, sql.Desc(golikesT.C(generationoutputlike.FieldCreatedAt)))
			} else {
				orderByCombined = append(orderByCombined, sql.Asc(golikesT.C(generationoutputlike.FieldCreatedAt)))
			}
		}

		for _, o := range orderByGeneration {
			if orderDir == "desc" {
				orderByCombined = append(orderByCombined, sql.Desc(gt.C(o)))
			} else {
				orderByCombined = append(orderByCombined, sql.Asc(gt.C(o)))
			}
		}
		for _, o := range orderByOutput {
			if orderDir == "desc" {
				orderByCombined = append(orderByCombined, sql.Desc(got.C(o)))
			} else {
				orderByCombined = append(orderByCombined, sql.Asc(got.C(o)))
			}
		}
		// Order by generation, then output
		s.OrderBy(orderByCombined...)
	}).Scan(r.Ctx, &gQueryResult)

	if err != nil {
		log.Error("Error getting user generations", "err", err)
		return nil, err
	}

	if len(gQueryResult) == 0 {
		meta := &GenerationQueryWithOutputsMeta[*time.Time]{
			Outputs: []GenerationQueryWithOutputsResultFormatted{},
		}
		// Only give total if we have no cursor
		if cursor == nil {
			zero := 0
			meta.Total = &zero
		}
		return meta, nil
	}

	meta := &GenerationQueryWithOutputsMeta[*time.Time]{}
	if len(gQueryResult) > per_page {
		// Remove last item
		gQueryResult = gQueryResult[:len(gQueryResult)-1]
		meta.Next = &gQueryResult[len(gQueryResult)-1].CreatedAt
	}

	// Get real image URLs for each
	for i, g := range gQueryResult {
		if g.ImageUrl != "" {
			parsed := utils.GetEnv().GetURLFromImagePath(g.ImageUrl)
			gQueryResult[i].ImageUrl = parsed
		}
		if g.UpscaledImageUrl != "" {
			parsed := utils.GetEnv().GetURLFromImagePath(g.UpscaledImageUrl)
			gQueryResult[i].UpscaledImageUrl = parsed
		}
	}

	// Get prompt text
	promptIDsMap := make(map[uuid.UUID]string)
	negativePromptIdsMap := make(map[uuid.UUID]string)
	for _, g := range gQueryResult {
		if g.PromptID != nil {
			promptIDsMap[*g.PromptID] = ""
		}
		if g.NegativePromptID != nil {
			negativePromptIdsMap[*g.NegativePromptID] = ""
		}
	}
	promptIDs := make([]uuid.UUID, len(promptIDsMap))
	negativePromptId := make([]uuid.UUID, len(negativePromptIdsMap))

	i := 0
	for k := range promptIDsMap {
		promptIDs[i] = k
		i++
	}
	i = 0
	for k := range negativePromptIdsMap {
		negativePromptId[i] = k
		i++
	}

	prompts, err := r.DB.Prompt.Query().Select(prompt.FieldText).Where(prompt.IDIn(promptIDs...)).All(r.Ctx)
	if err != nil {
		log.Error("Error retrieving prompts", "err", err)
		return nil, err
	}

	negativePrompts, err := r.DB.NegativePrompt.Query().Select(negativeprompt.FieldText).Where(negativeprompt.IDIn(negativePromptId...)).All(r.Ctx)
	if err != nil {
		log.Error("Error retrieving prompts", "err", err)
		return nil, err
	}
	for _, p := range prompts {
		promptIDsMap[p.ID] = p.Text
	}
	for _, p := range negativePrompts {
		negativePromptIdsMap[p.ID] = p.Text
	}

	// Figure out liked by
	likedByMap := make(map[uuid.UUID]struct{})
	if filters.UserID != nil && len(gQueryResult) > 0 {
		outputIds := make([]uuid.UUID, len(gQueryResult))
		for i, g := range gQueryResult {
			if g.OutputID != nil {
				outputIds[i] = *g.OutputID
			}
		}
		likedByMap, err = r.GetGenerationOutputsLikedByUser(*filters.UserID, outputIds)
		if err != nil {
			log.Error("Error getting liked by map", "err", err)
			return nil, err
		}
	}

	// Format to GenerationQueryWithOutputsResultFormatted
	generationOutputMap := make(map[uuid.UUID][]GenerationUpscaleOutput)
	for _, g := range gQueryResult {
		if g.OutputID == nil {
			log.Warn("Output ID is nil for generation, cannot include in result", "id", g.ID)
			continue
		}
		likedByUser := false
		if _, ok := likedByMap[*g.OutputID]; ok {
			likedByUser = true
		}
		gOutput := GenerationUpscaleOutput{
			ID:               *g.OutputID,
			ImageUrl:         g.ImageUrl,
			UpscaledImageUrl: g.UpscaledImageUrl,
			GalleryStatus:    g.GalleryStatus,
			WasAutoSubmitted: g.WasAutoSubmitted,
			IsFavorited:      g.IsFavorited,
			IsPublic:         g.IsPublic,
			LikeCount:        utils.ToPtr(g.LikeCount),
			IsLiked:          utils.ToPtr(likedByUser),
		}
		promptText, _ := promptIDsMap[*g.PromptID]
		output := GenerationQueryWithOutputsResultFormatted{
			GenerationUpscaleOutput: gOutput,
			Generation: GenerationQueryWithOutputsData{
				ID:               g.ID,
				Height:           g.Height,
				Width:            g.Width,
				InferenceSteps:   g.InferenceSteps,
				Seed:             g.Seed,
				Status:           g.Status,
				GuidanceScale:    g.GuidanceScale,
				SchedulerID:      g.SchedulerID,
				ModelID:          g.ModelID,
				PromptID:         g.PromptID,
				NegativePromptID: g.NegativePromptID,
				CreatedAt:        g.CreatedAt,
				UpdatedAt:        g.UpdatedAt,
				StartedAt:        g.StartedAt,
				CompletedAt:      g.CompletedAt,
				InitImageURL:     g.InitImageURL,
				Prompt: PromptType{
					Text: promptText,
					ID:   *g.PromptID,
				},
				User: &UserType{
					Username:   g.Username,
					Identifier: utils.Sha256(g.UserID.String()),
				},
				IsFavorited:    gOutput.IsFavorited,
				IsPublic:       gOutput.IsPublic,
				PromptStrength: g.PromptStrength,
			},
		}
		if g.NegativePromptID != nil {
			negativePromptText, _ := negativePromptIdsMap[*g.NegativePromptID]
			output.Generation.NegativePrompt = &PromptType{
				Text: negativePromptText,
				ID:   *g.NegativePromptID,
			}
		}
		generationOutputMap[g.ID] = append(generationOutputMap[g.ID], gOutput)
		meta.Outputs = append(meta.Outputs, output)
	}
	// Now loop through and add outputs to each generation
	for i, g := range meta.Outputs {
		meta.Outputs[i].Generation.Outputs = generationOutputMap[g.Generation.ID]
	}

	if cursor == nil {
		total, err := r.GetGenerationCount(filters)
		if err != nil {
			log.Error("Error getting user generation count", "err", err)
			return nil, err
		}
		meta.Total = &total
	}

	return meta, err
}

// Separate count function
var cacheIsUpdating bool = false

type CachedCount struct {
	Count    int       `json:"count"`
	CachedAt time.Time `json:"cached_at"`
}

func (r *Repository) UpdateGenerationCountCacheAdmin(filters *requests.QueryGenerationFilters) (int, error) {
	queryG := r.DB.Generation.Query().Select(generation.FieldID).Where(
		generation.StatusEQ(generation.StatusSucceeded),
	)
	queryG = r.ApplyUserGenerationsFilters(queryG, filters, false)
	queryG = queryG.Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	})
	total, err := queryG.Modify(func(s *sql.Selector) {
		got := sql.Table(generationoutput.Table).As("t1")
		s.LeftJoin(got).On(s.C(generation.FieldID), got.C(generationoutput.FieldGenerationID))
	}).Count(r.Ctx)

	// Set in cache
	marshalled, err := json.Marshal(filters)
	if err != nil {
		log.Error("Error marshalling filters", "err", err)
	}
	hash := utils.Sha256(string(marshalled))
	cachedCount := CachedCount{
		Count:    total,
		CachedAt: time.Now(),
	}
	marshalledCount, err := json.Marshal(cachedCount)
	if err != nil {
		log.Error("Error marshalling cached count", "err", err)
	}
	err = r.Redis.Client.Set(r.Ctx, "generation_count_"+hash, marshalledCount, 0).Err()
	if err != nil {
		log.Error("Error setting cached count", "err", err)
	}

	return total, err
}

func (r *Repository) GetGenerationCountAdmin(filters *requests.QueryGenerationFilters) (int, error) {
	// Get has of filters to see if we have this in cache
	marshalled, err := json.Marshal(filters)
	if err != nil {
		log.Error("Error marshalling filters", "err", err)
	}
	hash := utils.Sha256(string(marshalled))
	// Check cache
	var cachedCount CachedCount
	valStr, err := r.Redis.Client.Get(r.Ctx, "generation_count_"+hash).Result()
	if err != nil {
		log.Error("Error getting cached count", "err", err)
		return r.UpdateGenerationCountCacheAdmin(filters)
	}
	err = json.Unmarshal([]byte(valStr), &cachedCount)
	if err != nil {
		log.Error("Error unmarshalling cached count", "err", err)
		return r.UpdateGenerationCountCacheAdmin(filters)
	}
	// See if needs updating
	// If older than 30 minutes, refresh cache
	if cachedCount.CachedAt.Before(time.Now().Add(-30 * time.Minute)) {
		if !cacheIsUpdating {
			go func() {
				cacheIsUpdating = true
				r.UpdateGenerationCountCacheAdmin(filters)
				cacheIsUpdating = false
			}()
		}
	}
	return cachedCount.Count, nil
}

// Alternate version for performance when we can't index by user_id
func (r *Repository) QueryGenerationsAdmin(per_page int, cursor *time.Time, callingUserId *uuid.UUID, filters *requests.QueryGenerationFilters) (*GenerationQueryWithOutputsMeta[*time.Time], error) {
	var gQueryResult []GenerationQueryWithOutputsResult

	// Figure out order bys
	var orderByGeneration []string
	var orderByOutput []string
	if filters == nil || (filters != nil && filters.OrderBy == requests.OrderByCreatedAt) {
		orderByGeneration = []string{generation.FieldCreatedAt}
		orderByOutput = []string{generationoutput.FieldCreatedAt}
	} else {
		orderByGeneration = []string{generation.FieldCreatedAt, generation.FieldUpdatedAt}
		orderByOutput = []string{generationoutput.FieldCreatedAt, generationoutput.FieldUpdatedAt}
	}

	// Initial query to get output IDs
	var rawQ []struct {
		ID                uuid.UUID  `json:"id" sql:"id"`
		OutputID          uuid.UUID  `json:"output_id" sql:"output_id"`
		LikeCount         int        `json:"like_count" sql:"like_count"`
		OutputCreatedAt   time.Time  `json:"output_created_at" sql:"output_created_at"`
		DeletedAt         *time.Time `json:"deleted_at" sql:"deleted_at"`
		UpscaledImagePath *string    `json:"upscaled_image_path" sql:"upscaled_image_path"`
		GalleryStatus     string     `json:"gallery_status" sql:"gallery_status"`
	}

	queryG := r.DB.Generation.Query().Select(generation.FieldID).Where(
		generation.StatusEQ(generation.StatusSucceeded),
	)
	if filters != nil && len(filters.Username) > 0 {
		queryG = queryG.Where(func(s *sql.Selector) {
			uT := sql.Table(user.Table)
			v := make([]any, len(filters.Username))
			for i := range v {
				v[i] = filters.Username[i]
			}
			s.Where(sql.In(s.C(generation.FieldUserID), sql.Select(user.FieldID).From(uT).Where(sql.In(sql.Lower(user.FieldUsername), v...))))

		})
	}
	queryG = r.ApplyUserGenerationsFilters(queryG, filters, false)
	queryG = queryG.Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
	})
	err := queryG.Limit(per_page+1).Modify(func(s *sql.Selector) {
		got := sql.Table(generationoutput.Table).As("t1")
		s.LeftJoin(got).On(s.C(generation.FieldID), got.C(generationoutput.FieldGenerationID))
		s.AppendSelect(sql.As(got.C(generationoutput.FieldID), "output_id"), sql.As(got.C(generationoutput.FieldLikeCount), "like_count"), sql.As(got.C(generationoutput.FieldCreatedAt), "output_created_at"), sql.As(got.C(generationoutput.FieldDeletedAt), "deleted_at"),
			sql.As(got.C(generationoutput.FieldUpscaledImagePath), "upscaled_image_path"), sql.As(got.C(generationoutput.FieldGalleryStatus), "gallery_status"))
		orderDir := "asc"
		if filters == nil || (filters != nil && filters.Order == requests.SortOrderDescending) {
			orderDir = "desc"
		}
		var orderByGeneration2 []string
		var orderByOutput2 []string
		for _, o := range orderByGeneration {
			if orderDir == "desc" {
				orderByGeneration2 = append(orderByGeneration2, sql.Desc(s.C(o)))
			} else {
				orderByGeneration2 = append(orderByGeneration2, sql.Asc(s.C(o)))
			}
		}
		for _, o := range orderByOutput {
			if orderDir == "desc" {
				orderByOutput2 = append(orderByOutput2, sql.Desc(got.C(o)))
			} else {
				orderByOutput2 = append(orderByOutput2, sql.Asc(got.C(o)))
			}
		}
		// Order by generation, then output
		orderByCombined := append(orderByGeneration2, orderByOutput2...)
		s.OrderBy(orderByCombined...)
	}).Scan(r.Ctx, &rawQ)
	if err != nil {
		log.Error("Error querying generations", "err", err)
		return nil, err
	}

	outputIDs := make([]uuid.UUID, len(rawQ))
	for i, r := range rawQ {
		outputIDs[i] = r.OutputID
	}

	query := r.DB.GenerationOutput.Query().Where(generationoutput.IDIn(outputIDs...)).WithGenerations(func(s *ent.GenerationQuery) {
		s.WithPrompt()
		s.WithNegativePrompt()
		s.WithUser()
		s.WithGenerationOutputs(func(goq *ent.GenerationOutputQuery) {
			if filters == nil || (filters != nil && filters.Order == requests.SortOrderDescending) {
				goq = goq.Order(ent.Desc(orderByOutput...))
			} else {
				goq = goq.Order(ent.Asc(orderByOutput...))
			}
		})
		if filters == nil || (filters != nil && filters.Order == requests.SortOrderDescending) {
			s = s.Order(ent.Desc(orderByGeneration...))
		} else {
			s = s.Order(ent.Asc(orderByGeneration...))
		}
	})

	if filters == nil || (filters != nil && filters.Order == requests.SortOrderDescending) {
		query = query.Order(ent.Desc(orderByOutput...))
	} else {
		query = query.Order(ent.Asc(orderByOutput...))
	}

	// Limit
	query = query.Limit(per_page + 1)

	res, err := query.All(r.Ctx)

	if err != nil {
		log.Error("Error getting admin generations", "err", err)
		return nil, err
	}

	meta := &GenerationQueryWithOutputsMeta[*time.Time]{}

	if len(res) == 0 {
		meta := &GenerationQueryWithOutputsMeta[*time.Time]{
			Outputs: []GenerationQueryWithOutputsResultFormatted{},
		}
		// Only give total if we have no cursor
		if cursor == nil {
			zero := 0
			meta.Total = &zero
		}
		return meta, nil
	}

	if len(res) > per_page {
		// Remove last item
		res = res[:len(res)-1]
		meta.Next = &res[len(res)-1].CreatedAt
	}

	// Figure out liked by in another query, if calling user is provided
	likedByMap := make(map[uuid.UUID]struct{})
	if callingUserId != nil && len(outputIDs) > 0 {
		outputIds := make([]uuid.UUID, len(outputIDs))
		for i, g := range outputIDs {
			outputIds[i] = g
		}
		likedByMap, err = r.GetGenerationOutputsLikedByUser(*callingUserId, outputIds)
		if err != nil {
			log.Error("Error getting liked by map", "err", err)
			return nil, err
		}
	}

	// Get real image URLs for each
	// Since we do this weird format need to cap at per page
	for _, g := range res {
		// root generation data
		generationRoot := GenerationQueryWithOutputsData{
			ID:               g.ID,
			Width:            g.Edges.Generations.Width,
			Height:           g.Edges.Generations.Height,
			InferenceSteps:   g.Edges.Generations.InferenceSteps,
			Seed:             g.Edges.Generations.Seed,
			Status:           string(g.Edges.Generations.Status),
			GuidanceScale:    g.Edges.Generations.GuidanceScale,
			SchedulerID:      g.Edges.Generations.SchedulerID,
			ModelID:          g.Edges.Generations.ModelID,
			PromptID:         g.Edges.Generations.PromptID,
			NegativePromptID: g.Edges.Generations.NegativePromptID,
			CreatedAt:        g.Edges.Generations.CreatedAt,
			UpdatedAt:        g.Edges.Generations.UpdatedAt,
			StartedAt:        g.Edges.Generations.StartedAt,
			CompletedAt:      g.Edges.Generations.CompletedAt,
			WasAutoSubmitted: g.Edges.Generations.WasAutoSubmitted,
			IsFavorited:      g.IsFavorited,
			IsPublic:         g.IsPublic,
		}
		if g.Edges.Generations.InitImageURL != nil {
			generationRoot.InitImageURL = *g.Edges.Generations.InitImageURL
			generationRoot.PromptStrength = g.Edges.Generations.PromptStrength
		}
		if g.Edges.Generations.Edges.NegativePrompt != nil {
			generationRoot.NegativePrompt = &PromptType{
				Text: g.Edges.Generations.Edges.NegativePrompt.Text,
				ID:   *g.Edges.Generations.NegativePromptID,
			}
		}
		if g.Edges.Generations.Edges.Prompt != nil {
			generationRoot.Prompt = PromptType{
				Text: g.Edges.Generations.Edges.Prompt.Text,
				ID:   *g.Edges.Generations.PromptID,
			}
		}
		if g.Edges.Generations.Edges.User != nil {
			generationRoot.User = &UserType{
				Email:      g.Edges.Generations.Edges.User.Email,
				Username:   g.Edges.Generations.Edges.User.Username,
				Identifier: utils.Sha256(g.Edges.Generations.Edges.User.ID.String()),
			}
		}

		// Add outputs
		for _, o := range g.Edges.Generations.Edges.GenerationOutputs {
			output := GenerationUpscaleOutput{
				ID:                     o.ID,
				ImageUrl:               utils.GetEnv().GetURLFromImagePath(o.ImagePath),
				GalleryStatus:          o.GalleryStatus,
				CreatedAt:              &o.CreatedAt,
				IsFavorited:            o.IsFavorited,
				IsPublic:               o.IsPublic,
				WasAutoSubmitted:       generationRoot.WasAutoSubmitted,
				AestheticRatingScore:   utils.ToPtr(o.AestheticRatingScore),
				AestheticArtifactScore: utils.ToPtr((o.AestheticArtifactScore)),
			}
			if o.UpscaledImagePath != nil {
				output.UpscaledImageUrl = utils.GetEnv().GetURLFromImagePath(*o.UpscaledImagePath)
			}
			generationRoot.Outputs = append(generationRoot.Outputs, output)
		}

		likedByUser := false
		if _, ok := likedByMap[g.ID]; ok {
			likedByUser = true
		}

		ret := GenerationQueryWithOutputsResult{
			LikeCount:                      g.LikeCount,
			IsLiked:                        utils.ToPtr(likedByUser),
			OutputID:                       &g.ID,
			ImageUrl:                       utils.GetEnv().GetURLFromImagePath(g.ImagePath),
			GalleryStatus:                  g.GalleryStatus,
			DeletedAt:                      g.DeletedAt,
			GenerationQueryWithOutputsData: generationRoot,
			AestheticArtifactScore:         utils.ToPtr(g.AestheticArtifactScore),
			AestheticRatingScore:           utils.ToPtr(g.AestheticRatingScore),
		}
		if g.UpscaledImagePath != nil {
			ret.UpscaledImageUrl = utils.GetEnv().GetURLFromImagePath(*g.UpscaledImagePath)
		}

		gQueryResult = append(gQueryResult, ret)
	}

	// Format to GenerationQueryWithOutputsResultFormatted
	for _, g := range gQueryResult {
		if g.OutputID == nil {
			log.Warn("Output ID is nil for generation, cannot include in result", "id", g.ID)
			continue
		}
		likedByUser := false
		if _, ok := likedByMap[g.ID]; ok {
			likedByUser = true
		}

		gOutput := GenerationUpscaleOutput{
			ID:                     *g.OutputID,
			ImageUrl:               g.ImageUrl,
			UpscaledImageUrl:       g.UpscaledImageUrl,
			GalleryStatus:          g.GalleryStatus,
			LikeCount:              utils.ToPtr(g.LikeCount),
			IsLiked:                utils.ToPtr(likedByUser),
			AestheticRatingScore:   g.AestheticRatingScore,
			AestheticArtifactScore: g.AestheticArtifactScore,
		}
		for _, o := range g.Outputs {
			if o.ID == *g.OutputID {
				gOutput.CreatedAt = o.CreatedAt
				break
			}
		}
		output := GenerationQueryWithOutputsResultFormatted{
			GenerationUpscaleOutput: gOutput,
			Generation:              g.GenerationQueryWithOutputsData,
		}

		meta.Outputs = append(meta.Outputs, output)
	}

	// Get count when no cursor
	if cursor == nil {
		total, err := r.GetGenerationCountAdmin(filters)
		if err != nil {
			log.Error("Error getting user generation count", "err", err)
			return nil, err
		}
		meta.Total = &total
	}

	return meta, err
}

type GenerationUpscaleOutput struct {
	ID                     uuid.UUID                      `json:"id"`
	ImageUrl               string                         `json:"image_url,omitempty"`
	AudioFileUrl           string                         `json:"audio_file_url,omitempty"`
	VideoFileUrl           string                         `json:"video_file_url,omitempty"`
	AudioArray             []float64                      `json:"audio_array,omitempty"`
	AudioDuration          *float32                       `json:"audio_duration,omitempty"`
	UpscaledImageUrl       string                         `json:"upscaled_image_url,omitempty"`
	GalleryStatus          generationoutput.GalleryStatus `json:"gallery_status,omitempty"`
	OutputID               *uuid.UUID                     `json:"output_id,omitempty"`
	CreatedAt              *time.Time                     `json:"created_at,omitempty"`
	IsFavorited            bool                           `json:"is_favorited"`
	InitImageUrl           string                         `json:"init_image_url,omitempty"`
	WasAutoSubmitted       bool                           `json:"was_auto_submitted"`
	IsPublic               bool                           `json:"is_public"`
	DenoiseAudio           *bool                          `json:"denoise_audio,omitempty"`
	RemoveSilence          *bool                          `json:"remove_silence,omitempty"`
	LikeCount              *int                           `json:"like_count,omitempty"`
	IsLiked                *bool                          `json:"is_liked,omitempty"`
	AestheticRatingScore   *float32                       `json:"aesthetic_rating_score,omitempty"`
	AestheticArtifactScore *float32                       `json:"aesthetic_artifact_score,omitempty"`
}

type GenerationQueryWithOutputsMetaCursor interface {
	*uint | *time.Time
}

// Paginated meta for querying generations
type GenerationQueryWithOutputsMeta[T GenerationQueryWithOutputsMetaCursor] struct {
	Total   *int                                        `json:"total_count,omitempty"`
	Outputs []GenerationQueryWithOutputsResultFormatted `json:"outputs"`
	Next    T                                           `json:"next,omitempty"`
}

type PromptType struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type UserType struct {
	Email      string `json:"email,omitempty"`
	Username   string `json:"username"`
	Identifier string `json:"identifier,omitempty"`
}

type GenerationQueryWithOutputsData struct {
	ID                     uuid.UUID                 `json:"id" sql:"id"`
	Height                 int32                     `json:"height" sql:"height"`
	Width                  int32                     `json:"width" sql:"width"`
	InferenceSteps         int32                     `json:"inference_steps" sql:"inference_steps"`
	Seed                   int                       `json:"seed" sql:"seed"`
	Status                 string                    `json:"status" sql:"status"`
	GuidanceScale          float32                   `json:"guidance_scale" sql:"guidance_scale"`
	SchedulerID            uuid.UUID                 `json:"scheduler_id" sql:"scheduler_id"`
	ModelID                uuid.UUID                 `json:"model_id" sql:"model_id"`
	PromptID               *uuid.UUID                `json:"prompt_id,omitempty" sql:"prompt_id"`
	NegativePromptID       *uuid.UUID                `json:"negative_prompt_id,omitempty" sql:"negative_prompt_id"`
	CreatedAt              time.Time                 `json:"created_at" sql:"created_at"`
	UpdatedAt              time.Time                 `json:"updated_at" sql:"updated_at"`
	StartedAt              *time.Time                `json:"started_at,omitempty" sql:"started_at"`
	CompletedAt            *time.Time                `json:"completed_at,omitempty" sql:"completed_at"`
	Username               string                    `json:"username,omitempty" sql:"username"`
	UserID                 uuid.UUID                 `json:"user_id,omitempty" sql:"user_id"`
	IsFavorited            bool                      `json:"is_favorited" sql:"is_favorited"`
	IsPublic               bool                      `json:"is_public" sql:"is_public"`
	Outputs                []GenerationUpscaleOutput `json:"outputs"`
	Prompt                 PromptType                `json:"prompt"`
	NegativePrompt         *PromptType               `json:"negative_prompt,omitempty"`
	WasAutoSubmitted       bool                      `json:"was_auto_submitted" sql:"was_auto_submitted"`
	InitImageURL           string                    `json:"init_image_url,omitempty" sql:"init_image_url"`
	InitImageURLSigned     string                    `json:"init_image_url_signed,omitempty"`
	User                   *UserType                 `json:"user,omitempty"`
	PromptStrength         *float32                  `json:"prompt_strength,omitempty" sql:"prompt_strength"`
	AestheticRatingScore   *float32                  `json:"aesthetic_rating_score,omitempty"`
	AestheticArtifactScore *float32                  `json:"aesthetic_artifact_score,omitempty"`
}

type GenerationQueryWithOutputsResult struct {
	OutputID               *uuid.UUID                     `json:"output_id,omitempty" sql:"output_id"`
	IsLiked                *bool                          `json:"is_liked,omitempty" sql:"is_liked"`
	LikedAt                *time.Time                     `json:"liked_at,omitempty" sql:"liked_at"`
	AestheticArtifactScore *float32                       `json:"aesthetic_artifact_score,omitempty"`
	AestheticRatingScore   *float32                       `json:"aesthetic_rating_score,omitempty"`
	LikeCount              int                            `json:"like_count" sql:"like_count"`
	LikeCountTrending      *int                           `json:"like_count_trending,omitempty"`
	ImageUrl               string                         `json:"image_url,omitempty" sql:"image_path"`
	UpscaledImageUrl       string                         `json:"upscaled_image_url,omitempty" sql:"upscaled_image_path"`
	GalleryStatus          generationoutput.GalleryStatus `json:"gallery_status,omitempty" sql:"output_gallery_status"`
	DeletedAt              *time.Time                     `json:"deleted_at,omitempty" sql:"deleted_at"`
	GenerationQueryWithOutputsData
}

type GenerationQueryWithOutputsResultFormatted struct {
	GenerationUpscaleOutput
	Generation GenerationQueryWithOutputsData `json:"generation"`
}

func (r *Repository) GetGenerationsQueuedOrStarted() ([]*ent.Generation, error) {
	// Get generations that are started/queued and older than 5 minutes
	return r.DB.Generation.Query().
		Where(
			generation.StatusIn(
				generation.StatusQueued,
				generation.StatusStarted,
			),
			generation.CreatedAtLT(time.Now().Add(-5*time.Minute)),
		).
		Order(ent.Desc(generation.FieldCreatedAt)).
		Limit(100).
		All(r.Ctx)
}
