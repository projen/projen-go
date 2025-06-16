package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The configuration that is contained inside the file `biome.json`.
// Experimental.
type IConfiguration interface {
	// Specific configuration for assists.
	// Experimental.
	Assists() IAssistsConfiguration
	// Experimental.
	SetAssists(a IAssistsConfiguration)
	// Specific configuration for the Css language.
	// Experimental.
	Css() ICssConfiguration
	// Experimental.
	SetCss(c ICssConfiguration)
	// A list of paths to other JSON files, used to extends the current configuration.
	// Experimental.
	Extends() *[]*string
	// Experimental.
	SetExtends(e *[]*string)
	// The configuration of the filesystem.
	// Experimental.
	Files() IFilesConfiguration
	// Experimental.
	SetFiles(f IFilesConfiguration)
	// The configuration of the formatter.
	// Experimental.
	Formatter() IFormatterConfiguration
	// Experimental.
	SetFormatter(f IFormatterConfiguration)
	// Specific configuration for the GraphQL language.
	// Experimental.
	Graphql() IGraphqlConfiguration
	// Experimental.
	SetGraphql(g IGraphqlConfiguration)
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
	// The configuration for the linter.
	// Experimental.
	Linter() ILinterConfiguration
	// Experimental.
	SetLinter(l ILinterConfiguration)
	// The configuration of the import sorting.
	// Experimental.
	OrganizeImports() IOrganizeImports
	// Experimental.
	SetOrganizeImports(o IOrganizeImports)
	// A list of granular patterns that should be applied only to a sub set of files.
	// Experimental.
	Overrides() *[]IOverridePattern
	// Experimental.
	SetOverrides(o *[]IOverridePattern)
	// The configuration of the VCS integration.
	// Experimental.
	Vcs() IVcsConfiguration
	// Experimental.
	SetVcs(v IVcsConfiguration)
}

// The jsii proxy for IConfiguration
type jsiiProxy_IConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IConfiguration) Assists() IAssistsConfiguration {
	var returns IAssistsConfiguration
	_jsii_.Get(
		j,
		"assists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetAssists(val IAssistsConfiguration) {
	_jsii_.Set(
		j,
		"assists",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Css() ICssConfiguration {
	var returns ICssConfiguration
	_jsii_.Get(
		j,
		"css",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetCss(val ICssConfiguration) {
	_jsii_.Set(
		j,
		"css",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Extends() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extends",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetExtends(val *[]*string) {
	_jsii_.Set(
		j,
		"extends",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Files() IFilesConfiguration {
	var returns IFilesConfiguration
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetFiles(val IFilesConfiguration) {
	_jsii_.Set(
		j,
		"files",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Formatter() IFormatterConfiguration {
	var returns IFormatterConfiguration
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetFormatter(val IFormatterConfiguration) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Graphql() IGraphqlConfiguration {
	var returns IGraphqlConfiguration
	_jsii_.Get(
		j,
		"graphql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetGraphql(val IGraphqlConfiguration) {
	_jsii_.Set(
		j,
		"graphql",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Javascript() IJavascriptConfiguration {
	var returns IJavascriptConfiguration
	_jsii_.Get(
		j,
		"javascript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetJavascript(val IJavascriptConfiguration) {
	_jsii_.Set(
		j,
		"javascript",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Json() IJsonConfiguration {
	var returns IJsonConfiguration
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetJson(val IJsonConfiguration) {
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Linter() ILinterConfiguration {
	var returns ILinterConfiguration
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetLinter(val ILinterConfiguration) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) OrganizeImports() IOrganizeImports {
	var returns IOrganizeImports
	_jsii_.Get(
		j,
		"organizeImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetOrganizeImports(val IOrganizeImports) {
	_jsii_.Set(
		j,
		"organizeImports",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Overrides() *[]IOverridePattern {
	var returns *[]IOverridePattern
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetOverrides(val *[]IOverridePattern) {
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_IConfiguration) Vcs() IVcsConfiguration {
	var returns IVcsConfiguration
	_jsii_.Get(
		j,
		"vcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguration)SetVcs(val IVcsConfiguration) {
	_jsii_.Set(
		j,
		"vcs",
		val,
	)
}

