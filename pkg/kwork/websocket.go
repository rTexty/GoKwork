package kwork

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rtexty/gokwork/pkg/kwork/types"
)

// MessageListener слушает сообщения через WebSocket
func (c *Client) MessageListener(ctx context.Context, messageChan chan<- *types.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := c.listenMessagesOnce(ctx, messageChan); err != nil {
				log.Printf("WebSocket error: %v, reconnecting in 10 seconds...", err)
				time.Sleep(10 * time.Second)
				continue
			}
		}
	}
}

func (c *Client) listenMessagesOnce(ctx context.Context, messageChan chan<- *types.Message) error {
	channel, err := c.getChannel(ctx)
	if err != nil {
		return fmt.Errorf("failed to get channel: %w", err)
	}

	uri := fmt.Sprintf("wss://notice.kwork.ru/ws/public/%s", channel)

	conn, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %w", err)
	}
	defer conn.Close()

	// Канал для закрытия при отмене контекста
	done := make(chan struct{})
	defer close(done)

	// Горутина для обработки отмены контекста
	go func() {
		select {
		case <-ctx.Done():
			conn.Close()
		case <-done:
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("failed to read message: %w", err)
			}

			log.Printf("Received WebSocket data: %s", string(message))

			// Парсим внешний JSON
			var wsEvent struct {
				Text string `json:"text"`
			}
			if err := json.Unmarshal(message, &wsEvent); err != nil {
				log.Printf("Failed to unmarshal outer event: %v", err)
				continue
			}

			// Парсим внутренний JSON (данные события)
			var event types.BaseEvent
			if err := json.Unmarshal([]byte(wsEvent.Text), &event); err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			// Игнорируем события "печатает"
			if event.Event == types.EventTypeIsTyping {
				continue
			}

			// Обрабатываем новые сообщения
			msg := c.processEvent(&event)
			if msg != nil {
				messageChan <- msg
			}
		}
	}
}

func (c *Client) processEvent(event *types.BaseEvent) *types.Message {
	switch event.Event {
	case types.EventTypeNewMessage:
		return c.processNewMessageEvent(event)
	case types.EventTypeNotify:
		return c.processNotifyEvent(event)
	case types.EventTypePopUpNotify:
		return c.processPopUpNotifyEvent(event)
	default:
		return nil
	}
}

func (c *Client) processNewMessageEvent(event *types.BaseEvent) *types.Message {
	fromID, _ := event.Data["from"].(float64)
	text, _ := event.Data["inboxMessage"].(string)
	toUserID, _ := event.Data["to_user_id"].(float64)
	inboxID, _ := event.Data["inbox_id"].(float64)
	title, _ := event.Data["title"].(string)
	lastMessage, _ := event.Data["lastMessage"].(map[string]interface{})

	return types.NewMessage(
		c,
		int(fromID),
		text,
		int(toUserID),
		int(inboxID),
		title,
		lastMessage,
	)
}

func (c *Client) processNotifyEvent(event *types.BaseEvent) *types.Message {
	// Проверяем наличие нового сообщения
	if newMsg, ok := event.Data[types.NotifyNewMessage]; ok && newMsg != nil {
		// Проверяем наличие данных диалога
		if dialogData, ok := event.Data["dialog_data"]; !ok || dialogData == nil {
			// Получаем последний диалог
			ctx := context.Background()
			dialogs, err := c.GetAllDialogs(ctx)
			if err != nil || len(dialogs) == 0 {
				log.Printf("Failed to get dialogs: %v", err)
				return nil
			}

			lastDialog := dialogs[0]
			return types.NewMessage(
				c,
				lastDialog.UserID,
				lastDialog.LastMessageText,
				0,
				0,
				"",
				nil,
			)
		} else {
			// Есть данные диалога, получаем сообщение
			dialogDataSlice, ok := event.Data["dialog_data"].([]interface{})
			if !ok || len(dialogDataSlice) == 0 {
				return nil
			}

			dialogInfo, ok := dialogDataSlice[0].(map[string]interface{})
			if !ok {
				return nil
			}

			login, _ := dialogInfo["login"].(string)
			if login == "" {
				return nil
			}

			ctx := context.Background()
			messages, err := c.GetDialogWithUser(ctx, login)
			if err != nil || len(messages) == 0 {
				log.Printf("Failed to get messages: %v", err)
				return nil
			}

			msg := messages[0]
			return types.NewMessage(
				c,
				msg.FromID,
				msg.Message,
				msg.ToID,
				msg.MessageID,
				"",
				nil,
			)
		}
	}

	return nil
}

func (c *Client) processPopUpNotifyEvent(event *types.BaseEvent) *types.Message {
	popUpNotify, ok := event.Data["pop_up_notify"].(map[string]interface{})
	if !ok {
		return nil
	}

	data, ok := popUpNotify["data"].(map[string]interface{})
	if !ok {
		return nil
	}

	username, _ := data["username"].(string)
	if username == "" {
		return nil
	}

	ctx := context.Background()
	messages, err := c.GetDialogWithUser(ctx, username)
	if err != nil || len(messages) == 0 {
		log.Printf("Failed to get messages: %v", err)
		return nil
	}

	msg := messages[0]
	return types.NewMessage(
		c,
		msg.FromID,
		msg.Message,
		msg.ToID,
		msg.MessageID,
		"",
		nil,
	)
}
