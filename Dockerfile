FROM golang as builder

WORKDIR /app

# Копируем исходники и файлы зависимостей
COPY ./src ./src
COPY go.mod ./
COPY go.sum ./

# Устанавливаем зависимости
RUN go mod tidy

# Сборка исполняемого файла
RUN go build -o tender-service ./src/main.go

# Создаем финальный образ
FROM golang
WORKDIR /app

# Копируем скомпилированный бинарный файл
COPY --from=builder /app/tender-service .
COPY .env .

# Открываем порт для сервера
EXPOSE 8080

# Запуск сервиса
CMD ["./tender-service"]
