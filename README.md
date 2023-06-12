# Backend Master

## Project Name: Simple Bank

- Create and manage account (owner, balance, currency)
- Record all balance changes (create an account entry for each change)
- Money transfer transaction (perform money transfer between 2 accounts consistently within a transaction)

## Database Design using [dbdiagram](https://dbdiagram.io)

<details>
<summary>View contents</summary>

Design database tables using <https://dbdiagram.io>

### Setup dbdigram

```sh
# install dbdocs
npm i -g dbdocs

# check dbdocs
dbdocs

# create doc directory
mkdir doc

# copy dbml codes and paste in db.dbml file
# install "vscode-dbml" extension to highlight codes
touch db.dbml

# login to dbdocs
dbdocs login

# generate dbdocs view
dbdocs build doc/db.dbml

# visit: https://dbdocs.io/foyezar/simplebank

# set password
# dbdocs password --set <password> --project <project name>
dbdocs password --set secret --project simplebank

# remove a project
# dbdocs remove <project name>
dbdocs remove simplebank

# install dbml cli
npm i -g @dbml/cli

# convert a dbml file to sql
# dbml2sql <path-to-dbml-file> [--mysql|--postgres] -o <output-filepath>
dbml2sql doc/db.dbml --postgres -o doc/schema.sql

# convert a sql file to dbml
# sql2dbml <path-to-sql-file> [--mysql|--postgres] -o <output-filepath>
sql2dbml doc/schema.sql --postgres -o doc/db.dbml
```

</details>

## Install & use Docker + Postgres + Mysql + TablePlus

<details>
<summary>View contents</summary>

