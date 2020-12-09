package day09

func isSum(window []int, target int) bool {
	for i := 0; i<len(window); i++ {
		for j := i+1 ; j<len(window); j++ {
			if window[i] + window[j] == target {
				return true
			}
		}
	}
	return false
}

// FirstInvalid gets the first invalid code in the series
func FirstInvalid(series []int, preamble int) (int, int) {
	for i:=preamble; i<len(series); i++ {
		if !isSum(series[i-preamble:i], series[i]) {
			return series[i], i
		}
 	}
	return 0, -1
}

// EncryptionWeakness gets the encryption weakness
func EncryptionWeakness(series []int, preamble int) int {
	target, index := FirstInvalid(series, preamble)
	for l:=2; l<index; l++ {
		for i:=0; i<index-l; i++ {
			sum := 0
			min := 4294967295
			max := 0
			for j:=0; j<l; j++ {
				sum+=series[i+j]
				if series[i+j] < min {
					min = series[i+j]
				}
				if series[i+j] > max {
					max = series[i+j]
				}
			}
			if sum==target {
				return min + max
			}
		}
	}
	return 0;
}