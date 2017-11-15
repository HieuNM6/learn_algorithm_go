package array

// IncreaseSubSequence get array of integer and return longest increasing subsequence
// will return more than 1 if have more than one longest increasing subsequence have same length
func IncreaseSubSequence(input []int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}

	if len(input) == 1 {
		return [][]int{input}
	}

	counter := 1
	m := make(map[int][][]int)
	start, stop := 0, 0

	for i := 1; i < len(input); i++ {
		switch {
		case len(input) == i+1:
			counter++
			m[counter] = append(m[counter], input[start:i+1])
		case input[i] <= input[i-1]:
			m[counter] = append(m[counter], input[start:stop+1])
			counter = 1
			start = i
		default:
			counter++
			stop = i
		}
	}

	max := 0
	for k, _ := range m {
		if k > max {
			max = k
		}
	}

	return m[max]
}
