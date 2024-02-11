package spec

// pre defined A Instruction
var SymbolTable = map[string]string{
	"R0":     "0", // address 0..15 on memory is virtual register
	"R1":     "1",
	"R2":     "2",
	"R3":     "3",
	"R4":     "4",
	"R5":     "5",
	"R6":     "6",
	"R7":     "7",
	"R8":     "8",
	"R9":     "9",
	"R10":    "10",
	"R11":    "11",
	"R12":    "12",
	"R13":    "13",
	"R14":    "14",
	"R15":    "15",
	"SCREEN": "16384", // screen memory map
	"KBD":    "24576", // keyboard memory map
	"SP":     "0",     // stack pointer
	"LCL":    "1",
	"ARG":    "2",
	"THIS":   "3",
	"THAT":   "4",
}
