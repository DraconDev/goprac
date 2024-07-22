package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPeople(t *testing.T) {
	// [17233,32521,14087,42738,46669,65662,43204,8224]
	// ["WSpmqvb","EPCFFt","RPJOFYZUBFSIYp","VOYGWWNCf","Vk","Sgizfdfrims","IEO","QTASHKQ"]

	assert.Equal(t, sortPeople([]string{"WSpmqvb", "EPCFFt", "RPJOFYZUBFSIYp", "VOYGWWNCf", "Vk", "Sgizfdfrims", "IEO", "QTASHKQ"}, []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224}), []string{"EPCFFt", "RPJOFYZUBFSIYp", "VOYGWWNCf", "Vk", "Sgizfdfrims", "IEO", "QTASHKQ", "WSpmqvb"})

}
