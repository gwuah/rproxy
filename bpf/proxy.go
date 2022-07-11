package bpf

import (
	"net"
	"C"
	"github.com/cilium/ebpf/link"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf ../proxy.c

func LoadBPFObjectsIntoKernel() (*proxyObjects, error) {
	objs := proxyObjects{}
	err := loadProxyObjects(&objs, nil)
	return &objs, err
}

func AttachToHookpoint(o *proxyObjects, i *net.Interface) (link.Link, error) {
	return link.AttachXDP(link.XDPOptions{
		Program:   o.ReverseProxy,
		Interface: i.Index,
	})
}

func SetCurrentPort(obj *proxyObjects, port int) error {
	return obj.proxyMaps.Ports.Put(uint32(0), uint64(port))
}
