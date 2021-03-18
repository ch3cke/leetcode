package main

type MyHashMap struct {
	nums map[int]int
}

/** Initialize your data structure here. */
func Constructor_3() MyHashMap {
	s := MyHashMap{}
	if s.nums == nil {
		s.nums = make(map[int]int)
	}
	return s
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.nums[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	tmp, ok := this.nums[key]
	if ok == false {
		return -1
	} else {
		return tmp
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	delete(this.nums, key)
}
