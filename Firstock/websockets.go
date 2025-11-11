package Firstock

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SafeConn struct {
	ws *websocket.Conn
	mu sync.Mutex
}

// Global connection manager
var connections struct {
	connMap  map[*SafeConn]bool            // active connections
	indexMap map[*websocket.Conn]*SafeConn // fast lookup by *websocket.Conn
	mu       sync.Mutex
}

// Initialize connection manager (call once)
func initConnections() {
	connections.connMap = make(map[*SafeConn]bool)
	connections.indexMap = make(map[*websocket.Conn]*SafeConn)
}

func (fs *firstock) InitializeWebSockets(userId string, model WebSocketModel) (conn *websocket.Conn, errWebSockets *ErrorResponseModel) {
	var res *http.Response
	var err error
	// Build the WebSocket URL
	baseUrl, headers, errWebSockets := getUrlAndHeaderData(userId)
	if errWebSockets != nil {
		return
	}

	conn, res, err = websocket.DefaultDialer.Dial(baseUrl, headers)
	if err != nil {
		errWebSockets = &ErrorResponseModel{
			Error: ErrorDetail{
				Message: err.Error(),
			},
		}
		log.Print(res)
		return
	}

	var msg []byte
	_, msg, err = conn.ReadMessage()
	if err != nil {
		return
	}
	if strings.Contains(string(msg), "Authentication successful") {
		time.Sleep(1 * time.Second)
		go readMessage(userId, conn, model)
	} else if strings.Contains(string(msg), "Maximum sessions limit (3) reached") {
		errWebSockets = &ErrorResponseModel{
			Error: ErrorDetail{
				Message: string(msg),
			},
		}
		return
	}
	if connections.connMap == nil {
		initConnections()
	}
	addConnection(conn)

	if len(model.Tokens) > 0 {
		subscribe(conn, model.Tokens)
	}

	return
}

func (fs *firstock) CloseWebSocket(conn *websocket.Conn) (errRes *ErrorResponseModel) {
	if conn != nil {
		if checkIfConnectionExists(conn) {
			err := conn.Close()
			if err != nil && !strings.Contains(err.Error(), "closed") {
				errRes = &ErrorResponseModel{
					Error: ErrorDetail{
						Message: err.Error(),
					},
				}
			} else {
				deleteConnection(conn)
			}
		} else {
			errRes = &ErrorResponseModel{
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
		}
	} else {
		errRes = &ErrorResponseModel{
			Error: ErrorDetail{
				Message: "Connection does not exist",
			},
		}
	}

	return
}

func (fs *firstock) Subscribe(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		errRes = subscribe(conn, data)
	} else {
		errRes = &ErrorResponseModel{
			Error: ErrorDetail{
				Message: "Connection does not exist",
			},
		}
	}
	return
}

func (fs *firstock) Unsubscribe(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		errRes = unsubscribe(conn, data)
	} else {
		errRes = &ErrorResponseModel{
			Error: ErrorDetail{
				Message: "Connection does not exist",
			},
		}
	}
	return
}
