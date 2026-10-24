package main

import (
	"fmt"
	"strings"
	"sync"
)

func ValidarURL(url string) bool {

	if strings.HasPrefix(url, "https://") {
		return true
	}
	return false

}

func ValidarListaURLs(urls []string) map[string]bool {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	result := make(map[string]bool)
	
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			valid := ValidarURL(u)
			mu.Lock()
			result[u] = valid
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return result
}

func main() {

	fmt.Println("--- Ex 01 ---")

	urls := []string{
		"http://example.com/api/v1/recurso",  // false
		"https://proxy.golang.org/",          // true
		"https://golang.org/doc/",            // true
		"http://api.github.com/users/google", // false
	}

	for _, url := range urls {
		fmt.Println("Resultado de cada url se é valida", ValidarURL(url))
	}

	fmt.Println("-----")
	fmt.Println("Resultado do validador de urls concorrentes: ")
	for url, valid := range ValidarListaURLs(urls) {
		fmt.Printf("%-40s válida: %v\n", url, valid)
	}

}
