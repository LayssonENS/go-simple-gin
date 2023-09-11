swagger:
	swag init -g app/main.go
	swag init -g books/delivery/http/books_handler.go

mock:
	mockgen -destination=domain/mock/domain_mock.go -source=domain/books.go

test:
	go test ./...

build:
	docker-compose build
	docker-compose up -d

run:
	go run app/main.go
