
COMMONENVVAR		= GOOS=linux GOARCH=amd64
BUILDENVVAR			= CGO_ENABLED=0
GOTESTNOCACHE		?= GOCACHE=off
BIN_NAME			:= TelRobot
JSONITERTAG			:= -tags=jsoniter
GITTAG 				?= $(shell git describe --abbrev=0 --tags 2>/dev/null || echo NO_TAG)
GITTAG_NO_V			:= $(subst v,,$(GITTAG))
GITCOMMITHASH		?= $(shell git rev-parse --short HEAD)
DATE_TIME			:= $(shell date '+%Y-%m-%dT%H:%M:%S')
COMMIT_COUNT		:= $(shell git rev-list --all --count)
DOCKER_PROJECT		:= telrobot
DOCKER_VERSION		:= $(GITTAG_NO_V).$(GITCOMMITHASH)
VERSIONTAG			:= -ldflags "-X main.BuildTime=$(DATE_TIME) -X main.BuildGitHash=$(GITCOMMITHASH) -X main.BuildGitTag=$(GITTAG) -X main.BuildGitCount=$(COMMIT_COUNT)"
PACKAGES			:= `go list ./... | grep -v /vendor/`
VETPACKAGES			:= `go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES				:= `find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./carloc/tracking_pipeline/*"`


.PHONY: all

all: build

deps:
# 	go mod tidy

#build: deps
#	go build -mod=vendor $(JSONITERTAG) -o bin/$(BIN_NAME) $(VERSIONTAG) .

build: deps
	go mod download github.com/go-sql-driver/mysql
	go mod download github.com/lib/pq
	go build $(JSONITERTAG) -o bin/$(BIN_NAME) $(VERSIONTAG) .


vendor-build:
	go build $(JSONITERTAG) -o bin/$(BIN_NAME) $(VERSIONTAG) .

go-mod-vendor:
	go mod vendor

linux_build:
	$(COMMONENVVAR) $(BUILDENVVAR) go build $(JSONITERTAG) -o bin/$(BIN_NAME)-linux $(VERSIONTAG) .


docker:
	docker build --no-cache -t $(DOCKER_PROJECT):$(DOCKER_VERSION) -f Dockerfile .
#	docker tag $(DOCKER_PROJECT):$(DOCKER_VERSION) $(DOCKER_HUB_REGISTRY)/$(DOCKER_HUB_REGISTRY_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)
#	docker tag $(DOCKER_PROJECT):$(DOCKER_VERSION) $(DOCKER_HUB)/$(DOCKER_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)

quick-docker:go-mod-vendor
	docker build  -t $(DOCKER_PROJECT):$(DOCKER_VERSION) -f Dockerfile-quick .
	docker tag $(DOCKER_PROJECT):$(DOCKER_VERSION) $(DOCKER_HUB)/$(DOCKER_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)

docker-runner:
	docker build --no-cache --rm -t alpine_plus:latest - < Dockerfile_runner

docker-push-registry:
	docker push $(DOCKER_HUB_REGISTRY)/$(DOCKER_HUB_REGISTRY_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)

docker-push:
	docker push $(DOCKER_HUB)/$(DOCKER_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)


clean:
	@rm -rf bin
