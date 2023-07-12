package domainfiscalcondition_test

import (
	"github.com/murilocardoso/ddd-value-object-go/domain/domainfiscalcondition"
	"reflect"
	"testing"
)

func TestNewFromString(t *testing.T) {

	type expected struct {
		value string
		err   error
	}

	testCases := []struct {
		name     string
		input    string
		expected expected
	}{
		{
			name:  "Should get a consumidor final fiscal condition value from a description form",
			input: "Consumidor final",
			expected: expected{
				value: "CONSUMIDOR_FINAL",
				err:   nil,
			},
		},
		{
			name:  "Should get a consumidor final fiscal condition value from a normalized value",
			input: "CONSUMIDOR_FINAL",
			expected: expected{
				value: "CONSUMIDOR_FINAL",
				err:   nil,
			},
		},
		{
			name:  "Should fail to get a consumidor final fiscal condition when input is in a bad shape",
			input: "Consumidor  final",
			expected: expected{
				value: "",
				err:   domainfiscalcondition.InvalidFiscalCondition{},
			},
		},
		{
			name:  "Should get a monotributo fiscal condition value from a description form",
			input: "Monotributo",
			expected: expected{
				value: "MONOTRIBUTO",
				err:   nil,
			},
		},
		{
			name:  "Should get a monotributo fiscal condition value from a normalized value",
			input: "MONOTRIBUTO",
			expected: expected{
				value: "MONOTRIBUTO",
				err:   nil,
			},
		},
		{
			name:  "Should fail to get a monotributo fiscal condition when input is in a bad shape",
			input: "Monotibuto",
			expected: expected{
				value: "",
				err:   domainfiscalcondition.InvalidFiscalCondition{},
			},
		},
		{
			name:  "Should get a responsable inscripto fiscal condition value from a description form",
			input: "Responsable inscripto",
			expected: expected{
				value: "RESPONSABLE_INSCRIPTO",
				err:   nil,
			},
		},
		{
			name:  "Should get a responsable inscripto fiscal condition value from a normalized value",
			input: "RESPONSABLE_INSCRIPTO",
			expected: expected{
				value: "RESPONSABLE_INSCRIPTO",
				err:   nil,
			},
		},
		{
			name:  "Should fail to get a responsable inscripto fiscal condition when input is in a bad shape",
			input: "IsResponsableInscripto",
			expected: expected{
				value: "",
				err:   domainfiscalcondition.InvalidFiscalCondition{},
			},
		},
		{
			name:  "Should get an empty fiscal condition value",
			input: "",
			expected: expected{
				value: "",
				err:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := domainfiscalcondition.NewFromString(testCase.input)
			if (result.Value() != testCase.expected.value) ||
				reflect.TypeOf(err) != reflect.TypeOf(testCase.expected.err) {
				t.Errorf("NewFromString(%+v) got %+v %T, expected %+v %T", testCase.input, result.Value(),
					err, testCase.expected.value, testCase.expected.err)
			}
		})
	}
}

func TestConsumidorFinal(t *testing.T) {

	t.Run("Should get a consumidor final fiscal condition value", func(t *testing.T) {
		expected := "CONSUMIDOR_FINAL"
		result := domainfiscalcondition.ConsumidorFinal()
		if result.Value() != expected {
			t.Errorf("ConsumidorFinal() got %+v, expected %+v", result.Value(), expected)
		}
	})
}

func TestMonotributo(t *testing.T) {

	t.Run("Should get a monotributo fiscal condition value", func(t *testing.T) {
		expected := "MONOTRIBUTO"
		result := domainfiscalcondition.Monotributo()
		if result.Value() != expected {
			t.Errorf("Monotributo() got %+v, expected %+v", result.Value(), expected)
		}
	})
}

func TestResponsableInscripto(t *testing.T) {

	t.Run("Should get a responsable inscripto fiscal condition value", func(t *testing.T) {
		expected := "RESPONSABLE_INSCRIPTO"
		result := domainfiscalcondition.ResponsableInscripto()
		if result.Value() != expected {
			t.Errorf("IsResponsableInscripto() got %+v, expected %+v", result.Value(), expected)
		}
	})
}

func TestEmpty(t *testing.T) {

	t.Run("Should get an empty fiscal condition value", func(t *testing.T) {
		expected := ""
		result := domainfiscalcondition.Empty()
		if result.Value() != expected {
			t.Errorf("Empty() got %+v, expected %+v", result.Value(), expected)
		}
	})
}

