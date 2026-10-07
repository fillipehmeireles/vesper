.PHONY: lab-up lab-down lab-reset lab-clean

run:
	go run ./cmd/main.go --name node01

lab-up:
	docker compose up --build

lab-down:
	docker compose down --remove-orphans

lab-reset:
	docker compose down --remove-orphans --volumes
	docker compose up --build

lab-clean:
	docker compose down --remove-orphans --volumes --rmi local
