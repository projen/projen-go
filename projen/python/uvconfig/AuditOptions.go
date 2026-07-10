package uvconfig


// Experimental.
type AuditOptions struct {
	// A list of vulnerability IDs to ignore during auditing.
	//
	// Vulnerabilities matching any of the provided IDs (including aliases) will be excluded from
	// the audit results.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// A list of vulnerability IDs to ignore during auditing, but only while no fix is available.
	//
	// Vulnerabilities matching any of the provided IDs (including aliases) will be excluded from
	// the audit results as long as they have no known fix versions. Once a fix version becomes
	// available, the vulnerability will be reported again.
	// Experimental.
	IgnoreUntilFixed *[]*string `field:"optional" json:"ignoreUntilFixed" yaml:"ignoreUntilFixed"`
}

