package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simple-ozohub-prjct/internal/api"
	"simple-ozohub-prjct/internal/client"
	"simple-ozohub-prjct/internal/config"

	"github.com/gorilla/mux"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации: ", err)
	}

	// Инициализируем подключение к базе данных
	if err := config.InitDB(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer config.DB.Close() // Закрываем соединение с БД при завершении программы
	fmt.Println("Успешно подключено к базе данных.")

	// Инициализируем клиент
	client.InitializeClient(cfg.APIKey, cfg.ClientID)

	// Настраиваем маршрутизацию
	r := mux.NewRouter()
	r.HandleFunc("/products/{product_id}", api.GetProductHandler).Methods("GET")
	r.HandleFunc("/products/list", api.GetListOfProductsHandler).Methods("GET")
	r.HandleFunc("/products/list/", api.GetListOfProductsHandler).Methods("GET")

	// Создаём сервер
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Канал для обработки системных сигналов для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()
	fmt.Println("Сервер запущен на порту 8080")

	// Ожидаем сигнал завершения
	<-stop
	fmt.Println("\nПолучен сигнал завершения, закрытие сервера...")

	// Контекст завершения с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Завершение работы сервера
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка завершения работы сервера: %v", err)
	}
	fmt.Println("Сервер успешно остановлен.")
}
