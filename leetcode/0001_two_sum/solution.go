package main
import "fmt"

func twoSum(nums []int, target int) []int {
    d := make(map[int]int)
    for i :=0 ; i<len(nums) ; i++ {
        m := target - nums[i]
        val,ok := d[m]
        if ok {
            return []int{val,i}
        } else {
            d[nums[i]] = i
        }
    }
    return nil
}

func main() {
  nums := [2,7,11,13]
  target := 9
  a := twoSum(nums,target)
  fmt.Prinln(a)
}
