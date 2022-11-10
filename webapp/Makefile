DOCKER_CONTAINER_NAME := webapp
DOCKER_IMAGE_NAME := basebandit-webapp
TAG_NAME := mrparish/${DOCKER_IMAGE_NAME}

build:
	@docker build -t ${DOCKER_IMAGE_NAME} .; docker tag ${DOCKER_IMAGE_NAME} ${TAG_NAME}
run:
	@docker run --name ${DOCKER_CONTAINER_NAME} -p 8080:8080 -d ${TAG_NAME}
stop:
	@docker stop ${DOCKER_CONTAINER_NAME}; docker rm ${DOCKER_CONTAINER_NAME}

