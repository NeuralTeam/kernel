package kernel

import (
	"fmt"
	"github.com/Binject/debug/pe"
	"github.com/NeuralTeam/kernel/internal/hook"
	"github.com/NeuralTeam/kernel/pkg/dll"
	"github.com/awgh/rawreader"
	"path/filepath"
	"strings"
)

type Kernel struct {
	Hook *hook.Hook
	File *pe.File
}

func New(dll dll.Dll) (kernel *Kernel, err error) {
	kernel = new(Kernel)
	kernel.Hook = new(hook.Hook)
	kernel.File = new(pe.File)
	kernel.File.OptionalHeader = new(pe.OptionalHeader32)

	for path, image := range kernel.Images() {
		if strings.EqualFold(path, dll.Path()) || strings.EqualFold(dll.String(), filepath.Base(path)) {
			raw := rawreader.New(uintptr(image.BaseAddr), int(image.Size))
			kernel.Hook.SetMemLoc(uintptr(image.BaseAddr))
			kernel.File, err = pe.NewFileFromMemory(raw)
			return
		}
	}
	err = fmt.Errorf("module not found: %v", dll)
	return
}
