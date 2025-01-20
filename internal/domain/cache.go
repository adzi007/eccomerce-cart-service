package domain

type CacheRepository interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration int64) error
	Delete(key string) error
}
