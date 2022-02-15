GOOSE_DBSTRING = "user=admin dbname=with-tweet sslmode=disable password=admin"

migrate-%:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ./pkg/migration ${@:migrate-%=%}

postgress-rebuild:
	docker-compose down -v
	docker-compose build --no-cache
	docker-compose up -d

generate:
	@if [ ! -d "./vendor" ]; then\
  	go mod vendor;\
	fi
	go run ./gqlgen.go

run-server:
	go run ./cmd/tweet_dstribution/main.go

run-batch:
	go run ./cmd/tweet_submit/main.go