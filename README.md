## rproxy
🤟🏾 the world's greatest reverse proxy

## how to run
- Run `make deps` to install dependencies
- Run `make proxy-files` to generate bpf object files
- Run `go run main.go` to initialize application

## how it works
- the bpf program accepts a port & drops all packets to that port.
- for every packet whose destination port is a **reverse** of the accepted port, it modifies the packet and sets the destination port to the accepted port.
- and so for example, if our host server is listening on port 1234, the xdp program will drop all packets targetting 1234. But we'll redirect all packets targetting 4321 to 1234.

## but why
welll cos I can & ebpf is fun

## status
wip wip. currently facing some blockers attaching it to the network interface.
