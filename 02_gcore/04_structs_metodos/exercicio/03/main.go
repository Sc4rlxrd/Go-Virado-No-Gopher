package main

import "fmt"

type RecursoK8s struct {
	Nome      string
	Namespace string
}

type Deployment struct {
	RecursoK8s
	Replicas int
}

func Escalar(deploy *Deployment, novasReplicas int) {

	deploy.Replicas++

	fmt.Printf("[%s/%s] Escalado para %d réplicas\n", deploy.Namespace, deploy.Nome, deploy.Replicas)

}

func main() {


	fmt.Println("---- Exercício 03 -----")

	deploy := Deployment{
		RecursoK8s: RecursoK8s{
			Nome:      "Teste",
			Namespace: "Golang",
		},
		Replicas: 0,
	}

	fmt.Printf("Deploy config padrão: \nNome: %s\nNamespace: %s\nReplicas: %d\n", deploy.Nome, deploy.Namespace, deploy.Replicas)

	fmt.Println("Partiu escalar a infra")
	for i := 0; i < 4; i++ {
		Escalar(&deploy, 1)
	}
	fmt.Printf("Deploy config depois de escalar: \nNome: %s\nNamespace: %s\nReplicas: %d\n", deploy.Nome, deploy.Namespace, deploy.Replicas)


}
