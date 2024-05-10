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
	// Default: - no version is provided.
	//
	// Deprecated: - The top level `version` field is obsolete per the Compose Specification.
	// {@link https://github.com/compose-spec/compose-spec/blob/master/spec.md#version-and-name-top-level-elements Compose Specification}
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// Service descriptions.
	// Experimental.
	Services *map[string]*DockerComposeServiceDescription `field:"optional" json:"services" yaml:"services"`
}

