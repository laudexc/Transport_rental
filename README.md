# Структура всех файлов:
rental-system/
├── cmd/
│   └── main.go                   # точка входа (CLI / тестовый запуск)
│
├── internal/
│   ├── transport/                # всё, что связано с транспортом
│   │   ├── interfaces/
│   │   │   ├── vehicle.go        # базовый интерфейс транспорта
│   │   │   ├── rentable.go       # всё, что можно арендовать
│   │   │   └── repairable.go     # всё, что можно чинить
│   │   ├── entities/
│   │   │   ├── car.go
│   │   │   ├── boat.go
│   │   │   ├── plane.go
│   │   │   └── bike.go
│   │   ├── services/
│   │   │   ├── maintenance.go    # ремонт и обслуживание
│   │   │   └── rental.go         # базовая логика аренды (внутри транспорта)
│   │   └── storage/
│   │       └── memory.go         # хранилище состояния
│   │
│   ├── users/                    # всё про клиентов
│   │   ├── entities/
│   │   │   └── client.go         # структура клиента
│   │   ├── services/
│   │   │   ├── auth.go           # регистрация/поиск клиента
│   │   │   └── rental.go         # связывает клиента и транспорт
│   │   └── storage/
│   │       └── memory.go
│   │
│   ├── genid/
│   │   └── genid.go              # генерация уникальных ID (uuid / int64 / hash)
│   │
│   └── common/
│       └── errors.go             # общие ошибки, типы, утилиты
│
├── go.mod
└── go.sum
***
