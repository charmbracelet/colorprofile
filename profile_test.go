package colorprofile

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
)

func TestHexTo256(t *testing.T) {
	testCases := map[string]struct {
		input          colorful.Color
		expectedOutput ansi.ExtendedColor
	}{
		"white": {
			input:          colorful.Color{1, 1, 1},
			expectedOutput: 231,
		},
		"offwhite": {
			input:          colorful.Color{0.9333, 0.9333, 0.933},
			expectedOutput: 255,
		},
		"slightly brighter than offwhite": {
			input:          colorful.Color{0.95, 0.95, 0.95},
			expectedOutput: 255,
		},
		"red": {
			input:          colorful.Color{1, 0, 0},
			expectedOutput: 196,
		},
		"silver foil": {
			input:          colorful.Color{0.6863, 0.6863, 0.6863},
			expectedOutput: 145,
		},
		"silver chalice": {
			input:          colorful.Color{0.698, 0.698, 0.698},
			expectedOutput: 249,
		},
		"slightly closer to silver foil": {
			input:          colorful.Color{0.692, 0.692, 0.692},
			expectedOutput: 145,
		},
		"slightly closer to silver chalice": {
			input:          colorful.Color{0.694, 0.694, 0.694},
			expectedOutput: 249,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			output := hexToANSI256Color(testCase.input)
			if output != testCase.expectedOutput {
				t.Errorf("Expected truecolor %+v to map to 256 color %d, but instead received %d", testCase.input, testCase.expectedOutput, output)
			}
		})
	}
}
