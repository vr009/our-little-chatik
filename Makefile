.SILENT:
.PHONY: run migrate-Users migrate-Users-drop

run:
	docker-compose up --remove-orphans --build

migrate-users:
	migrate -path ./internal/auth/schema -database 'postgres://postgres:admin@0.0.0.0:5433/users?sslmode=disable' up

migrate-users-drop:
	migrate -path ./users-service/schema -database 'postgres://postgres:adminy@0.0.0.0:5433/users?sslmode=disable' drop