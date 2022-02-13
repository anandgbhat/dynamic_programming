// Grid traveler program
// Given a grid on (m,n) size, figure out number of paths to reach from top
// left corner to bottom right corner.
// Constraint: Only moves allowed are down and right.
// Time complexity = O(2^(m+n))
// Space complexity = O(n+m)

package main

import (
  "fmt"
)

func gridTravelerRecursive(m,n int) int {
  if m == 1 && n == 1 {
    return 1
  } else if m == 0 || n == 0 {
    return 0
  }

  return gridTravelerRecursive(m-1, n) + gridTravelerRecursive(m, n-1)
}

func main() {
  numPaths := gridTravelerRecursive(18, 18)
  fmt.Println("Number of paths for (18, 18) grid = ", numPaths)
}


