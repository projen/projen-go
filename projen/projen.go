// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Represents a project component.
// Experimental.
type Component interface {
	Project() Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Component
type jsiiProxy_Component struct {
	_ byte // padding
}

func (j *jsiiProxy_Component) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponent(project Project) Component {
	_init_.Initialize()

	j := jsiiProxy_Component{}

	_jsii_.Create(
		"projen.Component",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewComponent_Override(c Component, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Component",
		[]interface{}{project},
		c,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (c *jsiiProxy_Component) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (c *jsiiProxy_Component) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (c *jsiiProxy_Component) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type CreateProjectOptions struct {
	// Directory that the project will be generated in.
	// Experimental.
	Dir *string `json:"dir" yaml:"dir"`
	// Fully-qualified name of the project type (usually formatted as `module.ProjectType`).
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ProjectFqn *string `json:"projectFqn" yaml:"projectFqn"`
	// Project options.
	//
	// Only JSON-like values can be passed in (strings,
	// booleans, numbers, enums, arrays, and objects that are not
	// derived from classes).
	//
	// Consult the API reference of the project type you are generating for
	// information about what fields and types are available.
	// Experimental.
	ProjectOptions *map[string]interface{} `json:"projectOptions" yaml:"projectOptions"`
	// Should we render commented-out default options in the projenrc file?
	//
	// Does not apply to projenrc.json files.
	// Experimental.
	OptionHints InitProjectOptionHints `json:"optionHints" yaml:"optionHints"`
	// Should we execute post synthesis hooks?
	//
	// (usually package manager install).
	// Experimental.
	Post *bool `json:"post" yaml:"post"`
	// Should we call `project.synth()` or instantiate the project (could still have side-effects) and render the .projenrc file.
	// Experimental.
	Synth *bool `json:"synth" yaml:"synth"`
}

// The `Dependencies` component is responsible to track the list of dependencies a project has, and then used by project types as the model for rendering project-specific dependency manifests such as the dependencies section `package.json` files.
//
// To add a dependency you can use a project-type specific API such as
// `nodeProject.addDeps()` or use the generic API of `project.deps`:
// Experimental.
type Dependencies interface {
	Component
	All() *[]*Dependency
	Project() Project
	AddDependency(spec *string, type_ DependencyType, metadata *map[string]interface{}) *Dependency
	GetDependency(name *string, type_ DependencyType) *Dependency
	PostSynthesize()
	PreSynthesize()
	RemoveDependency(name *string, type_ DependencyType)
	Synthesize()
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

// Adds a dependency to this project.
// Experimental.
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

// Returns a dependency by name.
//
// Fails if there is no dependency defined by that name or if `type` is not
// provided and there is more then one dependency type for this dependency.
//
// Returns: a copy (cannot be modified)
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (d *jsiiProxy_Dependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (d *jsiiProxy_Dependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a dependency.
// Experimental.
func (d *jsiiProxy_Dependencies) RemoveDependency(name *string, type_ DependencyType) {
	_jsii_.InvokeVoid(
		d,
		"removeDependency",
		[]interface{}{name, type_},
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (d *jsiiProxy_Dependencies) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

// Returns a dependency by name.
//
// Returns `undefined` if there is no dependency defined by that name or if
// `type` is not provided and there is more then one dependency type for this
// dependency.
//
// Returns: a copy (cannot be modified) or undefined if there is no match
// Experimental.
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

// Represents a project dependency.
// Experimental.
type Dependency struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	//
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Semantic version version requirement.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
	// Which type of dependency this is (runtime, build-time, etc).
	// Experimental.
	Type DependencyType `json:"type" yaml:"type"`
	// Additional JSON metadata associated with the dependency (package manager specific).
	// Experimental.
	Metadata *map[string]interface{} `json:"metadata" yaml:"metadata"`
}

// Coordinates of the dependency (name and version).
// Experimental.
type DependencyCoordinates struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	//
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Semantic version version requirement.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Type of dependency.
// Experimental.
type DependencyType string

const (
	DependencyType_RUNTIME DependencyType = "RUNTIME"
	DependencyType_PEER DependencyType = "PEER"
	DependencyType_BUNDLED DependencyType = "BUNDLED"
	DependencyType_BUILD DependencyType = "BUILD"
	DependencyType_TEST DependencyType = "TEST"
	DependencyType_DEVENV DependencyType = "DEVENV"
)

// Experimental.
type DepsManifest struct {
	// All dependencies of this module.
	// Experimental.
	Dependencies *[]*Dependency `json:"dependencies" yaml:"dependencies"`
}

// Options for specifying the Docker image of the container.
// Experimental.
type DevEnvironmentDockerImage interface {
	DockerFile() *string
	Image() *string
}

// The jsii proxy struct for DevEnvironmentDockerImage
type jsiiProxy_DevEnvironmentDockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_DevEnvironmentDockerImage) DockerFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevEnvironmentDockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// The relative path of a Dockerfile that defines the container contents.
//
// TODO: EXAMPLE
//
// Experimental.
func DevEnvironmentDockerImage_FromFile(dockerFile *string) DevEnvironmentDockerImage {
	_init_.Initialize()

	var returns DevEnvironmentDockerImage

	_jsii_.StaticInvoke(
		"projen.DevEnvironmentDockerImage",
		"fromFile",
		[]interface{}{dockerFile},
		&returns,
	)

	return returns
}

// A publicly available Docker image.
//
// TODO: EXAMPLE
//
// Experimental.
func DevEnvironmentDockerImage_FromImage(image *string) DevEnvironmentDockerImage {
	_init_.Initialize()

	var returns DevEnvironmentDockerImage

	_jsii_.StaticInvoke(
		"projen.DevEnvironmentDockerImage",
		"fromImage",
		[]interface{}{image},
		&returns,
	)

	return returns
}

// Base options for configuring a container-based development environment.
// Experimental.
type DevEnvironmentOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImage `json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]Task `json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `json:"vscodeExtensions" yaml:"vscodeExtensions"`
}

// Create a docker-compose YAML file.
// Experimental.
type DockerCompose interface {
	Component
	Project() Project
	AddService(serviceName *string, description *DockerComposeServiceDescription) DockerComposeService
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for DockerCompose
type jsiiProxy_DockerCompose struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_DockerCompose) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerCompose(project Project, props *DockerComposeProps) DockerCompose {
	_init_.Initialize()

	j := jsiiProxy_DockerCompose{}

	_jsii_.Create(
		"projen.DockerCompose",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerCompose_Override(d DockerCompose, project Project, props *DockerComposeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.DockerCompose",
		[]interface{}{project, props},
		d,
	)
}

// Create a bind volume that binds a host path to the target path in the container.
// Experimental.
func DockerCompose_BindVolume(sourcePath *string, targetPath *string) IDockerComposeVolumeBinding {
	_init_.Initialize()

	var returns IDockerComposeVolumeBinding

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"bindVolume",
		[]interface{}{sourcePath, targetPath},
		&returns,
	)

	return returns
}

// Create a named volume and mount it to the target path.
//
// If you use this
// named volume in several services, the volume will be shared. In this
// case, the volume configuration of the first-provided options are used.
// Experimental.
func DockerCompose_NamedVolume(volumeName *string, targetPath *string, options *DockerComposeVolumeConfig) IDockerComposeVolumeBinding {
	_init_.Initialize()

	var returns IDockerComposeVolumeBinding

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"namedVolume",
		[]interface{}{volumeName, targetPath, options},
		&returns,
	)

	return returns
}

// Create a port mapping.
// Experimental.
func DockerCompose_PortMapping(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) *DockerComposeServicePort {
	_init_.Initialize()

	var returns *DockerComposeServicePort

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"portMapping",
		[]interface{}{publishedPort, targetPort, options},
		&returns,
	)

	return returns
}

// Depends on a service name.
// Experimental.
func DockerCompose_ServiceName(serviceName *string) IDockerComposeServiceName {
	_init_.Initialize()

	var returns IDockerComposeServiceName

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"serviceName",
		[]interface{}{serviceName},
		&returns,
	)

	return returns
}

// Add a service to the docker-compose file.
// Experimental.
func (d *jsiiProxy_DockerCompose) AddService(serviceName *string, description *DockerComposeServiceDescription) DockerComposeService {
	var returns DockerComposeService

	_jsii_.Invoke(
		d,
		"addService",
		[]interface{}{serviceName, description},
		&returns,
	)

	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (d *jsiiProxy_DockerCompose) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (d *jsiiProxy_DockerCompose) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (d *jsiiProxy_DockerCompose) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

// Build arguments for creating a docker image.
// Experimental.
type DockerComposeBuild struct {
	// Docker build context directory.
	// Experimental.
	Context *string `json:"context" yaml:"context"`
	// Build args.
	// Experimental.
	Args *map[string]*string `json:"args" yaml:"args"`
	// A dockerfile to build from.
	// Experimental.
	Dockerfile *string `json:"dockerfile" yaml:"dockerfile"`
}

// Options for port mappings.
// Experimental.
type DockerComposePortMappingOptions struct {
	// Port mapping protocol.
	// Experimental.
	Protocol DockerComposeProtocol `json:"protocol" yaml:"protocol"`
}

// Props for DockerCompose.
// Experimental.
type DockerComposeProps struct {
	// A name to add to the docker-compose.yml filename.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	NameSuffix *string `json:"nameSuffix" yaml:"nameSuffix"`
	// Docker Compose schema version do be used.
	// Experimental.
	SchemaVersion *string `json:"schemaVersion" yaml:"schemaVersion"`
	// Service descriptions.
	// Experimental.
	Services *map[string]*DockerComposeServiceDescription `json:"services" yaml:"services"`
}

// Network protocol for port mapping.
// Experimental.
type DockerComposeProtocol string

const (
	DockerComposeProtocol_TCP DockerComposeProtocol = "TCP"
	DockerComposeProtocol_UDP DockerComposeProtocol = "UDP"
)

// A docker-compose service.
// Experimental.
type DockerComposeService interface {
	IDockerComposeServiceName
	Command() *[]*string
	DependsOn() *[]IDockerComposeServiceName
	Environment() *map[string]*string
	Image() *string
	ImageBuild() *DockerComposeBuild
	Ports() *[]*DockerComposeServicePort
	ServiceName() *string
	Volumes() *[]IDockerComposeVolumeBinding
	AddDependsOn(serviceName IDockerComposeServiceName)
	AddEnvironment(name *string, value *string)
	AddPort(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions)
	AddVolume(volume IDockerComposeVolumeBinding)
}

// The jsii proxy struct for DockerComposeService
type jsiiProxy_DockerComposeService struct {
	jsiiProxy_IDockerComposeServiceName
}

func (j *jsiiProxy_DockerComposeService) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) DependsOn() *[]IDockerComposeServiceName {
	var returns *[]IDockerComposeServiceName
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) ImageBuild() *DockerComposeBuild {
	var returns *DockerComposeBuild
	_jsii_.Get(
		j,
		"imageBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Ports() *[]*DockerComposeServicePort {
	var returns *[]*DockerComposeServicePort
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Volumes() *[]IDockerComposeVolumeBinding {
	var returns *[]IDockerComposeVolumeBinding
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerComposeService(serviceName *string, serviceDescription *DockerComposeServiceDescription) DockerComposeService {
	_init_.Initialize()

	j := jsiiProxy_DockerComposeService{}

	_jsii_.Create(
		"projen.DockerComposeService",
		[]interface{}{serviceName, serviceDescription},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerComposeService_Override(d DockerComposeService, serviceName *string, serviceDescription *DockerComposeServiceDescription) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.DockerComposeService",
		[]interface{}{serviceName, serviceDescription},
		d,
	)
}

// Make the service depend on another service.
// Experimental.
func (d *jsiiProxy_DockerComposeService) AddDependsOn(serviceName IDockerComposeServiceName) {
	_jsii_.InvokeVoid(
		d,
		"addDependsOn",
		[]interface{}{serviceName},
	)
}

// Add an environment variable.
// Experimental.
func (d *jsiiProxy_DockerComposeService) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		d,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

// Add a port mapping.
// Experimental.
func (d *jsiiProxy_DockerComposeService) AddPort(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) {
	_jsii_.InvokeVoid(
		d,
		"addPort",
		[]interface{}{publishedPort, targetPort, options},
	)
}

// Add a volume to the service.
// Experimental.
func (d *jsiiProxy_DockerComposeService) AddVolume(volume IDockerComposeVolumeBinding) {
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{volume},
	)
}

// Description of a docker-compose.yml service.
// Experimental.
type DockerComposeServiceDescription struct {
	// Provide a command to the docker container.
	// Experimental.
	Command *[]*string `json:"command" yaml:"command"`
	// Names of other services this service depends on.
	// Experimental.
	DependsOn *[]IDockerComposeServiceName `json:"dependsOn" yaml:"dependsOn"`
	// Add environment variables.
	// Experimental.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// Use a docker image.
	//
	// Note: You must specify either `build` or `image` key.
	// See: imageBuild
	//
	// Experimental.
	Image *string `json:"image" yaml:"image"`
	// Build a docker image.
	//
	// Note: You must specify either `imageBuild` or `image` key.
	// See: image
	//
	// Experimental.
	ImageBuild *DockerComposeBuild `json:"imageBuild" yaml:"imageBuild"`
	// Map some ports.
	// Experimental.
	Ports *[]*DockerComposeServicePort `json:"ports" yaml:"ports"`
	// Mount some volumes into the service.
	//
	// Use one of the following to create volumes:
	// See: DockerCompose.namedVolume() to create & mount a named volume
	//
	// Experimental.
	Volumes *[]IDockerComposeVolumeBinding `json:"volumes" yaml:"volumes"`
}

// A service port mapping.
// Experimental.
type DockerComposeServicePort struct {
	// Port mapping mode.
	// Experimental.
	Mode *string `json:"mode" yaml:"mode"`
	// Network protocol.
	// Experimental.
	Protocol DockerComposeProtocol `json:"protocol" yaml:"protocol"`
	// Published port number.
	// Experimental.
	Published *float64 `json:"published" yaml:"published"`
	// Target port number.
	// Experimental.
	Target *float64 `json:"target" yaml:"target"`
}

// Volume configuration.
// Experimental.
type DockerComposeVolumeConfig struct {
	// Driver to use for the volume.
	// Experimental.
	Driver *string `json:"driver" yaml:"driver"`
	// Options to provide to the driver.
	// Experimental.
	DriverOpts *map[string]*string `json:"driverOpts" yaml:"driverOpts"`
	// Set to true to indicate that the volume is externally created.
	// Experimental.
	External *bool `json:"external" yaml:"external"`
	// Name of the volume for when the volume name isn't going to work in YAML.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Service volume mounting information.
// Experimental.
type DockerComposeVolumeMount struct {
	// Volume source.
	// Experimental.
	Source *string `json:"source" yaml:"source"`
	// Volume target.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
	// Type of volume.
	// Experimental.
	Type *string `json:"type" yaml:"type"`
}

// Experimental.
type FileBase interface {
	Component
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for FileBase
type jsiiProxy_FileBase struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_FileBase) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewFileBase_Override(f FileBase, project Project, filePath *string, options *FileBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.FileBase",
		[]interface{}{project, filePath, options},
		f,
	)
}

func (j *jsiiProxy_FileBase) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_FileBase) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func FileBase_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.FileBase",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (f *jsiiProxy_FileBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (f *jsiiProxy_FileBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (f *jsiiProxy_FileBase) Synthesize() {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
//
// Returns: the content to synthesize or undefined to skip the file
// Experimental.
func (f *jsiiProxy_FileBase) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Experimental.
type FileBaseOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
}

// Assign attributes to file names in a git repository.
// See: https://git-scm.com/docs/gitattributes
//
// Experimental.
type GitAttributesFile interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddAttributes(glob *string, attributes ...*string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_arg IResolver) *string
}

// The jsii proxy struct for GitAttributesFile
type jsiiProxy_GitAttributesFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_GitAttributesFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitAttributesFile(project Project) GitAttributesFile {
	_init_.Initialize()

	j := jsiiProxy_GitAttributesFile{}

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewGitAttributesFile_Override(g GitAttributesFile, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{project},
		g,
	)
}

func (j *jsiiProxy_GitAttributesFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_GitAttributesFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func GitAttributesFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.GitAttributesFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Maps a set of attributes to a set of files.
// Experimental.
func (g *jsiiProxy_GitAttributesFile) AddAttributes(glob *string, attributes ...*string) {
	args := []interface{}{glob}
	for _, a := range attributes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addAttributes",
		args,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (g *jsiiProxy_GitAttributesFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (g *jsiiProxy_GitAttributesFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (g *jsiiProxy_GitAttributesFile) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (g *jsiiProxy_GitAttributesFile) SynthesizeContent(_arg IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// The Gitpod component which emits .gitpod.yml.
// Experimental.
type Gitpod interface {
	Component
	IDevEnvironment
	Config() interface{}
	Project() Project
	AddCustomTask(options *GitpodTask)
	AddDockerImage(image DevEnvironmentDockerImage)
	AddPorts(ports ...*string)
	AddPrebuilds(config *GitpodPrebuilds)
	AddTasks(tasks ...Task)
	AddVscodeExtensions(extensions ...*string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Gitpod
type jsiiProxy_Gitpod struct {
	jsiiProxy_Component
	jsiiProxy_IDevEnvironment
}

func (j *jsiiProxy_Gitpod) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gitpod) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitpod(project Project, options *GitpodOptions) Gitpod {
	_init_.Initialize()

	j := jsiiProxy_Gitpod{}

	_jsii_.Create(
		"projen.Gitpod",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitpod_Override(g Gitpod, project Project, options *GitpodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Gitpod",
		[]interface{}{project, options},
		g,
	)
}

// Add a task with more granular options.
//
// By default, all tasks will be run in parallel. To run tasks in sequence,
// create a new `Task` and set the other tasks as subtasks.
// Experimental.
func (g *jsiiProxy_Gitpod) AddCustomTask(options *GitpodTask) {
	_jsii_.InvokeVoid(
		g,
		"addCustomTask",
		[]interface{}{options},
	)
}

// Add a custom Docker image or Dockerfile for the container.
// Experimental.
func (g *jsiiProxy_Gitpod) AddDockerImage(image DevEnvironmentDockerImage) {
	_jsii_.InvokeVoid(
		g,
		"addDockerImage",
		[]interface{}{image},
	)
}

// Add ports that should be exposed (forwarded) from the container.
// Experimental.
func (g *jsiiProxy_Gitpod) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addPorts",
		args,
	)
}

// Add a prebuilds configuration for the Gitpod App.
// Experimental.
func (g *jsiiProxy_Gitpod) AddPrebuilds(config *GitpodPrebuilds) {
	_jsii_.InvokeVoid(
		g,
		"addPrebuilds",
		[]interface{}{config},
	)
}

// Add tasks to run when gitpod starts.
//
// By default, all tasks will be run in parallel. To run tasks in sequence,
// create a new `Task` and specify the other tasks as subtasks.
// Experimental.
func (g *jsiiProxy_Gitpod) AddTasks(tasks ...Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addTasks",
		args,
	)
}

// Add a list of VSCode extensions that should be automatically installed in the container.
//
// These must be in the format defined in the Open VSX registry.
//
// TODO: EXAMPLE
//
// See: https://www.gitpod.io/docs/vscode-extensions/
//
// Experimental.
func (g *jsiiProxy_Gitpod) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addVscodeExtensions",
		args,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (g *jsiiProxy_Gitpod) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (g *jsiiProxy_Gitpod) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (g *jsiiProxy_Gitpod) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

// What to do when a service on a port is detected.
// Experimental.
type GitpodOnOpen string

const (
	GitpodOnOpen_OPEN_BROWSER GitpodOnOpen = "OPEN_BROWSER"
	GitpodOnOpen_OPEN_PREVIEW GitpodOnOpen = "OPEN_PREVIEW"
	GitpodOnOpen_NOTIFY GitpodOnOpen = "NOTIFY"
	GitpodOnOpen_IGNORE GitpodOnOpen = "IGNORE"
)

// Configure where in the IDE the terminal should be opened.
// Experimental.
type GitpodOpenIn string

const (
	GitpodOpenIn_BOTTOM GitpodOpenIn = "BOTTOM"
	GitpodOpenIn_LEFT GitpodOpenIn = "LEFT"
	GitpodOpenIn_RIGHT GitpodOpenIn = "RIGHT"
	GitpodOpenIn_MAIN GitpodOpenIn = "MAIN"
)

// Configure how the terminal should be opened relative to the previous task.
// Experimental.
type GitpodOpenMode string

const (
	GitpodOpenMode_TAB_AFTER GitpodOpenMode = "TAB_AFTER"
	GitpodOpenMode_TAB_BEFORE GitpodOpenMode = "TAB_BEFORE"
	GitpodOpenMode_SPLIT_RIGHT GitpodOpenMode = "SPLIT_RIGHT"
	GitpodOpenMode_SPLIT_LEFT GitpodOpenMode = "SPLIT_LEFT"
	GitpodOpenMode_SPLIT_TOP GitpodOpenMode = "SPLIT_TOP"
	GitpodOpenMode_SPLIT_BOTTOM GitpodOpenMode = "SPLIT_BOTTOM"
)

// Constructor options for the Gitpod component.
//
// By default, Gitpod uses the 'gitpod/workspace-full' docker image.
// See: https://github.com/gitpod-io/workspace-images/blob/master/full/Dockerfile
//
// By default, all tasks will be run in parallel. To run the tasks in sequence,
// create a new task and specify the other tasks as subtasks.
//
// Experimental.
type GitpodOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImage `json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]Task `json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `json:"vscodeExtensions" yaml:"vscodeExtensions"`
	// Optional Gitpod's Github App integration for prebuilds If this is not set and Gitpod's Github App is installed, then Gitpod will apply these defaults: https://www.gitpod.io/docs/prebuilds/#configure-the-github-app.
	// Experimental.
	Prebuilds *GitpodPrebuilds `json:"prebuilds" yaml:"prebuilds"`
}

// Options for an exposed port on Gitpod.
// Experimental.
type GitpodPort struct {
	// What to do when a service on a port is detected.
	// Experimental.
	OnOpen GitpodOnOpen `json:"onOpen" yaml:"onOpen"`
	// A port that should be exposed (forwarded) from the container.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Port *string `json:"port" yaml:"port"`
	// Whether the port visibility should be private or public.
	// Experimental.
	Visibility GitpodPortVisibility `json:"visibility" yaml:"visibility"`
}

// Whether the port visibility should be private or public.
// Experimental.
type GitpodPortVisibility string

const (
	GitpodPortVisibility_PUBLIC GitpodPortVisibility = "PUBLIC"
	GitpodPortVisibility_PRIVATE GitpodPortVisibility = "PRIVATE"
)

// Configure the Gitpod App for prebuilds.
//
// Currently only GitHub is supported.
// See: https://www.gitpod.io/docs/prebuilds/
//
// Experimental.
type GitpodPrebuilds struct {
	// Add a "Review in Gitpod" button to the pull request's description.
	// Experimental.
	AddBadge *bool `json:"addBadge" yaml:"addBadge"`
	// Add a check to pull requests.
	// Experimental.
	AddCheck *bool `json:"addCheck" yaml:"addCheck"`
	// Add a "Review in Gitpod" button as a comment to pull requests.
	// Experimental.
	AddComment *bool `json:"addComment" yaml:"addComment"`
	// Add a label once the prebuild is ready to pull requests.
	// Experimental.
	AddLabel *bool `json:"addLabel" yaml:"addLabel"`
	// Enable for all branches in this repo.
	// Experimental.
	Branches *bool `json:"branches" yaml:"branches"`
	// Enable for the master/default branch.
	// Experimental.
	Master *bool `json:"master" yaml:"master"`
	// Enable for pull requests coming from this repo.
	// Experimental.
	PullRequests *bool `json:"pullRequests" yaml:"pullRequests"`
	// Enable for pull requests coming from forks.
	// Experimental.
	PullRequestsFromForks *bool `json:"pullRequestsFromForks" yaml:"pullRequestsFromForks"`
}

// Configure options for a task to be run when opening a Gitpod workspace (e.g. running tests, or starting a dev server).
//
// Start Mode         | Execution
// Fresh Workspace    | before && init && command
// Restart Workspace  | before && command
// Snapshot           | before && command
// Prebuild           | before && init && prebuild
// Experimental.
type GitpodTask struct {
	// Required.
	//
	// The shell command to run
	// Experimental.
	Command *string `json:"command" yaml:"command"`
	// In case you need to run something even before init, that is a requirement for both init and command, you can use the before property.
	// Experimental.
	Before *string `json:"before" yaml:"before"`
	// The init property can be used to specify shell commands that should only be executed after a workspace was freshly cloned and needs to be initialized somehow.
	//
	// Such tasks are usually builds or downloading
	// dependencies. Anything you only want to do once but not when you restart a workspace or start a snapshot.
	// Experimental.
	Init *string `json:"init" yaml:"init"`
	// A name for this task.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// You can configure where in the IDE the terminal should be opened.
	// Experimental.
	OpenIn GitpodOpenIn `json:"openIn" yaml:"openIn"`
	// You can configure how the terminal should be opened relative to the previous task.
	// Experimental.
	OpenMode GitpodOpenMode `json:"openMode" yaml:"openMode"`
	// The optional prebuild command will be executed during prebuilds.
	//
	// It is meant to run additional long running
	// processes that could be useful, e.g. running test suites.
	// Experimental.
	Prebuild *string `json:"prebuild" yaml:"prebuild"`
}

// Abstract interface for container-based development environments, such as Gitpod and GitHub Codespaces.
// Experimental.
type IDevEnvironment interface {
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image DevEnvironmentDockerImage)
	// Adds ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports ...*string)
	// Adds tasks to run when the container starts.
	// Experimental.
	AddTasks(tasks ...Task)
	// Adds a list of VSCode extensions that should be automatically installed in the container.
	// Experimental.
	AddVscodeExtensions(extensions ...*string)
}

// The jsii proxy for IDevEnvironment
type jsiiProxy_IDevEnvironment struct {
	_ byte // padding
}

func (i *jsiiProxy_IDevEnvironment) AddDockerImage(image DevEnvironmentDockerImage) {
	_jsii_.InvokeVoid(
		i,
		"addDockerImage",
		[]interface{}{image},
	)
}

func (i *jsiiProxy_IDevEnvironment) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addPorts",
		args,
	)
}

func (i *jsiiProxy_IDevEnvironment) AddTasks(tasks ...Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTasks",
		args,
	)
}

func (i *jsiiProxy_IDevEnvironment) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addVscodeExtensions",
		args,
	)
}

