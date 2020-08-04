GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
CLANG := clang
CLANG_INCLUDE := -I./ebpf_prog
CFLAGS :=

GO_SOURCE := main.go
GO_BINARY := bio_tracepoint

EBPF_SOURCE := ebpf_prog/bio_tracepoint.c
EBPF_BINARY := ebpf_prog/bio_tracepoint.elf

all: build_bpf build_go

build_bpf: $(EBPF_BINARY)

build_go: $(GO_BINARY)

clean:
	$(GOCLEAN)
	rm -f $(GO_BINARY)
	rm -f $(EBPF_BINARY)

$(EBPF_BINARY): $(EBPF_SOURCE)
	$(CLANG) $(CLANG_INCLUDE) $(CFLAGS) -O2 -target bpf -c $^  -o $@
	go-bindata -pkg main -prefix "ebpf_prog" -o "bio_tracepoint_ebpf.go" "ebpf_prog/bio_tracepoint.elf"

$(GO_BINARY): $(GO_SOURCE)
	$(GOBUILD) -v -o $@
