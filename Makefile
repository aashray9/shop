# https://unix.stackexchange.com/a/470502
ifndef os
override os = linux
endif

# https://unix.stackexchange.com/a/470502
ifndef arch
override arch = amd64
endif

build:
	-docker rm lms
	-docker rmi -f lms:latest
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) docker compose -f ./docker-compose.yml build --no-cache lms

run:
	# https://stackoverflow.com/a/2670143/6670698
	-docker rm lms_dev
	-docker rmi -f lms_dev:latest
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) docker compose -f ./docker-compose.yml up --remove-orphans lms-dev