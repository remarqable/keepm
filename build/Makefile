UNAME_S := $(shell uname -s)

all:
ifeq ($(UNAME_S),Darwin)
    $(error Build only available for Linux)
endif
    # rest of the build steps go here

DOCKER_PS := $(shell docker ps -a -q)
DOCKER_IMAGES := $(shell docker images -a -q)
DOCKER_LOGS := $(shell docker ps|grep keepm|cut -d " " -f1)

.DEFAULT_GOAL:=help


help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "build	build the project"
	@echo "docker	build docker image"
	@echo "run	run docker container"
	@echo "logs	show the logs of the container"
	@echo "clean	stop/remove all containers and images"

build:
	@echo "Building for Linux"
	go build -ldflags="-extldflags=-static" -tags netgo -o keepm ../main.go


docker:
	@echo "Building docker image for keepm"
	cp ../db/dbinit.sql .
	@docker build --progress=plain -t keepm  -f keepm.dockerfile .	
	rm dbinit.sql

run:
ifndef DOCKER_PS
	@echo "Starting keepm container"
	@docker run --name keepm -h keepm-docker -d -p 5432:5432 -p 8080:8080  -t keepm
else
	@echo  "Container keepm already running"
endif

logs:
ifdef DOCKER_LOGS
	@docker logs $(DOCKER_LOGS)
endif

clean:
ifdef DOCKER_PS
	@docker rm -f $(DOCKER_PS) 
endif
ifdef DOCKER_IMAGES
	@docker rmi -f $(DOCKER_IMAGES)
endif
	@echo -n ""
