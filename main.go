package main

import (
	"fmt"
	"flag"
	dcilogparser "github.com/rajender-ufo/go-samples/smptelogparser"
)


func main(){

	var absfilepath string
	flag.StringVar(&absfilepath,"file","smptelog.xml","smpte log file")
	flag.Parse()
	fmt.Println("demo parsing "+absfilepath)
	dcilogparser.Parse(absfilepath)
}
