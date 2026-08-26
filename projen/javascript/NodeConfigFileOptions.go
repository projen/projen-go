package javascript


// Options for `NodeConfigFile`.
// See: https://nodejs.org/api/cli.html#configuration-via-nodeconfig
//
// Experimental.
type NodeConfigFileOptions struct {
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
	// The path of the generated Node.js configuration file.
	// Default: "node.config.json"
	//
	// Experimental.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
}

