// Copyright (c) [2025] [Firstock]
// SPDX-License-Identifier: MIT
package Firstock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func getUrlAndHeaderData(userId string) (baseUrl string, headers http.Header, err *ErrorResponseModel) {
	urlVal := url.URL{Scheme: scheme, Host: host, Path: path}
	jKey, err := readJkey(userId)
	if err != nil {
		return
	} else if jKey == "" {
		err = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "Please login first",
			},
		}
		return
	}

	q := urlVal.Query()
	q.Set("userId", userId)
	q.Set("jKey", jKey)
	q.Set("source", srcVal)

	urlVal.RawQuery = q.Encode()
	headers = http.Header{}
	headers.Add("accept-encoding", accept_encoding)
	headers.Add("accept-language", accept_language)
	headers.Add("cache-control", no_cache)
	headers.Add("origin", origin)
	headers.Add("pragma", no_cache)
	baseUrl = urlVal.String()
	return
}

func readMessage(userId string, conn *websocket.Conn, model WebSocketModel) {
	for {
		if !checkIfConnectionExists(conn) {
			return
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			for i := 0; i < maxWebsocketConnectionRetries; i++ {
				time.Sleep(timeInterval * time.Second)
				if !checkIfConnectionExists(conn) {
					return
				}
				fmt.Println("Attempting to reconnect...")

				// Build the WebSocket URL
				baseUrl, headers, errWebSockets := getUrlAndHeaderData(userId)
				if errWebSockets != nil {
					return
				}
				conn2, _, err := websocket.DefaultDialer.Dial(baseUrl, headers)
				if err != nil {
					continue // try again
				}
				if checkIfConnectionExists(conn) {
					deleteConnection(conn)
				} else {
					conn = nil
				}
				conn = conn2
				addConnection(conn)

				if model.SubscribeFeedData != nil && (len(model.SubscribeFeedTokens) > 0) {
					subscribe(conn, model.SubscribeFeedTokens)
				}
				if model.SubscribeOptionGreeksData != nil && (len(model.SubscribeOptionGreeksTokens) > 0) {
					subscribeOptionGreeks(conn, model.SubscribeOptionGreeksTokens)
				}
				fmt.Println("Reconnected successfully.")
				break
			}

			_, msg, err = conn.ReadMessage()
			if err != nil {
				continue
			}
		}
		if strings.Contains(string(msg), "Authentication successful") || strings.Contains(string(msg), `\"status\":\"failed\"`) {
			continue
		}
		if model.OrderData != nil && strings.Contains(string(msg), "norenordno") {
			var data map[string]string
			err = json.Unmarshal(msg, &data)
			if err == nil {
				model.OrderData(data)
			}
		} else if model.PositonData != nil && strings.Contains(string(msg), "brkname") {
			var data map[string]interface{}
			err = json.Unmarshal(msg, &data)
			if err == nil {
				model.PositonData(data)
			}
		} else if model.SubscribeFeedData != nil && !(strings.Contains(string(msg), "brkname") || strings.Contains(string(msg), "norenordno") ||
			strings.Contains(string(msg), "gamma")) {
			var data SubscribeFeedModel
			err = json.Unmarshal(msg, &data)
			if err == nil && data.CExchSeg != "" {
				model.SubscribeFeedData(data)
			} else {
				var res map[string]SubscribeFeedModel
				err = json.Unmarshal(msg, &res)
				if err == nil {
					for _, val := range res {
						model.SubscribeFeedData(val)
					}
				}
			}
		} else if model.SubscribeOptionGreeksData != nil && strings.Contains(string(msg), "gamma") {
			var res map[string]OptionGreeksModel
			err = json.Unmarshal(msg, &res)
			if err == nil {
				for _, val := range res {
					model.SubscribeOptionGreeksData(val)
				}
			}
		}
	}
}

// Add a connection (safe, avoids duplicates)
func addConnection(ws *websocket.Conn) *SafeConn {
	connections.mu.Lock()
	defer connections.mu.Unlock()

	// Check if already exists
	if existing, ok := connections.indexMap[ws]; ok {
		fmt.Println("Connection already exists")
		return existing
	}

	safe := &SafeConn{ws: ws}
	connections.connMap[safe] = true
	connections.indexMap[ws] = safe

	fmt.Println("Connection added")
	return safe
}

// Delete a connection
func deleteConnection(ws *websocket.Conn) {
	connections.mu.Lock()
	defer connections.mu.Unlock()

	if safe, ok := connections.indexMap[ws]; ok {
		delete(connections.connMap, safe)
		delete(connections.indexMap, ws)
		safe.ws.Close() // optional
		fmt.Println("Connection deleted")
		return
	}
}

// Check if connection exists
func checkIfConnectionExists(ws *websocket.Conn) bool {
	connections.mu.Lock()
	defer connections.mu.Unlock()

	_, exists := connections.indexMap[ws]
	return exists
}

// Safe write to a specific connection
func writeMessage(ws *websocket.Conn, data []byte) error {
	connections.mu.Lock()
	safe, ok := connections.indexMap[ws]
	connections.mu.Unlock()

	if !ok {
		return fmt.Errorf("connection not found")
	}

	safe.mu.Lock()
	defer safe.mu.Unlock()

	return safe.ws.WriteMessage(websocket.TextMessage, data)
}

func subscribe(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		if !checkIfConnectionExists(conn) {
			errRes = &ErrorResponseModel{
				Code:   "400",
				Status: "failed",
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
			return
		}
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
		return
	}
	tokens := strings.Join(data, "|")
	msg := fmt.Sprintf(`{"action":"subscribe","tokens":"%s"}`, tokens)

	_ = writeMessage(conn, []byte(msg))
	return
}

func unsubscribe(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		if !checkIfConnectionExists(conn) {
			errRes = &ErrorResponseModel{
				Code:   "400",
				Status: "failed",
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
			return
		}
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
		return
	}
	tokens := strings.Join(data, "|")
	msg := fmt.Sprintf(`{"action":"unsubscribe","tokens":"%s"}`, tokens)

	_ = writeMessage(conn, []byte(msg))
	return
}

func subscribeOptionGreeks(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		if !checkIfConnectionExists(conn) {
			errRes = &ErrorResponseModel{
				Code:   "400",
				Status: "failed",
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
			return
		}
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
		return
	}
	tokens := strings.Join(data, "|")
	msg := fmt.Sprintf(`{"action":"subscribe-option-greeks","tokens":"%s"}`, tokens)

	_ = writeMessage(conn, []byte(msg))
	return
}

func unsubscribeOptionGreeks(conn *websocket.Conn, data []string) (errRes *ErrorResponseModel) {
	if conn != nil {
		if !checkIfConnectionExists(conn) {
			errRes = &ErrorResponseModel{
				Code:   "400",
				Status: "failed",
				Error: ErrorDetail{
					Message: "Connection does not exist",
				},
			}
			return
		}
	} else {
		errRes = &ErrorResponseModel{
			Code:   "400",
			Status: "failed",
			Error: ErrorDetail{
				Message: "WebSocket connection is not initialized",
			},
		}
		return
	}
	tokens := strings.Join(data, "|")
	msg := fmt.Sprintf(`{"action":"unsubscribe-option-greeks","tokens":"%s"}`, tokens)

	_ = writeMessage(conn, []byte(msg))
	return
}
