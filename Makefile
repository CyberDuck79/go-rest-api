NAME := cat-api

all: gen.entity gen.server build

gen.run: gen run

gen: gen.entity gen.server

gen.entity:
	go generate ./ent

gen.server:
	swagger generate server -t swagger -f ./swagger/swagger.yml --exclude-main -A $(NAME)

run:
	go run ./swagger/cmd/$(NAME)-server/main.go --port 3000

build:
	go install ./swagger/cmd/$(NAME)-server/