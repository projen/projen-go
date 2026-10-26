package biomeconfig


// Experimental.
type OverridePattern struct {
	// Assist settings for matched files.
	// Experimental.
	Assist *OverrideAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// CSS-specific settings for matched files.
	// Experimental.
	Css *CssConfiguration `field:"optional" json:"css" yaml:"css"`
	// File-handling settings for matched files.
	// Experimental.
	Files *OverrideFilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// Formatter settings for matched files.
	// Experimental.
	Formatter *OverrideFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// GraphQL-specific settings for matched files.
	// Experimental.
	Graphql *GraphqlConfiguration `field:"optional" json:"graphql" yaml:"graphql"`
	// GritQL-specific settings for matched files.
	// Experimental.
	Grit *GritConfiguration `field:"optional" json:"grit" yaml:"grit"`
	// HTML-specific settings for matched files.
	// Experimental.
	Html *HtmlConfiguration `field:"optional" json:"html" yaml:"html"`
	// A list of glob patterns selecting the files to which this override applies.
	//
	// If omitted, the
	// override applies to every file. An empty list applies to no files.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// JavaScript-specific settings for matched files.
	// Experimental.
	Javascript *JsConfiguration `field:"optional" json:"javascript" yaml:"javascript"`
	// JSON-specific settings for matched files.
	// Experimental.
	Json *JsonConfiguration `field:"optional" json:"json" yaml:"json"`
	// Linter settings for matched files.
	// Experimental.
	Linter *OverrideLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// Specific configuration for the Markdown language.
	// Experimental.
	Markdown *MarkdownConfiguration `field:"optional" json:"markdown" yaml:"markdown"`
	// Additional plugins for matched files.
	// Experimental.
	Plugins *[]interface{} `field:"optional" json:"plugins" yaml:"plugins"`
	// Specific configuration for the YAML language.
	// Experimental.
	Yaml *YamlConfiguration `field:"optional" json:"yaml" yaml:"yaml"`
}

