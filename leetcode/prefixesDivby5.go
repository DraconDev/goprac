package leetcode

import "math/big"

func prefixesDivBy5(nums []int) []bool {
	result := []bool{}
	bigNum := big.NewInt(0)
	for _, v := range nums {
		bigNum.Lsh(bigNum, 1)
		bigNum.Add(bigNum, big.NewInt(int64(v)))
		result = append(result, bigNum.Mod(bigNum, big.NewInt(5)).Int64() == 0)
	}

	return result
}
