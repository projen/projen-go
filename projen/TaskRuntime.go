// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// The runtime component of the tasks engine.
// Experimental.
type TaskRuntime interface {
	// The contents of tasks.json.
	// Experimental.
	Manifest() *TasksManifest
	// The tasks in this project.
	// Experimental.
	Tasks() *[]*TaskSpec
	// The root directory of the project and the cwd for executing tasks.
	// Experimental.
	Workdir() *string
	// Runs the task.
	// Experimental.
	RunTask(name *string, parents *[]*string, args *[]interface{})
	// Find a task by name, or `undefined` if not found.
	// Experimental.
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

	if err := validateNewTaskRuntimeParameters(workdir); err != nil {
		panic(err)
	}
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

func TaskRuntime_MANIFEST_FILE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.TaskRuntime",
		"MANIFEST_FILE",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TaskRuntime) RunTask(name *string, parents *[]*string, args *[]interface{}) {
	if err := t.validateRunTaskParameters(name, args); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"runTask",
		[]interface{}{name, parents, args},
	)
}

func (t *jsiiProxy_TaskRuntime) TryFindTask(name *string) *TaskSpec {
	if err := t.validateTryFindTaskParameters(name); err != nil {
		panic(err)
	}
	var returns *TaskSpec

	_jsii_.Invoke(
		t,
		"tryFindTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

