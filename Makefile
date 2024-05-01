GOCMD=go
GODEV=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build

build:
	$(GOBUILD) -o ./dist/devforge .

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f mi-programa

dev:
	$(GODEV) ./main.go