package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"./day05"
)

func main() {
	file, err := os.Open("day05/input.txt")
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
	
	fmt.Println(day05.Missing(lines))
}