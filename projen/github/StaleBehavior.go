package github


// Stale behavior.
// Experimental.
type StaleBehavior struct {
	// The comment to add to the issue/PR when it's closed.
	// Default: "Closing this pull request as it hasn\'t seen activity for a while. Please add a comment
	//
	// Experimental.
	CloseMessage *string `field:"optional" json:"closeMessage" yaml:"closeMessage"`
	// Days until the issue/PR is closed after it is marked as "Stale".
	//
	// Set to -1 to disable.
	// Default: -.
	//
	// Experimental.
	DaysBeforeClose *float64 `field:"optional" json:"daysBeforeClose" yaml:"daysBeforeClose"`
	// How many days until the issue or pull request is marked as "Stale".
	//
	// Set to -1 to disable.
	// Default: -.
	//
	// Experimental.
	DaysBeforeStale *float64 `field:"optional" json:"daysBeforeStale" yaml:"daysBeforeStale"`
	// Determines if this behavior is enabled.
	//
	// Same as setting `daysBeforeStale` and `daysBeforeClose` to `-1`.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Label which exempt an issue/PR from becoming stale.
	//
	// Set to `[]` to disable.
	// Default: - ["backlog"].
	//
	// Experimental.
	ExemptLabels *[]*string `field:"optional" json:"exemptLabels" yaml:"exemptLabels"`
	// The label to apply to the issue/PR when it becomes stale.
	// Default: "stale".
	//
	// Experimental.
	StaleLabel *string `field:"optional" json:"staleLabel" yaml:"staleLabel"`
	// The comment to add to the issue/PR when it becomes stale.
	// Default: "This pull request is now marked as stale because hasn\'t seen activity for a while. Add a comment or it will be closed soon."
	//
	// Experimental.
	StaleMessage *string `field:"optional" json:"staleMessage" yaml:"staleMessage"`
}

