package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

// ConnectRedis initializes Redis connection
func ConnectRedis(redisURL string) error {
	if redisURL == "" {
		log.Println("⚠️  Redis URL not configured, running in degraded mode")
		return nil
	}

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("❌ Failed to parse Redis URL: %v\n", err)
		return err
	}

	Client = redis.NewClient(opts)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = Client.Ping(ctx).Err()
	if err != nil {
		log.Printf("❌ Failed to connect to Redis: %v\n", err)
		Client = nil
		return err
	}

	log.Println("✅ Redis connected successfully")
	return nil
}

// IsConnected checks if Redis is available
func IsConnected() bool {
	return Client != nil
}

// Close gracefully closes Redis connection
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

// Set stores a key-value pair with TTL
func Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if !IsConnected() {
		return nil // Gracefully handle missing Redis
	}

	return Client.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value by key
func Get(ctx context.Context, key string) (string, error) {
	if !IsConnected() {
		return "", nil
	}

	return Client.Get(ctx, key).Result()
}

// Delete removes a key
func Delete(ctx context.Context, key string) error {
	if !IsConnected() {
		return nil
	}

	return Client.Del(ctx, key).Err()
}

// Exists checks if a key exists
func Exists(ctx context.Context, key string) (bool, error) {
	if !IsConnected() {
		return false, nil
	}

	result, err := Client.Exists(ctx, key).Result()
	return result > 0, err
}

// IncrBy increments a counter by value
func IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	if !IsConnected() {
		return 0, nil
	}

	return Client.IncrBy(ctx, key, value).Result()
}

// ExpireAt sets expiration time
func ExpireAt(ctx context.Context, key string, tm time.Time) error {
	if !IsConnected() {
		return nil
	}

	return Client.ExpireAt(ctx, key, tm).Err()
}
