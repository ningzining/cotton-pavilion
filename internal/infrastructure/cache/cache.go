package cache

import "time"

type ICache interface {
	Get(key string) string
	Set(key string, value string, expiration time.Duration) error
	Remove(key ...string) error
}
