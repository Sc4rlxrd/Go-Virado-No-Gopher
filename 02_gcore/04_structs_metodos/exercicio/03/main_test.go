package main

import "testing"

func TestEscalar(t *testing.T) {

	deploy := Deployment{
		RecursoK8s: RecursoK8s{
			Nome:      "Teste",
			Namespace: "Golang",
		},
		Replicas: 0,
	}

	for i := 0; i < 4; i++ {
		Escalar(&deploy, 1)
	}

	esperado := 4

	if deploy.Replicas != esperado {
		t.Errorf("Deploy config depois de escalar: \nNome: %s\nNamespace: %s\nReplicas: %d\n", deploy.Nome, deploy.Namespace, deploy.Replicas)
	}
}
