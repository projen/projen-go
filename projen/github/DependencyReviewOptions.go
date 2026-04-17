package github

import (
	"github.com/projen/projen-go/projen"
)

// Options for the DependencyReview component.
// Experimental.
type DependencyReviewOptions struct {
	// GitHub Advisory Database IDs that can be skipped during detection.
	// Default: - no advisories are skipped.
	//
	// Experimental.
	AllowGhsas *[]*string `field:"optional" json:"allowGhsas" yaml:"allowGhsas"`
	// List of allowed SPDX license identifiers.
	// Default: - no license allow-list.
	//
	// Experimental.
	AllowLicenses *[]*string `field:"optional" json:"allowLicenses" yaml:"allowLicenses"`
	// Whether to post a comment summary on the PR.
	// Default: "always".
	//
	// Experimental.
	CommentSummaryInPr *string `field:"optional" json:"commentSummaryInPr" yaml:"commentSummaryInPr"`
	// Path to an external configuration file.
	// Default: - no external config.
	//
	// Experimental.
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	// Packages to block in a PR (in purl format).
	// Default: - no packages are denied.
	//
	// Experimental.
	DenyPackages *[]*string `field:"optional" json:"denyPackages" yaml:"denyPackages"`
	// Scopes of dependencies to fail on.
	// Default: - no scopes filter (action default is "runtime").
	//
	// Experimental.
	FailOnScopes *[]*string `field:"optional" json:"failOnScopes" yaml:"failOnScopes"`
	// The severity level at which the action will fail.
	// Default: - no minimum severity (action default is "low").
	//
	// Experimental.
	FailOnSeverity *string `field:"optional" json:"failOnSeverity" yaml:"failOnSeverity"`
	// Enable or disable the license check.
	// Default: true.
	//
	// Experimental.
	LicenseCheck *bool `field:"optional" json:"licenseCheck" yaml:"licenseCheck"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
	// Show OpenSSF Scorecard scores for dependencies.
	// Default: true.
	//
	// Experimental.
	ShowOpenSSFScorecard *bool `field:"optional" json:"showOpenSSFScorecard" yaml:"showOpenSSFScorecard"`
	// Enable or disable the vulnerability check.
	// Default: true.
	//
	// Experimental.
	VulnerabilityCheck *bool `field:"optional" json:"vulnerabilityCheck" yaml:"vulnerabilityCheck"`
	// When true, the action will only warn and not fail.
	// Default: false.
	//
	// Experimental.
	WarnOnly *bool `field:"optional" json:"warnOnly" yaml:"warnOnly"`
	// Score threshold for OpenSSF Scorecard warnings.
	// Default: 3.
	//
	// Experimental.
	WarnOnOpenSSFScorecardLevel *float64 `field:"optional" json:"warnOnOpenSSFScorecardLevel" yaml:"warnOnOpenSSFScorecardLevel"`
}

