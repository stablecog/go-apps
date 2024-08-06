package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/prompt"
	"github.com/stablecog/sc-go/database/enttypes"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/utils"
)

// CreateGeneration creates the initial generation in the database
// Takes in a userID (creator),  device info, countryCode, and a request body
func (r *Repository) CreateGeneration(userID uuid.UUID, deviceType, deviceOs, deviceBrowser, countryCode string, req requests.CreateGenerationRequest, productId *string, apiTokenId *uuid.UUID, sourceType enttypes.SourceType, DB *ent.Client) (*ent.Generation, error) {
	if DB == nil {
		DB = r.DB
	}
	// Get prompt, negative prompt, device info
	deviceInfoId, err := r.GetOrCreateDeviceInfo(deviceType, deviceOs, deviceBrowser, DB)
	if err != nil {
		return nil, err
	}
	insert := DB.Generation.Create().
		SetStatus(generation.StatusQueued).
		SetWidth(*req.Width).
		SetHeight(*req.Height).
		SetGuidanceScale(*req.GuidanceScale).
		SetInferenceSteps(*req.InferenceSteps).
		SetSeed(*req.Seed).
		SetModelID(*req.ModelId).
		SetSchedulerID(*req.SchedulerId).
		SetDeviceInfoID(deviceInfoId).
		SetCountryCode(countryCode).
		SetUserID(userID).
		SetWasAutoSubmitted(req.WasAutoSubmitted).
		SetNumOutputs(*req.NumOutputs).
		SetSourceType(sourceType).
		SetWebhookToken(uuid.New())
	if productId != nil {
		insert.SetStripeProductID(*productId)
	}
	if req.InitImageUrl != "" {
		insert.SetInitImageURL(req.InitImageUrl)
	}
	if req.MaskImageUrl != "" {
		insert.SetMaskImageURL(req.MaskImageUrl)
	}
	if req.PromptStrength != nil {
		insert.SetPromptStrength(*req.PromptStrength)
	}
	if apiTokenId != nil {
		insert.SetAPITokenID(*apiTokenId)
	}
	return insert.Save(r.Ctx)
}

func (r *Repository) SetGenerationStarted(generationID string) error {
	uid, err := uuid.Parse(generationID)
	if err != nil {
		log.Error("Error parsing generation id in SetGenerationStarted", "id", generationID, "err", err)
		return err
	}
	_, err = r.DB.Generation.Update().Where(generation.IDEQ(uid), generation.StatusEQ(generation.StatusQueued)).SetStatus(generation.StatusStarted).SetStartedAt(time.Now()).Save(r.Ctx)
	if err != nil {
		// Log error here since this might be happening in a goroutine
		log.Error("Error setting generation started", "id", generationID, "err", err)
	}
	return err
}

func (r *Repository) SetGenerationFailed(generationID string, reason string, nsfwCount int32, db *ent.Client) error {
	if db == nil {
		db = r.DB
	}

	uid, err := uuid.Parse(generationID)
	if err != nil {
		log.Error("Error parsing generation id in SetGenerationFailed", "id", generationID, "err", err)
		return err
	}
	_, err = db.Generation.UpdateOneID(uid).SetStatus(generation.StatusFailed).SetFailureReason(reason).SetNsfwCount(nsfwCount).SetCompletedAt(time.Now()).SetWebhookToken(uuid.New()).Save(r.Ctx)
	if err != nil {
		log.Error("Error setting generation failed", "id", generationID, "err", err)
	}
	return err
}

