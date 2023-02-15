package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk8s/internal"
)

// Manages dependencies on the CDK8s.
// Experimental.
type Cdk8sDeps interface {
	projen.Component
	// The major version of the CDK8s (e.g. 1, 2, ...).
	// Experimental.
	Cdk8sMajorVersion() *float64
	// The minimum version of the CDK8s (e.g. `2.0.0`).
	// Experimental.
	Cdk8sMinimumVersion() *string
	// The dependency requirement for CDK8s.
	// Experimental.
	Cdk8sVersion() *string
	// Experimental.
	Project() projen.Project
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *Cdk8sPackageNames
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

// The jsii proxy struct for Cdk8sDeps
type jsiiProxy_Cdk8sDeps struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Cdk8sDeps) Cdk8sMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdk8sMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDeps) Cdk8sMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdk8sMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDeps) Cdk8sVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdk8sVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDeps) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdk8sDeps_Override(c Cdk8sDeps, project projen.Project, options *Cdk8sDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk8s.Cdk8sDeps",
		[]interface{}{project, options},
		c,
	)
}

func (c *jsiiProxy_Cdk8sDeps) PackageNames() *Cdk8sPackageNames {
	var returns *Cdk8sPackageNames

	_jsii_.Invoke(
		c,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdk8sDeps) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sDeps) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sDeps) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

