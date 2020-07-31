package manager

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/florianl/go-tc"
	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"github.com/ismhong/ebpf"
	"github.com/ismhong/ebpf/internal"
)

// XdpAttachMode selects a way how XDP program will be attached to interface
type XdpAttachMode int

const (
	// XdpAttachModeNone stands for "best effort" - kernel automatically
	// selects best mode (would try Drv first, then fallback to Generic).
	// NOTE: Kernel will not fallback to Generic XDP if NIC driver failed
	//       to install XDP program.
	XdpAttachModeNone XdpAttachMode = 0
	// XdpAttachModeSkb is "generic", kernel mode, less performant comparing to native,
	// but does not requires driver support.
	XdpAttachModeSkb XdpAttachMode = (1 << 1)
	// XdpAttachModeDrv is native, driver mode (support from driver side required)
	XdpAttachModeDrv XdpAttachMode = (1 << 2)
	// XdpAttachModeHw suitable for NICs with hardware XDP support
	XdpAttachModeHw XdpAttachMode = (1 << 3)
)

type TrafficType uint32

const (
	Ingress TrafficType = tc.Ingress
	Egress  TrafficType = tc.Egress
)

type ProbeIdentificationPair struct {
	UID     string
	Section string
}

func (pip ProbeIdentificationPair) String() string {
	return fmt.Sprintf("[UID:%s, Section:%s]", pip.UID, pip.Section)
}

// Probe - Main eBPF probe wrapper. This structure is used to store the required data to attach a loaded eBPF
// program to its hook point.
type Probe struct {
	manager          *Manager
	program          *ebpf.Program
	programSpec      *ebpf.ProgramSpec
	perfEventFD      *internal.FD
	state            state
	stateLock        *sync.RWMutex
	manualLoadNeeded bool
	checkPin         bool
	funcName         string
	attachPID        int

	// UID - (optional) this field can be used to identify your probes when the same eBPF program is used on multiple
	// hook points. Keep in mind that the pair (probe section, probe UID) needs to be unique
	// system-wide for the kprobes and uprobes registration to work.
	UID string

	// Section - Section of the program, as defined in its section SEC("[section]"). This section is therefore made of
	// a prefix
	Section string

	// SyscallFuncName - Name of the syscall on which the program should be hooked. As the exact kernel symbol may
	// differ from one kernel version to the other, the right prefix will be computed automatically at runtime.
	// If a syscall name is not provided, the section name (without its probe type prefix) is assumed to be the
	// hook point.
	SyscallFuncName string

	// MatchFuncName - When this option is activated, the provided symbol is matched against the
	// list of available symbols in /sys/kernel/debug/tracing/available_filter_functions. If the exact function does not
	// exist, then the closest function will be used. This requires debugfs.
	MatchFuncName string

	// Enabled - Indicates if a probe should be enabled or not. This parameter can be set at runtime using the
	// Manager options (see ActivatedProbes)
	Enabled bool

	// PinPath - Once loaded, the eBPF program will be pinned to this path. If the eBPF program has already been pinned
	// and is already running in the kernel, then it will be loaded from this path.
	PinPath string

	// KProbeMaxActive - (kretprobes) With kretprobes, you can configure the maximum number of instances of the function that can be
	// probed simultaneously with maxactive. If maxactive is 0 it will be set to the default value: if CONFIG_PREEMPT is
	// enabled, this is max(10, 2*NR_CPUS); otherwise, it is NR_CPUS. For kprobes, maxactive is ignored.
	KProbeMaxActive int

	// BinaryPath - (uprobes) A Uprobe is attached to a specific symbol in a user space binary. The offset is
	// automatically computed for the symbol name provided in the uprobe section ( SEC("uprobe/[symbol_name]") ).
	BinaryPath string

	// CGrouPath - (cgroup family programs) All CGroup programs are attached to a CGroup (v2). This field provides the
	// path to the CGroup to which the probe should be attached. The attach type is determined by the section.
	CGroupPath string

	// SocketFD - (socket filter) Socket filter programs are bound to a socket and filter the packets they receive
	// before they reach user space. The probe will be bound to the provided file descriptor
	SocketFD int

	// Ifindex - (TC classifier & XDP) Interface index used to identify the interface on which the probe will be
	// attached. If not set, fall back to Ifname.
	Ifindex int32

	// Ifname - (TC Classifier & XDP) Interface name on which the probe will be attached.
	Ifname string

	// IfindexNetns - (TC Classifier & XDP) Network namespace in which the network interface lives
	IfindexNetns uint64

	// XDPAttachMode - (XDP) XDP attach mode. If not provided the kernel will automatically select the best available
	// mode.
	XDPAttachMode XdpAttachMode

	// NetworkDirection - (TC classifier) Network traffic direction of the classifier. Can be either Ingress or Egress. Keep
	// in mind that if you are hooking on the host side of a virtuel ethernet pair, Ingress and Egress are inverted.
	NetworkDirection TrafficType

	// tcObject - (TC classifier) TC object created when the classifier was attached. It will be reused to delete it on
	// exit.
	tcObject *tc.Object

	// PerfEventSampleFrequency - (perf event) The sample frequency for perf_event
	PerfEventSampleFrequency uint

	// PerfEventType - (perf event) The sample type for perf_event. ex. unix.PERF_TYPE_HARDWARE, unix.PERF_TYPE_SOFTWARE, PERF_TYPE_HW_CACHE
	PerfEventType uint

	// PerfEventConfig - (perf event) The sample config for perf_event. ex. unix.PERF_COUNT_HW_CPU_CYCLES, unix.PERF_COUNT_HW_INSTRUCTIONS
	PerfEventConfig uint

	// PerfEventPid - (perf event) The Pid for perf_event.
	PerfEventPid int

	// PerfEventCpuId - (perf event) The CPU id for perf_event. -1 for all CPUs
	PerfEventCpuId int
}

