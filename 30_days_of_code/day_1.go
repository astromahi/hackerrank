// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-1-data-types
// Day 1: Data Types!

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

	// Mapping Java7 data types
	m := map[string]string{
		"int":     "int",
		"int8":    "byte",
		"int16":   "short",
		"int32":   "int",
		"int64":   "long",
		"float32": "float",
		"float64": "double",
		"bool":    "boolean",
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
