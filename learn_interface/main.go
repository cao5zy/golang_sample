package main

import "fmt"

// this package aims to show how to convert interface to concrete types

func main() {
	data := make(map[string]interface{})

	data["a"] = 1
	data["b"] = "b"
	data["c"] = 0.8

	fmt.Println(data["a"].(int) + 1)       //convert interface{} to int
	fmt.Println(data["b"].(string) + "oo") //convert interface{} to string

	dog := Dog{ID: "2323223"}

	WhoCanWalk(dog)
}

type Animal interface {
	Walk()
}

type Dog struct {
	ID string
}

func (d Dog) Walk() {
	fmt.Println("dog can walk", d.ID)
}

func WhoCanWalk(animal Animal) {
	animal.Walk()
}
