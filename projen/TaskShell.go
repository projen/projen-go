package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// The shell used to run a task's commands - including its `condition` and any `$(...)` environment-variable evaluation.
//
// Choose one of the built-in shells, or provide an explicit shell invocation:
//
// - {@link TaskShell.projen} (the default) - the built-in cross-platform shell.
// - {@link TaskShell.system} - the operating system's native shell.
// - {@link TaskShell.bash} / {@link TaskShell.sh} - common POSIX shells.
// - {@link TaskShell.command} - an arbitrary shell invocation given as an
//   explicit argument list, e.g. `TaskShell.command(["npx", "-c"])`.
//
// A `shell` can be set at the project (tasks), task and step level, and the
// nearest declared level wins (it is a scalar override, not merged).
// Experimental.
type TaskShell interface {
}

// The jsii proxy struct for TaskShell
type jsiiProxy_TaskShell struct {
	_ byte // padding
}

// Runs commands through `bash -c`.
// Experimental.
func TaskShell_Bash() TaskShell {
	_init_.Initialize()

	var returns TaskShell

	_jsii_.StaticInvoke(
		"projen.TaskShell",
		"bash",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An arbitrary shell invocation, given as an explicit argument list.
//
// The task
// command is appended as the final argument, so the invocation must accept
// the command as its last argument (e.g. `bash -c`, `sh -c`, `npx -c`,
// `yarn exec`).
//
// Example:
//   TaskShell.command(["npx", "-c"]);
//
// Experimental.
func TaskShell_Command(command *[]*string) TaskShell {
	_init_.Initialize()

	if err := validateTaskShell_CommandParameters(command); err != nil {
		panic(err)
	}
	var returns TaskShell

	_jsii_.StaticInvoke(
		"projen.TaskShell",
		"command",
		[]interface{}{command},
		&returns,
	)

	return returns
}

// The built-in cross-platform shell (the default).
//
// It interprets POSIX-style
// task syntax (`mkdir -p`, `&&`, `$VAR`, ...) identically on every platform,
// including Windows.
// Experimental.
func TaskShell_Projen() TaskShell {
	_init_.Initialize()

	var returns TaskShell

	_jsii_.StaticInvoke(
		"projen.TaskShell",
		"projen",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Runs commands through `sh -c`.
// Experimental.
func TaskShell_Sh() TaskShell {
	_init_.Initialize()

	var returns TaskShell

	_jsii_.StaticInvoke(
		"projen.TaskShell",
		"sh",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The operating system's native shell (`/bin/sh` on POSIX, `cmd.exe` on Windows).
//
// Use this to opt out of the cross-platform shell and run commands through
// whatever shell the host provides.
// Experimental.
func TaskShell_System() TaskShell {
	_init_.Initialize()

	var returns TaskShell

	_jsii_.StaticInvoke(
		"projen.TaskShell",
		"system",
		nil, // no parameters
		&returns,
	)

	return returns
}

