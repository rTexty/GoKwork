# GoKwork

Простая асинхронная обёртка над закрытым API для фриланс биржи kwork.ru на языке Go

## Описание

GoKwork - это Go-клиент для работы с API kwork.ru. Библиотека предоставляет удобный интерфейс для взаимодействия с платформой, включая:

- 🔐 Аутентификацию
- 💬 Работу с сообщениями и диалогами
- 📊 Получение информации о проектах и заказах
- 🤖 Создание ботов для автоматических ответов
- 🔌 WebSocket подключение для real-time событий
- 🌐 Поддержку прокси (SOCKS5/SOCKS4)

## Установка

```bash
go get github.com/rtexty/gokwork
```

## Быстрый старт

### Простой API запрос

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rtexty/gokwork/pkg/kwork"
)

func main() {
    client, err := kwork.NewClient(kwork.Config{
        Login:    "login",
        Password: "password",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()

    // Получение своего профиля
    me, err := client.GetMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("My profile: %+v\n", me)
}
```

### Создание бота

```go
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
    bot, err := kwork.NewBot(kwork.Config{
        Login:    "login",
        Password: "password",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer bot.Close()

    // Обработчик для первого сообщения
    bot.MessageHandler("", true, "", func(ctx context.Context, msg *types.Message) error {
        text := "Здравствуйте, рад что вы обратились именно ко мне!"
        return msg.AnswerSimulation(ctx, text)
    })

    // Обработчик для сообщений содержащих слово "бот"
    bot.MessageHandler("", false, "бот", func(ctx context.Context, msg *types.Message) error {
        text := "Вам нужен бот? Могу помочь!"
        return msg.AnswerSimulation(ctx, text)
    })

    // Обработчик для точного совпадения "привет"
    bot.MessageHandler("привет", false, "", func(ctx context.Context, msg *types.Message) error {
        text := "И вам привет!"
        return msg.AnswerSimulation(ctx, text)
    })

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-sigChan
        cancel()
    }()

    log.Println("Bot is running!")
    if err := bot.Run(ctx); err != nil && err != context.Canceled {
        log.Fatal(err)
    }
}
```

## Основные возможности

### Аутентификация

```go
client, err := kwork.NewClient(kwork.Config{
    Login:     "login",
    Password:  "password",
    PhoneLast: "0102", // Последние 4 цифры телефона (если требуется)
})
```

### Работа с прокси

```go
client, err := kwork.NewClient(kwork.Config{
    Login:    "login",
    Password: "password",
    ProxyURL: "socks5://208.113.220.250:3420",
})
```

### API методы

```go
// Получение профиля
me, err := client.GetMe(ctx)
user, err := client.GetUser(ctx, userID)

// Диалоги
dialogs, err := client.GetAllDialogs(ctx)
messages, err := client.GetDialogWithUser(ctx, "username")

// Сообщения
err = client.SendMessage(ctx, userID, "текст")
err = client.DeleteMessage(ctx, messageID)
err = client.SetTyping(ctx, recipientID)

// Проекты
categories, err := client.GetCategories(ctx)
projects, err := client.GetProjects(ctx, kwork.ProjectsParams{
    CategoriesIDs: []int{11, 79},
    PriceFrom:     1000,
    PriceTo:       5000,
})

// Коннекты
connects, err := client.GetConnects(ctx)

// Заказы
workerOrders, err := client.GetWorkerOrders(ctx)
payerOrders, err := client.GetPayerOrders(ctx)

// Уведомления
notifications, err := client.GetNotifications(ctx)

// Статус
err = client.SetOffline(ctx)
```

### Обработчики бота

Бот поддерживает три типа обработчиков:

1. **OnStart** - срабатывает только на первое сообщение в диалоге
```go
bot.MessageHandler("", true, "", handlerFunc)
```

2. **Text** - точное совпадение текста
```go
bot.MessageHandler("привет", false, "", handlerFunc)
```

3. **TextContains** - содержит указанное слово
```go
bot.MessageHandler("", false, "бот", handlerFunc)
```

### Ответы на сообщения

```go
// С симуляцией набора текста (реалистично)
err := msg.AnswerSimulation(ctx, "Ответ")

// Быстрый ответ без симуляции
err := msg.FastAnswer(ctx, "Ответ")
```

## Структура проекта

```
GoKwork/
├── cmd/
│   └── examples/          # Примеры использования
│       ├── api_example.go
│       └── bot_example.go
├── pkg/
│   └── kwork/
│       ├── client.go      # Основной API клиент
│       ├── bot.go         # Бот с обработчиками
│       ├── websocket.go   # WebSocket слушатель
│       ├── types/         # Модели данных
│       └── errors/        # Кастомные ошибки
├── go.mod
├── go.sum
└── README.md
```

## Примечания

### Ошибка "Подтвердите что вы не робот"

Если получаете эту ошибку, используйте прокси:

```go
client, err := kwork.NewClient(kwork.Config{
    Login:    "login",
    Password: "password",
    ProxyURL: "socks5://your-proxy:port",
})
```

### Последние 4 цифры телефона

Если требуется подтверждение телефона:

```go
client, err := kwork.NewClient(kwork.Config{
    Login:     "login",
    Password:  "password",
    PhoneLast: "0102",
})
```

## Зависимости

- `github.com/gorilla/websocket` - WebSocket клиент
- `golang.org/x/net/proxy` - Поддержка SOCKS прокси

## Лицензия

MIT

## Автор

Портировано с Python версии [pykwork](https://github.com/kesha1225/pykwork)

## Вклад

Если у вас есть предложения или вы нашли баги, создавайте issue или pull request!
