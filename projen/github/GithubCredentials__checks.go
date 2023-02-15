//go:build !no_runtime_type_checking

package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateGithubCredentials_FromAppParameters(options *GithubCredentialsAppOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGithubCredentials_FromPersonalAccessTokenParameters(options *GithubCredentialsPersonalAccessTokenOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

