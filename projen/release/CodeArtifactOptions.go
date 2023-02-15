package release


// Options for publishing packages to AWS CodeArtifact.
// Experimental.
type CodeArtifactOptions struct {
	// GitHub secret which contains the AWS access key ID to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	AccessKeyIdSecret *string `field:"optional" json:"accessKeyIdSecret" yaml:"accessKeyIdSecret"`
	// Provider to use for authorizing requests to AWS CodeArtifact.
	// Experimental.
	AuthProvider CodeArtifactAuthProvider `field:"optional" json:"authProvider" yaml:"authProvider"`
	// ARN of AWS role to be assumed prior to get authorization token from AWS CodeArtifact This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	//
	// When using the `CodeArtifactAuthProvider.GITHUB_OIDC` auth provider, this value must be defined.
	// Experimental.
	RoleToAssume *string `field:"optional" json:"roleToAssume" yaml:"roleToAssume"`
	// GitHub secret which contains the AWS secret access key to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	SecretAccessKeySecret *string `field:"optional" json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
}

