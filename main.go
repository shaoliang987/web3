package main

import "fmt"

/**
 * 结构体练习
 */

type person struct {
	Name  string
	Age   int
	Call  func()
	Map   map[string]string
	Ch    chan string
	Arr   [5]uint8
	Slice []interface{}
	Ptr   *int
}

func main() {
	person1 := person{
		Name: "张三1",
		Age:  18,
		Call: func() {
			fmt.Println("call")
		},
		Map: map[string]string{
			"1": "1",
			"2": "2",
		},
		Ch: make(chan string),
		Arr: [5]uint8{
			1, 2, 3, 4, 5,
		},
		Slice: []interface{}{
			1, "3", true,
		},
		Ptr: new(int),
	}
	fmt.Println(person1)
}
