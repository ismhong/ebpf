package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"math"

	"github.com/ismhong/ebpf"
	"github.com/ismhong/ebpf/manager"
)

var m = &manager.Manager{}

func main() {
	options := manager.Options{
		MapSpecEditors: map[string]manager.MapSpecEditor{
			"cache": manager.MapSpecEditor{
				Type: ebpf.LRUHash,
				MaxEntries: 1000000,
				EditorFlag: manager.EditMaxEntries | manager.EditType,
			},
		},
		RLimit: &unix.Rlimit{
			Cur: math.MaxUint64,
			Max: math.MaxUint64,
		},
	}

	// Initialize the manager
	if err := m.InitWithOptions(recoverAssets(), options); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("successfully loaded, checkout the parameters of the map \"cache\" using bpftool")

	wait()

	// Close the manager
	if err := m.Stop(manager.CleanAll); err != nil {
		logrus.Fatal(err)
	}
}
