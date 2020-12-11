package day11

import ( 
	// "fmt"
) 

func adjacentOccupiedSeats(plan []string, x int, y int) int {
	result := 0
	for i:=x-1; i<=x+1; i++ {
		if 0<=i && i <len(plan[0]) { 
			for j:=y-1; j<=y+1; j++ {
				if 0<=j && j<len(plan) && !(i==x && j==y) {
					if plan[j][i] == '#' {
						result++
					}
				}
			}
		}
	}
	return result
}

// OccupiedSeats returns how many seats are occupied in a steady state
func OccupiedSeats(plan []string) int {
	changed := true
	result := 0

	for changed {
		result = 0
		next := []string{}
		changed = false
		for y:=0; y<len(plan); y++ {
			row := ""
			for x:=0; x<len(plan[y]); x++ {
				adjacent := adjacentOccupiedSeats(plan, x, y)
				current := plan[y][x]
				if current == '.' {
					row = row + string(current)
				} else if adjacent == 0 {
					row = row + "#"
					if current != '#' {
						changed = true
					}
					result++
				} else if adjacent >= 4 {
					row = row + "L"
					if current != 'L' {
						changed = true
					}
				} else {
					row = row + string(current)
					if current == '#' {
						result++
					}
				}
			}
			next = append(next, row)
		}
		plan = next
	}
	return result
}