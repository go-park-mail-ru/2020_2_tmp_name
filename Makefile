MAIN_BINARY=main_service
AUTH_BINARY=auth_service
COMMENT_BINARY=comment_service

PROJECT_DIR := ${CURDIR}

DOCKER_DIR := ${CURDIR}/docker

## build: Build compiles project
build:
	go build -o ${MAIN_BINARY} cmd/main/main.go
	go build -o ${AUTH_BINARY} cmd/auth/main_auth.go
	go build -o ${COMMENT_BINARY} cmd/comment/main_comment.go
	
## build-docker: Builds all docker containers
build-docker:
	docker build -t dependencies -f ${DOCKER_DIR}/builder.Dockerfile .
	docker build -t main_service -f ${DOCKER_DIR}/main.Dockerfile .
	docker build -t auth_service -f ${DOCKER_DIR}/auth.Dockerfile .
	docker build -t comment_service -f ${DOCKER_DIR}/comment.Dockerfile .
	docker build -t face_service -f ${DOCKER_DIR}/face.Dockerfile .

## run-and-build: Build and run docker
build-and-run: build-docker
	docker-compose up

## run: Build and run docker with new changes
run:
	docker rm -vf $$(docker ps -a -q) || true
	sudo rm -rf postgres_data/
	docker build -t dependencies -f ${DOCKER_DIR}/builder.Dockerfile .
	docker build -t main_service -f ${DOCKER_DIR}/main.Dockerfile .
	docker build -t auth_service -f ${DOCKER_DIR}/auth.Dockerfile .
	docker build -t comment_service -f ${DOCKER_DIR}/comment.Dockerfile .
	docker build -t face_service -f ${DOCKER_DIR}/face.Dockerfile .
	
	docker-compose up --build --no-deps

## coverage-html: generates HTML file with test coverage
tests:
	sudo rm -rf postgres_data/
	go test ./... -coverprofile cover; go tool cover -func cover
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