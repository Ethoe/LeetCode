package main

import (
	"LeetCode/problems"
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give problem number to solve")
		return
	}
	problems := new(problems.Problem)
	function := fmt.Sprintf("Run%s", os.Args[1])
	call(problems, function, nil)
}

func call(obj interface{}, fn string, args map[string]interface{}) (res []reflect.Value) {
	method := reflect.ValueOf(obj).MethodByName(fn)
	var inputs []reflect.Value
	for _, v := range args {
		inputs = append(inputs, reflect.ValueOf(v))
	}
	return method.Call(inputs)
}