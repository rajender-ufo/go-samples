package main

import "fmt"

type greeting string


func (g greeting) Greet(){
	fmt.Println("నమస్కారము అతిధి")
}

var Greeter greeting
