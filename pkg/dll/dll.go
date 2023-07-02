package dll

const (
	Ntdll = "ntdll.dll"
)

type Dll struct {
	error
	Name, Path string
}

func New(path, name string) (dll *Dll) {
	dll = new(Dll)
	dll.Name = name
	dll.Path = path + `/` + dll.Name
	return
}
