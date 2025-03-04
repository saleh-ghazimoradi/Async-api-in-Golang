db_login:
	psql ${DATABASE_URL}

db_create_migration:
	migrate create -ext sql -dir migrations -seq $(name)

db_migrate_up:
	migrate -database ${DATABASE_URL} -path migrations up

db_migrate_down:
	migrate -database ${DATABASE_URL} -path migrations down 1

db_migrate_drop:
	migrate -database ${DATABASE_URL} -path migrations drop


docker_up:
	docker compose up -d

docker_down:
	docker compose down

