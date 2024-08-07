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

func (r *Repository) SetGenerationSucceeded(generationID string, promptStr string, translatedPrompt string, negativePrompt string, translatedNegativePrompt string, submitToGallery bool, whOutput requests.CogWebhookOutput, nsfwCount int32) ([]*ent.GenerationOutput, error) {
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
		promptId, negativePromptId, err := r.GetOrCreatePrompts(promptStr, translatedPrompt, negativePrompt, translatedNegativePrompt, prompt.TypeImage, db)
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
		_, err = genUpdate.Save(r.Ctx)
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
				SetHasEmbeddings(false).
				SetIsPublic(isPublic).
				SetAestheticArtifactScore(0).
				SetAestheticRatingScore(0).
				Save(r.Ctx)
			if err != nil {
				log.Error("Error inserting generation output", "id", generationID, "err", err)
				return err
			}
			outputRet = append(outputRet, gOutput)
		}

		return nil
	}); err != nil {
		log.Error("Error starting transaction in SetGenerationSucceeded", "id", generationID, "err", err)
		return nil, err
	}

	return outputRet, nil
}
