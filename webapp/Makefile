ROOT := $(pwd)
DOCKER_REPOSITORY := basebandit-webapp
DOCKER_IMAGE_NAME := basebandit-webapp:latest
DOCKER_IMAGE_TAG := basebandit-webapp:v1.0

build:
	docker build --tag ${DOCKER_IMAGE_NAME} -t ${DOCKER_IMAGE_TAG} ${PWD}
run:
	docker run -p 8080:8080 ${DOCKER_REPOSITORY}