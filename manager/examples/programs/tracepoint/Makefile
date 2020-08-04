all: build-ebpf build run

build-ebpf:
	mkdir -p ebpf/bin
	clang -O2 -target bpf -c ebpf/main.c -o ebpf/bin/probe.o
	go-bindata -pkg main -prefix "ebpf/bin" -o "probe.go" "ebpf/bin/probe.o"

build:
	go build -o bin/main .

run:
	sudo bin/main
