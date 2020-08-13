GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
CLANG := clang
CLANG_INCLUDE := -I./ebpf_prog
CLANG_CFLAGS := -D__x86_64__

GO_SOURCE := main.go
GO_BINARY := ipcstat

EBPF_SOURCE := ebpf_prog/ipcstat.c
EBPF_BINARY := ebpf_prog/ipcstat.elf

all: build_bpf build_go

build_bpf: $(EBPF_BINARY)

build_go: $(GO_BINARY)

clean:
	$(GOCLEAN)
	rm -f $(GO_BINARY)
	rm -f $(EBPF_BINARY)

$(EBPF_BINARY): $(EBPF_SOURCE)
	$(CLANG) $(CLANG_INCLUDE) $(CLANG_CFLAGS) -O2 -target bpf -c $^  -o $@
	go-bindata -pkg main -prefix "ebpf_prog" -o "ipcstat_ebpf.go" "ebpf_prog/ipcstat.elf"

$(GO_BINARY): $(GO_SOURCE)
	$(GOBUILD) -v -o $@
