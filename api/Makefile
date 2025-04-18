.PHONY: run
run:
	@go run main.go

.PHONY: setup
setup:
	@echo "📦 Instalando ferramentas e preparando ambiente..."

	@if ! command -v mockery &> /dev/null; then \
		echo "🔧 Instalando mockery..."; \
		go install github.com/vektra/mockery/v2@v2.53.3; \
	else \
		echo "✅ mockery já está instalado."; \
	fi

.PHONY: generate-key
generate-key: 
	@echo "🔑 Gerando chaves ECDSA..."
	@openssl ecparam -name prime256v1 -genkey -noout -out ecdsa_private.pem
	@openssl ec -in ecdsa_private.pem -pubout -out ecdsa_public.pem

.PHONY: docker-up
docker-up: 
	@echo "🚀 Iniciando containers..."
	@sudo docker-compose up -d

.PHONY: docker-down
docker-down: 
	@echo "🛑 Parando containers..."
	@docker-compose down

.PHONY: mock
mock: 
	@echo "🔄 Gerando mocks..."
	@mockery

kill-port:
	@echo "🔫 Matando processo na porta 8080..."
	@fuser -k 8080/tcp || true