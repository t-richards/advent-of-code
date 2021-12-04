package conv

func Transpose(input []string) []string {
	rowLen := len(input)
	columnLen := len(input[0])
	result := make([]string, columnLen)

	for col := 0; col < columnLen; col++ {
		for row := 0; row < rowLen; row++ {
			result[col] += string(input[row][col])
		}
	}

	return result
}
