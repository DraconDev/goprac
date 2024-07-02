package kata

import (
	"fmt"
	"strconv"
)

func Hello() {
	str := "Hello, 世界!"

	result := strconv.QuoteToASCII(str)

	fmt.Println(result)
}
