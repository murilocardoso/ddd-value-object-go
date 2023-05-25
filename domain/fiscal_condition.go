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

type FiscalCondition struct {
	value string
}

const (
	consumidorFinal      string = "CONSUMIDOR_FINAL"
	monotributo          string = "MONOTRIBUTO"
	responsableInscripto string = "RESPONSABLE_INSCRIPTO"
	empty                string = ""
)

func (f *FiscalCondition) IsConsumidorFinal() bool {
	return f.value == consumidorFinal
}

func (f *FiscalCondition) IsMonotributo() bool {
	return f.value == monotributo
}

func (f *FiscalCondition) IsResponsableInscripto() bool {
	return f.value == responsableInscripto
}

func (f *FiscalCondition) IsEmpty() bool {
	return f.value == empty
}

func NewFiscalConditionFromString(aFiscalCondition string) (FiscalCondition, error) {
	normalizedFiscalCondition, err := normalizeFiscalCondition(aFiscalCondition)
	if err != nil {
		return FiscalCondition{}, err
	}

	return newFiscalCondition(normalizedFiscalCondition)
}

func newFiscalCondition(aFiscalCondition string) (FiscalCondition, error) {
	fiscalCondition := FiscalCondition{
		value: aFiscalCondition,
	}

	if !fiscalCondition.valid() {
		return FiscalCondition{}, errors.New(fmt.Sprintf("Invalid fiscal condition %s", fiscalCondition))
	}

	return fiscalCondition, nil
}

func (f *FiscalCondition) valid() bool {
	validValues := []string{consumidorFinal, monotributo, responsableInscripto, empty}
	// go get github.com/thoas/go-funk
	return funk.Contains(validValues, f.value)
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
