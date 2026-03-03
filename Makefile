include .env

.PHONY:	run tidy build

run:
	go run cmd/main.go

build: bin/main
bin/main: cmd/main.go
	@echo "Building binary..."
	go build -o bin/main cmd/main.go

tidy:
	go mod tidy

# make migrate-new-postgres name=create_users_table
migrate-new-postgres:
	migrate create -ext sql -dir $(POSTGRES_MIGRATE_PATH) $(name)

migrate-up-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose up

migrate-down-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose down 

migrate-force-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose force $(DIRTY)

migrate-drop:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose drop

migrate-fresh-postgres: migrate-down-postgres migrate-up-postgres

sqlc:
	sqlc -f sqlc.yaml generate


docker-build: sqlc 
	docker compose build

docker-up: sqlc
	docker compose up --build

# Menghentikan semua service tanpa menghapus data di database
docker-down:
	docker compose down

# Menghentikan service DAN menghapus volume (data database akan hilang/reset)
# Gunakan ini kalau kamu mau mulai database dari kondisi benar-benar kosong
docker-clean:
	docker compose down -v
