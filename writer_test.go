package colorprofile

import (
	"bytes"
	"io"
	"testing"
)

var writers = map[Profile]func(io.Writer) *Writer{
	TrueColor: func(w io.Writer) *Writer { return &Writer{w, TrueColor} },
	ANSI256:   func(w io.Writer) *Writer { return &Writer{w, ANSI256} },
	ANSI:      func(w io.Writer) *Writer { return &Writer{w, ANSI} },
	Ascii:     func(w io.Writer) *Writer { return &Writer{w, Ascii} },
	NoTTY:     func(w io.Writer) *Writer { return &Writer{w, NoTTY} },
}

var writer_cases = []struct {
	name              string
	input             string
	expectedTrueColor string
	expectedANSI256   string
	expectedANSI      string
	expectedAscii     string
	expectedNoTTY     string
}{
	{
		name: "empty",
	},
	{
		name:              "no styles",
		input:             "hello world",
		expectedTrueColor: "hello world",
		expectedANSI256:   "hello world",
		expectedANSI:      "hello world",
		expectedAscii:     "hello world",
		expectedNoTTY:     "hello world",
	},
	{
		name:              "simple style attributes",
		input:             "hello \x1b[1mworld\x1b[m",
		expectedTrueColor: "hello \x1b[1mworld\x1b[m",
		expectedANSI256:   "hello \x1b[1mworld\x1b[m",
		expectedANSI:      "hello \x1b[1mworld\x1b[m",
		expectedAscii:     "hello \x1b[1mworld\x1b[m",
		expectedNoTTY:     "hello world",
	},
	{
		name:              "simple basic color ",
		input:             "hello \x1b[31mworld\x1b[m",
		expectedTrueColor: "hello \x1b[31mworld\x1b[m",
		expectedANSI256:   "hello \x1b[31mworld\x1b[m",
		expectedANSI:      "hello \x1b[31mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
		expectedNoTTY:     "hello world",
	},
	{
		name:              "simple extended color",
		input:             "hello \x1b[38;5;196mworld\x1b[m",
		expectedTrueColor: "hello \x1b[38;5;196mworld\x1b[m",
		expectedANSI256:   "hello \x1b[38;5;196mworld\x1b[m",
		expectedANSI:      "hello \x1b[91mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
		expectedNoTTY:     "hello world",
	},
	{
		name:              "simple true color",
		input:             "hello \x1b[38;2;255;133;55mworld\x1b[m", // #ff8537
		expectedTrueColor: "hello \x1b[38;2;255;133;55mworld\x1b[m",
		expectedANSI256:   "hello \x1b[38;5;209mworld\x1b[m",
		expectedANSI:      "hello \x1b[91mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
		expectedNoTTY:     "hello world",
	},
	{
		name:              "simple background color",
		input:             "hello \x1b[48;5;196mworld\x1b[m",
		expectedTrueColor: "hello \x1b[48;5;196mworld\x1b[m",
		expectedANSI256:   "hello \x1b[48;5;196mworld\x1b[m",
		expectedANSI:      "hello \x1b[101mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
		expectedNoTTY:     "hello world",
	},
}

func TestWriter(t *testing.T) {
	for i, c := range writer_cases {
		for profile, writer := range writers {
			t.Run(c.name+"-"+profile.String(), func(t *testing.T) {
				var buf bytes.Buffer
				w := writer(&buf)
				_, err := w.Write([]byte(c.input))
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				var expected string
				switch profile {
				case TrueColor:
					expected = c.expectedTrueColor
				case ANSI256:
					expected = c.expectedANSI256
				case ANSI:
					expected = c.expectedANSI
				case Ascii:
					expected = c.expectedAscii
				case NoTTY:
					expected = c.expectedNoTTY
				}
				if got := buf.String(); got != expected {
					t.Errorf("case: %d, got: %q, expected: %q", i+1, got, expected)
				}
			})
		}
	}
}

func BenchmarkWriter(b *testing.B) {
	w := &Writer{&bytes.Buffer{}, ANSI}
	input := []byte("\x1b[1;3;59mhello\x1b[m \x1b[38;2;255;133;55mworld\x1b[m")
	for i := 0; i < b.N; i++ {
		_, _ = w.Write(input)
	}
}
