package metrics

import (
	"fmt"
)

type Coletor struct {
	TotalRequisicoes int
}

func (c *Coletor) RegistrarUsoCPU(porcentagem float64) {
	fmt.Printf("[METRICS INTERNAL] Uso de CPU registrado: %.2f%%\n", porcentagem)
}
