package main

type MinStack struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor_8() MinStack {
	s := MinStack{}
	return s
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	result := this.nums[0]
	for i := 0; i < len(this.nums); i++ {
		if result > this.nums[i] {
			result = this.nums[i]
		}
	}
	return result
}
