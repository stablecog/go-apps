package interactions

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/discobot/responses"
	"github.com/stablecog/sc-go/log"
	srvres "github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/utils"
)

func (c *DiscordInteractionWrapper) HandleTip(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Find the channel that the message came from.
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}
	if channel.GuildID == "" {
		// Is a DM
		return
	}

	// Get mentions
	for _, mention := range m.Mentions {
		if mention.ID == m.Author.ID {
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "You can't send tips to yourself.")
			return
		}
	}

	if len(m.Mentions) != 1 {
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		// Send DM
		dmChl, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			log.Error("Failed to create DM channel", "err", err)
			return
		}
		_, err = s.ChannelMessageSend(dmChl.ID, "You can only tip 1 person at a time.")
		return
	}

	// Get the users
	tippedBy, err := c.Repo.GetUserByDiscordID(m.Author.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "You need to register your account before you can tip. Try using /authenticate to get started, it's free!")
			return
		} else {
			log.Error("Failed to get user by discord id", "err", err)
			return
		}
	}
	tippedTo, err := c.Repo.GetUserByDiscordID(m.Mentions[0].ID)
	if err != nil && !ent.IsNotFound(err) {
		log.Error("Failed to get user by discord id", "err", err)
		return
	}
	var tippedToId *uuid.UUID
	if tippedTo != nil {
		tippedToId = utils.ToPtr(tippedTo.ID)
	}

	amt, err := utils.ExtractAmountsFromString(m.Content)
	if err != nil {
		switch err {
		case utils.AmountAmbiguousError:
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "You can only specify 1 amount in your message.")
			return
		case utils.AmountMissingError:
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "You need to specify an amount in your message.")
			return
		case utils.AmountNotIntegerError:
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "The amount you specified is not a valid number. It must be a whole number, example: `123.45` is not valid but `123` is.")
			return
		default:
			log.Error("Failed to extract amounts from string", "err", err)
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}
	}

	// Send tip
	success, err := c.Repo.TipCreditsToUser(tippedBy.ID, tippedToId, m.Mentions[0].ID, int32(amt))
	if err != nil || !success {
		if errors.Is(err, srvres.InsufficientCreditsErr) {
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			// Send DM
			dmChl, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Error("Failed to create DM channel", "err", err)
				return
			}
			_, err = s.ChannelMessageSend(dmChl.ID, "You don't have enough credits to send that tip. Use /info to see your available credits!")
			return
		}
		log.Error("Failed to tip credits to user", "err", err)
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		return
	}

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	s.MessageReactionAdd(m.ChannelID, m.ID, "🤑")

	// Send DM to the receiver
	dmChl, err := s.UserChannelCreate(m.Mentions[0].ID)
	if err != nil {
		log.Error("Failed to create DM channel", "err", err)
		return
	}

	// Different flows if registered or not
	if tippedToId != nil {

	}
	_, err = s.ChannelMessageSendComplex(dmChl.ID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{responses.NewEmbed(
			"Stablecog Credits Received!",
			fmt.Sprintf("You received %d credits from %s!", amt, m.Author.Username),
			"",
		),
		},
	},
	)
}
