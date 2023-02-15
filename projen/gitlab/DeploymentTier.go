package gitlab


// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
// Experimental.
type DeploymentTier string

const (
	// Experimental.
	DeploymentTier_DEVELOPMENT DeploymentTier = "DEVELOPMENT"
	// Experimental.
	DeploymentTier_OTHER DeploymentTier = "OTHER"
	// Experimental.
	DeploymentTier_PRODUCTION DeploymentTier = "PRODUCTION"
	// Experimental.
	DeploymentTier_STAGING DeploymentTier = "STAGING"
	// Experimental.
	DeploymentTier_TESTING DeploymentTier = "TESTING"
)

