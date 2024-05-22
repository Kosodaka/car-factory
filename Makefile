
#FOR LOCAL START
migrate:
	migrate create -ext sql -dir ./migrations -seq init
migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
create-container:
	docker run --name=factory-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres

#FOR DOCKER START
docker-up:
	docker-compose up -d
#enter in docker container to interact with tui
docker-in:
	docker-compose attach api
