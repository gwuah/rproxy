package main

import (
	"log"
	"net"
	"os"

	"github.com/gwuah/rproxy/bpf"
)

func main() {

	// Look up the network interface by name.
	ifaceName := os.Args[1]
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatalf("lookup network iface %q: %s", ifaceName, err)
	}

	// Load bpf objects(program, maps) into the kernel
	objs, err := bpf.LoadBPFObjectsIntoKernel()
	if err != nil {
		log.Fatalf("failed to load bpf objects into kernel. err= %v:", err)
	}
	defer objs.Close()

	// Attach the program.
	l, err := bpf.AttachToHookpoint(objs, iface)
	if err != nil {
		log.Fatalf("could not attach XDP program: %s", err)
	}
	defer l.Close()

	log.Printf("Attached XDP program to iface %q (index %d)", iface.Name, iface.Index)
	log.Printf("Press Ctrl-C to exit and remove the program")

}
