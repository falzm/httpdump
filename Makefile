VERSION := 0.1.1
BUILD_DATE := $(shell date +%F)

all: httpdump

httpdump:
	go build -mod vendor -ldflags " \
		-X main.version=$(VERSION) \
		-X main.buildDate=$(BUILD_DATE) \
		"

clean:
	rm -rf httpdump bin/ pkg/
