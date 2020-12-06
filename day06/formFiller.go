package day06

// FillForm fills the forms and gets the sum of the counts
func FillForm(entries []string) int {
	result := 0
	entries = append(entries, "")
	group := make(map[string]bool)

	for _, entry := range entries {
		if entry == "" {
			result += len(group)
			group = make(map[string]bool)
		} else {
			for i:=0; i<len(entry); i ++ {
				group[entry[i:i+1]] = true
			}
		}
	}


	return result
}