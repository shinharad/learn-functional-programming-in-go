package reduce

import (
	"github.com/alediaferia/go-collections"
	"fmt"
)

func main() {

	numbers := []interface{}{
		1,
		5,
		3,
		2,
	}

	collection := collections.NewFromSlice(numbers)
	max := collection.Reduce(0, func(a, b interface{}) interface{} {
		if a.(int) > b.(int) { return a } else { return b }
	})
	fmt.Println(max)

}


