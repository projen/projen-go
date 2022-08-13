package awscdk


// Which approval is required when deploying CDK apps.
// Experimental.
type ApprovalLevel string

const (
	// Approval is never required.
	// Experimental.
	ApprovalLevel_NEVER ApprovalLevel = "NEVER"
	// Requires approval on any IAM or security-group-related change.
	// Experimental.
	ApprovalLevel_ANY_CHANGE ApprovalLevel = "ANY_CHANGE"
	// Requires approval when IAM statements or traffic rules are added;
	//
	// removals don't require approval.
	// Experimental.
	ApprovalLevel_BROADENING ApprovalLevel = "BROADENING"
)

