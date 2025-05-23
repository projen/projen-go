package release

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Used to manage release strategies.
//
// This includes release
// and release artifact automation.
// Experimental.
type ReleaseTrigger interface {
	// Project-level changelog file path.
	// Experimental.
	ChangelogPath() *string
	// Override git-push command used when releasing manually.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand() *string
	// Whether or not this is a continuous release.
	// Experimental.
	IsContinuous() *bool
	// Whether or not this is a release trigger with a manual task run in a working copy.
	//
	// If the `ReleaseTrigger` is a GitHub-only manual task, this will return `false`.
	// Experimental.
	IsManual() *bool
	// Paths for which pushes will trigger a release when `isContinuous` is `true`.
	// Experimental.
	Paths() *[]*string
	// Cron schedule for releases.
	//
	// Only defined if this is a scheduled release.
	//
	// Example:
	//   '0 17 * * *' - every day at 5 pm
	//
	// Experimental.
	Schedule() *string
}

// The jsii proxy struct for ReleaseTrigger
type jsiiProxy_ReleaseTrigger struct {
	_ byte // padding
}

func (j *jsiiProxy_ReleaseTrigger) ChangelogPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changelogPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) GitPushCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPushCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) IsContinuous() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isContinuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) IsManual() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) Paths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}


// Creates a continuous release trigger.
//
// Automated releases will occur on every commit.
// Experimental.
func ReleaseTrigger_Continuous(options *ContinuousReleaseOptions) ReleaseTrigger {
	_init_.Initialize()

	if err := validateReleaseTrigger_ContinuousParameters(options); err != nil {
		panic(err)
	}
	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"continuous",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a manual release trigger.
//
// Use this option if you want totally manual releases.
//
// This will give you a release task that, in addition to the normal
// release activities will trigger a `publish:git` task. This task will
// handle project-level changelog management, release tagging, and pushing
// these artifacts to origin.
//
// The command used for pushing can be customised by specifying
// `gitPushCommand`. Set to an empty string to disable pushing entirely.
//
// Simply run `yarn release` to trigger a manual release.
// Experimental.
func ReleaseTrigger_Manual(options *ManualReleaseOptions) ReleaseTrigger {
	_init_.Initialize()

	if err := validateReleaseTrigger_ManualParameters(options); err != nil {
		panic(err)
	}
	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"manual",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a scheduled release trigger.
//
// Automated releases will occur based on the provided cron schedule.
// Experimental.
func ReleaseTrigger_Scheduled(options *ScheduledReleaseOptions) ReleaseTrigger {
	_init_.Initialize()

	if err := validateReleaseTrigger_ScheduledParameters(options); err != nil {
		panic(err)
	}
	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"scheduled",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// The release can only be triggered using the GitHub UI.
// Experimental.
func ReleaseTrigger_WorkflowDispatch() ReleaseTrigger {
	_init_.Initialize()

	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"workflowDispatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

