package sort

func QuickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	pivot := input[0]
	less := []int{}
	greater := []int{}

	parted := input[1:]

	for i := 0; i < len(parted); i++ {
		if parted[i] < pivot {
			less = append(less, parted[i])
		} else {
			greater = append(greater, parted[i])
		}
	}

	output := []int{}
	output = append(output, (QuickSort(less))...)
	output = append(output, pivot)
	output = append(output, (QuickSort(greater))...)

	return output
}
