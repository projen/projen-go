package projen


// Props for DockerCompose.
// Experimental.
type DockerComposeProps struct {
	// A name to add to the docker-compose.yml filename.
	//
	// Example:
	//   'myname' yields 'docker-compose.myname.yml'
	//
	// Experimental.
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	// Docker Compose schema version do be used.
	// Experimental.
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// Service descriptions.
	// Experimental.
	Services *map[string]*DockerComposeServiceDescription `field:"optional" json:"services" yaml:"services"`
}

