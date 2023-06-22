package kernel

import (
	"errors"
	"github.com/Binject/debug/pe"
	"github.com/NeuralTeam/kernel/pkg/windows"
	"github.com/awgh/rawreader"
	"unsafe"
)

// WriteMemory writes memory to the specified address
// May cause panic if memory is not writable
func (k *Kernel) WriteMemory(buf []byte, destination uintptr) {
	for index := uint32(0); index < uint32(len(buf)); index++ {
		writePtr := unsafe.Pointer(destination + uintptr(index))
		v := (*byte)(writePtr)
		*v = buf[index]
	}
}

// MemoryId takes the exported syscall name or ordinal and gets the ID it refers to
func (k *Kernel) MemoryId(name string) (uint16, error) {
	return k.memoryId(name, 0, false)
}

// ModuleOrder returns the start address of the module located at i in the load order
func (k *Kernel) ModuleOrder(i int) (start uintptr, size uintptr, path string) {
	var utf16 *windows.Utf16
	start, size, utf16 = k.GetModuleLoadedOrder(i)
	path = utf16.String()
	return
}

// Images return a map of loaded dll paths to current process offsets
func (k *Kernel) Images() (images map[string]windows.Image) {
	images = make(map[string]windows.Image)
	start, size, path := k.ModuleOrder(0)
	images[path] = windows.Image{BaseAddr: uint64(start), Size: uint64(size)}

	p := path
	i := 1
	for {
		start, size, path = k.ModuleOrder(i)
		if path != "" {
			images[path] = windows.Image{BaseAddr: uint64(start), Size: uint64(size)}
		}
		if path == p {
			break
		}
		i++
	}
	return
}

func (k *Kernel) memoryId(name string, ord uint32, useOrd bool) (id uint16, err error) {
	start, size := k.GetDllStart()
	raw := rawreader.New(start, int(size))
	file, err := pe.NewFileFromMemory(raw)
	fileBytes, err := file.Bytes()

	exports, err := file.Exports()
	for _, export := range exports {
		if err != nil {
			return
		}
		if (useOrd && export.Ordinal == ord) ||
			export.Name == name {
			offset := rvaToOffset(file, export.VirtualAddress)
			buff := fileBytes[offset : offset+10]
			id, err = idFromRawBytes(buff)
		}
	}
	if id == 0 {
		err = errors.New("could not find syscall id")
	}
	return
}

// rvaToOffset converts an RVA value from a PE file into the file offset
func rvaToOffset(file *pe.File, rva uint32) uint32 {
	baseOffset := uint64(rva)
	for _, section := range file.Sections {
		if baseOffset > uint64(section.VirtualAddress) &&
			baseOffset < uint64(section.VirtualAddress+section.VirtualSize) {
			rva = rva - section.VirtualAddress + section.Offset
		}
	}
	return rva
}
