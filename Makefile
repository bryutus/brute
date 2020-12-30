up:
	docker-compose up -d

stop:
	docker-compose stop

go-app:
	docker-compose exec app /bin/bash

go-db:
	docker-compose exec db /bin/bash

go-psql:
	docker exec -it db psql -U admin -h localhost -d app

migrate-status:
	docker exec -it app sql-migrate status

migrate-up:
	docker exec -it app sql-migrate up

migrate-down:
	docker exec -it app sql-migrate down

deploy:
	git push heroku main
