package main

import (
	"fmt"

	arraystack "github.com/agustfricke/dsa-go/stack/array_stack"
)

func main() {
  stack := arraystack.NewStack()
  stack.Push(2)
  stack.Push(69)
  val := stack.GetSize()
  fmt.Println(val)
}
