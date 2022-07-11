package bpf

import (
	"net"
	"github.com/cilium/ebpf/link"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf ../proxy.c

func LoadBPFObjectsIntoKernel() (*bpfObjects, error) {
	objs := bpfObjects{}
	err := loadBpfObjects(&objs, nil)
	return &objs, err
}

func AttachToHookpoint(o *bpfObjects, i *net.Interface) (link.Link, error) {
	return link.AttachXDP(link.XDPOptions{
		Program:   o.ReverseProxy,
		Interface: i.Index,
	})
}

func SetCurrentPort(obj *bpfObjects, port int) error {
	return obj.bpfMaps.Ports.Put(uint32(0), uint64(port))
}
