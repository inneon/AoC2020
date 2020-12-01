package day01

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