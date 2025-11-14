// Copyright (c) [2025] [Firstock]
// SPDX-License-Identifier: MIT
package Firstock

import (
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

func (fs *firstock) InitializeWebSockets(userId string, model WebSocketModel) (errWebSockets *ErrorResponseModel) {
	var err error

	if model.WebSocketConection == nil {
		errWebSockets = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection callback is nil",
			},
		}
		return
	}
	// Build the WebSocket URL
	baseUrl, headers, errWebSockets := getUrlAndHeaderData(userId)
	if errWebSockets != nil {
		return
	}

	var conn *websocket.Conn
	conn, _, err = websocket.DefaultDialer.Dial(baseUrl, headers)
	if err != nil {
		errWebSockets = &ErrorResponseModel{
			Code:   "500",
			Status: "failed",
			Error: ErrorDetail{
				Message: err.Error(),
			},
		}
		return
	}

	var msg []byte
	_, msg, err = conn.ReadMessage()
	if err != nil {
		errWebSockets = &ErrorResponseModel{
			Code:   "500",
			Status: "failed",
			Error: ErrorDetail{
				Message: err.Error(),
			},
		}
		return
	}

	if strings.Contains(string(msg), "Maximum sessions limit (3) reached") {
		errWebSockets = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: string(msg),
			},
		}
		return
	} else if strings.Contains(string(msg), "failed") {
		errWebSockets = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: string(msg),
			},
		}
		return
	}
	if connections.connMap == nil {
		initConnections()
	}
	if strings.Contains(string(msg), "Authentication successful") {
		conn.SetWriteDeadline(time.Now().Add(50 * time.Second))
		addConnection(conn)

		if len(model.SubscribeFeedTokens) > 0 {
			go subscribe(conn, model.SubscribeFeedTokens)
		}
		if len(model.SubscribeOptionGreeksTokens) > 0 {
			subscribeOptionGreeks(conn, model.SubscribeOptionGreeksTokens)
		}

		go readMessage(userId, conn, model)
	}
	model.WebSocketConection(conn)
	return
}

func (fs *firstock) CloseWebSocket(conn *websocket.Conn) (errRes *ErrorResponseModel) {
	if conn != nil {
		if checkIfConnectionExists(conn) {
			deleteConnection(conn)
		} else {
			errRes = &ErrorResponseModel{
				Code:   "400",
				Status: "failed",
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
		}
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
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
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
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
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
	}
	return
}

func (fs *firstock) SubscribeOptionGreeks(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		errRes = subscribeOptionGreeks(conn, data)
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
	}
	return
}

func (fs *firstock) UnsubscribeOptionGreeks(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		errRes = unsubscribeOptionGreeks(conn, data)
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
	}
	return
}
