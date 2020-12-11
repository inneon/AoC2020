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

func visibleOccupiedSeats(plan []string, x int, y int) int {
	result := 0
	vectors := []struct{
		dx int
		dy int
	}{
		{-1, -1},{-1, 0},{-1, 1},
		{ 0, -1}        ,{ 0, 1},
		{ 1, -1},{ 1, 0},{ 1, 1},
	}
	for _, v := range vectors {
		cont := true
		dist := 1
		for cont {
			lookingAtX := x + (v.dx*dist)
			lookingAtY := y + (v.dy*dist)
			if 0 <= lookingAtY && lookingAtY < len(plan) &&
				0 <= lookingAtX && lookingAtX < len(plan[lookingAtY]) {
					char := plan[lookingAtY][lookingAtX]
					if char == '#' {
						result++;
						cont = false
					} else if char == 'L' {
						cont = false
					} else {
						dist ++
					}
			} else {
				cont = false
			}
		}
	}
	return result
}

// OccupiedSeats returns how many seats are occupied in a steady state
func OccupiedSeats(plan []string) int {
	return occupiedSeats(plan, true, 4)
}

// OccupiedSeatsByVisibility returns how many seats are occupied in a steady state when strict visibility is used
func OccupiedSeatsByVisibility(plan []string) int {
	return occupiedSeats(plan, false, 5)
}

func occupiedSeats (plan []string, adjacentOnly bool, limit int) int {
	changed := true
	result := 0

	for changed {
		result = 0
		next := []string{}
		changed = false
		for y:=0; y<len(plan); y++ {
			row := ""
			for x:=0; x<len(plan[y]); x++ {
				occupied := 0
				if adjacentOnly {
					occupied = adjacentOccupiedSeats(plan, x, y)
				} else {
					occupied = visibleOccupiedSeats(plan, x, y)
				}
				current := plan[y][x]
				if current == '.' {
					row = row + string(current)
				} else if occupied == 0 {
					row = row + "#"
					if current != '#' {
						changed = true
					}
					result++
				} else if occupied >= limit {
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