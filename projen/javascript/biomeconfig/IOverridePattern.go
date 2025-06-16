package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IOverridePattern interface {
	// Specific configuration for the Css language.
	// Experimental.
	Css() ICssConfiguration
	// Experimental.
	SetCss(c ICssConfiguration)
	// Specific configuration for the Json language.
	// Experimental.
	Formatter() IOverrideFormatterConfiguration
	// Experimental.
	SetFormatter(f IOverrideFormatterConfiguration)
	// Specific configuration for the Graphql language.
	// Experimental.
	Graphql() IGraphqlConfiguration
	// Experimental.
	SetGraphql(g IGraphqlConfiguration)
	// A list of Unix shell style patterns.
	//
	// The formatter will ignore files/folders that will match these patterns.
	// Experimental.
	Ignore() *[]*string
	// Experimental.
	SetIgnore(i *[]*string)
	// A list of Unix shell style patterns.
	//
	// The formatter will include files/folders that will match these patterns.
	// Experimental.
	Include() *[]*string
	// Experimental.
	SetInclude(i *[]*string)
	// Specific configuration for the JavaScript language.
	// Experimental.
	Javascript() IJavascriptConfiguration
	// Experimental.
	SetJavascript(j IJavascriptConfiguration)
	// Specific configuration for the Json language.
	// Experimental.
	Json() IJsonConfiguration
	// Experimental.
	SetJson(j IJsonConfiguration)
	// Specific configuration for the Json language.
	// Experimental.
	Linter() IOverrideLinterConfiguration
	// Experimental.
	SetLinter(l IOverrideLinterConfiguration)
	// Specific configuration for the Json language.
	// Experimental.
	OrganizeImports() IOverrideOrganizeImportsConfiguration
	// Experimental.
	SetOrganizeImports(o IOverrideOrganizeImportsConfiguration)
}

// The jsii proxy for IOverridePattern
type jsiiProxy_IOverridePattern struct {
	_ byte // padding
}

func (j *jsiiProxy_IOverridePattern) Css() ICssConfiguration {
	var returns ICssConfiguration
	_jsii_.Get(
		j,
		"css",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetCss(val ICssConfiguration) {
	_jsii_.Set(
		j,
		"css",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Formatter() IOverrideFormatterConfiguration {
	var returns IOverrideFormatterConfiguration
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetFormatter(val IOverrideFormatterConfiguration) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Graphql() IGraphqlConfiguration {
	var returns IGraphqlConfiguration
	_jsii_.Get(
		j,
		"graphql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetGraphql(val IGraphqlConfiguration) {
	_jsii_.Set(
		j,
		"graphql",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Javascript() IJavascriptConfiguration {
	var returns IJavascriptConfiguration
	_jsii_.Get(
		j,
		"javascript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetJavascript(val IJavascriptConfiguration) {
	_jsii_.Set(
		j,
		"javascript",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Json() IJsonConfiguration {
	var returns IJsonConfiguration
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetJson(val IJsonConfiguration) {
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) Linter() IOverrideLinterConfiguration {
	var returns IOverrideLinterConfiguration
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetLinter(val IOverrideLinterConfiguration) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

func (j *jsiiProxy_IOverridePattern) OrganizeImports() IOverrideOrganizeImportsConfiguration {
	var returns IOverrideOrganizeImportsConfiguration
	_jsii_.Get(
		j,
		"organizeImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverridePattern)SetOrganizeImports(val IOverrideOrganizeImportsConfiguration) {
	_jsii_.Set(
		j,
		"organizeImports",
		val,
	)
}

