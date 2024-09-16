package github


// Options for `concurrency`.
// Experimental.
type ConcurrencyOptions struct {
	// When a workflow is triggered while another one (in the same group) is running, should GitHub cancel the running workflow?
	// Default: false.
	//
	// Experimental.
	CancelInProgress *bool `field:"optional" json:"cancelInProgress" yaml:"cancelInProgress"`
	// Concurrency group controls which workflow runs will share the same concurrency limit.
	//
	// For example, if you specify `${{ github.workflow }}-${{ github.ref }}`, workflow runs triggered
	// on the same branch cannot run concurrenty, but workflows runs triggered on different branches can.
	// See: https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/using-concurrency#example-concurrency-groups
	//
	// Default: - ${{ github.workflow }}
	//
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
}

