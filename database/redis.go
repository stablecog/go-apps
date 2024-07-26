package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
)

var logInfo = log.Info
var logError = log.Error

type RedisWrapper struct {
	Client *redis.Client
	Ctx    context.Context
}

// Should return render redis url if render is set
func getRedisURL() string {
	return utils.GetEnv().RedisConnectionString
}

// Returns our *RedisWrapper, since we wrap some useful methods with the redis client
func NewRedis(ctx context.Context) (*RedisWrapper, error) {
	var opts *redis.Options
	var err error
	if utils.GetEnv().MockRedis {
		logInfo("Using mock redis client because MOCK_REDIS=true is set in environment")
		mr, _ := miniredis.Run()
		opts = &redis.Options{
			Addr: mr.Addr(),
		}
	} else {
		opts, err = redis.ParseURL(getRedisURL())
		if err != nil {
			logError("Error parsing REDIS_CONNECTION_STRING", "err", err)
			return nil, err
		}
	}
	redis := redis.NewClient(opts)
	_, err = redis.Ping(ctx).Result()
	if err != nil {
		logError("Error pinging Redis", "err", err)
		return nil, err
	}
	return &RedisWrapper{
		Client: redis,
		Ctx:    ctx,
	}, nil
}

// Set generate and upscale count stats
func (r *RedisWrapper) SetOutputCount(generationOutputCount, upscaleOutputCount, voiceoverOutputCount int) error {
	stats := RedisStats{
		GenerationOutputCount: generationOutputCount,
		VoiceoverOutputCount:  voiceoverOutputCount,
		UpscaleOutputCount:    upscaleOutputCount,
	}
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return r.Client.Set(context.Background(), "stats", statsJSON, 0).Err()
}

// Get generate and upscale count stats
func (r *RedisWrapper) GetStatsCount() (stats *RedisStats, err error) {
	statsJSON, err := r.Client.Get(context.Background(), "stats").Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(statsJSON), &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

type RedisStats struct {
	GenerationOutputCount int `json:"generation_output_count"`
	UpscaleOutputCount    int `json:"upscale_output_count"`
	VoiceoverOutputCount  int `json:"voiceover_output_count"`
}

// Enqueues a request to sc-worker
func (r *RedisWrapper) EnqueueCogRequest(ctx context.Context, stream string, request interface{}) error {
	_, err := r.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		ID:     "*", // Asterisk auto-generates an ID for the item on the stream
		Values: []interface{}{"value", request},
	}).Result()
	return err
}

func (r *RedisWrapper) GetQueueSize() (int64, error) {
	return r.Client.XLen(r.Ctx, shared.COG_REDIS_QUEUE).Result()
}

// Get pending request IDs on queue that are stale
// olderThan will be subtracted from the current time to return requests older than that
func (r *RedisWrapper) GetPendingGenerationAndUpscaleIDs(olderThan time.Duration) (generationOutputIDs, upscaleOutputIDs []PendingCogRequestRedis, err error) {
	// Get current time in MS since epoch, minus endAt
	to := time.Now().UnixNano()/int64(time.Millisecond) - int64(olderThan/time.Millisecond)
	// Get from the redis stream COG_REDIS_QUEUE, we want to read them without deleting them
	messages, err := r.Client.XRange(r.Ctx, shared.COG_REDIS_QUEUE, "0-0", fmt.Sprintf("%d", to)).Result()
	if err != nil {
		log.Error("Error getting pending generation and upscale IDs", "err", err)
		return nil, nil, err
	}

	generationOutputIDs = make([]PendingCogRequestRedis, 0)
	upscaleOutputIDs = make([]PendingCogRequestRedis, 0)

	for _, message := range messages {
		// Get the request ID from the message
		input, ok := message.Values["value"].(string)
		if input == "" || !ok {
			log.Error("Error getting value from message", "message", message)
			continue
		}
		// Deserialize
		var request requests.CogQueueRequest
		err = json.Unmarshal([]byte(input), &request)
		if err != nil {
			log.Error("Error deserializing input", "err", err)
			continue
		}

		if request.Input.ProcessType == shared.UPSCALE {
			upscaleOutputIDs = append(upscaleOutputIDs, PendingCogRequestRedis{
				RedisMsgid: message.ID,
				Type:       request.Input.ProcessType,
				ID:         request.Input.ID,
			})
		}
		if request.Input.ProcessType == shared.GENERATE || request.Input.ProcessType == shared.GENERATE_AND_UPSCALE {
			generationOutputIDs = append(generationOutputIDs, PendingCogRequestRedis{
				RedisMsgid: message.ID,
				Type:       request.Input.ProcessType,
				ID:         request.Input.ID,
			})
		}
	}

	return generationOutputIDs, upscaleOutputIDs, nil
}

