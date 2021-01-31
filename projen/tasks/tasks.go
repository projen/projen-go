package tasks

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen"
)

// Class interface
type TaskIface interface {
	GetName() string
	GetSteps() []TaskStepIface
	GetCategory() TaskCategory
	GetCondition() string
	GetDescription() string
	Env(name string, value string)
	Exec(command string, options TaskStepOptionsIface)
	Prepend(shell string, options TaskStepOptionsIface)
	PrependExec(shell string, options TaskStepOptionsIface)
	PrependSay(message string, options TaskStepOptionsIface)
	PrependSpawn(subtask TaskIface, options TaskStepOptionsIface)
	Reset(command string)
	Say(message string, options TaskStepOptionsIface)
	Spawn(subtask TaskIface, options TaskStepOptionsIface)
	ToShellCommand() string
}

// A task that can be performed on the project.
// 
// Modeled as a series of shell
// commands and subtasks.
// Experimental.
// Struct proxy
type Task struct {
	// Task name.
	// Experimental.
	Name string `json:"name"`
	// Returns an immutable copy of all the step specifications of the task.
	// Experimental.
	Steps []TaskStepIface `json:"steps"`
	// The start menu category of the task.
	// Experimental.
	Category TaskCategory `json:"category"`
	// A command to execute which determines if the task should be skipped.
	// 
	// If it
	// returns a zero exit code, the task will not be executed.
	// Experimental.
	Condition string `json:"condition"`
	// The description of the task.
	// Experimental.
	Description string `json:"description"`
}

func (t *Task) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *Task) GetSteps() []TaskStepIface {
	var returns []TaskStepIface
	_jsii_.Get(
		t,
		"steps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskStepIface)(nil)).Elem(): reflect.TypeOf((*TaskStep)(nil)).Elem(),
		},
	)
	return returns
}

func (t *Task) GetCategory() TaskCategory {
	var returns TaskCategory
	_jsii_.Get(
		t,
		"category",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskCategory)(nil)).Elem(): reflect.TypeOf((*TaskCategory)(nil)).Elem(),
		},
	)
	return returns
}