func (r *Repository) SetGenerationSucceeded(generationID string, promptStr string, negativePrompt string, submitToGallery bool, whOutput requests.CogWebhookOutput, nsfwCount int32) ([]*ent.GenerationOutput, error) {
	uid, err := uuid.Parse(generationID)
	if err != nil {
		log.Error("Error parsing generation id in SetGenerationSucceeded", "id", generationID, "err", err)
		return nil, err
	}

	var outputRet []*ent.GenerationOutput

	// Wrap in transaction
	if err := r.WithTx(func(tx *ent.Tx) error {
		if err != nil {
			log.Error("Error starting transaction in SetGenerationSucceeded", "id", generationID, "err", err)
			return err
		}
		db := tx.Client()

		// Get prompt IDs
		promptId, negativePromptId, err := r.GetOrCreatePrompts(promptStr, negativePrompt, prompt.TypeImage, db)
		if err != nil {
			log.Error("Error getting or creating prompts", "id", generationID, "err", err, "prompt", promptStr, "negativePrompt", negativePrompt)
			return err
		}

		// Update the generation
		genUpdate := db.Generation.UpdateOneID(uid).SetStatus(generation.StatusSucceeded).SetCompletedAt(time.Now()).SetWebhookToken(uuid.New()).SetNsfwCount(nsfwCount)
		if promptId != nil {
			genUpdate.SetPromptID(*promptId)
		}
		if negativePromptId != nil {
			genUpdate.SetNegativePromptID(*negativePromptId)
		}
		generation, err := genUpdate.Save(r.Ctx)
		if err != nil {
			log.Error("Error setting generation succeeded", "id", generationID, "err", err)
			return err
		}

		// If this generation was created with "submit_to_gallery", then submit all outputs to gallery
		var galleryStatus generationoutput.GalleryStatus
		isPublic := false
		if submitToGallery {
			galleryStatus = generationoutput.GalleryStatusSubmitted
			isPublic = true
		} else {
			galleryStatus = generationoutput.GalleryStatusNotSubmitted
		}

		// Insert all generation outputs
		for _, output := range whOutput.Images {
			parsedS3, err := utils.GetPathFromS3URL(output.Image)
			if err != nil {
				log.Error("Error parsing s3 url", "output", output, "err", err)
				parsedS3 = output.Image
			}
			gOutput, err := db.GenerationOutput.Create().
				SetGenerationID(uid).
				SetImagePath(parsedS3).
				SetGalleryStatus(galleryStatus).
				SetHasEmbeddings(true).
				SetIsPublic(isPublic).
				SetAestheticArtifactScore(output.AestheticArtifactScore).
				SetAestheticRatingScore(output.AestheticRatingScore).
				Save(r.Ctx)
			if err != nil {
				log.Error("Error inserting generation output", "id", generationID, "err", err)
				return err
			}
			outputRet = append(outputRet, gOutput)
			if r.Qdrant != nil {
				payload := map[string]interface{}{
					"image_path":               gOutput.ImagePath,
					"gallery_status":           gOutput.GalleryStatus,
					"is_favorited":             gOutput.IsFavorited,
					"created_at":               gOutput.CreatedAt.Unix(),
					"updated_at":               gOutput.UpdatedAt.Unix(),
					"was_auto_submitted":       generation.WasAutoSubmitted,
					"guidance_scale":           generation.GuidanceScale,
					"inference_steps":          generation.InferenceSteps,
					"prompt_strength":          generation.PromptStrength,
					"height":                   generation.Height,
					"width":                    generation.Width,
					"model":                    generation.ModelID.String(),
					"scheduler":                generation.SchedulerID.String(),
					"user_id":                  generation.UserID.String(),
					"generation_id":            generation.ID.String(),
					"prompt":                   promptStr,
					"prompt_id":                generation.PromptID.String(),
					"is_public":                isPublic,
					"aesthetic_rating_score":   gOutput.AestheticRatingScore,
					"aesthetic_artifact_score": gOutput.AestheticArtifactScore,
				}
				if gOutput.UpscaledImagePath != nil {
					payload["upscaled_image_path"] = *gOutput.UpscaledImagePath
				}
				if generation.InitImageURL != nil {
					payload["init_image_url"] = *generation.InitImageURL
				}
				if negativePrompt != "" {
					payload["negative_prompt"] = negativePrompt
				}
				err = r.Qdrant.Upsert(
					gOutput.ID,
					payload,
					output.ImageEmbed,
					false,
				)
				if err != nil {
					log.Error("Error upserting to qdrant", "id", generationID, "err", err)
					return err
				}
			} else {
				log.Warn("Qdrant client not initialized, not adding to qdrant")
			}
		}

		return nil
	}); err != nil {
		log.Error("Error starting transaction in SetGenerationSucceeded", "id", generationID, "err", err)
		return nil, err
	}

	return outputRet, nil
}

func GetEmbeddingsAndInsertIntoQdrant(r *Repository, generation *ent.Generation, outputs []*ent.GenerationOutput, generationID string, promptStr string, negativePrompt string, isPublic bool) (err error) {
	imageEmbedPlaceholder := []float32{0.0}
	// Transaction
	if err := r.WithTx(func(tx *ent.Tx) error {
		if err != nil {
			log.Error("Error starting transaction in GetEmbeddingsAndInsertIntoQdrant", "id", generationID, "err", err)
			return err
		}
		db := tx.Client()
		// Insert all generation outputs
		for _, output := range outputs {
			_, err := db.GenerationOutput.Update().
				Where(generationoutput.IDEQ(output.ID)).
				SetHasEmbeddings(true).
				SetAestheticArtifactScore(output.AestheticArtifactScore).
				SetAestheticRatingScore(output.AestheticRatingScore).
				Save(r.Ctx)
			if err != nil {
				log.Error("Error updating generation output", "id", output.ID, "err", err)
				return err
			}
			if r.Qdrant != nil {
				payload := map[string]interface{}{
					"image_path":               output.ImagePath,
					"gallery_status":           output.GalleryStatus,
					"is_favorited":             output.IsFavorited,
					"created_at":               output.CreatedAt.Unix(),
					"updated_at":               output.UpdatedAt.Unix(),
					"was_auto_submitted":       generation.WasAutoSubmitted,
					"guidance_scale":           generation.GuidanceScale,
					"inference_steps":          generation.InferenceSteps,
					"prompt_strength":          generation.PromptStrength,
					"height":                   generation.Height,
					"width":                    generation.Width,
					"model":                    generation.ModelID.String(),
					"scheduler":                generation.SchedulerID.String(),
					"user_id":                  generation.UserID.String(),
					"generation_id":            generation.ID.String(),
					"prompt":                   promptStr,
					"prompt_id":                generation.PromptID.String(),
					"is_public":                isPublic,
					"aesthetic_rating_score":   output.AestheticRatingScore,
					"aesthetic_artifact_score": output.AestheticArtifactScore,
				}
				if output.UpscaledImagePath != nil {
					payload["upscaled_image_path"] = *output.UpscaledImagePath
				}
				if generation.InitImageURL != nil {
					payload["init_image_url"] = *generation.InitImageURL
				}
				if negativePrompt != "" {
					payload["negative_prompt"] = negativePrompt
				}
				err = r.Qdrant.Upsert(
					output.ID,
					payload,
					imageEmbedPlaceholder,
					false,
				)
				if err != nil {
					log.Error("Error upserting to qdrant", "id", generationID, "err", err)
					return err
				}
			} else {
				log.Warn("Qdrant client not initialized, not adding to qdrant")
			}
		}
		return nil
	}); err != nil {
		log.Error("Error starting transaction in SetGenerationSucceeded", "id", generationID, "err", err)
		return err
	}
	return nil
}
