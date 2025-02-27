## image сервис на Go 

## Требования.

- **Go**
- **Docker**: Docker используется для деплоя сервисов с помощью `docker-compose`.
- **Компилятор Protobuf (protoc)**: Необходим для генерации Go-кода из `.proto` файлов.

## Настройка окружения.

Проект использует переменные окружения для конфигурации различных сервисов. Шаблонный файл переменных окружения предоставлен как `.env.example`. 
Перед запуском проекта необходимо создать файл `.env` и заполнить его реальными значениями.

## Отредактируйте файл `.env`** и замените значения переменных на реальные:

   ```env
    LOG_LEVEL=debug   # Уровень вывода логов: info,debug,error 
    GRPC_PORT=50051  # Порт для gRPC сервера
    PG_HOST=postgres #Host PostgreSQL
    PG_PORT=5432 #Порт PostgreSQL
    PG_USER=template_service  # Имя пользователя PostgreSQL
    PG_PASSWORD=template_service  # Имя пользователя PostgreSQL
    PG_DATABASE=template_service # Имя базы данных PostgreSQL
   ```

## Структура проекта

Проект организован по следующим директориям:

###  `cmd/service/`

Содержит точку входа в приложение — файл `main.go`. Здесь определяется основной запуск сервиса.

- **`main.go`**: Главный файл, отвечающий за запуск и конфигурацию приложения.

###  `config/`

Содержит файлы конфигурации проекта.

- **`config.go`**: Здесь определяется структура конфигурационных данных и логика их загрузки.

###  `db/`

Содержит файлы, связанные с базой данных.

- **`migrations/`**: Содержит SQL файлы миграций базы данных. 
- 
###  `deploy/`

Содержит файлы, связанные с деплоем приложения с помощью Docker.

- **`docker-compose/`**:
    - **`docker-compose.yaml`**: Конфигурация для запуска сервисов через Docker Compose.
    - **`Dockerfile`**: Файл для сборки Docker-образа приложения.

###  `internal/`

Содержит основные пакеты сервиса, которые не должны быть доступны извне.

- **`application`**:
  **Слой приложения**: 

    - grpc/ - Реализация gRPC сервера
    - dto/ -  Обработка DTO

- **`models` && `modules`**:
  **Бизнес-логика и сущности **: 

    - models/ - Определение доменных сущностей.
    - modules/ - Реализация бизнес-логики и сценариев взаимодействия с БД

- **`generated`**:
  **Содержит автоматически сгенерированный код.**

###  `pkg/`

**Содержит вспомогательные пакеты для работы с БД и логгером**

### `schemas/proto/`

Содержит исходные файлы `.proto` для генерации gRPC и Protobuf кода.

### Makefile

- **build - сборка исполняемого файла в корне проекта*
- **lint - запуск линтера golangci-lint**
- **docker-build - сборка образа приложения(имя образа задается командой)** 
- **run - запуск всех контейнеров в фоновов режиме**
- **stop - остановка всех контейнеров в фоновом режиме**

### Запуск приложения

   1. git clone
   2. Копирование файла .env.example и внесение актуальных переменных окружения
   3. Сборка образа приложения make docker-build
   4. Запуск всех контейнеров make run  
   5. Остановка контейнеров через make stop 
      


