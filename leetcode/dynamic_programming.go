package main

import (
	"fmt"
	"math"
)

func fibBottomToUp(n int) int {
	slice := make([]int, n+1)
	slice[0] = 0
	slice[1] = 1
	for i := 2; i <= n; i++ {
		slice[i] = slice[i-2] + slice[i-1]
	}
	return slice[n]
}

var fibMap = make(map[int]int)

func fibUpToBottom(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	value, exist := fibMap[n]
	if exist {
		return value
	} else {
		value = fibUpToBottom(n-1) + fibUpToBottom(n-2)
		fibMap[n] = value
	}
	return fibMap[n]
}

func rob(nums []int) int {
	return robBottomToUp(nums)
}

func robUpToBottom(i int, nums []int) int {

	if i == 0 {
		return nums[0]
	}

	if i == 1 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	return int(math.Max(float64(robUpToBottom(i-1, nums)),
		float64(robUpToBottom(i-2, nums)+nums[i])))
}

func robBottomToUp(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	slice := make([]int, length)
	slice[0] = nums[0]
	slice[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < length; i++ {
		slice[i] = int(math.Max(float64(slice[i-1]),
			float64(slice[i-2]+nums[i])))
	}
	return slice[length-1]
}

func minCostClimbingStairs(cost []int) int {
	//return minCostClimbingStairsTopToBottom(len(cost), cost)
	pre, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, int(math.Min(float64(pre+cost[i-2]),
			float64(cur+cost[i-1])))
	}
	return cur
}

func minCostClimbingStairsTopToBottom(i int, cost []int) int {

	if i == 0 || i == 1 {
		return 0
	}
	minCost := int(math.Min(float64(minCostClimbingStairsTopToBottom(i-1, cost)+cost[i-1]),
		float64(minCostClimbingStairsTopToBottom(i-2, cost)+cost[i-2])))
	return minCost
}

func main() {
	//fmt.Println(fibBottomToUp(1))
	//fmt.Println(fibUpToBottom(1))
	//fmt.Println(rob([]int{2, 7, 9, 3, 1}[:]))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}[:]))
}
