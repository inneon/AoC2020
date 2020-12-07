package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"./day07"
)

func main() {
	file, err := os.Open("day07/input.txt")
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
	
	fmt.Println(day07.CountOutermost(lines))
}