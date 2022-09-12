package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "deposito feito",
			c.saldo
	} else {
		return "deposito insuficiete",
			c.saldo
	}
}
func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 300}

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDoGustavo.saldo)
	status, valor := contaDaSilvia.Depositar(2000)
	fmt.Println(status)
	fmt.Println(valor)
}
