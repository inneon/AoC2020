package day15

import ( 
) 

// MemoryGame gets the number for the given turn in the game
func MemoryGame(initial []int, turn int) int {
	previous := initial[len(initial) - 1]
	for i:=len(initial); i<turn; i++ {
		next := 0
		for age:=1; age<len(initial); age++ {
			if initial[len(initial) - age - 1] == previous {
				next = age
				break
			}
		}
		previous = next
		initial = append(initial, next)
	}
	return initial[turn-1]
}