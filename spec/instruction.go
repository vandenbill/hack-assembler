package spec

// C Instruction - 111accccccdddjjj
var Dest = map[string]string{
	"null": "000", // The value is not stored anywhere
	"M":    "001", // Memory[A] (memory register addressed by A)
	"D":    "010", // D register
	"MD":   "011", // Memory[A] and D register
	"A":    "100", // A register
	"AM":   "101", // A register and Memory[A]
	"AD":   "110", // A register and D register
	"AMD":  "111", // A register, Memory[A], and D register
}

var Comp = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

var Jump = map[string]string{
	"null": "000", // no jump
	"JGT":  "001", // if comp > 0 jump
	"JEQ":  "010", // if comp = 0 jump
	"JGE":  "011", // if comp ≥ 0 jump
	"JLT":  "100", // if comp < 0 jump
	"JNE":  "101", // if comp ≠ 0 jump
	"JLE":  "110", // if comp ≤ 0 jump
	"JMP":  "111", // unconditional jump
}
