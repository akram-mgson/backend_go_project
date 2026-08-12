# Backend — личный кабинет клиента

## Требования
Go 1.23+

## Запуск
bash
go run .


# Тесты

go test ./...
go test -race ./...

# Проверка кода

gofmt -l *.go
go vet ./...

# Авторизация

Authorization: Bearer fake-token


# Основные эндпоинты

POST /api/cabinet/auth/login — логин
POST /api/cabinet/auth/logout — выход
GET /api/cabinet/auth/me — текущий пользователь
GET /api/cabinet/orders — список заказов
GET /api/cabinet/orders/{id} — карточка заказа
GET /api/cabinet/profile — профиль
POST /api/cabinet/orders/{id}/comment — комментарий
GET /api/cabinet/orders/{id}/documents — документы