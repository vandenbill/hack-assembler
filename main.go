package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/vandenbill/hack-assembler/parser"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		line = parser.DeleteWhiteSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	lines = parser.FristPass(lines)
	parser.SecondPass(lines)
	instructions := parser.ToInstructions(lines)

	if err := os.WriteFile("out.hack", []byte(strings.Join(instructions, "\n")), 0644); err != nil {
		log.Fatalln(err.Error())
	}
}
