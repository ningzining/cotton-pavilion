package qr_code_conn_cache

import (
	"github.com/gorilla/websocket"
)

var qrCodeConnMap = make(map[*websocket.Conn]string)

func Get(conn *websocket.Conn) string {
	return qrCodeConnMap[conn]
}

func Save(conn *websocket.Conn, ticket string) {
	qrCodeConnMap[conn] = ticket
}

func Remove(conn *websocket.Conn) {
	delete(qrCodeConnMap, conn)
}
