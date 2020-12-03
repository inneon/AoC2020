package day02

import (
	"regexp"
	"strconv"
)

type passwordLine struct {
	lowerBound int
	upperBound int
	character  string
	password   string
}


func parse (line string) (*passwordLine, error) {
	regexp, err := regexp.Compile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")
	if err != nil {
		return nil, err
	}
	matches := regexp.FindAllStringSubmatch(line, -1)
	lowerBound, _ := strconv.Atoi(matches[0][1])
	upperBound, _ := strconv.Atoi(matches[0][2])
	return &passwordLine {
		lowerBound: lowerBound,
		upperBound: upperBound,
		character: matches[0][3],
		password: matches[0][4],
	}, nil
}

func checkRule (password passwordLine) bool {
	regexp, err := regexp.Compile(password.character)
	if err != nil {
		return false
	}
	matches := regexp.FindAllString(password.password, -1)
	occurences := len(matches)
	return password.lowerBound <= occurences && occurences <= password.upperBound
}

func checkSecondRule (password passwordLine) bool {
	firstMatch := password.password[password.lowerBound-1:password.lowerBound] == password.character
	secondMatch := password.password[password.upperBound-1:password.upperBound] == password.character
	return firstMatch != secondMatch
}

// NumberOfMatches returns the number of lines from the password database that match their rule
func NumberOfMatches (lines []string) int {
	result := 0
	for _, line := range lines {
		pass, err := parse(line)
		if err != nil {
			continue
		}
		if checkRule(*pass) {
			result ++
		}
	}
	return result
}

// NumberOfSecondaryMatches returns the number of lines from the password database that match their rule
func NumberOfSecondaryMatches (lines []string) int {
	result := 0
	for _, line := range lines {
		pass, err := parse(line)
		if err != nil {
			continue
		}
		if checkSecondRule(*pass) {
			result ++
		}
	}
	return result
}