# Variáveis
BINARY_NAME=gocrud-auth-api
MAIN_FILES=main.go init_dependencies.go
DOCKER_IMAGE=gocrud-auth-api

# Comandos principais
.PHONY: help
help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@echo "  install              - Instala as dependências do projeto"
	@echo "  build                - Compila o projeto"
	@echo "  run                  - Roda a aplicação localmente"
	@echo "  run-build            - Compila e roda a aplicação"
	@echo "  test                 - Executa os testes"
	@echo "  test-verbose         - Executa os testes com output detalhado"
	@echo "  test-coverage        - Executa os testes com cobertura"
	@echo "  clean                - Remove arquivos de build"
	@echo "  docker-build         - Build da imagem Docker"
	@echo "  docker-run           - Roda com Docker Compose"
	@echo "  docker-run-detached  - Roda com Docker Compose em background"
	@echo "  docker-stop          - Para os containers Docker"
	@echo "  docker-clean         - Para e remove volumes dos containers"
	@echo "  docker-logs          - Mostra logs dos containers"
	@echo "  dev                  - Roda em modo desenvolvimento"
	@echo "  lint                 - Executa linter"
	@echo "  fmt                  - Formata o código"
	@echo "  vet                  - Executa go vet"
	@echo "  release              - Build para produção (Linux)"
	@echo "  release-windows      - Build para Windows"
	@echo "  release-mac          - Build para macOS"
	@echo "  mongo-start          - Inicia MongoDB localmente"
	@echo "  mongo-stop           - Para MongoDB local"
	@echo "  logs                 - Mostra logs da aplicação"
	@echo "  status               - Mostra status dos containers"

.PHONY: install
install: ## Instala as dependências do projeto
	go mod download
	go mod tidy

.PHONY: build
build: ## Compila o projeto
	go build -o $(BINARY_NAME) $(MAIN_FILES)

.PHONY: run
run: ## Roda a aplicação localmente
	go run $(MAIN_FILES)

.PHONY: run-build
run-build: build ## Compila e roda a aplicação
	./$(BINARY_NAME)

.PHONY: test
test: ## Executa os testes
	go test ./...

.PHONY: test-verbose
test-verbose: ## Executa os testes com output detalhado
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Executa os testes com cobertura
	go test -cover ./...

.PHONY: clean
clean: ## Remove arquivos de build
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe
	go clean

# Comandos Docker
.PHONY: docker-build
docker-build: ## Build da imagem Docker
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run: ## Roda com Docker Compose
	docker-compose up --build

.PHONY: docker-run-detached
docker-run-detached: ## Roda com Docker Compose em background
	docker-compose up -d --build

.PHONY: docker-stop
docker-stop: ## Para os containers Docker
	docker-compose down

.PHONY: docker-clean
docker-clean: ## Para e remove volumes dos containers
	docker-compose down -v

.PHONY: docker-logs
docker-logs: ## Mostra logs dos containers
	docker-compose logs -f

# Comandos de desenvolvimento
.PHONY: dev
dev: ## Roda em modo desenvolvimento (com hot reload se disponível)
	go run $(MAIN_FILES)

.PHONY: lint
lint: ## Executa linter (se disponível)
	@echo "Linting..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint não encontrado. Instale com: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

.PHONY: fmt
fmt: ## Formata o código
	go fmt ./...

.PHONY: vet
vet: ## Executa go vet
	go vet ./...

# Comandos de deploy
.PHONY: release
release: ## Build para produção
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $(BINARY_NAME) $(MAIN_FILES)

.PHONY: release-windows
release-windows: ## Build para Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o $(BINARY_NAME).exe $(MAIN_FILES)

.PHONY: release-mac
release-mac: ## Build para macOS
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o $(BINARY_NAME) $(MAIN_FILES)

# Comandos de banco de dados
.PHONY: mongo-start
mongo-start: ## Inicia MongoDB localmente com Docker
	docker run -d -p 27017:27017 --name mongo mongo:latest

.PHONY: mongo-stop
mongo-stop: ## Para MongoDB local
	docker stop mongo && docker rm mongo

# Comandos de monitoramento
.PHONY: logs
logs: ## Mostra logs da aplicação (se rodando com Docker)
	docker-compose logs -f app

.PHONY: status
status: ## Mostra status dos containers
	docker-compose ps 