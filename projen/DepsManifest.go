// CDK for software projects
package projen


// Experimental.
type DepsManifest struct {
	// All dependencies of this module.
	// Experimental.
	Dependencies *[]*Dependency `field:"required" json:"dependencies" yaml:"dependencies"`
}

