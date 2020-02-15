package leetcode

// 1、直接计算两个邻近值的差异 2、记录开始结束后计算

// FindPoisonedDuration link https://leetcode-cn.com/problems/teemo-attacking/
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := duration
	for i, max := 1, len(timeSeries); i < max; i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			sum += duration
		} else {
			sum += timeSeries[i] - timeSeries[i-1]
		}
	}
	return sum
}

// func FindPoisonedDuration(timeSeries []int, duration int) int{
// 	sum, start, end := 0, 0, 0
// 	for _, v := range timeSeries {
// 		if v <= end {
// 			end = v + duration
// 		} else {
// 			sum += end - start
// 			start, end = v, v+duration
// 		}
// 	}
// 	sum += end - start
// 	return sum
// }

//
