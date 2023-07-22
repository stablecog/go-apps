package interactions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/enttypes"
	"github.com/stablecog/sc-go/discobot/aspectratio"
	"github.com/stablecog/sc-go/discobot/components"
	"github.com/stablecog/sc-go/discobot/responses"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	srvresponses "github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/server/scworker"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
)

func (c *DiscordInteractionWrapper) NewImageCommand() *DiscordInteraction {
	// Build model choices
	// Ensure default is at the top
	defaultModel := shared.GetCache().GetDefaultGenerationModel()
	modelChoices := []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  fmt.Sprintf("%s (default)", defaultModel.NameInWorker),
			Value: defaultModel.ID.String(),
		},
	}
	for _, model := range shared.GetCache().GenerateModels {
		if model.ID == defaultModel.ID {
			continue
		}
		if model.IsActive && !model.IsHidden {
			modelChoices = append(modelChoices, &discordgo.ApplicationCommandOptionChoice{
				Name:  model.NameInWorker,
				Value: model.ID.String(),
			})
		}
	}

	// Build aspect ratio choices
	aspectRatioChoices := []*discordgo.ApplicationCommandOptionChoice{}
	for _, ratio := range aspectratio.AvailableRatios {
		aspectRatioChoices = append(aspectRatioChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  ratio.String(),
			Value: ratio,
		})
	}

	return &DiscordInteraction{
		// Command spec
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "image",
			Description: "Create an image with Stablecog.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "The prompt for the generation.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "negative-prompt",
					Description: "The negative prompt for the generation.",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "image-count",
					Description: "The number of images to generate.",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "model",
					Description: "The model for the generation.",
					Required:    false,
					Choices:     modelChoices,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "aspect-ratio",
					Description: "The aspect ratio for the generation.",
					Required:    false,
					Choices:     aspectRatioChoices,
				},
				{
					Type:        discordgo.ApplicationCommandOptionAttachment,
					Name:        "init-image",
					Description: "Use an initial image.",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "seed",
					Description: "The seed for the generation.",
					Required:    false,
				},
			},
		},
		// The handler for the command
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var discordUserId string
			if i.Member != nil {
				discordUserId = i.Member.User.ID
			} else {
				discordUserId = i.User.ID
			}
			if u := c.Disco.CheckAuthorization(s, i); u != nil {
				// Access options in the order provided by the user.
				options := i.ApplicationCommandData().Options

				// Parse options
				var prompt string
				var negativePrompt string
				var modelId *uuid.UUID
				var aspectRatio *aspectratio.AspectRatio
				numOutputs := 4

				// Attachment of init image
				var attachmentId string
				var initImage string
				var seed *int

				for _, option := range options {
					switch option.Name {
					case "prompt":
						prompt = option.StringValue()
					case "negative-prompt":
						negativePrompt = option.StringValue()
					case "image-count":
						numOutputs = int(option.IntValue())
					case "model":
						modelId = utils.ToPtr[uuid.UUID](uuid.MustParse(option.StringValue()))
					case "aspect-ratio":
						aspectRatio = utils.ToPtr(aspectratio.AspectRatio(option.IntValue()))
					case "init-image":
						id, ok := option.Value.(string)
						if !ok {
							log.Errorf("Invalid image attachment for upscale command: %v", i.ApplicationCommandData())
							responses.ErrorResponseInitial(s, i, responses.PRIVATE)
							return
						}
						attachmentId = id
					case "seed":
						seed = utils.ToPtr[int](int(option.IntValue()))
					}
				}

				if attachmentId != "" {
					attachment, ok := i.ApplicationCommandData().Resolved.Attachments[attachmentId]
					if !ok {
						log.Errorf("No image attachment for generate command: %v", i.ApplicationCommandData())
						responses.ErrorResponseInitial(s, i, responses.PRIVATE)
						return
					}
					initImage = attachment.URL

					if attachment.ContentType != "image/png" && attachment.ContentType != "image/jpeg" && attachment.ContentType != "image/jpg" && attachment.ContentType != "image/webp" {
						responses.InitialInteractionResponse(s, i, &responses.InteractionResponseOptions{
							Privacy:      responses.PRIVATE,
							EmbedTitle:   "❌ Attachment type is not supported",
							EmbedContent: "The attachment can be a PNG, JPEG, or WEBP image.",
						})
						return
					}
				}

				if modelId == nil {
					modelId = utils.ToPtr(shared.GetCache().GetDefaultGenerationModel().ID)
				}

				if aspectRatio == nil {
					aspectRatio = utils.ToPtr(aspectratio.DefaultAspectRatio)
				}

				// Validate req/apply defaults
				req := requests.CreateGenerationRequest{
					Prompt:         prompt,
					NegativePrompt: negativePrompt,
					ModelId:        modelId,
					NumOutputs:     utils.ToPtr[int32](int32(numOutputs)),
					InitImageUrl:   initImage,
					Seed:           seed,
				}
				if aspectRatio != nil {
					width, height := aspectRatio.GetWidthHeightForModel(*modelId)
					req.Width = utils.ToPtr[int32](width)
					req.Height = utils.ToPtr[int32](height)
				}
				err := req.Validate(true)
				if err != nil {
					responses.ErrorResponseInitialValidation(s, i, err.Error(), responses.PRIVATE)
					return
				}

				credits, err := c.Repo.GetNonExpiredCreditTotalForUser(u.ID, nil)
				if err != nil {
					log.Errorf("Error getting credits for user: %v", err)
					responses.ErrorResponseInitial(s, i, responses.PRIVATE)
					return
				}
				if credits < int(req.Cost()) {
					responses.InitialInteractionResponse(s, i, responses.InsufficientCreditsResponseOptions(req.Cost(), int32(credits)))
					return
				}

				// Always create initial message
				responses.InitialLoadingResponse(s, i, responses.PUBLIC)

				// Create context
				ctx := context.Background()
				res, err := scworker.CreateGeneration(
					ctx,
					enttypes.SourceTypeDiscord,
					nil,
					c.SafetyChecker,
					c.Repo,
					c.Redis,
					c.SMap,
					c.QThrottler,
					u,
					nil,
					req,
				)
				if err != nil {
					if errors.Is(err, srvresponses.InsufficientCreditsErr) {
						credits, err := c.Repo.GetNonExpiredCreditTotalForUser(u.ID, nil)
						if err != nil {
							log.Errorf("Error getting credits for user: %v", err)
							responses.ErrorResponseEdit(s, i)
							return
						}
						responses.InteractionEdit(s, i, responses.InsufficientCreditsResponseOptions(req.Cost(), int32(credits)))
						return
					}
					log.Errorf("Error creating generation: %v", err)
					responses.ErrorResponseEdit(s, i)
					return
				}

				var imageUrls []string
				var actionRowOne []*components.SCDiscordComponent
				for i, output := range res.Outputs {
					if output.ImageURL != nil {
						imageUrls = append(imageUrls, *output.ImageURL)
						actionRowOne = append(actionRowOne, components.NewButton(fmt.Sprintf("Upscale #%d", i+1), fmt.Sprintf("upscale:%s:number:%d", output.ID.String(), i+1), "✨"))
					}
				}

				// Send the image
				_, err = responses.InteractionEdit(s, i, &responses.InteractionResponseOptions{
					Content:   utils.ToPtr(fmt.Sprintf("<@%s> **%s**", discordUserId, prompt)),
					ImageURLs: imageUrls,
					Embeds: []*discordgo.MessageEmbed{
						responses.NewGenerationMetadataEmbed(shared.GetCache().GetGenerationModelNameFromID(*req.ModelId)),
					},
					ActionRowOne: actionRowOne,
				},
				)
				if err != nil {
					log.Error(err)
					responses.ErrorResponseEdit(s, i)
				}
			} else {
				c.LoginInteractionMap.Put(discordUserId, &LoginInteraction{
					Session:     s,
					Interaction: i,
					InsertedAt:  time.Now(),
				})
			}
		},
	}
}

