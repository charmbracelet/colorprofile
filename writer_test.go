package colorprofile

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/charmbracelet/x/ansi"
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
	},
	{
		name:              "simple style attributes",
		input:             "hello \x1b[1mworld\x1b[m",
		expectedTrueColor: "hello \x1b[1mworld\x1b[m",
		expectedANSI256:   "hello \x1b[1mworld\x1b[m",
		expectedANSI:      "hello \x1b[1mworld\x1b[m",
		expectedAscii:     "hello \x1b[1mworld\x1b[m",
	},
	{
		name:              "simple ansi color fg",
		input:             "hello \x1b[31mworld\x1b[m",
		expectedTrueColor: "hello \x1b[31mworld\x1b[m",
		expectedANSI256:   "hello \x1b[31mworld\x1b[m",
		expectedANSI:      "hello \x1b[31mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
	},
	{
		name:              "default fg color after ansi color",
		input:             "\x1b[31mhello \x1b[39mworld\x1b[m",
		expectedTrueColor: "\x1b[31mhello \x1b[39mworld\x1b[m",
		expectedANSI256:   "\x1b[31mhello \x1b[39mworld\x1b[m",
		expectedANSI:      "\x1b[31mhello \x1b[39mworld\x1b[m",
		expectedAscii:     "\x1b[mhello \x1b[mworld\x1b[m",
	},
	{
		name:              "ansi color fg and bg",
		input:             "\x1b[31;42mhello world\x1b[m",
		expectedTrueColor: "\x1b[31;42mhello world\x1b[m",
		expectedANSI256:   "\x1b[31;42mhello world\x1b[m",
		expectedANSI:      "\x1b[31;42mhello world\x1b[m",
		expectedAscii:     "\x1b[mhello world\x1b[m",
	},
	{
		name:              "bright ansi color fg and bg",
		input:             "\x1b[91;102mhello world\x1b[m",
		expectedTrueColor: "\x1b[91;102mhello world\x1b[m",
		expectedANSI256:   "\x1b[91;102mhello world\x1b[m",
		expectedANSI:      "\x1b[91;102mhello world\x1b[m",
		expectedAscii:     "\x1b[mhello world\x1b[m",
	},
	{
		name:              "simple 256 color fg",
		input:             "hello \x1b[38;5;196mworld\x1b[m",
		expectedTrueColor: "hello \x1b[38;5;196mworld\x1b[m",
		expectedANSI256:   "hello \x1b[38;5;196mworld\x1b[m",
		expectedANSI:      "hello \x1b[91mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
	},
	{
		name:              "256 color bg",
		input:             "\x1b[48;5;196mhello world\x1b[m",
		expectedTrueColor: "\x1b[48;5;196mhello world\x1b[m",
		expectedANSI256:   "\x1b[48;5;196mhello world\x1b[m",
		expectedANSI:      "\x1b[101mhello world\x1b[m",
		expectedAscii:     "\x1b[mhello world\x1b[m",
	},
	{
		name:              "simple true color bg",
		input:             "hello \x1b[38;2;255;133;55mworld\x1b[m", // #ff8537
		expectedTrueColor: "hello \x1b[38;2;255;133;55mworld\x1b[m",
		expectedANSI256:   "hello \x1b[38;5;209mworld\x1b[m",
		expectedANSI:      "hello \x1b[91mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
	},
	{
		name:              "itu true color bg",
		input:             "hello \x1b[38:2::255:133:55mworld\x1b[m", // #ff8537
		expectedTrueColor: "hello \x1b[38:2::255:133:55mworld\x1b[m",
		expectedANSI256:   "hello \x1b[38;5;209mworld\x1b[m",
		expectedANSI:      "hello \x1b[91mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
	},
	{
		name:              "simple ansi 256 color bg",
		input:             "hello \x1b[48:5:196mworld\x1b[m",
		expectedTrueColor: "hello \x1b[48:5:196mworld\x1b[m",
		expectedANSI256:   "hello \x1b[48;5;196mworld\x1b[m",
		expectedANSI:      "hello \x1b[101mworld\x1b[m",
		expectedAscii:     "hello \x1b[mworld\x1b[m",
	},
	{
		name:              "simple missing param",
		input:             "\x1b[31mhello \x1b[;1mworld",
		expectedTrueColor: "\x1b[31mhello \x1b[;1mworld",
		expectedANSI256:   "\x1b[31mhello \x1b[;1mworld",
		expectedANSI:      "\x1b[31mhello \x1b[;1mworld",
		expectedAscii:     "\x1b[mhello \x1b[;1mworld",
	},
	{
		name:              "color with other attributes",
		input:             "\x1b[1;38;5;204mhello \x1b[38;5;204mworld\x1b[m",
		expectedTrueColor: "\x1b[1;38;5;204mhello \x1b[38;5;204mworld\x1b[m",
		expectedANSI256:   "\x1b[1;38;5;204mhello \x1b[38;5;204mworld\x1b[m",
		expectedANSI:      "\x1b[1;91mhello \x1b[91mworld\x1b[m",
		expectedAscii:     "\x1b[1mhello \x1b[mworld\x1b[m",
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
					expected = ansi.Strip(c.input)
				}
				if got := buf.String(); got != expected {
					t.Errorf("case: %d, got: %q, expected: %q", i+1, got, expected)
				}
			})
		}
	}
}

func TestNewWriterPanic(t *testing.T) {
	_ = NewWriter(io.Discard, []string{"TERM=dumb"})
}

func TestNewWriterOsEnviron(t *testing.T) {
	w := NewWriter(io.Discard, os.Environ())
	if w.Profile != NoTTY {
		t.Errorf("expected NoTTY, got %v", w.Profile)
	}
}

func BenchmarkWriter(b *testing.B) {
	w := &Writer{&bytes.Buffer{}, ANSI}
	input := []byte("\x1b[1;3;59mhello\x1b[m \x1b[38;2;255;133;55mworld\x1b[m")
	for i := 0; i < b.N; i++ {
		_, _ = w.Write(input)
	}
}
