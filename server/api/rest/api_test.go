// Test setup for the controller package
package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stablecog/sc-go/database"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/server/analytics"
	"github.com/stablecog/sc-go/server/api/sse"
	"github.com/stablecog/sc-go/server/scworker"
	"github.com/stablecog/sc-go/server/translator"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/shared/queue"
	"github.com/stablecog/sc-go/utils"
)

// A valid sse stream ID that will be acceptable by APIs
const MockSSEId = "e08abf9698f7d27e634de0d36cc974a0d908ec41c0a7e5e5738d2431f9a700e3"

var MockController *RestAPI

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	ctx := context.Background()
	dbconn, err := database.GetSqlDbConn(!utils.GetEnv().GithubActions)
	if err != nil {
		log.Fatal("Failed to connect to database", "err", err)
		os.Exit(1)
	}
	entClient, err := database.NewEntClient(dbconn)
	defer entClient.Close()
	if err != nil {
		log.Fatal("Failed to create ent client", "err", err)
		os.Exit(1)
	}

	// Redis setup
	origMockRedis := utils.GetEnv().MockRedis
	utils.GetEnv().MockRedis = true
	defer func() {
		utils.GetEnv().MockRedis = origMockRedis
	}()

	redis, err := database.NewRedis(ctx)
	if err != nil {
		log.Fatal("Error connecting to redis", "err", err)
		os.Exit(1)
	}

	//Create schema
	if err := entClient.Schema.Create(ctx); err != nil {
		log.Fatal("Failed to run migrations", "err", err)
		os.Exit(1)
	}

	qThrottler := shared.NewQueueThrottler(ctx, redis.Client, time.Hour)

	repo := &repository.Repository{
		DB:             entClient,
		ConnInfo:       dbconn,
		Redis:          redis,
		Ctx:            ctx,
		QueueThrottler: qThrottler,
	}

	// Mock data
	if err := repo.CreateMockData(ctx); err != nil {
		log.Fatal("Failed to create mock data", "err", err)
		os.Exit(1)
	}

	// Populate cache
	os.Setenv("SKIP_VAST", "true")
	defer os.Unsetenv("SKIP_VAST")
	if err := repo.UpdateCache(); err != nil {
		log.Fatal("Failed to populate cache", "err", err)
		os.Exit(1)
	}

	// Setup fake sse hub
	hub := sse.NewHub(redis, repo)
	go hub.Run()
	// Add user to hub
	hub.Register <- &sse.Client{
		Uid:  MockSSEId,
		Send: make(chan []byte, 256),
	}

	// Create mock client
	mockClient := &queue.MockRabbitMQClient{
		PublishFunc: func(id string, msg any, priority uint8) error {
			return nil
		},
	}

	// Setup controller
	MockController = &RestAPI{
		Client:         &http.Client{Timeout: 10 * time.Second},
		Repo:           repo,
		Redis:          redis,
		Hub:            hub,
		Track:          analytics.NewAnalyticsService(),
		QueueThrottler: qThrottler,
		SafetyChecker:  translator.NewTranslatorSafetyChecker(ctx, "", true, redis),
		SCWorker: &scworker.SCWorker{
			Repo:           repo,
			Redis:          redis,
			QueueThrottler: qThrottler,
			Track:          analytics.NewAnalyticsService(),
			SafetyChecker:  translator.NewTranslatorSafetyChecker(ctx, "", true, redis),
			MQClient:       mockClient,
		},
	}

	return m.Run()
}
