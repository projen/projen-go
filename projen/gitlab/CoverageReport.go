package gitlab


// Code coverage report interface.
// Experimental.
type CoverageReport struct {
	// Experimental.
	CoverageFormat *string `field:"required" json:"coverageFormat" yaml:"coverageFormat"`
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

