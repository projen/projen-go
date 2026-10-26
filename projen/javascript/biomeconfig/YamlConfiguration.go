package biomeconfig


// Options applied to Yaml files.
// Experimental.
type YamlConfiguration struct {
	// Formatter options.
	// Experimental.
	Formatter *YamlFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Experimental.
	Linter *YamlLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

