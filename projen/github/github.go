package github

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen"
)

// Class interface
type AutoMergeIface interface {
	GetProject() projen.ProjectIface
	GetAutoMergeLabel() string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Sets up mergify to merging approved pull requests.
// 
// If `buildJob` is specified, the specified GitHub workflow job ID is required
// to succeed in order for the PR to be merged.
// 
// `approvedReviews` specified the number of code review approvals required for
// the PR to be merged.
// Experimental.
// Struct proxy
type AutoMerge struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Experimental.
	AutoMergeLabel string `json:"autoMergeLabel"`
}

func (a *AutoMerge) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		a,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AutoMerge) GetAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		a,
		"autoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewAutoMerge(project projen.ProjectIface, options AutoMergeOptionsIface) AutoMergeIface {
	_init_.Initialize()
	self := AutoMerge{}
	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (a *AutoMerge) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AutoMerge) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AutoMerge) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// AutoMergeOptionsIface is the public interface for the custom type AutoMergeOptions
// Experimental.
type AutoMergeOptionsIface interface {
	GetApprovedReviews() float64
	GetAutoMergeLabel() string
	GetBuildJob() string
}

// Experimental.
// Struct proxy
type AutoMergeOptions struct {
	// Number of approved code reviews.
	// Experimental.
	ApprovedReviews float64 `json:"approvedReviews"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	AutoMergeLabel string `json:"autoMergeLabel"`
	// The GitHub job ID of the build workflow.
	// Experimental.
	BuildJob string `json:"buildJob"`
}

func (a *AutoMergeOptions) GetApprovedReviews() float64 {
	var returns float64
	_jsii_.Get(
		a,
		"approvedReviews",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AutoMergeOptions) GetAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		a,
		"autoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AutoMergeOptions) GetBuildJob() string {
	var returns string
	_jsii_.Get(
		a,
		"buildJob",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type DependabotIface interface {
	GetProject() projen.ProjectIface
	GetConfig() interface{}
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddIgnore(dependencyName string, versions string)
}

// Defines dependabot configuration for node projects.
// 
// Since module versions are managed in projen, the versioning strategy will be
// configured to "lockfile-only" which means that only updates that can be done
// on the lockfile itself will be proposed.
// Experimental.
// Struct proxy
type Dependabot struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The raw dependabot configuration.
	// See: https://docs.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
	//
	// Experimental.
	Config interface{} `json:"config"`
}

func (d *Dependabot) GetProject() projen.ProjectIface {
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

func (d *Dependabot) GetConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		d,
		"config",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewDependabot(github GitHubIface, options DependabotOptionsIface) DependabotIface {
	_init_.Initialize()
	self := Dependabot{}
	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (d *Dependabot) PostSynthesize() {
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

func (d *Dependabot) PreSynthesize() {
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

func (d *Dependabot) Synthesize() {
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

func (d *Dependabot) AddIgnore(dependencyName string, versions string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addIgnore",
		[]interface{}{dependencyName, versions},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// DependabotIgnoreIface is the public interface for the custom type DependabotIgnore
// Experimental.
type DependabotIgnoreIface interface {
	GetDependencyName() string
	GetVersions() []string
}

// You can use the `ignore` option to customize which dependencies are updated.
// 
// The ignore option supports the following options.
// Experimental.
// Struct proxy
type DependabotIgnore struct {
	// Use to ignore updates for dependencies with matching names, optionally using `*` to match zero or more characters.
	// 
	// For Java dependencies, the format of the dependency-name attribute is:
	// `groupId:artifactId`, for example: `org.kohsuke:github-api`.
	// Experimental.
	DependencyName string `json:"dependencyName"`
	// Use to ignore specific versions or ranges of versions.
	// 
	// If you want to
	// define a range, use the standard pattern for the package manager (for
	// example: `^1.0.0` for npm, or `~> 2.0` for Bundler).
	// Experimental.
	Versions []string `json:"versions"`
}

func (d *DependabotIgnore) GetDependencyName() string {
	var returns string
	_jsii_.Get(
		d,
		"dependencyName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DependabotIgnore) GetVersions() []string {
	var returns []string
	_jsii_.Get(
		d,
		"versions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// DependabotOptionsIface is the public interface for the custom type DependabotOptions
// Experimental.
type DependabotOptionsIface interface {
	GetAutoMerge() bool
	GetIgnore() []DependabotIgnoreIface
	GetIgnoreProjen() bool
	GetScheduleInterval() DependabotScheduleInterval
	GetVersioningStrategy() VersioningStrategy
}

// Experimental.
// Struct proxy
type DependabotOptions struct {
	// Automatically merge dependabot PRs if build CI build passes.
	// Experimental.
	AutoMerge bool `json:"autoMerge"`
	// You can use the `ignore` option to customize which dependencies are updated.
	// 
	// The ignore option supports the following options.
	// Experimental.
	Ignore []DependabotIgnoreIface `json:"ignore"`
	// Ignores updates to `projen`.
	// 
	// This is required since projen updates may cause changes in committed files
	// and anti-tamper checks will fail.
	// 
	// Projen upgrades are covered through the `ProjenUpgrade` class.
	// Experimental.
	IgnoreProjen bool `json:"ignoreProjen"`
	// How often to check for new versions and raise pull requests.
	// Experimental.
	ScheduleInterval DependabotScheduleInterval `json:"scheduleInterval"`
	// The strategy to use when edits manifest and lock files.
	// Experimental.
	VersioningStrategy VersioningStrategy `json:"versioningStrategy"`
}

func (d *DependabotOptions) GetAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		d,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DependabotOptions) GetIgnore() []DependabotIgnoreIface {
	var returns []DependabotIgnoreIface
	_jsii_.Get(
		d,
		"ignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependabotIgnoreIface)(nil)).Elem(): reflect.TypeOf((*DependabotIgnore)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DependabotOptions) GetIgnoreProjen() bool {
	var returns bool
	_jsii_.Get(
		d,
		"ignoreProjen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DependabotOptions) GetScheduleInterval() DependabotScheduleInterval {
	var returns DependabotScheduleInterval
	_jsii_.Get(
		d,
		"scheduleInterval",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependabotScheduleInterval)(nil)).Elem(): reflect.TypeOf((*DependabotScheduleInterval)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DependabotOptions) GetVersioningStrategy() VersioningStrategy {
	var returns VersioningStrategy
	_jsii_.Get(
		d,
		"versioningStrategy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VersioningStrategy)(nil)).Elem(): reflect.TypeOf((*VersioningStrategy)(nil)).Elem(),
		},
	)
	return returns
}


// How often to check for new versions and raise pull requests for version updates.
// Experimental.
type DependabotScheduleInterval string

const (
	DependabotScheduleIntervalDaily DependabotScheduleInterval = "DAILY"
	DependabotScheduleIntervalWeekly DependabotScheduleInterval = "WEEKLY"
	DependabotScheduleIntervalMonthly DependabotScheduleInterval = "MONTHLY"
)

// Class interface
type GitHubIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddDependabot(options DependabotOptionsIface) DependabotIface
	AddMergifyRules(rules MergifyRuleIface)
	AddPullRequestTemplate(content string) PullRequestTemplateIface
	AddWorkflow(name string) GithubWorkflowIface
}

// Experimental.
// Struct proxy
type GitHub struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (g *GitHub) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		g,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewGitHub(project projen.ProjectIface) GitHubIface {
	_init_.Initialize()
	self := GitHub{}
	_jsii_.Create(
		"projen.github.GitHub",
		[]interface{}{project},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (g *GitHub) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GitHub) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GitHub) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GitHub) AddDependabot(options DependabotOptionsIface) DependabotIface {
	var returns DependabotIface
	_jsii_.Invoke(
		g,
		"addDependabot",
		[]interface{}{options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependabotIface)(nil)).Elem(): reflect.TypeOf((*Dependabot)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitHub) AddMergifyRules(rules MergifyRuleIface) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addMergifyRules",
		[]interface{}{rules},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GitHub) AddPullRequestTemplate(content string) PullRequestTemplateIface {
	var returns PullRequestTemplateIface
	_jsii_.Invoke(
		g,
		"addPullRequestTemplate",
		[]interface{}{content},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PullRequestTemplateIface)(nil)).Elem(): reflect.TypeOf((*PullRequestTemplate)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitHub) AddWorkflow(name string) GithubWorkflowIface {
	var returns GithubWorkflowIface
	_jsii_.Invoke(
		g,
		"addWorkflow",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

// Class interface
type GithubWorkflowIface interface {
	GetProject() projen.ProjectIface
	GetFile() projen.YamlFileIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddJobs(jobs map[string]interface{})
	On(events map[string]interface{})
}

// Experimental.
// Struct proxy
type GithubWorkflow struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Experimental.
	File projen.YamlFileIface `json:"file"`
}

func (g *GithubWorkflow) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		g,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GithubWorkflow) GetFile() projen.YamlFileIface {
	var returns projen.YamlFileIface
	_jsii_.Get(
		g,
		"file",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.YamlFileIface)(nil)).Elem(): reflect.TypeOf((*projen.YamlFile)(nil)).Elem(),
		},
	)
	return returns
}


func NewGithubWorkflow(github GitHubIface, name string) GithubWorkflowIface {
	_init_.Initialize()
	self := GithubWorkflow{}
	_jsii_.Create(
		"projen.github.GithubWorkflow",
		[]interface{}{github, name},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (g *GithubWorkflow) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GithubWorkflow) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GithubWorkflow) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GithubWorkflow) AddJobs(jobs map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addJobs",
		[]interface{}{jobs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *GithubWorkflow) On(events map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"on",
		[]interface{}{events},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type MergifyIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddRule(rule MergifyRuleIface)
}

// Experimental.
// Struct proxy
type Mergify struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (m *Mergify) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		m,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewMergify(github GitHubIface, options MergifyOptionsIface) MergifyIface {
	_init_.Initialize()
	self := Mergify{}
	_jsii_.Create(
		"projen.github.Mergify",
		[]interface{}{github, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (m *Mergify) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *Mergify) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *Mergify) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *Mergify) AddRule(rule MergifyRuleIface) {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"addRule",
		[]interface{}{rule},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// MergifyOptionsIface is the public interface for the custom type MergifyOptions
// Experimental.
type MergifyOptionsIface interface {
	GetRules() []MergifyRuleIface
}

// Experimental.
// Struct proxy
type MergifyOptions struct {
	// Experimental.
	Rules []MergifyRuleIface `json:"rules"`
}

func (m *MergifyOptions) GetRules() []MergifyRuleIface {
	var returns []MergifyRuleIface
	_jsii_.Get(
		m,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MergifyRuleIface)(nil)).Elem(): reflect.TypeOf((*MergifyRule)(nil)).Elem(),
		},
	)
	return returns
}


// MergifyRuleIface is the public interface for the custom type MergifyRule
// Experimental.
type MergifyRuleIface interface {
	GetActions() map[string]interface{}
	GetConditions() []string
	GetName() string
}

// Experimental.
// Struct proxy
type MergifyRule struct {
	// Experimental.
	Actions map[string]interface{} `json:"actions"`
	// Experimental.
	Conditions []string `json:"conditions"`
	// Experimental.
	Name string `json:"name"`
}

func (m *MergifyRule) GetActions() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		m,
		"actions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (m *MergifyRule) GetConditions() []string {
	var returns []string
	_jsii_.Get(
		m,
		"conditions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MergifyRule) GetName() string {
	var returns string
	_jsii_.Get(
		m,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type PullRequestTemplateIface interface {
	GetProject() projen.ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_ projen.IResolverIface) string
	AddLine(line string)
}

// Template for GitHub pull requests.
// Experimental.
// Struct proxy
type PullRequestTemplate struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (p *PullRequestTemplate) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		p,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PullRequestTemplate) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		p,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PullRequestTemplate) GetPath() string {
	var returns string
	_jsii_.Get(
		p,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PullRequestTemplate) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		p,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PullRequestTemplate) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		p,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewPullRequestTemplate(github GitHubIface, options PullRequestTemplateOptionsIface) PullRequestTemplateIface {
	_init_.Initialize()
	self := PullRequestTemplate{}
	_jsii_.Create(
		"projen.github.PullRequestTemplate",
		[]interface{}{github, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *PullRequestTemplate) SetExecutable(val bool) {
	_jsii_.Set(
		p,
		"executable",
		val,
	)
}

func (p *PullRequestTemplate) SetReadonly(val bool) {
	_jsii_.Set(
		p,
		"readonly",
		val,
	)
}

func PullRequestTemplate_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.github.PullRequestTemplate",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PullRequestTemplate) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *PullRequestTemplate) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *PullRequestTemplate) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *PullRequestTemplate) SynthesizeContent(_ projen.IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		p,
		"synthesizeContent",
		[]interface{}{_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PullRequestTemplate) AddLine(line string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addLine",
		[]interface{}{line},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// PullRequestTemplateOptionsIface is the public interface for the custom type PullRequestTemplateOptions
// Experimental.
type PullRequestTemplateOptionsIface interface {
	GetLines() []string
}

// Options for `PullRequestTemplate`.
// Experimental.
// Struct proxy
type PullRequestTemplateOptions struct {
	// The contents of the template.
	// 
	// You can use `addLine()` to add additional lines.
	// Experimental.
	Lines []string `json:"lines"`
}

func (p *PullRequestTemplateOptions) GetLines() []string {
	var returns []string
	_jsii_.Get(
		p,
		"lines",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// The strategy to use when edits manifest and lock files.
// Experimental.
type VersioningStrategy string

const (
	VersioningStrategyLockfileOnly VersioningStrategy = "LOCKFILE_ONLY"
	VersioningStrategyAuto VersioningStrategy = "AUTO"
	VersioningStrategyWiden VersioningStrategy = "WIDEN"
	VersioningStrategyIncrease VersioningStrategy = "INCREASE"
	VersioningStrategyIncreaseIfNecessary VersioningStrategy = "INCREASE_IF_NECESSARY"
)

