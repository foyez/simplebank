# Backend Master

## Project Name: Simple Bank

- Create and manage account (owner, balance, currency)
- Record all balance changes (create an account entry for each change)
- Money transfer transaction (perform money transfer between 2 accounts consistently within a transaction)

## Database Design using <dbdiagram.io>

Design database tables using <dbdiagram.io>

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
