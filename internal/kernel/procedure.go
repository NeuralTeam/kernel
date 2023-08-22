package kernel

import (
	"fmt"
	"github.com/NeuralTeam/kernel/pkg/windows"
	"strings"
)

func (k *Kernel) FuncPtr(name string) (ptr uint64, err error) {
	exports, err := k.File.Exports()
	for _, export := range exports {
		if strings.EqualFold(name, export.Name) {
			ptr = uint64(k.Hook.GetMemLoc()) + uint64(export.VirtualAddress)
		}
	}
	if ptr == 0 {
		err = fmt.Errorf("could not find function: %s", name)
	}
	return
}

func (k *Kernel) NewProc(name string) *windows.Procedure {
	address, _ := k.FuncPtr(name)
	return &windows.Procedure{Address: uintptr(address)}
}
