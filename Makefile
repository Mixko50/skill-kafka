run:
	docker compose up

build api:
	echo "Building skill-api"
	cd api && make build

build consumer:
	echo "Building skill-consumer"
	cd consumer && make build

run api:
	docker compose up skill-api

run consumer:
	docker compose up skill-consumer