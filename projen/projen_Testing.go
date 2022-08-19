// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// A Testing static class with a .synth helper for getting a snapshots of construct outputs. Useful for snapshot testing with Jest.
//
// Example:
//   `expect(Testing.synth(someProject)).toMatchSnapshot()`
//
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Produces a simple JS object that represents the contents of the projects with field names being file paths.
//
// Returns: : any }.
// Experimental.
func Testing_Synth(project Project, options *SnapshotOptions) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"projen.Testing",
		"synth",
		[]interface{}{project, options},
		&returns,
	)

	return returns
}