func (t *Task) GetCondition() string {
	var returns string
	_jsii_.Get(
		t,
		"condition",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *Task) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewTask(tasks TasksIface, name string, props TaskOptionsIface) TaskIface {
	_init_.Initialize()
	self := Task{}
	_jsii_.Create(
		"projen.tasks.Task",
		[]interface{}{tasks, name, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (t *Task) Env(name string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"env",
		[]interface{}{name, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) Exec(command string, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"exec",
		[]interface{}{command, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) Prepend(shell string, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"prepend",
		[]interface{}{shell, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) PrependExec(shell string, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"prependExec",
		[]interface{}{shell, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) PrependSay(message string, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"prependSay",
		[]interface{}{message, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) PrependSpawn(subtask TaskIface, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"prependSpawn",
		[]interface{}{subtask, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) Reset(command string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"reset",
		[]interface{}{command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) Say(message string, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"say",
		[]interface{}{message, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) Spawn(subtask TaskIface, options TaskStepOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"spawn",
		[]interface{}{subtask, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Task) ToShellCommand() string {
	var returns string
	_jsii_.Invoke(
		t,
		"toShellCommand",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Experimental.
type TaskCategory string

const (
	TaskCategoryBuild TaskCategory = "BUILD"
	TaskCategoryTest TaskCategory = "TEST"
	TaskCategoryRelease TaskCategory = "RELEASE"
	TaskCategoryMaintain TaskCategory = "MAINTAIN"
	TaskCategoryMisc TaskCategory = "MISC"
)

// TaskCommonOptionsIface is the public interface for the custom type TaskCommonOptions
// Experimental.
type TaskCommonOptionsIface interface {
	GetCategory() TaskCategory
	GetCondition() string
	GetCwd() string
	GetDescription() string
	GetEnv() map[string]string
}

// Experimental.
// Struct proxy
type TaskCommonOptions struct {
	// Category for start menu.
	// Experimental.
	Category TaskCategory `json:"category"`
	// A shell command which determines if the this task should be executed.
	// 
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition string `json:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd string `json:"cwd"`
	// The description of this build command.
	// Experimental.
	Description string `json:"description"`
	// Defines environment variables for the execution of this task.
	// 
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env map[string]string `json:"env"`
}

func (t *TaskCommonOptions) GetCategory() TaskCategory {
	var returns TaskCategory
	_jsii_.Get(
		t,
		"category",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskCategory)(nil)).Elem(): reflect.TypeOf((*TaskCategory)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskCommonOptions) GetCondition() string {
	var returns string
	_jsii_.Get(
		t,
		"condition",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskCommonOptions) GetCwd() string {
	var returns string
	_jsii_.Get(
		t,
		"cwd",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskCommonOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskCommonOptions) GetEnv() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// TaskOptionsIface is the public interface for the custom type TaskOptions
// Experimental.
type TaskOptionsIface interface {
	GetCategory() TaskCategory
	GetCondition() string
	GetCwd() string
	GetDescription() string
	GetEnv() map[string]string
	GetExec() string
}

// Experimental.
// Struct proxy
type TaskOptions struct {
	// Category for start menu.
	// Experimental.
	Category TaskCategory `json:"category"`
	// A shell command which determines if the this task should be executed.
	// 
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition string `json:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd string `json:"cwd"`
	// The description of this build command.
	// Experimental.
	Description string `json:"description"`
	// Defines environment variables for the execution of this task.
	// 
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env map[string]string `json:"env"`
	// Shell command to execute as the first command of the task.
	// Experimental.
	Exec string `json:"exec"`
}

func (t *TaskOptions) GetCategory() TaskCategory {
	var returns TaskCategory
	_jsii_.Get(
		t,
		"category",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskCategory)(nil)).Elem(): reflect.TypeOf((*TaskCategory)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskOptions) GetCondition() string {
	var returns string
	_jsii_.Get(
		t,
		"condition",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskOptions) GetCwd() string {
	var returns string
	_jsii_.Get(
		t,
		"cwd",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskOptions) GetEnv() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskOptions) GetExec() string {
	var returns string
	_jsii_.Get(
		t,
		"exec",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type TaskRuntimeIface interface {
	GetManifest() TasksManifestIface
	GetTasks() []TaskSpecIface
	GetWorkdir() string
	RunTask(name string, parents []string)
	TryFindTask(name string) TaskSpecIface
}

// The runtime component of the tasks engine.
// Experimental.
// Struct proxy
type TaskRuntime struct {
	// The contents of tasks.json.
	// Experimental.
	Manifest TasksManifestIface `json:"manifest"`
	// The tasks in this project.
	// Experimental.
	Tasks []TaskSpecIface `json:"tasks"`
	// The root directory of the project and the cwd for executing tasks.
	// Experimental.
	Workdir string `json:"workdir"`
}

func (t *TaskRuntime) GetManifest() TasksManifestIface {
	var returns TasksManifestIface
	_jsii_.Get(
		t,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TasksManifestIface)(nil)).Elem(): reflect.TypeOf((*TasksManifest)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskRuntime) GetTasks() []TaskSpecIface {
	var returns []TaskSpecIface
	_jsii_.Get(
		t,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskSpecIface)(nil)).Elem(): reflect.TypeOf((*TaskSpec)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskRuntime) GetWorkdir() string {
	var returns string
	_jsii_.Get(
		t,
		"workdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewTaskRuntime(workdir string) TaskRuntimeIface {
	_init_.Initialize()
	self := TaskRuntime{}
	_jsii_.Create(
		"projen.tasks.TaskRuntime",
		[]interface{}{workdir},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (t *TaskRuntime) RunTask(name string, parents []string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"runTask",
		[]interface{}{name, parents},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TaskRuntime) TryFindTask(name string) TaskSpecIface {
	var returns TaskSpecIface
	_jsii_.Invoke(
		t,
		"tryFindTask",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskSpecIface)(nil)).Elem(): reflect.TypeOf((*TaskSpec)(nil)).Elem(),
		},
	)
	return returns
}

// TaskSpecIface is the public interface for the custom type TaskSpec
// Experimental.
type TaskSpecIface interface {
	GetCategory() TaskCategory
	GetCondition() string
	GetCwd() string
	GetDescription() string
	GetEnv() map[string]string
	GetName() string
	GetSteps() []TaskStepIface
}

// Specification of a single task.
// Experimental.
// Struct proxy
type TaskSpec struct {
	// Category for start menu.
	// Experimental.
	Category TaskCategory `json:"category"`
	// A shell command which determines if the this task should be executed.
	// 
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition string `json:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd string `json:"cwd"`
	// The description of this build command.
	// Experimental.
	Description string `json:"description"`
	// Defines environment variables for the execution of this task.
	// 
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env map[string]string `json:"env"`
	// Task name.
	// Experimental.
	Name string `json:"name"`
	// Task steps.
	// Experimental.
	Steps []TaskStepIface `json:"steps"`
}

func (t *TaskSpec) GetCategory() TaskCategory {
	var returns TaskCategory
	_jsii_.Get(
		t,
		"category",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskCategory)(nil)).Elem(): reflect.TypeOf((*TaskCategory)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskSpec) GetCondition() string {
	var returns string
	_jsii_.Get(
		t,
		"condition",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskSpec) GetCwd() string {
	var returns string
	_jsii_.Get(
		t,
		"cwd",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskSpec) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskSpec) GetEnv() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TaskSpec) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskSpec) GetSteps() []TaskStepIface {
	var returns []TaskStepIface
	_jsii_.Get(
		t,
		"steps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskStepIface)(nil)).Elem(): reflect.TypeOf((*TaskStep)(nil)).Elem(),
		},
	)
	return returns
}


// TaskStepIface is the public interface for the custom type TaskStep
// Experimental.
type TaskStepIface interface {
	GetCwd() string
	GetName() string
	GetExec() string
	GetSay() string
	GetSpawn() string
}

// A single step within a task.
// 
// The step could either be  the execution of a
// shell command or execution of a sub-task, by name.
// Experimental.
// Struct proxy
type TaskStep struct {
	// The working directory for this step.
	// Experimental.
	Cwd string `json:"cwd"`
	// Step name.
	// Experimental.
	Name string `json:"name"`
	// Shell command to execute.
	// Experimental.
	Exec string `json:"exec"`
	// Print a message.
	// Experimental.
	Say string `json:"say"`
	// Subtask to execute.
	// Experimental.
	Spawn string `json:"spawn"`
}

func (t *TaskStep) GetCwd() string {
	var returns string
	_jsii_.Get(
		t,
		"cwd",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskStep) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskStep) GetExec() string {
	var returns string
	_jsii_.Get(
		t,
		"exec",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskStep) GetSay() string {
	var returns string
	_jsii_.Get(
		t,
		"say",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskStep) GetSpawn() string {
	var returns string
	_jsii_.Get(
		t,
		"spawn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// TaskStepOptionsIface is the public interface for the custom type TaskStepOptions
// Experimental.
type TaskStepOptionsIface interface {
	GetCwd() string
	GetName() string
}

// Options for task steps.
// Experimental.
// Struct proxy
type TaskStepOptions struct {
	// The working directory for this step.
	// Experimental.
	Cwd string `json:"cwd"`
	// Step name.
	// Experimental.
	Name string `json:"name"`
}

func (t *TaskStepOptions) GetCwd() string {
	var returns string
	_jsii_.Get(
		t,
		"cwd",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TaskStepOptions) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type TasksIface interface {
	GetProject() projen.ProjectIface
	GetAll() []TaskIface
	GetEnv() map[string]string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddEnvironment(name string, value string)
	AddTask(name string, options TaskOptionsIface) TaskIface
	TryFind(name string) TaskIface
}

// Defines project tasks.
// 
// Tasks extend the projen CLI by adding subcommands to it. Task definitions are
// synthesized into `.projen/tasks.json`.
// Experimental.
// Struct proxy
type Tasks struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// All tasks.
	// Experimental.
	All []TaskIface `json:"all"`
	// Returns a copy of the currently global environment for this project.
	// Experimental.
	Env map[string]string `json:"env"`
}

func (t *Tasks) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		t,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *Tasks) GetAll() []TaskIface {
	var returns []TaskIface
	_jsii_.Get(
		t,
		"all",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskIface)(nil)).Elem(): reflect.TypeOf((*Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *Tasks) GetEnv() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


func NewTasks(project projen.ProjectIface) TasksIface {
	_init_.Initialize()
	self := Tasks{}
	_jsii_.Create(
		"projen.tasks.Tasks",
		[]interface{}{project},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Tasks_ManifestFile() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.tasks.Tasks",
		"MANIFEST_FILE",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *Tasks) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Tasks) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Tasks) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Tasks) AddEnvironment(name string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addEnvironment",
		[]interface{}{name, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *Tasks) AddTask(name string, options TaskOptionsIface) TaskIface {
	var returns TaskIface
	_jsii_.Invoke(
		t,
		"addTask",
		[]interface{}{name, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskIface)(nil)).Elem(): reflect.TypeOf((*Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *Tasks) TryFind(name string) TaskIface {
	var returns TaskIface
	_jsii_.Invoke(
		t,
		"tryFind",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskIface)(nil)).Elem(): reflect.TypeOf((*Task)(nil)).Elem(),
		},
	)
	return returns
}

// TasksManifestIface is the public interface for the custom type TasksManifest
// Experimental.
type TasksManifestIface interface {
	GetEnv() map[string]string
	GetTasks() map[string]TaskSpecIface
}

// Schema for `tasks.json`.
// Experimental.
// Struct proxy
type TasksManifest struct {
	// Environment for all tasks.
	// Experimental.
	Env map[string]string `json:"env"`
	// All tasks available for this project.
	// Experimental.
	Tasks map[string]TaskSpecIface `json:"tasks"`
}

func (t *TasksManifest) GetEnv() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TasksManifest) GetTasks() map[string]TaskSpecIface {
	var returns map[string]TaskSpecIface
	_jsii_.Get(
		t,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TaskSpecIface)(nil)).Elem(): reflect.TypeOf((*TaskSpec)(nil)).Elem(),
		},
	)
	return returns
}


