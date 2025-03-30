func twoSum(nums []int, target int) []int {
    differenceToIndex := make(map[int]int)
    
    for i, num := range nums {
        if index, found := differenceToIndex[num]; found {
            return []int{index, i}
        }
        differenceToIndex[target-num] = i
    }
    
    return []int{}
}