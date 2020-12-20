package day19

import ( 
	"fmt"
	"strconv"
	"strings"
) 

// GetMatchingRules gets the strings that match the rules
func GetMatchingRules(input []string) int {
	rules := make(map[int]string)
	candidates := make([]string, 0)
	mode := "rules"

	for _, line := range input {
		if line == "" {
			mode = "candidates"
		} else if mode == "rules" {
			halves := strings.Split(line, ":")
			ruleNum, _ := strconv.Atoi(halves[0])
			sanitised := strings.Replace(halves[1], "\"", "", 2)
			sanitised = strings.Trim(sanitised, " ")
			rules[ruleNum] = sanitised
		} else if mode == "candidates" {
			candidates = append(candidates, line)
		}
	}

	result := 0
	for _, candidate := range candidates {
		fmt.Println(candidate)
		matches := doesMatch(candidate, rules, 0)
		if any(matches, func(match string) bool {return match == ""}) {
			result++
		}
	}

	return result
}

func matchOption(candidate string, rules map[int]string, option string) (bool, string) {
	// fmt.Println("option", option)
	for _, token := range strings.Split(option, " ") {
		// fmt.Println("token", token, "rest", candidate)
		nextRule, err := strconv.Atoi(token)
		if err != nil {
			// fmt.Println("cannot parse")
			return false, candidate
		}
		foo := doesMatch(candidate, rules, nextRule)
		subMatch := len(foo) > 0
		// fmt.Println(subMatch, candidate)
		if !subMatch {
			return false, candidate
		}
		candidate = foo[0]
	
	}
	return true, candidate
}

func doesMatch(candidate string, rules map[int]string, rule int) ([]string) {
	if candidate == "" {
		// fmt.Println("empty candidate")
		return []string{}
	}

	if candidate[0:1] == rules[rule] {
		// fmt.Println("basic match true")
		return []string{candidate[1:]}
	}
	for _, option := range strings.Split(rules[rule], "|") {
		rest, submatch := matchOption(candidate, rules, strings.Trim(option, " "))
		if rest {
			return []string{submatch}
		}
	}

	return []string{}
}
	
func any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}