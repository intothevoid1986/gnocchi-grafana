BINARY_NAME=gnocchi-token-renewal
BUILD_PATH=build
CONTAINER_REGISTRY_PATH=intothevoid
TAG=latest

USERNAME=intothevoid
PASSWORD=dckr_pat_Pv8n0nDgQ7qcGu0udhur1lqFR2o

build:
	CGO_ENABLED=0 go build -o ${BUILD_PATH}/${BINARY_NAME} -v main.go
	chmod +x ${BUILD_PATH}/${BINARY_NAME}

docker:
	docker build -t ${CONTAINER_REGISTRY_PATH}/${BINARY_NAME}:${TAG} .

docker-push:
	docker login -u ${USERNAME} -p ${PASSWORD} 
	docker push ${CONTAINER_REGISTRY_PATH}/${BINARY_NAME}:${TAG}

clean:
		if [ -f ${BUILD_PATH}/${BINARY_NAME} ] ; then rm -rf ${BUILD_PATH} ; fi

all: clean build docker docker-push