MAIN_BINARY=main_service
AUTH_BINARY=auth_service
COMMENT_BINARY=comment_service

PROJECT_DIR := ${CURDIR}

DOCKER_DIR := ${CURDIR}/docker

## build: Build compiles project
build:
	go build -o ${AUTH_BINARY} cmd/auth/main_auth.go
	go build -o ${COMMENT_BINARY} cmd/comment/main_comment.go
	go build -o ${MAIN_BINARY} cmd/main/main.go

## build-docker: Builds all docker containers
build-docker:
	docker build -t dependencies -f ${DOCKER_DIR}/builder.Dockerfile .
	docker build -t main_service -f ${DOCKER_DIR}/main_service.Dockerfile .
	docker build -t auth_service -f ${DOCKER_DIR}/auth.Dockerfile .
	docker build -t comment_service -f ${DOCKER_DIR}/comment.Dockerfile .

## run-and-build: Build and run docker
build-and-run: build-docker
	docker-compose up

## run: Build and run docker with new changes
run:
	docker rm -vf $$(docker ps -a -q) || true
	docker build -t dependencies -f ${DOCKER_DIR}/builder.Dockerfile .
	docker build -t auth_service -f ${DOCKER_DIR}/auth.Dockerfile .
	docker build -t comment_service -f ${DOCKER_DIR}/comment.Dockerfile .
	docker build -t main_service -f ${DOCKER_DIR}/main.Dockerfile .
	docker-compose up --build --no-deps

## test-coverage: get final code coverage
coverage:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	rm -rf cover

## coverage-html: generates HTML file with test coverage
test-html:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	go tool cover -html=cover
	rm -rf cover

## run-background: run process in background(available after build)
run-background:
	docker-compose up -d

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | col