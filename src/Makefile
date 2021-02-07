GOCMD=go
export GOPATH=$(CURDIR)/..
export GOBIN=$(GOPATH)/bin
export PATH=$(GOPATH):$(GOBIN):$(shell printenv PATH)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
#GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROTOC=protoc
BINDIR=../bin
COPYFILES= cp ../../EWS/micro/copyfiles/*
PROTODIR=generated
PROTOINPUT=proto

.PHONY: all

all: deps build

build: $(BINDIR)/client $(BINDIR)/server


$(BINDIR)/server: $(PROTODIR)/Sample.pb.go   server/*.go
	$(GOBUILD) -o $@ $(GOPATH)/src/server

$(BINDIR)/client: $(PROTODIR)/Sample.pb.go  client/*.go
	$(GOBUILD) -o $@ $(GOPATH)/src/client

$(PROTODIR)/Sample.pb.go: $(PROTOINPUT)/Sample.proto
	$(PROTOC) -I $(PROTOINPUT)/ $(PROTOINPUT)/Sample.proto --go_out=plugins=grpc:$(PROTODIR)
#	$(COPYFILES) $(MICRODIR)

deps:
	echo "Gopath =$(GOPATH)"
	echo "GoBin  =$(GOBIN)"
	echo "path   =$(PATH)"
	$(GOGET) -u google.golang.org/grpc
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go


clean: ## remove previous builds
	echo "@rm -rf $(BINDIR)* $(PROTODIR)/*.go  "
	@rm -rf $(BINDIR)* $(PROTODIR)/*.go
