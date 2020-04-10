setup: up

build-up: 
	docker-compose build --no-cache && docker-compose up -d

up:
	docker-compose up -d

down:
	docker-compose down

inspect-redis:
	docker exec -it redis redis-cli

inspect-shop:
	docker exec -it shop /bin/sh

inspect-dashboard:
	docker exec -it dashboard /bin/sh

inspect-customer:
	docker exec -it customer /bin/sh