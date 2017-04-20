SHELL			:=	/bin/bash
PKGS			:=	reppa
LD_FLAGS		:= -ldflags "-w -s -X main.Version=$(shell cat ./VERSION)-$(shell git rev-parse --short HEAD)"

all: $(addsuffix .out, $(PKGS))

image: $(addsuffix .linux.amd64, $(PKGS))
	docker build -t beldpro/reppa .


install: $(addprefix install-, $(PKGS))


deps:
	glide install


install-%:
	cd $* && go install -v


clean:
	find . -name "*.out" -type f -delete
	find . -name "*.linux.amd64" -type f -delete


%.linux.amd64:
	cd $* && gofmt -s -w .
	cd $* && GOOS=linux GOARCH=amd64 GCO_ENABLED=0 go build $(LD_FLAGS) -v -o $@

%.out:
	cd $* && gofmt -s -w .
	cd $* && go build $(LD_FLAGS) -v -o $@


.PHONY: deps install lint test clean image

