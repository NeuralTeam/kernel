
// func getPeb() uintptr
TEXT ·getPeb(SB), $0-8
     MOVQ 	0x60(GS), AX
     MOVQ	AX, ret+0(FP)
     RET

// func getDllStart() (address uintptr, size uintptr)
TEXT ·getDllStart(SB), $0-16
	// All operations push values into AX
	// PEB
	MOVQ 0x60(GS), AX
	// PEB -> LDR
	MOVQ 0x18(AX),AX

	// LDR -> InMemoryOrderModuleList
	MOVQ 0x20(AX),AX

	// Flink (get next element)
	MOVQ (AX),AX

	// Flink - 0x10 -> _LDR_DATA_TABLE_ENTRY
	// _LDR_DATA_TABLE_ENTRY -> DllBase (offset 0x30)
	MOVQ 0x20(AX),CX
	MOVQ CX, start+0(FP)
	
	MOVQ 0x30(AX),CX
	MOVQ CX, size+8(FP)
		
	RET 

// func getModuleLoadedOrder(i int) (start uintptr, size uintptr)
TEXT ·getModuleLoadedOrder(SB), $0-32
	// All operations push values into AX
	// PEB
	MOVQ 0x60(GS), AX
	// PEB -> LDR
	MOVQ 0x18(AX),AX

	// LDR -> InMemoryOrderModuleList
	MOVQ 0x20(AX),AX

	XORQ R10,R10
startloop:
	CMPQ R10,i+0(FP)
	JE endloop
	// Flink (get next element)
	MOVQ (AX),AX
	INCQ R10
	JMP startloop
endloop:
	// Flink - 0x10 -> _LDR_DATA_TABLE_ENTRY
	// _LDR_DATA_TABLE_ENTRY -> DllBase (offset 0x30)
	MOVQ 0x20(AX),CX
	MOVQ CX, start+8(FP)
	
	MOVQ 0x30(AX),CX
	MOVQ CX, size+16(FP)
	MOVQ AX,CX
	ADDQ $0x38,CX
	MOVQ CX, modulepath+24(FP)
	// SYSCALL
	RET 

// func getModuleLoadedOrderPtr(i int) *win.LdrDataTableEntry
TEXT ·getModuleLoadedOrderPtr(SB), $0-16
	// All operations push values into AX
	// PEB
	MOVQ 0x60(GS), AX
	// PEB -> LDR
	MOVQ 0x18(AX),AX

	// LDR -> InMemoryOrderModuleList
	MOVQ 0x20(AX),AX

	XORQ R10,R10
startloop:
	CMPQ R10,i+0(FP)
	JE endloop
	// Flink (get next element)
	MOVQ (AX),AX
	INCQ R10
	JMP startloop
endloop:
	MOVQ AX,CX
	SUBQ $0x10,CX
	MOVQ CX, ret+8(FP)
	
	RET 

// Based on https://golang.org/src/runtime/sys_windows_amd64.s
#define maxargs 16
// func _syscall(id uint16, argh ...uintptr) uint32
TEXT ·_syscall(SB), $0-56
	XORQ AX,AX
	MOVW id+0(FP), AX
	PUSHQ CX
	// Put variadic size into CX
	MOVQ argh_len+16(FP),CX
	// Put variadic pointer into SI
	MOVQ argh_base+8(FP),SI
	// SetLastError(0)
	MOVQ	0x30(GS), DI
	MOVL	$0, 0x68(DI)
	// Room for args
	SUBQ	$(maxargs*8), SP
	// Fast version, do not store args on the stack
	CMPL	CX, $4
	JLE	loadregs
	// Check we have enough room for args
	CMPL	CX, $maxargs
	JLE	2(PC)
	// Not enough room -> crash
	INT	$3
	// Copy args to the stack
	MOVQ	SP, DI
	CLD
	REP; MOVSQ
	MOVQ	SP, SI
loadregs:
	SUBQ	$8, SP
	// Load first 4 args into correspondent registers
	MOVQ	0(SI), CX
	MOVQ	8(SI), DX
	MOVQ	16(SI), R8
	MOVQ	24(SI), R9
	// https://msdn.microsoft.com/en-us/library/zthk2dkh.aspx
	// Floating point arguments are passed in the XMM registers
	// Set them here in case any of the arguments are floating point values
	MOVQ	CX, X0
	MOVQ	DX, X1
	MOVQ	R8, X2
	MOVQ	R9, X3
	// MOVW id+0(FP), AX
	MOVQ CX, R10
	SYSCALL
	ADDQ	$((maxargs+1)*8), SP
	// Return result
	POPQ	CX
	MOVL	AX, errcode+32(FP)
	RET
