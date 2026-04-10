package github


// Defines a cooldown period for dependency version updates.
// See: https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#cooldown-
//
// Experimental.
type DependabotCooldown struct {
	// Default cooldown period (in days) for all dependencies without specific semver rules.
	// Default: - no default cooldown.
	//
	// Experimental.
	DefaultDays *float64 `field:"optional" json:"defaultDays" yaml:"defaultDays"`
	// List of dependencies excluded from cooldown.
	//
	// Supports wildcards.
	// Takes precedence over `include`.
	// Default: - no exclusions.
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of dependencies to apply cooldown to.
	//
	// Supports wildcards.
	// Default: - all dependencies.
	//
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// Cooldown period (in days) for major version updates.
	// Default: - uses defaultDays.
	//
	// Experimental.
	SemverMajorDays *float64 `field:"optional" json:"semverMajorDays" yaml:"semverMajorDays"`
	// Cooldown period (in days) for minor version updates.
	// Default: - uses defaultDays.
	//
	// Experimental.
	SemverMinorDays *float64 `field:"optional" json:"semverMinorDays" yaml:"semverMinorDays"`
	// Cooldown period (in days) for patch version updates.
	// Default: - uses defaultDays.
	//
	// Experimental.
	SemverPatchDays *float64 `field:"optional" json:"semverPatchDays" yaml:"semverPatchDays"`
}

