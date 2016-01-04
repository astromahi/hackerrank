package main

import (
	"fmt"
	"reflect"
)

func main() {

	t := []interface{}{
		5.35,
		"a",
		false,
		100,
		"I am a code monkey",
		true,
		17.3,
		"c",
		"derp",
	}

	m := map[string]string{
		"float32": "double",
		"float64": "double",
		"bool":    "boolean",
		"int":     "int",
		"int8":    "int",
		"int16":   "int",
		"int32":   "int",
		"int64":   "int",
	}

	for _, v := range t {
		kind := reflect.TypeOf(v).Kind().String()
		if kind == "string" && len(v.(string)) == 1 {
			fmt.Println("Primitive : char")
		} else if kind == "string" {
			fmt.Println("Referenced : String")
		} else {
			fmt.Printf("Primitive : %s\n", m[kind])
		}
	}
}
