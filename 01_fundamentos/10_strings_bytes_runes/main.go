package main

import "fmt"

func main() {

	// Em Go, strings são sequências imutáveis de BYTES codificadas em UTF-8 por convenção

	text := "Golang 🚀"
	fmt.Println("--- Tamanho em Bytes vs. Caracteres ---")

	// len() retorna o número de BYTES da string, não o número de letras!

	fmt.Printf("String: %s\n", text)
	fmt.Printf("Tamanho via len(texto): %d bytes (Emoji ocupa 4 bytes em UTF-8!)\n", len(text))

	//  Iteração por Índice (Acessa BYTES individuais)
	fmt.Println("--- Iteração por Índice (Bytes brutos em Hexadecimal) ---")

	for i := 0; i < len(text); i++ {
		fmt.Printf("Byte %d: %x ('%c')\n", i, text[i], text[i])
	}

	//  Iteração com 'range' (Decodifica automaticamente cada RUNE / Unicode Code Point)
	// 'rune' é um alias nativo para 'int32' e representa um caractere Unicode

	fmt.Println("--- Iteração com 'range' (Decodificação UTF-8 em Runes) ---")
	for bytePosition, charRune := range text {
		fmt.Printf("Posição byte %2d -> Rune: %U | Caractere: %c\n", bytePosition, charRune, charRune)
	}

	//  Conversões de Tipos entre string, []byte e []rune

	fmt.Println("--- Conversões ---")
	sliceBytes := []byte("Fedora")
	sliceBytes[0] = 'f' // Modificando o byte
	modifiedText := string(sliceBytes)
	fmt.Printf("Original alterado via []byte: %s\n", modifiedText)
	runes := []rune(text)
	fmt.Printf("Quantidade REAL de caracteres Unicode (runes): %d\n", len(runes))
}
