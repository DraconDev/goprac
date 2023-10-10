package main

func GetSize(w,h,d int) [2]int {
	result := [2]int{
		2*w*h+2*h*d+2*d*w,
		w*h*d,
	}
	return result
}

