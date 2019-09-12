#An example of the rest api collector. 
It receives data from REST API from a service and save data in the postgres. 
The receiving is scheduled via cron packet. 
And through web api is possible to update data.

## Scheduled tasks
* MasterTask

## Api methods
* `/stat/{date-from}/{date-to}/` - manual run MasterTask task for any period (date examples: `YYYY-MM-DD|today|yesterday|day-before-yesterday|week-ago|two-week-ago|three-week-ago|month-ago|two-month-ago|quarter-ago|half-year-ago|year-ago`)

---
## Installation
To install collector, add lines to ~/.gitconfig:
```
[url "ssh://git@gitlab.domain.com"]
insteadOf = https://gitlab.domain.com
```
then execute:
```sh
go build
```
To running migration, execute:
```sh
cd migrations/
go build

export PG_RAW_DSN="host=localhost port=5432 user=postgres dbname=postgres password="
export MAIL_SERVER_ADDR="localhost:25"
export ERR_MAIL_FROM="sender@domain"
export ERR_MAIL_TO="recipient@domain"
export LOG_FORMAT="text"
export LOG_LEVEL="debug"

./migrations init
./migrations  
```
To running collector, execute:
```sh
export PG_RAW_DSN="host=localhost port=5432 user=postgres dbname=postgres password="
export PG_REMTRACK_DSN="host=localhost port=5432 user=postgres dbname=postgres password="
export MAIL_SERVER_ADDR="localhost:25"
export ERR_MAIL_FROM="sender@domain"
export ERR_MAIL_TO="recipient@domain"
export LISTEN_ADDR="127.0.0.1:8080"
export LOG_FORMAT="json"
export LOG_LEVEL="info"

./collector 
```
To running all tests, execute:
```sh
go test ./...
```
---
## Configuration environment variables

* PG_RAW_DSN - dsn for connection to raw "database" in Postgres 
* PG_REMTRACK_DSN - dsn for connection to "remtrack" database in Postgres
* MAIL_SERVER_ADDR - mail server host with port for errors messages
* ERR_MAIL_FROM - sender email for errors
* ERR_MAIL_TO - recipient email for errors
* LISTEN_ADDR - ip and port for web api
* LOG_FORMAT - log output format to stdout [text|json]
* LOG_LEVEL - log verbosity [panic|fatal|error|warn|info|debug|trace]

##Dependencies
* `github.com/go-pg/pg` - PostgresSQL ORM
* `github.com/sirupsen/logrus` - Logger with hooks
* `github.com/gorilla/mux` - Lightweight router for web server
* `github.com/jasonlvhit/gocron`- scheduler (crontab)
* `github.com/stretchr/testify/assert` - asserts for tests

