package day03

// CountTrees counts the number of trees that we would bump into
func CountTrees(grid []string) int {
	trees := 0
	position := 0 
	right := 3

	for i := 0; i < len(grid); i++ {
		current := grid[i][position:position+1]
		if current == "#" {
			trees++
		}
		position+=right
		if position >= len(grid[i]) {
			position-=len(grid[i])
		}
	}

	return trees
}