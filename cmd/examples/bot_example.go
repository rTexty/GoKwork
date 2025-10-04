package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rtexty/gokwork/pkg/kwork"
	"github.com/rtexty/gokwork/pkg/kwork/types"
)

func main() {
	// Создаем бота
	bot, err := kwork.NewBot(kwork.Config{
		Login:    "login",
		Password: "password",
		// ProxyURL: "socks5://64.90.53.198:46088", // Если нужен прокси
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	defer bot.Close()

	// Регистрируем обработчик для первого сообщения от юзера
	bot.MessageHandler("", true, "", func(ctx context.Context, msg *types.Message) error {
		text := "Здравствуйте, рад что вы обратились именно ко мне, опишите ваше желание подробнее!"
		return msg.AnswerSimulation(ctx, text)
	})

	// Регистрируем обработчик для сообщений содержащих слово "бот"
	bot.MessageHandler("", false, "бот", func(ctx context.Context, msg *types.Message) error {
		text := "Вам нужен бот? Можете посмотреть на примеры уже сделанных:..."
		return msg.AnswerSimulation(ctx, text)
	})

	// Регистрируем обработчик для точного совпадения текста "привет"
	bot.MessageHandler("привет", false, "", func(ctx context.Context, msg *types.Message) error {
		text := "И вам привет!"
		return msg.AnswerSimulation(ctx, text)
	})

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received interrupt signal, shutting down...")
		cancel()
	}()

	// Запускаем бота
	log.Println("Starting bot...")
	if err := bot.Run(ctx); err != nil && err != context.Canceled {
		log.Fatalf("Bot error: %v", err)
	}

	log.Println("Bot stopped gracefully")
}