// IdentificationPairMatches - Returns true if the identification pair (probe uid, probe section) match.
func (p *Probe) IdentificationPairMatches(id ProbeIdentificationPair) bool {
	return p.UID == id.UID && p.Section == id.Section
}

// GetIdentificationPair - Returns the identification pair (probe section, probe UID)
func (p *Probe) GetIdentificationPair() ProbeIdentificationPair {
	return ProbeIdentificationPair{p.UID, p.Section}
}

// IsRunning - Returns true if the probe was successfully initialized, started and is currently running.
func (p *Probe) IsRunning() bool {
	p.stateLock.RLock()
	defer p.stateLock.RUnlock()
	return p.state == running
}

// IsInitialized - Returns true if the probe was successfully initialized, started and is currently running.
func (p *Probe) IsInitialized() bool {
	p.stateLock.RLock()
	defer p.stateLock.RUnlock()
	return p.state >= initialized
}

// Test - Triggers the probe with the provided test data. Returns the length of the output, the raw output or an error.
func (p *Probe) Test(in []byte) (uint32, []byte, error) {
	return p.program.Test(in)
}

// Benchmark - Benchmark runs the Program with the given input for a number of times and returns the time taken per
// iteration.
//
// Returns the result of the last execution of the program and the time per run or an error. reset is called whenever
// the benchmark syscall is interrupted, and should be set to testing.B.ResetTimer or similar.
func (p *Probe) Benchmark(in []byte, repeat int, reset func()) (uint32, time.Duration, error) {
	return p.program.Benchmark(in, repeat, reset)
}

// InitWithOptions - Initializes a probe with options
func (p *Probe) InitWithOptions(manager *Manager, manualLoadNeeded bool, checkPin bool) error {
	p.stateLock = &sync.RWMutex{}
	if !p.Enabled {
		return nil
	}
	p.manager = manager
	p.stateLock.Lock()
	defer p.stateLock.Unlock()
	p.state = reset
	p.manualLoadNeeded = manualLoadNeeded
	p.checkPin = checkPin
	return p.init()
}

// Init - Initialize a probe
func (p *Probe) Init(manager *Manager) error {
	p.stateLock = &sync.RWMutex{}
	if !p.Enabled {
		return nil
	}
	p.manager = manager
	p.stateLock.Lock()
	defer p.stateLock.Unlock()
	p.state = reset
	return p.init()
}

