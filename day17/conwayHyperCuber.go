package day17

import ( 
	"math"
) 

type pocketHyperDimension = map[int]map[int]map[int]map[int]bool
type hyperBounds struct {
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
	minW int
	maxW int
}

func parseHyperInput(input []string)(pocketHyperDimension, hyperBounds) {
	result := make(pocketHyperDimension, 0)
	cube := make(map[int]map[int]map[int]bool, 0)
	result[0] = cube
	plane := make(map[int]map[int]bool)
	cube[0] = plane
	
	for y, line := range input {
		row := make(map[int]bool, 0)
		plane[y] = row

		for x, cell := range line {
			if cell == '#' {
				row[x] = true
			}
		}
	}
	
	return result, hyperBounds{
		minX: 0,
		maxX: len(input[0])-1,
		minY: 0,
		maxY: len(input)-1,
		minZ: 0,
		maxZ: 0,
		minW: 0,
		maxW: 0,
	}
}

func hyperNeigbours(pocketHyperDimension pocketHyperDimension, x, y, z, w int) int {
	result := 0
	for dw:=-1; dw<=1; dw++ {
		if cube, ok := pocketHyperDimension[w+dw]; ok {
			for dz:=-1; dz<=1; dz++ {
				if plane, ok := cube[z+dz]; ok {
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
		}
	}
	if pocketHyperDimension[w][z][y][x] {
		result--
	}
	return result
}

func hyperGet(dim pocketHyperDimension, x, y, z, w int) bool {
	cube, ok := dim[w] 
	if !ok {
		return false
	}
	plane, ok := cube[z]
	if !ok {
		return false
	}
	row, ok := plane[y]
	if !ok {
		return false
	}

	return row[x]
}

func hyperPut(dim pocketHyperDimension, x, y, z, w int) {
	cube, ok := dim[w] 
	if !ok {
		cube = make(map[int]map[int]map[int]bool)
		dim[w] = cube
	}
	plane, ok := cube[z]
	if !ok {
		plane = make(map[int]map[int]bool)
		cube[z] = plane
	}
	row, ok := plane[y]
	if !ok {
		row = make(map[int]bool)
		plane[y] = row
	}

	row[x] = true
}

func hyperCount(dim pocketHyperDimension, b hyperBounds) int {
	result := 0
	for w:=b.minW; w<=b.maxW; w++ {
		for z:=b.minZ; z<=b.maxZ; z++ {
			for y:=b.minY; y<=b.maxY; y++ {
				for x:=b.minX; x<=b.maxX; x++ {
					if (hyperGet(dim, x, y, z, w)) {
						result++
					}
				}
			}
		}
	}
	return result
}

func hyperGeneration(dim pocketHyperDimension, currhyperBounds hyperBounds) (pocketHyperDimension, hyperBounds) {
	result := make(pocketHyperDimension, 0)
	nexthyperBounds := hyperBounds{
		minX: math.MaxInt32,
		maxX: math.MinInt32,
		minY: math.MaxInt32,
		maxY: math.MinInt32,
		minZ: math.MaxInt32,
		maxZ: math.MinInt32,
		minW: math.MaxInt32,
		maxW: math.MinInt32,
	}

	for w:=currhyperBounds.minW-1; w<=currhyperBounds.maxW+1; w++ {
		for z:=currhyperBounds.minZ-1; z<=currhyperBounds.maxZ+1; z++ {
			for y:=currhyperBounds.minY-1; y<=currhyperBounds.maxY+1; y++ {
				for x:=currhyperBounds.minX-1; x<=currhyperBounds.maxX+1; x++ {
					numhyperNeigbours := hyperNeigbours(dim, x, y, z, w)
					if numhyperNeigbours == 3 || 
						(numhyperNeigbours == 2 && hyperGet(dim, x, y, z, w)){
						hyperPut(result, x, y, z, w)
						if x > nexthyperBounds.maxX {
							nexthyperBounds.maxX = x
						}
						if x < nexthyperBounds.minX {
							nexthyperBounds.minX = x
						}
						if y > nexthyperBounds.maxY {
							nexthyperBounds.maxY = y
						}
						if y < nexthyperBounds.minY {
							nexthyperBounds.minY = y
						}
						if z > nexthyperBounds.maxZ {
							nexthyperBounds.maxZ = z
						}
						if z < nexthyperBounds.minZ {
							nexthyperBounds.minZ = z
						}
						if w > nexthyperBounds.maxW {
							nexthyperBounds.maxW = w
						}
						if w < nexthyperBounds.minW {
							nexthyperBounds.minW = w
						}
					}
				}
			}
		}
	}
	return result, nexthyperBounds
}

// RunHyperCube runs the initial cube for a given number of generations and then reurns the active hyperCount
func RunHyperCube(initial []string, generations int) int {
	dim, b := parseHyperInput(initial)
	
	for i:=0; i<generations; i++ {
		dim, b = hyperGeneration(dim, b)
	}
	return hyperCount(dim, b)
}