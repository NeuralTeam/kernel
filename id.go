package kernel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/NeuralTeam/kernel/hook"
)

func (k *Kernel) Id(name string) (id uint16, err error) {
	id, err = k.id(name, 0, false, true)
	return
}

func (k *Kernel) IdOrdinal(ordinal uint32) (id uint16, err error) {
	id, err = k.id("", ordinal, true, true)
	return
}

func (k *Kernel) id(name string, ord uint32, useOrd, useNeighbor bool) (id uint16, err error) {
	exports, err := k.file.Exports()
	if err != nil {
		return
	}
	fileBytes, err := k.file.Bytes()
	for _, export := range exports {
		if err != nil {
			return
		}
		if !(useOrd && export.Ordinal == ord) &&
			export.Name != name {
			return
		}
		offset := rvaToOffset(k.file, export.VirtualAddress)
		buffer := fileBytes[offset : offset+10]
		id, err = idFromRawBytes(buffer)

		if !errors.As(err, &k.hook) && !useNeighbor {
			return
		}
		start, size := k.GetDllStart()
		distanceNeighbor := 0
		for i := uintptr(offset); i < start+size; i += 1 {
			if fileBytes[i] == byte('\x0f') &&
				fileBytes[i+1] == byte('\x05') &&
				fileBytes[i+2] == byte('\xc3') {
				distanceNeighbor++

				id, err = idFromRawBytes(fileBytes[i+14 : i+14+8])
				if !errors.As(err, &k.hook) {
					id -= uint16(distanceNeighbor)
					return
				}
			}
		}
		distanceNeighbor = 1
		for i := uintptr(offset) - 1; i > 0; i -= 1 {
			if fileBytes[i] == byte('\x0f') &&
				fileBytes[i+1] == byte('\x05') &&
				fileBytes[i+2] == byte('\xc3') {
				distanceNeighbor++
				id, err = idFromRawBytes(fileBytes[i+14 : i+14+8])
				if !errors.As(err, &k.hook) {
					id += uint16(distanceNeighbor) - 1
					return
				}
			}
		}
	}
	err = errors.New("could not find syscall id")
	return
}

// idFromRawBytes takes a byte slice and determines if there is a SysID in the expected location
func idFromRawBytes(b []byte) (id uint16, err error) {
	h := new(hook.Hook)
	h.Bytes(b)
	if !bytes.HasPrefix(b, h.Start()) {
		err = h
	}
	id = binary.LittleEndian.Uint16(b[4:8])
	return
}
