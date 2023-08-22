package asm

import (
	"errors"
	"github.com/NeuralTeam/kernel/pkg/windows"
)

type Asm64 interface {
	// GetPeb returns the in-memory start address of PEB while making no API calls
	GetPeb() uintptr
	// GetNtdllStart returns the start address of ntdll in memory
	GetNtdllStart() (start uintptr, size uintptr)
	// GetModuleLoadedOrder returns the start address of the module located at i in the load order.
	GetModuleLoadedOrder(i int) (start uintptr, size uintptr, path *windows.Utf16)
	// GetModuleLoadedOrderPtr returns a pointer to the LDR data table entry
	GetModuleLoadedOrderPtr(i int) *windows.LdrDataTableEntry
	// Syscall calls the system function specified by ID with arguments
	Syscall(id uint16, args ...uintptr) (error uint32, err error)
}

type asm64 uintptr

func (a asm64) Syscall(id uint16, args ...uintptr) (error uint32, err error) {
	if error = a.syscall(id, args...); error != 0 {
		err = errors.New("non-zero return from syscall")
	}
	return
}

func (a asm64) GetPeb() uintptr {
	return getPeb()
}

func (a asm64) GetNtdllStart() (start uintptr, size uintptr) {
	return getNtdllStart()
}

func (a asm64) GetModuleLoadedOrder(i int) (start uintptr, size uintptr, path *windows.Utf16) {
	return getModuleLoadedOrder(i)
}

func (a asm64) GetModuleLoadedOrderPtr(i int) *windows.LdrDataTableEntry {
	return getModuleLoadedOrderPtr(i)
}

func (a asm64) syscall(id uint16, args ...uintptr) (error uint32) {
	return syscall(id, args...)
}

func getPeb() uintptr
func getNtdllStart() (start uintptr, size uintptr)
func getModuleLoadedOrder(i int) (start uintptr, size uintptr, path *windows.Utf16)
func getModuleLoadedOrderPtr(i int) *windows.LdrDataTableEntry
func syscall(id uint16, args ...uintptr) (error uint32)
