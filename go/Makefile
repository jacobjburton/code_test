PROJECT_NAME := bracketizer

.DEFAULT := help
.PHONY: local down reset app all-reset nginx nginx-reset

help:
	@echo '# Dev Targets'
	@echo
	@echo '    local            build the application containers/services locally'
	@echo '    down             tear down the application containers/services locally'
	@echo '    reset            rebuild the application container and restart the service'
	@echo
	@echo '    all-reset        reset the application containers/services'
	@echo '    app              build the application container'
	@echo '    nginx            build the nginx container'
	@echo '    nginx-reset      rebuild the nginx container and restart the service'
	@echo
	@echo '    help             show this help message'
	@echo

local: app nginx
	docker compose up -d --force-recreate \
	&& docker compose logs -f

down:
	docker compose down -v

all-reset: down local

reset: app
	docker compose up -d --force-recreate ${PROJECT_NAME}-app \
	&& docker compose logs -f

app:
	docker compose build ${PROJECT_NAME}-app

nginx:
	docker compose build ${PROJECT_NAME}-nginx

nginx-reset: nginx
	docker compose up -d --force-recreate ${PROJECT_NAME}-nginx
