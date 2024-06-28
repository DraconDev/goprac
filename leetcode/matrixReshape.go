package leetcode

func MatrixReshape(mat [][]int, r int, c int) [][]int {

	elems := make([]int, 0)
	for _, row := range mat {
		elems = append(elems, row...)
	}

	if len(elems) != r*c {
		return mat
	}

	// flat matrix
	result := make([][]int, r)
	for i := range result {
		result[i] = elems[i*c : (i+1)*c]
	}

	return result
}
