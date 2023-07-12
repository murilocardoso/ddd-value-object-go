package domainfiscalcondition

import (
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

// Value returns the fiscal condition value
func (f Model) Value() string {
	return f.value
}

// ConsumidorFinal returns a consumidor final fiscal condition
func ConsumidorFinal() Model {
	value, _ := newFiscalCondition(consumidorFinal)
	return value
}

// Monotributo returns a monotributo fiscal condition
func Monotributo() Model {
	value, _ := newFiscalCondition(monotributo)
	return value
}

// ResponsableInscripto returns a responsable inscripto fiscal condition
func ResponsableInscripto() Model {
	value, _ := newFiscalCondition(responsableInscripto)
	return value
}

// Empty returns an empty fiscal condition
func Empty() Model {
	value, _ := newFiscalCondition(empty)
	return value
}

// Equals return if the fiscal condition is equal or not
func (f Model) Equals(fiscalCondition Model) bool {
	return f.value == fiscalCondition.value
}

// EqualsString return if the string fiscal condition value are equivalent or not
func (f Model) EqualsString(stringFiscalCondition string) bool {
	normalizedFiscalCondition, err := normalizeFiscalCondition(stringFiscalCondition)
	if err != nil {
		return false
	}

	return f.value == normalizedFiscalCondition
}

// IsConsumidorFinal returns if the fiscal condition is consumidor final or not
func (f Model) IsConsumidorFinal() bool {
	return f.value == consumidorFinal
}

// IsMonotributo returns if the fiscal condition is monotributo or not
func (f Model) IsMonotributo() bool {
	return f.value == monotributo
}

// IsResponsableInscripto returns if the fiscal condition is responsable inscripto or not
func (f Model) IsResponsableInscripto() bool {
	return f.value == responsableInscripto
}

// IsEmpty returns if the fiscal condition is empty or not
func (f Model) IsEmpty() bool {
	return f.value == empty
}

// NewFromString returns a new fiscal condition
func NewFromString(aFiscalCondition string) (Model, error) {
	normalizedFiscalCondition, err := normalizeFiscalCondition(aFiscalCondition)
	if err != nil {
		return Empty(), err
	}

	return newFiscalCondition(normalizedFiscalCondition)
}

func newFiscalCondition(aFiscalCondition string) (Model, error) {
	fiscalCondition := Model{
		value: aFiscalCondition,
	}

	if !fiscalCondition.valid() {
		return Empty(), NewInvalidFiscalCondition(aFiscalCondition)
	}

	return fiscalCondition, nil
}

func (f Model) valid() bool {
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
