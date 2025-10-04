package kwork

import (
	"context"
	"log"
	"strings"

	"github.com/rtexty/gokwork/pkg/kwork/errors"
	"github.com/rtexty/gokwork/pkg/kwork/types"
)

// HandlerFunc функция-обработчик сообщения
type HandlerFunc func(ctx context.Context, msg *types.Message) error

// Handler представляет обработчик сообщения
type Handler struct {
	Func         HandlerFunc
	Text         string
	OnStart      bool
	TextContains string
}

// Bot представляет бота Kwork
type Bot struct {
	*Client
	handlers []Handler
}

// NewBot создает нового бота
func NewBot(cfg Config) (*Bot, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Client:   client,
		handlers: make([]Handler, 0),
	}, nil
}

// MessageHandler регистрирует обработчик сообщений
func (b *Bot) MessageHandler(text string, onStart bool, textContains string, handler HandlerFunc) {
	b.handlers = append(b.handlers, Handler{
		Func:         handler,
		Text:         text,
		OnStart:      onStart,
		TextContains: textContains,
	})
}

// Run запускает бота
func (b *Bot) Run(ctx context.Context) error {
	if len(b.handlers) == 0 {
		return errors.NewKworkBotError("no handlers registered")
	}

	log.Println("Bot is running!")

	messageChan := make(chan *types.Message, 100)

	// Запускаем слушатель сообщений в отдельной горутине
	go func() {
		if err := b.MessageListener(ctx, messageChan); err != nil {
			log.Printf("Message listener error: %v", err)
		}
	}()

	// Обрабатываем входящие сообщения
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-messageChan:
			// Обрабатываем сообщение всеми подходящими хендлерами
			for _, handler := range b.handlers {
				if b.shouldHandleMessage(ctx, msg, &handler) {
					log.Printf("Found handler for message: %s", msg.Text)
					if err := handler.Func(ctx, msg); err != nil {
						log.Printf("Handler error: %v", err)
					}
				}
			}
		}
	}
}

// shouldHandleMessage проверяет, должен ли хендлер обработать сообщение
func (b *Bot) shouldHandleMessage(ctx context.Context, msg *types.Message, handler *Handler) bool {
	// Если не установлены никакие условия, обрабатываем все сообщения
	if !handler.OnStart && handler.Text == "" && handler.TextContains == "" {
		return true
	}

	// Проверяем условие "первое сообщение"
	if handler.OnStart {
		dialogs, err := b.GetAllDialogs(ctx)
		if err != nil || len(dialogs) == 0 {
			return false
		}

		fromUsername := dialogs[0].Username
		dialog, err := b.GetDialogWithUser(ctx, fromUsername)
		if err != nil {
			return false
		}

		// Если в диалоге только одно сообщение, это первое сообщение
		return len(dialog) == 1
	}

	// Проверяем точное совпадение текста
	if handler.Text != "" && strings.EqualFold(handler.Text, msg.Text) {
		return true
	}

	// Проверяем вхождение текста
	if handler.TextContains != "" && b.dispatchTextContains(handler.TextContains, msg.Text) {
		return true
	}

	return false
}

// dispatchTextContains проверяет содержит ли сообщение указанный текст
func (b *Bot) dispatchTextContains(text, messageText string) bool {
	lowerText := strings.ToLower(text)
	words := strings.Fields(strings.ToLower(messageText))

	for _, word := range words {
		// Убираем знаки препинания
		cleanWord := strings.Trim(word, "...!.?-")
		if cleanWord == lowerText {
			return true
		}
	}

	return false
}
