package dll

import (
	"golang.org/x/sys/windows"
)

const (
	Ntdll = "ntdll.dll"
)

type Dll struct {
	Name, Path string
	error
}

func New(name string) (dll *Dll) {
	dll = new(Dll)
	dll.Name = name
	dll.Path, dll.error = windows.GetSystemDirectory()
	dll.Path = dll.Path + "\\" + dll.Name
	return
}
