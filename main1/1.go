package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for j, num := range nums {
		if i, ok := hashTable[target-num]; ok {
			return []int{i, j}
		}
		hashTable[num] = j
	}
	return nil
}

func main() {
	//nums := []int{2, 7, 11, 15}
	//res := twoSum(nums, 9)
	unixTime := int64(1713238505)
	t := time.Unix(unixTime, 0)
	fmt.Println("当前时间: ", t)
	//fmt.Println("result: ", res)
}
