package main

import (
	"fmt"
	"time"
)

func AguardarResultadoComTimeout[T any](chIn <-chan T, timeout time.Duration) (T, error) {
	var zero T
	select {
	case res := <-chIn:
		return res, nil
	case <-time.After(timeout):
		return zero, fmt.Errorf("operacao expirou apos %v", timeout)
	}

}

func main() {
	fmt.Println(" --- Ex 01 --- ")

	rapido := make(chan string, 1)

	go func() {

		time.Sleep(55 * time.Millisecond)
		rapido <- "Amazon: R$199.90"

	}()

	res, err := AguardarResultadoComTimeout(rapido, 500*time.Millisecond)
	fmt.Printf("rápido -> res=%q err=%v\n", res, err)

	lento := make(chan string, 1)

	go func() {

		time.Sleep(300 * time.Millisecond)
		lento <- "Kabum: R$205.90"

	}()

	res, err = AguardarResultadoComTimeout(lento, 100*time.Millisecond)
	fmt.Printf("lento -> res=%q err=%v\n", res, err)

}
