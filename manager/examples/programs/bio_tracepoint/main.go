package main

import (
	"bytes"
	"fmt"
	"github.com/ismhong/ebpf/manager"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"math"
	"os"
	"os/signal"
	//"time"
)

var m = &manager.Manager{
	Probes: []*manager.Probe{
		&manager.Probe{
			Section: "tracepoint/block/block_rq_issue",
		},
		&manager.Probe{
			Section: "tracepoint/block/block_rq_complete",
		},
	},
}

func main() {
	fmt.Println("bio exporter")

	options := manager.Options{
		RLimit: &unix.Rlimit{
			Cur: math.MaxUint64,
			Max: math.MaxUint64,
		},
	}

	// Get reader for ebpf program
	ebpfProg, err := Asset("bio_tracepoint.elf")
	if err != nil {
		logrus.Fatal(err)
	}

	// Initialize the manager
	if err := m.InitWithOptions(bytes.NewReader(ebpfProg), options); err != nil {
		logrus.Fatal(err)
	}

	// Start the manager
	if err := m.Start(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("successfully started")
	logrus.Println("=> Cmd+C to exit")

	//go func() {
	//cpuInstructionsMap, ok, err := m.GetMap("instructions")
	//if err != nil || ok != true {
	//logrus.Println("Can't get instructions map")
	//}

	//cpuCyclesMap, ok, err := m.GetMap("cycles")
	//if err != nil || ok != true {
	//logrus.Println("Can't get cycles map")
	//}

	//for {
	//time.Sleep(time.Second)

	//var instructionCount, cycleCount uint64
	//cpuInstructionsMap.Lookup(uint32(0), &instructionCount)
	//cpuCyclesMap.Lookup(uint32(0), &cycleCount)

	//logrus.Printf("instructions/cycles [%8d/%8d]\n", instructionCount, cycleCount)
	//}
	//}()

	wait()

	// Close the manager
	if err := m.Stop(manager.CleanAll); err != nil {
		logrus.Fatal(err)
	}
}

// wait - Waits until an interrupt or kill signal is sent
func wait() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println()
}
