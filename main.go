package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDoBruna := ContaCorrente{"Bruna", 222, 111222, 55.5}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDoBruna)
}