func TestModel_IsConsumidorFinal(t *testing.T) {

	testCases := []struct {
		name     string
		input    domainfiscalcondition.Model
		expected bool
	}{
		{
			name:     "Should get true when fiscal condition is consumidor final",
			input:    domainfiscalcondition.ConsumidorFinal(),
			expected: true,
		},
		{
			name:     "Should get false when fiscal condition is monotributo",
			input:    domainfiscalcondition.Monotributo(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is responsable inscripto",
			input:    domainfiscalcondition.ResponsableInscripto(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is empty",
			input:    domainfiscalcondition.Empty(),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.IsConsumidorFinal()
			if result != testCase.expected {
				t.Errorf("IsConsumidorFinal() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}

func TestModel_IsMonotributo(t *testing.T) {

	testCases := []struct {
		name     string
		input    domainfiscalcondition.Model
		expected bool
	}{
		{
			name:     "Should get true when fiscal condition is monotributo",
			input:    domainfiscalcondition.Monotributo(),
			expected: true,
		},
		{
			name:     "Should get false when fiscal condition is consumidor final",
			input:    domainfiscalcondition.ConsumidorFinal(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is responsable inscripto",
			input:    domainfiscalcondition.ResponsableInscripto(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is empty",
			input:    domainfiscalcondition.Empty(),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.IsMonotributo()
			if result != testCase.expected {
				t.Errorf("IsMonotributo() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}

func TestModel_IsResponsableInscripto(t *testing.T) {

	testCases := []struct {
		name     string
		input    domainfiscalcondition.Model
		expected bool
	}{
		{
			name:     "Should get true when fiscal condition is responsable inscripto",
			input:    domainfiscalcondition.ResponsableInscripto(),
			expected: true,
		},
		{
			name:     "Should get false when fiscal condition is consumidor final",
			input:    domainfiscalcondition.ConsumidorFinal(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is monotributo",
			input:    domainfiscalcondition.Monotributo(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is empty",
			input:    domainfiscalcondition.Empty(),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.IsResponsableInscripto()
			if result != testCase.expected {
				t.Errorf("IsResponsableInscripto() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}

func TestModel_IsEmpty(t *testing.T) {

	testCases := []struct {
		name     string
		input    domainfiscalcondition.Model
		expected bool
	}{
		{
			name:     "Should get true when fiscal condition has no value",
			input:    domainfiscalcondition.Empty(),
			expected: true,
		},
		{
			name:     "Should get false when fiscal condition is consumidor final",
			input:    domainfiscalcondition.ConsumidorFinal(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is monotributo",
			input:    domainfiscalcondition.Monotributo(),
			expected: false,
		},
		{
			name:     "Should get false when fiscal condition is responsable inscripto",
			input:    domainfiscalcondition.ResponsableInscripto(),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.IsEmpty()
			if result != testCase.expected {
				t.Errorf("IsEmpty() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}

func TestModel_Equals(t *testing.T) {

	type input struct {
		oneFiscalCondition   domainfiscalcondition.Model
		otherFiscalCondition domainfiscalcondition.Model
	}

	testCases := []struct {
		name     string
		input    input
		expected bool
	}{
		{
			name: "Should be true when fiscal conditions are the same",
			input: input{
				oneFiscalCondition:   domainfiscalcondition.ConsumidorFinal(),
				otherFiscalCondition: domainfiscalcondition.ConsumidorFinal(),
			},
			expected: true,
		},
		{
			name: "Should be false when fiscal conditions are not the same",
			input: input{
				oneFiscalCondition:   domainfiscalcondition.ConsumidorFinal(),
				otherFiscalCondition: domainfiscalcondition.Monotributo(),
			},
			expected: false,
		},
		{
			name: "Should be false when fiscal conditions are not the same",
			input: input{
				oneFiscalCondition:   domainfiscalcondition.ConsumidorFinal(),
				otherFiscalCondition: domainfiscalcondition.Empty(),
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.oneFiscalCondition.Equals(testCase.input.otherFiscalCondition)
			if result != testCase.expected {
				t.Errorf("Equals() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}

func TestModel_EqualsString(t *testing.T) {
	type input struct {
		oneFiscalCondition           domainfiscalcondition.Model
		otherFiscalConditionAsString string
	}

	testCases := []struct {
		name     string
		input    input
		expected bool
	}{
		{
			name: "Should be true when string fiscal condition representation is equivalent",
			input: input{
				oneFiscalCondition:           domainfiscalcondition.ConsumidorFinal(),
				otherFiscalConditionAsString: "Consumidor final",
			},
			expected: true,
		},
		{
			name: "Should be true when string fiscal condition representation is equivalent",
			input: input{
				oneFiscalCondition:           domainfiscalcondition.ConsumidorFinal(),
				otherFiscalConditionAsString: "CONSUMIDOR_FINAL",
			},
			expected: true,
		},
		{
			name: "Should be false when string fiscal condition representation is not equivalent",
			input: input{
				oneFiscalCondition:           domainfiscalcondition.ConsumidorFinal(),
				otherFiscalConditionAsString: "Consumidor  final",
			},
			expected: false,
		},
		{
			name: "Should be false when string fiscal condition representation is not equivalent",
			input: input{
				oneFiscalCondition:           domainfiscalcondition.ConsumidorFinal(),
				otherFiscalConditionAsString: "Monotributo",
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.input.oneFiscalCondition.EqualsString(testCase.input.otherFiscalConditionAsString)
			if result != testCase.expected {
				t.Errorf("EqualsString() got %+v , expected %+v", result, testCase.expected)
			}
		})
	}
}
