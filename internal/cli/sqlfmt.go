// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// TODO(mjibson): This subcommand has more flags than I would prefer. My
// goal is to have it have just -e and nothing else. gofmt initially started
// with tabs/spaces and width specifiers but later regretted that decision
// and removed them. I would like to get to the point where we achieve SQL
// formatting nirvana and we make this an opinionated formatter with few or
// zero options, hopefully before 2.1.

// This file contains definitions for data types suitable for use by
// the flag+pflag packages.

// statementsValue is an implementation of pflag.Value that appends any
// argument to a slice.
type statementsValue []string

// Type implements the pflag.Value interface.
func (s *statementsValue) Type() string { return "statementsValue" }

// String implements the pflag.Value interface.
func (s *statementsValue) String() string {
	return strings.Join(*s, ";")
}

// Set implements the pflag.Value interface.
func (s *statementsValue) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// sqlfmtCtx captures the command-line parameters of the `sqlfmt` command.
// Defaults set by InitCLIDefaults() above.
var sqlfmtCtx struct {
	len        int
	useSpaces  bool
	tabWidth   int
	noSimplify bool
	align      bool
	execStmts  statementsValue
}

// NewSqlfmtCmd create and returns the sqlfmt command
func NewSqlfmtCmd() *cobra.Command {
	sqlfmtCmd := &cobra.Command{
		Use:   "sqlfmt",
		Short: "format SQL statements",
		Long:  "Formats SQL statements from stdin to line length n.",
		RunE:  runSQLFmt,
	}

	cfg := tree.DefaultPrettyCfg()
	sqlfmtCmd.Flags().IntVar(&sqlfmtCtx.len, "len", cfg.LineWidth, "The line length where sqlfmt will try to wrap.")
	sqlfmtCmd.Flags().BoolVar(&sqlfmtCtx.useSpaces, "use-spaces", !cfg.UseTabs, "Indent with spaces instead of tabs.")
	sqlfmtCmd.Flags().IntVar(&sqlfmtCtx.tabWidth, "tab-width", cfg.TabWidth, "Number of spaces per indentation level.")
	sqlfmtCmd.Flags().BoolVar(&sqlfmtCtx.noSimplify, "no-simplify", !cfg.Simplify, "Don't simplify output.")
	sqlfmtCmd.Flags().BoolVar(&sqlfmtCtx.align, "align", cfg.Align != tree.PrettyNoAlign, "Align the output.")

	return sqlfmtCmd
}

func runSQLFmt(cmd *cobra.Command, args []string) error {
	if sqlfmtCtx.len < 1 {
		return errors.Errorf("line length must be > 0: %d", sqlfmtCtx.len)
	}
	if sqlfmtCtx.tabWidth < 1 {
		return errors.Errorf("tab width must be > 0: %d", sqlfmtCtx.tabWidth)
	}

	var sl parser.Statements
	if len(sqlfmtCtx.execStmts) != 0 {
		for _, exec := range sqlfmtCtx.execStmts {
			stmts, err := parser.Parse(exec)
			if err != nil {
				return err
			}
			sl = append(sl, stmts...)
		}
	} else {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		sl, err = parser.Parse(string(in))
		if err != nil {
			return err
		}
	}

	cfg := tree.DefaultPrettyCfg()
	cfg.UseTabs = !sqlfmtCtx.useSpaces
	cfg.LineWidth = sqlfmtCtx.len
	cfg.TabWidth = sqlfmtCtx.tabWidth
	cfg.Simplify = !sqlfmtCtx.noSimplify
	cfg.Align = tree.PrettyNoAlign
	if sqlfmtCtx.align {
		cfg.Align = tree.PrettyAlignAndDeindent
	}

	for i := range sl {
		fmt.Print(cfg.Pretty(sl[i].AST))
		if len(sl) > 1 {
			fmt.Print(";")
		}
		fmt.Println()
	}
	return nil
}
