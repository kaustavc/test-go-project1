package main

import (
	"fmt"
	"github.com/kaustavc/test-go-project1/packagea/mathlib"
	"github.com/kaustavc/test-go-project1/packageb"
)

func main() {
	x, y := 4, 5
	z := mathlib.Add(x, y)
	s := packageb.ToCaps("hello")
	fmt.Println(z)
	fmt.Println(s)
}
