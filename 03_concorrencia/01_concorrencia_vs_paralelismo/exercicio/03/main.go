package main

import (
	
	"runtime"
	
)

func DetectarVazamento(limiteMaximo int) (bool, int) {

	goativas := runtime.NumGoroutine()

	if goativas > limiteMaximo {
		return true, goativas
	} else {
		return false, goativas
	}

}