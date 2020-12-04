package day04

import (
	"strings"
)

func valid(candidate map[string]string) bool {
	return (candidate["byr"] != "" && candidate["iyr"] != "" && candidate["eyr"] != "" && candidate["hgt"] != "" && candidate["hcl"] != "" && candidate["ecl"] != "" && candidate["pid"] != "")
}

// ValidPassports returns how many passports are valid
func ValidPassports(lines []string) int {
	result := 0
	partial := make(map[string]string)
	for _, line := range lines {
		if (line == "") {
			if valid(partial) {
				result ++
			}
			partial = make(map[string]string)
		} else {
			fragments := strings.Split(line, " ")
			for _, fragment := range fragments {
				kvp := strings.Split(fragment, ":")
				partial[kvp[0]] = kvp[1]
			}
		}
	}
	if valid(partial) {
		result++
	}
	return result
}