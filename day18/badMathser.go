package day18

import ( 
	"strconv"
	"strings"
) 

// DoHomework does the homework for the kid on the plane
func DoHomework(lines []string) int {
	result := 0
	for _, line := range lines {
		lineResult := interpretRightToLeft(line)
		result += lineResult
	}
	return result
}

func interpretRightToLeft (line string) int {
	tokens := strings.Split(line, " ")
	result, _, _ := runCalculations(tokens)
	return result
}
func runCalculations (tokens []string) (int, []string, int) {
	operand := tokens[0]
	result := 0
	pops := 0
	tokens = tokens[1:]
	if strings.HasPrefix(operand, "(") {
		result, tokens, pops = runCalculations(append([]string{operand[1:]}, tokens...))
	} else {
		result, _ = strconv.Atoi(operand)
	}
	if pops > 0 {
		return result, tokens, pops-1
	}

	for len(tokens) > 0 {
		operator := tokens[0]
		operand := tokens[1]
		parsedOperand := 0
		tokens = tokens[2:]
		pops := 0
		if strings.HasSuffix(operand, ")") {
			pops = strings.Count(operand, ")")
			operand = operand[:len(operand) - pops]
		}
		if strings.HasPrefix(operand, "(") {
			parsedOperand, tokens, pops = runCalculations(append([]string{operand[1:]}, tokens...))
		} else {
			parsedOperand, _ = strconv.Atoi(operand)
		}
		switch operator {
		case "+":
			result += parsedOperand
			break 
		case "-":
			result -= parsedOperand
			break 
		case "*":
			result *= parsedOperand
			break 
		case "/":
			result /= parsedOperand
			break 
		}
		if pops > 0 {
			return result, tokens, pops-1
		}
	}
	
	return result, []string{}, 0
}

