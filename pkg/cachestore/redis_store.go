package cachestore

import (
	"cart-service/internal/domain"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a new Redis cache instance
func NewRedisCache(ctx context.Context, addr string, password string, db int) domain.CacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisCache{
		client: client,
		ctx:    ctx,
	}
}

func (r *redisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key not found
	}
	return val, err
}

func (r *redisCache) Set(key string, value string, expiration int64) error {
	return r.client.Set(r.ctx, key, value, time.Duration(expiration)*time.Second).Err()
}

func (r *redisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// MSetStructs stores multiple ProductCart structs in Redis
func (r *redisCache) MSetProductsCache(data map[string]domain.ProductServiceResponse, expiration int64) error {
	// Prepare a flat map for MSET
	flatData := make(map[string]string)
	for key, product := range data {
		// Serialize each struct to JSON
		jsonData, err := json.Marshal(product)
		if err != nil {
			return err
		}
		flatData[key] = string(jsonData)
	}

	// Use MSET to store all data
	_, err := r.client.MSet(r.ctx, flatData).Result()
	if err != nil {
		return err
	}

	// Set expiration for each key (MSET doesn't support expiration directly)
	for key := range flatData {
		err := r.client.Expire(r.ctx, key, time.Duration(expiration)*time.Second).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// MGetStructs retrieves multiple ProductCart structs from Redis
func (r *redisCache) MGetProductsCache(keys []string, keyPrefix string) ([]domain.ProductServiceResponse, []uint, error) {

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = keyPrefix + key
	}

	// Use MGET to fetch all values
	values, err := r.client.MGet(r.ctx, prefixedKeys...).Result()

	if err != nil {
		pp.Println("Error fetching values from Redis:", err)
		return nil, nil, err
	}

	// Prepare the result map
	var result []domain.ProductServiceResponse
	var missingKeyProducts []uint

	// Iterate through the keys and deserialize values
	for i, key := range keys {
		if values[i] == nil {
			key, _ := strconv.ParseUint(key, 10, 64)
			missingKeyProducts = append(missingKeyProducts, uint(key))
			// pp.Println("Key not found in Redis:", key)
			continue // Skip missing keys
		}

		// Check type assertion
		strValue, ok := values[i].(string)
		if !ok {
			pp.Println("Error: Value is not a string:", values[i])
			continue
		}

		// Validate JSON
		if !json.Valid([]byte(strValue)) {
			// pp.Println("Error: Invalid JSON for key:", key, "Value:", strValue)
			continue
		}

		// Deserialize JSON to ProductCart struct
		var product domain.ProductServiceResponse
		err := json.Unmarshal([]byte(strValue), &product)
		if err != nil {
			pp.Println("Error unmarshaling JSON for key:", key, "Error:", err)
			return nil, nil, err
		}

		// result[key] = product
		result = append(result, product)
	}

	return result, missingKeyProducts, nil
}
