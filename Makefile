.PHONY: build
build:
	go build -o ./bin/app ./cmd/app

.PHONY: run
run: build
	go run ./cmd/app

bin/migrate:
	go build -tags 'sqlite3' -o ./bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate

.PHONY: migrate
migrate: bin/migrate
	./bin/migrate -source file://migrations -database sqlite3://./tasks.db up

.PHONY: rollback
rollback: bin/migrate
	./bin/migrate -source file://migrations -database sqlite3://./tasks.db down

.PHONY: create-task
create-task:
	curl -XPOST http://localhost:8080/tasks/ -H 'Content-Type: application/json' -d '{"title": "test", "description": "test"}'

.PHONY: create-migration
create-migration: bin/migrate
	./bin/migrate create -ext sql -dir ./migrations -seq $(name)