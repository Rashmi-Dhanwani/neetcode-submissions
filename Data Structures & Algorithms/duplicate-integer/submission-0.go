func hasDuplicate(nums []int) bool {
    duplicateMap := make(map[int]bool)
    for _, value := range nums{
        ok := duplicateMap[value]; if ok {
           return true
        }
        duplicateMap[value]= true
    }
    return false
} 
