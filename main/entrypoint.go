package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Outra necessidade comum é a comparação de igualdade entre objetos de valor.
Para resolver isso podemos pensar como javeiros rsrsrs e usar um método Equals.
Abre portas para outros tipos de comparação de igualdade, por exemplo EqualsString
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition, err := domain.NewFiscalConditionFromString("")
	otherFiscalCondition, err := domain.NewFiscalConditionFromString("Monotributo")
	equalsString, err := otherFiscalCondition.EqualsString("MONOTRIBUTO")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("IS EMPTY: %t", fiscalCondition.IsEmpty()))
	fmt.Println(fmt.Sprintf("EQUALS: %t", fiscalCondition.Equals(otherFiscalCondition)))
	fmt.Println(fmt.Sprintf("EQUALS STRING: %t", equalsString))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
