package kernel

import (
	"github.com/Binject/debug/pe"
	"github.com/NeuralTeam/kernel/dll"
	"github.com/NeuralTeam/kernel/hook"
	"github.com/NeuralTeam/kernel/windows/asm"
	"github.com/awgh/rawreader"
	"path/filepath"
	"strings"
)

type Kernel struct {
	hook *hook.Hook
	*asm.Asm
	file *pe.File
}

func New(dll *dll.Dll) (kernel *Kernel, err error) {
	kernel = new(Kernel)
	kernel.hook = new(hook.Hook)
	kernel.file = new(pe.File)
	kernel.file.OptionalHeader = new(pe.OptionalHeader32)
	kernel.Asm = asm.New()

	images := kernel.Images()
	for k, image := range images {
		if !strings.EqualFold(k, dll.Path) || !strings.EqualFold(dll.Name, filepath.Base(k)) {
			continue
		}
		raw := rawreader.New(uintptr(image.BaseAddr), int(image.Size))
		kernel.hook.SetMemLoc(uintptr(image.BaseAddr))
		kernel.file, err = pe.NewFileFromMemory(raw)
		break
	}
	return
}
