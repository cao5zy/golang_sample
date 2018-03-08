package main

import (
    "fmt"
)

type Struct1 struct {
    id int
}

type Struct2 struct {
    *Struct1
}

func main() {
     p := &Struct1 {id : 1}
     nd := &Struct2{ Struct1: p }
     fmt.Println(nd.id)
}