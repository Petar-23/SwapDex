.PHONY: sdx sdx-cross evm all test clean
.PHONY: sdx-linux sdx-linux-386 sdx-linux-amd64 sdx-linux-mips64 sdx-linux-mips64le
.PHONY: sdx-darwin sdx-darwin-386 sdx-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.13.1
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

sdx:
	go run build/ci.go install ./cmd/sdx
	@echo "Done building."d
	@echo "Run \"$(GOBIN)/sdx\" to launch sdx."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

sdx-cross: sdx-windows-amd64 sdx-darwin-amd64 sdx-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/sdx-*

sdx-linux: sdx-linux-386 sdx-linux-amd64 sdx-linux-mips64 sdx-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-*

sdx-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/sdx
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep 386

sdx-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/sdx
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep amd64

sdx-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/sdx
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep mips

sdx-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/sdx
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep mipsle

sdx-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/sdx
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep mips64

sdx-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/sdx
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/sdx-linux-* | grep mips64le

sdx-darwin: sdx-darwin-386 sdx-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/sdx-darwin-*

sdx-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/sdx
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-darwin-* | grep 386

sdx-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/sdx
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-darwin-* | grep amd64

sdx-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) -buildmode=mode -x --targets=windows/amd64 -v ./cmd/sdx
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/sdx-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
