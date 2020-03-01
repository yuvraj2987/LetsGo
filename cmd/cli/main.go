package main

import (
	"flag"
	"fmt"
)

func main() {
	ptrO := flag.Bool("o", false, "Boolean option")
	ptrI := flag.Int("i", 10, "Interger option")
	ptrK := flag.Bool("k", true, "Boolean option")
	flag.Parse()
	fmt.Println("-o: ", *ptrO)
	fmt.Println("-i: ", *ptrI)
	fmt.Println("-k: ", *ptrK)
	fmt.Println("=================")
	for index, value := range flag.Args() {
		fmt.Println(index, ":", value)
	}

}
