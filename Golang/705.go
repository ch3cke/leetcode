package main

type MyHashSet struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor_2() MyHashSet {
	s := MyHashSet{}
	return s
}

func (this *MyHashSet) Add(key int) {
	this.nums = append(this.nums, key)
}

func (this *MyHashSet) Remove(key int) {
	i := 0
	for {
		if i == len(this.nums) {
			break
		}
		if i == len(this.nums)-1 {
			if this.nums[i] == key {
				this.nums = this.nums[:i]
				break
			}
			break
		}
		if this.nums[i] == key {
			this.nums = append(this.nums[:i], this.nums[i+1:]...)
			i -= 1
		}
		i += 1
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == key {
			return true
		}
	}
	return false
}
