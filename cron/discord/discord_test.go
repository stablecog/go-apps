package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/stablecog/sc-go/cron/models"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
	"github.com/stretchr/testify/assert"
)

var MockDiscordHealthTracker *DiscordHealthTracker

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	// Setup
	oldWebhookUrl := utils.GetEnv().DiscordWebhookUrl
	utils.GetEnv().DiscordWebhookUrl = "http://localhost:123456"
	defer func() { utils.GetEnv().DiscordWebhookUrl = oldWebhookUrl }()

	MockDiscordHealthTracker = NewDiscordHealthTracker(context.Background())

	return m.Run()
}

func TestSendDiscordNotificationIfNeeded(t *testing.T) {
	// Mock logger
	orgLogInfo := logInfo
	defer func() { logInfo = orgLogInfo }()

	// Write log output to string
	logs := []string{}
	logInfo = func(format interface{}, args ...interface{}) {
		logs = append(logs, format.(string)+fmt.Sprint(args...))
	}

	// Mock generations
	generations := []*ent.Generation{}
	failedStatus := generation.StatusFailed
	startedStatus := generation.StatusStarted
	successStatus := generation.StatusSucceeded
	queuedStatus := generation.StatusQueued
	nsfw := "NSFW"
	generations = append(generations, &ent.Generation{
		ID:            uuid.New(),
		FailureReason: &nsfw,
		Status:        failedStatus,
	})
	generations = append(generations, &ent.Generation{
		ID:            uuid.New(),
		FailureReason: nil,
		Status:        failedStatus,
	})
	generations = append(generations, &ent.Generation{
		ID:            uuid.New(),
		FailureReason: nil,
		Status:        queuedStatus,
	})
	generations = append(generations, &ent.Generation{
		ID:            uuid.New(),
		FailureReason: nil,
		Status:        startedStatus,
	})
	generations = append(generations, &ent.Generation{
		ID:            uuid.New(),
		FailureReason: nil,
		Status:        successStatus,
	})

	// ! Test notification not needed
	MockDiscordHealthTracker.lastStatus = shared.UNKNOWN
	err := MockDiscordHealthTracker.SendDiscordNotificationIfNeeded(shared.HEALTHY, generations, time.Now(), time.Now(), false, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Skipping Discord notification, not needed", logs[0])

	MockDiscordHealthTracker.lastNotificationTime = time.Now()
	MockDiscordHealthTracker.lastHealthyNotificationTime = time.Now()
	MockDiscordHealthTracker.lastUnhealthyNotificationTime = time.Now()

	MockDiscordHealthTracker.lastStatus = shared.UNHEALTHY
	err = MockDiscordHealthTracker.SendDiscordNotificationIfNeeded(shared.UNHEALTHY, generations, time.Now(), time.Now(), false, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Skipping Discord notification, not needed", logs[1])

	// Reset keys
	MockDiscordHealthTracker.lastNotificationTime = time.Time{}
	MockDiscordHealthTracker.lastHealthyNotificationTime = time.Time{}
	MockDiscordHealthTracker.lastUnhealthyNotificationTime = time.Time{}

	// ! Test notification needed
	// Setup httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://localhost:123456",
		func(req *http.Request) (*http.Response, error) {
			var request models.DiscordWebhookBody
			err := json.NewDecoder(req.Body).Decode(&request)
			assert.Nil(t, err)
			assert.Equal(t, 11437547, request.Embeds[0].Color)
			assert.Equal(t, "```🟢👌🟢```", request.Embeds[0].Fields[0].Value)
			assert.Equal(t, "```🌶️🔴⏲️🟡🟢```", request.Embeds[0].Fields[1].Value)
			assert.Equal(t, "```Just now```", request.Embeds[0].Fields[2].Value)

			resp, err := httpmock.NewJsonResponse(200, map[string]interface{}{
				"status": "ok",
			})
			return resp, err
		},
	)

	MockDiscordHealthTracker.lastStatus = shared.UNHEALTHY
	err = MockDiscordHealthTracker.SendDiscordNotificationIfNeeded(shared.HEALTHY, generations, time.Now(), time.Now(), false, nil)
	assert.Nil(t, err)
}
