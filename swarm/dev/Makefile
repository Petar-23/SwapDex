.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/69th-byte/SmartDex-Chain/cmd/swarm
	go build -o bin/sdx     github.com/69th-byte/SmartDex-Chain/cmd/sdx
	go build -o bin/bootnode github.com/69th-byte/SmartDex-Chain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/69th-byte/SmartDex-Chain/swarm/...
