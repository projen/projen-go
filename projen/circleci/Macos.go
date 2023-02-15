package circleci


// CircleCI supports running jobs on macOS, to allow you to build, test, and deploy apps for macOS, iOS, tvOS and watchOS.
//
// To run a job in a macOS virtual machine,
// you must add the macos key to the top-level configuration for the job and specify
// the version of Xcode you would like to use.
// See: https://circleci.com/docs/2.0/configuration-reference/#macos
//
// Experimental.
type Macos struct {
	// The version of Xcode that is installed on the virtual machine.
	// Experimental.
	Xcode *string `field:"required" json:"xcode" yaml:"xcode"`
}

