package main

import (
	"github.com/NeuralTeam/kernel"
	"github.com/NeuralTeam/kernel/dll"
	"sort"
)

type Modules struct {
	error
	paths  []string
	kernel *kernel.Kernel
}

func New() (modules *Modules, err error) {
	modules = new(Modules)
	modules.kernel, modules.error = kernel.New(dll.New(dll.Ntdll))
	if err = modules.error; err != nil {
		return
	}
	modules.paths = modules.Get()
	return
}

func (m *Modules) Get() (paths []string) {
	images := m.kernel.Images()
	for p := range images {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	return
}
