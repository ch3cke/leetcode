package main

type NestedIterator struct {
	nums []int
}

type NestedInteger struct {
}

func (i NestedInteger) IsInteger() bool {
	return true
}

func (i NestedInteger) GetInteger() int {
	return 0
}

func (i NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

func Constructor_5(nestedList []*NestedInteger) *NestedIterator {
	nums := []int{}
	var dfs func(nestedList []*NestedInteger)
	dfs = func(nestedList []*NestedInteger) {
		for _, i := range nestedList {
			if i.IsInteger() {
				nums = append(nums, i.GetInteger())
			} else {
				dfs(i.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{nums: nums}
}

func (this *NestedIterator) Next() int {
	result := this.nums[0]
	this.nums = this.nums[1:]
	return result
}

func (this *NestedIterator) HasNext() bool {
	return len(this.nums) > 0
}