// An interface providing the name of a docker compose service.
// Experimental.
type IDockerComposeServiceName interface {
	// The name of the docker compose service.
	// Experimental.
	ServiceName() *string
}

// The jsii proxy for IDockerComposeServiceName
type jsiiProxy_IDockerComposeServiceName struct {
	_ byte // padding
}

func (j *jsiiProxy_IDockerComposeServiceName) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

// Volume binding information.
// Experimental.
type IDockerComposeVolumeBinding interface {
	// Binds the requested volume to the docker-compose volume configuration and provide mounting instructions for synthesis.
	//
	// Returns: mounting instructions for the service.
	// Experimental.
	Bind(volumeConfig IDockerComposeVolumeConfig) *DockerComposeVolumeMount
}

// The jsii proxy for IDockerComposeVolumeBinding
type jsiiProxy_IDockerComposeVolumeBinding struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeVolumeBinding) Bind(volumeConfig IDockerComposeVolumeConfig) *DockerComposeVolumeMount {
	var returns *DockerComposeVolumeMount

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{volumeConfig},
		&returns,
	)

	return returns
}

// Storage for volume configuration.
// Experimental.
type IDockerComposeVolumeConfig interface {
	// Add volume configuration to the repository.
	// Experimental.
	AddVolumeConfiguration(volumeName *string, configuration *DockerComposeVolumeConfig)
}

