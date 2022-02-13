// Fibonacci series - memoization

package main

import (
  "fmt"
)

func fib(n int, mem map[int]int) int {
  if _, ok := mem[n]; ok {
    return mem[n]
  }
  if n <= 2 {
    return 1
  }

  mem[n] = fib(n-1, mem) + fib(n-2, mem)
  return mem[n]
}

func main() {
  mem := make(map[int]int)
  fmt.Println("50th fibonacci number  = ", fib(50, mem))
}

