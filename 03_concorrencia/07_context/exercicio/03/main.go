package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ResultadoBusca struct {
	Loja  string
	Preco float64
	Err   error
}

var (
	ErrNenhumaLojaFornecida = errors.New("nenhuma loja fornecida")
	ErrTodasFalharam        = errors.New("todas as lojas falharam")
)

// buscador permite trocar a busca por uma versão controlada nos testes.
type buscador func(ctx context.Context, loja string) ResultadoBusca

// Lojas simuladas: quanto cada uma demora e que preço responde.
var lojasSimuladas = map[string]struct {
	atraso time.Duration
	preco  float64
}{
	"Amazon":       {150 * time.Millisecond, 210.90},
	"MercadoLivre": {30 * time.Millisecond, 199.90},
	"Kabum":        {200 * time.Millisecond, 205.00},
	"Magalu":       {250 * time.Millisecond, 215.50},
	"Shopee":       {300 * time.Millisecond, 189.99},
}

func buscarPrecoSimulado(ctx context.Context, loja string) ResultadoBusca {
	cfg, ok := lojasSimuladas[loja]
	if !ok {
		cfg.atraso, cfg.preco = 100*time.Millisecond, 199.90
	}

	timer := time.NewTimer(cfg.atraso)
	defer timer.Stop()

	select {
	case <-timer.C: // terminou a busca
		return ResultadoBusca{Loja: loja, Preco: cfg.preco}
	case <-ctx.Done(): // alguém já ganhou (ou o chamador desistiu): aborta
		return ResultadoBusca{Loja: loja, Err: ctx.Err()}
	}
}

func BuscarPrimeiroPreco(ctx context.Context, lojas []string) (ResultadoBusca, error) {
	return buscarPrimeiro(ctx, lojas, buscarPrecoSimulado)
}

func buscarPrimeiro(ctx context.Context, lojas []string, buscar buscador) (ResultadoBusca, error) {
	if len(lojas) == 0 {
		return ResultadoBusca{}, ErrNenhumaLojaFornecida
	}

	ctxChild, cancel := context.WithCancel(ctx)
	defer cancel() // garante a limpeza em qualquer caminho de saída

	// Buffer len(lojas): as perdedoras sempre conseguem depositar e terminar
	chRes := make(chan ResultadoBusca, len(lojas))

	for _, loja := range lojas {
		go func(l string) {
			chRes <- buscar(ctxChild, l)
		}(loja)
	}

	var ultimoErr error
	for range lojas {
		select {
		case r := <-chRes:
			if r.Err == nil {
				cancel() // avisa as outras para abortarem já
				return r, nil
			}
			ultimoErr = r.Err // essa falhou: espera a próxima
		case <-ctx.Done(): // o chamador cancelou ou o prazo dele estourou
			return ResultadoBusca{}, ctx.Err()
		}
	}
	return ResultadoBusca{}, fmt.Errorf("%w: %v", ErrTodasFalharam, ultimoErr)
}

func main() {
	fmt.Println(" --- Ex 03 --- ")

	lojas := []string{"Amazon", "MercadoLivre", "Kabum", "Magalu", "Shopee", "TeraByte"}

	inicio := time.Now()
	r, err := BuscarPrimeiroPreco(context.Background(), lojas)
	fmt.Printf("vencedora: %s R$ %.2f | err=%v | levou %v\n", r.Loja, r.Preco, err, time.Since(inicio).Round(time.Millisecond))
}
