postgres:
	docker run --name es-postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=addison -d postgres:12-alpine

createdb:
	docker exec -it es-postgres-12 createdb --username=root --owner=root elo-system

dropdb:
	docker exec -it es-postgres-12 dropdb elo-system

migrateup:
	migrate -path database/migrate -database "postgresql://root:addison@172.18.48.243:5432/elo-system?sslmode=disable" -verbose up 

migratedown:
	migrate -path database/migrate -database "postgresql://root:addison@172.18.48.243:5432/elo-system?sslmode=disable" -verbose down
	
.PHONY: postgres createdb dropdb migrateup migratedown