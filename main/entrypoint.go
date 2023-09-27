package main

import (
	"fmt"
	"github.com/murilocardoso/ddd-value-object-go/domain"
	"github.com/murilocardoso/ddd-value-object-go/domain/domainfiscalcondition"
	"strings"
)

/*
Por fim, implementar um método simples de "Get" para poder pegar o valor, deixei esse método por último de forma
intencional, a ideia de existir métodos getters é por conta dos tráficos de conteúdos entre camadas, por exemplo,
para sair da camada de domínio e ir para o repositório, um getter vai ser necessário, e este é o ponto de atenção,
sempre que for necessário fazer get, devemos nos perguntar se não há alguma forma de pedir para que o objeto faça algo
ao invés de eu como client pegar o valor e aplicar alguma regra.
*/

func main() {
	var fiscalCondition domainfiscalcondition.Model
	fiscalCondition, err := domainfiscalcondition.NewFromString("Monotributo")

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(fmt.Sprintf("FISCAL_CONDITION: %s", fiscalCondition))
	fmt.Println(fmt.Sprintf("IS CONSUMIDOR FINAL: %t", fiscalCondition.IsConsumidorFinal()))
	fmt.Println(fmt.Sprintf("IS MONOTRIBUTO: %t", fiscalCondition.IsMonotributo()))
	fmt.Println(fmt.Sprintf("IS RESPONSABLE INSCRIPTO: %t", fiscalCondition.IsResponsableInscripto()))
	fmt.Println(fmt.Sprintf("IS EMPTY: %t", fiscalCondition.IsEmpty()))
	fmt.Println(fmt.Sprintf("VALUE: %s", fiscalCondition.Value()))
	fmt.Println(fmt.Sprintf("ERROR: %+v", err))
	fmt.Println(strings.Repeat("-", 30))

	domain.NewFiscalIdentityID()
	domain.NewFiscalIdentityIDFromString("")
}

type Taxpayer struct {
	Padrones PadronesByCodeAndPeriod
}

type PadronesByCodeAndPeriod map[PadronCode]map[PadronPeriod]Padron

type PadronCode struct {
	Value string
}

type PadronPeriod struct {
	Value string
}

type Padron struct {
	SomeValue string
}
