package main

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if x1 == x2 {
		return "YES"
	}
	var firstPos, firstSpeed = x1, v1
	var secondPos, secondSpeed = x2, v2
	if secondPos > firstPos {
		firstPos, firstSpeed = x2, v2
		secondPos, secondSpeed = x1, v1
	}

	if firstSpeed >= secondSpeed {
		return "NO"
	}
	for firstPos > secondPos {
		firstPos += firstSpeed
		secondPos += secondSpeed
		if firstPos == secondPos {
			return "YES"
		}
	}
	return "NO"
}
