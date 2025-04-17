# Project settings
PROJECT_NAME = trader.pro-be
COMPOSE = docker-compose
DOCKER_RUN = docker compose run --rm

# Services
BACKEND_SERVICE = backend
BREEZESVC_SERVICE = breezesvc
NGINX_SERVICE = nginx

# Targets

.PHONY: all build up down restart logs backend breeze test deploy

## Default target (build, up)
all: build up

## Build all services
build:
	$(COMPOSE) build

## Start all services
up:
	$(COMPOSE) up -d

## Stop all services
down:
	$(COMPOSE) down

## Restart all services
restart: down up

## Show logs for all services
logs:
	$(COMPOSE) logs -f

## Run Go backend locally inside container
backend:
	$(DOCKER_RUN) $(BACKEND_SERVICE) go run main.go

## Run Flask service locally inside container
breeze:
	$(DOCKER_RUN) $(BREEZESVC_SERVICE) python app.py
