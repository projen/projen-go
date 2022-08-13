package github


// Stale behavior.
// Experimental.
type StaleBehavior struct {
	// The comment to add to the issue/PR when it's closed.
	// Experimental.
	CloseMessage *string `field:"optional" json:"closeMessage" yaml:"closeMessage"`
	// Days until the issue/PR is closed after it is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeClose *float64 `field:"optional" json:"daysBeforeClose" yaml:"daysBeforeClose"`
	// How many days until the issue or pull request is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeStale *float64 `field:"optional" json:"daysBeforeStale" yaml:"daysBeforeStale"`
	// Determines if this behavior is enabled.
	//
	// Same as setting `daysBeforeStale` and `daysBeforeClose` to `-1`.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Label which exempt an issue/PR from becoming stale.
	//
	// Set to `[]` to disable.
	// Experimental.
	ExemptLabels *[]*string `field:"optional" json:"exemptLabels" yaml:"exemptLabels"`
	// The label to apply to the issue/PR when it becomes stale.
	// Experimental.
	StaleLabel *string `field:"optional" json:"staleLabel" yaml:"staleLabel"`
	// The comment to add to the issue/PR when it becomes stale.
	// Experimental.
	StaleMessage *string `field:"optional" json:"staleMessage" yaml:"staleMessage"`
}

