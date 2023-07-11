package interactions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stablecog/sc-go/database"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/discobot/domain"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
)

// Create new wrapper and register interactions
func NewDiscordInteractionWrapper(repo *repository.Repository, redis *database.RedisWrapper, supabase *database.SupabaseAuth, sMap *shared.SyncMap[chan requests.CogWebhookMessage], qThrottler *shared.UserQueueThrottlerMap, safetyChecker *utils.TranslatorSafetyChecker) *DiscordInteractionWrapper {
	// Create wrapper
	wrapper := &DiscordInteractionWrapper{
		Disco:         &domain.DiscoDomain{Repo: repo, Redis: redis, SupabaseAuth: supabase},
		Repo:          repo,
		SupabseAuth:   supabase,
		SMap:          sMap,
		Redis:         redis,
		QThrottler:    qThrottler,
		SafetyChecker: safetyChecker,
	}
	// Register commands
	commands := []*DiscordInteraction{
		wrapper.NewAuthenticateCommand(),
		wrapper.NewImageCommand(),
		wrapper.NewVoiceoverCommand(),
	}
	// Register component responses
	components := []*DiscordInteraction{}
	// Set commands
	wrapper.Commands = commands
	// Set components
	wrapper.Components = components
	return wrapper
}

// Wrapper for all interactions
type DiscordInteractionWrapper struct {
	Disco         *domain.DiscoDomain
	Repo          *repository.Repository
	SupabseAuth   *database.SupabaseAuth
	Redis         *database.RedisWrapper
	SMap          *shared.SyncMap[chan requests.CogWebhookMessage]
	QThrottler    *shared.UserQueueThrottlerMap
	Commands      []*DiscordInteraction
	Components    []*DiscordInteraction
	SafetyChecker *utils.TranslatorSafetyChecker
}

// Specification for specific interactions
type DiscordInteraction struct {
	ApplicationCommand *discordgo.ApplicationCommand
	ComponentID        string
	Handler            func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (w *DiscordInteractionWrapper) GetHandlerForCommand(command string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, c := range w.Commands {
		if c.ApplicationCommand.Name == command {
			return c.Handler
		}
	}
	return nil
}

func (w *DiscordInteractionWrapper) GetHandlerForComponent(component string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, c := range w.Components {
		if c.ComponentID == component {
			return c.Handler
		}
	}
	return nil
}