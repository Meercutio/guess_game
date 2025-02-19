# 1) Базовый образ для сборки (Stage: builder)
FROM golang:1.22-alpine AS builder

# Создадим рабочую директорию внутри контейнера
WORKDIR /app

# Скопируем go.mod и go.sum, чтобы заранее загрузить зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем наше приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /guess-game ./cmd/server

# 2) Минимальный образ для запуска (Stage: runner)
FROM alpine:3.17

# Создаём рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарник из builder
COPY --from=builder /guess-game /app/guess-game

# Указываем, что контейнер будет слушать на порту 8080
EXPOSE 8080

# Команда по умолчанию: запуск нашего приложения
CMD ["/app/guess-game"]
