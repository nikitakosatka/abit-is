package services

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)

type WSService struct {
	upgrader    websocket.Upgrader
	connections map[*websocket.Conn]bool
	broadcast   chan Message
	mutex       *sync.Mutex
}

func NewWSService(upgrader websocket.Upgrader) *WSService {
	return &WSService{
		upgrader:    upgrader,
		connections: make(map[*websocket.Conn]bool),
		broadcast:   make(chan Message),
		mutex:       &sync.Mutex{},
	}
}

type Message struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (s *WSService) Connect(
	w http.ResponseWriter, r *http.Request,
) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("upgrade: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	defer ws.Close()

	s.mutex.Lock()
	s.connections[ws] = true
	s.mutex.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			s.mutex.Lock()
			delete(s.connections, ws)
			s.mutex.Unlock()
			break
		}
		s.broadcast <- msg
	}
}
func (s *WSService) HandleMessages() {
	for {
		msg := <-s.broadcast
		s.mutex.Lock()
		for conn := range s.connections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				conn.Close()
				delete(s.connections, conn)
			}
		}
		s.mutex.Unlock()
	}
}
