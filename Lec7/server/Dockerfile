# 1. Ступень сборки
# Сборка на основе goruntime
FROM golang:alpine as builder 


# Обновляю конфиги слепка ОС в конейтенере
# И загружаю гит, для поддержки go.mod
RUN apk update 
RUN apk upgrade 
RUN apk add --no-cache git 

# СОздаем рабочую директорию проекта в контейнере
RUN mkdir /app 
# Переходим в рабочую директорию в контейнере
WORKDIR /app

# Копируем все из server в /app
COPY . .

COPY go.mod .
COPY go.sum . 

# Позволяет получить все зависимости из go.mod
RUN go mod download 
# CGO_ENABLED - позволяет использовать/не использовать компилятор Cgo
# для сборки низкоуровненвых библиотек языка go написанных на C

# GOOS - под какое окружение используем компилятор языка Go
# go build -a -installsuffix cgo - собрать проект с учетом обхода cgo
# -o server - как назвать выходный бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server 
# На этом этапе мы имеем ГОРУ ненужного мусора в goruntime + server.bin


# 2. Ступень сборки
# Подготавливаем рабочий контейнер 
# В котором будет минимум функционала необходимого для выполнения .bin файла
FROM alpine:latest 

RUN apk --no-cache add ca-certificates 
# Настраиваем рабочее месо в новом контейнере
RUN mkdir /app
# Переходим в рабочее место
WORKDIR /app 
# Хотим перебросить из контейнера builder
# Копируем из builder наш erver.bin

COPY --from=builder /app/server .

# Запускаем бинарник
CMD ["./server"]