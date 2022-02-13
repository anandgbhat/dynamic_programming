// Given an integer arrray and a target sum, check if addition of any
// combination of numbers in the array results in the given sum.
// If so, return true. Else return false.
//
// Memoize all intermediate results to avoid expanding subtrees in the
// decision tree.

// Time complexity: O(m * n)
// Space complexity: O(m)

package main

import (
  "fmt"
)

func canSum(targetSum int, numbers []int, mem map[int]bool) bool {

  if _, ok := mem[targetSum]; ok {
    return mem[targetSum]
  }

  if targetSum == 0 {
    return true
  }

  if targetSum < 0 {
    return false
  }

  for _, num := range numbers {
    remainder := targetSum - num
    if canSum(remainder, numbers, mem) == true {
      mem[targetSum] = true
      return true
    }
  }

  mem[targetSum] = false
  return false
}

func main() {
  mem := make(map[int]bool)
  targetSum := 7
  numbers := []int {7, 3, 4, 5}

  fmt.Println("CanSum ", canSum(targetSum, numbers, mem))
}
