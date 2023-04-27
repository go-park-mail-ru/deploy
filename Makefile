.PHONY: build
build:
	go build -o ./bin/app ./app


#.PHONY: migrate
#migrate:
#	migrate -source file://path/to/migrations -database postgres://postgres:mysecretpassword@localhost:5432/testdb up
