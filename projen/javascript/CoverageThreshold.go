package javascript


// Experimental.
type CoverageThreshold struct {
	// Experimental.
	Branches *float64 `field:"optional" json:"branches" yaml:"branches"`
	// Experimental.
	Functions *float64 `field:"optional" json:"functions" yaml:"functions"`
	// Experimental.
	Lines *float64 `field:"optional" json:"lines" yaml:"lines"`
	// Experimental.
	Statements *float64 `field:"optional" json:"statements" yaml:"statements"`
}

