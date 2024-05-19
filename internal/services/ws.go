package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

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
	pool        *pgxpool.Pool
}

func NewWSService(upgrader websocket.Upgrader, pool *pgxpool.Pool) *WSService {
	return &WSService{
		upgrader:    upgrader,
		connections: make(map[*websocket.Conn]bool),
		broadcast:   make(chan Message),
		mutex:       &sync.Mutex{},
		pool:        pool,
	}
}

type Message struct {
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *WSService) Connect(w http.ResponseWriter, r *http.Request) {
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

	messages, err := s.loadMessages(r.Context())
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("load messages: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	for _, msg := range messages {
		if err := ws.WriteJSON(msg); err != nil {
			log.Printf("error sending message history: %v", err)
		}
	}

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
		msg.Timestamp = time.Now()
		if err := s.saveMessage(r.Context(), msg); err != nil {
			http.Error(
				w, fmt.Sprintf("save message: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		s.broadcast <- msg
	}
}

func (s *WSService) HandleMessages() {
	for {
		msg := <-s.broadcast
		s.mutex.Lock()
		for conn := range s.connections {
			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				conn.Close()
				delete(s.connections, conn)
			}
		}
		s.mutex.Unlock()
	}
}

func (s *WSService) saveMessage(ctx context.Context, msg Message) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquire database connection: %w", err)
	}
	defer conn.Release()

	if _, err = conn.Exec(context.Background(),
		"INSERT INTO messages (email, timestamp, message) VALUES ($1, $2, $3)",
		msg.Email, msg.Timestamp, msg.Message,
	); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func (s *WSService) loadMessages(ctx context.Context) ([]Message, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("acquire database connection: %w", err)
	}
	defer conn.Release()

	rows, err := conn.Query(
		ctx,
		"SELECT email, timestamp, message FROM messages ORDER BY timestamp",
	)
	if err != nil {
		return nil, fmt.Errorf("select messages: %w", err)
	}
	defer rows.Close()

	messages := make([]Message, 0)
	for rows.Next() {
		var msg Message
		if err := rows.Scan(
			&msg.Email, &msg.Timestamp, &msg.Message,
		); err != nil {
			continue
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
