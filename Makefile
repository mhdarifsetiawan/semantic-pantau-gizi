# ============================================================
# Makefile — semantic-pantau-gizi
# Usage: make <target>
# ============================================================

APP_NAME    := semantic-pantau-gizi
COMPOSE     := docker compose -f deployments/docker-compose.yml
COMPOSE_DEV := docker compose -f deployments/docker-compose.yml -f deployments/docker-compose.dev.yml

# ============================================================
# HELP
# ============================================================

.PHONY: help
help:
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@echo "  CHECK"
	@echo "    check            Check semua tools yang dibutuhkan (docker, go, air)"
	@echo ""
	@echo "  DEVELOPMENT"
	@echo "    dev              Jalankan stack development dengan Air live-reload"
	@echo "    dev-build        Build ulang image dev lalu jalankan"
	@echo "    dev-down         Stop stack development"
	@echo "    dev-logs         Tail logs development"
	@echo ""
	@echo "  PRODUCTION"
	@echo "    up               Jalankan stack production"
	@echo "    up-build         Build ulang image lalu jalankan production"
	@echo "    down             Stop stack production"
	@echo "    logs             Tail logs production"
	@echo ""
	@echo "  VALIDATE"
	@echo "    validate         Validasi semua docker-compose file"
	@echo "    lint             Jalankan golangci-lint"
	@echo "    test             Jalankan unit test"
	@echo ""
	@echo "  DATABASE"
	@echo "    db-only          Jalankan hanya postgres dan redis"
	@echo "    db-shell         Masuk ke psql shell"
	@echo ""
	@echo "  CLEAN"
	@echo "    clean            Hapus container, volume, dan image"
	@echo "    prune            Docker system prune"
	@echo ""

# ============================================================
# CHECK — pastikan semua tools terinstall
# ============================================================

.PHONY: check
check:
	@echo "Checking required tools..."
	@echo ""

	@# Check Docker
	@if ! command -v docker > /dev/null 2>&1; then \
		echo "[MISSING] docker — install: https://docs.docker.com/get-docker/"; \
	else \
		echo "[OK]      docker $$(docker --version)"; \
	fi

	@# Check Docker daemon running
	@if ! docker info > /dev/null 2>&1; then \
		echo "[ERROR]   docker daemon is not running — start Docker Desktop or dockerd"; \
	else \
		echo "[OK]      docker daemon is running"; \
	fi

	@# Check docker compose
	@if ! docker compose version > /dev/null 2>&1; then \
		echo "[MISSING] docker compose plugin — update Docker Desktop or install manually"; \
	else \
		echo "[OK]      $$(docker compose version)"; \
	fi

	@# Check Go
	@if ! command -v go > /dev/null 2>&1; then \
		echo "[MISSING] go — install: https://go.dev/dl/"; \
	else \
		echo "[OK]      go $$(go version | awk '{print $$3}')"; \
	fi

	@# Check Air
	@if ! command -v air > /dev/null 2>&1; then \
		echo "[MISSING] air — install: go install github.com/air-verse/air@latest"; \
	else \
		echo "[OK]      air $$(air -v 2>&1 | head -1)"; \
	fi

	@# Check .env file
	@if [ ! -f .env ]; then \
		echo "[MISSING] .env — run: cp .env.example .env"; \
	else \
		echo "[OK]      .env file exists"; \
	fi

	@echo ""

# ============================================================
# VALIDATE — cek docker-compose config tanpa menjalankan apapun
# ============================================================

.PHONY: validate
validate:
	@echo "Validating docker-compose.yml..."
	$(COMPOSE) config --quiet && echo "[OK] docker-compose.yml is valid"
	@echo ""
	@echo "Validating docker-compose.dev.yml..."
	$(COMPOSE_DEV) config --quiet && echo "[OK] docker-compose.dev.yml is valid"
	@echo ""

# ============================================================
# DEVELOPMENT
# ============================================================

.PHONY: dev
dev: check-env
	$(COMPOSE_DEV) up

.PHONY: dev-build
dev-build: check-env
	$(COMPOSE_DEV) up --build

.PHONY: dev-down
dev-down:
	$(COMPOSE_DEV) down

.PHONY: dev-logs
dev-logs:
	$(COMPOSE_DEV) logs -f api

# ============================================================
# PRODUCTION
# ============================================================

.PHONY: up
up: check-env
	$(COMPOSE) up -d

.PHONY: up-build
up-build: check-env
	$(COMPOSE) up -d --build

.PHONY: down
down:
	$(COMPOSE) down

.PHONY: logs
logs:
	$(COMPOSE) logs -f api

# ============================================================
# DATABASE
# ============================================================

.PHONY: db-only
db-only: check-env
	$(COMPOSE) up -d postgres redis
	@echo ""
	@echo "postgres and redis are running."
	@echo "run 'air' to start the API with live-reload."
	@echo ""

.PHONY: db-shell
db-shell:
	$(COMPOSE) exec postgres psql -U $${DB_USER:-postgres} -d $${DB_NAME:-pantau_gizi}

# ============================================================
# CODE QUALITY
# ============================================================

.PHONY: lint
lint:
	@if ! command -v golangci-lint > /dev/null 2>&1; then \
		echo "golangci-lint not found — install: https://golangci-lint.run/usage/install/"; \
		exit 1; \
	fi
	golangci-lint run ./...

.PHONY: test
test:
	go test ./... -race -count=1

.PHONY: test-cover
test-cover:
	go test ./... -race -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "coverage report: coverage.html"

# ============================================================
# CLEAN
# ============================================================

.PHONY: clean
clean:
	$(COMPOSE) down -v --rmi local
	$(COMPOSE_DEV) down -v --rmi local 2>/dev/null || true

.PHONY: prune
prune:
	docker system prune -f

# ============================================================
# INTERNAL HELPERS
# ============================================================

.PHONY: check-env
check-env:
	@if [ ! -f .env ]; then \
		echo ""; \
		echo "[ERROR] .env file not found."; \
		echo "        run: cp .env.example .env"; \
		echo ""; \
		exit 1; \
	fi