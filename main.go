package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoBruno := ContaCorrente{titular: "Bruno Luiz", numeroAgencia: 589, numeroConta: 5978, saldo: 178.98}
	contaDaBianca := ContaCorrente{"Bianca Santos", 578, 5778, 149.67}

	fmt.Println(contaDoBruno)
	fmt.Println(contaDaBianca)

}
