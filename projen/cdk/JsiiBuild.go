package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen/cdk/internal"
)

// A mixin that adds jsii compilation, multi-language packaging, and publishing capabilities to any TypeScript project.
//
// This implements the constructs `IMixin` interface and is applied using the
// `.with()` method on any construct.
//
// Example:
//   const project = new TypeScriptProject({ disableTsconfig: true, ... });
//   project.with(new JsiiBuild({
//     jsiiVersion: '~5.9.0',
//     publishToMaven: { ... },
//   }));
//
// Experimental.
type JsiiBuild interface {
	constructs.IMixin
	// Applies jsii configuration to the target TypeScriptProject.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Returns true if the construct is a TypeScriptProject.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for JsiiBuild
type jsiiProxy_JsiiBuild struct {
	internal.Type__constructsIMixin
}

// Experimental.
func NewJsiiBuild(options *JsiiBuildOptions, extraJobOptions *map[string]interface{}) JsiiBuild {
	_init_.Initialize()

	if err := validateNewJsiiBuildParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsiiBuild{}

	_jsii_.Create(
		"projen.cdk.JsiiBuild",
		[]interface{}{options, extraJobOptions},
		&j,
	)

	return &j
}

// Experimental.
func NewJsiiBuild_Override(j JsiiBuild, options *JsiiBuildOptions, extraJobOptions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.JsiiBuild",
		[]interface{}{options, extraJobOptions},
		j,
	)
}

func (j *jsiiProxy_JsiiBuild) ApplyTo(construct constructs.IConstruct) {
	if err := j.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"applyTo",
		[]interface{}{construct},
	)
}

func (j *jsiiProxy_JsiiBuild) Supports(construct constructs.IConstruct) *bool {
	if err := j.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		j,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

