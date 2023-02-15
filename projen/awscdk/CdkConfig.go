package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Represents cdk.json file.
// Experimental.
type CdkConfig interface {
	projen.Component
	// Name of the cdk.out directory.
	// Experimental.
	Cdkout() *string
	// List of glob patterns to be excluded by CDK.
	// Experimental.
	Exclude() *[]*string
	// List of glob patterns to be included by CDK.
	// Experimental.
	Include() *[]*string
	// Represents the JSON file.
	// Experimental.
	Json() projen.JsonFile
	// Experimental.
	Project() projen.Project
	// Add excludes to `cdk.json`.
	// Experimental.
	AddExcludes(patterns ...*string)
	// Add includes to `cdk.json`.
	// Experimental.
	AddIncludes(patterns ...*string)
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

// The jsii proxy struct for CdkConfig
type jsiiProxy_CdkConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CdkConfig) Cdkout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Json() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkConfig(project projen.Project, options *CdkConfigOptions) CdkConfig {
	_init_.Initialize()

	if err := validateNewCdkConfigParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdkConfig{}

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkConfig_Override(c CdkConfig, project projen.Project, options *CdkConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		c,
	)
}

func (c *jsiiProxy_CdkConfig) AddExcludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addExcludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) AddIncludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addIncludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

