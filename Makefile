PHONY: migrate expose-pg-pass

include .env

migrate:
	@psql "$(POSTGRESQL_URL)" -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	@export PG_PASSWORD="$(PG_PASSWORD)"
	tern migrate --migrations sql/migrations/
