.PHONY: *

backend:
	cd backend && go run cmd/server/main.go

frontend:
	cd frontend && bun dev
