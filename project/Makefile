BROKER_BINARY=brokerApp
POST_BINARY=postApp
LOGGER_BINARY=loggerApp
MAIL_BINARY=mailApp
AUTH_BINARY=authApp
IMAGE_BINARY=imageApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_post build_mail build_logger build_auth build_image
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../back-end/broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
install_front:
	@echo "Building front end binary..."
	cd ../front-end && npm install
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ../back-end/logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

## build_post: builds the authentication binary as a linux executable
build_auth:
	@echo "Building authentication binary..."
	cd ../back-end/authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_post: builds the image binary as a linux executable
build_image:
	@echo "Building image binary..."
	cd ../back-end/image-service && env GOOS=linux CGO_ENABLED=0 go build -o ${IMAGE_BINARY} ./cmd/api
	@echo "Done!"

## build_post: builds the post binary as a linux executable
build_post:
	@echo "Building post binary..."
	cd ../back-end/post-service && env GOOS=linux CGO_ENABLED=0 go build -o ${POST_BINARY} ./cmd/api
	@echo "Done!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ../back-end/mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

## start: starts the front end, install_front if it was never realized
start:
	@echo "Starting front end"
	cd ../front-end && npm run dev &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	killall node
	@echo "Stopped front end!"