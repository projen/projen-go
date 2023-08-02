package workflows


// Supported tools.
// Experimental.
type Tools struct {
	// Setup .NET Core.
	// Default: - not installed.
	//
	// Experimental.
	Dotnet *ToolRequirement `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Setup golang.
	// Default: - not installed.
	//
	// Experimental.
	Go *ToolRequirement `field:"optional" json:"go" yaml:"go"`
	// Setup java (temurin distribution).
	// Default: - not installed.
	//
	// Experimental.
	Java *ToolRequirement `field:"optional" json:"java" yaml:"java"`
	// Setup node.js.
	// Default: - not installed.
	//
	// Experimental.
	Node *ToolRequirement `field:"optional" json:"node" yaml:"node"`
	// Setup python.
	// Default: - not installed.
	//
	// Experimental.
	Python *ToolRequirement `field:"optional" json:"python" yaml:"python"`
}

