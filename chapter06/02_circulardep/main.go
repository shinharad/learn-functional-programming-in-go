package main

import (
	a "github.com/shinharad/learn-functional-programming-in-go/chapter06/02_circulardep/src/packagea"
)

// import cycle not allowed
func main() {
	a.Atask()
}
