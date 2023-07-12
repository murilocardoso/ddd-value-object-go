package domainfiscalcondition

import (
	errors "errors"
	"fmt"
)

type InvalidFiscalCondition struct {
	value string
}

func NewInvalidFiscalCondition(value string) InvalidFiscalCondition {
	return InvalidFiscalCondition{
		value: value,
	}
}

func IsInvalidFiscalCondition(err error) bool {
	var valueErrorChecker InvalidFiscalCondition
	return errors.As(err, &valueErrorChecker)
}

func (i InvalidFiscalCondition) Error() string {
	return fmt.Sprintf("Invalid fiscal condition %s", i.value)
}
