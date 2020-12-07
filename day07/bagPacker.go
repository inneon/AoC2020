package day07

import (
	"strings"
	"regexp"
	// "fmt"
)

func parseRules(rules []string) map[string][]string {
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
	parsed := parseRules(rules)
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