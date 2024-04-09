package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/stablecog/sc-go/cron/models"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/utils"
)

// Sends a discord notification on either the healthy/unhealthy interval depending on status
func FireServerReadyWebhook(version string, msg string, buildStart string) error {
	webhookUrl := utils.GetEnv().DiscordWebhookUrlDeploy
	if webhookUrl == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL_DEPLOY not set")
	}
	// Parse build start as int
	buildStartInt, err := strconv.Atoi(buildStart)
	buildStartStr := ""
	if err != nil {
		log.Error("Error parsing build start", "err", err)
	} else {
		buildStartStr = fmt.Sprintf(" in %ds", int(time.Now().Sub(utils.SecondsSinceEpochToTime(int64(buildStartInt))).Seconds()))
	}
	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf(`%s  •  %s`, msg, version),
				Color: 5763719,
				Fields: []models.DiscordWebhookField{
					{
						Value: fmt.Sprintf("```Deployed%s```", buildStartStr),
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

// Sends a discord notification when a new subscriber signs up
func NewSubscriberWebhook(repo *repository.Repository, user *ent.User, productId string) error {
	webhookUrl := utils.GetEnv().DiscordWebhookUrlNewSub
	if webhookUrl == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL_NEWSUB not set")
	}
	nSubs, err := repo.GetNSubscribers()
	if err != nil {
		log.Error("Error getting nSubs", "err", err)
		return err
	}
	// Get credit type by product ID
	ctype, err := repo.GetCreditTypeByStripeProductID(productId)
	if err != nil || ctype == nil {
		log.Error("Error getting credit type", "err", err, "ctype", ctype)
		return err
	}
	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("🎉 New Sub #%d • %s", nSubs, ctype.Name),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "Email",
						Value: user.Email,
					},
					{
						Name:  "Plan",
						Value: ctype.Name,
					},
					{
						Name:  "Supabase ID",
						Value: user.ID.String(),
					},
					{
						Name:  "Stripe ID",
						Value: user.StripeCustomerID,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

func SubscriptionUpgradeWebhook(
	repo *repository.Repository,
	user *ent.User,
	productIdOld string,
	productIdNew string,
) error {
	webhookUrl := utils.GetEnv().DiscordWebhookUrlNewSub
	if webhookUrl == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL_NEWSUB not set")
	}
	// Get credit type by product ID
	creditTypeOld, err := repo.GetCreditTypeByStripeProductID(productIdOld)
	if err != nil || creditTypeOld == nil {
		log.Error("Error getting credit type", "err", err, "creditTypeOld", creditTypeOld)
		return err
	}

	creditTypeNew, err := repo.GetCreditTypeByStripeProductID(productIdNew)
	if err != nil || creditTypeNew == nil {
		log.Error("Error getting credit type", "err", err, "creditTypeNew", creditTypeNew)
		return err
	}

	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("🎉 Sub Upgrade • %s", creditTypeNew.Name),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "Email",
						Value: user.Email,
					},
					{
						Name:  "Old Plan",
						Value: creditTypeOld.Name,
					},
					{
						Name:  "New Plan",
						Value: creditTypeNew.Name,
					},
					{
						Name:  "Supabase ID",
						Value: user.ID.String(),
					},
					{
						Name:  "Stripe ID",
						Value: user.StripeCustomerID,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

// Sends a discord notification when adhoc credits purchased
func AdhocCreditsPurchasedWebhook(repo *repository.Repository, user *ent.User, creditType *ent.CreditType) error {
	webhookUrl := utils.GetEnv().DiscordWebhookUrlNewSub
	if webhookUrl == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL_NEWSUB not set")
	}
	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("🎉 Cred Purchase • %s", creditType.Name),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "Email",
						Value: user.Email,
					},
					{
						Name:  "Pack",
						Value: creditType.Name,
					},
					{
						Name:  "Supabase ID",
						Value: user.ID.String(),
					},
					{
						Name:  "Stripe ID",
						Value: user.StripeCustomerID,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

func FireGeoIPBannedUserWebhook(ip string, email string, domain string, userid string, countryCode string, thumbmarkID string) error {
	webhookUrl := utils.GetEnv().GeoIpWebhook
	if webhookUrl == "" {
		return fmt.Errorf("GEOIP_WEBHOOK not set")
	}

	thumbmark := "Unknown"
	if thumbmarkID != "" {
		thumbmark = thumbmarkID
	}

	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("%s IP", countryCode),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "IP",
						Value: ip,
					},
					{
						Name:  "Banned User ID",
						Value: userid,
					},
					{
						Name:  "Banned Email",
						Value: email,
					},
					{
						Name:  "Banned Domain",
						Value: domain,
					},
					{
						Name:  "Thumbmark ID",
						Value: thumbmark,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

func FireGeoIPSuspiciousUserWebhook(ip string, email string, domain string, userid string, countryCode string, thumbmarkID string) error {
	webhookUrl := utils.GetEnv().GeoIpWebhook
	if webhookUrl == "" {
		return fmt.Errorf("GEOIP_WEBHOOK not set")
	}

	thumbmark := "Unknown"
	if thumbmarkID != "" {
		thumbmark = thumbmarkID
	}

	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("%s IP (Suspicious)", countryCode),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "IP",
						Value: ip,
					},
					{
						Name:  "User ID",
						Value: userid,
					},
					{
						Name:  "Email",
						Value: email,
					},
					{
						Name:  "Domain",
						Value: domain,
					},
					{
						Name:  "Thumbmark ID",
						Value: thumbmark,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}

func FireBannedUserWebhook(ip string, email string, domain string, userid string, countryCode string, thumbmarkID string, banReasons []string) error {
	webhookUrl := utils.GetEnv().GeoIpWebhook
	if webhookUrl == "" {
		return fmt.Errorf("GEOIP_WEBHOOK not set")
	}

	thumbmark := "Unknown"
	if thumbmarkID != "" {
		thumbmark = thumbmarkID
	}

	banReasonsString := ""
	for i, reason := range banReasons {
		number := i + 1
		str := fmt.Sprintf("%d- %s", number, reason)
		if i != 0 {
			str = fmt.Sprintf("\n%s", str)
		}
		banReasonsString = fmt.Sprintf("%s%s", banReasonsString, str)
	}

	// Build webhook body
	body := models.DiscordWebhookBody{
		Embeds: []models.DiscordWebhookEmbed{
			{
				Title: fmt.Sprintf("Banned User!"),
				Color: 11437567,
				Fields: []models.DiscordWebhookField{
					{
						Name:  "Banned Email",
						Value: email,
					},
					{
						Name:  "Banned User ID",
						Value: userid,
					},
					{
						Name:  "IP",
						Value: ip,
					},
					{
						Name:  "Domain",
						Value: domain,
					},
					{
						Name:  "Thumbmark ID",
						Value: thumbmark,
					},
					{
						Name:  "Ban Reasons",
						Value: banReasonsString,
					},
				},
				Footer: models.DiscordWebhookEmbedFooter{
					Text: fmt.Sprintf("%s", time.Now().Format(time.RFC1123)),
				},
			},
		},
		Attachments: []models.DiscordWebhookAttachment{},
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshalling webhook body", "err", err)
		return err
	}
	res, postErr := http.Post(webhookUrl, "application/json", bytes.NewBuffer(reqBody))
	if postErr != nil {
		log.Error("Error sending webhook", "err", postErr)
		return postErr
	}
	defer res.Body.Close()

	return nil
}
