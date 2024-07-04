```markdown
# CHECKPROXY
## Project Structure
checkproxy
├── cmd                          cmd.(Точка входа)
│ └── cli
│ └── main.go
│
├── internal                     internal.(Инкапсуляция и защита данных)
│ │
│ ├── app                        app.(Логика приложения)
│ │ └── app.go
│ │
│ ├── repository                 repository.(Изоляция и хранение данных)
│ │ ├── config                   config.(Хранение различных конфигов и БД)
│ │ │ │
│ │ │ ├── GeoLite2-Country.mmdb
│ │ │ ├── IP-COUNTRY-SAMPLE.BIN 
│ │ │ └── config.yml
│ │ │
│ │ └── proxy                    proxy.(Запрос в БД)
│ │ └── allproxies.go
│ │
│ └── service                    service.(Бизнес логика)
│ └── checkproxy.go
│
└── pkg                          pkg.(Вспомогательные утилиты и зависимости)
├── config                       config.(Взаимодействие с конфигом)
│ │
│ ├── config.go
│ └── path.go
│
├── db                           db.(Структура конфига, Открытие баз)
│ │
│ ├── dbsource.go
│ ├── openip2ldb.go
│ └── openmmdb.go
│
└── mysql                        mysql.(Подключение к БД)
└── connector.go
  
