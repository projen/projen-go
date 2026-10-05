package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Describes where a pull request comes from.
//
// Used to select which pull requests a workflow acts on, e.g. which pull requests are auto-approved.
// Experimental.
type PullRequestSource interface {
}

// The jsii proxy struct for PullRequestSource
type jsiiProxy_PullRequestSource struct {
	_ byte // padding
}

// Pull requests from a head branch in this repository.
//
// Anyone who can push to the head branch controls the content of the pull request.
// Protect the matching branches with a ruleset that only allows the intended automation
// to create, update and delete them.
// Experimental.
func PullRequestSource_FromBranch(options *PullRequestSourceBranchOptions) PullRequestSource {
	_init_.Initialize()

	if err := validatePullRequestSource_FromBranchParameters(options); err != nil {
		panic(err)
	}
	var returns PullRequestSource

	_jsii_.StaticInvoke(
		"projen.github.PullRequestSource",
		"fromBranch",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Pull requests authored by any of the given users.
//
// Anyone with access to the credentials of these users can author a matching pull request.
// Experimental.
func PullRequestSource_FromUsers(options *PullRequestSourceUsersOptions) PullRequestSource {
	_init_.Initialize()

	if err := validatePullRequestSource_FromUsersParameters(options); err != nil {
		panic(err)
	}
	var returns PullRequestSource

	_jsii_.StaticInvoke(
		"projen.github.PullRequestSource",
		"fromUsers",
		[]interface{}{options},
		&returns,
	)

	return returns
}

