package day05

func getID(instruction string) int {
	result := 0
	for i:=0; i <len(instruction); i++ {
		result*=2
		char := instruction[i:i+1]
		if char == "B" || char == "R" || char == "1" {
			result++
		}
	}
	return result
}

// Range gets the instruction with the biggest and smallest id
func Range(instructions []string) (int, int) {
	max := 0
	min := 4294967295
	for _, instruction := range instructions {
		id := getID(instruction)
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}
	return max, min
}

func triangleNumber(n int) int {
	return (n * (n+1))/2
}

// Missing gets the missing ticket from the instructions
func Missing(instructions []string) int {
	sum := 0
	for _, instruction := range instructions {
		sum += getID(instruction)
	}
	max, min := Range(instructions)

	return triangleNumber(max) - sum - triangleNumber(min-1)
}