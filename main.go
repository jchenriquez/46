package main

import "fmt"

func permuteH (nums[]int, currentIndex int) [][]int {
  if currentIndex == (len(nums)-1) {
    return [][]int{[]int{nums[currentIndex]}}
  }

  lastGenerated := permuteH(nums, currentIndex+1)
  result := make([][]int, 0, len(lastGenerated) * 2)

  for _, generated := range lastGenerated {
    
    for i := 0; i <= len(generated); i++ {
      newGen := make([]int, len(generated)+1)
      copy(newGen, generated[:i])
      newGen[i] = nums[currentIndex]
      copy(newGen[i+1:], generated[i:])
      result = append(result, newGen)
    }

  }

  return result
}

func permute (nums []int) [][]int {
  return permuteH(nums, 0)
}


func main() {
  fmt.Printf("%v\n", permute([]int{1,2,3}))
}