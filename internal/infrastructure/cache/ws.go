package cache

import "github.com/gorilla/websocket"

var qrCodeConnMap = make(map[*websocket.Conn]string)

func GetQrCodeConnTicket(conn *websocket.Conn) string {
	return qrCodeConnMap[conn]
}

func SaveQrCodeConnTicket(conn *websocket.Conn, ticket string) {
	qrCodeConnMap[conn] = ticket
}

func RemoveQrCodeConnTicket(conn *websocket.Conn) {
	delete(qrCodeConnMap, conn)
}
