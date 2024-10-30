package colorprofile

import (
	"testing"
)

var cases = []struct {
	name     string
	environ  []string
	expected Profile
}{
	{
		name:     "empty",
		environ:  []string{},
		expected: NoTTY,
	},
	{
		name:     "no tty",
		environ:  []string{"TERM=dumb"},
		expected: NoTTY,
	},
	{
		name: "dumb term, truecolor, not forced",
		environ: []string{
			"TERM=dumb",
			"COLORTERM=truecolor",
		},
		expected: NoTTY,
	},
	{
		name: "dumb term, truecolor, forced",
		environ: []string{
			"TERM=dumb",
			"COLORTERM=truecolor",
			"CLICOLOR_FORCE=1",
		},
		expected: TrueColor,
	},
	{
		name: "dumb term, CLICOLOR_FORCE=1",
		environ: []string{
			"TERM=dumb",
			"CLICOLOR_FORCE=1",
		},
		expected: ANSI,
	},
	{
		name: "dumb term, CLICOLOR=1",
		environ: []string{
			"TERM=dumb",
			"CLICOLOR=1",
		},
		expected: NoTTY,
	},
	{
		name: "xterm-256color",
		environ: []string{
			"TERM=xterm-256color",
		},
		expected: ANSI256,
	},
	{
		name: "xterm-256color, CLICOLOR=1",
		environ: []string{
			"TERM=xterm-256color",
			"CLICOLOR=1",
		},
		expected: ANSI256,
	},
	{
		name: "xterm-256color, COLORTERM=yes",
		environ: []string{
			"TERM=xterm-256color",
			"COLORTERM=yes",
		},
		expected: TrueColor,
	},
	{
		name: "xterm-256color, NO_COLOR=1",
		environ: []string{
			"TERM=xterm-256color",
			"NO_COLOR=1",
		},
		expected: Ascii,
	},
	{
		name: "xterm",
		environ: []string{
			"TERM=xterm",
		},
		expected: Ascii,
	},
	{
		name: "xterm, NO_COLOR=1",
		environ: []string{
			"TERM=xterm",
			"NO_COLOR=1",
		},
		expected: Ascii,
	},
	{
		name: "xterm, CLICOLOR=1",
		environ: []string{
			"TERM=xterm",
			"CLICOLOR=1",
		},
		expected: ANSI,
	},
	{
		name: "xterm, CLICOLOR_FORCE=1",
		environ: []string{
			"TERM=xterm",
			"CLICOLOR_FORCE=1",
		},
		expected: ANSI,
	},
	{
		name: "xterm-16color",
		environ: []string{
			"TERM=xterm-16color",
		},
		expected: ANSI,
	},
	{
		name: "xterm-color",
		environ: []string{
			"TERM=xterm-color",
		},
		expected: ANSI,
	},
	{
		name: "xterm-256color, NO_COLOR=1, CLICOLOR_FORCE=1",
		environ: []string{
			"TERM=xterm-256color",
			"NO_COLOR=1",
			"CLICOLOR_FORCE=1",
		},
		expected: Ascii,
	},
	{
		name: "Windows Terminal",
		environ: []string{
			"WT_SESSION=1",
		},
		expected: TrueColor,
	},
}

func TestEnvColorProfile(t *testing.T) {
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Env(tc.environ)
			if p != tc.expected {
				t.Errorf("case %d: expected %v, got %v", i, tc.expected, p)
			}
		})
	}
}
