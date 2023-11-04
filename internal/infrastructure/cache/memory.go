package cache

var memoryCache = make(map[string]any)

type QrCodeValue struct {
	Status         string // 二维码状态
	TemporaryToken string // 临时token
	Token          string
}

func Save(key string, value any) {
	memoryCache[key] = value
}

func Get(key string) any {
	return memoryCache[key]
}

func Remove(key string) {
	delete(memoryCache, key)
}
