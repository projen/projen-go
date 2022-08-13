package workflows


// Supported tools.
// Experimental.
type Tools struct {
	// Setup .NET Core.
	// Experimental.
	Dotnet *ToolRequirement `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Setup golang.
	// Experimental.
	Go *ToolRequirement `field:"optional" json:"go" yaml:"go"`
	// Setup java (temurin distribution).
	// Experimental.
	Java *ToolRequirement `field:"optional" json:"java" yaml:"java"`
	// Setup node.js.
	// Experimental.
	Node *ToolRequirement `field:"optional" json:"node" yaml:"node"`
	// Setup python.
	// Experimental.
	Python *ToolRequirement `field:"optional" json:"python" yaml:"python"`
}

