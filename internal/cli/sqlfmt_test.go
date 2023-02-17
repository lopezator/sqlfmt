package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/lopezator/sqlfmt/internal/cli"
)

func TestRunSQLFmt(t *testing.T) {
	var buf = new(bytes.Buffer)

	tests := map[string]struct {
		args  []string
		input string
		want  string
	}{
		"simple": {
			args:  []string{"--use-spaces"},
			input: "SElect * from foo",
			want:  "SELECT * FROM foo",
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			sqlfmtCmd := cli.NewSqlfmtCmd()

			sqlfmtCmd.SetIn(strings.NewReader(test.input))
			sqlfmtCmd.SetOut(buf)
			sqlfmtCmd.SetErr(buf)
			sqlfmtCmd.SetArgs(test.args)

			c, err := sqlfmtCmd.ExecuteC()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if c.Name() != "sqlfmt" {
				t.Errorf(`invalid command returned from ExecuteC: expected "set"', got: %q`, c.Name())
			}
			got := buf.String()[:len(buf.String())-1] // remove trailing newline

			if test.want != got {
				t.Errorf("Unexpected output: expected: %v, got: %v", test.want, got)
			}

			buf.Reset()
		})
	}
}
