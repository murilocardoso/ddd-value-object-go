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
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition, err := domain.NewFiscalConditionFromString("Monotributo")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
