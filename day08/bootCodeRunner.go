package day08

import (
	"strings"
	"strconv"
)

// ExitCode the code that the program exited with
type ExitCode int

const (
	// Success - halted succesfully
	Success ExitCode = iota
	// Loop - went into an infinite loop
	Loop
	// OutOfBounds - attempted to access a location out of bounds
	OutOfBounds
	// Err - something unknown went wrong
	Err
)

// Run runs the instructions and returns the value in the accumulator
func Run(instrs []string) (int, ExitCode) {
	instructionCounter := 0
	accumulator := 0
	const halt = false
	visited := make(map[int]bool)

	for true {
		if visited[instructionCounter] == true {
			return accumulator, Loop
		}
		if instructionCounter == len(instrs) {
			return accumulator, Success
		}
		if instructionCounter > len(instrs) || instructionCounter < 0 {
			return -1, OutOfBounds
		}

		visited[instructionCounter] = true
		operands := strings.Split(instrs[instructionCounter], " ")
		inst := operands[0]
		arg, err := strconv.Atoi(operands[1])
		if err != nil {
			return -1, Err
		}

		if inst == "nop" {
			instructionCounter++
		} else if inst == "acc" {
			accumulator += arg
			instructionCounter++
		} else if inst == "jmp" {
			instructionCounter += arg
		}
	}

	return -1, Err
}

// Fix attempts to fix the program so it will halt succesfully
func Fix(instrs []string) int {
	for i, inst := range instrs {
		if strings.HasPrefix(inst, "acc") {
			continue
		}
		candidateFix := make([]string, len(instrs))
		copy(candidateFix, instrs)
		operands := strings.Split(instrs[i], " ")
		inst := operands[0]
		if inst == "nop" {
			candidateFix[i] = strings.Join([]string{"jmp", operands[1]}, " ")
		} else {
			candidateFix[i] = strings.Join([]string{"nop", operands[1]}, " ")
		}
		result, code := Run(candidateFix)
		if code == Success {
			return result
		}
	}

	return -1
}