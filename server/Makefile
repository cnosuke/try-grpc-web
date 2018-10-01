NAME     := todo-server
VERSION  := $(shell git describe --tags 2>/dev/null)
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null)
SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: clean cross-build release-pack gen-proto

bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

clean:
	rm -rf bin/* dist/* vendor/*

cross-build:
	for os in darwin linux; do \
		GOOS=$$os GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-amd64/$(NAME); \
	done

release-pack: cross-build
	for os in darwin linux; do \
		zip -j dist/$$os-amd64.zip dist/$$os-amd64/$(NAME); \
	done

gen-proto:
	@mkdir -p pb/
	@protoc \
		-I proto/ \
		--go_out=plugins=grpc:pb/ \
		proto/*.proto
