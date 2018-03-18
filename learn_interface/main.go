package main

import "fmt"
// this package aims to show how to convert interface to concrete types

func main(){
     data := make(map[string]interface{})

     data["a"] = 1
     data["b"] = "b"
     data["c"] = 0.8

     fmt.Println(data["a"].(int) + 1) //convert interface{} to int
     fmt.Println(data["b"].(string) + "oo") //convert interface{} to string
}