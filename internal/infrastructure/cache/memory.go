package cache

var memoryCache = make(map[string]string)

func Save(key, value string) {
	memoryCache[key] = value
}

func Get(key string) string {
	return memoryCache[key]
}

func Remove(key string) {
	delete(memoryCache, key)
}
