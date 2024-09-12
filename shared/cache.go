package shared

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"golang.org/x/exp/slices"
)

// A singleton that caches the features available to free users
// Avoids having to query the database every time a user requests the features
type Cache struct {
	// Models and options available to free users //
	generateModels       []*ent.GenerationModel
	upscaleModels        []*ent.UpscaleModel
	schedulers           []*ent.Scheduler
	voiceoverModels      []*ent.VoiceoverModel
	voiceoverSpeakers    []*ent.VoiceoverSpeaker
	adminIDs             []uuid.UUID
	iPBlacklist          []string
	thumbmarkIDBlacklist []string
	usernameBlacklist    []string
	bannedWords          []*ent.BannedWords
	clipUrls             []string
	httpClient           *http.Client
	sync.RWMutex
}

var lock = &sync.Mutex{}
var singleCache *Cache

func newCache() *Cache {
	return &Cache{
		httpClient: &http.Client{
			Timeout: time.Second * 30, // Set a timeout for all requests
		},
	}
}

func GetCache() *Cache {
	if singleCache == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleCache == nil {
			singleCache = newCache()
		}
	}
	return singleCache
}

func (f *Cache) UpdateGenerationModels(models []*ent.GenerationModel) {
	f.Lock()
	defer f.Unlock()
	f.generateModels = models
}

func (f *Cache) GenerationModels() []*ent.GenerationModel {
	f.RLock()
	defer f.RUnlock()
	return f.generateModels
}

func (f *Cache) UpdateUpscaleModels(models []*ent.UpscaleModel) {
	f.Lock()
	defer f.Unlock()
	f.upscaleModels = models
}

func (f *Cache) UpscaleModels() []*ent.UpscaleModel {
	f.RLock()
	defer f.RUnlock()
	return f.upscaleModels
}

func (f *Cache) UpdateSchedulers(schedulers []*ent.Scheduler) {
	f.Lock()
	defer f.Unlock()
	f.schedulers = schedulers
}

func (f *Cache) Schedulers() []*ent.Scheduler {
	f.RLock()
	defer f.RUnlock()
	return f.schedulers
}

func (f *Cache) UpdateVoiceoverModels(models []*ent.VoiceoverModel) {
	f.Lock()
	defer f.Unlock()
	f.voiceoverModels = models
}

func (f *Cache) VoiceoverModels() []*ent.VoiceoverModel {
	f.RLock()
	defer f.RUnlock()
	return f.voiceoverModels
}

func (f *Cache) UpdateVoiceoverSpeakers(speakers []*ent.VoiceoverSpeaker) {
	f.Lock()
	defer f.Unlock()
	f.voiceoverSpeakers = speakers
}

func (f *Cache) VoiceoverSpeakers() []*ent.VoiceoverSpeaker {
	f.RLock()
	defer f.RUnlock()
	return f.voiceoverSpeakers
}

func (f *Cache) UpdateBannedWords(bannedWords []*ent.BannedWords) {
	f.Lock()
	defer f.Unlock()
	f.bannedWords = bannedWords
}

func (f *Cache) BannedWords() []*ent.BannedWords {
	f.RLock()
	defer f.RUnlock()
	return f.bannedWords
}

func (f *Cache) IsValidGenerationModelID(id uuid.UUID) bool {
	for _, model := range f.GenerationModels() {
		if model.ID == id && model.IsActive {
			return true
		}
	}
	return false
}

func (f *Cache) GetDefaultInferenceStepsForModel(id uuid.UUID) (int32, error) {
	for _, model := range f.GenerationModels() {
		if model.ID == id {
			return model.DefaultInferenceSteps, nil
		}
	}
	return 0, fmt.Errorf("Couldn't find default inference steps for the model")
}

func (f *Cache) IsValidUpscaleModelID(id uuid.UUID) bool {
	for _, model := range f.UpscaleModels() {
		if model.ID == id {
			return true
		}
	}
	return false
}

func (f *Cache) IsValidShedulerID(id uuid.UUID) bool {
	for _, scheduler := range f.Schedulers() {
		if scheduler.ID == id {
			return true
		}
	}
	return false
}

func (f *Cache) IsValidVoiceoverModelID(id uuid.UUID) bool {
	for _, model := range f.VoiceoverModels() {
		if model.ID == id {
			return true
		}
	}
	return false
}

func (f *Cache) IsValidVoiceoverSpeakerID(speakerId uuid.UUID, modelId uuid.UUID) bool {
	for _, speaker := range f.VoiceoverSpeakers() {
		if speaker.ID == speakerId && speaker.ModelID == modelId {
			return true
		}
	}
	return false
}

func (f *Cache) GetVoiceoverSpeakersForModel(modelId uuid.UUID) []*ent.VoiceoverSpeaker {
	speakers := []*ent.VoiceoverSpeaker{}
	for _, speaker := range f.VoiceoverSpeakers() {
		if speaker.ModelID == modelId {
			speakers = append(speakers, speaker)
		}
	}
	return speakers
}

func (f *Cache) GetAllGenerationModels() []*ent.GenerationModel {
	f.RLock()
	defer f.RUnlock()
	return f.GenerationModels()
}

func (f *Cache) GetGenerationModelFromID(id uuid.UUID) *ent.GenerationModel {
	for _, model := range f.GenerationModels() {
		if model.ID == id {
			return model
		}
	}
	return nil
}

func (f *Cache) GetGenerationModelNameFromID(id uuid.UUID) string {
	for _, model := range f.GenerationModels() {
		if model.ID == id {
			return model.NameInWorker
		}
	}
	return ""
}

func (f *Cache) GetUpscaleModelFromID(id uuid.UUID) *ent.UpscaleModel {
	for _, model := range f.UpscaleModels() {
		if model.ID == id {
			return model
		}
	}
	return nil
}

