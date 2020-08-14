package manager

import (
	"github.com/ismhong/ebpf"
	"golang.org/x/sys/unix"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/ismhong/ebpf/internal"
)

func perfEventOpenTracepoint(id int, progFd int) ([]*internal.FD, error) {
	retEventFD := []*internal.FD{}

	attr := unix.PerfEventAttr{
		Type:        unix.PERF_TYPE_TRACEPOINT,
		Sample_type: unix.PERF_SAMPLE_RAW,
		Sample:      1,
		Wakeup:      1,
		Config:      uint64(id),
	}
	attr.Size = uint32(unsafe.Sizeof(attr))

	efd, err := unix.PerfEventOpen(&attr, -1, 0, -1, unix.PERF_FLAG_FD_CLOEXEC)
	if efd < 0 {
		return nil, errors.Wrap(err, "perf_event_open error")
	}

	if _, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(efd), unix.PERF_EVENT_IOC_ENABLE, 0); err != 0 {
		return nil, errors.Wrap(err, "error enabling perf event")
	}

	if _, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(efd), unix.PERF_EVENT_IOC_SET_BPF, uintptr(progFd)); err != 0 {
		return nil, errors.Wrap(err, "error attaching bpf program to perf event")
	}
	retEventFD = append(retEventFD, internal.NewFD(uint32(efd)))
	return retEventFD, nil
}

func perfEventOpenRawEvent(eventType, eventConfig uint, eventFreq, eventPeriod uint64, eventPid, eventCpuId, progFd int) ([]*internal.FD, error) {
	retEventFD := []*internal.FD{}

	attr := unix.PerfEventAttr{
		Type:        uint32(eventType),
		Sample_type: unix.PERF_SAMPLE_RAW,
		Config:      uint64(eventConfig),
	}
	if eventFreq != 0 {
		attr.Sample = eventFreq
		attr.Bits |= 1 << 10 // use frequency, not period
	} else if eventPeriod != 0 {
		attr.Sample = eventPeriod
	} else {
		var err error
		return nil, errors.Wrap(err, "perf_event_open error, need assign period or frequency")
	}
	attr.Size = uint32(unsafe.Sizeof(attr))

	cpuNumber, err := internal.PossibleCPUs()
	if err != nil {
		cpuNumber = 1
	}
	cpuIndex := 0

	if eventCpuId != -1 {
		cpuIndex = eventCpuId
		cpuNumber = eventCpuId + 1
	}

	for ; cpuIndex < cpuNumber; cpuIndex++ {
		efd, err := unix.PerfEventOpen(&attr, eventPid, cpuIndex, -1, unix.PERF_FLAG_FD_CLOEXEC)
		if efd < 0 {
			return nil, errors.Wrap(err, "perf_event_open error")
		}

		if _, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(efd), unix.PERF_EVENT_IOC_ENABLE, 0); err != 0 {
			return nil, errors.Wrap(err, "error enabling perf event")
		}

		if _, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(efd), unix.PERF_EVENT_IOC_SET_BPF, uintptr(progFd)); err != 0 {
			return nil, errors.Wrap(err, "error attaching bpf program to perf event")
		}

		retEventFD = append(retEventFD, internal.NewFD(uint32(efd)))
	}
	return retEventFD, nil
}

type bpfRawTracepointAttr struct {
	name    uint64
	prog_fd uint32
}

type bpfProgAttachAttr struct {
	targetFD    uint32
	attachBpfFD uint32
	attachType  uint32
	attachFlags uint32
}

const (
	_ProgAttach        = 8
	_ProgDetach        = 9
	_RawTracepointOpen = 17
)

func bpfRawTracepointOpen(progFd int, functionName string) error {
	attachPoint := []byte(functionName)
	attachPoint = append(attachPoint, 0)

	attr := bpfRawTracepointAttr{
		name:    uint64((*(*uintptr)(unsafe.Pointer(&attachPoint)))),
		prog_fd: uint32(progFd),
	}

	// Use bpf systemcall to open raw tracepoint
	_, err := internal.BPF(_RawTracepointOpen, unsafe.Pointer(&attr), unsafe.Sizeof(attr))
	if err != nil {
		return errors.Wrapf(err, "couldn't open raw tracepoint %s", functionName)
	}

	return nil
}

func bpfProgAttach(progFd int, targetFd int, attachType ebpf.AttachType) (int, error) {
	attr := bpfProgAttachAttr{
		targetFD:    uint32(targetFd),
		attachBpfFD: uint32(progFd),
		attachType:  uint32(attachType),
	}
	ptr, err := internal.BPF(_ProgAttach, unsafe.Pointer(&attr), unsafe.Sizeof(attr))
	if err != nil {
		return -1, errors.Wrapf(err, "can't attach program id %d to target fd %d", progFd, targetFd)
	}
	return int(ptr), nil
}

func bpfProgDetach(progFd int, targetFd int, attachType ebpf.AttachType) (int, error) {
	attr := bpfProgAttachAttr{
		targetFD:    uint32(targetFd),
		attachBpfFD: uint32(progFd),
		attachType:  uint32(attachType),
	}
	ptr, err := internal.BPF(_ProgDetach, unsafe.Pointer(&attr), unsafe.Sizeof(attr))
	if err != nil {
		return -1, errors.Wrapf(err, "can't detach program id %d to target fd %d", progFd, targetFd)
	}
	return int(ptr), nil
}

func sockAttach(sockFd int, progFd int) error {
	return syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, unix.SO_ATTACH_BPF, progFd)
}

func sockDetach(sockFd int, progFd int) error {
	return syscall.SetsockoptInt(sockFd, syscall.SOL_SOCKET, unix.SO_DETACH_BPF, progFd)
}
