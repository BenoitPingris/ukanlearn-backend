help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[ 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

db-dev: ## Launch a psql db in a docker container for development
	docker run --rm -e POSTGRES_DB=ukanlearn -e POSTGRES_USER=ukanlearn -e POSTGRES_PASSWORD=docker -p 5432:5432 --name psql-uk-dev postgres

db-test: ## Launch a psql db in a docker container for testing
	docker run --rm -e POSTGRES_DB=db_test -e POSTGRES_USER=gorm -e POSTGRES_PASSWORD=gorm -p 5432:5432 --name psql-uk-test postgres
	
.PHONY: db-test db-dev