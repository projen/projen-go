package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithUseComponentExportOnlyModulesOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IUseComponentExportOnlyModulesOptions
	// Experimental.
	SetOptions(o IUseComponentExportOnlyModulesOptions)
}

// The jsii proxy for IRuleWithUseComponentExportOnlyModulesOptions
type jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions) Options() IUseComponentExportOnlyModulesOptions {
	var returns IUseComponentExportOnlyModulesOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions)SetOptions(val IUseComponentExportOnlyModulesOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

