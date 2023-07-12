package domainfiscalcondition_test

import (
	"errors"
	"github.com/murilocardoso/ddd-value-object-go/domain/domainfiscalcondition"
	"testing"
)

func TestNewInvalidFiscalCondition(t *testing.T) {

	t.Run("Should get InvalidFiscalCondition", func(t *testing.T) {
		input := "InvalidValue"
		expected := "Invalid fiscal condition InvalidValue"
		result := domainfiscalcondition.NewInvalidFiscalCondition(input)
		if result.Error() != expected {
			t.Errorf("NewInvalidFiscalCondition(%+v) got %+v , expected %+v", input, result.Error(),
				expected)
		}
	})
}

func TestIsInvalidFiscalCondition(t *testing.T) {

	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "Should be true when error is InvalidFiscalCondition",
			err:      domainfiscalcondition.NewInvalidFiscalCondition("invalid"),
			expected: true,
		},
		{
			name:     "Should be false when error is not InvalidFiscalCondition",
			err:      errors.New("invalid"),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := domainfiscalcondition.IsInvalidFiscalCondition(testCase.err)
			if result != testCase.expected {
				t.Errorf("IsInvalidFiscalCondition(%q) got %+v , expected %+v", testCase.err, result,
					testCase.expected)
			}
		})
	}
}
