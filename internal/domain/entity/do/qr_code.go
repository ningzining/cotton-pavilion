package do

import (
	"time"
	"user-center/internal/domain/entity/enum"
)

type QrCode struct {
	Ticket         string // 二维码
	Status         string // 二维码状态
	TemporaryToken string // 临时token
	Token          string
	ExpiredAt      time.Time // 过期时间
}

func (q *QrCode) IsExpired() bool {
	return q.ExpiredAt.Before(time.Now())
}

func (q *QrCode) IsUnauthorized() bool {
	return q.Status == enum.QrCodeStatusUnauthorized
}

func (q *QrCode) IsAuthorizing() bool {
	return q.Status == enum.QrCodeStatusAuthorizing
}

func (q *QrCode) IsAuthorized() bool {
	return q.Status == enum.QrCodeStatusAuthorized && q.Token != ""
}

func (q *QrCode) UpdateAuthorizing(temporaryToken string) {
	q.Status = enum.QrCodeStatusAuthorizing
	q.TemporaryToken = temporaryToken
}

func (q *QrCode) UpdateAuthorized(token string) {
	q.Status = enum.QrCodeStatusAuthorized
	q.Token = token
}
