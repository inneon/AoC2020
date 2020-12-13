package day12

import ( 
	"strconv"
	"math"
) 

// FollowDirections follows the directions and returns the location
func FollowDirections(directions []string)(int, int) {
	x, y := 0, 0
	facing := 0
	forwardsMovements := make(map[int]func(int))
	// East
	forwardsMovements[0] = func(value int) {
		x+=value
	}
	// South
	forwardsMovements[1] = func(value int) {
		y-=value
	}
	// West
	forwardsMovements[2] = func(value int) {
		x-=value
	}
	// North
	forwardsMovements[3] = func(value int) {
		y+=value
	}

	for _, dir := range directions {
		action := dir[0]
		value, _ := strconv.Atoi(dir[1:len(dir)])

		switch action {
			case 'N': {
				y += value
				break
			} 
			case 'S': {
				y -= value
				break
			}
			case 'E': {
				x += value
				break
			} 
			case 'W': {
				x -= value
				break
			}
			case 'F': {
				forwardsMovements[facing](value)
				break
			}
			case 'R': {
				facing = (facing + int(value/90)) % 4
				break	
			}
			case 'L': {
				facing = ((facing - int(value/90)) % 4 + 4) %4
				break	
			}
		}
	}
	return x, y
}

func sin (angle int) int {
	rads:= float64(angle) * math.Pi /180
	return int(math.Sin(rads))
}
func cos (angle int) int {
	rads:= float64(angle) * math.Pi /180
	return int(math.Cos(rads))
}

// FollowWaypoint follows the directions by following the waypoint
func FollowWaypoint(directions []string)(int, int) {
	x, y := 0, 0
	waypointx, waypointy := 10, 1

	for _, dir := range directions {
		action := dir[0]
		value, _ := strconv.Atoi(dir[1:len(dir)])

		switch action {
			case 'F': {
				x += value * waypointx
				y += value * waypointy
				break
			} 
			case 'N': {
				waypointy += value
				break
			}
			case 'S': {
				waypointy -= value
				break
			}
			case 'E': {
				waypointx += value
				break
			}
			case 'W': {
				waypointx -= value
				break
			}
			case 'L': {
				temp := waypointx * cos(value) - waypointy * sin(value)
				waypointy = waypointy * cos(value) + waypointx * sin(value)
				waypointx = temp
				break
			}
			case 'R': {
				temp := waypointx * cos(-value) - waypointy * sin(-value)
				waypointy = waypointy * cos(-value) + waypointx * sin(-value)
				waypointx = temp
				break
			}
		}
	}

	return x, y
}