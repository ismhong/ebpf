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
	"time"
)

var m = &manager.Manager{
	Probes: []*manager.Probe{
		&manager.Probe{
			Section: "raw_tracepoint/timer_start",
		},
	},
}

func main() {
	fmt.Println("raw_tracepoint example")

	options := manager.Options{
		RLimit: &unix.Rlimit{
			Cur: math.MaxUint64,
			Max: math.MaxUint64,
		},
	}

	// Get reader for ebpf program
	ebpfProg, err := Asset("timers_raw_tracepoints.elf")
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

	go func() {
		bpfMap, ok, err := m.GetMap("counts")
		if err != nil || ok != true {
			logrus.Println("Can't get counts map")
		}

		for {
			entries := bpfMap.Iterate()
			var key, val uint64
			for entries.Next(&key, &val) {
				logrus.Printf("[%x] -> %d", key, val)
			}

			logrus.Println("================================")
			time.Sleep(time.Second)
		}
	}()

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
