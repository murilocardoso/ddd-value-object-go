package domainfiscalcondition

import "fmt"

type InvalidFiscalCondition struct {
	invalidValue string
}

func NewInvalidFiscalCondition(value string) InvalidFiscalCondition {
	return InvalidFiscalCondition{
		invalidValue: value,
	}
}

func (i InvalidFiscalCondition) Error() string {
	return fmt.Sprintf("Invalid fiscal condition %s", i.invalidValue)
}
