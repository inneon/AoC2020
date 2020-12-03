package day03

// CountTrees counts the number of trees that we would bump into
func CountTrees(grid []string, right int, down int) int {
	trees := 0
	position := 0 

	for i := 0; i < len(grid); i+=down {
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

type route struct {
	right int
	down  int
}

// CheckRoutes checks the routes for how many trees will be hit
func CheckRoutes(grid []string) int {
	result := 1

	routes := []route{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	for _, route := range routes {
		result *= CountTrees(grid, route.right, route.down)
	}

	return result
}