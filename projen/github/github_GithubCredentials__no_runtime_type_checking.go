//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateGithubCredentials_FromAppParameters(options *GithubCredentialsAppOptions) error {
	return nil
}

func validateGithubCredentials_FromPersonalAccessTokenParameters(options *GithubCredentialsPersonalAccessTokenOptions) error {
	return nil
}

