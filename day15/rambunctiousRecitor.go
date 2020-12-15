package day15

import ( 
	// "fmt"
) 

// MemoryGame gets the number for the given turn in the game
func MemoryGame(initial []int, turn int) int {
	previous := initial[len(initial) - 1]
	gameHistory := make(map[int]*int)
	for i:=0; i<len(initial) - 1; i++ {
		turn := i
		gameHistory[initial[i]] = &turn
	}
	for i:=len(initial); i<turn; i++ {
		next := 0
		ageOfLast := gameHistory[previous]
		if ageOfLast != nil {
			next = i - (*ageOfLast) - 1
		}
		previousTurn := i-1
		gameHistory[previous] = &previousTurn
		previous = next
		initial = append(initial, next)
	}
	return initial[turn-1]
}