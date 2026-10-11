package main

import "fmt"

func main() {

	// Matriz estática 3x3 (Array de Arrays)
	var days [3][3]int

	days[0][0] = 11
	days[0][1] = 21
	days[0][2] = 31

	days[1][0] = 41
	days[1][1] = 51
	days[1][2] = 61

	days[2][0] = 71
	days[2][1] = 81
	days[2][2] = 91

	fmt.Println("--- Percorrendo Matriz com Range Aninhado ---")

	for i, linha := range days {
		for j, valor := range linha {
			fmt.Printf("[%d][%d] = %-2d ", i, j, valor)
		}
		fmt.Println()
	}
}
