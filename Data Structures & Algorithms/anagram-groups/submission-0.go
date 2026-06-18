func groupAnagrams(strs []string) [][]string {
 
 groups := make(map[[26]int][]string)
 for _ , values := range strs{
   var key [26]int
   for _, val := range values{
      key[val-'a']++
   }
 groups[key] = append(groups[key],values)

 }

result := make([][]string, 0, len(groups))
 for _ ,v := range groups{
  result = append(result, v)
 }
return result
 }


