package hook

import "fmt"

type Hook struct {
	memLoc uintptr
	bytes  []byte
	error
}

// Start are the bytes expected to be seen at the start of the function
func (h *Hook) Start() []byte {
	// mov r10, rcx ;(4c 8b d1)
	// mov eax, sysid ;(b8 sysid)
	return []byte{0x4c, 0x8b, 0xd1, 0xb8}
}

// Bytes returned when trying to extract the ID
func (h *Hook) Bytes(b []byte) []byte {
	if b != nil {
		h.bytes = b
	}
	return h.bytes
}

func (h *Hook) GetMemLoc() uintptr {
	return h.memLoc
}

func (h *Hook) SetMemLoc(u uintptr) uintptr {
	h.memLoc = u
	return h.GetMemLoc()
}

func (h *Hook) Error() string {
	return fmt.Sprintf("may be hooked: wanted %x got %x", h.Start(), h.Bytes(nil))
}
