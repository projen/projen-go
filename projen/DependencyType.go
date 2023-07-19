package projen


// Type of dependency.
// Experimental.
type DependencyType string

const (
	// The dependency is required for the program/library during runtime.
	// Experimental.
	DependencyType_RUNTIME DependencyType = "RUNTIME"
	// The dependency is required at runtime but expected to be installed by the consumer.
	// Experimental.
	DependencyType_PEER DependencyType = "PEER"
	// The dependency is bundled and shipped with the module, so consumers are not required to install it.
	// Experimental.
	DependencyType_BUNDLED DependencyType = "BUNDLED"
	// The dependency is required to run the `build` task.
	// Experimental.
	DependencyType_BUILD DependencyType = "BUILD"
	// The dependency is required to run the `test` task.
	// Experimental.
	DependencyType_TEST DependencyType = "TEST"
	// The dependency is required for development (e.g. IDE plugins).
	// Experimental.
	DependencyType_DEVENV DependencyType = "DEVENV"
	// Transient dependency that needs to be overwritten.
	//
	// Available for Node packages.
	// Experimental.
	DependencyType_OVERRIDE DependencyType = "OVERRIDE"
	// An optional dependency that may be used at runtime if available, but is not required.
	//
	// It is expected to be installed by the consumer.
	// Experimental.
	DependencyType_OPTIONAL DependencyType = "OPTIONAL"
)

