package biomeconfig


// The configuration contained in `biome.json`.
// Experimental.
type BiomeConfiguration struct {
	// The assist configuration.
	// Experimental.
	Assist *AssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Configuration specific to CSS.
	// Experimental.
	Css *CssConfiguration `field:"optional" json:"css" yaml:"css"`
	// A list of paths to other JSON files, used to extend the current configuration.
	// Experimental.
	Extends *[]*string `field:"optional" json:"extends" yaml:"extends"`
	// The file handling configuration.
	// Experimental.
	Files *FilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// The formatter configuration.
	// Experimental.
	Formatter *FormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Configuration specific to GraphQL.
	// Experimental.
	Graphql *GraphqlConfiguration `field:"optional" json:"graphql" yaml:"graphql"`
	// Configuration specific to GritQL.
	// Experimental.
	Grit *GritConfiguration `field:"optional" json:"grit" yaml:"grit"`
	// Configuration specific to HTML.
	// Experimental.
	Html *HtmlConfiguration `field:"optional" json:"html" yaml:"html"`
	// Configuration specific to JavaScript.
	// Experimental.
	Javascript *JsConfiguration `field:"optional" json:"javascript" yaml:"javascript"`
	// Configuration specific to JSON.
	// Experimental.
	Json *JsonConfiguration `field:"optional" json:"json" yaml:"json"`
	// The linter configuration.
	// Experimental.
	Linter *LinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// A list of granular patterns applied only to a subset of files.
	// Experimental.
	Overrides *[]*OverridePattern `field:"optional" json:"overrides" yaml:"overrides"`
	// List of plugins to load.
	// Experimental.
	Plugins *[]interface{} `field:"optional" json:"plugins" yaml:"plugins"`
	// Indicates whether this configuration file is at the root of a Biome project.
	//
	// By default, this is `true`.
	// Experimental.
	Root *bool `field:"optional" json:"root" yaml:"root"`
	// A field for the JSON schema specification: https://json-schema.org/.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// The version control integration configuration.
	// Experimental.
	Vcs *VcsConfiguration `field:"optional" json:"vcs" yaml:"vcs"`
}

