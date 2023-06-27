package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Find commits that should be considered releasable to decide if a release is required.
// Experimental.
type ReleasableCommits interface {
	// Experimental.
	Cmd() *string
	// Experimental.
	SetCmd(val *string)
}

// The jsii proxy struct for ReleasableCommits
type jsiiProxy_ReleasableCommits struct {
	_ byte // padding
}

func (j *jsiiProxy_ReleasableCommits) Cmd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmd",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_ReleasableCommits)SetCmd(val *string) {
	if err := j.validateSetCmdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmd",
		val,
	)
}

// Release every commit.
//
// This will only not release if the most recent commit is tagged with the latest matching tag.
// Experimental.
func ReleasableCommits_EveryCommit(path *string) ReleasableCommits {
	_init_.Initialize()

	var returns ReleasableCommits

	_jsii_.StaticInvoke(
		"projen.ReleasableCommits",
		"everyCommit",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Use an arbitrary shell command to find releasable commits since the latest tag.
//
// A new release will be initiated, if the number of returned commits is greater than zero.
// Must return a newline separate list of commits that should considered releasable.
// `$LATEST_TAG` will be replaced with the actual latest tag for the given prefix.*
//
// Example:
//   "git log --oneline $LATEST_TAG..HEAD -- ."
//
// Experimental.
func ReleasableCommits_Exec(cmd *string) ReleasableCommits {
	_init_.Initialize()

	if err := validateReleasableCommits_ExecParameters(cmd); err != nil {
		panic(err)
	}
	var returns ReleasableCommits

	_jsii_.StaticInvoke(
		"projen.ReleasableCommits",
		"exec",
		[]interface{}{cmd},
		&returns,
	)

	return returns
}

// Release only features and fixes.
//
// Shorthand for `ReleasableCommits.onlyOfType(['feat', 'fix'])`.
// Experimental.
func ReleasableCommits_FeaturesAndFixes(path *string) ReleasableCommits {
	_init_.Initialize()

	var returns ReleasableCommits

	_jsii_.StaticInvoke(
		"projen.ReleasableCommits",
		"featuresAndFixes",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Limit commits by their conventional commit type.
//
// This will only release commit that match one of the provided types.
// Commits are required to follow the conventional commit spec and will be ignored otherwise.
// Experimental.
func ReleasableCommits_OfType(types *[]*string, path *string) ReleasableCommits {
	_init_.Initialize()

	if err := validateReleasableCommits_OfTypeParameters(types); err != nil {
		panic(err)
	}
	var returns ReleasableCommits

	_jsii_.StaticInvoke(
		"projen.ReleasableCommits",
		"ofType",
		[]interface{}{types, path},
		&returns,
	)

	return returns
}

