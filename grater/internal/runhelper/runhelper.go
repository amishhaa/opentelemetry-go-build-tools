// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package runhelper provides helper functions for running commands.
package runhelper

import (
	"go.opentelemetry.io/build-tools/grater/internal/container"
	"go.opentelemetry.io/build-tools/grater/internal/dependent"
	"go.opentelemetry.io/build-tools/grater/internal/testenvironment"
	"go.opentelemetry.io/build-tools/grater/internal/workspace"
)


func runTests(ws workspace.Workspace, moduleName string, dependents []dependent.Dependent) {
	
	for _, dependent := range dependents {
		report.write(runTest(ws, moduleName, dependent))
	}
}

func runTest(moduleName string, dependent dependent.Dependent) {
	// call test enviornment helper to run test.
}
