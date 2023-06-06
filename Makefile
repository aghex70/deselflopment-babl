.PHONY: local
local: ## Run the application locally
	docker-compose up --build

automigrate: ## Run the automigration
	docker-compose run babl-backend sh -c "go build -o deselflopment-babl && ./deselflopment-babl automigrate"

stop: ## Stop the application
	docker-compose down