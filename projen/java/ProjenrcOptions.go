package java


// Options for `Projenrc`.
// Experimental.
type ProjenrcOptions struct {
	// The name of the Java class which contains the `main()` method for projen.
	// Experimental.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The projen version to use.
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Defines projenrc under the test scope instead of the main scope, which is reserved to the app.
	//
	// This means that projenrc will be under
	// `src/test/java/projenrc.java` and projen will be defined as a test
	// dependency. This enforces that application code does not take a dependency
	// on projen code.
	//
	// If this is disabled, projenrc should be under
	// `src/main/java/projenrc.java`.
	// Experimental.
	TestScope *bool `field:"optional" json:"testScope" yaml:"testScope"`
}

