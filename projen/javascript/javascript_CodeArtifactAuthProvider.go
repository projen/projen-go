package javascript


// Options for authorizing requests to a AWS CodeArtifact npm repository.
// Experimental.
type CodeArtifactAuthProvider string

const (
	// Fixed credentials provided via Github secrets.
	// Experimental.
	CodeArtifactAuthProvider_ACCESS_AND_SECRET_KEY_PAIR CodeArtifactAuthProvider = "ACCESS_AND_SECRET_KEY_PAIR"
	// Ephemeral credentials provided via Github's OIDC integration with an IAM role.
	//
	// See:
	// https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc.html
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services
	// Experimental.
	CodeArtifactAuthProvider_GITHUB_OIDC CodeArtifactAuthProvider = "GITHUB_OIDC"
)

