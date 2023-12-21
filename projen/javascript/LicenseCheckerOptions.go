package javascript


// Options to configure the license checker.
// Experimental.
type LicenseCheckerOptions struct {
	// List of SPDX license identifiers that are allowed to be used.
	//
	// For the license check to pass, all detected licenses MUST be in this list.
	// Only one of `allowedLicenses` and `prohibitedLicenses` can be provided and must not be empty.
	// Default: - no licenses are allowed.
	//
	// Experimental.
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// List of SPDX license identifiers that are prohibited to be used.
	//
	// For the license check to pass, no detected licenses can be in this list.
	// Only one of `allowedLicenses` and `prohibitedLicenses` can be provided and must not be empty.
	// Default: - no licenses are prohibited.
	//
	// Experimental.
	Deny *[]*string `field:"optional" json:"deny" yaml:"deny"`
	// Check development dependencies.
	// Default: false.
	//
	// Experimental.
	Development *bool `field:"optional" json:"development" yaml:"development"`
	// Check production dependencies.
	// Default: true.
	//
	// Experimental.
	Production *bool `field:"optional" json:"production" yaml:"production"`
	// The name of the task that is added to check licenses.
	// Default: "check-licenses".
	//
	// Experimental.
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
}

