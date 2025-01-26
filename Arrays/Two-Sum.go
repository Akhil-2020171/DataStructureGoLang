package main

func twoSum(nums []int, target int) []int {
    // initialising map for storing the nums
    numMap := make(map[int]int);

    for i, num := range nums {
        // first check if the element is present in the map
        if idx, ok := numMap[target-num];  ok {
            return []int{idx,i};
        } else{
            numMap[num] = i;
        }
    }

    return nil;
}