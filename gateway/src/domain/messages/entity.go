package messages

import (
	"amarolio-gateway/src/logger"
	"context"
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/redis/go-redis/v9"
)

type (
	Message struct {
		ID         int        `json:"id"`
		ChatroomID string     `json:"chatroom_id"`
		Role       string     `json:"role"`
		Content    string     `json:"content"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
	}
	GetMessages struct {
		Messages []Message `json:"messages"`
	}
	ChatClient struct {
		conn *websocket.Conn
		ps   *redis.PubSub
		wg   *sync.WaitGroup
		cf   context.CancelFunc
		lg   logger.Logger
	}
)

func NewChatClient(conn *websocket.Conn, wg *sync.WaitGroup, cf context.CancelFunc, lg logger.Logger) *ChatClient {
	return &ChatClient{conn, nil, wg, cf, lg}
}

func (cc *ChatClient) Subscribe(ps *redis.PubSub) {
	cc.ps = ps
}

func (cc *ChatClient) cleanUp() {
	cc.cf()
	cc.ps.Close()
	cc.wg.Done()
	cc.conn.Close()
}

func (cc *ChatClient) Read(ctx context.Context, userID string, chatroomID string, publisher func(ctx context.Context, userID string, chatroomID string, message string)) {
	defer cc.cleanUp()
	for {
		mt, msg, err := cc.conn.ReadMessage()
		if err != nil {
			cc.lg.Errorln(err)
			return
		}
		if mt == websocket.TextMessage {
			publisher(ctx, userID, chatroomID, string(msg))
		}
	}
}

func (cc *ChatClient) Write(ctx context.Context) {
	ticker := time.NewTicker(55 * time.Second)
	ch := cc.ps.Channel()

	defer func() {
		ticker.Stop()
		cc.cleanUp()
	}()

	for {
		select {
		case <-ctx.Done():
			cc.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		case msg, ok := <-ch:
			if !ok {
				cc.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := cc.conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				cc.lg.Errorln(err)
				return
			}
		case <-ticker.C:
			if err := cc.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				cc.lg.Errorln(err)
				return
			}
		}
	}
}
