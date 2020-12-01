package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"strconv"
	"./day01"
)

func main() {
	file, err := os.Open("day01/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
        lines = append(lines, num)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	fmt.Println(day01.FindPair(lines, 2020))
}