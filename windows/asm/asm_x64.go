package asm

import (
	"errors"
	"github.com/NeuralTeam/kernel/windows"
)

type Asm struct {
	// GetPeb returns the in-memory start address of PEB while making no API calls
	GetPeb func() uintptr
	// GetDllStart returns the start address of ntdll in memory
	GetDllStart func() (address uintptr, size uintptr)
	// GetModuleLoadedOrder returns the start address of the module located at i in the load order.
	GetModuleLoadedOrder func(i int) (start uintptr, size uintptr, path *windows.Utf16)
	// GetModuleLoadedOrderPtr returns a pointer to the LDR data table entry
	GetModuleLoadedOrderPtr func(i int) *windows.LdrDataTableEntry
	// Syscall calls the system function specified by ID with arguments
	Syscall func(id uint16, argh ...uintptr) (err error)
}

func New() *Asm {
	asm := new(Asm)
	asm.GetPeb = getPeb
	asm.GetDllStart = getDllStart
	asm.GetModuleLoadedOrder = getModuleLoadedOrder
	asm.GetModuleLoadedOrderPtr = getModuleLoadedOrderPtr
	asm.Syscall = syscall
	return asm
}

func syscall(id uint16, argh ...uintptr) (err error) {
	result := _syscall(id, argh...)
	if result != 0 {
		err = errors.New("non-zero return from syscall")
	}
	return err
}

func getPeb() uintptr
func getDllStart() (address uintptr, size uintptr)
func getModuleLoadedOrder(i int) (start uintptr, size uintptr, path *windows.Utf16)
func getModuleLoadedOrderPtr(i int) *windows.LdrDataTableEntry
func _syscall(id uint16, argh ...uintptr) uint32
