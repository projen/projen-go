package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Checks if an existing dependency satisfies a dependency requirement.
	//
	// Returns: `true` if the dependency exists and its version satisfies the provided constraint. `false` otherwise.
	// Notably returns `false` if a dependency exists, but has no version.
	// Experimental.
	IsDependencySatisfied(name *string, type_ DependencyType, expectedRange *string) *bool
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
	// Request a dependency. Unlike `addDependency`, this merges intelligently with existing dependencies of the same name and type:.
	//
	// - If the dep exists with a version that already satisfies the request,
	//   the version is not changed.
	// - If the dep doesn't exist, it is added with the requested type/version.
	// - If the dep exists but the versions don't intersect, an error is thrown.
	// - If no type is provided, an existing dependency of any type will satisfy
	//   the request. If none exists, it is added as BUILD.
	//
	// Returns: The resulting dependency after merging.
	// Experimental.
	RequestDependency(request *DependencyRequest) *Dependency
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Returns a dependency by name.
	//
	// Returns `undefined` if there is no dependency defined by that name or if
	// `type` is not provided and there is more then one dependency type for this
	// dependency.
	//
	// Returns: a copy (cannot be modified) or undefined if there is no match.
	// Experimental.
	TryGetDependency(name *string, type_ DependencyType) *Dependency
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
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

func (j *jsiiProxy_Dependencies) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

	if err := validateNewDependenciesParameters(project); err != nil {
		panic(err)
	}
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

// Test whether the given construct is a component.
// Experimental.
func Dependencies_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDependencies_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Dependencies",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Dependencies_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDependencies_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Dependencies",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the coordinates of a dependency spec.
//
// Given `foo@^3.4.0` returns `{ name: "foo", version: "^3.4.0" }`.
// Given `bar@npm:@bar/legacy` returns `{ name: "bar", version: "npm:@bar/legacy" }`.
// Experimental.
func Dependencies_ParseDependency(spec *string) *DependencyCoordinates {
	_init_.Initialize()

	if err := validateDependencies_ParseDependencyParameters(spec); err != nil {
		panic(err)
	}
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
	if err := d.validateAddDependencyParameters(spec, type_); err != nil {
		panic(err)
	}
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
	if err := d.validateGetDependencyParameters(name); err != nil {
		panic(err)
	}
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"getDependency",
		[]interface{}{name, type_},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) IsDependencySatisfied(name *string, type_ DependencyType, expectedRange *string) *bool {
	if err := d.validateIsDependencySatisfiedParameters(name, type_, expectedRange); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		d,
		"isDependencySatisfied",
		[]interface{}{name, type_, expectedRange},
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
	if err := d.validateRemoveDependencyParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"removeDependency",
		[]interface{}{name, type_},
	)
}

func (d *jsiiProxy_Dependencies) RequestDependency(request *DependencyRequest) *Dependency {
	if err := d.validateRequestDependencyParameters(request); err != nil {
		panic(err)
	}
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"requestDependency",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependencies) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) TryGetDependency(name *string, type_ DependencyType) *Dependency {
	if err := d.validateTryGetDependencyParameters(name); err != nil {
		panic(err)
	}
	var returns *Dependency

	_jsii_.Invoke(
		d,
		"tryGetDependency",
		[]interface{}{name, type_},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependencies) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

