package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain/domainfiscalcondition"
	"strings"
)

/*
Podemos também tipar o nosso erro, dessa forma quem for usar o objeto de valor poderá saber exatamente qual
erro ocorreu.
*/

func main() {
	var fiscalCondition domainfiscalcondition.Model

	fiscalCondition, err := domainfiscalcondition.NewFromString("Valor invalido")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("IS EMPTY: %t", fiscalCondition.IsEmpty()))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
