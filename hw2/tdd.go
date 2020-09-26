package hw2

func merge(input []int, s1 int, e1 int, s2 int, e2 int) {
	var output []int

	len1 := e1 - s1 + 1
	len2 := e2 - s2 + 1
	save1, save2 := s1, s2
	for s1<=e1 && s2 <= e2 {
		if input[s1] < input[s2] {
			output = append(output, input[s1])
			s1++
		} else {
			output = append(output, input[s2])
			s2++
		}
	}

	for s1 <= e1 {
		output = append(output, input[s1])
		s1++
	}

	for s2 <= e2 {
		output = append(output, input[s2])
		s2++
	}
	i := 0
	for i = 0; i < len1; i++ {
		input[save1 + i] = output[i]
	}

	for i = len1; i<len2+len1;i++ {
		input[save2 + i - len1] = output[i]
	}
}
func MergeSort(input []int, start int, end int){
	if end - start <= 0 {
		return
	}

	if end - start == 1 {
		if input[end] < input[start] {
			temp := input[end]
			input[end] = input[start]
			input[start] = temp
		}
		return
	}


	mid := (start + end) / 2
	MergeSort(input, start, mid)
	MergeSort(input, mid+1, end)
	merge(input, start, mid, mid+1, end)
 }
