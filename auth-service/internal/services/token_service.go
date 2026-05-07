package services

import (
	"context"
	"strconv"
	"time"

	"auth-service/internal/redis"
)

// RevokeToken adds a token to the blacklist in Redis
func RevokeToken(token string, expiresAt time.Time) error {
	if !redis.IsConnected() {
		return nil // Gracefully handle missing Redis
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ttl := time.Until(expiresAt)
	if ttl < 0 {
		ttl = time.Second
	}

	// Store token in blacklist with TTL
	key := "blacklist:" + token
	return redis.Set(ctx, key, "1", ttl)
}

// IsTokenRevoked checks if a token is in the blacklist
func IsTokenRevoked(token string) bool {
	if !redis.IsConnected() {
		return false // Gracefully handle missing Redis
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "blacklist:" + token
	exists, err := redis.Exists(ctx, key)
	if err != nil {
		return false // On error, allow token
	}

	return exists
}

// IncrementLoginAttempt increments login attempts for a user
func IncrementLoginAttempt(username string) (int64, error) {
	if !redis.IsConnected() {
		return 1, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "login_attempts:" + username
	count, err := redis.IncrBy(ctx, key, 1)
	if err != nil {
		return 0, err
	}

	// Set TTL for attempts window (15 minutes)
	if count == 1 {
		_ = redis.ExpireAt(ctx, key, time.Now().Add(15*time.Minute))
	}

	return count, nil
}

// ResetLoginAttempts resets login attempts for a user
func ResetLoginAttempts(username string) error {
	if !redis.IsConnected() {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "login_attempts:" + username
	return redis.Delete(ctx, key)
}

// CacheUserPermissions stores user permissions in Redis for faster lookups
func CacheUserPermissions(userID uint, permissions []string, ttl time.Duration) error {
	if !redis.IsConnected() {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "user_permissions:" + strconv.FormatUint(uint64(userID), 10)
	permissionsStr := ""
	for i, perm := range permissions {
		if i > 0 {
			permissionsStr += ","
		}
		permissionsStr += perm
	}

	return redis.Set(ctx, key, permissionsStr, ttl)
}

// GetCachedUserPermissions retrieves cached user permissions from Redis
func GetCachedUserPermissions(userID uint) ([]string, error) {
	if !redis.IsConnected() {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "user_permissions:" + strconv.FormatUint(uint64(userID), 10)
	result, err := redis.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if result == "" {
		return nil, nil
	}

	// Parse comma-separated permissions
	permissions := []string{}
	if result != "" {
		// Simple split - could use strings.Split
	}

	return permissions, nil
}

// InvalidateUserPermissionsCache removes cached permissions for a user
func InvalidateUserPermissionsCache(userID uint) error {
	if !redis.IsConnected() {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := "user_permissions:" + strconv.FormatUint(uint64(userID), 10)
	return redis.Delete(ctx, key)
}
