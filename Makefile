JOBDATE		?= $(shell date -u +%Y-%m-%dT%H%M%SZ)
GIT_REVISION	= $(shell git rev-parse --short HEAD)
VERSION		?= $(shell git describe --tags --abbrev=0)

LDFLAGS		+= -X github.com/enablecloud/kulbe/version.Version=$(VERSION)
LDFLAGS		+= -X github.com/enablecloud/kulbe/version.Revision=$(GIT_REVISION)
LDFLAGS		+= -X github.com/enablecloud/kulbe/version.BuildDate=$(JOBDATE)

.PHONY: release

test:
	go test -v `go list ./... | egrep -v /vendor/`

build:
	@echo "++ Building keel"
	@echo "$(GOPATH)"
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "$(LDFLAGS)" -o kulbe .

image:
	docker build -t enablecloud/kulbe:alpha -f Dockerfile .

alpha: image
	@echo "++ Pushing keel alpha"	
	docker push enablecloud/kulbe:alpha