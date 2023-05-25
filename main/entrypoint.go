package main

import (
	"fmt"
	"strings"
)

func main() {
	var fiscalCondition string

	fiscalCondition = "Monotributo"

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fiscalCondition)
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
