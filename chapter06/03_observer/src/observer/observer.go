package observer

type Observable struct {
	Name string
}


type Callback interface {
	Notify(o *Observable)
}

// The observer implements the Callback interface
type Observer struct {}

func (ob *Observer) Notify(o *Observable) {
	println(o.Name)
}


