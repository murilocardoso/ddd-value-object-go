package domainfiscalcondition

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

type Model struct {
	value string
}

const (
	consumidorFinal      string = "CONSUMIDOR_FINAL"
	monotributo          string = "MONOTRIBUTO"
	responsableInscripto string = "RESPONSABLE_INSCRIPTO"
	empty                string = ""
)

func ConsumidorFinal() Model {
	value, _ := new(consumidorFinal)
	return value
}

func Monotributo() Model {
	value, _ := new(monotributo)
	return value
}

func IVAReponsableInscripto() Model {
	value, _ := new(responsableInscripto)
	return value
}

func Empty() Model {
	value, _ := new(empty)
	return value
}

func (f *Model) Equals(fiscalCondition Model) bool {
	return f.value == fiscalCondition.value
}

func (f *Model) EqualsString(stringFiscalCondition string) (bool, error) {
	normalizedFiscalCondition, err := normalizeFiscalCondition(stringFiscalCondition)
	if err != nil {
		return false, err
	}

	return f.value == normalizedFiscalCondition, nil
}

func (f *Model) IsConsumidorFinal() bool {
	return f.value == consumidorFinal
}

func (f *Model) IsMonotributo() bool {
	return f.value == monotributo
}

func (f *Model) IsResponsableInscripto() bool {
	return f.value == responsableInscripto
}

func (f *Model) IsEmpty() bool {
	return f.value == empty
}

func NewFromString(aFiscalCondition string) (Model, error) {
	normalizedFiscalCondition, err := normalizeFiscalCondition(aFiscalCondition)
	if err != nil {
		return Model{}, err
	}

	return new(normalizedFiscalCondition)
}

func new(aFiscalCondition string) (Model, error) {
	fiscalCondition := Model{
		value: aFiscalCondition,
	}

	if !fiscalCondition.valid() {
		return Model{}, errors.New(fmt.Sprintf("Invalid fiscal condition %s", fiscalCondition))
	}

	return fiscalCondition, nil
}

func (f *Model) valid() bool {
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