// init - Internal initialization function
func (p *Probe) init() error {
	// Load spec if necessary
	if p.manualLoadNeeded {
		prog, err := ebpf.NewProgramWithOptions(p.programSpec, p.manager.options.VerifierOptions.Programs)
		if err != nil {
			return errors.Wrapf(err, "couldn't load new probe %v", p.GetIdentificationPair())
		}
		p.program = prog
	}

	// Retrieve eBPF program if one isn't already set
	if p.program == nil {
		prog, ok := p.manager.collection.Programs[p.Section]
		if !ok {
			return errors.Wrapf(ErrUnknownSection, "couldn't find program %s", p.Section)
		}
		p.program = prog
		p.checkPin = true
	}

	if p.checkPin {
		// Pin program if needed
		if p.PinPath != "" {
			if err := p.program.Pin(p.PinPath); err != nil {
				return errors.Wrapf(err, "couldn't pin program %s at %s", p.Section, p.PinPath)
			}
		}
		p.checkPin = false
	}

	// Update syscall function name with the correct arch prefix
	if p.SyscallFuncName != "" {
		var err error
		p.funcName, err = GetSyscallFnNameWithSymFile(p.SyscallFuncName, p.manager.options.SymFile)
		if err != nil {
			return err
		}
	}

	// Find function name match if required
	if p.MatchFuncName != "" {
		var err error
		p.funcName, err = FindFilterFunction(p.MatchFuncName)
		if err != nil {
			return err
		}
	}

	// Resolve interface index if one is provided
	if p.Ifindex == 0 && p.Ifname != "" {
		inter, err := net.InterfaceByName(p.Ifname)
		if err != nil {
			return errors.Wrapf(err, "couldn't find interface %v", p.Ifname)
		}
		p.Ifindex = int32(inter.Index)
	}

	// update probe state
	p.state = initialized
	return nil
}

// Attach - Attaches the probe to the right hook point in the kernel depending on the program type and the provided
// parameters.
func (p *Probe) Attach() error {
	p.stateLock.Lock()
	defer p.stateLock.Unlock()
	if p.state >= running || !p.Enabled {
		return nil
	}
	if p.state < initialized {
		return ErrProbeNotInitialized
	}

	// Per program type start
	var err error
	switch p.programSpec.Type {
	case ebpf.UnspecifiedProgram:
		err = errors.Wrap(ErrSectionFormat, "invalid program type, make sure to use the right section prefix")
	case ebpf.Kprobe:
		err = p.attachKprobe()
	case ebpf.TracePoint:
		err = p.attachTracepoint()
	case ebpf.PerfEvent:
		err = p.attachPerfEvent()
	case ebpf.CGroupDevice, ebpf.CGroupSKB, ebpf.CGroupSock, ebpf.CGroupSockAddr, ebpf.CGroupSockopt, ebpf.CGroupSysctl:
		err = p.attachCGroup()
	case ebpf.SocketFilter:
		err = p.attachSocket()
	case ebpf.SchedCLS:
		err = p.attachTCCLS()
	case ebpf.XDP:
		err = p.attachXDP()
	default:
		err = fmt.Errorf("program type %s not implemented yet", p.programSpec.Type)
	}
	if err != nil {
		// Clean up any progress made in the attach attempt
		_ = p.stop()
		return errors.Wrapf(err, "couldn't start probe %s", p.Section)
	}

	// update probe state
	p.state = running
	return nil
}

// Detach - Detaches the probe from its hook point depending on the program type and the provided parameters. This
// method does not close the underlying eBPF program, which means that Attach can be called again later.
func (p *Probe) Detach() error {
	p.stateLock.Lock()
	defer p.stateLock.Unlock()
	if p.state < running || !p.Enabled {
		return nil
	}

	// detach from hook point
	err := p.detach()

	// update state of the probe
	if err != nil {
		p.state = initialized
	}
	return err
}

// detach - Thread unsafe version of Detach.
func (p *Probe) detach() error {
	var err error
	// Remove pin if needed
	if p.PinPath != "" {
		err = ConcatErrors(err, os.Remove(p.PinPath))
	}

	// Shared with all probes: close the perf event file descriptor
	if p.perfEventFD != nil {
		err = p.perfEventFD.Close()
	}

	// Per program type cleanup
	switch p.programSpec.Type {
	case ebpf.UnspecifiedProgram:
		// nothing to do
		break
	case ebpf.Kprobe:
		err = ConcatErrors(err, p.detachKprobe())
	case ebpf.CGroupDevice, ebpf.CGroupSKB, ebpf.CGroupSock, ebpf.CGroupSockAddr, ebpf.CGroupSockopt, ebpf.CGroupSysctl:
		err = ConcatErrors(err, p.detachCgroup())
	case ebpf.SocketFilter:
		err = ConcatErrors(err, p.detachSocket())
	case ebpf.SchedCLS:
		err = ConcatErrors(err, p.detachTCCLS())
	case ebpf.XDP:
		err = ConcatErrors(err, p.detachXDP())
	default:
		// unsupported section, nothing to do either
		break
	}
	return err
}

