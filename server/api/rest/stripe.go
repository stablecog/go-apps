package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/discord"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	scstripe "github.com/stablecog/sc-go/server/stripe"
	"github.com/stablecog/sc-go/utils"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"
	"golang.org/x/exp/slices"
)

// For creating customer portal session
func (c *RestAPI) HandleCreatePortalSession(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var stripeReq requests.StripePortalRequest
	err := json.Unmarshal(reqBody, &stripeReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	// Create portal session
	session, err := c.StripeClient.BillingPortalSessions.New(&stripe.BillingPortalSessionParams{
		Customer:  stripe.String(user.StripeCustomerID),
		ReturnURL: stripe.String(stripeReq.ReturnUrl),
	})

	if err != nil {
		log.Error("Error creating portal session", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	sessionResponse := responses.StripeSessionResponse{
		CustomerPortalURL: session.URL,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, sessionResponse)
}

// For creating a new subscription or upgrading one
// Rejects, if they have a subscription that is at a higher level than the target priceID
func (c *RestAPI) HandleCreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var stripeReq requests.StripeCheckoutRequest
	err := json.Unmarshal(reqBody, &stripeReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	// Make sure price ID exists in map
	var targetPriceID string
	var targetPriceLevel int
	adhocPrice := false
	for level, priceID := range scstripe.GetPriceIDs() {
		if priceID == stripeReq.TargetPriceID {
			targetPriceID = priceID
			targetPriceLevel = level
			break
		}
	}
	if targetPriceID == "" {
		// Check if it's a single purchase price
		for priceID := range scstripe.GetSinglePurchasePriceIDs() {
			if priceID == stripeReq.TargetPriceID {
				targetPriceID = priceID
				adhocPrice = true
				break
			}
		}
	}
	if targetPriceID == "" {
		responses.ErrBadRequest(w, r, "invalid_price_id", "")
		return
	}

	// Validate currency
	if !slices.Contains([]string{"usd", "eur"}, stripeReq.Currency) {
		responses.ErrBadRequest(w, r, "invalid_currency", "")
		return
	}

	// Get subscription
	customer, err := c.StripeClient.Customers.Get(user.StripeCustomerID, &stripe.CustomerParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("subscriptions"),
			},
		},
	})

	if err != nil {
		log.Error("Error getting customer", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	var currentPriceID string
	if customer.Subscriptions != nil {
		for _, sub := range customer.Subscriptions.Data {
			if sub.Status == stripe.SubscriptionStatusActive {
				for _, item := range sub.Items.Data {
					if item.Price.ID == targetPriceID {
						responses.ErrBadRequest(w, r, "already_subscribed", "")
						return
					}
					// If price ID is in map it's valid
					for _, priceID := range scstripe.GetPriceIDs() {
						if item.Price.ID == priceID {
							currentPriceID = item.Price.ID
							break
						}
					}
				}
				break
			}
		}
	}

	// If they don't have one, cannot buy adhoc
	if currentPriceID == "" && adhocPrice {
		responses.ErrBadRequest(w, r, "no_subscription", "")
		return
	}

	// If they have a current one, make sure they are upgrading
	if currentPriceID != "" && !adhocPrice {
		// Annual cannot upgrade
		if scstripe.IsAnnualPriceID(currentPriceID) {
			responses.ErrBadRequest(w, r, "cannot_upgrade", "Must cancel and re-subscribe with annual plan")
			return
		}
		var currentPriceLevel int
		for level, priceID := range scstripe.GetPriceIDs() {
			if priceID == currentPriceID {
				currentPriceLevel = level
				break
			}
		}
		if currentPriceLevel >= targetPriceLevel {
			responses.ErrBadRequest(w, r, "cannot_downgrade", "")
			return
		}
	}

	mode := stripe.CheckoutSessionModeSubscription
	if adhocPrice {
		mode = stripe.CheckoutSessionModePayment
	}
	// Create checkout session
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(user.StripeCustomerID),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(targetPriceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(mode)),
		SuccessURL: stripe.String(stripeReq.SuccessUrl),
		CancelURL:  stripe.String(stripeReq.CancelUrl),
		Currency:   stripe.String(stripeReq.Currency),
	}
	if adhocPrice {
		params.PaymentIntentData = &stripe.CheckoutSessionPaymentIntentDataParams{
			Metadata: map[string]string{
				"product": scstripe.GetSinglePurchasePriceIDs()[targetPriceID],
			},
		}
	}
	if stripeReq.PromotionCodeID != "" {
		params.Discounts = []*stripe.CheckoutSessionDiscountParams{
			{
				PromotionCode: stripe.String(stripeReq.PromotionCodeID),
			},
		}
	} else {
		params.AllowPromotionCodes = stripe.Bool(true)
	}

	session, err := c.StripeClient.CheckoutSessions.New(params)
	if err != nil {
		log.Error("Error creating checkout session", "err", err)
		if stripeErr, ok := err.(*stripe.Error); ok {
			responses.ErrInternalServerError(w, r, string(stripeErr.Code))
			return
		}
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	sessionResponse := responses.StripeSessionResponse{
		CheckoutURL: session.URL,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, sessionResponse)
}

// HTTP Post - handle stripe subscription downgrade
// Rejects if they don't have a subscription, or if they are not downgrading
func (c *RestAPI) HandleSubscriptionDowngrade(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var stripeReq requests.StripeDowngradeRequest
	err := json.Unmarshal(reqBody, &stripeReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	// Make sure price ID exists in map
	var targetPriceID string
	var targetPriceLevel int
	for level, priceID := range scstripe.GetPriceIDs() {
		if priceID == stripeReq.TargetPriceID {
			targetPriceID = priceID
			targetPriceLevel = level
			break
		}
	}
	if targetPriceID == "" {
		responses.ErrBadRequest(w, r, "invalid_price_id", "")
		return
	}

	// Get subscription
	customer, err := c.StripeClient.Customers.Get(user.StripeCustomerID, &stripe.CustomerParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("subscriptions"),
			},
		},
	})

	if err != nil {
		log.Error("Error getting customer", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	if customer.Subscriptions == nil || len(customer.Subscriptions.Data) == 0 || customer.Subscriptions.TotalCount == 0 {
		responses.ErrBadRequest(w, r, "no_active_subscription", "")
		return
	}

	var currentPriceID string
	var currentSubId string
	var currentItemId string
	for _, sub := range customer.Subscriptions.Data {
		if sub.Status == stripe.SubscriptionStatusActive && sub.CancelAt == 0 {
			for _, item := range sub.Items.Data {
				// If price ID is in map it's valid
				for _, priceID := range scstripe.GetPriceIDs() {
					if item.Price.ID == priceID {
						currentPriceID = item.Price.ID
						currentSubId = sub.ID
						currentItemId = item.ID
						break
					}
				}
				break
			}
		}
	}

	if currentPriceID == "" {
		responses.ErrBadRequest(w, r, "no_active_subscription", "")
		return
	}

	if currentPriceID == targetPriceID {
		responses.ErrBadRequest(w, r, "not_lower", "")
		return
	}

	// Can't downgrade from annual
	if scstripe.IsAnnualPriceID(currentPriceID) {
		responses.ErrBadRequest(w, r, "cannot_downgrade", "Cancel and re-subscribe")
		return
	}

	// Make sure this is a downgrade
	for level, priceID := range scstripe.GetPriceIDs() {
		if priceID == currentPriceID {
			if level <= targetPriceLevel {
				responses.ErrBadRequest(w, r, "not_lower", "")
				return
			}
			break
		}
	}

	// Execute subscription update
	_, err = c.StripeClient.Subscriptions.Update(currentSubId, &stripe.SubscriptionParams{
		ProrationBehavior: stripe.String("none"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				ID:    stripe.String(currentItemId),
				Price: stripe.String(targetPriceID),
			},
		},
	})

	if err != nil {
		log.Error("Error updating subscription", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	highestProductID, highestPriceID, cancelsAt, renewsAt, err := c.GetAndSyncStripeSubscriptionInfo(user.StripeCustomerID)

	if err != nil {
		log.Error("Error getting and syncing Stripe subscription info", "err", err)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"success":            true,
		"highest_product_id": highestProductID,
		"highest_price_id":   highestPriceID,
		"cancels_at":         cancelsAt,
		"renews_at":          renewsAt,
	})
}

func (c *RestAPI) HandleStripeWebhookSubscription(w http.ResponseWriter, r *http.Request) {
	s := time.Now()

	stripePaymentIntentEvents := []string{
		"payment_intent.created",
		"payment_intent.canceled",
		"payment_intent.payment_failed",
		"payment_intent.processing",
		"payment_intent.succeeded",
	}

	stripeInvoiceEvents := []string{
		"invoice.created",
		"invoice.updated",
		"invoice.deleted",
		"invoice.finalization_failed",
		"invoice.payment_failed",
		"invoice.finalized",
		"invoice.paid",
	}

	stripeSubscriptionEvents := []string{
		"customer.subscription.created",
		"customer.subscription.deleted",
		"customer.subscription.updated",
	}

	// Parse request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("🪝 🔴 Unable reading Stripe webhook body", err)
		responses.ErrBadRequest(w, r, "invalid stripe webhook body", "")
		return
	}

	// Verify signature
	endpointSecret := utils.GetEnv().StripeWebhookSubscriptionSecret

	event, err := webhook.ConstructEvent(reqBody, r.Header.Get("Stripe-Signature"), endpointSecret)
	log.Infof("🪝 🟡 Processing Stripe webhook event: %s", event.Type)

	if err != nil {
		log.Error("🪝 🔴 Unable verifying stripe webhook signature", err)
		responses.ErrBadRequest(w, r, "invalid stripe webhook signature", "")
		return
	}

	if !slices.Contains(stripeSubscriptionEvents, event.Type) &&
		!slices.Contains(stripeInvoiceEvents, event.Type) &&
		!slices.Contains(stripePaymentIntentEvents, event.Type) {
		log.Infof(`🪝 🔵 Stripe webhook event is not a registered event, not handling: %s`, event.Type)
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "OK")
	}

	var customerID string

	if slices.Contains(stripePaymentIntentEvents, event.Type) {
		paymentIntent, err := stripeObjectMapToPaymentIntent(event.Data.Object)
		if err != nil || paymentIntent == nil {
			log.Error("🪝 Unable parsing Stripe payment intent object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		customerID = paymentIntent.Customer
	} else if slices.Contains(stripeInvoiceEvents, event.Type) {
		invoice, err := stripeObjectMapToInvoiceObject(event.Data.Object)
		if err != nil || invoice == nil {
			log.Error("🪝 Unable parsing Stripe invoice object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		customerID = invoice.Customer
	} else {
		subscription, err := stripeObjectMapToCustomSubscriptionObject(event.Data.Object)
		if err != nil || subscription == nil {
			log.Error("🪝 Unable parsing Stripe subscription object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		customerID = subscription.Customer
	}

	user, userErr := c.Repo.GetUserByStripeCustomerId(customerID)
	if userErr != nil {
		log.Error("🪝 Error getting user from stripe customer id", "err", userErr)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	highestProductID, highestPriceID, cancelsAt, renewsAt, err := c.GetAndSyncStripeSubscriptionInfo(customerID)

	if err != nil {
		log.Error("🪝 🔴 Unable getting and syncing stripe subscription info", err)
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	log.Infof("🪝 🟢 Updated Stripe subscription info in DB | %dms | %s | userID: %s, customerID: %s, highestProductID: %s, highestPriceID: %s, cancelsAt: %v, renewsAt: %v", time.Since(s).Milliseconds(), event.Type, user.ID, customerID, highestProductID, highestPriceID, cancelsAt, renewsAt)

	render.Status(r, http.StatusOK)
	render.PlainText(w, r, "OK")
}

func (c *RestAPI) GetAndSyncStripeSubscriptionInfo(stripeCustomerID string) (string, string, *time.Time, *time.Time, error) {
	s := time.Now()
	customer, err := c.StripeClient.Customers.Get(stripeCustomerID, &stripe.CustomerParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("subscriptions"),
			},
		},
	})

	if err != nil {
		log.Error("🔴 GetAndSyncStripeSubscriptionInfo: Unable getting customer", err)
		return "", "", nil, nil, err
	}

	// Get subscription info
	highestProductID, highestPriceID, cancelsAt, renewsAt := extractSubscriptionInfoFromCustomer(customer)

	user, userErr := c.Repo.GetUserByStripeCustomerId(stripeCustomerID)
	if userErr != nil {
		log.Error("🔴 GetAndSyncStripeSubscriptionInfo: Unable getting user from stripe customer id", userErr)
		return "", "", nil, nil, userErr
	}

	_, updateErr := c.Repo.UpdateUserStripeSubscriptionInfo(user.ID, highestProductID, highestPriceID, cancelsAt, renewsAt, time.Now())
	if updateErr != nil {
		log.Error("🔴 GetAndSyncStripeSubscriptionInfo: Unable updating user stripe subscription info", updateErr)
		return "", "", nil, nil, updateErr
	}

	log.Infof("🟢 GetAndSyncStripeSubscriptionInfo: %dms", time.Since(s).Milliseconds())
	return highestProductID, highestPriceID, cancelsAt, renewsAt, nil
}

// Handle stripe webhooks in the following ways:
// invoice.payment_succeeded
//   - Apply credits to user depending on type (subscription, adhoc)
//   - For subscriptions, set active_product_id
//
// customer.subscription.deleted"
//   - For an immediate cancellation, we set active_product_id to nil if this is a cancellation
//     of the product ID we currently have set for them. (In case they upgraded, it won't unset their upgrade)
//
// customer.subscription.created
//   - For a subscription upgrade, we cancel all old subscriptions
//
// payment_intent.succeeded
//   - For adhoc payments, we apply credits to the user
func (c *RestAPI) HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Unable reading stripe webhook body", "err", err)
		responses.ErrBadRequest(w, r, "invalid stripe webhook body", "")
		return
	}

	// Verify signature
	endpointSecret := utils.GetEnv().StripeEndpointSecret

	event, err := webhook.ConstructEvent(reqBody, r.Header.Get("Stripe-Signature"), endpointSecret)
	if err != nil {
		log.Error("Unable verifying stripe webhook signature", "err", err)
		responses.ErrBadRequest(w, r, "invalid stripe webhook signature", "")
		return
	}

	switch event.Type {
	// For subscription upgrades, we want to cancel all old subscriptions
	case "customer.subscription.created":
		newSub, err := stripeObjectMapToSubscriptionObject(event.Data.Object)
		var newProduct string
		var oldProduct string
		if err != nil || newSub == nil {
			log.Error("Unable parsing stripe subscription object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		if newSub.Items != nil && len(newSub.Items.Data) > 0 && newSub.Items.Data[0].Price != nil && newSub.Items.Data[0].Price.Product != nil {
			newProduct = newSub.Items.Data[0].Price.Product.ID
		}
		// We need to see if they have more than one subscription
		subIter := c.StripeClient.Subscriptions.List(&stripe.SubscriptionListParams{
			Customer: stripe.String(newSub.Customer.ID),
		})
		for subIter.Next() {
			sub := subIter.Subscription()
			if sub.ID != newSub.ID {
				if sub.Items != nil && len(sub.Items.Data) > 0 && sub.Items.Data[0].Price != nil && sub.Items.Data[0].Price.Product != nil {
					oldProduct = sub.Items.Data[0].Price.Product.ID
				}
				// We need to cancel this subscription
				_, err := c.StripeClient.Subscriptions.Cancel(sub.ID, &stripe.SubscriptionCancelParams{
					Prorate: stripe.Bool(false),
				})
				if err != nil {
					log.Error("Unable canceling stripe subscription", "err", err)
					responses.ErrInternalServerError(w, r, err.Error())
					return
				}
			}
		}
		// Analytics
		if newProduct != "" && oldProduct != "" {
			go func() {
				user, err := c.Repo.GetUserByStripeCustomerId(newSub.Customer.ID)
				if err != nil {
					log.Error("Unable getting user from stripe customer id in upgrade subscription event", "err", err)
					return
				}
				go c.Track.SubscriptionUpgraded(user, oldProduct, newProduct)
				go discord.SubscriptionUpgradeWebhook(c.Repo, user, oldProduct, newProduct)
			}()
		}
	case "customer.subscription.deleted":
		sub, err := stripeObjectMapToCustomSubscriptionObject(event.Data.Object)
		if err != nil || sub == nil {
			log.Error("Unable parsing stripe subscription object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		user, err := c.Repo.GetUserByStripeCustomerId(sub.Customer)
		if err != nil {
			log.Error("Unable getting user from stripe customer id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		} else if user == nil {
			log.Error("User does not exist with stripe customer id: %s", sub.Customer)
			responses.ErrInternalServerError(w, r, "User does not exist with stripe customer id")
			return
		}
		// Get product Id from subscription
		if sub.Items != nil && len(sub.Items.Data) > 0 && sub.Items.Data[0].Price != nil {
			go func() {
				// Delay to avoid race with upgrades
				time.Sleep(30 * time.Second)
				affected, err := c.Repo.UnsetActiveProductID(user.ID, sub.Items.Data[0].Price.Product, nil)
				if err != nil {
					log.Error("Unable unsetting stripe product id", "err", err)
					return
				}
				if affected > 0 {
					// Subscription cancelled
					go c.Track.SubscriptionCancelled(user, sub.Items.Data[0].Price.Product)
				}
			}()
		}
	// Remove credits if necessary
	case "invoice.finalization_failed", "invoice.payment_failed":
		// We can parse the object as an invoice since that's the only thing we care about
		invoice, err := stripeObjectMapToInvoiceObject(event.Data.Object)
		if err != nil || invoice == nil {
			log.Error("Unable parsing stripe invoice object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}

		c.RevertCreditsInvoice(invoice, w, r)
		return
	// Subcription payments
	case "invoice.finalized", "invoice.paid", "invoice.updated":
		// We can parse the object as an invoice since that's the only thing we care about
		invoice, err := stripeObjectMapToInvoiceObject(event.Data.Object)
		if err != nil || invoice == nil {
			log.Error("Unable parsing stripe invoice object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}

		if invoice.Status == InvoiceStatusVoid || invoice.Status == InvoiceStatusDraft || invoice.Status == InvoiceStatusUncollectible {
			c.RevertCreditsInvoice(invoice, w, r)
			return
		}

		if invoice.Status == InvoiceStatusOpen {
			// Get payment intent
			pi, err := c.StripeClient.PaymentIntents.Get(invoice.PaymentIntent, nil)
			if err != nil {
				log.Error("Unable getting payment intent", "err", err)
				responses.ErrInternalServerError(w, r, err.Error())
				return
			}
			if pi.Status == stripe.PaymentIntentStatusRequiresConfirmation || pi.Status == stripe.PaymentIntentStatusRequiresAction || pi.Status == stripe.PaymentIntentStatusCanceled || pi.Status == stripe.PaymentIntentStatusRequiresPaymentMethod {
				c.RevertCreditsInvoice(invoice, w, r)
				return
			}
		}

		// We only care about renewal (cycle), create, and manual
		if invoice.BillingReason != InvoiceBillingReasonSubscriptionCycle && invoice.BillingReason != InvoiceBillingReasonSubscriptionCreate {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, "OK")
			return
		}

		if invoice.Lines == nil {
			log.Error("Stripe invoice lines is nil %s", invoice.ID)
			responses.ErrInternalServerError(w, r, "Stripe invoice lines is nil")
			return
		}

		for _, line := range invoice.Lines.Data {
			var product string
			if line.Plan == nil {
				log.Error("Stripe plan is nil in line item %s", line.ID)
				responses.ErrInternalServerError(w, r, "Stripe plan is nil in line item")
				return
			}

			product = line.Plan.Product

			if product == "" {
				log.Error("Stripe product is nil in line item %s", line.ID)
				responses.ErrInternalServerError(w, r, "Stripe product is nil in line item")
				return
			}

			// Check if this is an annual subscription
			isAnnual := scstripe.IsAnnualPriceID(line.Price.ID)

			// old pro to starter
			if product == "prod_NDpntRHZ5BK7jJ" {
				product = "prod_NTzD6l0KByWfLm"
			}

			// Get user from customer ID
			user, err := c.Repo.GetUserByStripeCustomerId(invoice.Customer)
			if err != nil {
				log.Error("Unable getting user from stripe customer id", "err", err)
				responses.ErrInternalServerError(w, r, err.Error())
				return
			} else if user == nil {
				log.Error("User does not exist with stripe customer id: %s", invoice.Customer)
				responses.ErrInternalServerError(w, r, "User does not exist with stripe customer id")
				return
			}

			// Get the credit type for this plan
			creditType, err := c.Repo.GetCreditTypeByStripeProductID(product)
			if err != nil {
				log.Error("Unable getting credit type from stripe product id", "err", err)
				responses.ErrInternalServerError(w, r, err.Error())
				return
			} else if creditType == nil {
				log.Error("Credit type does not exist with stripe product id: %s", line.Plan.Product)
				responses.ErrInternalServerError(w, r, "Credit type does not exist with stripe product id")
				return
			}

			expiresAt := utils.SecondsSinceEpochToTime(line.Period.End)
			// Update user credit
			if err := c.Repo.WithTx(func(tx *ent.Tx) error {
				client := tx.Client()
				added, err := c.Repo.AddCreditsIfEligible(creditType, user.ID, expiresAt, isAnnual, line.ID, client)
				if err != nil {
					log.Error("Unable adding credits to user %s: %v", user.ID.String(), err)
					return err
				}
				if user.ActiveProductID == nil && added {
					// Set a key in redis indicating we should track, to check later
					err = c.Redis.Client.SetEx(c.Redis.Ctx, invoice.ID, user.ID, time.Minute*60).Err()
					if err != nil {
						log.Error("Unable setting redis key for user %s: %v", user.ID.String(), err)
					}
					go func() {
						// See if key exists in redis still and notify
						time.Sleep(time.Minute * 5)
						_, err := c.Redis.Client.Get(c.Redis.Ctx, invoice.ID).Result()
						if err == redis.Nil || err != nil {
							return
						}
						// Remove key
						err = c.Redis.Client.Del(c.Redis.Ctx, invoice.ID).Err()
						if err != nil {
							log.Error("Unable deleting redis key for user %s: %v", user.ID.String(), err)
						}
						// Notify
						c.Track.Subscription(user, product)
						discord.NewSubscriberWebhook(c.Repo, user, product)
					}()
				} else if added && event.Type == "invoice.paid" {
					// Renewal
					go c.Track.SubscriptionRenewal(user, product)
				} else {
					// Probably already added
					return nil
				}
				err = c.Repo.SetActiveProductID(user.ID, product, client)
				if err != nil {
					log.Error("Unable setting stripe product id for user %s: %v", user.ID.String(), err)
					return err
				}
				return nil
			}); err != nil {
				log.Error("Unable adding credits to user %s: %v", user.ID.String(), err)
				if ent.IsConstraintError(err) {
					// Ignore
					render.Status(r, http.StatusOK)
					render.PlainText(w, r, "OK")
					return
				}
				responses.ErrInternalServerError(w, r, err.Error())
				return
			}
		}
	// Revoke
	case "payment_intent.canceled", "payment_intent.payment_failed":
		pi, err := stripeObjectMapToPaymentIntent(event.Data.Object)
		if err != nil || pi == nil {
			log.Error("Unable parsing stripe payment intent object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		if pi == nil || pi.Invoice != nil {
			// Not an adhoc payment
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, "OK")
			return
		}
		// Get product from metadata
		_, ok := pi.Metadata["product"]
		if !ok {
			log.Error("Stripe payment intent metadata is missing product", "payment_intent_id", pi.ID)
			responses.ErrInternalServerError(w, r, "Stripe payment intent metadata is missing product")
			return
		}

		// Remove credits
		err = c.Repo.DeleteCreditsWithLineItemID(pi.ID)
		if err != nil {
			log.Error("Unable deleting credits with stripe line item id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}

	// Adhoc credit purchases
	case "payment_intent.succeeded", "payment_intent.processing":
		pi, err := stripeObjectMapToPaymentIntent(event.Data.Object)
		if err != nil || pi == nil {
			log.Error("Unable parsing stripe payment intent object", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		if pi == nil || pi.Invoice != nil {
			// Not an adhoc payment
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, "OK")
			return
		}

		// Get product from metadata
		product, ok := pi.Metadata["product"]
		if !ok {
			log.Error("Stripe payment intent metadata is missing product", "payment_intent_id", pi.ID)
			responses.ErrInternalServerError(w, r, "Stripe payment intent metadata is missing product")
			return
		}

		// Get the credit type for this plan
		creditType, err := c.Repo.GetCreditTypeByStripeProductID(product)
		if err != nil {
			log.Error("Unable getting credit type from stripe product id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		} else if creditType == nil {
			log.Error("Credit type does not exist with stripe product id: %s", product)
			responses.ErrInternalServerError(w, r, "Credit type does not exist with stripe product id")
			return
		}

		// Get user by customer id
		user, err := c.Repo.GetUserByStripeCustomerId(pi.Customer)
		if err != nil {
			log.Error("Unable getting user from stripe customer id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		} else if user == nil {
			log.Error("User does not exist with stripe customer id: %s", pi.Customer)
			responses.ErrInternalServerError(w, r, "User does not exist with stripe customer id")
			return
		}

		// Ad-hoc credit add
		added, err := c.Repo.AddAdhocCreditsIfEligible(creditType, user.ID, pi.ID)
		if err != nil {
			log.Error("Unable adding credits to user %s: %v", user.ID.String(), err)
			if ent.IsConstraintError(err) {
				// Ignore
				render.Status(r, http.StatusOK)
				render.PlainText(w, r, "OK")
				return
			}
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		if added {
			go c.Track.CreditPurchase(user, product, int(creditType.Amount))
			go discord.AdhocCreditsPurchasedWebhook(c.Repo, user, creditType)
		}
	}

	render.Status(r, http.StatusOK)
	render.PlainText(w, r, "OK")
}

func (c *RestAPI) RevertCreditsInvoice(invoice *Invoice, w http.ResponseWriter, r *http.Request) {
	// We only care about renewal (cycle), create, and manual
	if invoice.BillingReason != InvoiceBillingReasonSubscriptionCycle && invoice.BillingReason != InvoiceBillingReasonSubscriptionCreate {
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "OK")
		return
	}

	if invoice.Lines == nil {
		log.Error("Stripe invoice lines is nil %s", invoice.ID)
		responses.ErrInternalServerError(w, r, "Stripe invoice lines is nil")
		return
	}

	u, err := c.Repo.GetUserByStripeCustomerId(invoice.Customer)
	if err != nil {
		log.Error("Unable getting user from stripe customer id", "err", err)
		responses.ErrInternalServerError(w, r, err.Error())
		return
	} else if u == nil {
		log.Error("User does not exist with stripe customer id: %s", invoice.Customer)
		responses.ErrInternalServerError(w, r, "User does not exist with stripe customer id")
		return
	}

	for _, line := range invoice.Lines.Data {
		var product string
		if line.Plan == nil {
			log.Error("Stripe plan is nil in line item %s", line.ID)
			responses.ErrInternalServerError(w, r, "Stripe plan is nil in line item")
			return
		}

		product = line.Plan.Product

		if product == "" {
			log.Error("Stripe product is nil in line item %s", line.ID)
			responses.ErrInternalServerError(w, r, "Stripe product is nil in line item")
			return
		}

		err = c.Repo.DeleteCreditsWithLineItemID(line.ID)
		if err != nil {
			log.Error("Unable deleting credits with line item id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}

		_, err := c.Repo.UnsetActiveProductID(u.ID, product, nil)
		if err != nil {
			log.Error("Unable unsetting stripe product id", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
	}
	// Remove from redis
	err = c.Redis.Client.Del(c.Redis.Ctx, invoice.ID).Err()
}

// Parse generic object into stripe invoice struct
func stripeObjectMapToInvoiceObject(obj map[string]interface{}) (*Invoice, error) {
	marshalled, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var invoice Invoice
	err = json.Unmarshal(marshalled, &invoice)
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func extractStripeSubscriptionInfoFromCustomer(customer *stripe.Customer) (string, string, *time.Time, *time.Time) {
	now := time.Now().UnixNano() / int64(time.Second)

	var highestProductID string
	var highestPriceID string
	var cancelsAt *time.Time
	var renewsAt *time.Time

	if customer != nil && customer.Subscriptions != nil && customer.Subscriptions.Data != nil {
		// Find highest subscription tier
		for _, subscription := range customer.Subscriptions.Data {
			if subscription.Items == nil || subscription.Items.Data == nil {
				continue
			}

			for _, item := range subscription.Items.Data {
				if item.Price == nil || item.Price.Product == nil {
					continue
				}
				// Not expired or cancelled
				if now > subscription.CurrentPeriodEnd || subscription.CanceledAt > subscription.CurrentPeriodEnd {
					continue
				}
				highestPriceID = item.Price.ID
				highestProductID = item.Price.Product.ID
				// If not scheduled to be cancelled, we are done
				if !subscription.CancelAtPeriodEnd {
					cancelsAt = nil
					break
				}
				cancelsAsTime := utils.SecondsSinceEpochToTime(subscription.CancelAt)
				cancelsAt = &cancelsAsTime
			}
			if cancelsAt == nil && highestProductID != "" {
				renewsAtTime := utils.SecondsSinceEpochToTime(subscription.CurrentPeriodEnd)
				renewsAt = &renewsAtTime
				break
			}
		}
	}
	return highestProductID, highestPriceID, cancelsAt, renewsAt
}

// Parse generic object into stripe subscription struct
func stripeObjectMapToSubscriptionObject(obj map[string]interface{}) (*stripe.Subscription, error) {
	marshalled, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var subscription stripe.Subscription
	err = json.Unmarshal(marshalled, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

// Parse generic object into custom stripe subscription struct with correct types
func stripeObjectMapToCustomSubscriptionObject(obj map[string]interface{}) (*Subscription, error) {
	marshalled, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var subscription Subscription
	err = json.Unmarshal(marshalled, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

// Parse generic object into stripe invoice struct
func stripeObjectMapToPaymentIntent(obj map[string]interface{}) (*PaymentIntent, error) {
	marshalled, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var pi PaymentIntent
	err = json.Unmarshal(marshalled, &pi)
	if err != nil {
		return nil, err
	}
	return &pi, nil
}

// ! Stripe types are busted so we modify the ones included in their lib
// InvoiceBillingReason is the reason why a given invoice was created
type InvoiceBillingReason string

// List of values that InvoiceBillingReason can take.
const (
	InvoiceBillingReasonSubscription          InvoiceBillingReason = "subscription"
	InvoiceBillingReasonSubscriptionCreate    InvoiceBillingReason = "subscription_create"
	InvoiceBillingReasonSubscriptionCycle     InvoiceBillingReason = "subscription_cycle"
	InvoiceBillingReasonSubscriptionThreshold InvoiceBillingReason = "subscription_threshold"
	InvoiceBillingReasonSubscriptionUpdate    InvoiceBillingReason = "subscription_update"
	InvoiceBillingReasonUpcoming              InvoiceBillingReason = "upcoming"
)

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	HasMore    bool   `json:"has_more"`
	TotalCount uint32 `json:"total_count"`
	URL        string `json:"url"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type Plan struct {
	Product string `json:"product"`
}

type Price struct {
	ID      string `json:"id"`
	Product string `json:"product"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	ID     string  `json:"id"`
	Period *Period `json:"period"`
	Plan   *Plan   `json:"plan"`
	Price  *Price  `json:"price"`
}

type InvoiceLineList struct {
	ListMeta
	Data []*InvoiceLine `json:"data"`
}

type InvoiceStatus string

const (
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusUncollectible InvoiceStatus = "uncollectible"
	InvoiceStatusVoid          InvoiceStatus = "void"
)

type Invoice struct {
	ID            string               `json:"id"`
	BillingReason InvoiceBillingReason `json:"billing_reason"`
	Lines         *InvoiceLineList     `json:"lines"`
	Customer      string               `json:"customer"`
	PaymentIntent string               `json:"payment_intent"`
	Status        InvoiceStatus        `json:"status"`
	Currency      string               `json:"currency"`
}

// Subscription object is also pbroken in stripe
type SubscriptionItem struct {
	Price *Price `json:"price"`
}
type SubscriptionItemList struct {
	Data []*SubscriptionItem `json:"data"`
}
type Subscription struct {
	Items    *SubscriptionItemList `json:"items"`
	Customer string                `json:"customer"`
}

// PaymentIntent is also broken
type PaymentIntent struct {
	ID       string            `json:"id"`
	Invoice  *string           `json:"invoice,omitempty"`
	Metadata map[string]string `json:"metadata"`
	Customer string            `json:"customer"`
}
