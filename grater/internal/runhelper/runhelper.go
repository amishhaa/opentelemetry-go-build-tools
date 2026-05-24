// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package runhelper runs tests for all dependents.
package runhelper

import (
	"context"
	"fmt"

	"go.opentelemetry.io/build-tools/grater/internal/dockercontainer"
	"go.opentelemetry.io/build-tools/grater/internal/environment"
	"go.opentelemetry.io/build-tools/grater/internal/module"
	"go.opentelemetry.io/build-tools/grater/internal/report"
	"go.opentelemetry.io/build-tools/grater/internal/workspace"
)

// RunTestsForModule runs tests for all dependents and writes reports to the workspace.
func RunTestsForModule(ws *workspace.Workspace, mainModuleBase, mainModuleHead module.Module) ([]report.Result, error) {
	fmt.Println("called")
	ctx := context.Background()

	dc, err := dockercontainer.NewDockerContainer()
	if err != nil {
		return nil, err
	}

	env := environment.NewEnvironment(dc)

	fmt.Println(ws.GetDependents())
	results, err := env.RunTests(ctx, mainModuleBase, mainModuleHead, ws.GetDependents(), ws.GetReplacements())
	if err != nil {
		return nil, err
	}

	dependents := ws.GetDependents()

	if err = ws.WriteReport(report.GetReport(dependents, results)); err != nil {
		return nil, err
	}

	regressions := report.GetRegressionReport(dependents, results)

	if err = ws.WriteRegressionReport(regressions); err != nil {
		return nil, err
	}

	return regressions, nil
}