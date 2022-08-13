package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Implements JUnit-based testing.
// Experimental.
type Junit interface {
	projen.Component
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for Junit
type jsiiProxy_Junit struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Junit) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewJunit(project projen.Project, options *JunitOptions) Junit {
	_init_.Initialize()

	j := jsiiProxy_Junit{}

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJunit_Override(j Junit, project projen.Project, options *JunitOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		j,
	)
}

func (j *jsiiProxy_Junit) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Junit) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Junit) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

