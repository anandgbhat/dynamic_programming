// Return array of numbers that sum upto a target value given an array of
// numbers.
// Any number can be used any number of times

package main

import (
  "fmt"
)

func howSum(targetSum int, numbers []int) []int {
  // Base case
  if targetSum == 0 {
    return make([]int, 0)
  }

  if targetSum < 0 {
    return nil
  }

  for _, num := range numbers {
    remainder := targetSum - num
    fmt.Println(remainder, num)
    result := howSum(remainder, numbers)
    fmt.Println(result)
    if result != nil {
      // If the result is non-null (edge case), append the number that we used
      return append(result, num)
    }
  }
  return nil
}

func main() {
  targetSum := 7
  numbers := []int {3, 4, 6, 7}
  fmt.Println("howSum = ", howSum(targetSum, numbers))
}
