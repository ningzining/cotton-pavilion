package qr_code_conn_cache

import (
	"github.com/gorilla/websocket"
	"sync"
)

// 保存每个ws连接对应的二维码
var qrCodeConnMap sync.Map

func Get(conn *websocket.Conn) (any, bool) {
	return qrCodeConnMap.Load(conn)
}

func Save(conn *websocket.Conn, value string) {
	qrCodeConnMap.Store(conn, value)
}

func Remove(conn *websocket.Conn) {
	qrCodeConnMap.Delete(conn)
}
