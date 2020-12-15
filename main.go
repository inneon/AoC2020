package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"strconv"
	"./day15"
)

func getLines(fileName string) []string {
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	return lines
}

func getNumbericLines(fileName string) []int {
	var numbers []int
	lines := getLines(fileName)
	for _, line := range lines {
		parsed, _ := strconv.Atoi(line)
        numbers = append(numbers, parsed)
	}
	return numbers
}

func main() {
	// lines := getLines("day15/input.txt")
	fmt.Println(day15.MemoryGame([]int{8,11,0,19,1,2}, 30000000))
}