package day10

import ( 
    "sort"
) 

// BuildAdaptorChain gets the number of small and large gaps
func BuildAdaptorChain(adaptors []int) (int, int) {
	small, large := 0, 1
	adaptors = append(adaptors,0)
	sort.Ints(adaptors)

	for i:=0; i<len(adaptors)-1; i++ {
		diff := adaptors[i+1] - adaptors[i]
		if diff == 1 {
			small++
		} else if diff == 3 {
			large++
		}
	}
	return small, large
}
