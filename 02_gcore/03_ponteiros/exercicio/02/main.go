package  main


import(
	"fmt"
	
)

type PodK8s struct {
	Nome     string
	Replicas int
}

func EscalarPod(p *PodK8s, novasReplicas int) error {

	if p == nil {
		return fmt.Errorf("pod não pode ser nil")
	}

	p.Replicas = novasReplicas
	return nil
}

func main(){

	fmt.Println("---- Exercício 02 -----")
	podK8s := PodK8s{

		Nome:     "BookCommerce",
		Replicas: 3,
	}
	fmt.Println("Pod atual: ", podK8s.Replicas)
	err := EscalarPod(&podK8s, 66)
	if err != nil {
		fmt.Println("Erro ao escalar:", err)
		return
	}
	fmt.Println("Pod depois de escalado: ", podK8s.Replicas)
}
