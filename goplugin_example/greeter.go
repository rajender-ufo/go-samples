package main

import (
	"os"
	"fmt"
	"flag"
	"plugin"
)

type Greeter interface{
	Greet()
}

func main(){
	var lang string
	flag.StringVar(&lang,"lang","english","languages : english/telugu")
	flag.Parse()


	//1. identify plugin to load
	var mod string
	switch lang {
	case "english":
		mod ="./eng/eng.so"
	case "telugu":
		mod="./tel/tel.so"
	default:
		fmt.Println("unsupported language for this application")
		os.Exit(1)
	}

	//2.1 load module
	plug,err := plugin.Open(mod)
	if(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}
	//2.2 lookup for symbol in the plugin
	symGreeter ,err := plug.Lookup("Greeter")
	if(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}
	
	//2.3 assert that loaded symbol is of desired type i.e Greeter
	var greeter Greeter
	greeter,ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	greeter.Greet()

}
