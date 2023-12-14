package qr_code_info_cache

import (
	"sync"
	"time"
	"user-center/internal/domain/model/do"
	"user-center/internal/infrastructure/util/timerutil"
)

func init() {
	timerutil.RunPeriodTask(ClearExpiredCache, time.Hour)
}

// 保存二维码对应的信息
var qrCodeInfoCache sync.Map

func Save(ticket string, value *do.QrCode) {
	qrCodeInfoCache.Store(ticket, value)
}

func Get(ticket string) (any, bool) {
	return qrCodeInfoCache.Load(ticket)
}

func Remove(ticket string) {
	qrCodeInfoCache.Delete(ticket)
}

func ClearExpiredCache() {
	var deleteKeys []string
	qrCodeInfoCache.Range(func(key, value any) bool {
		qrCode, ok := value.(*do.QrCode)
		if !ok {
			return true
		}
		if qrCode.IsExpired() {
			deleteKeys = append(deleteKeys, key.(string))
		}
		return true
	})

	for _, key := range deleteKeys {
		qrCodeInfoCache.Delete(key)
	}
}
