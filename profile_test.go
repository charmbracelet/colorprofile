package colorprofile

import (
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lucasb-eyer/go-colorful"
)

func TestHexTo256(t *testing.T) {
	testCases := map[string]struct {
		input          colorful.Color
		expectedHex    string
		expectedOutput ansi.ExtendedColor
	}{
		"white": {
			input:          colorful.Color{R: 1, G: 1, B: 1},
			expectedHex:    "#ffffff",
			expectedOutput: 231,
		},
		"offwhite": {
			input:          colorful.Color{R: 0.9333, G: 0.9333, B: 0.933},
			expectedHex:    "#eeeeee",
			expectedOutput: 255,
		},
		"slightly brighter than offwhite": {
			input:          colorful.Color{R: 0.95, G: 0.95, B: 0.95},
			expectedHex:    "#f2f2f2",
			expectedOutput: 255,
		},
		"red": {
			input:          colorful.Color{R: 1, G: 0, B: 0},
			expectedHex:    "#ff0000",
			expectedOutput: 196,
		},
		"silver foil": {
			input:          colorful.Color{R: 0.6863, G: 0.6863, B: 0.6863},
			expectedHex:    "#afafaf",
			expectedOutput: 145,
		},
		"silver chalice": {
			input:          colorful.Color{R: 0.698, G: 0.698, B: 0.698},
			expectedHex:    "#b2b2b2",
			expectedOutput: 249,
		},
		"slightly closer to silver foil": {
			input:          colorful.Color{R: 0.692, G: 0.692, B: 0.692},
			expectedHex:    "#b0b0b0",
			expectedOutput: 145,
		},
		"slightly closer to silver chalice": {
			input:          colorful.Color{R: 0.694, G: 0.694, B: 0.694},
			expectedHex:    "#b1b1b1",
			expectedOutput: 249,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			// hex := fmt.Sprintf("#%02x%02x%02x", uint8(testCase.input.R*255), uint8(testCase.input.G*255), uint8(testCase.input.B*255))
			output := hexToANSI256Color(testCase.input)
			if testCase.input.Hex() != testCase.expectedHex {
				t.Errorf("Expected %+v to map to %s, but instead received %s", testCase.input, testCase.expectedHex, testCase.input.Hex())
			}
			if output != testCase.expectedOutput {
				t.Errorf("Expected truecolor %+v to map to 256 color %d, but instead received %d", testCase.input, testCase.expectedOutput, output)
			}
		})
	}
}

func TestDetectionByEnvironment(t *testing.T) {
	testCases := map[string]struct {
		environ  []string
		expected Profile
	}{
		"TERM is set to dumb": {
			environ:  []string{"TERM=dumb"},
			expected: NoTTY,
		},
		"TERM set to xterm": {
			environ:  []string{"TERM=xterm"},
			expected: ANSI,
		},
		"TERM is set to rio": {
			environ:  []string{"TERM=rio"},
			expected: TrueColor,
		},
		"TERM set to xterm-256color": {
			environ:  []string{"TERM=xterm-256color"},
			expected: ANSI256,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			profile := Env(testCase.environ)
			if profile != testCase.expected {
				t.Errorf("Expected profile to be %s, but instead received %s", testCase.expected, profile)
			}
		})
	}
}
