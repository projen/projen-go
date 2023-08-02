package projen


// Props for DockerCompose.
// Experimental.
type DockerComposeProps struct {
	// A name to add to the docker-compose.yml filename.
	//
	// Example:
	//   'myname' yields 'docker-compose.myname.yml'
	//
	// Default: - no name is added.
	//
	// Experimental.
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	// Docker Compose schema version do be used.
	// Default: 3.3
	//
	// Experimental.
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// Service descriptions.
	// Experimental.
	Services *map[string]*DockerComposeServiceDescription `field:"optional" json:"services" yaml:"services"`
}

