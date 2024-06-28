package example

import "os"

func Panics() {

	// panic("a problem")

	_, err := os.Create("/tmp/file.txt")
	if err != nil {
		panic(err)
	}
}
