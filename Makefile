include .env

.PHONY: proto proto-lint

build:
	docker-compose -p seeder-app build

up:
	docker-compose -p seeder-app up -d

tidy:
	docker-compose exec seeder-app-app go mod tidy

api-goget:
	docker-compose exec seeder-app go get ${MOD}

api-lint:
	docker-compose exec seeder-app go fmt ./...
	docker-compose exec seeder-app sh -c 'staticcheck -go 1.0 $$(go list ./... | grep -v "seeder_api_service/gen")'

api-test:
	docker-compose exec seeder-app go test -cover ./... -coverprofile=../cover.out

destroy:
	docker-compose down --rmi all --volumes --remove-orphans