// Stop - Detaches the probe from its hook point and close the underlying eBPF program.
func (p *Probe) Stop() error {
	p.stateLock.Lock()
	defer p.stateLock.Unlock()
	if p.state < running || !p.Enabled {
		return nil
	}
	return p.stop()
}

func (p *Probe) stop() error {
	// detach from hook point
	err := p.detach()

	// close the loaded program
	err = ConcatErrors(err, p.program.Close())

	// update state of the probe
	if err != nil {
		p.state = 0
	}
	return errors.Wrapf(err, "coulnd't cleanup probe %s", p.Section)
}

// attachKprobe - Attaches the probe to its kprobe
func (p *Probe) attachKprobe() error {
	// Prepare kprobe_events line parameters
	var probeType, maxactiveStr string
	var err error
	funcName := p.funcName
	if strings.HasPrefix(p.Section, "kretprobe/") {
		if funcName == "" {
			funcName = strings.TrimPrefix(p.Section, "kretprobe/")
		}
		if p.KProbeMaxActive > 0 {
			maxactiveStr = fmt.Sprintf("%d", p.KProbeMaxActive)
		}
		probeType = "r"
	} else if strings.HasPrefix(p.Section, "kprobe/") {
		if funcName == "" {
			funcName = strings.TrimPrefix(p.Section, "kprobe/")
		}
		probeType = "p"
	} else {
		// this might actually be a Uprobe
		return p.attachUprobe()
	}
	p.attachPID = os.Getpid()

	// Write kprobe_events line to register kprobe
	kprobeID, err := EnableKprobeEvent(probeType, funcName, p.UID, maxactiveStr, p.attachPID)
	// fallback without KProbeMaxActive
	if err == ErrKprobeIDNotExist {
		kprobeID, err = EnableKprobeEvent(probeType, funcName, p.UID, "", p.attachPID)
	}
	if err != nil {
		return errors.Wrapf(err, "couldn't enable kprobe %s", p.Section)
	}

	// Activate perf event
	p.perfEventFD, err = perfEventOpenTracepoint(kprobeID, p.program.FD())
	return errors.Wrapf(err, "couldn't enable kprobe %s", p.Section)
}

// detachKprobe - Detaches the probe from its kprobe
func (p *Probe) detachKprobe() error {
	// Prepare kprobe_events line parameters
	funcName := p.funcName
	probeType := ""
	if strings.HasPrefix(p.Section, "kretprobe/") {
		if funcName == "" {
			funcName = strings.TrimPrefix(p.Section, "kretprobe/")
		}
		probeType = "r"
	} else if strings.HasPrefix(p.Section, "kprobe/") {
		if funcName == "" {
			funcName = strings.TrimPrefix(p.Section, "kprobe/")
		}
		probeType = "p"
	} else {
		// this might be a Uprobe
		return p.detachUprobe()
	}

	// Write kprobe_events line to remove hook point
	return DisableKprobeEvent(probeType, funcName, p.UID, p.attachPID)
}

// attachTracepoint - Attaches the probe to its tracepoint
func (p *Probe) attachTracepoint() error {
	// Parse section
	traceGroup := strings.SplitN(p.Section, "/", 3)
	if len(traceGroup) != 3 {
		return errors.Wrapf(ErrSectionFormat, "expected SEC(\"tracepoint/[category]/[name]\") got %s", p.Section)
	}
	category := traceGroup[1]
	name := traceGroup[2]

	// Get the ID of the tracepoint to activate
	tracepointID, err := GetTracepointID(category, name)
	if err != nil {
		return errors.Wrapf(err, "couldn's activate tracepoint %s", p.Section)
	}

	// Hook the eBPF program to the tracepoint
	p.perfEventFD, err = perfEventOpenTracepoint(tracepointID, p.program.FD())
	return errors.Wrapf(err, "couldn't enable tracepoint %s", p.Section)
}

