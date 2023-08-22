package dll

import (
	"golang.org/x/sys/windows"
	"os"
)

var (
	SystemDirectory, _ = windows.GetSystemDirectory()
)

type Dll int

const (
	Ntdll Dll = iota
)

func (d Dll) String() string {
	return [...]string{
		`ntdll.dll`,
	}[d]
}

func (d Dll) Path() string {
	SystemDirectory += string(os.PathSeparator)
	return [...]string{
		SystemDirectory + Ntdll.String(),
	}[d]
}

// New example:
//
//	dll.Dll(0).New(`path`, `name`)
func (d Dll) New(path, name string) string {
	path += string(os.PathSeparator)
	return path + name
}
