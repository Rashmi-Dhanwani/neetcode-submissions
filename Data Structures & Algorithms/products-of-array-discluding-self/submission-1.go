func productExceptSelf(nums []int) []int {
result := make([]int, len(nums))
result[0] =1
for i:=1; i<=len(nums)-1 ; i++{
    result[i] = result[i-1]*nums[i-1]
}

r := 1
for j:=len(nums)-1 ; j>=0;j--{
  result[j] = result[j]*r
  r = r *nums[j] 
}

return result
}
