package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rtexty/gokwork/pkg/kwork"
)

func main() {
	// Создаем клиента
	client, err := kwork.NewClient(kwork.Config{
		Login:    "login",
		Password: "password",
		// PhoneLast: "0102", // Если требуется подтверждение последних 4 цифр телефона
		// ProxyURL: "socks5://208.113.220.250:3420", // Если нужен прокси
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Получение своего профиля
	me, err := client.GetMe(ctx)
	if err != nil {
		log.Fatalf("Failed to get me: %v", err)
	}
	fmt.Printf("My profile: %+v\n\n", me)

	// Получение профиля юзера
	user, err := client.GetUser(ctx, 1456898)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	fmt.Printf("User profile: %+v\n\n", user)

	// Получение ваших коннектов
	connects, err := client.GetConnects(ctx)
	if err != nil {
		log.Fatalf("Failed to get connects: %v", err)
	}
	fmt.Printf("Connects: %+v\n\n", connects)

	// Получение всех диалогов
	allDialogs, err := client.GetAllDialogs(ctx)
	if err != nil {
		log.Fatalf("Failed to get dialogs: %v", err)
	}
	fmt.Printf("All dialogs: %+v\n\n", allDialogs)

	// Получение диалога с пользователем
	if len(allDialogs) > 0 {
		dialogWithUser, err := client.GetDialogWithUser(ctx, allDialogs[0].Username)
		if err != nil {
			log.Fatalf("Failed to get dialog with user: %v", err)
		}
		fmt.Printf("Dialog with user: %+v\n\n", dialogWithUser)
	}

	// Получение категорий заказов на бирже
	categories, err := client.GetCategories(ctx)
	if err != nil {
		log.Fatalf("Failed to get categories: %v", err)
	}
	fmt.Printf("Categories: %+v\n\n", categories)

	// Получение проектов с биржи по ID категорий
	projects, err := client.GetProjects(ctx, kwork.ProjectsParams{
		CategoriesIDs: []int{11, 79},
	})
	if err != nil {
		log.Fatalf("Failed to get projects: %v", err)
	}
	fmt.Printf("Projects: %+v\n\n", projects)

	// Получение выполненных и отменённых заказов (работник)
	workerOrders, err := client.GetWorkerOrders(ctx)
	if err != nil {
		log.Fatalf("Failed to get worker orders: %v", err)
	}
	fmt.Printf("Worker orders: %+v\n\n", workerOrders)

	// Получение выполненных и отменённых заказов (заказчик)
	payerOrders, err := client.GetPayerOrders(ctx)
	if err != nil {
		log.Fatalf("Failed to get payer orders: %v", err)
	}
	fmt.Printf("Payer orders: %+v\n\n", payerOrders)

	// Отправка сообщения
	if err := client.SendMessage(ctx, 123, "привет!"); err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		fmt.Println("Message sent successfully")
	}

	// Удаление сообщения
	if err := client.DeleteMessage(ctx, 123); err != nil {
		log.Printf("Failed to delete message: %v", err)
	} else {
		fmt.Println("Message deleted successfully")
	}

	// Установка статуса "печатает"
	if err := client.SetTyping(ctx, 123); err != nil {
		log.Printf("Failed to set typing: %v", err)
	} else {
		fmt.Println("Typing status set successfully")
	}

	// Установка статуса оффлайн
	if err := client.SetOffline(ctx); err != nil {
		log.Printf("Failed to set offline: %v", err)
	} else {
		fmt.Println("Offline status set successfully")
	}

	// Получение уведомлений
	notifications, err := client.GetNotifications(ctx)
	if err != nil {
		log.Fatalf("Failed to get notifications: %v", err)
	}
	fmt.Printf("Notifications: %+v\n\n", notifications)
}
