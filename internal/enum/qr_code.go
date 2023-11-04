package enum

type QrCodeStatus = string

const (
	QrCodeStatusUnauthorized QrCodeStatus = "UNAUTHORIZED" // 未授权
	QrCodeStatusAuthorizing  QrCodeStatus = "AUTHORIZING"  // 授权中
	QrCodeStatusAuthorized   QrCodeStatus = "AUTHORIZED"   // 已授权
)
