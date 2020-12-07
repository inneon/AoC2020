package day07

import (
	"strings"
	"regexp"
	"strconv"
)

func parseRulesOuter(rules []string) map[string][]string {
	parsed := make(map[string][]string)
	matchBagsAndQuantity, _ := regexp.Compile("(\\d+) ([a-z]+ [a-z]+) bags?")
	matchBags, _ := regexp.Compile("([a-z]+ [a-z]+) bags?")
	for _, rule := range rules {
		halves := strings.Split(rule, " contain ")
		contents := halves[1]
		if contents != "no other bags." {
			for _, bagType := range strings.Split(contents, ", ") {
				match := matchBagsAndQuantity.FindAllStringSubmatch(bagType, -1)
				partial := parsed[match[0][2]]
				if partial == nil {
					partial = []string{}
				}
				parsed[match[0][2]] = append(partial, matchBags.FindAllStringSubmatch(halves[0], -1)[0][1])
			}
		}
	}
	return parsed
}

// CountOutermost calculates how many bags can eventually contain the gold bag
func CountOutermost(rules []string) int {
	result := make(map[string]bool)
	parsed := parseRulesOuter(rules)
	queue := []string {"shiny gold"}
	for len(queue) > 0 {
		containers := parsed[queue[0]]
		for _, container := range containers {
			if result[container] == false {
				result[container] = true
				queue = append(queue, container)
			}
		}
		queue = queue[1:]
	}
 	return len(result)
}

// Contents is a content item of a bag
type Contents struct {
	name string
	number int
}

func parseRulesInner(rules []string) map[string][]Contents {
	parsed := make(map[string][]Contents)
	matchBagsAndQuantity, _ := regexp.Compile("(\\d+) ([a-z]+ [a-z]+) bags?")
	matchBags, _ := regexp.Compile("([a-z]+ [a-z]+) bags?")
	for _, rule := range rules {
		halves := strings.Split(rule, " contain ")
		from := matchBags.FindAllStringSubmatch(halves[0], -1)[0][1]
		contents := halves[1]
		if contents != "no other bags." {
			for _, bagType := range strings.Split(contents, ", ") {
				match := matchBagsAndQuantity.FindAllStringSubmatch(bagType, -1)
				partial := parsed[from]
				if partial == nil {
					partial = make([]Contents, 0)
				}
				num, _ := strconv.Atoi(match[0][1])
				parsed[from] = append(partial, Contents{
					name: match[0][2],
					number: num,
				})
			}
		}
	}
	return parsed
}

func countInside(rules map[string][]Contents, current string, multiplier int) int {
	if rules[current] == nil {
		return 1
	}
	result := 0
	for _, rule := range rules[current]{
		result += countInside(rules, rule.name, multiplier) * rule.number
	}
	result++
	return result
} 

// CountInside calculates how many bags will go inside the gold bag
func CountInside(rules []string) int {
	return countInside(parseRulesInner(rules), "shiny gold", 1) -1
}