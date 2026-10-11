package main

import "fmt"

func main(){

   // Em Go, o tamanho do Array faz parte do seu TIPO. 
   // O tipo abaixo é '[4]int', que é totalmente DIFERENTE de '[5]int'.

   var ages[9]int

   fmt.Println("--- Zero Values em Arrays ---")
   fmt.Printf("Array sem inicialização: %v\n", ages)

   // Atribuição individual

   ages[0] = 1
   ages[1] = 2
   ages[2] = 3
   ages[3] = 4
   ages[4] = 5
   ages[5] = 6
   ages[6] = 7
   ages[7] = 8
   ages[8] = 9


   fmt.Printf("Array com valores atribuídos: %v\n", ages)
   fmt.Printf("Array preenchido: %v\n", ages) 
   for i := 0; i < len(ages); i++ {
	  fmt.Printf("Elemento no índice: [%d] = %d\n", i, ages[i])
   }
}
