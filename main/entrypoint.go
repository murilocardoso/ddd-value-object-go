package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Ainda dependemos da convenção de sempre usar um construtor e também temos outros pontos...
Seguindo as boas práticas para modelagem orientada a objetos, os valores estão expostos quando deveríamos encapsulá-los
e apenas expor métodos para que os clientes possam fazer o que queiram.

Ao encapsular os valores representados pelas constantes parece que perdemos algumas possibilidades, por exemplo?
(i) Como posso saber se uma determinada condição fiscal é Monotributo ou é Consumidor Final?
(i.i) Uma forma é criar os dois tipos e compará-los, mas nesse exemplo quero apenas saber se um determinado valor é
monotributo ou não.

Para resolver usando os famosos IsSomething()...
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition, err := domain.NewFiscalConditionFromString("Monotributo")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}

/*
Os atributos estão encapsulados e a forma que estes valores são usados internamente é uma coisa e os métodos expostos
é outra. Inclusive, depender de um método ou função é muito melhor que depender de um valor (mesmo que seja uma constante)
Em um método ou função o comportamento é encapsulado e uma eventual tratativa pode ser realizada.
*/
