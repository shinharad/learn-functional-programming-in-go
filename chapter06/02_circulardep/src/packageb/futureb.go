package packageb

import a "github.com/shinharad/learn-functional-programming-in-go/chapter06/02_circulardep/src/packagea"

func Btask() {
	println("B")
	a.Atask()
}
