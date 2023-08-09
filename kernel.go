package kernel

import (
	"github.com/NeuralTeam/kernel/internal/kernel"
	"github.com/NeuralTeam/kernel/pkg/dll"
	"github.com/NeuralTeam/kernel/pkg/windows"
)

type Kernel interface {
	Id(name string) (id uint16, err error)
	IdOrdinal(ordinal uint32) (id uint16, err error)

	MemoryId(name string) (uint16, error)
	ModuleOrder(i int) (start uintptr, size uintptr, path string)
	Images() (images map[string]windows.Image)
	WriteMemory(buf []byte, destination uintptr)

	FuncPtr(name string) (ptr uint64, err error)
	NewProc(name string) *windows.Procedure
}

func New(dll *dll.Dll) (Kernel, error) {
	return kernel.New(dll)
}
