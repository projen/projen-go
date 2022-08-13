// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// The `Dependencies` component is responsible to track the list of dependencies a project has, and then used by project types as the model for rendering project-specific dependency manifests such as the dependencies section `package.json` files.
//
// To add a dependency you can use a project-type specific API such as
// `nodeProject.addDeps()` or use the generic API of `project.deps`:
// Experimental.
type Dependencies interface {
	Component
	// A copy of all dependencies recorded for this project.
	//
	// The list is sorted by type->name->version.
	// Experimental.
	All() *[]*Dependency
	// Experimental.
	Project() Project
	// Adds a dependency to this project.
	// Experimental.
	AddDependency(spec *string, type_ DependencyType, metadata *map[string]interface{}) *Dependency
	// Returns a dependency by name.
	//
	// Fails if there is no dependency defined by that name or if `type` is not
	// provided and there is more then one dependency type for this dependency.
	//
	// Returns: a copy (cannot be modified).
	// Experimental.
	GetDependency(name *string, type_ DependencyType) *Dependency
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes a dependency.
	// Experimental.
	RemoveDependency(name *string, type_ DependencyType)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a dependency by name.
	//
	// Returns `undefined` if there is no dependency defined by that name or if
	// `type` is not provided and there is more then one dependency type for this
	// dependency.
	//
	// Returns: a copy (cannot be modified) or undefined if there is no match.
	// Experimental.
	TryGetDependency(name *string, type_ DependencyType) *Dependency
}

// The jsii proxy struct for Dependencies
type jsiiProxy_Dependencies struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Dependencies) All() *[]*Dependency {
	var returns *[]*Dependency
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependencies) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Adds a dependencies component to the project.
// Experimental.
func NewDependencies(project Project) Dependencies {
	_init_.Initialize()

	j := jsiiProxy_Dependencies{}

	_jsii_.Create(
		"projen.Dependencies",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Adds a dependencies component to the project.
// Experimental.
func NewDependencies_Override(d Dependencies, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Dependencies",
		[]interface{}{project},
		d,
	)
}

// Returns the coordinates of a dependency spec.
//
// Given `foo@^3.4.0` returns `{ name: "foo", version: "^3.4.0" }`.
// Given `bar@npm:@bar/legacy` returns `{ name: "bar", version: "npm:@bar/legacy" }`.
// Experimental.
func Dependencies_ParseDependency(spec *string) *DependencyCoordinates {
	_init_.Initialize()

	var returns *DependencyCoordinates

	_jsii_.StaticInvoke(
		"projen.Dependencies",
		"parseDependency",
		[]interface{}{spec},
		&returns,
	)

	return returns
}

func Dependencies_MANIFEST_FILE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Dependencies",
		"MANIFEST_FILE",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Dependencies) AddDependency(spec *string, type_ DependencyType, metadata *map[string]interface{}) *Dependency {
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"addDependency",
		[]interface{}{spec, type_, metadata},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) GetDependency(name *string, type_ DependencyType) *Dependency {
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"getDependency",
		[]interface{}{name, type_},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependencies) RemoveDependency(name *string, type_ DependencyType) {
	_jsii_.InvokeVoid(
		d,
		"removeDependency",
		[]interface{}{name, type_},
	)
}

func (d *jsiiProxy_Dependencies) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependencies) TryGetDependency(name *string, type_ DependencyType) *Dependency {
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"tryGetDependency",
		[]interface{}{name, type_},
		&returns,
	)

	return returns
}