- Download & install docker: [link](https://docs.docker.com/desktop/install/mac-install)

Postgresql

```sh
# Pull postgres image
docker pull postgres:15:2-alpine

# Start postgres container
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=testpass -d postgres:15.2-alpine

# Run command in container
docker exec -it postgres15 psql -U root

# Test connection
SELECT now();
```

Postgres commands <sup>[ref](https://hasura.io/blog/top-psql-commands-and-flags-you-need-to-know-postgresql/)</sup>

```sh
# Connect to a database (same host)
# -W - forces for the user password
psql -d db_name -U username -W

# Connect to a database (different host)
psql -h db_address -d db_name -U username -W

# Connect to a database (different host in SSL mode)
psql "sslmode=require host=db_address dbname=my_db user=root"

# Know all available psql commands
\?

# List all databases
\l

# Clear screen
# Ctrl + L
\! clear
\! cls

# Create a database
create database mydb;

# Switch to another database
\c db_name

# List database tables
\dt

# Create a table
CREATE TABLE accounts (
  id serial PRIMARY KEY,
  username varchar NOT NULL
);

# Insert data in a able
INSERT INTO accounts (username) VALUES ('foyez');

# Select data from a table
SELECT * FROM accounts;

# describe a table
\d table_name
\d+ table_name # more information

# Delete a database
drop database mydb;

# List all schemas
\dn

# List users and their roles
\du

# Retrieve a specific user
\du username

# Quit psql
\q
```

Mysql

```sh
# Pull mysql image
docker pull mysql:8

# Start mysql container
docker run --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=testpass -d mysql:8

# Run command in container
docker exec -it mysql8 mysql -uroot -ptestpass
```

Create a Postgres database from command line

```sh
# enter postgres shell & create a database
docker exec -it postres15 sh
createdb --username=root --owner=root simple_bank
dropdb simple_bank

# create a database
docker exec -it postres15 createdb --username=root --owner=root simple_bank

# login to db cli
docker exec -it postgres15 psql -U root simple_bank

# exit from db cli
\q
```

Mysql commands <sup>[ref](http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm)</sup>

```sh
# Connect to database
mysql -h hostname -u username -p
mysql -uroot -ptestpass

# Create a database from command line
mysql -e "create database db_name" -u username -p

# Create a database
create database db_name;

# Show database list
show databases;

# Switch to a database
use db_name;

# Show table list
show tables;

# Create a table
CREATE TABLE accounts (
  id INT(50) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(100) NOT NULL
);

# Insert data in a table
INSERT INTO accounts (username) VALUES ('foyez');

# Select data from a table
SELECT * FROM accounts;

# Describe a table
describe table_name;

# Delete a database
drop database db_name;

# Delete a table
drop table table_name;

# Quit mysql
exit;
```

Create a Mysql database from command line

```sh
# create a database
docker exec -it mysql8 mysql -e "create database db_name" -u username -p

# delete a database
docker exec -it mysql8 mysql -e "drop database db_name" -u username -p

# login to db cli
docker exec -it mysql8 mysql db_name -u username -p

# exit from db cli
\q
```

Show docker logs

```sh
# Postgres
docker logs postgres15

# Mysql
docker logs mysql8
```

Searching ran commands starting with `docker run`

```sh
history | grep "docker run"
```

- Download & install database management tool [TablePlus](https://tableplus.com/)

</details>

## Database Migration using [golang-migrate](https://github.com/golang-migrate/migrate)

<details>
<summary>View contents</summary>

Install migrate cli: [link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```sh
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
# OR
brew install golang-migrate

# migrate help command
migrate -help

# create migration files
migrate create -ext sql -dir db/migration -seq init_schema
```

</details>

## Database CRUD

<details>
<summary>View contents</summary>

- Create: insert new records to the database
- Read: select or search for records in the database
- Update: change some fields of the record in the database
- Delete: remove records from the database

### CRUD Tools

- Database/SQL: t
- ORM: GORM
- SQLX
- SQLC

### Setup [SQLC](https://sqlc.dev/)

```sh
# install sqlc
brew install sqlc

# to know sqlc commands
sqlc help

# Create an empty sqlc.yaml settings file
# schema_path: db/migration
# query path: db/query
# output path: db/sqlc
sqlc init

# Generate Go code from SQL
sqlc generate
```

</details>

## Unit tests for Database

<details>
<summary>View contents</summary>

- Install a pure postgres driver for Go's database/sql package

```sh
go get github.com/lib/pq
```

`main_test.go`

```go
package db

import (
 "database/sql"
 "log"
 "os"
 "testing"

 _ "github.com/lib/pq"
)

const (
 dbDriver = "postgres"
 dbSource = "postgresql://root:testpass@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
 db, err := sql.Open(dbDriver, dbSource)
 if err != nil {
  log.Fatal("cannot connect to db: ", err)
 }

 testQueries = New(db)

 os.Exit(m.Run())
}
```

- Run `go mod tidy` to add dependency in `go.mod` file
- Install [testify](https://github.com/stretchr/testify) - `A toolkit for assertions and mocks`

```sh
go get github.com/stretchr/testify
```

`account_test.go`

```go
package db

import (
 "context"
 "testing"

 "github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
 arg := CreateAccountParams{
  Owner:    "Mithu",
  Balance:  20,
  Currency: "USD",
 }
 account, err := testQueries.CreateAccount(context.Background(), arg)
 require.NoError(t, err)
 require.NotEmpty(t, account)

 require.Equal(t, account.Owner, arg.Owner)
 require.Equal(t, account.Balance, arg.Balance)
 require.Equal(t, account.Currency, arg.Currency)

 require.NotZero(t, account.ID)
 require.NotZero(t, account.CreatedAt)
}
```

- Run `go mod tidy` to add _testify_ dependency

</details>

## Database Transaction

<details>
<summary>View contents</summary>

### DB Transaction

- A single unit of work
- Often made up of multiple db operations

**Example:** Transfer 10 USD from bank account 1 to bank account 2.

```txt
1. Create a transfer record with amount = 10
2. Create an account entry for account 1 with amount = -10
3. Create an account entry for account 2 with amount = +10
4. Subtract 10 from the balance of account 1
5. Add 10 to the balance of account 2
```

### Why do we need db transaction?

1. To provide a reliable and consistent unit of work, even in case of system failure
2. To provide isolation between programs that access the database concurrently

A transaction in a database system must maintain **ACID** (Atomicity, Consistency, Isolation and Durability) in order to ensure accuracy, completeness and data integrity.

1. **Atomicity**
   Either all operations complete successfully or if the transaction fails, everything will be rolled back and the db will be unchanged.

2. **Consistency**
   The db state must be valid after the transaction. All constraints must be satisfied. More precisely, all data written to the database must be valid according to predefined rules, including constraints, cascade, and triggers.

3. **Isolation**
   Concurrent transaction must not affect each other.

4. **Durability**
   Data written by a successful transaction must be recorded in persistent storage, even in case of system failure.

### How to run SQL TX?

```sql
BEGIN;
COMMIT;

-- if the transaction is failed
BEGIN;
ROLLBACK;
```

### Deadlock

- a situation in which two or more transactions are waiting for one another to give up locks

Deadlocks can happen in multi-user environments when two or more transactions are running concurrently and try to access the same data in a different order. When this happens, one transaction may hold a lock on a resource that another transaction needs, while the second transaction may hold a lock on a resource that the first transaction needs. Both transactions are then blocked, waiting for the other to release the resource they need.

DBMSs often use various techniques to detect and resolve deadlocks automatically. These techniques include timeout mechanisms, where a transaction is forced to release its locks after a certain period of time, and deadlock detection algorithms, which periodically scan the transaction log for deadlock cycles and then choose a transaction to abort to resolve the deadlock.

It is also possible to prevent deadlocks by careful design of transactions, such as always acquiring locks in the same order or releasing locks as soon as possible. Proper design of the database schema and application can also help to minimize the likelihood of deadlocks

**ref:** [Deadlock in DBMS](https://www.geeksforgeeks.org/deadlock-in-dbms/)

### Update accounts concurrently

```sql
BEGIN;

SELECT * FROM accounts WHERE id = 15 FOR UPDATE;
UPDATE accounts SET balance = 500 WHERE id = 15;

COMMIT;
```

### Check deadlocks

```sql
SELECT
   a.application_name,
   l.relation::regclass,
   l.transactionid,
   l.mode,
   l.locktype,
   l.GRANTED,
   a.username,
   a.query,
   a.pid
FROM pq_stat_activity a
JOIN pg_locks l ON l.pid = a.pid
WHERE a.application_name = 'psql'
ORDER BY a.pid;
```

- [DB transaction lock & How to handle deadlock](https://www.youtube.com/watch?v=G2aggv_3Bbg&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=10)
- [How to avoid deadlock in DB transaction? Queries order matters!](https://www.youtube.com/watch?v=qn3-5wdOfoA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=8)

</details>

## Github Actions

<details>
<summary>View contents</summary>

- We can trigger a workflow by 3 ways: `event`, `schedule`, or `manually`
- A workflow consists of one or multiple jobs
- A job is composed of multiple steps
- Each step has one or more actions
- All jobs inside a workflow normally run in parallel, unless they depend on each other
- If some jobs depend on each other, they run serially
- Each job will be run separately by a specific runner
- The runners will report progress, logs, and results of the jobs back to github

<img width="1552" alt="image" src="https://github.com/foyez/simplebank/assets/11992095/5954c678-bdf0-45cc-bf84-7db9a383bf58">

### Setup a workflow for Golang and Postgres

- Goto `Actions` tab
- Then, in `Go` action click `configure`
- Create github workflows directory: `mkdir -p .github/workflows`
- Create workflow file: `touch .github/workflows/test.yml`
- Then, copy and paste the template from github for go
- [Creating PostgreSQL service containers](https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers)
- [How to setup Github Actions for Go + Postgres to run automated tests](https://dev.to/techschoolguru/how-to-setup-github-actions-for-go-postgres-to-run-automated-tests-81o)

</details>

## Implement HTTP API using Gin

<details>
<summary>View contents</summary>

### Popular web frameworks

- [Gin](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#build-with-json-replacement)
- Beego
- Echo
- Revel
- Martini
- Fiber
- Buffalo

### Popular HTTP routers

- FastHttp
- Gorilla Mux
- HttpRouter
- Chi

Install `gin` package:

```sh
https://github.com/gin-gonic/gin
```

- A POST api:

<details>
<summary>View contents</summary>

`db/query/account.sql`

```sql
-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;
```

`api/server.go`

```go
package api

import (
 db "github.com/foyez/simplebank/db/sqlc"
 "github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
 store  *db.Store
 router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
 server := &Server{store: store}
 router := gin.Default()

 router.POST("/accounts", server.createAccount)

 server.router = router
 return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
 return server.router.Run(address)
}

func errorResponse(err error) gin.H {
 return gin.H{"error": err.Error()}
}
```

`api/account.go`

```go
package api

import (
 "net/http"

 db "github.com/foyez/simplebank/db/sqlc"
 "github.com/gin-gonic/gin"
)

type createAccountRequest struct {
 // json tag to de-serialize json body
 Owner    string `json:"owner" binding:"required"`
 Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
 var req createAccountRequest
 if err := ctx.ShouldBindJSON(&req); err != nil {
  ctx.JSON(http.StatusBadRequest, errorResponse(err))
  return
 }

 arg := db.CreateAccountParams{
  Owner:    req.Owner,
  Currency: req.Currency,
  Balance:  0,
 }

 account, err := server.store.CreateAccount(ctx, arg)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, errorResponse(err))
  return
 }

 ctx.JSON(http.StatusCreated, account)
}
```

`main.go`

```go
package main

import (
 "database/sql"
 "log"

 "github.com/foyez/simplebank/api"
 db "github.com/foyez/simplebank/db/sqlc"
 _ "github.com/lib/pq"
)

const (
 dbDriver = "postgres"
 dbSource = "postgresql://root:testpass@localhost:5432/simplebank?sslmode=disable"
 address  = "0.0.0.0:8080"
)

func main() {
 conn, err := sql.Open(dbDriver, dbSource)
 if err != nil {
  log.Fatal("cannot connect to db: ", err)
 }

 store := db.NewStore(conn)
 server := api.NewServer(store)

 err = server.Start(address)
 if err != nil {
  log.Fatal("cannot start server: ", err)
 }
}
```

</details>

- A GET api:

<details>
<summary>View contents</summary>

`db/query/account.sql`

```sql
-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;
```

`api/account.go`

```go
type getAccountRequest struct {
 ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
 var req getAccountRequest
 if err := ctx.ShouldBindUri(&req); err != nil {
  ctx.JSON(http.StatusBadRequest, errorResponse(err))
  return
 }

 account, err := server.store.GetAccount(ctx, req.ID)
 if err != nil {
  if err == sql.ErrNoRows {
   ctx.JSON(http.StatusNotFound, errorResponse(err))
   return
  }
  ctx.JSON(http.StatusInternalServerError, errorResponse(err))
  return
 }

 ctx.JSON(http.StatusOK, account)
}
```

`api/server.go`

```go
router.GET("/accounts/:id", server.getAccount)
```

API:

```txt
http://localhost:8080/accounts/1
```

</details>

- A GET api with offset pagination:

<details>
<summary>View contents</summary>

`db/query/account.sql`

```sql
-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;
```

`api/account.go`

```go
type listAccountsRequest struct {
 PageID   int32 `form:"page_id" binding:"required,min=1"`
 PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
 var req listAccountsRequest
 if err := ctx.ShouldBindQuery(&req); err != nil {
  ctx.JSON(http.StatusBadRequest, errorResponse(err))
  return
 }

 arg := db.ListAccountsParams{
  Limit:  req.PageSize,
  Offset: (req.PageID - 1) * req.PageSize,
 }

 accounts, err := server.store.ListAccounts(ctx, arg)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, errorResponse(err))
  return
 }

 ctx.JSON(http.StatusOK, accounts)
}
```

`api/server.go`

```go
router.GET("/accounts", server.listAccounts)
```

API:

```txt
http://localhost:8080/accounts?page_id=1&page_size=10
```

</details>

- A GET api with cursor pagination:

<details>
<summary>View contents</summary>

`db/query/account.sql`

```sql
-- name: ListAccountWithCursor :many
SELECT * FROM accounts
WHERE created_at < sqlc.narg('cursor') OR sqlc.narg('cursor') IS NULL
ORDER BY created_at DESC
LIMIT sqlc.arg('limit');
```

`api/account.go`

```go
type listAccountsWithCursorRequest struct {
 Cursor time.Time `form:"cursor"`
 Limit  int32     `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) listAccountsWithCursor(ctx *gin.Context) {
 var req listAccountsWithCursorRequest
 if err := ctx.ShouldBindQuery(&req); err != nil {
  ctx.JSON(http.StatusBadRequest, errorResponse(err))
  return
 }

 limitPlusOne := req.Limit + 1

 arg := db.ListAccountWithCursorParams{
  Limit: limitPlusOne,
  Cursor: sql.NullTime{
   Time:  req.Cursor,
   Valid: !req.Cursor.IsZero(),
  },
 }

 accounts, err := server.store.ListAccountWithCursor(ctx, arg)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, errorResponse(err))
  return
 }

 newAccounts := accounts
 if int32(len(accounts)) > req.Limit {
  newAccounts = accounts[0:req.Limit]
 }

 rsp := gin.H{
  "accounts": newAccounts,
  "has_more": int32(len(accounts)) == limitPlusOne,
 }

 ctx.JSON(http.StatusOK, rsp)
}
```

`api/server.go`

```go
router.GET("/accountsWithCursor", server.listAccountsWithCursor)
```

API:

```txt
http://localhost:8080/accountsWithCursor?limit=5&cursor=2023-06-05T02%3A36%3A19.167614Z
```

</details>

- [Gin binding in Go: A tutorial with examples](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/)
- [Build RESTful API using Go Gin](https://www.golinuxcloud.com/golang-gin/)

</details>

## Load config from file & environment variables with [Viper](https://github.com/spf13/viper)

<details>
<summary>View contents</summary>

Install viper:

```sh
go get github.com/spf13/viper
```

`app.env`

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:testpass@localhost:5432/simplebank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
```

`util/config.go`

```go
package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
type Config struct {
 DBDriver      string `mapstructure:"DB_DRIVER"`
 DBSource      string `mapstructure:"DB_SOURCE"`
 ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
 viper.AddConfigPath(path)
 viper.SetConfigName("app")
 viper.SetConfigType("env")

 viper.AutomaticEnv()

 err = viper.ReadInConfig()
 if err != nil {
  return
 }

 err = viper.Unmarshal(&config)
 return
}
```

`main.go`

```go
package main

import (
 "database/sql"
 "log"

 "github.com/foyez/simplebank/util"
 _ "github.com/lib/pq"
)

func main() {
 config, err := util.LoadConfig(".")
 if err != nil {
  log.Fatal("cannot load config: ", err)
 }

 conn, err := sql.Open(config.DBDriver, config.DBSource)

//  ...
}
```

</details>

## Mock DB for testing HTTP API

<details>
<summary>View contents</summary>

[Source](https://www.youtube.com/watch?v=rL0aeMutoJ0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=13)

### Why mock database?

- Independent tests: avoid conflicts
- Faster tests
- 100% coverage: easily setup edge cases

Install [gomock](https://github.com/golang/mock):

```sh
go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/golang/mock/mockgen@v1.6.0
```

`db/sqlc/store.go`

```go
package db

import (
 "context"
 "database/sql"
)

// Store provides all the function to exec
type Store interface {
 Querier
 TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functionalities to execute SQL queries and transaction
type SQLStore struct {
 *Queries
 db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
 return &SQLStore{
  db:      db,
  Queries: New(db),
 }
}
```

Generate mock interfaces:

```sh
mockgen -package mockdb -destination db/mock/store.go github.com/foyez/simplebank/db/sqlc Store
```

`api/account_test.go`

```go
package api

import (
 "bytes"
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "net/http/httptest"
 "testing"

 mockdb "github.com/foyez/simplebank/db/mock"
 db "github.com/foyez/simplebank/db/sqlc"
 "github.com/foyez/simplebank/util"
 "github.com/golang/mock/gomock"
 "github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
 account := randomAccount()

 testCases := []struct {
  name          string
  accountID     int64
  buildStubs    func(store *mockdb.MockStore)
  checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
 }{
  {
   name:      "OK",
   accountID: account.ID,
   buildStubs: func(store *mockdb.MockStore) {
    store.EXPECT().
     GetAccount(gomock.Any(), gomock.Eq(account.ID)).
     Times(1).
     Return(account, nil)
   },
   checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
    require.Equal(t, http.StatusOK, recorder.Code)
    requireBodyMatchAccount(t, recorder.Body, account)
   },
  },
 }

 for i := range testCases {
  tc := testCases[i]

  t.Run(tc.name, func(t *testing.T) {
   ctrl := gomock.NewController(t)
   defer ctrl.Finish()

   store := mockdb.NewMockStore(ctrl)
   tc.buildStubs(store)

   // start test server and send request
   server := NewServer(store)
   recorder := httptest.NewRecorder()

   url := fmt.Sprintf("/accounts/%d", account.ID)
   request, err := http.NewRequest(http.MethodGet, url, nil)
   require.NoError(t, err)

   server.router.ServeHTTP(recorder, request)
   tc.checkResponse(t, recorder)
  })

 }
}

func randomAccount() db.Account {
 return db.Account{
  ID:       util.RandomInt(1, 1000),
  Owner:    util.RandomOwner(),
  Balance:  util.RandomMoney(),
  Currency: util.RandomCurrency(),
 }
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
 data, err := ioutil.ReadAll(body)
 require.NoError(t, err)

 var gotAccount db.Account
 err = json.Unmarshal(data, &gotAccount)
 require.NoError(t, err)
 require.Equal(t, account, gotAccount)
}
```

`api/main_test.go`

```go
package api

import (
 "os"
 "testing"

 "github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
 gin.SetMode(gin.TestMode)
 os.Exit(m.Run())
}
```

</details>

## Dockerize the app

<details>
<summary>View contents</summary>

`Dockerfile`

```Dockerfile
# Build stage
FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
```

Build and Run docker container:

```sh
# build image
docker build -t foyezar/simplebank:latest .

# run container
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release foyezar/simplebank:latest
```

### Connect postgres with simplebank using network

<details>
<summary>View contents</summary>

```sh
# Get container details information
docker container inspect postgres15
```

```json
[
  {
    "Id": "efebd1beb2f417887655d482767644a4c816a28154ba1bb8f5cd1cb5cf2ad150",
    "Created": "2023-05-22T06:54:07.034351837Z",
    "Path": "docker-entrypoint.sh",
    "Args": ["postgres"],
    "Name": "/postgres15",
    "NetworkSettings": {
      "Ports": {
        "5432/tcp": [
          {
            "HostIp": "0.0.0.0",
            "HostPort": "5432"
          }
        ]
      },
      "Gateway": "172.17.0.1",
      "IPAddress": "172.17.0.2",
      "MacAddress": "02:42:ac:11:00:02",
      "Networks": {
        "bridge": {
          "Gateway": "172.17.0.1",
          "IPAddress": "172.17.0.2"
        }
      }
    }
  }
]
```

```sh
docker container inspect simplebank
```

```json
[
  {
    "Id": "928356064d037aef04f027b4bef1580b4381866cd8eb0cb02fd9b9675772eb26",
    "Created": "2023-06-12T17:30:41.745803471Z",
    "Path": "/app/main",
    "Name": "/simplebank",
    "NetworkSettings": {
      "Ports": {
        "8080/tcp": [
          {
            "HostIp": "0.0.0.0",
            "HostPort": "8080"
          }
        ]
      },
      "Gateway": "172.17.0.1",
      "IPAddress": "172.17.0.3",
      "MacAddress": "02:42:ac:11:00:03",
      "Networks": {
        "bridge": {
          "Gateway": "172.17.0.1",
          "IPAddress": "172.17.0.3"
        }
      }
    }
  }
]
```

Here, the IP address of `postgres15` container (`172.17.0.2`) is different than the IP address of `simplebank` container (`172.17.0.3`).

```sh
# Get network list
docker network ls

# Get network details information
docker network inspect bridge
```

```json
[
  {
    "Name": "bridge",
    "Driver": "bridge",
    "Containers": {
      "928356064d037aef04f027b4bef1580b4381866cd8eb0cb02fd9b9675772eb26": {
        "Name": "simplebank",
        "MacAddress": "02:42:ac:11:00:03",
        "IPv4Address": "172.17.0.3/16",
        "IPv6Address": ""
      },
      "efebd1beb2f417887655d482767644a4c816a28154ba1bb8f5cd1cb5cf2ad150": {
        "Name": "postgres15",
        "MacAddress": "02:42:ac:11:00:02",
        "IPv4Address": "172.17.0.2/16",
        "IPv6Address": ""
      }
    }
  }
]
```

```sh
# Get docker network COMMAND
docker network --help

# Create a network
docker network create bank-network

# Get docker network connect COMMAND
docker network connect --help

# Connect a container with a network
docker network connect bank-network postgres15

# Get network details information
docker network inspect bank-network
```

```json
[
  {
    "Name": "bank-network",
    "Scope": "local",
    "Driver": "bridge",
    "Containers": {
      "efebd1beb2f417887655d482767644a4c816a28154ba1bb8f5cd1cb5cf2ad150": {
        "Name": "postgres15",
        "MacAddress": "02:42:ac:12:00:02",
        "IPv4Address": "172.18.0.2/16",
        "IPv6Address": ""
      }
    }
  }
]
```

```sh
# Get container details information
docker container inspect postgres15
```

```json
[
  {
    "Name": "/postgres15",
    "NetworkSettings": {
      "Gateway": "172.17.0.1",
      "IPAddress": "172.17.0.2",
      "MacAddress": "02:42:ac:11:00:02",
      "Networks": {
        "bank-network": {
          "Gateway": "172.18.0.1",
          "IPAddress": "172.18.0.2",
          "MacAddress": "02:42:ac:12:00:02"
        },
        "bridge": {
          "Gateway": "172.17.0.1",
          "IPAddress": "172.17.0.2",
          "MacAddress": "02:42:ac:11:00:02"
        }
      }
    }
  }
]
```

```sh
# Run a container in a specific network
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:testpass@postgres15:5432/simplebank?sslmode=disable" foyezar/simplebank:latest
```

```sh
# Get network details information
docker network inspect bank-network
```

```json
[
  {
    "Name": "bank-network",
    "Scope": "local",
    "Driver": "bridge",
    "Containers": {
      "bc239683e762e39e6d3d368f16c377ddacc3e6a02e6f0efd5c50bf8aed138ded": {
        "Name": "simplebank",
        "MacAddress": "02:42:ac:12:00:03",
        "IPv4Address": "172.18.0.3/16"
      },
      "efebd1beb2f417887655d482767644a4c816a28154ba1bb8f5cd1cb5cf2ad150": {
        "Name": "postgres15",
        "MacAddress": "02:42:ac:12:00:02",
        "IPv4Address": "172.18.0.2/16"
      }
    }
  }
]
```

</details>

</details>
