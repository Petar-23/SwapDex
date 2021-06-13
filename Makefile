.PHONY: SDX SDX-cross evm all test clean
.PHONY: SDX-linux SDX-linux-386 SDX-linux-amd64 SDX-linux-mips64 SDX-linux-mips64le
.PHONY: SDX-darwin SDX-darwin-386 SDX-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

SDX:
	build/env.sh go run build/ci.go install ./cmd/SDX
	@echo "Done building."
	@echo "Run \"$(GOBIN)/SDX\" to launch SDX."

gc:
	build/env.sh go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

SDX-cross: SDX-linux SDX-darwin
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/SDX-*

SDX-linux: SDX-linux-386 SDX-linux-amd64 SDX-linux-mips64 SDX-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-*

SDX-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/SDX
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep 386

SDX-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/SDX
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep amd64

SDX-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/SDX
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep mips

SDX-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/SDX
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep mipsle

SDX-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/SDX
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep mips64

SDX-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/SDX
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/SDX-linux-* | grep mips64le

SDX-darwin: SDX-darwin-386 SDX-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/SDX-darwin-*

SDX-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/SDX
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/SDX-darwin-* | grep 386

SDX-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/SDX
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/SDX-darwin-* | grep amd64

gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
