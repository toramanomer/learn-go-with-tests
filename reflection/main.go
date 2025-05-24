package main

import (
	"fmt"
	"reflect"
)

func walk(x any, fn func(input string)) {
	reflectValue := getValue(x)

	switch reflectValue.Kind() {
	case reflect.Struct:
		for i := range reflectValue.NumField() {
			walk(reflectValue.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := range reflectValue.Len() {
			walk(reflectValue.Index(i).Interface(), fn)
		}
	case reflect.Map:
		iter := reflectValue.MapRange()
		for iter.Next() {
			walk(iter.Value().Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := reflectValue.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		for _, out := range reflectValue.Call(nil) {
			walk(out.Interface(), fn)
		}
	case reflect.String:
		fn(reflectValue.String())
	}
}

func getValue(x any) reflect.Value {
	reflectValue := reflect.ValueOf(x)
	if reflectValue.Kind() == reflect.Pointer {
		reflectValue = reflectValue.Elem()
	}

	return reflectValue
}

func main() {
	ch := make(chan string)
	go func() {
		ch <- "1"
		ch <- "2"
		close(ch)
	}()

	walk(ch, func(input string) {
		fmt.Println("input:", input)
	})

	fmt.Println("lol?")
}
