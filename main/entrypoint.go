package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Uma possível forma de resolver é convencionar usar uma função que encapsule as validações dos valores válidos, nesse
caso seria uma função construtora do tipo FiscalCondition

Aqui há uma oportunidade de encapsular este regra de normalização, pode ser que alguns casos de uso venham
normalizado e em outros não, por não ser uma operação cara, pode-se encapsular e sempre normalizar antes de criar.

Pronto, agora conseguimos lidar também com valores desnormalizados e está encapsulado!
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition, err := domain.NewFiscalConditionFromString("Monotributo")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}

/*
Nesse caso, como Monotributo não está normalizado da forma esperada é retornado um erro.
Se passarmos o valor normalizado corretamente, não teremos problemas...
*/
