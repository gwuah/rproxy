#include <stddef.h>
#include <linux/bpf.h>
#include <linux/in.h>
#include <linux/if_ether.h>
#include <linux/ip.h>

#include <bpf/bpf_endian.h>
#include <bpf/bpf_helpers.h>

struct bpf_map_def SEC("maps") ports = {
	.type		= BPF_MAP_TYPE_HASH,
	.max_entries	= 1,
	.key_size	= sizeof(__u16),
	.value_size	= sizeof(__u8),
};

SEC("xdp")
int reverse_proxy(struct xdp_md *ctx)
{
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;

    struct ethhdr *eth = data;
    if (data + sizeof(struct ethhdr) > data_end)
        return XDP_ABORTED;

    if (bpf_ntohs(eth->h_proto) != ETH_P_IP)
        return XDP_PASS;

    struct iphdr *iph = data + sizeof(struct ethhdr);
    if (data + sizeof(struct ethhdr) + sizeof(struct iphdr) > data_end)
        return XDP_ABORTED;

    if (iph->protocol != IPPROTO_TCP)
        return XDP_PASS;
    
    // struct tcphdr *tcp = (void*)iph + sizeof(*iph);
    // if ((void*)tcp + sizeof(*tcp) <= data_end) 

    bpf_printk("Got TCP packet from %x", iph->saddr);

    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";