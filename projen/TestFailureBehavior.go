package projen


// Experimental.
type TestFailureBehavior string

const (
	// Skip the current patch operation and continue with the next operation.
	// Experimental.
	TestFailureBehavior_SKIP TestFailureBehavior = "SKIP"
	// Fail the whole file synthesis.
	// Experimental.
	TestFailureBehavior_FAIL_SYNTHESIS TestFailureBehavior = "FAIL_SYNTHESIS"
)

