
#FOR LOCAL START
migrate:
	migrate create -ext sql -dir ./migrations -seq init
migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' down
create-container:
	docker run --name=factory-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres

#FOR DOCKER START
docker-up:
	docker-compose up -d