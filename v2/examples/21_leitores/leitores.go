package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_leitores()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _leitores() {
	fmt.Println("_leitores()")
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
