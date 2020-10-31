/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

// @lc code=end

