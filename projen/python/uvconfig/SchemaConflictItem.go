package uvconfig


// A single item in a conflicting set.
//
// Each item is a pair of an (optional) package and a corresponding extra or group name for that
// package.
// Experimental.
type SchemaConflictItem struct {
	// Experimental.
	Extra *string `field:"optional" json:"extra" yaml:"extra"`
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Experimental.
	Package *string `field:"optional" json:"package" yaml:"package"`
}

