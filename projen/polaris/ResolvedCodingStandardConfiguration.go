package polaris


// Experimental.
type ResolvedCodingStandardConfiguration struct {
	// Name of this code compliance configuration.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// List of deviations for this standard.
	// Experimental.
	Deviations *[]*CodingStandardDeviation `field:"optional" json:"deviations" yaml:"deviations"`
	// Version of this code compliance configuration.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

