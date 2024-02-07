package leetcode
import (
	"math"
)
func jump(nums []int) int {
	var jumps int = 0
	var curEnd  int = 0;
	var curFarthest int = 0;
	for i := 0; i < len(nums)-1 ; i++ {
		curFarthest = int(math.Max(float64(curFarthest),float64(i + nums[i])));
		if (i == curEnd) {
			jumps++;
			curEnd = curFarthest;
		}
	}
	return jumps;
}
