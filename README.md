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