// The jsii proxy for IDockerComposeVolumeConfig
type jsiiProxy_IDockerComposeVolumeConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeVolumeConfig) AddVolumeConfiguration(volumeName *string, configuration *DockerComposeVolumeConfig) {
	_jsii_.InvokeVoid(
		i,
		"addVolumeConfiguration",
		[]interface{}{volumeName, configuration},
	)
}

// Experimental.
type IResolvable interface {
	// Resolves and returns content.
	// Experimental.
	ToJSON() interface{}
}

// The jsii proxy for IResolvable
type jsiiProxy_IResolvable struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolvable) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

// API for resolving tokens when synthesizing file content.
// Experimental.
type IResolver interface {
	// Given a value (object/string/array/whatever, looks up any functions inside the object and returns an object where all functions are called.
	// Experimental.
	Resolve(value interface{}, options *ResolveOptions) interface{}
}

// The jsii proxy for IResolver
type jsiiProxy_IResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolver) Resolve(value interface{}, options *ResolveOptions) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Experimental.
type IgnoreFile interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddPatterns(patterns ...*string)
	Exclude(patterns ...*string)
	Include(patterns ...*string)
	PostSynthesize()
	PreSynthesize()
	RemovePatterns(patterns ...*string)
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for IgnoreFile
type jsiiProxy_IgnoreFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_IgnoreFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewIgnoreFile(project Project, filePath *string) IgnoreFile {
	_init_.Initialize()

	j := jsiiProxy_IgnoreFile{}

	_jsii_.Create(
		"projen.IgnoreFile",
		[]interface{}{project, filePath},
		&j,
	)

	return &j
}

