package types

import (
	"context"
	"time"
)

// Message представляет сообщение от бота
type Message struct {
	FromID      int
	Text        string
	ToUserID    int
	InboxID     int
	Title       string
	LastMessage map[string]interface{}
	api         MessageSender
}

// MessageSender интерфейс для отправки сообщений
type MessageSender interface {
	SendMessage(ctx context.Context, userID int, text string) error
	SetTyping(ctx context.Context, recipientID int) error
}

// NewMessage создает новое сообщение
func NewMessage(api MessageSender, fromID int, text string, toUserID, inboxID int, title string, lastMessage map[string]interface{}) *Message {
	return &Message{
		FromID:      fromID,
		Text:        text,
		ToUserID:    toUserID,
		InboxID:     inboxID,
		Title:       title,
		LastMessage: lastMessage,
		api:         api,
	}
}

// AnswerSimulation отправляет реалистичный ответ с симуляцией набора текста
func (m *Message) AnswerSimulation(ctx context.Context, text string) error {
	time.Sleep(2 * time.Second)
	if err := m.api.SetTyping(ctx, m.FromID); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	return m.api.SendMessage(ctx, m.FromID, text)
}

// FastAnswer отправляет быстрый ответ без симуляции
func (m *Message) FastAnswer(ctx context.Context, text string) error {
	return m.api.SendMessage(ctx, m.FromID, text)
}
