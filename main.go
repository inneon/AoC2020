package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"./day02"
)

func main() {
	file, err := os.Open("day02/input.txt")
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
	
	fmt.Println(day02.NumberOfSecondaryMatches(lines))
}