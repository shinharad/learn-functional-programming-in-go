package packagea

import b "github.com/shinharad/learn-functional-programming-in-go/chapter06/02_circulardep/src/packageb"

func Atask() {
	println("A")
	b.Btask()
}
