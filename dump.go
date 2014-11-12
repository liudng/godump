// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"reflect"
	"fmt"
	"strconv"
)

var indent int64 = -1

// Print to standard out the value that is passed as the argument with indentation.
// Pointers are dereferenced.
func Dump(v interface{}) {
	val := reflect.ValueOf(v)
	dump(val, "")
}

func dump(val reflect.Value, name string) {
	indent++

	typ := val.Type()

	switch typ.Kind() {
	case reflect.Array, reflect.Slice:
		l := val.Len()
		for i := 0; i < l; i++ {
			dump(val.Index(i), strconv.Itoa(i))
		}
	case reflect.Map:
		printType(name, val.Interface())
		//l := val.Len()
		keys := val.MapKeys()
		for _, k := range keys {
			dump(val.MapIndex(k), k.Interface().(string))
		}
	case reflect.Ptr:
		printType(name, val.Interface())
		dump(val.Elem(), name)
	case reflect.Struct:
		printType(name, val.Interface())
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			dump(val.FieldByIndex([]int{i}), field.Name)
		}
	default:
		printValue(name, val.Interface())
	}

	indent--
}

func printIndent(){
	var i int64
	for i = 0; i < indent; i++ {
		fmt.Printf("  ")
	}
}

func printType(name string, v interface{}){
	printIndent()
	fmt.Printf("%s(%T)\n", name, v)
}

func printValue(name string, v interface{}){
	printIndent()
	fmt.Printf("%s(%T) %#v\n", name, v, v)
}
