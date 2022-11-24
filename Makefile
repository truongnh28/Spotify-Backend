COVERAGE := TRUE
DISPLAY_COVERAGE := TRUE
HTML_FILE:= coverage.html
JUNIT_FILE:= rspec.xml
COBERTURA_FILE:= coverage.xml
REPORT_PATH:= ./test-reports

gen:
	## Go generate
	@go generate ./...

test:
	@echo "==> Running unit tests with coverage <=="
	@bash ./scripts/test.sh --coverage $(COVERAGE) --display $(DISPLAY_COVERAGE) --html-report "$(HTML_FILE)" --junit-report "$(JUNIT_FILE)" --cobertura-report "$(COBERTURA_FILE)" --report-path "$(REPORT_PATH)"

ut:
	@go test ./... -failfast

coverage: test
	@go tool cover -html=test-results/.testcoverage.txt -o test-results/coverage.html && open test-results/coverage.html

run:
	cd cmd && go run main.go

##@ Development

PROJECT_NAME ?= spotify
DOCKER_COMPOSE ?= cd dev && docker compose -p $(PROJECT_NAME)

dev-up: ## Run local environment for developing locally
	$(DOCKER_COMPOSE) up -d

dev-down: ## Shutdown the local environment
	$(DOCKER_COMPOSE) down
