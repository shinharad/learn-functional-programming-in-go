package packagea

import b "github.com/shinharad/learn-functional-programming-in-go/chapter06/01_dependency-rule-good/src/packageb"

func Atask() {
	println("A")
	b.Btask()
}
