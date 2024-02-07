package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// A task that can be performed on the project.
//
// Modeled as a series of shell
// commands and subtasks.
// Experimental.
type Task interface {
	// A command to execute which determines if the task should be skipped.
	//
	// If it
	// returns a zero exit code, the task will not be executed.
	// Experimental.
	Condition() *string
	// Returns the working directory for this task.
	//
	// Sets the working directory for this task.
	// Experimental.
	Cwd() *string
	// Experimental.
	SetCwd(val *string)
	// Returns the description of this task.
	//
	// Sets the description of this task.
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Returns all environment variables in the task level.
	// Experimental.
	EnvVars() *map[string]*string
	// Task name.
	// Experimental.
	Name() *string
	// Returns an immutable copy of all the step specifications of the task.
	// Experimental.
	Steps() *[]*TaskStep
	// Add a command to execute which determines if the task should be skipped.
	//
	// If a condition already exists, the new condition will be appended with ` && ` delimiter.
	// See: {@link Task.condition }
	//
	// Experimental.
	AddCondition(condition ...*string)
	// Execute a builtin task.
	//
	// Builtin tasks are programs bundled as part of projen itself and used as
	// helpers for various components.
	//
	// In the future we should support built-in tasks from external modules.
	// Experimental.
	Builtin(name *string)
	// Adds an environment variable to this task.
	// Experimental.
	Env(name *string, value *string)
	// Executes a shell command.
	// Experimental.
	Exec(command *string, options *TaskStepOptions)
	// Forbid additional changes to this task.
	// Experimental.
	Lock()
	// Adds a command at the beginning of the task.
	// Deprecated: use `prependExec()`.
	Prepend(shell *string, options *TaskStepOptions)
	// Adds a command at the beginning of the task.
	// Experimental.
	PrependExec(shell *string, options *TaskStepOptions)
	// Says something at the beginning of the task.
	// Experimental.
	PrependSay(message *string, options *TaskStepOptions)
	// Adds a spawn instruction at the beginning of the task.
	// Experimental.
	PrependSpawn(subtask Task, options *TaskStepOptions)
	// Experimental.
	RemoveStep(index *float64)
	// Reset the task so it no longer has any commands.
	// Experimental.
	Reset(command *string, options *TaskStepOptions)
	// Say something.
	// Experimental.
	Say(message *string, options *TaskStepOptions)
	// Spawns a sub-task.
	// Experimental.
	Spawn(subtask Task, options *TaskStepOptions)
	// Experimental.
	UpdateStep(index *float64, step *TaskStep)
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

func (j *jsiiProxy_Task) Cwd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cwd",
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

func (j *jsiiProxy_Task) EnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVars",
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

	if err := validateNewTaskParameters(name, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_Task)SetCwd(val *string) {
	_jsii_.Set(
		j,
		"cwd",
		val,
	)
}

func (j *jsiiProxy_Task)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (t *jsiiProxy_Task) AddCondition(condition ...*string) {
	args := []interface{}{}
	for _, a := range condition {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addCondition",
		args,
	)
}

func (t *jsiiProxy_Task) Builtin(name *string) {
	if err := t.validateBuiltinParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"builtin",
		[]interface{}{name},
	)
}

func (t *jsiiProxy_Task) Env(name *string, value *string) {
	if err := t.validateEnvParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"env",
		[]interface{}{name, value},
	)
}

func (t *jsiiProxy_Task) Exec(command *string, options *TaskStepOptions) {
	if err := t.validateExecParameters(command, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"exec",
		[]interface{}{command, options},
	)
}

func (t *jsiiProxy_Task) Lock() {
	_jsii_.InvokeVoid(
		t,
		"lock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) Prepend(shell *string, options *TaskStepOptions) {
	if err := t.validatePrependParameters(shell, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"prepend",
		[]interface{}{shell, options},
	)
}

func (t *jsiiProxy_Task) PrependExec(shell *string, options *TaskStepOptions) {
	if err := t.validatePrependExecParameters(shell, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"prependExec",
		[]interface{}{shell, options},
	)
}

func (t *jsiiProxy_Task) PrependSay(message *string, options *TaskStepOptions) {
	if err := t.validatePrependSayParameters(message, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"prependSay",
		[]interface{}{message, options},
	)
}

func (t *jsiiProxy_Task) PrependSpawn(subtask Task, options *TaskStepOptions) {
	if err := t.validatePrependSpawnParameters(subtask, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"prependSpawn",
		[]interface{}{subtask, options},
	)
}

func (t *jsiiProxy_Task) RemoveStep(index *float64) {
	if err := t.validateRemoveStepParameters(index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"removeStep",
		[]interface{}{index},
	)
}

func (t *jsiiProxy_Task) Reset(command *string, options *TaskStepOptions) {
	if err := t.validateResetParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"reset",
		[]interface{}{command, options},
	)
}

func (t *jsiiProxy_Task) Say(message *string, options *TaskStepOptions) {
	if err := t.validateSayParameters(message, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"say",
		[]interface{}{message, options},
	)
}

func (t *jsiiProxy_Task) Spawn(subtask Task, options *TaskStepOptions) {
	if err := t.validateSpawnParameters(subtask, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"spawn",
		[]interface{}{subtask, options},
	)
}

func (t *jsiiProxy_Task) UpdateStep(index *float64, step *TaskStep) {
	if err := t.validateUpdateStepParameters(index, step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"updateStep",
		[]interface{}{index, step},
	)
}

