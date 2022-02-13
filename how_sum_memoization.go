// Return array of numbers that sum upto a target value given an array of
// numbers.
// Any number can be used any number of times

package main

import (
  "fmt"
)

func howSum(targetSum int, numbers []int, mem map[int][]int) []int {
  // Base case
  if _, ok := mem[targetSum]; ok {
    return mem[targetSum]
  }

  if targetSum == 0 {
    return make([]int, 0)
  }

  if targetSum < 0 {
    return nil
  }

  for _, num := range numbers {
    remainder := targetSum - num
    fmt.Println(remainder, num)
    result := howSum(remainder, numbers, mem)
    fmt.Println(result)
    if result != nil {
      // If the result is non-null (edge case), append the number that we used
      mem[targetSum] = append(result, num)
      return mem[targetSum]
    }
  }

  mem[targetSum] = nil
  return nil
}

func main() {
  targetSum := 7
  numbers := []int {3, 4, 6, 7}
  mem := make(map[int][]int)
  fmt.Println("howSum = ", howSum(targetSum, numbers, mem))
  targetSum = 300
  numbers = []int {7, 14}
  mem = make(map[int][]int)
  fmt.Println("howSum = ", howSum(targetSum, numbers, mem))
  targetSum = 7
  numbers = []int {2, 3}
  mem = make(map[int][]int)
  fmt.Println("howSum = ", howSum(targetSum, numbers, mem))
}
