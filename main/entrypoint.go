package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Um ponto acabou ficando para trás, tudo isso até o momento depende de uma convenção de usar um construtor, se alguém
ainda quiser burlar e associar um valor completamente inválido para umq condição fiscal é possível.
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition = "Murilo"

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	//fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