// attachPerfEvent - Attaches the probe to its perf event
func (p *Probe) attachPerfEvent() error {
	// Parse section
	if strings.HasPrefix(p.Section, "perf_event/") == false {
		// unknown type
		return errors.Wrapf(ErrSectionFormat, "program type unrecognized in section %v", p.Section)
	}

	// Hook the eBPF program to the perf event
	var err error
	p.perfEventFD, err = perfEventOpenRawEvent(p.PerfEventType, p.PerfEventConfig, p.PerfEventSampleFrequency, p.PerfEventPid, p.PerfEventCpuId, p.program.FD())
	return errors.Wrapf(err, "couldn't enable raw perf event %s", p.Section)
}

// attachUprobe - Attaches the probe to its Uprobe
func (p *Probe) attachUprobe() error {
	// Prepare uprobe_events line parameters
	var probeType, funcName string
	if strings.HasPrefix(p.Section, "uretprobe/") {
		funcName = strings.TrimPrefix(p.Section, "uretprobe/")
		probeType = "r"
	} else if strings.HasPrefix(p.Section, "uprobe/") {
		funcName = strings.TrimPrefix(p.Section, "uprobe/")
		probeType = "p"
	} else {
		// unknown type
		return errors.Wrapf(ErrSectionFormat, "program type unrecognized in section %v", p.Section)
	}
	p.attachPID = os.Getpid()

	// Write uprobe_events line to register uprobe
	uprobeID, err := EnableUprobeEvent(probeType, funcName, p.BinaryPath, p.UID, p.attachPID)
	if err != nil {
		return errors.Wrapf(err, "couldn't enable uprobe %s", p.Section)
	}

	// Activate perf event
	p.perfEventFD, err = perfEventOpenTracepoint(uprobeID, p.program.FD())
	return errors.Wrapf(err, "couldn't enable kprobe %s", p.Section)
}

// detachUprobe - Detaches the probe from its Uprobe
func (p *Probe) detachUprobe() error {
	// Prepare uprobe_events line parameters
	var probeType, funcName string
	if strings.HasPrefix(p.Section, "uretprobe/") {
		funcName = strings.TrimPrefix(p.Section, "uretprobe/")
		probeType = "r"
	} else if strings.HasPrefix(p.Section, "uprobe/") {
		funcName = strings.TrimPrefix(p.Section, "uprobe/")
		probeType = "p"
	} else {
		// unknown type
		return errors.Wrapf(ErrSectionFormat, "program type unrecognized in section %v", p.Section)
	}

	// Write uprobe_events line to remove hook point
	return DisableUprobeEvent(probeType, funcName, p.BinaryPath, p.UID, p.attachPID)
}

// attachCGroup - Attaches the probe to a cgroup hook point
func (p *Probe) attachCGroup() error {
	// open CGroupPath
	f, err := os.Open(p.CGroupPath)
	if err != nil {
		return errors.Wrapf(err, "error opening cgroup %s from probe %s", p.CGroupPath, p.Section)
	}
	defer f.Close()

	// Attach CGroup
	ret, err := bpfProgAttach(p.program.FD(), int(f.Fd()), p.programSpec.AttachType)
	if ret < 0 {
		return errors.Wrapf(err, "failed to attach probe %v to cgroup %s", p.GetIdentificationPair(), p.CGroupPath)
	}
	return nil
}

// detachCGroup - Detaches the probe from its cgroup hook point
func (p *Probe) detachCgroup() error {
	// open CGroupPath
	f, err := os.Open(p.CGroupPath)
	if err != nil {
		return errors.Wrapf(err, "error opening cgroup %s from probe %s", p.CGroupPath, p.Section)
	}

	// Detach CGroup
	ret, err := bpfProgDetach(p.program.FD(), int(f.Fd()), p.programSpec.AttachType)
	if ret < 0 {
		return errors.Wrapf(err, "failed to detach probe %v from cgroup %s", p.GetIdentificationPair(), p.CGroupPath)
	}
	return nil
}

// attachSocket - Attaches the probe to the provided socket
func (p *Probe) attachSocket() error {
	return sockAttach(p.SocketFD, p.program.FD())
}

// detachSocket - Detaches the probe from its socket
func (p *Probe) detachSocket() error {
	return sockDetach(p.SocketFD, p.program.FD())
}

