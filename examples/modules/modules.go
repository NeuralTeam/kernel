package main

import (
	"github.com/NeuralTeam/kernel"
	"github.com/NeuralTeam/kernel/pkg/dll"
	"sync"
)

type Modules struct {
	Images *sync.Map
	Kernel kernel.Kernel
}

func New() (modules *Modules, err error) {
	modules = new(Modules)
	if modules.Kernel, err = kernel.New(dll.Ntdll); err != nil {
		return
	}
	modules.Get()
	return
}

func (m *Modules) Get() *sync.Map {
	m.Images = new(sync.Map)
	for p, i := range m.Kernel.Images() {
		m.Images.Store(p, i)
	}
	return m.Images
}