func (f *Cache) GetUpscaleModelNameFromID(id uuid.UUID) string {
	for _, model := range f.UpscaleModels() {
		if model.ID == id {
			return model.NameInWorker
		}
	}
	return ""
}

func (f *Cache) GetSchedulerNameFromID(id uuid.UUID) string {
	for _, scheduler := range f.Schedulers() {
		if scheduler.ID == id {
			return scheduler.NameInWorker
		}
	}
	return ""
}

func (f *Cache) GetVoiceoverModelNameFromID(id uuid.UUID) string {
	for _, model := range f.VoiceoverModels() {
		if model.ID == id {
			return model.NameInWorker
		}
	}
	return ""
}

func (f *Cache) GetVoiceoverSpeakerNameFromID(speakerId uuid.UUID) string {
	for _, speaker := range f.VoiceoverSpeakers() {
		if speaker.ID == speakerId {
			return speaker.NameInWorker
		}
	}
	return ""
}

func (f *Cache) IsAdmin(id uuid.UUID) bool {
	for _, adminID := range f.AdminIDs() {
		if adminID == id {
			return true
		}
	}
	return false
}

func (f *Cache) SetAdminUUIDs(ids []uuid.UUID) {
	f.Lock()
	defer f.Unlock()
	f.adminIDs = ids
}

func (f *Cache) AdminIDs() []uuid.UUID {
	f.RLock()
	defer f.RUnlock()
	return f.adminIDs
}

func (f *Cache) UpdateIPBlacklist(ips []string) {
	f.Lock()
	defer f.Unlock()
	f.iPBlacklist = ips
}

func (f *Cache) IPBlacklist() []string {
	f.RLock()
	defer f.RUnlock()
	return f.iPBlacklist
}

func (f *Cache) UpdateThumbmarkIDBlacklist(tm []string) {
	f.Lock()
	defer f.Unlock()
	f.thumbmarkIDBlacklist = tm
}

func (f *Cache) ThumbmarkIDBlacklist() []string {
	f.RLock()
	defer f.RUnlock()
	return f.thumbmarkIDBlacklist
}

func (f *Cache) UpdateUsernameBlacklist(blacklist []string) {
	f.Lock()
	defer f.Unlock()
	f.usernameBlacklist = blacklist
}

func (f *Cache) IsUsernameBlacklisted(username string) bool {
	f.RLock()
	defer f.RUnlock()
	for _, blacklistedUsername := range f.usernameBlacklist {
		if username == blacklistedUsername {
			return true
		}
	}
	return false
}

func (f *Cache) IsIPBanned(ip string) bool {
	return slices.Contains(f.IPBlacklist(), ip)
}

func (f *Cache) IsThumbmarkIDBanned(ip string) bool {
	return slices.Contains(f.ThumbmarkIDBlacklist(), ip)
}

func (f *Cache) GetDefaultGenerationModel() *ent.GenerationModel {
	var defaultModel *ent.GenerationModel
	for _, model := range f.GenerationModels() {
		// always set at least 1 as default
		if defaultModel == nil {
			defaultModel = model
		}
		if model.IsDefault {
			defaultModel = model
		}
	}
	return defaultModel
}

func (f *Cache) GetDefaultUpscaleModel() *ent.UpscaleModel {
	var defaultModel *ent.UpscaleModel
	for _, model := range f.UpscaleModels() {
		// always set at least 1 as default
		if defaultModel == nil {
			defaultModel = model
		}
		if model.IsDefault {
			defaultModel = model
		}
	}
	return defaultModel
}

func (f *Cache) GetDefaultVoiceoverModel() *ent.VoiceoverModel {
	var defaultModel *ent.VoiceoverModel
	for _, model := range f.VoiceoverModels() {
		// always set at least 1 as default
		if defaultModel == nil {
			defaultModel = model
		}
		if model.IsDefault {
			defaultModel = model
		}
	}
	return defaultModel
}

func (f *Cache) GetDefaultVoiceoverSpeaker() *ent.VoiceoverSpeaker {
	var defaultSpeaker *ent.VoiceoverSpeaker
	for _, speaker := range f.VoiceoverSpeakers() {
		// always set at least 1 as default
		if defaultSpeaker == nil {
			defaultSpeaker = speaker
		}
		if speaker.IsDefault {
			defaultSpeaker = speaker
		}
	}
	return defaultSpeaker
}

func (f *Cache) GetDefaultScheduler() *ent.Scheduler {
	var defaultScheduler *ent.Scheduler
	for _, scheduler := range f.Schedulers() {
		// always set at least 1 as default
		if defaultScheduler == nil {
			defaultScheduler = scheduler
		}
		if scheduler.IsDefault {
			defaultScheduler = scheduler
		}
	}
	return defaultScheduler
}

func (f *Cache) GetGenerationModelByID(id uuid.UUID) *ent.GenerationModel {
	for _, model := range f.GenerationModels() {
		if model.ID == id {
			return model
		}
	}
	return nil
}

func (f *Cache) GetCompatibleSchedulerIDsForModel(ctx context.Context, modelId uuid.UUID) []uuid.UUID {
	m := f.GetGenerationModelByID(modelId)
	if m == nil {
		return []uuid.UUID{}
	}
	schedulerIds := make([]uuid.UUID, len(m.Edges.Schedulers))
	for i, scheduler := range m.Edges.Schedulers {
		schedulerIds[i] = scheduler.ID
	}
	return schedulerIds
}

func (f *Cache) GetDefaultSchedulerIDForModel(modelId uuid.UUID) uuid.UUID {
	m := f.GetGenerationModelByID(modelId)
	if m == nil {
		return uuid.Nil
	}
	return *m.DefaultSchedulerID
}