// Handle upscaling
func (c *DiscordInteractionWrapper) HandleUpscaleGeneration(s *discordgo.Session, i *discordgo.InteractionCreate, outputId uuid.UUID, number int) {
	var discordUserId string
	if i.Member != nil {
		discordUserId = i.Member.User.ID
	} else {
		discordUserId = i.User.ID
	}
	if u := c.Disco.CheckAuthorization(s, i); u != nil {
		// Disable the button
		// if len(i.Message.Components) < 1 {
		// 	log.Errorf("Error getting action row")
		// 	responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		// 	return
		// }
		// actionRowRaw, err := i.Message.Components[0].MarshalJSON()
		// if err != nil {
		// 	log.Errorf("Error getting action row: %v", err)
		// 	responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		// 	return
		// }

		// // Parse as SCDiscordComponent
		// var actionRow components.SCDiscordActionRow
		// err = json.Unmarshal(actionRowRaw, &actionRow)
		// if err != nil {
		// 	log.Errorf("Error getting action row: %v", err)
		// 	responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		// 	return
		// }

		// // Get button from action row
		// if len(actionRow.Components) < number {
		// 	log.Errorf("Error getting button")
		// 	responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		// 	return
		// }
		// actionRow.Components[number-1].Disabled = utils.ToPtr(true)
		// marshalled, err := actionRow.AsMessageComponent()
		// if err != nil {
		// 	log.Errorf("Error getting action row: %v", err)
		// 	responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		// 	return
		// }

		// err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 	Type: discordgo.InteractionResponseUpdateMessage,
		// 	Data: &discordgo.InteractionResponseData{
		// 		Components: []discordgo.MessageComponent{marshalled},
		// 	},
		// })

		// See if the output is already upscaled, send private response to avoid pollution
		existingOutput, err := c.Repo.GetGenerationOutput(outputId)
		if err != nil {
			log.Errorf("Error getting output: %v", err)
			responses.ErrorResponseInitial(s, i, responses.PRIVATE)
			return
		}
		if existingOutput.UpscaledImagePath != nil {
			// Send the image
			err = responses.InitialInteractionResponse(s, i, &responses.InteractionResponseOptions{
				Content: utils.ToPtr(fmt.Sprintf("<@%s> ✨ Image has already been upscaled #%d \n%s", discordUserId, number, utils.GetURLFromImagePath(*existingOutput.UpscaledImagePath))),
				Embeds:  nil,
				Privacy: responses.PRIVATE,
			})
			return
		}

		req := requests.CreateUpscaleRequest{
			Input: outputId.String(),
		}

		credits, err := c.Repo.GetNonExpiredCreditTotalForUser(u.ID, nil)
		if err != nil {
			log.Errorf("Error getting credits for user: %v", err)
			responses.ErrorResponseInitial(s, i, responses.PRIVATE)
			return
		}
		if credits < int(req.Cost()) {
			responses.InitialInteractionResponse(s, i, responses.InsufficientCreditsResponseOptions(req.Cost(), int32(credits)))
			return
		}

		// Always create initial message
		responses.InitialLoadingResponse(s, i, responses.PUBLIC)

		// Create context
		ctx := context.Background()
		res, err := scworker.CreateUpscale(
			ctx,
			enttypes.SourceTypeDiscord,
			nil,
			c.Repo,
			c.Redis,
			c.SMap,
			c.QThrottler,
			u,
			req,
		)
		if err != nil {
			if errors.Is(err, srvresponses.InsufficientCreditsErr) {
				credits, err := c.Repo.GetNonExpiredCreditTotalForUser(u.ID, nil)
				if err != nil {
					log.Errorf("Error getting credits for user: %v", err)
					responses.ErrorResponseEdit(s, i)
					return
				}
				responses.InteractionEdit(s, i, responses.InsufficientCreditsResponseOptions(req.Cost(), int32(credits)))
				return
			}
			log.Errorf("Error creating upscale for output: %v", err)
			responses.ErrorResponseEdit(s, i)
			return
		}

		var upscaledImageUrl string
		for _, output := range res.Outputs {
			if output.UpscaledImageURL != nil {
				upscaledImageUrl = *output.UpscaledImageURL
			}
		}
		if upscaledImageUrl == "" {
			log.Errorf("Error getting upscaled image url")
			responses.ErrorResponseEdit(s, i)
			return
		}

		// Send the image
		_, err = responses.InteractionEdit(s, i, &responses.InteractionResponseOptions{
			Content: utils.ToPtr(fmt.Sprintf("<@%s> ✨ Upscaled #%d \n%s", discordUserId, number, upscaledImageUrl)),
			Embeds:  nil,
		},
		)
		if err != nil {
			log.Error(err)
			responses.ErrorResponseEdit(s, i)
		}
	} else {
		c.LoginInteractionMap.Put(discordUserId, &LoginInteraction{
			Session:     s,
			Interaction: i,
			InsertedAt:  time.Now(),
		})
	}
}
