SHELL := /bin/bash
.PHONY: build
build:
	@echo 'Building Docker image'
	docker build -t arena-add .

run:
	@echo 'Running application'
	docker-compose run  arena-addition --firstNumber=$(firstnumber) --secondNumber=$(secondnumber)
