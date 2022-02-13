// Grid traveler problem - memoization

package main

import (
  "fmt"
  "strconv"
)

func gridTravelerMemoization(n,m int, memo map[string]int) int {
  key := strconv.Itoa(m) + "," + strconv.Itoa(n)
  if _, ok := memo[key]; ok {
    return memo[key]
  }

  if m == 1 && n == 1 {
    return 1
  } else if m == 0 || n == 0 {
    return 0
  }
  memo[key] = gridTravelerMemoization(m-1, n, memo) + gridTravelerMemoization(n, m-1, memo)
  return memo[key]
}

func main() {
  memo := make(map[string]int)
  numPaths := gridTravelerMemoization(18, 18, memo)
  fmt.Println("Number of paths for (18, 18) grid = ", numPaths)
}
