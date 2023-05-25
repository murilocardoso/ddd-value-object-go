package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Uma evolução para o nosso modelo seria criar um tipo FiscalCondition e definir constantes para os possíveis valores,
olhando muitos códigos go, essa abordagem é bastante usada
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition = domain.Monotributo

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fiscalCondition)
	fmt.Println(strings.Repeat("-", 30))
}
