package hw2


func Repeat(character string, times int) string {
	var output string
	for i := 0;i<times;i++ {
		output += character
	}
	return output
}

