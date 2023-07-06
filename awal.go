package main

import "fmt"

type Person struct {
	name string
	age  int
}
type People struct {
	name string
	age  int
	alias string
}

type TestStruct struct {
	id int
	name string
}

var structs []TestStruct

func main(){
	fmt.Println("awal belajar");

	structs = append(structs, TestStruct{id: 1, name: "test 1"})
	structs = append(structs, TestStruct{id: 2, name: "test 2"})
	structs = append(structs, TestStruct{id: 3, name: "test 3"})
	structs = append(structs, TestStruct{id: 4, name: "test 4"})
	structs = append(structs, TestStruct{id: 5, name: "test 5"})

	fmt.Println("teststructs", structs)

	structs = append(structs[:2], structs[3:]...)

	fmt.Println("teststructs after", structs)

}