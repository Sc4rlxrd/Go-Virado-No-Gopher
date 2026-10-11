package main

import "fmt"

func main() {

	//  Criando um Map com Literal Composto (map[TipoChave]TipoValor)
	servicePorts := map[int]string{
		80:   "HTTP",
		443:  "HTTPS",
		22:   "SSH",
		5432: "PostgreSQL",
		8080: "HTTP-ALT",
	}

	fmt.Println("--- Leitura e Exibição do Map ---")
	fmt.Println("Map completo:", servicePorts)
	fmt.Println("Porta 22 corresponde a:", servicePorts[22])

	//  Adicionando e Atualizando elementos [5]
	servicePorts[3306] = "MySQL"         // Adicionando um novo elemento
	servicePorts[80] = "Tomcat/Alt HTTP" // Atualizando um elemento existente
	fmt.Println("Após adição/atualização:", servicePorts)

	//  Removendo uma chave com delete(mapa, chave)
	delete(servicePorts, 8080)
	fmt.Println("Após excluir a porta 8080:", servicePorts)

	//  Acessando chave inexistente e o Idioma "comma ok"
	// Ao buscar uma chave ausente, o Go NÃO lança erro: ele retorna o 'zero value' do valor.
	serviceUnavailable := servicePorts[9000]
	fmt.Printf("Busca da porta 9000 (inexistente): %q (zero value)\n", serviceUnavailable)

	// Para checar se a chave REALMENTE existe no map, usamos o "comma ok"

	if service, ok := servicePorts[443]; ok {

		fmt.Printf("A porta 443 existe e o serviço é: %s\n", service)

	}

	if _, ok := servicePorts[9000]; !ok {
		fmt.Println("A porta 9000 NÃO está cadastrada no mapeamento!")
	}

	// 5 Inicialização via make()
	// // Importante: Um map declarando apenas 'var m map[string]int' é nil e causa PANIC ao tentar escrever.
	// Deve-se usar make() para alocar o map na memória [
	settingsVm := make(map[string]string)
	settingsVm["distro"] = "Fedora 44"
	settingsVm["region"] = "sa-saopaulo-1"
	settingsVm["cpu"] = "02"
	settingsVm["Kernel"] = "Linux 6.17.0-1020-oracle"
	fmt.Println("Map alocado via make():", settingsVm)

}