// attachTCCLS - Attaches the probe to its TC classifier hook point
func (p *Probe) attachTCCLS() error {
	var err error
	// Make sure Ifindex is properly set
	if p.Ifindex == 0 && p.Ifname == "" {
		return ErrInterfaceNotSet
	}

	// Recover the netlink socket of the interface from the manager
	ntl, ok := p.manager.netlinkCache[netlinkCacheKey{p.Ifindex, p.IfindexNetns}]
	if !ok {
		// Set up new netlink connection
		ntl, err = p.manager.newNetlinkConnection(p.Ifindex, p.IfindexNetns)
		if err != nil {
			return err
		}
	}

	// Create a Qdisc for the provided interface
	qdisc := &tc.Object{
		Msg: tc.Msg{
			Family:  unix.AF_UNSPEC,
			Ifindex: uint32(p.Ifindex),
			Handle:  tc.BuildHandle(0xFFFF, 0x0000),
			Parent:  0xFFFFFFF1,
			Info:    0,
		},
		Attribute: tc.Attribute{
			Kind: "clsact",
		},
	}

	// Add the Qdisc
	err = ntl.rtNetlink.Qdisc().Add(qdisc)
	if err != nil {
		if err.Error() != "netlink receive: file exists" {
			return errors.Wrapf(err, "couldn't add a \"clsact\" qdisc to interface %v", p.Ifindex)
		}
	}

	// Create qdisc filter
	filter := tc.Object{
		Msg: tc.Msg{
			Family:  unix.AF_UNSPEC,
			Ifindex: uint32(p.Ifindex),
			Handle:  0,
			Parent:  uint32(p.NetworkDirection),
			Info:    0x300,
		},
		Attribute: tc.Attribute{
			Kind: "bpf",

			BPF: &tc.Bpf{
				FD:    uint32(p.program.FD()),
				Name:  p.Section,
				Flags: 0x1,
			},
		},
	}

	// Add qdisc filter
	err = ntl.rtNetlink.Filter().Add(&filter)
	if err == nil {
		p.tcObject = qdisc
		ntl.schedClsCount += 1
	}
	return errors.Wrapf(err, "couldn't add a %v filter to interface %v: %v", p.NetworkDirection, p.Ifindex, err)
}

// detachTCCLS - Detaches the probe from its TC classifier hook point
func (p *Probe) detachTCCLS() error {
	// Recover the netlink socket of the interface from the manager
	ntl, ok := p.manager.netlinkCache[netlinkCacheKey{p.Ifindex, p.IfindexNetns}]
	if !ok {
		return fmt.Errorf("coulnd't find qdisc from which the probe %v was meant to be detached", p.GetIdentificationPair())
	}

	if ntl.schedClsCount >= 2 {
		ntl.schedClsCount -= 1
		// another classifier is still using the qdisc, do not delete it yet
		return nil
	}

	// Delete qdisc
	err := ntl.rtNetlink.Qdisc().Delete(p.tcObject)
	return errors.Wrapf(err, "couldn't detach TC classifier of probe %v", p.GetIdentificationPair())
}

// attachXDP - Attaches the probe to an interface with an XDP hook point
func (p *Probe) attachXDP() error {
	// Lookup interface
	link, err := netlink.LinkByIndex(int(p.Ifindex))
	if err != nil {
		return errors.Wrapf(err, "couldn't retrieve interface %v", p.Ifindex)
	}

	// Attach program
	err = netlink.LinkSetXdpFdWithFlags(link, p.program.FD(), int(p.XDPAttachMode))
	return errors.Wrapf(err, "couldn't attach XDP program %v to interface %v", p.GetIdentificationPair(), p.Ifindex)
}

// detachXDP - Detaches the probe from its XDP hook point
func (p *Probe) detachXDP() error {
	// Lookup interface
	link, err := netlink.LinkByIndex(int(p.Ifindex))
	if err != nil {
		return errors.Wrapf(err, "couldn't retrieve interface %v", p.Ifindex)
	}

	// Detach program
	err = netlink.LinkSetXdpFdWithFlags(link, -1, int(p.XDPAttachMode))
	return errors.Wrapf(err, "couldn't detach XDP program %v from interface %v", p.GetIdentificationPair(), p.Ifindex)
}
