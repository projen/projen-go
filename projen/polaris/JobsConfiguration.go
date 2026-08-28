package polaris


// Experimental.
type JobsConfiguration struct {
	// If true, the number of analysis workers to run in parallel is based on the amount of memory and number of logical processors in the machine.
	//
	// This is the default for a non-Flexnet license. This key is mutually exclusive with the "count" and "max" keys.
	// Experimental.
	Auto *bool `field:"optional" json:"auto" yaml:"auto"`
	// Number of analysis workers to run in parallel.
	//
	// This key is mutually exclusive with the "auto" and "max" keys.
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Maximum number of analysis worker to run in parallel, subject to limits on the amount of memory and number of logical processors in the machine.
	//
	// A value of 8 is the default for a Flexnet license. This key is mutually exclusive with the "auto" and "count" keys.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Allows the number of analysis workers to exceed the recommended value.
	//
	// This key may only be used with the "count" key.
	// Experimental.
	OverrideWorkerLimit *bool `field:"optional" json:"overrideWorkerLimit" yaml:"overrideWorkerLimit"`
}

