func longestConsecutive(nums []int) int {
set := make(map[int]bool)
for _, n := range nums{
  set[n]=true
}

best := 0

for n := range set{
if !set[n-1]{
     length:=1
    for set[n+1]{
         length ++
        n++
    }
    if length > best{
        best = length
    }
}
}
return best

}
