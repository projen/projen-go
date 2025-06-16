package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Set of properties to integrate Biome with a VCS software.
// Experimental.
type IVcsConfiguration interface {
	// The kind of client.
	// Experimental.
	ClientKind() *string
	// Experimental.
	SetClientKind(c *string)
	// The main branch of the project.
	// Experimental.
	DefaultBranch() *string
	// Experimental.
	SetDefaultBranch(d *string)
	// Whether Biome should integrate itself with the VCS client.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// The folder where Biome should check for VCS files.
	//
	// By default, Biome will use the same folder where `biome.json` was found.
	//
	// If Biome can't find the configuration, it will attempt to use the current working directory. If no current working directory can't be found, Biome won't use the VCS integration, and a diagnostic will be emitted
	// Experimental.
	Root() *string
	// Experimental.
	SetRoot(r *string)
	// Whether Biome should use the VCS ignore file.
	//
	// When [true], Biome will ignore the files specified in the ignore file.
	// Experimental.
	UseIgnoreFile() *bool
	// Experimental.
	SetUseIgnoreFile(u *bool)
}

// The jsii proxy for IVcsConfiguration
type jsiiProxy_IVcsConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IVcsConfiguration) ClientKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVcsConfiguration)SetClientKind(val *string) {
	_jsii_.Set(
		j,
		"clientKind",
		val,
	)
}

func (j *jsiiProxy_IVcsConfiguration) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVcsConfiguration)SetDefaultBranch(val *string) {
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_IVcsConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVcsConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IVcsConfiguration) Root() *string {
	var returns *string
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVcsConfiguration)SetRoot(val *string) {
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_IVcsConfiguration) UseIgnoreFile() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useIgnoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVcsConfiguration)SetUseIgnoreFile(val *bool) {
	_jsii_.Set(
		j,
		"useIgnoreFile",
		val,
	)
}

