package gitlab


// Exit code that are not considered failure.
//
// The job fails for any other exit code.
// You can list which exit codes are not considered failures. The job fails for any other
// exit code.
// See: https://docs.gitlab.com/ee/ci/yaml/#allow_failure
//
// Experimental.
type AllowFailure struct {
	// Experimental.
	ExitCodes interface{} `field:"required" json:"exitCodes" yaml:"exitCodes"`
}

