package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// File representing the local NPM config in .npmrc.
// Experimental.
type NpmConfig interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// configure a generic property.
	// Experimental.
	AddConfig(name *string, value *string)
	// configure a scoped registry.
	// Experimental.
	AddRegistry(url *string, scope *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for NpmConfig
type jsiiProxy_NpmConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NpmConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNpmConfig(project NodeProject, options *NpmConfigOptions) NpmConfig {
	_init_.Initialize()

	j := jsiiProxy_NpmConfig{}

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNpmConfig_Override(n NpmConfig, project NodeProject, options *NpmConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		n,
	)
}

func (n *jsiiProxy_NpmConfig) AddConfig(name *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"addConfig",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NpmConfig) AddRegistry(url *string, scope *string) {
	_jsii_.InvokeVoid(
		n,
		"addRegistry",
		[]interface{}{url, scope},
	)
}

func (n *jsiiProxy_NpmConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

