package interactions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"github.com/stablecog/sc-go/database"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/discobot/domain"
	"github.com/stablecog/sc-go/server/analytics"
	"github.com/stablecog/sc-go/server/clip"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/scworker"
	"github.com/stablecog/sc-go/server/translator"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/shared/queue"
	"github.com/stablecog/sc-go/utils"
)

// Create new wrapper and register interactions
func NewDiscordInteractionWrapper(
	repo *repository.Repository,
	redis *database.RedisWrapper,
	supabase *database.SupabaseAuth,
	sMap *shared.SyncMap[chan requests.CogWebhookMessage],
	qThrottler *shared.UserQueueThrottlerMap,
	safetyChecker *translator.TranslatorSafetyChecker,
	track *analytics.AnalyticsService,
	LoginInteractionMap *shared.SyncMap[*LoginInteraction],
	MQClient queue.MQClient,
) *DiscordInteractionWrapper {
	// Setup S3 Client
	region := utils.GetEnv().S3Img2ImgRegion
	accessKey := utils.GetEnv().S3Img2ImgAccessKey
	secretKey := utils.GetEnv().S3Img2ImgSecretKey

	s3ConfigImg := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(utils.GetEnv().S3Img2ImgEndpoint),
		Region:      aws.String(region),
	}

	newSessionImg := session.New(s3ConfigImg)
	s3ClientImg := s3.New(newSessionImg)

	// Setup S3 Client regular
	region = utils.GetEnv().S3Region
	accessKey = utils.GetEnv().S3AccessKey
	secretKey = utils.GetEnv().S3SecretKey

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(utils.GetEnv().S3Endpoint),
		Region:      aws.String(region),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	// Create wrapper
	wrapper := &DiscordInteractionWrapper{
		Disco:               &domain.DiscoDomain{Repo: repo, Redis: redis, SupabaseAuth: supabase},
		SupabseAuth:         supabase,
		LoginInteractionMap: LoginInteractionMap,
		SCWorker: &scworker.SCWorker{
			Repo:           repo,
			Redis:          redis,
			QueueThrottler: qThrottler,
			Track:          track,
			SMap:           sMap,
			SafetyChecker:  safetyChecker,
			MQClient:       MQClient,
			S3Img:          s3ClientImg,
			S3:             s3Client,
		},
		Clip: clip.NewClipService(redis, safetyChecker),
		Repo: repo,
	}
	// Register commands
	commands := []*DiscordInteraction{
		wrapper.NewHelpCommand(),
		wrapper.NewAuthenticateCommand(),
		wrapper.NewInfoCommand(),
		wrapper.NewImageCommand(),
		wrapper.NewUpscaleCommand(),
		// wrapper.NewVoiceoverCommand(),
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
	Disco               *domain.DiscoDomain
	SupabseAuth         *database.SupabaseAuth
	LoginInteractionMap *shared.SyncMap[*LoginInteraction]
	Commands            []*DiscordInteraction
	Components          []*DiscordInteraction
	SCWorker            *scworker.SCWorker
	Repo                *repository.Repository
	Clip                *clip.ClipService
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
