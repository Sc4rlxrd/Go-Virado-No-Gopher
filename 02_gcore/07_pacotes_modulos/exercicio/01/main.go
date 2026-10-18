package main

import "fmt"

var VersaoSistema string

func init() {
	VersaoSistema = "v1.0.0-fedora"
	fmt.Println("Dentro do init:", VersaoSistema)
}

func main() {
	fmt.Println(" === Ex01 ===")
}
