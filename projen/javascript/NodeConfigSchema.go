package javascript


// Experimental.
type NodeConfigSchema struct {
	// Experimental.
	NodeOptions *NodeConfigSchemaNodeOptions `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// Experimental.
	Permission *NodeConfigSchemaPermission `field:"optional" json:"permission" yaml:"permission"`
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Experimental.
	Test *NodeConfigSchemaTest `field:"optional" json:"test" yaml:"test"`
	// Experimental.
	Watch *NodeConfigSchemaWatch `field:"optional" json:"watch" yaml:"watch"`
}

