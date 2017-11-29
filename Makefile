REPO = packer-post-processor-openstack-image-management
BIN = $(REPO)
VERSION = 0.0.1
ARTIFACTS_DIR = dist/snapshot/

all: clean test build

test: dep
	go test ./...

build: dep
	go build -o bin/$(BIN) main.go

cross: dep
	goxc

deploy: cross
	ghr -u $(GITHUB_USER) -t $(GITHUB_TOKEN) -r $(REPO) -delete v$(VERSION) $(ARTIFACTS_DIR)

dep:
	dep ensure

depup:
	dep ensure -update

builddep:
	go get -v -u github.com/golang/dep/cmd/dep github.com/laher/goxc github.com/tcnksm/ghr

clean:
	rm -rf bin dist

.PHONY: test build cross deploy dep depup clean
