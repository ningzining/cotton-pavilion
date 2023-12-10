package qr_code_info_cache

import (
	"sync"
	"user-center/internal/domain/entity/do"
)

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
