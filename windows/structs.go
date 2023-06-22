package windows

type ListEntry struct {
	Flink, Blink *ListEntry
}

type LdrDataTableEntry struct {
	InLoadOrderLinks           ListEntry
	InMemoryOrderLinks         ListEntry
	InInitializationOrderLinks ListEntry
	DllBase                    *uintptr
	EntryPoint                 *uintptr
	SizeOfImage                *uintptr
	DllName                    Utf16
	BaseDllName                Utf16
	Flags                      uint32
	LoadCount                  uint16
	TlsIndex                   uint16
	HashLinks                  ListEntry
	TimeDateStamp              uint64
}

type Image struct {
	BaseAddr,
	Size uint64
}

type Utf16 struct {
	Length    uint16
	MaxLength uint16
	PwStr     *uint16
}

type Procedure struct {
	Address uintptr
}
