package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int

	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name()) // returns int

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())        // returns empty string
	fmt.Println(xpt.Kind())        // returns reflect.Ptr
	fmt.Println(xpt.Elem().Name()) // retruns int
	fmt.Println(xpt.Elem().Kind()) // returns reflect.Int

	type Foo struct {
		A string `myTag:"value1"`
		B int    `myTag:"value2"`
	}

	var y Foo

	yt := reflect.TypeOf(y)

	for i := 0; i < yt.NumField(); i++ {
		field := yt.Field(i)
		fmt.Println("field name", field.Name)
		fmt.Println("field type name", field.Type.Name())
		fmt.Println("field tag value", field.Tag.Get("myTag"))
	}

	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	fmt.Println("interface of s:", sv.Interface()) // returns an interface\

	stringType := reflect.TypeOf((*string)(nil)).Elem()
	w := reflect.New(stringType).Elem()
	w.SetString("mahdi")
	fmt.Println(w.Interface().(string))

	j := []string(nil)
	var k interface{}
	k = j

	fmt.Println(k == nil)

	kv := reflect.ValueOf(k)

	fmt.Println(kv.IsValid())
	fmt.Println(kv.IsNil())
}
