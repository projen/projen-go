package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IUseComponentExportOnlyModulesOptions interface {
	// Allows the export of constants.
	//
	// This option is for environments that support it, such as [Vite](https://vitejs.dev/)
	// Experimental.
	AllowConstantExport() *bool
	// Experimental.
	SetAllowConstantExport(a *bool)
	// A list of names that can be additionally exported from the module This option is for exports that do not hinder [React Fast Refresh](https://github.com/facebook/react/tree/main/packages/react-refresh), such as [`meta` in Remix](https://remix.run/docs/en/main/route/meta).
	// Experimental.
	AllowExportNames() *[]*string
	// Experimental.
	SetAllowExportNames(a *[]*string)
}

// The jsii proxy for IUseComponentExportOnlyModulesOptions
type jsiiProxy_IUseComponentExportOnlyModulesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IUseComponentExportOnlyModulesOptions) AllowConstantExport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowConstantExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseComponentExportOnlyModulesOptions)SetAllowConstantExport(val *bool) {
	_jsii_.Set(
		j,
		"allowConstantExport",
		val,
	)
}

func (j *jsiiProxy_IUseComponentExportOnlyModulesOptions) AllowExportNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowExportNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseComponentExportOnlyModulesOptions)SetAllowExportNames(val *[]*string) {
	_jsii_.Set(
		j,
		"allowExportNames",
		val,
	)
}

