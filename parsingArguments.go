package main

import (
	//"bufio"
	"flag"
	"fmt"
	//"io/ioutil"
	//"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	species := flag.String("species", "gopher", "the species we are studying")
	flag.Parse()
	fmt.Println(*species)
	//fmt.Print(string(dat))
}
