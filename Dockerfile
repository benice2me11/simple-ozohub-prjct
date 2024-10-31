# Используем базовый образ Golang
FROM golang:1.21-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Аргументы сборки
ARG API_KEY
ARG CLIENT_ID

# Устанавливаем переменные окружения на этапе сборки
ENV API_KEY=$API_KEY
ENV CLIENT_ID=$CLIENT_ID

# Копируем и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код и собираем приложение
COPY . .
RUN go build -o main ./cmd

# Финишный образ
FROM golang:1.21-alpine

# Создаем директорию и копируем бинарник из builder
WORKDIR /app
COPY --from=builder /app/main .

# Устанавливаем переменные окружения
ENV API_KEY=$API_KEY
ENV CLIENT_ID=$CLIENT_ID

# Запуск приложения
CMD ["./main"]
