package deps

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen"
)

// Class interface
type DependenciesIface interface {
	GetProject() projen.ProjectIface
	GetAll() []DependencyIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddDependency(spec string, type_ DependencyType, metadata map[string]interface{}) DependencyIface
	GetDependency(name string, type_ DependencyType) DependencyIface
	RemoveDependency(name string, type_ DependencyType)
}

// The `Dependencies` component is responsible to track the list of dependencies a project has, and then used by project types as the model for rendering project-specific dependency manifests such as the dependencies section `package.json` files.
// 
// To add a dependency you can use a project-type specific API such as
// `nodeProject.addDeps()` or use the generic API of `project.deps`:
// Experimental.
// Struct proxy
type Dependencies struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// A copy of all dependencies recorded for this project.
	// 
	// The list is sorted by type->name->version
	// Experimental.
	All []DependencyIface `json:"all"`
}

func (d *Dependencies) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		d,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Dependencies) GetAll() []DependencyIface {
	var returns []DependencyIface
	_jsii_.Get(
		d,
		"all",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyIface)(nil)).Elem(): reflect.TypeOf((*Dependency)(nil)).Elem(),
		},
	)
	return returns
}


// Adds a dependencies component to the project.
func NewDependencies(project projen.ProjectIface) DependenciesIface {
	_init_.Initialize()
	self := Dependencies{}
	_jsii_.Create(
		"projen.deps.Dependencies",
		[]interface{}{project},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Dependencies_ParseDependency(spec string) DependencyCoordinatesIface {
	_init_.Initialize()
	var returns DependencyCoordinatesIface
	_jsii_.InvokeStatic(
		"projen.deps.Dependencies",
		"parseDependency",
		[]interface{}{spec},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyCoordinatesIface)(nil)).Elem(): reflect.TypeOf((*DependencyCoordinates)(nil)).Elem(),
		},
	)
	return returns
}

func Dependencies_ManifestFile() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.deps.Dependencies",
		"MANIFEST_FILE",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Dependencies) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Dependencies) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Dependencies) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Dependencies) AddDependency(spec string, type_ DependencyType, metadata map[string]interface{}) DependencyIface {
	var returns DependencyIface
	_jsii_.Invoke(
		d,
		"addDependency",
		[]interface{}{spec, type_, metadata},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyIface)(nil)).Elem(): reflect.TypeOf((*Dependency)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Dependencies) GetDependency(name string, type_ DependencyType) DependencyIface {
	var returns DependencyIface
	_jsii_.Invoke(
		d,
		"getDependency",
		[]interface{}{name, type_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyIface)(nil)).Elem(): reflect.TypeOf((*Dependency)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Dependencies) RemoveDependency(name string, type_ DependencyType) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"removeDependency",
		[]interface{}{name, type_},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// DependencyIface is the public interface for the custom type Dependency
// Experimental.
type DependencyIface interface {
	GetName() string
	GetVersion() string
	GetType() DependencyType
	GetMetadata() map[string]interface{}
}

// Represents a project dependency.
// Experimental.
// Struct proxy
type Dependency struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	// 
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name string `json:"name"`
	// Semantic version version requirement.
	// Experimental.
	Version string `json:"version"`
	// Which type of dependency this is (runtime, build-time, etc).
	// Experimental.
	Type DependencyType `json:"type"`
	// Additional JSON metadata associated with the dependency (package manager specific).
	// Experimental.
	Metadata map[string]interface{} `json:"metadata"`
}

func (d *Dependency) GetName() string {
	var returns string
	_jsii_.Get(
		d,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Dependency) GetVersion() string {
	var returns string
	_jsii_.Get(
		d,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Dependency) GetType() DependencyType {
	var returns DependencyType
	_jsii_.Get(
		d,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyType)(nil)).Elem(): reflect.TypeOf((*DependencyType)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Dependency) GetMetadata() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		d,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}


// DependencyCoordinatesIface is the public interface for the custom type DependencyCoordinates
// Experimental.
type DependencyCoordinatesIface interface {
	GetName() string
	GetVersion() string
}

// Coordinates of the dependency (name and version).
// Experimental.
// Struct proxy
type DependencyCoordinates struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	// 
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name string `json:"name"`
	// Semantic version version requirement.
	// Experimental.
	Version string `json:"version"`
}

func (d *DependencyCoordinates) GetName() string {
	var returns string
	_jsii_.Get(
		d,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DependencyCoordinates) GetVersion() string {
	var returns string
	_jsii_.Get(
		d,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Type of dependency.
// Experimental.
type DependencyType string

const (
	DependencyTypeRuntime DependencyType = "RUNTIME"
	DependencyTypePeer DependencyType = "PEER"
	DependencyTypeBundled DependencyType = "BUNDLED"
	DependencyTypeBuild DependencyType = "BUILD"
	DependencyTypeTest DependencyType = "TEST"
	DependencyTypeDevenv DependencyType = "DEVENV"
)

// DepsManifestIface is the public interface for the custom type DepsManifest
// Experimental.
type DepsManifestIface interface {
	GetDependencies() []DependencyIface
}

// Experimental.
// Struct proxy
type DepsManifest struct {
	// All dependencies of this module.
	// Experimental.
	Dependencies []DependencyIface `json:"dependencies"`
}

func (d *DepsManifest) GetDependencies() []DependencyIface {
	var returns []DependencyIface
	_jsii_.Get(
		d,
		"dependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyIface)(nil)).Elem(): reflect.TypeOf((*Dependency)(nil)).Elem(),
		},
	)
	return returns
}


