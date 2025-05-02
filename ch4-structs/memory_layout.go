package main

import (
	"fmt"
	"reflect"
)

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

func main() {
	typ := reflect.TypeOf(stats{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())
}
