
package main

import "testing"

func TestLimparCacheIPs(t *testing.T) {
    resultado := []string{
        "214.11.156.108",
        "197.222.252.246",
        "52.5.113.164",
    }

    LimparCacheIPs(&resultado)

    if resultado != nil {
        t.Errorf("esperava que o slice fosse nil, mas recebeu: %v", resultado)
    }
}
