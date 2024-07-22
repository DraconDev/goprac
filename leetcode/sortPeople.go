package leetcode

import (
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	// zip names and heights

	type Person struct {
		Name   string
		Height int
	}
	zipped := make([]Person, len(names))

	for i := 0; i < len(names); i++ {
		zipped[i] = Person{names[i], heights[i]}
	}
	// sort by height
	sort.Slice(zipped, func(i, j int) bool {
		return zipped[i].Height > zipped[j].Height
	})

	// return names
	ret := make([]string, len(names))
	for i := 0; i < len(names); i++ {
		ret[i] = zipped[i].Name
	}
	return ret
}
