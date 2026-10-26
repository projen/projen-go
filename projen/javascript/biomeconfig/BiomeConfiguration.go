package biomeconfig


// The configuration contained in `biome.json` or `biome.jsonc`.
// Experimental.
type BiomeConfiguration struct {
	// The assist configuration.
	// Experimental.
	Assist *AssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Configuration specific to CSS.
	// Experimental.
	Css *CssConfiguration `field:"optional" json:"css" yaml:"css"`
	// Extends the current configuration with settings from one or more other Biome configurations.
	//
	// Use `"//"` in a nested configuration to extend the root configuration at any nesting depth.
	// Use an array to extend configurations by relative path or installed package specifier.
	//
	// Biome merges configurations from left to right. Later configurations take precedence for
	// single-value options, list entries are combined, and the current file is applied last.
	// Experimental.
	Extends interface{} `field:"optional" json:"extends" yaml:"extends"`
	// Configures which files Biome can discover and process.
	// Experimental.
	Files *FilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// Configures formatting for selected files.
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
	// Configuration specific to Markdown.
	// Experimental.
	Markdown *MarkdownConfiguration `field:"optional" json:"markdown" yaml:"markdown"`
	// Configures settings for files selected by each override.
	//
	// When multiple overrides configure
	// the same single-value setting, the last matching override takes precedence.
	// Experimental.
	Overrides *[]*OverridePattern `field:"optional" json:"overrides" yaml:"overrides"`
	// A list of GritQL plugins to load.
	// Experimental.
	Plugins *[]interface{} `field:"optional" json:"plugins" yaml:"plugins"`
	// Marks whether this file is the root configuration for a Biome project.
	//
	// The root configuration
	// establishes the project's base settings, while `false` marks a configuration found inside the
	// project as nested rather than a second project root.
	//
	// Defaults to `true`. Biome implicitly treats `root` as `false` when a nested configuration uses
	// `"extends": "//"`. All other nested configurations must explicitly set `root` to `false`.
	// Default: true`. Biome implicitly treats `root` as `false` when a nested configuration uses
	//
	// Experimental.
	Root *bool `field:"optional" json:"root" yaml:"root"`
	// The JSON Schema used to validate the configuration and provide editor completion.
	//
	// Biome
	// emits a diagnostic when a `biomejs.dev` schema URL specifies a version that differs from the
	// running CLI version.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Configures how Biome recognizes Tailwind class strings.
	// Experimental.
	Tailwind *TailwindConfiguration `field:"optional" json:"tailwind" yaml:"tailwind"`
	// The version control integration configuration.
	// Experimental.
	Vcs *VcsConfiguration `field:"optional" json:"vcs" yaml:"vcs"`
	// Configuration specific to YAML.
	// Experimental.
	Yaml *YamlConfiguration `field:"optional" json:"yaml" yaml:"yaml"`
}

