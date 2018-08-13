package main

import (
	"github.com/alediaferia/go-collections"
	"strings"
	"fmt"
)

func main() {
	names := []interface{}{
		"Alice",
		"Bob",
		"Cindy",
	}
	collection := collections.NewFromSlice(names)
	collection = collection.Map(func(v interface{}) interface{} {
		return strings.Join([]string{"Hey", v.(string), "\n"}, " ")
	})
	fmt.Println(collection)
}
