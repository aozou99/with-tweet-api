GOOSE_DBSTRING = "user=admin dbname=with-tweet sslmode=disable password=admin"

migration-%:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ./migration ${@:migration-%=%}

postgress-rebuild:
	docker-compose down -v
	docker-compose build --no-cache
	docker-compose up -d