package day14

import ( 
	"regexp"
	"strconv"
	"strings"
) 

func applyMask (val int, mask string) int {
	for i, c := range mask {
		power := 1 << (len(mask) - i - 1)
		switch c {
			case '1': 
				if power & val == 0 {
					val += power
				}
				break
			case '0': 
				if power & val != 0 {
					val -= power
				}
				break
		}
	}
	return val
}

// RunMaskedProgram runs the program and returns the sum of memory locations
func RunMaskedProgram(program []string) int {
	memRegex, _ := regexp.Compile("^mem\\[(\\d+)\\] = (\\d+)$")
	maskRegex, _ := regexp.Compile("^mask = ([\\dX]+)$")
	mem := make(map[int]int)
	mask := "X"

	for _, line := range program {
		if strings.HasPrefix(line, "mask = "){ 
			matches := maskRegex.FindAllStringSubmatch(line, -1)
			mask = matches[0][1]
		} else if strings.HasPrefix(line, "mem[") {
			matches := memRegex.FindAllStringSubmatch(line, -1)
			loc, _ := strconv.Atoi(matches[0][1])
			val, _ := strconv.Atoi(matches[0][2])
			mem[loc] = applyMask(val, mask)
		}
	}

	result := 0
	for _, v := range mem {
		result += v
	}
	return result
}

func applyV2Mask (val int, mask string, i int) []int {
	if i == len(mask) {
		return []int{val}
	}
	result := applyV2Mask(val, mask, i+1)
	c := mask[i]
	power := 1 << (len(mask) - i - 1)
	switch c {
		case '0': 
			return result
		case '1':
			for i := range result {
				result[i] |= power
			}
			return result
		case 'X':
			lower := append([]int(nil), result...)
			upper := append([]int(nil), result...)
			for i := range result {
				lower[i] &= ^power
				upper[i] |= power
			}
			return append(lower, upper...)
	}

	return result
}

// RunV2Program v2 runs the program and returns the sum of memory locations
func RunV2Program(program []string) int {
	memRegex, _ := regexp.Compile("^mem\\[(\\d+)\\] = (\\d+)$")
	maskRegex, _ := regexp.Compile("^mask = ([\\dX]+)$")
	mem := make(map[int]int)
	mask := "X"

	for _, line := range program {
		if strings.HasPrefix(line, "mask = "){ 
			matches := maskRegex.FindAllStringSubmatch(line, -1)
			mask = matches[0][1]
		} else if strings.HasPrefix(line, "mem[") {
			matches := memRegex.FindAllStringSubmatch(line, -1)
			loc, _ := strconv.Atoi(matches[0][1])
			val, _ := strconv.Atoi(matches[0][2])
			for _, maskedLoc := range applyV2Mask(loc, mask, 0) {
				mem[maskedLoc] = val
			}
		}
	}

	result := 0
	for _, v := range mem {
		result += v
	}
	return result
}