func hasDuplicate(nums []int) bool {
    setResult := make(map[int]struct{})

    for _,value := range nums {
        setResult[value] =  struct{}{}
    }

    if len(nums) == len(setResult) {
   return false
}
return true

}
