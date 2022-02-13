// Fibonacci series - recursive implementation

package main

import (
  "fmt"
)

func fib(n int) int {
  if n <= 2 {
    return 1
  }

  return fib(n-1) + fib(n-2)
}

func main() {
  fmt.Println("50th fibonacci number is", fib(50))
}
