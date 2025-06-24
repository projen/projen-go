package biomeconfig


// The configuration that is contained inside the file `biome.json`.
// Experimental.
type BiomeConfiguration struct {
	// Specific configuration for assists.
	// Experimental.
	Assist *AssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Specific configuration for the Css language.
	// Experimental.
	Css *CssConfiguration `field:"optional" json:"css" yaml:"css"`
	// A list of paths to other JSON files, used to extends the current configuration.
	// Experimental.
	Extends *[]*string `field:"optional" json:"extends" yaml:"extends"`
	// The configuration of the filesystem.
	// Experimental.
	Files *FilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// The configuration of the formatter.
	// Experimental.
	Formatter *FormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Specific configuration for the GraphQL language.
	// Experimental.
	Graphql *GraphqlConfiguration `field:"optional" json:"graphql" yaml:"graphql"`
	// Specific configuration for the GraphQL language.
	// Experimental.
	Grit *GritConfiguration `field:"optional" json:"grit" yaml:"grit"`
	// Specific configuration for the HTML language.
	// Experimental.
	Html *HtmlConfiguration `field:"optional" json:"html" yaml:"html"`
	// Specific configuration for the JavaScript language.
	// Experimental.
	Javascript *JsConfiguration `field:"optional" json:"javascript" yaml:"javascript"`
	// Specific configuration for the Json language.
	// Experimental.
	Json *JsonConfiguration `field:"optional" json:"json" yaml:"json"`
	// The configuration for the linter.
	// Experimental.
	Linter *LinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// A list of granular patterns that should be applied only to a sub set of files.
	// Experimental.
	Overrides *[]*OverridePattern `field:"optional" json:"overrides" yaml:"overrides"`
	// List of plugins to load.
	// Experimental.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Indicates whether this configuration file is at the root of a Biome project.
	//
	// By default, this is `true`.
	// Experimental.
	Root *bool `field:"optional" json:"root" yaml:"root"`
	// A field for the [JSON schema](https://json-schema.org/) specification.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// The configuration of the VCS integration.
	// Experimental.
	Vcs *VcsConfiguration `field:"optional" json:"vcs" yaml:"vcs"`
}

