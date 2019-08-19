package main

import (
	"fmt"
	"reflect"
)

type Object1 struct{}
type Object2 struct{}
type Object3 struct{}

var typeRegistry = map[string]reflect.Type{
	"one":   reflect.TypeOf(Object1{}),
	"two":   reflect.TypeOf(Object2{}),
	"three": reflect.TypeOf(Object3{}),
}

func GetStructByName(name string) interface{} {
	return reflect.ValueOf(typeRegistry[name]).Interface()
}

func main() {
	obj1 := GetStructByName("one")
	fmt.Println(obj1)

	obj2 := GetStructByName("two")
	fmt.Println(obj2)

	obj3 := GetStructByName("three")
	fmt.Println(obj3)
}
