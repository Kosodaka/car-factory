
#FOR LOCAL START
migrate:
	migrate create -ext sql -dir ./migrations -seq init
migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
create-container:
	docker run --name=factory-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres

#local run
run-http:
	go run cmd/app/http/main.go
run-tui:
	go run cmd/app/tui/main.go

#FOR DOCKER START
docker-up:
	docker-compose up -d
#enter in docker container to interact with tui
docker-in:
	docker-compose attach api
#FOR Tests
.PHONY:mock-gen
mock-gen:
	mockgen -source=app/repo/repo/repo.go -destination=pkg/mocks/api/repo/repo_mock.go
	mockgen -source=app/service/service.go -destination=pkg/mocks/api/service/service_mock.go

test:
	go test -cover ./app/service/...