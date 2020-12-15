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

// AlignedTimestamps gets the first timestamp when the departures are aligned with the schedule
func AlignedTimestamps(schedule string) int {
	offsets := make([]positionAndID,0)
	for i, id := range strings.Split(schedule, ",") {
		num, err := strconv.Atoi(id)
		if err == nil {
			offsets = append(offsets, positionAndID {id: num, position: i})
		}
	}
	return alignedTimestamps(offsets)
}

type positionAndID struct{
	id int
	position int
}
func alignedTimestamps(schedule []positionAndID) int {
	if len(schedule) == 1 {
		return schedule[0].position
	} 

	first := schedule[0]
	second := schedule[1]
	i:=first.position
	for ; i<math.MaxInt64; i+=first.id {
		if (i + second.position) % second.id == 0 {
			return alignedTimestamps(append([]positionAndID{{ id: first.id * second.id, position: i}}, schedule[2:]...))
		}
	}
	return 0
}