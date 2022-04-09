GO:=go
GIT_COMMIT=$(shell git rev-list -1 HEAD)
GIT_TAG=$(shell git describe --all)
BUILD_DATE=$(shell date -uIs)
VERMOUTH=github.com/akihiro/vermouth
BUILD_INFO=-X $(VERMOUTH).GitCommit=$(GIT_COMMIT) -X $(VERMOUTH).Version=$(GIT_TAG) -X $(VERMOUTH).BuildDate=$(BUILD_DATE)
BUILD_FLAGS=-v -ldflags '-X main.Message=hello $(BUILD_INFO)'
PROG=example

.PHONY: build clean

build: $(PROG)
$(PROG): *.go
	$(GO) build -o $(PROG) $(BUILD_FLAGS) .

clean:
	$(RM) $(PROG)