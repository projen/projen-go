package biomeconfig


// Experimental.
type Rules struct {
	// Experimental.
	A11Y interface{} `field:"optional" json:"a11Y" yaml:"a11Y"`
	// Experimental.
	Complexity interface{} `field:"optional" json:"complexity" yaml:"complexity"`
	// Experimental.
	Correctness interface{} `field:"optional" json:"correctness" yaml:"correctness"`
	// Experimental.
	Nursery interface{} `field:"optional" json:"nursery" yaml:"nursery"`
	// Experimental.
	Performance interface{} `field:"optional" json:"performance" yaml:"performance"`
	// It enables the lint rules recommended by Biome.
	//
	// `true` by default.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Experimental.
	Security interface{} `field:"optional" json:"security" yaml:"security"`
	// Experimental.
	Style interface{} `field:"optional" json:"style" yaml:"style"`
	// Experimental.
	Suspicious interface{} `field:"optional" json:"suspicious" yaml:"suspicious"`
}

