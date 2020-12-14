package day13

import ( 
	"strconv"
	"math"
	"strings"
) 

// EarliestBus gets the ealiest bus after the departing timestamp
func EarliestBus (departing int, schedule string) int {
	earliestTime, earliestID := math.MaxInt32, 0
	for _, id := range strings.Split(schedule, ",") {
		num, err := strconv.Atoi(id)
		if err == nil {
			waiting := num - (departing % num)
			if waiting < earliestTime {
				earliestID = num
				earliestTime = waiting
			}
		}
	}
	return earliestID * earliestTime
}