.DEFAULT_GOAL := help
help:
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'
build:  ## build and up console container
		docker-compose up -d  --build
up:     ## up console container
		docker-compose up -d
start:  ## start console container
		docker-compose start
down:   ## down console container
		docker-compose down
destroy: ## destroy console container
		docker-compose down -v
stop:   ## stop  console container
		docker-compose stop
restart: ## restart console container
		docker-compose stop
		docker-compose up -d
ps: 	## list containers
		docker-compose ps
console: ## goto console container bash
		docker-compose  exec console bash
adder-bin: ## run binary adder with args. For example: make adder-bin args='12312341234123412341234.1234123412341234123 9872394582347958237459882345324.29824234234234254'
		docker-compose exec console adder $(args)
adder-src:  ## run adder from source with args
		docker-compose exec console adder $(args)
