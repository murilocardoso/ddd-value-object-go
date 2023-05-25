package domain

import (
	"errors"
	"fmt"
	"github.com/thoas/go-funk"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

type FiscalCondition string

const (
	consumidorFinal      FiscalCondition = "CONSUMIDOR_FINAL"
	monotributo          FiscalCondition = "MONOTRIBUTO"
	responsableInscripto FiscalCondition = "RESPONSABLE_INSCRIPTO"
)

func (f *FiscalCondition) IsConsumidorFinal() bool {
	return *f == consumidorFinal
}

func (f *FiscalCondition) IsMonotributo() bool {
	return *f == monotributo
}

func (f *FiscalCondition) IsResponsableInscripto() bool {
	return *f == responsableInscripto
}

func NewFiscalConditionFromString(aFiscalCondition string) (FiscalCondition, error) {
	normalizedFiscalCondition, err := normalizeFiscalCondition(aFiscalCondition)
	if err != nil {
		return "", err
	}

	return newFiscalCondition(normalizedFiscalCondition)
}

func newFiscalCondition(aFiscalCondition string) (FiscalCondition, error) {
	fiscalCondition := FiscalCondition(aFiscalCondition)

	if !fiscalCondition.valid() {
		return "", errors.New(fmt.Sprintf("Invalid fiscal condition %s", fiscalCondition))
	}

	return fiscalCondition, nil
}

func (f *FiscalCondition) valid() bool {
	validValues := []string{string(consumidorFinal), string(monotributo), string(responsableInscripto)}
	// go get github.com/thoas/go-funk
	return funk.Contains(validValues, string(*f))
}

func normalizeFiscalCondition(aFiscalCondition string) (string, error) {
	// go get golang.org/x/text/transform | golang.org/x/text/unicode/norm | golang.org/x/text/runes
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalizedFiscalCondition, _, err := transform.String(t, aFiscalCondition)
	if err != nil {
		return "", err
	}

	normalizedFiscalCondition = strings.ReplaceAll(normalizedFiscalCondition, " ", "_")
	normalizedFiscalCondition = strings.ToUpper(normalizedFiscalCondition)

	return normalizedFiscalCondition, err
}
