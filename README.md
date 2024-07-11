# CHECKPROXY

CheckProxy - это инструмент для проверки прокси-серверов, написанный на Go.

## Project Structure
```
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
| |
│ └── test                       test.(Тестирование)
│ └── checkproxy_test.go
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
```
## Config structure
Основной конфиг находится в **checkproxy/internal/repository/config/config.yml**, выглядит следующим образом
```
```yaml
db:
  host: "127.0.0.1"
  port: 3306
  user: "test"
  password: "test"
  name: "test"
  ip2l: "Путь до базы IP2Location"
  mm: "Путь до базы MaxMind"
  ```

## Run locally
Для работы потребуется версия Go >= **1.22.3**
### Сборка 
Собрать можно командой ```make build```

### Тестирование
Запустить тестирование можно командой ```make it``` , убедитесь что у вас установле **Docker** и **Docker-Compose**



### Запуск в CLI
```sh
go run cmd/cli/main.go -c checkproxy/internal/repository/config/config.yml
```
Путь передается с помощью ключа ```-c``` или указывается в переменной ```PROXY_GEO_CONFIG``` на системном уровне, или в файле **Makefile** с последующим запуском ```make run```
