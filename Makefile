## db_docs: generate dbdocs view
db_docs:
	@echo "generating dbdocs view..."
	dbdocs build doc/db.dbml

## db_schema: convert a dbml file to sql
db_schema:
	@echo "converting a dbml file to sql..."
	dbml2sql doc/db.dbml --postgres -o doc/schema.sql

.PHONY: db_docs db_schema