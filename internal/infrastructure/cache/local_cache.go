package cache

var qrCodeCache = make(map[string]any)

func Save(key string, value any) {
	qrCodeCache[key] = value
}

func Get(key string) any {
	return qrCodeCache[key]
}

func Remove(key string) {
	delete(qrCodeCache, key)
}
