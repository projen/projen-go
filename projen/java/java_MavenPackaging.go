package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Configures a maven project to produce a .jar archive with sources and javadocs.
// Experimental.
type MavenPackaging interface {
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

// The jsii proxy struct for MavenPackaging
type jsiiProxy_MavenPackaging struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenPackaging) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenPackaging(project projen.Project, pom Pom, options *MavenPackagingOptions) MavenPackaging {
	_init_.Initialize()

	if err := validateNewMavenPackagingParameters(project, pom, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MavenPackaging{}

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenPackaging_Override(m MavenPackaging, project projen.Project, pom Pom, options *MavenPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		m,
	)
}

func (m *jsiiProxy_MavenPackaging) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenPackaging) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenPackaging) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

