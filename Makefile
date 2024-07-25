.PHONY: build build-api build-consumer run run-api run-consumer

run:
	docker compose up

build: build-api build-consumer

build-api:
	@echo "Building skill-api"
	@cd api && make build

build-consumer:
	@echo "Building skill-consumer"
	@cd consumer && make build

run-api:
	docker compose up skill-api

run-consumer:
	docker compose up skill-consumer

test:
	@echo "Testing skill-api"
	@cd api && make test
	@echo "-----------------------------------------------------------------------"
	@echo "Testing skill-consumer"
	@cd consumer && make test