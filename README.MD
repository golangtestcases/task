# Task API

HTTP API для управления длительными I/O bound задачами (3-5 минут выполнения)

# Особенности
- Чистая архитектура
- Хранение данных в памяти
- Поддержка операций: создание, проверка статуса, удаление задач
- Автоматическое обновление статусов задач (pending → processing → completed)

# Запуск сервиса
go run cmd/server/main.go

# Клонирование репозитория
git clone https://github.com/golangtestcases/task.git
cd task

# Запуск сервера
go run cmd/server/main.go

# API Endpoints
POST	/api/v1/tasks	Создать новую задачу
GET	/api/v1/tasks/{id}	Получить статус задачи
DELETE	/api/v1/tasks/{id}	Удалить задачу

# Примеры запросов
## Создание задачи:
POST http://localhost:8080/api/v1/tasks

Ответ:

json
{
    "ID": "3a9f1c7a-5e8f-4b2d-9c3a-1e6f5d2b8c7a",
    "Status": "pending",
    "CreatedAt": "2023-05-20T14:30:45Z",
    "StartedAt": null,
    "CompletedAt": null,
    "Result": null,
    "Error": null,
    "DurationMs": 0
}
## Проверка статуса

curl http://localhost:8080/api/v1/tasks/3a9f1c7a-5e8f-4b2d-9c3a-1e6f5d2b8c7a
# Возможные статусы:

- pending - задача создана

- processing - задача выполняется

- completed - задача успешно завершена

## Удаление задачи

DELETE http://localhost:8080/api/v1/tasks/3a9f1c7a-5e8f-4b2d-9c3a-1e6f5d2b8c7a

# Последовательность тестирования:

- Создайте задачу (POST)

- Проверяйте статус (GET) каждые 30 секунд

- Удалите задачу после завершения (DELETE)