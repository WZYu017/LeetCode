package leetcode

import "sort"

// FindRadius link https://leetcode-cn.com/problems/heaters/
func FindRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	findMin := func(house int, heaters []int) int {
		// todo
	}
}

// 使用两层循环算出最近加热器的最大值

// func FindRadius(houses []int, heaters []int) int {
// 	if len(houses) == 0 || len(heaters) == 0 {
// 		return 0
// 	}
// 	findMin := func(house int, heaters []int) int {
// 		abs := func(a int) int {
// 			if a >= 0 {
// 				return a
// 			}
// 			return -a
// 		}
// 		min := abs(heaters[0] - house)
// 		for _, v := range heaters {
// 			r := abs(v - house)
// 			if min > r {
// 				min = r
// 			}
// 		}
// 		return min
// 	}
// 	max := 0
// 	for i, size := 0, len(houses); i < size; i++ {
// 		r := findMin(houses[i], heaters)
// 		if r > max {
// 			max = r
// 		}
// 	}
// 	return max
// }
