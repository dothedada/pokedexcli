package main

import (
	"math"
	"testing"
)

func TestCatchProbability(t *testing.T) {
	cases := []struct {
		caseTxt        string
		input          int
		expectedOutput float64
	}{
		{
			caseTxt:        "minimum value",
			input:          36,
			expectedOutput: 0.9,
		},
		{
			caseTxt:        "maximum value",
			input:          608,
			expectedOutput: 0.05,
		},
	}

	for caseIndx, values := range cases {
		tolerance := 0.01
		result := catchProbability(values.input)

		if math.Abs(result-values.expectedOutput) > tolerance {
			t.Errorf("%d) %s expected: %f, returned: %f", caseIndx, values.caseTxt, values.expectedOutput, result)
		}

	}
}