type PendingCogRequestRedis struct {
	RedisMsgid string
	Type       shared.ProcessType
	ID         uuid.UUID
}

func (r *RedisWrapper) XDelListOfIDs(ids []string) (deleted int64, err error) {
	_, err = r.Client.XAck(r.Ctx, shared.COG_REDIS_QUEUE, shared.COG_REDIS_QUEUE, ids...).Result()
	if err != nil {
		log.Error("Error acking from redis", "err", err)
		return 0, err
	}
	deleted, err = r.Client.XDel(r.Ctx, shared.COG_REDIS_QUEUE, ids...).Result()
	if err != nil {
		log.Error("Error deleting from redis", "err", err)
	}
	return deleted, err
}

// Keep track of request ID to cog, with stream ID of the client, for timeout tracking
func (r *RedisWrapper) SetCogRequestStreamID(ctx context.Context, requestID string, streamID string) error {
	_, err := r.Client.Set(ctx, requestID, streamID, 1*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}

// Get the stream ID of the client for a given request ID
func (r *RedisWrapper) GetCogRequestStreamID(ctx context.Context, requestID string) (string, error) {
	return r.Client.Get(ctx, requestID).Result()
}

// Delete the stream ID of the client for a given request ID
func (r *RedisWrapper) DeleteCogRequestStreamID(ctx context.Context, requestID string) (int64, error) {
	return r.Client.Del(ctx, requestID).Result()
}

// Caching embeddings
func (r *RedisWrapper) CacheEmbeddings(ctx context.Context, key string, embedding []float32) error {
	// Convert embedding to string
	b, err := json.Marshal(embedding)
	if err != nil {
		log.Error("Error converting embedding to string", "err", err)
		return err
	}
	// Set embedding in redis
	err = r.Client.Set(ctx, key, b, 3*time.Minute).Err()
	if err != nil {
		log.Error("Error caching embedding", "err", err)
		return err
	}
	return nil
}

// Retrieve from cache
func (r *RedisWrapper) GetEmbeddings(ctx context.Context, key string) ([]float32, error) {
	// Get embedding from redis
	b, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.Error("Error getting embedding from cache", "err", err)
		}
		return nil, err
	}
	// Convert string to embedding
	var embedding []float32
	err = json.Unmarshal(b, &embedding)
	if err != nil {
		log.Error("Error converting embedding string to embedding", "err", err)
		return nil, err
	}
	return embedding, nil
}

// Caching translations
func (r *RedisWrapper) CacheTranslation(ctx context.Context, key string, translation string) error {
	// Set translation in redis
	err := r.Client.Set(ctx, fmt.Sprintf("translate:%s", key), translation, 24*time.Hour).Err()
	if err != nil {
		log.Error("Error caching embedding", "err", err)
		return err
	}
	return nil
}

// Retrieve from cache
func (r *RedisWrapper) GetTranslation(ctx context.Context, key string) (string, error) {
	// Get embedding from redis
	b, err := r.Client.Get(ctx, fmt.Sprintf("translate:%s", key)).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.Error("Error getting embedding from cache", "err", err)
		}
		return "", err
	}
	return string(b), nil
}

// Discord bot

// Set verify token and return token
func (r *RedisWrapper) SetDiscordVerifyToken(discordId string) (string, error) {
	// Generate random token
	token, err := utils.GenerateRandomHex(nil, 32)
	if err != nil {
		return "", err
	}

	// Set token in redis
	err = r.Client.Set(r.Ctx, fmt.Sprintf("disco_verify:%s", discordId), token, shared.DISCORD_VERIFY_TOKEN_EXPIRY).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// Get discord ID from token
func (r *RedisWrapper) GetDiscordTokenFromID(discordId string) (string, error) {
	return r.Client.Get(r.Ctx, fmt.Sprintf("disco_verify:%s", discordId)).Result()
}

// Delete token
func (r *RedisWrapper) DeleteDiscordToken(discordId string) error {
	return r.Client.Del(r.Ctx, fmt.Sprintf("disco_verify:%s", discordId)).Err()
}
