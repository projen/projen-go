package projen


// Experimental.
type LicenseOptions struct {
	// License type (SPDX).
	// See: https://github.com/projen/projen/tree/main/license-text for list of supported licenses
	//
	// Experimental.
	Spdx *string `field:"required" json:"spdx" yaml:"spdx"`
	// Copyright owner.
	//
	// If the license text has $copyright_owner, this option must be specified.
	// Experimental.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// Period of license (e.g. "1998-2023").
	//
	// The string `$copyright_period` will be substituted with this string.
	// Experimental.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
}

