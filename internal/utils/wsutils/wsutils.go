package wsutils

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func UpGrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
