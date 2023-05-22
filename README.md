# Backend Master

## Project Name: Simple Bank

- Create and manage account (owner, balance, currency)
- Record all balance changes (create an account entry for each change)
- Money transfer transaction (perform money transfer between 2 accounts consistently within a transaction)

## Database Design using <https://dbdiagram.io>

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

- Download & install docker from <https://docs.docker.com/desktop/install/mac-install/>

Postgresql

```sh
# Pull postgres image
docker pull postgres:15:2-alpine

# Start postgres container
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=testpass -d postgres:15.2-alpine

# Run command in container
docker exec -it postgres15 psql -U root
```

Postgres commands

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

Mysql commands [Handy MySQL Commands](http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm)

```sh
# Connect to database
mysql -h hostname -u username -p
mysql -uroot -ptestpass

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

Show docker logs

```sh
# Postgres
docker logs postgres15

# Mysql
docker logs mysql8
```

</details>
