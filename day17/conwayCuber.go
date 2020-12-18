package day17

import ( 
	"math"
	// "fmt"
	// "strconv"
	// "strings"
) 

type pocketDimension = map[int]map[int]map[int]bool
type bounds struct {
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func parseInput(input []string)(pocketDimension, bounds) {
	result := make(pocketDimension, 0)
	plane := make(map[int]map[int]bool, 0)
	result[0] = plane

	for y, line := range input {
		row := make(map[int]bool, 0)
		plane[y] = row

		for x, cell := range line {
			if cell == '#' {
				row[x] = true
			}
		}
	}
	
	return result, bounds{
		minX: 0,
		maxX: len(input[0])-1,
		minY: 0,
		maxY: len(input)-1,
		minZ: 0,
		maxZ: 0,
	}
}

func neigbours(pocketDimension pocketDimension, x int, y int, z int) int {
	result := 0
	for dz:=-1; dz<=1; dz++ {
		if plane, ok := pocketDimension[z+dz]; ok {
			for dy:=-1; dy<=1; dy++ {
				if row, ok := plane[y+dy]; ok {
					for dx:=-1; dx<=1; dx++ {
						if cell, ok := row[x+dx]; ok && cell {
							result++
						}
					}
				}
			}
		}
	}
	if pocketDimension[z][y][x] {
		result--
	}
	return result
}

func get(dim pocketDimension, x, y, z int) bool {
	plane, ok := dim[z]
	if !ok {
		return false
	}
	row, ok := plane[y]
	if !ok {
		return false
	}

	return row[x]
}

func put(dim pocketDimension, x, y, z int) {
	plane, ok := dim[z]
	if !ok {
		plane = make(map[int]map[int]bool)
		dim[z] = plane
	}
	row, ok := plane[y]
	if !ok {
		row = make(map[int]bool)
		plane[y] = row
	}

	row[x] = true
	// fmt.Println("put", x, y, z, dim)
}

func count(dim pocketDimension, b bounds) int {
	result := 0
	for z:=b.minZ; z<=b.maxZ; z++ {
		for y:=b.minY; y<=b.maxY; y++ {
			for x:=b.minX; x<=b.maxX; x++ {
				if (get(dim, x, y, z)) {
					result++
				}
			}
		}
	}
	return result
}

func generation(dim pocketDimension, currBounds bounds) (pocketDimension, bounds) {
	result := make(pocketDimension, 0)
	nextBounds := bounds{
		minX: math.MaxInt32,
		maxX: math.MinInt32,
		minY: math.MaxInt32,
		maxY: math.MinInt32,
		minZ: math.MaxInt32,
		maxZ: math.MinInt32,
	}

	for z:=currBounds.minZ-1; z<=currBounds.maxZ+1; z++ {
		for y:=currBounds.minY-1; y<=currBounds.maxY+1; y++ {
			for x:=currBounds.minX-1; x<=currBounds.maxX+1; x++ {
				numNeigbours := neigbours(dim, x, y, z)
				// fmt.Println(numNeigbours, x, y, z)
				if numNeigbours == 3 || 
					(numNeigbours == 2 && get(dim, x, y, z)){
					put(result, x, y, z)
					if x > nextBounds.maxX {
						nextBounds.maxX = x
					}
					if x < nextBounds.minX {
						nextBounds.minX = x
					}
					if y > nextBounds.maxY {
						nextBounds.maxY = y
					}
					if y < nextBounds.minY {
						nextBounds.minY = y
					}
					if z > nextBounds.maxZ {
						nextBounds.maxZ = z
					}
					if z < nextBounds.minZ {
						nextBounds.minZ = z
					}
				}
			}
		}
	}
	return result, nextBounds
}

// RunCube runs the initial cube for a given number of generations and then reurns the active count
func RunCube(initial []string, generations int) int {
	dim, b := parseInput(initial)

	for i:=0; i<generations; i++ {
		dim, b = generation(dim, b)
	}
	return count(dim, b)
}