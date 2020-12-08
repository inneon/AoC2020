package day08

import (
	"strings"
	"strconv"
)

// Run runs the instructions and returns the value in the accumulator
func Run(instrs []string) int {
	instructionCounter := 0
	accumulator := 0
	halt := false
	visited := make(map[int]bool)

	for !halt {
		if visited[instructionCounter] == true {
			halt = true
		} else {
			visited[instructionCounter] = true
			operands := strings.Split(instrs[instructionCounter], " ")
			inst := operands[0]
			arg, _ := strconv.Atoi(operands[1])
			if inst == "nop" {
				instructionCounter++
			} else if inst == "acc" {
				accumulator += arg
				instructionCounter++
			} else if inst == "jmp" {
				instructionCounter += arg
			}
		}
	}

	return accumulator
}