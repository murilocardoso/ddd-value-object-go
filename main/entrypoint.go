package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Outra necessidade é a de construtores para valores específicos, em um determinado caso de uso talvez se deseje uma
condição fiscal específica, por exemplo Responsable Inscirpto... Então vale encapsular essa regra de como criar uma
condição fiscal com este valor específico.
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition = domain.Monotributo()

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("IS EMPTY: %t", fiscalCondition.IsEmpty()))
	// fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
