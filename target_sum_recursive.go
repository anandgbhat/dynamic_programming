// Given a array of integers, check if there is a possibility of generating
// targetSum. Any number in the array can be used multiple times.
// Function returns true or false
// Logic is to take target number and subtract it with each number in the
// array till one of the base cases are hit.
// Base case: if the remaining target sum is 0, then there is at least one
// possible set that results in the given sum. Return true in this case.
// If the target sum is negative, return false.

// Time complexity: O(n^m)
// Space complexity : O(m)

package main

import (
  "fmt"
)

func canSum(targetSum int, numbers []int) bool {
  if targetSum == 0 {
    return true
  }

  if targetSum < 0 {
    return false
  }

  for _, num := range numbers {
    remainder := targetSum - num
    fmt.Println(remainder, num)
    if canSum(remainder, numbers) == true {
      return true
    }
  }

  return false
}

func main() {
  numbers := []int {5, 3, 4, 6, 7}
  targetSum := 7

  fmt.Println(" CanSum = ", canSum(targetSum, numbers))
}
