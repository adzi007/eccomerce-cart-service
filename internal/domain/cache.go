package domain

type CacheRepository interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration int64) error
	Delete(key string) error
	MSetProductsCache(data map[string]ProductServiceResponse, expiration int64) error
	MGetProductsCache(keys []string, keyPrefix string) ([]ProductServiceResponse, []uint, error)
}
