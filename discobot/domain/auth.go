package domain

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/discobot/components"
	"github.com/stablecog/sc-go/discobot/responses"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/shared"
)

var ErrNotAuthorized = errors.New("not authorized")

// Shared auth wrapper, returns nil if unauthorized
func (d *DiscoDomain) CheckAuthorization(s *discordgo.Session, i *discordgo.InteractionCreate) *ent.User {
	if i.Member == nil {
		return nil
	}
	u, err := d.Repo.GetUserByDiscordID(i.Member.User.ID)
	if err != nil && !ent.IsNotFound(err) {
		log.Errorf("Failed to get user by discord ID %v", err)
		responses.ErrorResponseInitial(s, i, responses.PRIVATE)
		return nil
	}
	if err != nil && ent.IsNotFound(err) {
		// Set token in redis
		token, err := d.Redis.SetDiscordVerifyToken(i.Member.User.ID)
		if err != nil {
			log.Errorf("Failed to set discord verify token in redis %v", err)
			responses.ErrorResponseInitial(s, i, responses.PRIVATE)
			return nil
		}

		// Create URL params for login
		params := url.Values{}
		params.Add("platform_token", token)
		params.Add("platform_user_id", i.Member.User.ID)
		params.Add("platform_username", i.Member.User.Username)
		params.Add("platform_avatar_url", url.QueryEscape(i.Member.AvatarURL("128")))

		// Auth msg
		err = responses.InitialInteractionResponse(s,
			i,
			&responses.InteractionResponseOptions{
				EmbedTitle:   "🚀 Sign in to get started",
				EmbedContent: "Connect your Stablecog account to your Discord account to get started.\n\n",
				EmbedFooter:  "By signing in you agree to our Terms of Service and Privacy Policy.",
				ActionRowOne: []*components.SCDiscordComponent{
					components.NewLinkButton("Sign in", fmt.Sprintf("https://stablecog.com/connect/discord?%s", params.Encode()), "🔑"),
					components.NewLinkButton("Terms & Privacy", "https://stablecog.com/legal", ""),
				},
				Privacy: responses.PRIVATE,
			},
		)
		if err != nil {
			responses.ErrorResponseInitial(s, i, responses.PRIVATE)
			return nil
		}
		// Delete message when link expires
		time.AfterFunc(shared.DISCORD_VERIFY_TOKEN_EXPIRY, func() {
			s.InteractionResponseDelete(i.Interaction)
		})
		return nil
	}

	return u
}