// Experimental.
func NewIgnoreFile_Override(i IgnoreFile, project Project, filePath *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.IgnoreFile",
		[]interface{}{project, filePath},
		i,
	)
}

func (j *jsiiProxy_IgnoreFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_IgnoreFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func IgnoreFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.IgnoreFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Add ignore patterns.
//
// Files that match this pattern will be ignored. If the
// pattern starts with a negation mark `!`, files that match will _not_ be
// ignored.
//
// Comment lines (start with `#`) are ignored.
// Experimental.
func (i *jsiiProxy_IgnoreFile) AddPatterns(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addPatterns",
		args,
	)
}

// Ignore the files that match these patterns.
// Experimental.
func (i *jsiiProxy_IgnoreFile) Exclude(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"exclude",
		args,
	)
}

// Always include the specified file patterns.
// Experimental.
func (i *jsiiProxy_IgnoreFile) Include(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"include",
		args,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (i *jsiiProxy_IgnoreFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (i *jsiiProxy_IgnoreFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes patterns previously added from the ignore file.
//
// If `addPattern()` is called after this, the pattern will be added again.
// Experimental.
func (i *jsiiProxy_IgnoreFile) RemovePatterns(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"removePatterns",
		args,
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (i *jsiiProxy_IgnoreFile) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (i *jsiiProxy_IgnoreFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Represents an INI file.
// Experimental.
type IniFile interface {
	ObjectFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for IniFile
type jsiiProxy_IniFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_IniFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IniFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewIniFile(project Project, filePath *string, options *IniFileOptions) IniFile {
	_init_.Initialize()

	j := jsiiProxy_IniFile{}

	_jsii_.Create(
		"projen.IniFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIniFile_Override(i IniFile, project Project, filePath *string, options *IniFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.IniFile",
		[]interface{}{project, filePath, options},
		i,
	)
}

func (j *jsiiProxy_IniFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_IniFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_IniFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func IniFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.IniFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (i *jsiiProxy_IniFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		i,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (i *jsiiProxy_IniFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (i *jsiiProxy_IniFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (i *jsiiProxy_IniFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (i *jsiiProxy_IniFile) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (i *jsiiProxy_IniFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `IniFile`.
// Experimental.
type IniFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// Information passed from `projen new` to the project object when the project is first created.
//
// It is used to generate projenrc files in various languages.
// Experimental.
type InitProject struct {
	// Initial arguments passed to `projen new`.
	// Experimental.
	Args *map[string]interface{} `json:"args" yaml:"args"`
	// Include commented out options.
	//
	// Does not apply to projenrc.json files.
	// Experimental.
	Comments InitProjectOptionHints `json:"comments" yaml:"comments"`
	// The JSII FQN of the project type.
	// Experimental.
	Fqn *string `json:"fqn" yaml:"fqn"`
	// Project metadata.
	// Experimental.
	Type ProjectType `json:"type" yaml:"type"`
}

// Choices for how to display commented out options in projenrc files.
//
// Does not apply to projenrc.json files.
// Experimental.
type InitProjectOptionHints string

const (
	InitProjectOptionHints_ALL InitProjectOptionHints = "ALL"
	InitProjectOptionHints_FEATURED InitProjectOptionHints = "FEATURED"
	InitProjectOptionHints_NONE InitProjectOptionHints = "NONE"
)

// Represents a JSON file.
// Experimental.
type JsonFile interface {
	ObjectFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for JsonFile
type jsiiProxy_JsonFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_JsonFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewJsonFile(project Project, filePath *string, options *JsonFileOptions) JsonFile {
	_init_.Initialize()

	j := jsiiProxy_JsonFile{}

	_jsii_.Create(
		"projen.JsonFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJsonFile_Override(j JsonFile, project Project, filePath *string, options *JsonFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.JsonFile",
		[]interface{}{project, filePath, options},
		j,
	)
}

func (j *jsiiProxy_JsonFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_JsonFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_JsonFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func JsonFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.JsonFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (j *jsiiProxy_JsonFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		j,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (j *jsiiProxy_JsonFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (j *jsiiProxy_JsonFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (j *jsiiProxy_JsonFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (j *jsiiProxy_JsonFile) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (j *jsiiProxy_JsonFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `JsonFile`.
// Experimental.
type JsonFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
	// Adds a newline at the end of the file.
	// Experimental.
	Newline *bool `json:"newline" yaml:"newline"`
}

// Experimental.
type License interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_arg IResolver) *string
}

// The jsii proxy struct for License
type jsiiProxy_License struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_License) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_License) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_License) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_License) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_License) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_License) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewLicense(project Project, options *LicenseOptions) License {
	_init_.Initialize()

	j := jsiiProxy_License{}

	_jsii_.Create(
		"projen.License",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLicense_Override(l License, project Project, options *LicenseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.License",
		[]interface{}{project, options},
		l,
	)
}

func (j *jsiiProxy_License) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_License) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func License_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.License",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (l *jsiiProxy_License) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (l *jsiiProxy_License) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (l *jsiiProxy_License) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (l *jsiiProxy_License) SynthesizeContent(_arg IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Experimental.
type LicenseOptions struct {
	// License type (SPDX).
	// See: https://github.com/projen/projen/tree/main/license-text for list of supported licenses
	//
	// Experimental.
	Spdx *string `json:"spdx" yaml:"spdx"`
	// Copyright owner.
	//
	// If the license text has $copyright_owner, this option must be specified.
	// Experimental.
	CopyrightOwner *string `json:"copyrightOwner" yaml:"copyrightOwner"`
	// Period of license (e.g. "1998-2023").
	//
	// The string `$copyright_period` will be substituted with this string.
	// Experimental.
	CopyrightPeriod *string `json:"copyrightPeriod" yaml:"copyrightPeriod"`
}

// Logging verbosity.
// Experimental.
type LogLevel string

const (
	LogLevel_OFF LogLevel = "OFF"
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_WARN LogLevel = "WARN"
	LogLevel_INFO LogLevel = "INFO"
	LogLevel_DEBUG LogLevel = "DEBUG"
	LogLevel_VERBOSE LogLevel = "VERBOSE"
)

// Project-level logging utilities.
// Experimental.
type Logger interface {
	Component
	Project() Project
	Debug(text ...interface{})
	Error(text ...interface{})
	Info(text ...interface{})
	Log(level LogLevel, text ...interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	Verbose(text ...interface{})
	Warn(text ...interface{})
}

// The jsii proxy struct for Logger
type jsiiProxy_Logger struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Logger) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogger(project Project, options *LoggerOptions) Logger {
	_init_.Initialize()

	j := jsiiProxy_Logger{}

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLogger_Override(l Logger, project Project, options *LoggerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{project, options},
		l,
	)
}

// Log a message to stderr with DEBUG severity.
// Experimental.
func (l *jsiiProxy_Logger) Debug(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"debug",
		args,
	)
}

// Log a message to stderr with ERROR severity.
// Experimental.
func (l *jsiiProxy_Logger) Error(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"error",
		args,
	)
}

// Log a message to stderr with INFO severity.
// Experimental.
func (l *jsiiProxy_Logger) Info(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"info",
		args,
	)
}

// Log a message to stderr with a given logging level.
//
// The message will be
// printed as long as `logger.level` is set to the message's severity or higher.
// Experimental.
func (l *jsiiProxy_Logger) Log(level LogLevel, text ...interface{}) {
	args := []interface{}{level}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"log",
		args,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (l *jsiiProxy_Logger) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (l *jsiiProxy_Logger) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (l *jsiiProxy_Logger) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Log a message to stderr with VERBOSE severity.
// Experimental.
func (l *jsiiProxy_Logger) Verbose(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"verbose",
		args,
	)
}

// Log a message to stderr with WARN severity.
// Experimental.
func (l *jsiiProxy_Logger) Warn(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"warn",
		args,
	)
}

// Options for logging utilities.
// Experimental.
type LoggerOptions struct {
	// The logging verbosity.
	//
	// The levels available (in increasing verbosity) are
	// OFF, ERROR, WARN, INFO, DEBUG, and VERBOSE.
	// Experimental.
	Level LogLevel `json:"level" yaml:"level"`
	// Include a prefix for all logging messages with the project name.
	// Experimental.
	UsePrefix *bool `json:"usePrefix" yaml:"usePrefix"`
}

// Minimal Makefile.
// Experimental.
type Makefile interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	Rules() *[]*Rule
	AddAll(target *string) Makefile
	AddAlls(targets ...*string) Makefile
	AddRule(rule *Rule) Makefile
	AddRules(rules ...*Rule) Makefile
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for Makefile
type jsiiProxy_Makefile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_Makefile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Rules() *[]*Rule {
	var returns *[]*Rule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Experimental.
func NewMakefile(project Project, filePath *string, options *MakefileOptions) Makefile {
	_init_.Initialize()

	j := jsiiProxy_Makefile{}

	_jsii_.Create(
		"projen.Makefile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMakefile_Override(m Makefile, project Project, filePath *string, options *MakefileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Makefile",
		[]interface{}{project, filePath, options},
		m,
	)
}

func (j *jsiiProxy_Makefile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_Makefile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func Makefile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Makefile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Add a target to all.
// Experimental.
func (m *jsiiProxy_Makefile) AddAll(target *string) Makefile {
	var returns Makefile

	_jsii_.Invoke(
		m,
		"addAll",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Add multiple targets to all.
// Experimental.
func (m *jsiiProxy_Makefile) AddAlls(targets ...*string) Makefile {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	var returns Makefile

	_jsii_.Invoke(
		m,
		"addAlls",
		args,
		&returns,
	)

	return returns
}

// Add a rule to the Makefile.
// Experimental.
func (m *jsiiProxy_Makefile) AddRule(rule *Rule) Makefile {
	var returns Makefile

	_jsii_.Invoke(
		m,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Add multiple rules to the Makefile.
// Experimental.
func (m *jsiiProxy_Makefile) AddRules(rules ...*Rule) Makefile {
	args := []interface{}{}
	for _, a := range rules {
		args = append(args, a)
	}

	var returns Makefile

	_jsii_.Invoke(
		m,
		"addRules",
		args,
		&returns,
	)

	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (m *jsiiProxy_Makefile) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (m *jsiiProxy_Makefile) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (m *jsiiProxy_Makefile) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (m *jsiiProxy_Makefile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for Makefiles.
// Experimental.
type MakefileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// List of targets to build when Make is invoked without specifying any targets.
	// Experimental.
	All *[]*string `json:"all" yaml:"all"`
	// Rules to include in the Makefile.
	// Experimental.
	Rules *[]*Rule `json:"rules" yaml:"rules"`
}

// Represents an Object file.
// Experimental.
type ObjectFile interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for ObjectFile
type jsiiProxy_ObjectFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_ObjectFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewObjectFile_Override(o ObjectFile, project Project, filePath *string, options *ObjectFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ObjectFile",
		[]interface{}{project, filePath, options},
		o,
	)
}

func (j *jsiiProxy_ObjectFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_ObjectFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_ObjectFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func ObjectFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.ObjectFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (o *jsiiProxy_ObjectFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		o,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (o *jsiiProxy_ObjectFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (o *jsiiProxy_ObjectFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		o,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (o *jsiiProxy_ObjectFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		o,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (o *jsiiProxy_ObjectFile) Synthesize() {
	_jsii_.InvokeVoid(
		o,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (o *jsiiProxy_ObjectFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `ObjectFile`.
// Experimental.
type ObjectFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// Base project.
// Experimental.
type Project interface {
	BuildTask() Task
	CompileTask() Task
	Components() *[]Component
	DefaultTask() Task
	Deps() Dependencies
	Files() *[]FileBase
	Gitattributes() GitAttributesFile
	Gitignore() IgnoreFile
	InitProject() *InitProject
	Logger() Logger
	Name() *string
	Outdir() *string
	PackageTask() Task
	Parent() Project
	PostCompileTask() Task
	PreCompileTask() Task
	ProjectBuild() ProjectBuild
	ProjenCommand() *string
	Root() Project
	Tasks() Tasks
	TestTask() Task
	AddExcludeFromCleanup(globs ...*string)
	AddGitIgnore(pattern *string)
	AddPackageIgnore(_pattern *string)
	AddTask(name *string, props *TaskOptions) Task
	AddTip(message *string)
	AnnotateGenerated(_glob *string)
	PostSynthesize()
	PreSynthesize()
	RemoveTask(name *string) Task
	RunTaskCommand(task Task) *string
	Synth()
	TryFindFile(filePath *string) FileBase
	TryFindJsonFile(filePath *string) JsonFile
	TryFindObjectFile(filePath *string) ObjectFile
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	_ byte // padding
}

func (j *jsiiProxy_Project) BuildTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Components() *[]Component {
	var returns *[]Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DefaultTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Deps() Dependencies {
	var returns Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Files() *[]FileBase {
	var returns *[]FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Gitattributes() GitAttributesFile {
	var returns GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Gitignore() IgnoreFile {
	var returns IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InitProject() *InitProject {
	var returns *InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Logger() Logger {
	var returns Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PackageTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Parent() Project {
	var returns Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PostCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectBuild() ProjectBuild {
	var returns ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Root() Project {
	var returns Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Tasks() Tasks {
	var returns Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TestTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewProject(options *ProjectOptions) Project {
	_init_.Initialize()

	j := jsiiProxy_Project{}

	_jsii_.Create(
		"projen.Project",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewProject_Override(p Project, options *ProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Project",
		[]interface{}{options},
		p,
	)
}

func Project_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Project",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (p *jsiiProxy_Project) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addExcludeFromCleanup",
		args,
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (p *jsiiProxy_Project) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (p *jsiiProxy_Project) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (p *jsiiProxy_Project) AddTask(name *string, props *TaskOptions) Task {
	var returns Task

	_jsii_.Invoke(
		p,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (p *jsiiProxy_Project) AddTip(message *string) {
	_jsii_.InvokeVoid(
		p,
		"addTip",
		[]interface{}{message},
	)
}

// Consider a set of files as "generated".
//
// This method is implemented by
// derived classes and used for example, to add git attributes to tell GitHub
// that certain files are generated.
// Experimental.
func (p *jsiiProxy_Project) AnnotateGenerated(_glob *string) {
	_jsii_.InvokeVoid(
		p,
		"annotateGenerated",
		[]interface{}{_glob},
	)
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Project) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (p *jsiiProxy_Project) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (p *jsiiProxy_Project) RemoveTask(name *string) Task {
	var returns Task

	_jsii_.Invoke(
		p,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the shell command to execute in order to run a task.
//
// By default, this is `npx projen@<version> <task>`
// Experimental.
func (p *jsiiProxy_Project) RunTaskCommand(task Task) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Synthesize all project files into `outdir`.
//
// 1. Call "this.preSynthesize()"
// 2. Delete all generated files
// 3. Synthesize all sub-projects
// 4. Synthesize all components of this project
// 5. Call "postSynthesize()" for all components of this project
// 6. Call "this.postSynthesize()"
// Experimental.
func (p *jsiiProxy_Project) Synth() {
	_jsii_.InvokeVoid(
		p,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (p *jsiiProxy_Project) TryFindFile(filePath *string) FileBase {
	var returns FileBase

	_jsii_.Invoke(
		p,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (p *jsiiProxy_Project) TryFindJsonFile(filePath *string) JsonFile {
	var returns JsonFile

	_jsii_.Invoke(
		p,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (p *jsiiProxy_Project) TryFindObjectFile(filePath *string) ObjectFile {
	var returns ObjectFile

	_jsii_.Invoke(
		p,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Manages a standard build process for all projects.
//
// Build spawns these tasks in order:
// 1. default
// 2. pre-compile
// 3. compile
// 4. post-compile
// 5. test
// 6. package
// Experimental.
type ProjectBuild interface {
	Component
	BuildTask() Task
	CompileTask() Task
	PackageTask() Task
	PostCompileTask() Task
	PreCompileTask() Task
	Project() Project
	TestTask() Task
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for ProjectBuild
type jsiiProxy_ProjectBuild struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ProjectBuild) BuildTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) CompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PackageTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PostCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PreCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) TestTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjectBuild(project Project) ProjectBuild {
	_init_.Initialize()

	j := jsiiProxy_ProjectBuild{}

	_jsii_.Create(
		"projen.ProjectBuild",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewProjectBuild_Override(p ProjectBuild, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjectBuild",
		[]interface{}{project},
		p,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_ProjectBuild) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_ProjectBuild) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_ProjectBuild) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Project`.
// Experimental.
type ProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *LoggerOptions `json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent Project `json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *ProjenrcOptions `json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
}

// Which type of project this is.
// Deprecated: no longer supported at the base project level
type ProjectType string

const (
	ProjectType_UNKNOWN ProjectType = "UNKNOWN"
	ProjectType_LIB ProjectType = "LIB"
	ProjectType_APP ProjectType = "APP"
)

// Programmatic API for projen.
// Experimental.
type Projects interface {
}

// The jsii proxy struct for Projects
type jsiiProxy_Projects struct {
	_ byte // padding
}

// Creates a new project with defaults.
//
// This function creates the project type in-process (with in VM) and calls
// `.synth()` on it (if `options.synth` is not `false`).
//
// At the moment, it also generates a `.projenrc.js` file with the same code
// that was just executed. In the future, this will also be done by the project
// type, so we can easily support multiple languages of projenrc.
// Experimental.
func Projects_CreateProject(options *CreateProjectOptions) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"projen.Projects",
		"createProject",
		[]interface{}{options},
	)
}

// Sets up a project to use JSON for projenrc.
// Experimental.
type Projenrc interface {
	Component
	Project() Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Projenrc) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrc(project Project, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.Projenrc",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project Project, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Projenrc",
		[]interface{}{project, options},
		p,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_Projenrc) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Experimental.
	Filename *string `json:"filename" yaml:"filename"`
}

// Resolve options.
// Experimental.
type ResolveOptions struct {
	// Context arguments.
	// Experimental.
	Args *[]interface{} `json:"args" yaml:"args"`
	// Omits empty arrays and objects.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// A Make rule.
// Experimental.
type Rule struct {
	// Files to be created or updated by this rule.
	//
	// If the rule is phony then instead this represents the command's name(s).
	// Experimental.
	Targets *[]*string `json:"targets" yaml:"targets"`
	// Marks whether the target is phony.
	// Experimental.
	Phony *bool `json:"phony" yaml:"phony"`
	// Files that are used as inputs to create a target.
	// Experimental.
	Prerequisites *[]*string `json:"prerequisites" yaml:"prerequisites"`
	// Commands that are run (using prerequisites as inputs) to create a target.
	// Experimental.
	Recipe *[]*string `json:"recipe" yaml:"recipe"`
}

// Renders the given files into the directory if the directory does not exist.
//
// Use this to create sample code files
// Experimental.
type SampleDir interface {
	Component
	Project() Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for SampleDir
type jsiiProxy_SampleDir struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SampleDir) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Create sample files in the given directory if the given directory does not exist.
// Experimental.
func NewSampleDir(project Project, dir *string, options *SampleDirOptions) SampleDir {
	_init_.Initialize()

	j := jsiiProxy_SampleDir{}

	_jsii_.Create(
		"projen.SampleDir",
		[]interface{}{project, dir, options},
		&j,
	)

	return &j
}

// Create sample files in the given directory if the given directory does not exist.
// Experimental.
func NewSampleDir_Override(s SampleDir, project Project, dir *string, options *SampleDirOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleDir",
		[]interface{}{project, dir, options},
		s,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_SampleDir) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_SampleDir) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (s *jsiiProxy_SampleDir) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// SampleDir options.
// Experimental.
type SampleDirOptions struct {
	// The files to render into the directory.
	//
	// These files get added after
	// any files from `source` if that option is specified (replacing if names
	// overlap).
	// Experimental.
	Files *map[string]*string `json:"files" yaml:"files"`
	// Absolute path to a directory to copy files from (does not need to be text files).
	//
	// If your project is typescript-based and has configured `testdir` to be a
	// subdirectory of `src`, sample files should outside of the `src` directory
	// otherwise they may not be copied. For example:
	// ```
	// new SampleDir(this, 'public', { source: path.join(__dirname, '..', 'sample-assets') });
	// ```
	// Experimental.
	SourceDir *string `json:"sourceDir" yaml:"sourceDir"`
}

// Produces a file with the given contents but only once, if the file doesn't already exist.
//
// Use this for creating example code files or other resources.
// Experimental.
type SampleFile interface {
	Component
	Project() Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for SampleFile
type jsiiProxy_SampleFile struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SampleFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Creates a new SampleFile object.
// Experimental.
func NewSampleFile(project Project, filePath *string, options *SampleFileOptions) SampleFile {
	_init_.Initialize()

	j := jsiiProxy_SampleFile{}

	_jsii_.Create(
		"projen.SampleFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Creates a new SampleFile object.
// Experimental.
func NewSampleFile_Override(s SampleFile, project Project, filePath *string, options *SampleFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleFile",
		[]interface{}{project, filePath, options},
		s,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_SampleFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_SampleFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (s *jsiiProxy_SampleFile) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// Options for the SampleFile object.
// Experimental.
type SampleFileOptions struct {
	// The contents of the file to write.
	// Experimental.
	Contents *string `json:"contents" yaml:"contents"`
	// Absolute path to a file to copy the contents from (does not need to be a text file).
	//
	// If your project is Typescript-based and has configured `testdir` to be a
	// subdirectory of `src`, sample files should outside of the `src` directory,
	// otherwise they may not be copied. For example:
	// ```
	// new SampleFile(this, 'assets/icon.png', { source: path.join(__dirname, '..', 'sample-assets', 'icon.png') });
	// ```
	// Experimental.
	SourcePath *string `json:"sourcePath" yaml:"sourcePath"`
}

// Represents a README.md sample file. You are expected to manage this file after creation.
// Experimental.
type SampleReadme interface {
	SampleFile
	Project() Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for SampleReadme
type jsiiProxy_SampleReadme struct {
	jsiiProxy_SampleFile
}

func (j *jsiiProxy_SampleReadme) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSampleReadme(project Project, props *SampleReadmeProps) SampleReadme {
	_init_.Initialize()

	j := jsiiProxy_SampleReadme{}

	_jsii_.Create(
		"projen.SampleReadme",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSampleReadme_Override(s SampleReadme, project Project, props *SampleReadmeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleReadme",
		[]interface{}{project, props},
		s,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_SampleReadme) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_SampleReadme) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (s *jsiiProxy_SampleReadme) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// SampleReadme Properties.
// Experimental.
type SampleReadmeProps struct {
	// The contents.
	// Experimental.
	Contents *string `json:"contents" yaml:"contents"`
	// The name of the README.md file.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Filename *string `json:"filename" yaml:"filename"`
}

// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
type Semver interface {
	Mode() *string
	Spec() *string
	Version() *string
}

// The jsii proxy struct for Semver
type jsiiProxy_Semver struct {
	_ byte // padding
}

func (j *jsiiProxy_Semver) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semver) Spec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semver) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Accept any minor version.
//
// >= version
// < next major version
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Caret(version *string) Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"caret",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Latest version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Latest() Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"latest",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Of(spec *string) Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"of",
		[]interface{}{spec},
		&returns,
	)

	return returns
}

// Accept only an exact version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Pinned(version *string) Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"pinned",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Accept patches.
//
// >= version
// < next minor version
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Tilde(version *string) Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"tilde",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Represents a source file.
// Experimental.
type SourceCode interface {
	Component
	FilePath() *string
	Project() Project
	Close(code *string)
	Line(code *string)
	Open(code *string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for SourceCode
type jsiiProxy_SourceCode struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SourceCode) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceCode) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSourceCode(project Project, filePath *string, options *SourceCodeOptions) SourceCode {
	_init_.Initialize()

	j := jsiiProxy_SourceCode{}

	_jsii_.Create(
		"projen.SourceCode",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSourceCode_Override(s SourceCode, project Project, filePath *string, options *SourceCodeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SourceCode",
		[]interface{}{project, filePath, options},
		s,
	)
}

// Decreases the indentation level and closes a code block.
// Experimental.
func (s *jsiiProxy_SourceCode) Close(code *string) {
	_jsii_.InvokeVoid(
		s,
		"close",
		[]interface{}{code},
	)
}

// Emit a line of code.
// Experimental.
func (s *jsiiProxy_SourceCode) Line(code *string) {
	_jsii_.InvokeVoid(
		s,
		"line",
		[]interface{}{code},
	)
}

// Opens a code block and increases the indentation level.
// Experimental.
func (s *jsiiProxy_SourceCode) Open(code *string) {
	_jsii_.InvokeVoid(
		s,
		"open",
		[]interface{}{code},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_SourceCode) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_SourceCode) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (s *jsiiProxy_SourceCode) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `SourceCodeFile`.
// Experimental.
type SourceCodeOptions struct {
	// Indentation size.
	// Experimental.
	Indent *float64 `json:"indent" yaml:"indent"`
}

// A task that can be performed on the project.
//
// Modeled as a series of shell
// commands and subtasks.
// Experimental.
type Task interface {
	Condition() *string
	Description() *string
	SetDescription(val *string)
	Name() *string
	Steps() *[]*TaskStep
	Builtin(name *string)
	Env(name *string, value *string)
	Exec(command *string, options *TaskStepOptions)
	Lock()
	Prepend(shell *string, options *TaskStepOptions)
	PrependExec(shell *string, options *TaskStepOptions)
	PrependSay(message *string, options *TaskStepOptions)
	PrependSpawn(subtask Task, options *TaskStepOptions)
	Reset(command *string, options *TaskStepOptions)
	Say(message *string, options *TaskStepOptions)
	Spawn(subtask Task, options *TaskStepOptions)
}

// The jsii proxy struct for Task
type jsiiProxy_Task struct {
	_ byte // padding
}

func (j *jsiiProxy_Task) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Steps() *[]*TaskStep {
	var returns *[]*TaskStep
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}


// Experimental.
func NewTask(name *string, props *TaskOptions) Task {
	_init_.Initialize()

	j := jsiiProxy_Task{}

	_jsii_.Create(
		"projen.Task",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTask_Override(t Task, name *string, props *TaskOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Task",
		[]interface{}{name, props},
		t,
	)
}

func (j *jsiiProxy_Task) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

// Execute a builtin task.
//
// Builtin tasks are programs bundled as part of projen itself and used as
// helpers for various components.
//
// In the future we should support built-in tasks from external modules.
// Experimental.
func (t *jsiiProxy_Task) Builtin(name *string) {
	_jsii_.InvokeVoid(
		t,
		"builtin",
		[]interface{}{name},
	)
}

// Adds an environment variable to this task.
// Experimental.
func (t *jsiiProxy_Task) Env(name *string, value *string) {
	_jsii_.InvokeVoid(
		t,
		"env",
		[]interface{}{name, value},
	)
}

// Executes a shell command.
// Experimental.
func (t *jsiiProxy_Task) Exec(command *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"exec",
		[]interface{}{command, options},
	)
}

// Forbid additional changes to this task.
// Experimental.
func (t *jsiiProxy_Task) Lock() {
	_jsii_.InvokeVoid(
		t,
		"lock",
		nil, // no parameters
	)
}

// Adds a command at the beginning of the task.
// Deprecated: use `prependExec()`
func (t *jsiiProxy_Task) Prepend(shell *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"prepend",
		[]interface{}{shell, options},
	)
}

// Adds a command at the beginning of the task.
// Experimental.
func (t *jsiiProxy_Task) PrependExec(shell *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"prependExec",
		[]interface{}{shell, options},
	)
}

// Says something at the beginning of the task.
// Experimental.
func (t *jsiiProxy_Task) PrependSay(message *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"prependSay",
		[]interface{}{message, options},
	)
}

// Adds a spawn instruction at the beginning of the task.
// Experimental.
func (t *jsiiProxy_Task) PrependSpawn(subtask Task, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"prependSpawn",
		[]interface{}{subtask, options},
	)
}

// Reset the task so it no longer has any commands.
// Experimental.
func (t *jsiiProxy_Task) Reset(command *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"reset",
		[]interface{}{command, options},
	)
}

// Say something.
// Experimental.
func (t *jsiiProxy_Task) Say(message *string, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"say",
		[]interface{}{message, options},
	)
}

// Spawns a sub-task.
// Experimental.
func (t *jsiiProxy_Task) Spawn(subtask Task, options *TaskStepOptions) {
	_jsii_.InvokeVoid(
		t,
		"spawn",
		[]interface{}{subtask, options},
	)
}

// Experimental.
type TaskCommonOptions struct {
	// A shell command which determines if the this task should be executed.
	//
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition *string `json:"condition" yaml:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd *string `json:"cwd" yaml:"cwd"`
	// The description of this build command.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Defines environment variables for the execution of this task.
	//
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// A set of environment variables that must be defined in order to execute this task.
	//
	// Task execution will fail if one of these is not defined.
	// Experimental.
	RequiredEnv *[]*string `json:"requiredEnv" yaml:"requiredEnv"`
}

// Experimental.
type TaskOptions struct {
	// A shell command which determines if the this task should be executed.
	//
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition *string `json:"condition" yaml:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd *string `json:"cwd" yaml:"cwd"`
	// The description of this build command.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Defines environment variables for the execution of this task.
	//
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// A set of environment variables that must be defined in order to execute this task.
	//
	// Task execution will fail if one of these is not defined.
	// Experimental.
	RequiredEnv *[]*string `json:"requiredEnv" yaml:"requiredEnv"`
	// Shell command to execute as the first command of the task.
	// Experimental.
	Exec *string `json:"exec" yaml:"exec"`
}

// The runtime component of the tasks engine.
// Experimental.
type TaskRuntime interface {
	Manifest() *TasksManifest
	Tasks() *[]*TaskSpec
	Workdir() *string
	RunTask(name *string, parents *[]*string)
	TryFindTask(name *string) *TaskSpec
}

// The jsii proxy struct for TaskRuntime
type jsiiProxy_TaskRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_TaskRuntime) Manifest() *TasksManifest {
	var returns *TasksManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskRuntime) Tasks() *[]*TaskSpec {
	var returns *[]*TaskSpec
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskRuntime) Workdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workdir",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaskRuntime(workdir *string) TaskRuntime {
	_init_.Initialize()

	j := jsiiProxy_TaskRuntime{}

	_jsii_.Create(
		"projen.TaskRuntime",
		[]interface{}{workdir},
		&j,
	)

	return &j
}

// Experimental.
func NewTaskRuntime_Override(t TaskRuntime, workdir *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.TaskRuntime",
		[]interface{}{workdir},
		t,
	)
}

// Runs the task.
// Experimental.
func (t *jsiiProxy_TaskRuntime) RunTask(name *string, parents *[]*string) {
	_jsii_.InvokeVoid(
		t,
		"runTask",
		[]interface{}{name, parents},
	)
}

// Find a task by name, or `undefined` if not found.
// Experimental.
func (t *jsiiProxy_TaskRuntime) TryFindTask(name *string) *TaskSpec {
	var returns *TaskSpec

	_jsii_.Invoke(
		t,
		"tryFindTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Specification of a single task.
// Experimental.
type TaskSpec struct {
	// A shell command which determines if the this task should be executed.
	//
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition *string `json:"condition" yaml:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd *string `json:"cwd" yaml:"cwd"`
	// The description of this build command.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Defines environment variables for the execution of this task.
	//
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// A set of environment variables that must be defined in order to execute this task.
	//
	// Task execution will fail if one of these is not defined.
	// Experimental.
	RequiredEnv *[]*string `json:"requiredEnv" yaml:"requiredEnv"`
	// Task name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Task steps.
	// Experimental.
	Steps *[]*TaskStep `json:"steps" yaml:"steps"`
}

// A single step within a task.
//
// The step could either be  the execution of a
// shell command or execution of a sub-task, by name.
// Experimental.
type TaskStep struct {
	// The working directory for this step.
	// Experimental.
	Cwd *string `json:"cwd" yaml:"cwd"`
	// Step name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The name of a built-in task to execute.
	//
	// Built-in tasks are node.js programs baked into the projen module and as
	// component runtime helpers.
	//
	// The name is a path relative to the projen lib/ directory (without the .task.js extension).
	// For example, if your built in builtin task is under `src/release/resolve-version.task.ts`,
	// then this would be `release/resolve-version`.
	// Experimental.
	Builtin *string `json:"builtin" yaml:"builtin"`
	// Shell command to execute.
	// Experimental.
	Exec *string `json:"exec" yaml:"exec"`
	// Print a message.
	// Experimental.
	Say *string `json:"say" yaml:"say"`
	// Subtask to execute.
	// Experimental.
	Spawn *string `json:"spawn" yaml:"spawn"`
}

// Options for task steps.
// Experimental.
type TaskStepOptions struct {
	// The working directory for this step.
	// Experimental.
	Cwd *string `json:"cwd" yaml:"cwd"`
	// Step name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Defines project tasks.
//
// Tasks extend the projen CLI by adding subcommands to it. Task definitions are
// synthesized into `.projen/tasks.json`.
// Experimental.
type Tasks interface {
	Component
	All() *[]Task
	Env() *map[string]*string
	Project() Project
	AddEnvironment(name *string, value *string)
	AddTask(name *string, options *TaskOptions) Task
	PostSynthesize()
	PreSynthesize()
	RemoveTask(name *string) Task
	Synthesize()
	TryFind(name *string) Task
}

// The jsii proxy struct for Tasks
type jsiiProxy_Tasks struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Tasks) All() *[]Task {
	var returns *[]Task
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tasks) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tasks) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewTasks(project Project) Tasks {
	_init_.Initialize()

	j := jsiiProxy_Tasks{}

	_jsii_.Create(
		"projen.Tasks",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewTasks_Override(t Tasks, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Tasks",
		[]interface{}{project},
		t,
	)
}

func Tasks_MANIFEST_FILE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Tasks",
		"MANIFEST_FILE",
		&returns,
	)
	return returns
}

// Adds global environment.
// Experimental.
func (t *jsiiProxy_Tasks) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		t,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

// Adds a task to a project.
// Experimental.
func (t *jsiiProxy_Tasks) AddTask(name *string, options *TaskOptions) Task {
	var returns Task

	_jsii_.Invoke(
		t,
		"addTask",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (t *jsiiProxy_Tasks) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (t *jsiiProxy_Tasks) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (t *jsiiProxy_Tasks) RemoveTask(name *string) Task {
	var returns Task

	_jsii_.Invoke(
		t,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Synthesizes files to the project output directory.
// Experimental.
func (t *jsiiProxy_Tasks) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

// Finds a task by name.
//
// Returns `undefined` if the task cannot be found.
// Experimental.
func (t *jsiiProxy_Tasks) TryFind(name *string) Task {
	var returns Task

	_jsii_.Invoke(
		t,
		"tryFind",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Schema for `tasks.json`.
// Experimental.
type TasksManifest struct {
	// Environment for all tasks.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// All tasks available for this project.
	// Experimental.
	Tasks *map[string]*TaskSpec `json:"tasks" yaml:"tasks"`
}

// A Testing static class with a .synth helper for getting a snapshots of construct outputs. Useful for snapshot testing with Jest.
//
// TODO: EXAMPLE
//
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Produces a simple JS object that represents the contents of the projects with field names being file paths.
//
// Returns: : any }
// Experimental.
func Testing_Synth(project Project) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"projen.Testing",
		"synth",
		[]interface{}{project},
		&returns,
	)

	return returns
}

// A text file.
// Experimental.
type TextFile interface {
	FileBase
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddLine(line *string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_arg IResolver) *string
}

// The jsii proxy struct for TextFile
type jsiiProxy_TextFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_TextFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Defines a text file.
// Experimental.
func NewTextFile(project Project, filePath *string, options *TextFileOptions) TextFile {
	_init_.Initialize()

	j := jsiiProxy_TextFile{}

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Defines a text file.
// Experimental.
func NewTextFile_Override(t TextFile, project Project, filePath *string, options *TextFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{project, filePath, options},
		t,
	)
}

func (j *jsiiProxy_TextFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_TextFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func TextFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.TextFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Adds a line to the text file.
// Experimental.
func (t *jsiiProxy_TextFile) AddLine(line *string) {
	_jsii_.InvokeVoid(
		t,
		"addLine",
		[]interface{}{line},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (t *jsiiProxy_TextFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (t *jsiiProxy_TextFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (t *jsiiProxy_TextFile) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (t *jsiiProxy_TextFile) SynthesizeContent(_arg IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Options for `TextFile`.
// Experimental.
type TextFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// The contents of the text file.
	//
	// You can use `addLine()` to append lines.
	// Experimental.
	Lines *[]*string `json:"lines" yaml:"lines"`
}

// Represents a TOML file.
// Experimental.
type TomlFile interface {
	ObjectFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for TomlFile
type jsiiProxy_TomlFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_TomlFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TomlFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewTomlFile(project Project, filePath *string, options *TomlFileOptions) TomlFile {
	_init_.Initialize()

	j := jsiiProxy_TomlFile{}

	_jsii_.Create(
		"projen.TomlFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTomlFile_Override(t TomlFile, project Project, filePath *string, options *TomlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.TomlFile",
		[]interface{}{project, filePath, options},
		t,
	)
}

func (j *jsiiProxy_TomlFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_TomlFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_TomlFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func TomlFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.TomlFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (t *jsiiProxy_TomlFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		t,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (t *jsiiProxy_TomlFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (t *jsiiProxy_TomlFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (t *jsiiProxy_TomlFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (t *jsiiProxy_TomlFile) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (t *jsiiProxy_TomlFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `TomlFile`.
// Experimental.
type TomlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// Experimental.
type Version interface {
	Component
	BumpTask() Task
	ChangelogFileName() *string
	Project() Project
	ReleaseTagFileName() *string
	UnbumpTask() Task
	VersionFileName() *string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Version
type jsiiProxy_Version struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Version) BumpTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"bumpTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) ChangelogFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changelogFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) ReleaseTagFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTagFileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) UnbumpTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"unbumpTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) VersionFileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionFileName",
		&returns,
	)
	return returns
}


// Experimental.
func NewVersion(project Project, options *VersionOptions) Version {
	_init_.Initialize()

	j := jsiiProxy_Version{}

	_jsii_.Create(
		"projen.Version",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewVersion_Override(v Version, project Project, options *VersionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Version",
		[]interface{}{project, options},
		v,
	)
}

func Version_STANDARD_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Version",
		"STANDARD_VERSION",
		&returns,
	)
	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (v *jsiiProxy_Version) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (v *jsiiProxy_Version) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (v *jsiiProxy_Version) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Version`.
// Experimental.
type VersionOptions struct {
	// The name of the directory into which `changelog.md` and `version.txt` files are emitted.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// A name of a .json file to set the `version` field in after a bump.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	VersionInputFile *string `json:"versionInputFile" yaml:"versionInputFile"`
	// Custom configuration for versionrc file used by standard-release.
	// Experimental.
	VersionrcOptions *map[string]interface{} `json:"versionrcOptions" yaml:"versionrcOptions"`
}

// Represents an XML file.
//
// Objects passed in will be synthesized using the npm "xml" library.
// See: https://www.npmjs.com/package/xml
//
// Experimental.
type XmlFile interface {
	ObjectFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for XmlFile
type jsiiProxy_XmlFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_XmlFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XmlFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewXmlFile(project Project, filePath *string, options *XmlFileOptions) XmlFile {
	_init_.Initialize()

	j := jsiiProxy_XmlFile{}

	_jsii_.Create(
		"projen.XmlFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewXmlFile_Override(x XmlFile, project Project, filePath *string, options *XmlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.XmlFile",
		[]interface{}{project, filePath, options},
		x,
	)
}

func (j *jsiiProxy_XmlFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_XmlFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_XmlFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func XmlFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.XmlFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (x *jsiiProxy_XmlFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		x,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (x *jsiiProxy_XmlFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		x,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (x *jsiiProxy_XmlFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		x,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (x *jsiiProxy_XmlFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		x,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (x *jsiiProxy_XmlFile) Synthesize() {
	_jsii_.InvokeVoid(
		x,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (x *jsiiProxy_XmlFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `XmlFile`.
// Experimental.
type XmlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

// Represents a YAML file.
// Experimental.
type YamlFile interface {
	ObjectFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *bool
	SetMarker(val *bool)
	OmitEmpty() *bool
	Path() *string
	Project() Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddDeletionOverride(path *string)
	AddOverride(path *string, value interface{})
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for YamlFile
type jsiiProxy_YamlFile struct {
	jsiiProxy_ObjectFile
}

func (j *jsiiProxy_YamlFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Marker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) OmitEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_YamlFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewYamlFile(project Project, filePath *string, options *YamlFileOptions) YamlFile {
	_init_.Initialize()

	j := jsiiProxy_YamlFile{}

	_jsii_.Create(
		"projen.YamlFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewYamlFile_Override(y YamlFile, project Project, filePath *string, options *YamlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.YamlFile",
		[]interface{}{project, filePath, options},
		y,
	)
}

func (j *jsiiProxy_YamlFile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_YamlFile) SetMarker(val *bool) {
	_jsii_.Set(
		j,
		"marker",
		val,
	)
}

func (j *jsiiProxy_YamlFile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func YamlFile_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.YamlFile",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (y *jsiiProxy_YamlFile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		y,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Adds an override to the synthesized object file.
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// project.tsconfig.file.addOverride('compilerOptions.alwaysStrict', true);
// project.tsconfig.file.addOverride('compilerOptions.lib', ['dom', 'dom.iterable', 'esnext']);
// ```
// would add the overrides
// ```json
// "compilerOptions": {
//    "alwaysStrict": true,
//    "lib": [
//      "dom",
//      "dom.iterable",
//      "esnext"
//    ]
//    ...
// }
// ...
// ```
// Experimental.
func (y *jsiiProxy_YamlFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		y,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (y *jsiiProxy_YamlFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (y *jsiiProxy_YamlFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (y *jsiiProxy_YamlFile) Synthesize() {
	_jsii_.InvokeVoid(
		y,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (y *jsiiProxy_YamlFile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

// Options for `JsonFile`.
// Experimental.
type YamlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly" yaml:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `json:"marker" yaml:"marker"`
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty" yaml:"omitEmpty"`
}

