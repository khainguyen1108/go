# name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

# migrations
MIGRATE=migrate
DB_URL=mysql://baonk:baonk123@tcp(localhost:33306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local
MIGRATION_PATH=migrations

migrate-create:
	$(MIGRATE) create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migrate-up:
	$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" down 1