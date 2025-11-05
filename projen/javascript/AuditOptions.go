package javascript


// Options for security audit configuration.
// Experimental.
type AuditOptions struct {
	// Minimum vulnerability level to check for during audit.
	// Default: "high".
	//
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Only audit production dependencies.
	//
	// When false, both production and development dependencies are audited.
	// This is recommended as build dependencies can also contain security vulnerabilities.
	// Default: false.
	//
	// Experimental.
	ProdOnly *bool `field:"optional" json:"prodOnly" yaml:"prodOnly"`
	// When to run the audit task.
	//
	// - "build": Run during every build (default)
	// - "release": Only run during release workflow
	// - "manual": Create the task but don't run it automatically.
	// Default: "build".
	//
	// Experimental.
	RunOn *string `field:"optional" json:"runOn" yaml:"runOn"`
}

