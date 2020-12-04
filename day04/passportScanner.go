package day04

import (
	"strconv"
	"strings"
	"regexp"
)

func validBirthYear(entry string) bool {
	num, err := strconv.Atoi(entry)
	if err != nil {
		return false
	}
	return 1920 <= num && num <= 2002
}

func validIssueYear(entry string) bool {
	num, err := strconv.Atoi(entry)
	if err != nil {
		return false
	}
	return 2010 <= num && num <= 2020
}

func validExpirationYear(entry string) bool {
	num, err := strconv.Atoi(entry)
	if err != nil {
		return false
	}
	return 2020 <= num && num <= 2030
}

func validHeight(entry string) bool {
	regex, _ := regexp.Compile("^(?:(\\d{2})in)$")
	matches := regex.FindAllStringSubmatch(entry, -1)
	if len(matches) == 1 {
		height, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return false
		}
		return 59 <= height && height <= 76
	}

	regex, _ = regexp.Compile("^(?:(\\d{3})cm)$")
	matches = regex.FindAllStringSubmatch(entry, -1)
	if len(matches) == 1 {
		height, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return false
		}
		return 150 <= height && height <= 193
	}
	return false
}

func validHairColour(entry string) bool {
	matched, err := regexp.MatchString("^#[a-f0-9]{6}$", entry)
	if err != nil {
		return false
	}
	return matched
}

func validEyeColour(entry string) bool {
	return entry == "amb" || entry == "blu" || entry == "brn" || entry == "gry" || entry == "grn" || entry == "hzl" || entry == "oth"
}

func validPassportNumber(entry string) bool {
	matched, err := regexp.MatchString("^[0-9]{9}$", entry)
	if err != nil {
		return false
	}
	return matched 
}

func valid(candidate map[string]string, strict bool) bool {
	if strict {
		return validBirthYear(candidate["byr"]) &&
			validIssueYear(candidate["iyr"]) &&
			validExpirationYear(candidate["eyr"]) &&
			validHeight(candidate["hgt"]) &&
			validHairColour(candidate["hcl"]) &&
			validEyeColour(candidate["ecl"]) &&
			validPassportNumber(candidate["pid"])

	}
	return (candidate["byr"] != "" && candidate["iyr"] != "" && candidate["eyr"] != "" && candidate["hgt"] != "" && 
		candidate["hcl"] != "" && candidate["ecl"] != "" && candidate["pid"] != "")
}

func validPassports(lines []string, strict bool) int {
	result := 0
	partial := make(map[string]string)
	for _, line := range lines {
		if (line == "") {
			if valid(partial, strict) {
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
	if valid(partial, strict) {
		result++
	}
	return result
}

// ValidPassports returns how many passports are valid
func ValidPassports(lines []string) int {
	return validPassports(lines, false)
}

// ValidPassportsStrict returns how many passports are valid using strict validation
func ValidPassportsStrict(lines []string) int {
	return validPassports(lines, true)
}