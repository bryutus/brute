up:
	docker-compose up -d

stop:
	docker-compose stop

go-app:
	docker-compose exec app /bin/bash

go-db:
	docker-compose exec db /bin/bash

deploy:
	git push heroku main
