DIR 		= ./build
EXECUTABLE  	= rpserve

GOARCH		= amd64
GOOSWIN		= windows
GOOSX		= darwin
GOOSLINUX	= linux
GOMOD		= on
CGO_ENABLED 	= 0

WINBIN 		= $(DIR)/$(EXECUTABLE)-win-$(GOARCH).exe
OSXBIN 		= $(DIR)/$(EXECUTABLE)-darwin-$(GOARCH)
LINUXBIN 	= $(DIR)/$(EXECUTABLE)-linux-$(GOARCH)

CC 		= go build
CFLAGS		= -trimpath
LDFLAGS		= all=-w -s
GCFLAGS 	= all=
ASMFLAGS 	= all=

.PHONY: deps
deps:
	go get ./...

.PHONY: all
all: darwin linux win64

.PHONY: darwin
darwin: $(OSXBIN)
	chmod +x $(OSXBIN)

.PHONY: $(OSXBIN)
$(OSXBIN):
	GO111MODULE=$(GOMOD) GOARCH=$(GOARCH) GOOS=$(GOOSX) CGO_ENABLED=$(CGO_ENABLED) $(CC) $(CFLAGS) -o $(OSXBIN) -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" cmd/main.go

.PHONY: linux
linux: $(LINUXBIN)
	chmod +x $(LINUXBIN)

.PHONY: $(LINUXBIN)
$(LINUXBIN):
	GO111MODULE=$(GOMOD) GOARCH=$(GOARCH) GOOS=$(GOOSLINUX) CGO_ENABLED=$(CGO_ENABLED) $(CC) $(CFLAGS) -o $(LINUXBIN) -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" cmd/main.go

.PHONY: win64
win64: $(WINBIN)

.PHONY: $(WINBIN)
$(WINBIN):
	GO111MODULE=$(GOMOD) GOARCH=$(GOARCH) GOOS=$(GOOSWIN) CGO_ENABLED=$(CGO_ENABLED) $(CC) $(CFLAGS) -o $(WINBIN) -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" cmd/main.go

.PHONY: clean
clean:
	rm -rf $(DIR)/*

.PHONY: run-dev
run-dev: deps run-containers run-server

.PHONY: run-server
runserver:
	go run cmd/main.go


## run: Start demo http services for development
.PHONY: run-containers
run-containers:
	docker run --rm -d -p 9001:80 --name server1 kennethreitz/httpbin
	docker run --rm -d -p 9002:80 --name server2 kennethreitz/httpbin
	docker run --rm -d -p 9003:80 --name server3 kennethreitz/httpbin

## stop: stops all demo services
.PHONY: stop
stop:
	docker stop server1
	docker stop server2
	docker stop server3


