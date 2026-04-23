// // Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testenvironment

import (
	"go.opentelemetry.io/build-tools/grater/internal/container"
)

type testEnvironment struct {
	container        container.Container
	injectDependents []dependent.Dependent
	moduleDependents []dependent.Dependent
	volumeName      string
}

func NewTestEnvironment(container container.Container, injectDependents, moduleDependents []dependent.Dependent) *testEnvironment {
	// Create base env
	return &testEnvironment{
		container:        container,
		injectDependents: injectDependents,
		moduleDependents: moduleDependents,
		volumeName:      volumeName,
	}
}

func (te *testEnvironment) injectDependents() {
	for _, dependent := range te.injectDependents {
		// Inject depedents
	}
}

func (te *testEnvironment) cloneModuleInVolume() {
	// Clone the module into the volume
}

func (te *testEnvironment) RunTest() {
	// Calls inject deps on base te and executes test commands.
	// Returns report of tests that we ran
}
