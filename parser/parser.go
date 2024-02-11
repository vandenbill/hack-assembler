package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/vandenbill/hack-assembler/spec"
)

// delete comment and whitespace
func DeleteWhiteSpace(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	idx := strings.IndexAny(line, "//")
	if idx != -1 {
		return line[:idx]
	}
	return line
}

// extract label and pre defined variable
func FristPass(lines []string) []string {
	res := make([]string, 0, 100)
	for i, line := range lines {
		if rune(line[0]) == '@' {
			symbol := strings.Split(line, "@")[1]
			if spec.SymbolTable[symbol] != "" {
				res = append(res, fmt.Sprintf("@%s", spec.SymbolTable[symbol]))
			} else {
				res = append(res, line)
			}
		} else if rune(line[0]) == '(' {
			label := line[1 : len(line)-1] // (LABEL) get only LABEL
			if spec.SymbolTable[label] == "" {
				spec.SymbolTable[label] = fmt.Sprint(i + 1)
			}
		} else { // C Instruction
			res = append(res, line)
		}
	}
	return res
}

func SecondPass(lines []string) {
	c := 16 // from spec, the variable is located at 16 onward
	for i, line := range lines {
		if rune(line[0]) == '@' {
			variable := strings.Split(line, "@")[1]
			if unicode.IsLetter(rune(variable[0])) { // check if it already a constant
				symbol := spec.SymbolTable[variable]
				if symbol != "" {
					lines[i] = fmt.Sprintf("@%s", symbol)
					continue
				}

				address := fmt.Sprint(c) // new memory address
				spec.SymbolTable[variable] = address
				lines[i] = strings.Replace(line, variable, address, 1)
				c++
			}
		}
	}
}

func ToInstructions(lines []string) (instructions []string) {
	instructions = make([]string, 0, 100)
	for _, v := range lines {
		if rune(v[0]) == '@' {
			i := aInstruction(v)
			instructions = append(instructions, i)
		} else {
			i := cInstruction(v)
			instructions = append(instructions, i)
		}
	}
	return
}

func aInstruction(line string) string {
	address := strings.Split(line, "@")[1]
	aInt, err := strconv.Atoi(address)
	if err != nil {
		log.Fatalln(err)
	}

	addrBit := strconv.FormatInt(int64(aInt), 2)

	diff := 16 - 1 - len(addrBit)

	return fmt.Sprintf("%d%s%s", 0, strings.Repeat("0", diff), addrBit)
}

func cInstruction(line string) string {
	raw := strings.Split(line, ";")
	jump := "null"
	if len(raw) > 1 {
		jump = raw[1]
	}

	raw = strings.Split(raw[0], "=")

	dest := raw[0]
	comp := ""

	if len(raw) > 1 {
		comp = raw[1]
	}

	if len(raw) == 1 {
		dest = "null"
		comp = raw[0]
	}

	// 111accccccdddjjj
	return fmt.Sprintf("111%s%s%s",
		spec.Comp[comp], spec.Dest[dest], spec.Jump[jump])
}
