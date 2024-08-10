package leetcode

func validMountainArray(arr []int) bool {

	if len(arr) < 3 {
		return false
	}
	down := false

	if arr[0] >= arr[1] {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		} else if down == true && arr[i] > arr[i-1] {
			return false
		} else if down == false && arr[i] < arr[i-1] {
			down = true
		}
	}
	if down == false {
		return false
	}
	return true

}
