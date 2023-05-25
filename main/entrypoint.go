package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"strings"
)

/*
Estamos quase lá, agora que temos uma estrutura por valor, ainda podemos ter um valor como FiscalCondition{},
traduzindo seria uma condição fiscal em branco, lidar com o null ou em branco sempre é um desafio, e se pararmos para
pensar, muitos objetos de valor podem ser vazios, a própria condição fiscal é um exemplo, podemos ter duas interpretações,
tratar o vazio como vazio mesmo ou assumir algum valor default. Independente da situação devemos tratar isso dentro
do nosso modelo de condição fiscal. Para fins de exemplo, vou considerar o vazio como vazio mesmo.
*/

func main() {
	var fiscalCondition domain.FiscalCondition

	fiscalCondition, err := domain.NewFiscalConditionFromString("")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("IS EMPTY: %t", fiscalCondition.IsEmpty()))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))
}
