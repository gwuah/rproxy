CLANG ?= clang
CFLAGS := -O2 -g -Wall -Werror $(CFLAGS)

export BPF_CFLAGS := $(CFLAGS)
export BPF_CLANG := $(CLANG)

deps:
	go install

proxy-files:
	cd proxy; \
	go generate ./...; \
	rm -rf *_bpfeb.go && rm -rf *_bpfeb.o