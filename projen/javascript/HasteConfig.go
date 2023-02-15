package javascript


// Experimental.
type HasteConfig struct {
	// Experimental.
	ComputeSha1 *bool `field:"optional" json:"computeSha1" yaml:"computeSha1"`
	// Experimental.
	DefaultPlatform *string `field:"optional" json:"defaultPlatform" yaml:"defaultPlatform"`
	// Experimental.
	HasteImplModulePath *string `field:"optional" json:"hasteImplModulePath" yaml:"hasteImplModulePath"`
	// Experimental.
	Platforms *[]*string `field:"optional" json:"platforms" yaml:"platforms"`
	// Experimental.
	ThrowOnModuleCollision *bool `field:"optional" json:"throwOnModuleCollision" yaml:"throwOnModuleCollision"`
}

