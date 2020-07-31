/*#include <linux/ptrace.h>*/
/*#include <uapi/linux/bpf_perf_event.h>*/
#include "bpf_helpers.h"

/*const int max_cpus = 128;*/

/*BPF_ARRAY(instructions, u64, max_cpus);*/
BPF_MAP_DEF(instructions) = {
    .map_type = BPF_MAP_TYPE_ARRAY,
    .max_entries = 1,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u64),
};
BPF_MAP_ADD(instructions);

/*BPF_ARRAY(cycles, u64, max_cpus);*/
BPF_MAP_DEF(cycles) = {
    .map_type = BPF_MAP_TYPE_ARRAY,
    .max_entries = 1,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u64),
};
BPF_MAP_ADD(cycles);

SEC("perf_event/on_cpu_instruction")
int on_cpu_instruction(void *pt_regs) {
    /*instructions.increment(bpf_get_smp_processor_id(), ctx->sample_period);*/
    __u32 index = 0;
    __u64 *ins = (__u64 *)bpf_map_lookup_elem(&instructions, &index);
    __u64 new_ins;
    if (ins) {
        new_ins = *ins + 99;
        bpf_map_update_elem(&instructions, &index, &new_ins, 0);
    }
    return 0;
}

SEC("perf_event/on_cpu_cycle")
int on_cpu_cycle(void *pt_regs) {
    /*cycles.increment(bpf_get_smp_processor_id(), ctx->sample_period);*/
    __u32 index = 0;
    __u64 *ins = (__u64 *)bpf_map_lookup_elem(&cycles, &index);
    __u64 new_ins;
    if (ins) {
        new_ins = *ins + 99;
        bpf_map_update_elem(&cycles, &index, &new_ins, 0);
    }
    return 0;
}

char _license[] SEC("license") = "GPL";
