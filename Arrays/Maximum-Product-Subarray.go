package main

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	currentMax := nums[0]
	currentMin := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = max(nums[i], currentMax * nums[i])
		currentMin = min(nums[i], currentMin * nums[i])
		maxProduct = max(maxProduct, currentMax)
	}
	return maxProduct
}