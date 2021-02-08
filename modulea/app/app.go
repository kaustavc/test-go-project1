package main

import (
	"fmt"
	"github.com/kaustavc/test-go-project/modulea/mathlib"
	"github.com/kaustavc/test-go-project/moduleb"
)

func main() {
	x, y := 4, 5
	z := mathlib.Add(x, y)
	s := moduleb.ToCaps("hello")
	fmt.Println(z)
	fmt.Println(s)
}