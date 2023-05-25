package domain

import (
	"errors"
	"fmt"
	"github.com/thoas/go-funk"
)

type FiscalCondition string

const (
	ConsumidorFinal      FiscalCondition = "CONSUMIDOR_FINAL"
	Monotributo          FiscalCondition = "MONOTRIBUTO"
	ResponsableInscripto FiscalCondition = "RESPONSABLE_INSCRIPTO"
)

func NewFiscalConditionFromString(aFiscalCondition string) (FiscalCondition, error) {

	fiscalCondition := FiscalCondition(aFiscalCondition)

	if !fiscalCondition.valid() {
		return "", errors.New(fmt.Sprintf("Invalid fiscal condition %s", fiscalCondition))
	}

	return fiscalCondition, nil
}

func (f *FiscalCondition) valid() bool {
	validValues := []string{string(ConsumidorFinal), string(Monotributo), string(ResponsableInscripto)}
	// go get github.com/thoas/go-funk
	return funk.Contains(validValues, string(*f))
}
