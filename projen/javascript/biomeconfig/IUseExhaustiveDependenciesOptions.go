package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for the rule `useExhaustiveDependencies`.
// Experimental.
type IUseExhaustiveDependenciesOptions interface {
	// List of hooks of which the dependencies should be validated.
	// Experimental.
	Hooks() *[]IHook
	// Experimental.
	SetHooks(h *[]IHook)
	// Whether to report an error when a hook has no dependencies array.
	// Experimental.
	ReportMissingDependenciesArray() *bool
	// Experimental.
	SetReportMissingDependenciesArray(r *bool)
	// Whether to report an error when a dependency is listed in the dependencies array but isn't used.
	//
	// Defaults to true.
	// Experimental.
	ReportUnnecessaryDependencies() *bool
	// Experimental.
	SetReportUnnecessaryDependencies(r *bool)
}

// The jsii proxy for IUseExhaustiveDependenciesOptions
type jsiiProxy_IUseExhaustiveDependenciesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions) Hooks() *[]IHook {
	var returns *[]IHook
	_jsii_.Get(
		j,
		"hooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions)SetHooks(val *[]IHook) {
	_jsii_.Set(
		j,
		"hooks",
		val,
	)
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions) ReportMissingDependenciesArray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reportMissingDependenciesArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions)SetReportMissingDependenciesArray(val *bool) {
	_jsii_.Set(
		j,
		"reportMissingDependenciesArray",
		val,
	)
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions) ReportUnnecessaryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reportUnnecessaryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseExhaustiveDependenciesOptions)SetReportUnnecessaryDependencies(val *bool) {
	_jsii_.Set(
		j,
		"reportUnnecessaryDependencies",
		val,
	)
}

