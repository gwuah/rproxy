CLANG ?= clang
CFLAGS := -O2 -g -Wall -Werror $(CFLAGS)

export BPF_CFLAGS := $(CFLAGS)
export BPF_CLANG := $(CLANG)

deps:
	go install

proxy-files:
	cd bpf; \
	go generate ./...; \
	rm -rf bpf_bpfeb.go && rm -rf bpf_bpfeb.o