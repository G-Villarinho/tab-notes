.PHONY: run
run:
	@echo "Running the program..."
	@go run cmd/api/main.go

.PHONY: docker-up
docker-up: 
	@docker-compose up -d

.PHONY: docker-down
docker-down: 
	@docker-compose down

.PHONY: worker
worker:
	@echo "🔥 Starting email worker..."
	@go run cmd/cli/send_email/main.go