package repository

import (
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/deviceinfo"
	"github.com/stablecog/sc-go/database/ent/negativeprompt"
	"github.com/stablecog/sc-go/database/ent/prompt"
)

// Create device info and prompts if they don't exist, otherwise get existing
func (r *Repository) GetOrCreatePrompts(promptText, translatedPromptText string, negativePromptText string, translatedNegativePromptText string, promptType prompt.Type, DB *ent.Client) (promptId *uuid.UUID, negativePromptId *uuid.UUID, err error) {
	if DB == nil {
		DB = r.DB
	}
	// Check if prompt exists
	var dbPrompt *ent.Prompt
	dbPrompt, err = DB.Prompt.Query().Where(prompt.TextEQ(promptText), prompt.TypeEQ(promptType)).First(r.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			query := DB.Prompt.Create().SetText(promptText).SetRanTranslation(true).SetType(promptType)
			if promptText != translatedPromptText {
				query.SetTranslatedText(translatedPromptText)
			}
			dbPrompt, err = query.Save(r.Ctx)
			if err != nil {
				return promptId, negativePromptId, err
			}
		} else {
			return promptId, negativePromptId, err
		}
	}

	// Check if negative prompt exists
	var dbNegativePrompt *ent.NegativePrompt
	if negativePromptText != "" {
		dbNegativePrompt, err = DB.NegativePrompt.Query().Where(negativeprompt.TextEQ(negativePromptText)).First(r.Ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				query := DB.NegativePrompt.Create().SetText(negativePromptText).SetRanTranslation(true)
				if negativePromptText != translatedNegativePromptText {
					query.SetTranslatedText(translatedNegativePromptText)
				}
				dbNegativePrompt, err = query.Save(r.Ctx)
				if err != nil {
					return promptId, negativePromptId, err
				}
			} else {
				return promptId, negativePromptId, err
			}
		}
	}

	// Negative prompt is optional so
	if dbPrompt != nil {
		promptId = &dbPrompt.ID
	}
	if dbNegativePrompt != nil {
		negativePromptId = &dbNegativePrompt.ID
	}
	return promptId, negativePromptId, nil
}

// Get a device_info ID given inputs
func (r *Repository) GetOrCreateDeviceInfo(deviceType, deviceOs, deviceBrowser string, DB *ent.Client) (deviceInfoId uuid.UUID, err error) {
	if DB == nil {
		DB = r.DB
	}
	// Check if device info combo exists
	var dbDeviceInfo *ent.DeviceInfo
	dbDeviceInfo, err = DB.DeviceInfo.Query().Where(deviceinfo.Type(deviceType), deviceinfo.Os(deviceOs), deviceinfo.Browser(deviceBrowser)).First(r.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// Create device info
			dbDeviceInfo, err = DB.DeviceInfo.Create().SetType(deviceType).SetOs(deviceOs).SetBrowser(deviceBrowser).Save(r.Ctx)
			if err != nil {
				return deviceInfoId, err
			}
		} else {
			return deviceInfoId, err
		}
	}

	return dbDeviceInfo.ID, nil
}
