package biomeconfig


// Experimental.
type OverridePattern struct {
	// Specific configuration for the Json language.
	// Experimental.
	Assist *OverrideAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Specific configuration for the CSS language.
	// Experimental.
	Css *CssConfiguration `field:"optional" json:"css" yaml:"css"`
	// Specific configuration for the filesystem.
	// Experimental.
	Files *OverrideFilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// Specific configuration for the Json language.
	// Experimental.
	Formatter *OverrideFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Specific configuration for the Graphql language.
	// Experimental.
	Graphql *GraphqlConfiguration `field:"optional" json:"graphql" yaml:"graphql"`
	// Specific configuration for the GritQL language.
	// Experimental.
	Grit *GritConfiguration `field:"optional" json:"grit" yaml:"grit"`
	// Specific configuration for the GritQL language.
	// Experimental.
	Html *HtmlConfiguration `field:"optional" json:"html" yaml:"html"`
	// A list of glob patterns.
	//
	// Biome will include files/folders that will
	// match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// Specific configuration for the JavaScript language.
	// Experimental.
	Javascript *JsConfiguration `field:"optional" json:"javascript" yaml:"javascript"`
	// Specific configuration for the Json language.
	// Experimental.
	Json *JsonConfiguration `field:"optional" json:"json" yaml:"json"`
	// Specific configuration for the Json language.
	// Experimental.
	Linter *OverrideLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// Specific configuration for additional plugins.
	// Experimental.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
}

