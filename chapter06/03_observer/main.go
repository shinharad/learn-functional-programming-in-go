package main

import (
	. "github.com/shinharad/learn-functional-programming-in-go/chapter06/03_observer/src/observer"
)

func main() {

	subject := Subject{}
	oa := Observable{Name: "A"}
	ob := Observable{Name: "B"}
	subject.AddObserver(&Observer{})
	subject.NotifyObservers(oa, ob)

	oc := Observable{Name: "C"}
	subject.NotifyObservers(oa, ob, oc)

	subject.DeleteObserver(&Observer{})
	subject.NotifyObservers(oa, ob, oc)

	od := Observable{Name: "D"}
	subject.NotifyObservers(oa, ob, oc, od)

}
