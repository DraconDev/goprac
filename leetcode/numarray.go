package leetcode

type NumArray struct {
	nums []int
}

func Consty(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int

	for _, v := range this.nums[left : right+1] {
		sum += v
	}
	return sum

}
