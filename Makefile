VERSION := 0.1.0
BUILD_DATE := $(shell date +%F)

all: httpdump

httpdump:
	gb build -ldflags " \
		-X main.version=$(VERSION) \
		-X main.buildDate='$(BUILD_DATE)' \
		"

clean:
	rm -rf bin/ pkg/
