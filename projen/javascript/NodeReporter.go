package javascript


// A single reporter/destination pair for the Node.js native test runner.
// See: https://nodejs.org/api/test.html#test-reporters
//
// Experimental.
type NodeReporter struct {
	// Where the reporter's output is written.
	// See: https://github.com/nodejs/node/blob/4215cc35e25c44f9f4fea5a4541afc862db7ef0a/test/parallel/test-runner-reporters.js#L46-L77
	//
	// Default: Destination.STDOUT
	//
	// Experimental.
	Destination Destination `field:"required" json:"destination" yaml:"destination"`
	// The name/kind of the reporter.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

