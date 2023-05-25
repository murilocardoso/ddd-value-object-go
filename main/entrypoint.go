package main

import (
	"fmt"
	"strings"
)

/*
Olhando o exemplo de uma string solta parece um pouco forçado, mas usando dentro de uma estrutura eu diria que é bem
comum encontrar.
*/
type Taxpayer struct {
	FiscalCondition string
}

func main() {

	var taxpayer Taxpayer

	taxpayer.FiscalCondition = "Monotributo"

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(taxpayer.FiscalCondition)
	fmt.Println(strings.Repeat("-", 30))
}

/*
Modelando como uma simples string...
*/

/*
O problema dessa abordagem é que praticamente não existe tipagem, assim como coloquei Monotributo
poderia ter colocado Murilo. Ou seja, é um modelo que diz muito pouco a respeito do que é uma condição fiscal e
mais importante, não garante consistência
*/
