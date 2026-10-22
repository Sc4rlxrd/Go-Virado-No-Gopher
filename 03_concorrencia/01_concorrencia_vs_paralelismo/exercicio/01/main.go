package main

import (
	"fmt"
	"runtime"
	"time"
)

type DiagnosticoSistema struct {
	CPUsLogical     int
	MaxProcs        int
	GoroutinesVivas int
}

func ObterDiagnostico() DiagnosticoSistema {
	return DiagnosticoSistema{
		CPUsLogical:     runtime.NumCPU(),
		MaxProcs:        runtime.GOMAXPROCS(0),
		GoroutinesVivas: runtime.NumGoroutine(),
	}
}

func ChecarPreco(item int) {
	fmt.Printf("🔍 Checando preço do item %d...\n", item)
	time.Sleep(200 * time.Millisecond) // simula I/O e mantém a goroutine viva
}

func main() {
	d := ObterDiagnostico()
	fmt.Println("=== Diagnóstico inicial ===")
	fmt.Printf("CPUs lógicas     : %d\n", d.CPUsLogical)
	fmt.Printf("GOMAXPROCS       : %d\n", d.MaxProcs)
	fmt.Printf("Goroutines vivas : %d\n", d.GoroutinesVivas)

	for i := 1; i <= 10; i++ {
		go ChecarPreco(i)
	}

	// Logo após o disparo, antes de qualquer sleep na main
	fmt.Printf("Goroutines vivas após o disparo: %d\n", ObterDiagnostico().GoroutinesVivas)

	time.Sleep(500 * time.Millisecond) // espera os coletores terminarem
	fmt.Printf("Goroutines vivas ao final: %d\n", ObterDiagnostico().GoroutinesVivas)
}