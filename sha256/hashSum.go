package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1=%x\nc2=%x\nc1==c2 -> %t\ntype(c1) -> %T\n", c1, c2, c1 == c2, c1)
}
