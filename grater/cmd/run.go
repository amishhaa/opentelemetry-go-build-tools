// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Package cmd provides the CLI commands for grater.
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/build-tools/grater/internal/module"
	"go.opentelemetry.io/build-tools/grater/internal/runhelper"
	"go.opentelemetry.io/build-tools/grater/internal/workspace"
)

// runCmd represents the run command.
func runCmd(ws *workspace.Workspace) *cobra.Command {
	return &cobra.Command{
		Use:   "run [base@version] [head@version]",
		Short: "Runs tests for all dependents.",
		Long:  "Runs tests for all dependents of the module and writes reports to the workspace.",
		Example: `
grater run go.opentelemetry.io/otel@v1.23.0 go.opentelemetry.io/otel@v1.24.0
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			basePath, baseVersion, _ := strings.Cut(args[0], "@")
			mainModuleBase := *module.NewModule(basePath, baseVersion)

			headPath, headVersion, _ := strings.Cut(args[1], "@")
			mainModuleHead := *module.NewModule(headPath, headVersion)

			regressions, err := runhelper.RunTestsForModule(ws, mainModuleBase, mainModuleHead)
			if err != nil {
				return err
			}
			fmt.Println(ws.GetDependents())
			fmt.Println("Called")
			if len(regressions) > 0 {
				cmd.Printf("Regressions found:\n")
				for _, r := range regressions {
					cmd.Printf("  %s: %s\n", r.Dependent, r.Status)
				}
				return fmt.Errorf("%d regression(s) found", len(regressions))
			}

			cmd.Printf("All tests passed.\n")
			return nil
		},
	}
}