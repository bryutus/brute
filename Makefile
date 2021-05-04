up:
	docker compose up -d

stop:
	docker compose stop

go-app:
	docker compose exec app /bin/bash

app-log:
	docker logs -f app

test:
	docker exec -it app go test -v ./... -count=1

test-coverage:
	docker exec -it app go test -race -timeout 30m -coverprofile=coverage.txt -covermode=atomic ./...
	docker exec -it app go tool cover -html=coverage.txt -o coverage.html

go-db:
	docker compose exec db /bin/bash

go-psql:
	docker exec -it db psql -U admin -h localhost -d app

migrate-status:
	docker exec -it app goose status

migrate-up:
	docker exec -it app goose up

migrate-down:
	docker exec -it app goose down

deploy:
	git push heroku main

go-prd-app:
	heroku run bash

go-prd-psql:
	heroku pg:psql
