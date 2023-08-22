package kernel

import (
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

	images := kernel.Images()
	for k, image := range images {
		if strings.EqualFold(k, dll.Path()) || strings.EqualFold(dll.String(), filepath.Base(k)) {
			raw := rawreader.New(uintptr(image.BaseAddr), int(image.Size))
			kernel.Hook.SetMemLoc(uintptr(image.BaseAddr))
			kernel.File, err = pe.NewFileFromMemory(raw)
			break
		}
	}
	return
}
