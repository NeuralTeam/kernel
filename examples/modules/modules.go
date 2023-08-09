package main

import (
	"github.com/NeuralTeam/kernel"
	"github.com/NeuralTeam/kernel/pkg/dll"
	"golang.org/x/sys/windows"
	"sort"
)

type Modules struct {
	error
	paths  []string
	kernel kernel.Kernel
}

func New() (modules *Modules, err error) {
	modules = new(Modules)

	directory, err := windows.GetSystemDirectory()
	if err != nil {
		return
	}
	n := dll.New(directory, dll.Ntdll)

	modules.kernel, modules.error = kernel.New(n)
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
