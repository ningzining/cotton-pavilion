package cache

type ICache interface {
	Save(key string, value any)
	Get(key string, value any)
	Remove(key string)
}
