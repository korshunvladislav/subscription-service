# Subscription Management Service

Простое веб-приложение на **Go (Golang)** для управления подписками.  
Проект поддерживает два режима запуска: **development** и **production** (через Docker).

---

## 🚀 Функциональность

- Подключение к базе данных PostgreSQL
- REST API для управления подписками
- Миграции базы данных
- Работа в двух режимах:
    - **Dev** — локальная разработка со сборкой из `dev.dockerfile`
    - **Prod** — использование контейнера из DockerHub и `prod.dockerfile`

---

## 📂 Структура проекта

```bash
.
├── controllers/                # Контроллеры (обработка запросов)
│   └── subscriptionsController.go
│
├── initializers/               # Инициализация приложения
│   ├── database.go             # Подключение к БД
│   └── loadEnvVariables.go     # Загрузка переменных окружения
│
├── migrate/                    # Миграции базы данных
│   └── migrate.go
│
├── models/                     # Определение моделей
│   └── subscriptionModel.go
│
├── utils/                      # Утилиты
│   └── utils.go
│
├── vendor/                     # Зависимости (go mod vendor)
│
├── .env                        # Переменные окружения
├── .gitignore                  # Игнорируемые файлы Git
├── dev.dockerfile              # Dockerfile для разработки
├── prod.dockerfile             # Dockerfile для продакшена
├── docker-compose.dev.yml      # docker-compose для разработки
├── docker-compose.prod.yml     # docker-compose для продакшена
├── go.mod                      # Go модули
├── go.sum
├── main.go                     # Точка входа в приложение
```

---

## ⚙️ Установка и запуск

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/korshunvladislav/subscription-service
cd subscription-service
```

### 2. Создайте файл .env

Пример содержимого:
```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=subscriptions
PORT=8000
```

---

## 🛠 Режим разработки

Сборка из dev.dockerfile:
```bash
docker compose -f docker-compose.dev.yml up --build
```
API будет доступно по адресу:
```
http://localhost:8000
```

---

## 🏗 Продакшен режим

Используется образ с DockerHub и prod.dockerfile
```bash
docker compose -f docker-compose.prod.yml up --build
```

---

## 🗄 Миграции базы данных

Запуск миграций
```bash
go run migrate/migrate.go
```

---

## 📌 API Endpoints (пример)

- GET /subscriptions — список подписок
- GET /subscriptions/{id} — получить подписку по ID
- POST /subscriptions — создать новую подписку
- PUT /subscriptions/{id} — обновить подписку
- DELETE /subscriptions/{id} — удалить подписку
- GET /subscriptions/summary — получить сумму подписок за указанный период

---

## 🛠 Технологии

- Go 1.22+
- PostgreSQL
- Docker / Docker Compose
- GORM (ORM для работы с БД)