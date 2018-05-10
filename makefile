SRCPATH=$(GOPATH)/src/github.com/dwarvesf/devops-assignment
.PHONY: remove-infras init backend-up backend-build backend-package frontend-up

remove-infras:
	docker-compose stop; docker-compose rm -f

init:
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec accident_db mysqladmin --user=test --password=test --host "127.0.0.1" ping --silent &> /dev/null ; do \
		sleep 1; \
	done
	@echo "Seeding data..."
	@./setup.sh

backend-up:
	cd $(SRCPATH)/backend && make build && \
	docker rm -f accident_backend | true && \
	docker-compose up -d --build --force-recreate

backend-build:
	cd ./backend && make build

backend-package:
	cd ./backend && make package

frontend-up:
	cd ./frontend && make up