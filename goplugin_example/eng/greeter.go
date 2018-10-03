package main

import (
	"fmt"
)

type greeting string

func (g greeting) Greet(){
	fmt.Println("Hello Guest")
}

var Greeter greeting
