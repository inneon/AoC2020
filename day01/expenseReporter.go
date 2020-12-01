package day01

// FindPair find the product of 2 numbers in candiates that sum to target
func FindPair(candidates []int, target int) int {
	for i := 0; i<len(candidates); i++ {
		for j := i+1 ; j<len(candidates); j++ {
			if candidates[i] + candidates[j] == target {
				return candidates[i] * candidates[j]
			}
		}
	}
	return target;
}

// FindTriple find the product of 3 numbers in candiates that sum to target
func FindTriple(candidates []int, target int) int {
    for i := 0; i<len(candidates); i++ {
        for j := i+1 ; j<len(candidates); j++ {
            for k := j+1 ; k<len(candidates); k++ {
                if candidates[i] + candidates[j] + candidates[k] == target {
					return candidates[i] * candidates[j] * candidates[k]
                }
            }
        }
    }
    return target;
}