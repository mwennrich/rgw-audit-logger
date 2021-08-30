GO111MODULE := on
DOCKER_TAG := $(or ${GITHUB_TAG_NAME}, latest)

all: rgw-audit-logger

.PHONY: rgw-audit-logger
rgw-audit-logger:
	go build -o bin/rgw-audit-logger
	strip bin/rgw-audit-logger

.PHONY: dockerimages
dockerimages:
	docker build -t mwennrich/rgw-audit-logger:${DOCKER_TAG} .

.PHONY: dockerpush
dockerpush:
	docker push mwennrich/rgw-audit-logger:${DOCKER_TAG}

.PHONY: clean
clean:
	rm -f bin/*
