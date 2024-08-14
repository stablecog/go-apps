package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/qdrant"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
	"github.com/stripe/stripe-go/v74"
)

// HTTP Get - user info
func (c *RestAPI) HandleGetUserV2(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	m := time.Now()

	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	log.Infof("HandleGetUserV2 - GetUserIDAndEmailIfAuthenticated: %dms", time.Since(m).Milliseconds())

	if userID == nil || email == "" {
		return
	}

	var lastSignIn *time.Time
	lastSignInStr, ok := r.Context().Value("user_last_sign_in").(string)
	if ok {
		lastSignInP, err := time.Parse(time.RFC3339, lastSignInStr)
		if err == nil {
			lastSignIn = &lastSignInP
		}
	}

	// Get user with roles
	m = time.Now()
	user, err := c.Repo.GetUserWithRoles(*userID)
	log.Infof("HandleGetUserV2 - GetUserWithRoles: %dms", time.Since(m).Milliseconds())

	if err != nil {
		log.Error("Error getting user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	} else if user == nil {
		// Handle create user flow
		err := createNewUser(email, userID, lastSignIn, c)
		if err != nil {
			log.Error("Error creating user", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		go c.Track.SignUp(*userID, email, utils.GetIPAddress(r), utils.GetClientDeviceInfo(r))

		user, err = c.Repo.GetUserWithRoles(*userID)
		if err != nil {
			log.Error("Error getting user with roles", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
	}

	type result struct {
		totalRemaining     int
		customer           *stripe.Customer
		paidCreditCount    int
		err                error
		stripeHadError     bool
		updateLastSeenErr  error
		paymentsMadeByUser int
		duration           time.Duration
		operation          string
	}

	ch := make(chan result, 5)
	var wg sync.WaitGroup

	wg.Add(5)

	// Get total credits
	go func() {
		defer wg.Done()
		m := time.Now()
		totalRemaining, err := c.Repo.GetNonExpiredCreditTotalForUser(*userID, nil)
		ch <- result{totalRemaining: totalRemaining, err: err, duration: time.Since(m), operation: "GetNonExpiredCreditTotalForUser"}
	}()

	// Get customer from Stripe
	go func() {
		defer wg.Done()
		m := time.Now()
		customer, err := c.StripeClient.Customers.Get(user.StripeCustomerID, &stripe.CustomerParams{
			Params: stripe.Params{
				Expand: []*string{
					stripe.String("subscriptions"),
				},
			},
		})
		stripeHadError := err != nil
		ch <- result{customer: customer, stripeHadError: stripeHadError, duration: time.Since(m), operation: "GetStripeCustomer"}
	}()

	// Get paid credits
	go func() {
		defer wg.Done()
		m := time.Now()
		paidCreditCount, err := c.Repo.GetNonFreeCreditSum(*userID)
		ch <- result{paidCreditCount: paidCreditCount, err: err, duration: time.Since(m), operation: "GetNonFreeCreditSum"}
	}()

	// Update last seen
	go func() {
		defer wg.Done()
		m := time.Now()
		err := c.Repo.UpdateLastSeenAt(*userID)
		ch <- result{updateLastSeenErr: err, duration: time.Since(m), operation: "UpdateLastSeenAt"}
	}()

	// Get payments made by customer
	go func() {
		defer wg.Done()
		m := time.Now()
		paymentsMade := getPaymentsMadeByCustomer(user.StripeCustomerID, c)
		ch <- result{paymentsMadeByUser: paymentsMade, duration: time.Since(m), operation: "GetPaymentMadeByCustomer"}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res result
	for goroutineResult := range ch {
		if goroutineResult.err != nil {
			log.Error("Error in goroutine", "err", goroutineResult.err, "operation", goroutineResult.operation)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		log.Infof("HandleGetUserV2 - %s: %dms", goroutineResult.operation, goroutineResult.duration.Milliseconds())
		if goroutineResult.totalRemaining != 0 {
			res.totalRemaining = goroutineResult.totalRemaining
		}
		if goroutineResult.customer != nil {
			res.customer = goroutineResult.customer
		}
		if goroutineResult.paidCreditCount != 0 {
			res.paidCreditCount = goroutineResult.paidCreditCount
		}
		if goroutineResult.stripeHadError {
			res.stripeHadError = true
		}
		if goroutineResult.updateLastSeenErr != nil {
			log.Warn("Error updating last seen at", "err", goroutineResult.updateLastSeenErr, "user", userID.String())
		}
		if goroutineResult.paymentsMadeByUser != 0 {
			res.paymentsMadeByUser = goroutineResult.paymentsMadeByUser
		}
	}

	m = time.Now()
	highestProduct, highestPrice, cancelsAt, renewsAt := extractSubscriptionInfoFromCustomer(res.customer)
	log.Infof("HandleGetUserV2 - extractSubscriptionInfoFromCustomer: %dms", time.Since(m).Milliseconds())

	m = time.Now()
	moreCreditsAt, moreCreditsAtAmount, renewsAtAmount, freeCreditAmount := getMoreCreditsInfo(*userID, highestProduct, renewsAt, res.stripeHadError, c)
	log.Infof("HandleGetUserV2 - getMoreCreditsInfo: %dms", time.Since(m).Milliseconds())

	roles := make([]string, len(user.Edges.Roles))
	for i, role := range user.Edges.Roles {
		roles[i] = role.Name
	}

	log.Infof("HandleGetUserV2 - Total: %dms", time.Since(s).Milliseconds())

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responses.GetUserResponse{
		UserID:                  userID,
		TotalRemainingCredits:   res.totalRemaining,
		HasNonfreeCredits:       res.paidCreditCount > 0,
		ProductID:               highestProduct,
		PriceID:                 highestPrice,
		CancelsAt:               cancelsAt,
		RenewsAt:                renewsAt,
		RenewsAtAmount:          renewsAtAmount,
		FreeCreditAmount:        freeCreditAmount,
		StripeHadError:          res.stripeHadError,
		Roles:                   roles,
		MoreCreditsAt:           moreCreditsAt,
		MoreCreditsAtAmount:     moreCreditsAtAmount,
		MoreFreeCreditsAt:       moreCreditsAt,
		MoreFreeCreditsAtAmount: moreCreditsAtAmount,
		WantsEmail:              user.WantsEmail,
		Username:                user.Username,
		CreatedAt:               user.CreatedAt,
		UsernameChangedAt:       user.UsernameChangedAt,
		PurchaseCount:           res.paymentsMadeByUser,
	})
}

func (c *RestAPI) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	m := time.Now()

	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	log.Infof("HandleGetUser - GetUserIDAndEmailIfAuthenticated: %dms", time.Since(m).Milliseconds())

	if userID == nil || email == "" {
		return
	}
	var lastSignIn *time.Time
	lastSignInStr, ok := r.Context().Value("user_last_sign_in").(string)
	if ok {
		lastSignInP, err := time.Parse(time.RFC3339, lastSignInStr)
		if err == nil {
			lastSignIn = &lastSignInP
		}
	}

	// Get customer ID for user
	m = time.Now()
	user, err := c.Repo.GetUserWithRoles(*userID)
	log.Infof("HandleGetUser - GetUserWithRoles: %dms", time.Since(m).Milliseconds())

	if err != nil {
		log.Error("Error getting user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	} else if user == nil {
		// Handle create user flow
		err := createNewUser(email, userID, lastSignIn, c)
		if err != nil {
			log.Error("Error creating user", "err", err)
			responses.ErrInternalServerError(w, r, err.Error())
			return
		}
		go c.Track.SignUp(*userID, email, utils.GetIPAddress(r), utils.GetClientDeviceInfo(r))
	}

	if user == nil {
		user, err = c.Repo.GetUserWithRoles(*userID)
		if err != nil {
			log.Error("Error getting user with roles", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
	}

	// Get total credits
	m = time.Now()
	totalRemaining, err := c.Repo.GetNonExpiredCreditTotalForUser(*userID, nil)
	if err != nil {
		log.Error("Error getting credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}
	log.Infof("HandleGetUser - GetNonExpiredCreditTotalForUser: %dms", time.Since(m).Milliseconds())

	m = time.Now()
	customer, err := c.StripeClient.Customers.Get(user.StripeCustomerID, &stripe.CustomerParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("subscriptions"),
			},
		},
	})
	stripeHadError := false
	if err != nil {
		log.Error("Error getting customer from stripe, unknown error", "err", err)
		stripeHadError = true
	}
	log.Infof("HandleGetUser - GetStripeCustomer: %dms", time.Since(m).Milliseconds())

	// Get subscription info
	highestProduct, highestPrice, cancelsAt, renewsAt := extractSubscriptionInfoFromCustomer(customer)

	m = time.Now()
	err = c.Repo.UpdateLastSeenAt(*userID)
	if err != nil {
		log.Warn("Error updating last seen at", "err", err, "user", userID.String())
	}
	log.Infof("HandleGetUser - UpdateLastSeenAt: %dms", time.Since(m).Milliseconds())

	// Figure out when free credits will be replenished
	m = time.Now()
	moreCreditsAt, moreCreditsAtAmount, renewsAtAmount, freeCreditAmount := getMoreCreditsInfo(*userID, highestProduct, renewsAt, stripeHadError, c)
	log.Infof("HandleGetUser - getMoreCreditsInfo: %dms", time.Since(m).Milliseconds())

	// Get paid credits for user
	m = time.Now()
	paidCreditCount, err := c.Repo.GetNonFreeCreditSum(*userID)
	if err != nil {
		log.Error("Error getting paid credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}
	log.Infof("HandleGetUser - GetNonFreeCreditSum: %dms", time.Since(m).Milliseconds())

	roles := make([]string, len(user.Edges.Roles))
	for i, role := range user.Edges.Roles {
		roles[i] = role.Name
	}

	m = time.Now()
	paymentsMadeByCustomer := getPaymentsMadeByCustomer(user.StripeCustomerID, c)
	log.Infof("HandleGetUser - GetPaymentMadeByCustomer: %dms", time.Since(m).Milliseconds())

	log.Infof("HandleGetUser - Total: %dms", time.Since(s).Milliseconds())

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responses.GetUserResponse{
		UserID:                  userID,
		TotalRemainingCredits:   totalRemaining,
		HasNonfreeCredits:       paidCreditCount > 0,
		ProductID:               highestProduct,
		PriceID:                 highestPrice,
		CancelsAt:               cancelsAt,
		RenewsAt:                renewsAt,
		RenewsAtAmount:          renewsAtAmount,
		FreeCreditAmount:        freeCreditAmount,
		StripeHadError:          stripeHadError,
		Roles:                   roles,
		MoreCreditsAt:           moreCreditsAt,
		MoreCreditsAtAmount:     moreCreditsAtAmount,
		MoreFreeCreditsAt:       moreCreditsAt,
		MoreFreeCreditsAtAmount: moreCreditsAtAmount,
		WantsEmail:              user.WantsEmail,
		Username:                user.Username,
		CreatedAt:               user.CreatedAt,
		UsernameChangedAt:       user.UsernameChangedAt,
		PurchaseCount:           paymentsMadeByCustomer,
	})
}

func getMoreCreditsInfo(userID uuid.UUID, highestProduct string, renewsAt *time.Time, stripeHadError bool, c *RestAPI) (*time.Time, *int, *int, *int) {
	var moreCreditsAt *time.Time
	var moreCreditsAtAmount *int
	var renewsAtAmount *int
	var fcredit *ent.Credit
	var ctype *ent.CreditType
	var freeCreditAmount *int
	var err error
	if highestProduct == "" && !stripeHadError {
		moreCreditsAt, fcredit, ctype, err = c.Repo.GetFreeCreditReplenishesAtForUser(userID)
		if err != nil {
			log.Error("Error getting next free credit replenishment time", "err", err, "user", userID.String())
		}
		moreCreditsAtAmount = utils.ToPtr(shared.FREE_CREDIT_AMOUNT_DAILY)

		if fcredit != nil && ctype != nil {
			if shared.FREE_CREDIT_AMOUNT_DAILY+fcredit.RemainingAmount > ctype.Amount {
				am := int(shared.FREE_CREDIT_AMOUNT_DAILY + fcredit.RemainingAmount - ctype.Amount)
				freeCreditAmount = &am
			} else {
				am := shared.FREE_CREDIT_AMOUNT_DAILY
				freeCreditAmount = &am
			}
		}
	} else if !stripeHadError && renewsAt != nil {
		creditType, err := c.Repo.GetCreditTypeByStripeProductID(highestProduct)
		if err != nil {
			log.Warnf("Error getting credit type from product id '%s' %v", highestProduct, err)
		} else {
			renewsAtAmount = utils.ToPtr(int(creditType.Amount))
		}
	}
	return moreCreditsAt, moreCreditsAtAmount, renewsAtAmount, freeCreditAmount
}

func createNewUser(email string, userID *uuid.UUID, lastSignIn *time.Time, c *RestAPI) error {
	unknownError := errors.New("An unknown error has occurred")
	freeCreditType, err := c.Repo.GetOrCreateFreeCreditType(nil)
	if err != nil {
		log.Error("Error getting free credit type", "err", err)
		return unknownError
	}
	if freeCreditType == nil {
		log.Error("Server misconfiguration: a credit_type with the name 'free' must exist")
		return unknownError
	}
	tippableCreditType, err := c.Repo.GetOrCreateTippableCreditType(nil)
	if err != nil {
		log.Error("Error getting tippable credit type", "err", err)
		return unknownError
	}
	if tippableCreditType == nil {
		log.Error("Server misconfiguration: a credit_type with the name 'tippable' must exist")
		return unknownError
	}

	// See if email exists
	_, exists, err := c.Repo.CheckIfEmailExists(email)
	if err != nil {
		log.Error("Error checking if email exists", "err", err)
		return unknownError
	} else if exists {
		log.Error("Email already exists", "email", email)
		return errors.New("Email already exists")
	}

	var customer *stripe.Customer
	if err := c.Repo.WithTx(func(tx *ent.Tx) error {
		client := tx.Client()

		customer, err = c.StripeClient.Customers.New(&stripe.CustomerParams{
			Email: stripe.String(email),
			Params: stripe.Params{
				Metadata: map[string]string{
					"supabase_id": (*userID).String(),
				},
			},
		})
		if err != nil {
			log.Error("Error creating stripe customer", "err", err)
			return err
		}

		u, err := c.Repo.CreateUser(*userID, email, customer.ID, lastSignIn, client)
		if err != nil {
			log.Error("Error creating user", "err", err)
			return err
		}

		// Add free credits
		added, err := c.Repo.GiveFreeCredits(u.ID, client)
		if err != nil || !added {
			log.Error("Error adding free credits", "err", err)
			return err
		}

		// Add free tippable credits
		added, err = c.Repo.GiveFreeTippableCredits(u.ID, client)
		if err != nil || !added {
			log.Error("Error adding free tippable credits", "err", err)
			return err
		}

		return nil
	}); err != nil {
		log.Error("Error creating user", "err", err)
		// Delete stripe customer
		if customer != nil {
			_, err := c.StripeClient.Customers.Del(customer.ID, nil)
			if err != nil {
				log.Error("Error deleting stripe customer", "err", err)
			}
		}
		return unknownError
	}
	return nil
}

func getPaymentsMadeByCustomer(customerId string, c *RestAPI) int {
	paymentsMadeByCustomer := 0
	paymentIntents := c.StripeClient.PaymentIntents.List(&stripe.PaymentIntentListParams{
		Customer: stripe.String(customerId),
	})
	for paymentIntents.Next() {
		intent := paymentIntents.PaymentIntent()
		if intent != nil && intent.Status == stripe.PaymentIntentStatusSucceeded {
			paymentsMadeByCustomer++
		}
	}
	return paymentsMadeByCustomer
}

func extractSubscriptionInfoFromCustomer(customer *stripe.Customer) (string, string, *time.Time, *time.Time) {
	now := time.Now().UnixNano() / int64(time.Second)

	var highestProduct string
	var highestPrice string
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
				highestPrice = item.Price.ID
				highestProduct = item.Price.Product.ID
				// If not scheduled to be cancelled, we are done
				if !subscription.CancelAtPeriodEnd {
					cancelsAt = nil
					break
				}
				cancelsAsTime := utils.SecondsSinceEpochToTime(subscription.CancelAt)
				cancelsAt = &cancelsAsTime
			}
			if cancelsAt == nil && highestProduct != "" {
				renewsAtTime := utils.SecondsSinceEpochToTime(subscription.CurrentPeriodEnd)
				renewsAt = &renewsAtTime
				break
			}
		}
	}
	return highestProduct, highestPrice, cancelsAt, renewsAt
}

// HTTP Get - generations for user
// Takes query paramers for pagination
// per_page: number of generations to return
// cursor: cursor for pagination, it is an iso time string in UTC
func (c *RestAPI) HandleQueryGenerations(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	// Validate query parameters
	perPage := DEFAULT_PER_PAGE
	var err error
	if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil {
			responses.ErrBadRequest(w, r, "per_page must be an integer", "")
			return
		} else if perPage < 1 || perPage > MAX_PER_PAGE {
			responses.ErrBadRequest(w, r, fmt.Sprintf("per_page must be between 1 and %d", MAX_PER_PAGE), "")
			return
		}
	}

	cursorStr := r.URL.Query().Get("cursor")
	search := r.URL.Query().Get("search")

	filters := &requests.QueryGenerationFilters{}
	err = filters.ParseURLQueryParameters(r.URL.Query())
	if err != nil {
		responses.ErrBadRequest(w, r, err.Error(), "")
		return
	}

	// For search, use qdrant semantic search
	if search != "" {
		// get embeddings from clip service
		e, err := c.Clip.GetEmbeddingFromText(search, true)
		if err != nil {
			log.Error("Error getting embedding from clip service", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}

		// Parse as qdrant filters
		qdrantFilters, scoreThreshold := filters.ToQdrantFilters(false)
		// Append user_id requirement, unless liked
		if filters.IsLiked == nil {
			qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
				Key:   "user_id",
				Match: &qdrant.SCValue{Value: user.ID.String()},
			})
		} else {
			// Get this users likes
			likedIds, err := c.Repo.GetGenerationOutputIDsLikedByUser(user.ID, 10000)
			if err != nil {
				log.Error("Error getting liked ids", "err", err)
				responses.ErrInternalServerError(w, r, "An unknown error has occurred")
				return
			}
			qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
				HasId: likedIds,
			})
		}
		// Deleted at not empty
		qdrantFilters.Must = append(qdrantFilters.Must, qdrant.SCMatchCondition{
			IsEmpty: &qdrant.SCIsEmpty{Key: "deleted_at"},
		})

		// Get cursor str as uint
		var offset *uint
		var total *uint
		if cursorStr != "" {
			cursoru64, err := strconv.ParseUint(cursorStr, 10, 64)
			if err != nil {
				responses.ErrBadRequest(w, r, "cursor must be a valid uint", "")
				return
			}
			cursoru := uint(cursoru64)
			offset = &cursoru
		} else {
			count, err := c.Qdrant.CountWithFilters(qdrantFilters, false)
			if err != nil {
				log.Error("Error counting qdrant", "err", err)
				responses.ErrInternalServerError(w, r, "An unknown error has occurred")
				return
			}
			total = &count
		}

		// Query qdrant
		qdrantRes, err := c.Qdrant.QueryGenerations(e, perPage, offset, scoreThreshold, filters.Oversampling, qdrantFilters, false, false)
		if err != nil {
			log.Error("Error querying qdrant", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}

		// Get generation output ids
		var outputIds []uuid.UUID
		for _, hit := range qdrantRes.Result {
			outputId, err := uuid.Parse(hit.Id)
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			outputIds = append(outputIds, outputId)
		}

		// Get user generation data in correct format
		generationsUnsorted, err := c.Repo.RetrieveGalleryDataWithOutputIDs(outputIds, utils.ToPtr(user.ID), repository.GalleryDataFromHistory)
		if err != nil {
			log.Error("Error getting generations", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}

		gDataMap := make(map[uuid.UUID]repository.GalleryData)
		for _, gData := range generationsUnsorted {
			gDataMap[gData.ID] = gData
		}
		generationsSorted := make([]repository.GalleryData, len(qdrantRes.Result))

		for i, hit := range qdrantRes.Result {
			outputId, err := uuid.Parse(hit.Id)
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			item, ok := gDataMap[outputId]
			if !ok {
				log.Error("Error retrieving gallery data", "output_id", outputId)
				continue
			}
			generationsSorted[i] = item
		}

		// Return generations
		render.Status(r, http.StatusOK)
		render.JSON(w, r, GalleryResponseV3[*uint]{
			Next:    qdrantRes.Next,
			Outputs: c.Repo.ConvertRawGalleryDataToV3Results(generationsSorted),
			Total:   total,
		})
		return
	}

	// Otherwise, query postgres
	var cursor *time.Time
	if cursorStr := r.URL.Query().Get("cursor"); cursorStr != "" {
		cursorTime, err := utils.ParseIsoTime(cursorStr)
		if err != nil {
			responses.ErrBadRequest(w, r, "cursor must be a valid iso time string", "")
			return
		}
		cursor = &cursorTime
	}

	// Ensure user ID is set to only include this users generations
	filters.UserID = &user.ID
	filters.ForHistory = true

	// Test flag
	generations, nextCursor, _, err := c.Repo.RetrieveMostRecentGalleryDataV3(filters, filters.UserID, perPage, cursor, nil)
	if err != nil {
		log.Error("Error getting generations for user", "err", err)
		responses.ErrInternalServerError(w, r, "Error getting generations")
		return
	}

	// Presign init image URLs
	signedMap := make(map[string]string)
	for _, g := range generations {
		if g.InitImageURL != nil {
			// See if we have already signed this URL
			signedInitImageUrl, ok := signedMap[*g.InitImageURL]
			if !ok {
				g.InitImageURLSigned = &signedInitImageUrl
				continue
			}
			// remove s3:// prefix
			if strings.HasPrefix(*g.InitImageURL, "s3://") {
				prefixRemoved := (*g.InitImageURL)[5:]
				// Sign object URL to pass to worker
				req, _ := c.S3.GetObjectRequest(&s3.GetObjectInput{
					Bucket: aws.String(utils.GetEnv().S3Img2ImgBucketName),
					Key:    aws.String(prefixRemoved),
				})
				urlStr, err := req.Presign(1 * time.Hour)
				if err != nil {
					log.Error("Error signing init image URL", "err", err)
					continue
				}
				// Add to map
				signedMap[*g.InitImageURL] = urlStr
				g.InitImageURLSigned = &urlStr
			}
		}
	}

	// Get total if no cursor
	var total *uint
	if cursor == nil {
		totalI, err := c.Repo.GetGenerationCount(filters)
		if err != nil {
			log.Error("Error getting user generation count", "err", err)
			responses.ErrInternalServerError(w, r, "Error getting generations")
			return
		}
		// Convert int to uint
		totalUInt := uint(totalI)
		// Assign the address of the uint to the total pointer
		total = &totalUInt
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, GalleryResponseV3[*time.Time]{
		Next:    nextCursor,
		Outputs: c.Repo.ConvertRawGalleryDataToV3Results(generations),
		Total:   total,
	})
}

// HTTP Get - credits for user
func (c *RestAPI) HandleQueryCredits(w http.ResponseWriter, r *http.Request) {
	// See if authenticated
	userIDStr, authenticated := r.Context().Value("user_id").(string)
	// This should always be true because of the auth middleware, but check it anyway
	if !authenticated || userIDStr == "" {
		responses.ErrUnauthorized(w, r)
		return
	}
	// Parse to UUID
	userId, err := uuid.Parse(userIDStr)
	if err != nil {
		responses.ErrUnauthorized(w, r)
		return
	}

	// Get credits
	credits, err := c.Repo.GetCreditsForUser(userId)
	if err != nil {
		log.Error("Error getting credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "Error getting credits")
		return
	}

	// Format as a nicer response
	var totalRemaining int32
	for _, credit := range credits {
		totalRemaining += credit.RemainingAmount
	}

	creditsFormatted := make([]responses.Credit, len(credits))
	for i, credit := range credits {
		creditsFormatted[i] = responses.Credit{
			ID:              credit.ID,
			RemainingAmount: credit.RemainingAmount,
			ExpiresAt:       credit.ExpiresAt,
			Type: responses.CreditType{
				ID:          credit.CreditTypeID,
				Name:        credit.CreditTypeName,
				Description: credit.CreditTypeDescription,
				Amount:      credit.CreditTypeAmount,
			},
		}
	}

	creditsResponse := responses.QueryCreditsResponse{
		TotalRemainingCredits: totalRemaining,
		Credits:               creditsFormatted,
	}

	// Return credits
	render.Status(r, http.StatusOK)
	render.JSON(w, r, creditsResponse)
}

// HTTP DELETE - delete generation
func (c *RestAPI) HandleDeleteGenerationOutputForUser(w http.ResponseWriter, r *http.Request) {
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
	var deleteReq requests.DeleteGenerationRequest
	err := json.Unmarshal(reqBody, &deleteReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	count, err := c.Repo.MarkGenerationOutputsForDeletionForUser(deleteReq.GenerationOutputIDs, user.ID)
	if err != nil {
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	res := responses.DeletedResponse{
		Deleted: count,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP POST - favorite generation
func (c *RestAPI) HandleFavoriteGenerationOutputsForUser(w http.ResponseWriter, r *http.Request) {
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
	var favReq requests.FavoriteGenerationRequest
	err := json.Unmarshal(reqBody, &favReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	if favReq.Action != requests.AddFavoriteAction && favReq.Action != requests.RemoveFavoriteAction {
		responses.ErrBadRequest(w, r, "action must be either 'add' or 'remove'", "")
		return
	}

	count, err := c.Repo.SetFavoriteGenerationOutputsForUser(favReq.GenerationOutputIDs, user.ID, favReq.Action)
	if err != nil {
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	res := responses.FavoritedResponse{
		Favorited: count,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP DELETE - delete voiceover
func (c *RestAPI) HandleDeleteVoiceoverOutputForUser(w http.ResponseWriter, r *http.Request) {
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
	var deleteReq requests.DeleteVoiceoverRequest
	err := json.Unmarshal(reqBody, &deleteReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	count, err := c.Repo.MarkVoiceoverOutputsForDeletionForUser(deleteReq.OutputIDs, user.ID)
	if err != nil {
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	res := responses.DeletedResponse{
		Deleted: count,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP POST - set email preferences
func (c *RestAPI) HandleUpdateEmailPreferences(w http.ResponseWriter, r *http.Request) {
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
	var emailReq requests.EmailPreferencesRequest
	err := json.Unmarshal(reqBody, &emailReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	// Update email preferences
	err = c.Repo.SetWantsEmail(user.ID, emailReq.WantsEmail)
	if err != nil {
		log.Error("Error setting email preferences", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	res := responses.UpdatedResponse{
		Updated: 1,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP POST - set email preferences
func (c *RestAPI) HandleUpdateUsername(w http.ResponseWriter, r *http.Request) {
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
	var usernameReq requests.ChangeUsernameRequest
	err := json.Unmarshal(reqBody, &usernameReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	// Check if valid
	if err := utils.IsValidUsername(usernameReq.Username); err != nil {
		responses.ErrBadRequest(w, r, err.Error(), "")
		return
	}

	// Update username
	err = c.Repo.SetUsername(user.ID, usernameReq.Username)
	if err != nil {
		if errors.Is(err, repository.UsernameExistsErr) {
			responses.ErrBadRequest(w, r, "username_taken", "")
			return
		}
		log.Error("Error setting username", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"username": usernameReq.Username,
	})
}

// HTTP POST - like/unlike generation
func (c *RestAPI) HandleLikeGenerationOutputsForUser(w http.ResponseWriter, r *http.Request) {
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
	var likeReq requests.LikeUnlikeActionRequest
	err := json.Unmarshal(reqBody, &likeReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	if likeReq.Action != requests.LikeAction && likeReq.Action != requests.UnlikeAction {
		responses.ErrBadRequest(w, r, "action must be either 'like' or 'unlike'", "")
		return
	}

	err = c.Repo.SetOutputsLikedForUser(likeReq.GenerationOutputIDs, user.ID, likeReq.Action)
	// Error check required due to https://github.com/ent/ent/issues/2176
	// This shouldn't return an error if it fails, due to on conflict do nothing behavior
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Error("Error setting outputs liked for user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"success": true,
	})
}
