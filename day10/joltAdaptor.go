package day10

import ( 
	"sort"
	"math"
	"fmt"
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

func subChainPermutation(length int) int {
	if length == 3 {
		return 2
	} else if length == 4 {
		return 4
	} else if length == 5 {
		return 7
	}
	if length > 5 {
		fmt.Println("too long", length) // todo - or not
	}
	return 1
}

// ChainPermutations returns the number of permutations of valid chains
func ChainPermutations(adaptors []int) int {
	sort.Ints(adaptors)
	subChains := make(map[int]int)

	subChainLen := 1 
	prev := 0 
	for  i:=0; i<len(adaptors); i++ {
		curr := adaptors[i]
		if curr == prev + 1 {
			subChainLen++
		} else if curr == prev + 3 {
			subChains[subChainLen]++
			subChainLen = 1
		} else {
			fmt.Println("Jump not 1 or 3", prev, curr)
		}

		prev = curr
	}
	subChains[subChainLen]++

	result := 1
	for k, v := range subChains {
		result *= int(math.Pow(float64(subChainPermutation(k)), float64(v)))
	}

	return result
}