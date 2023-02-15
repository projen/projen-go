package python


// Experimental.
type RequirementsFileOptions struct {
	// Provide a list of packages that can be dynamically updated.
	// Experimental.
	PackageProvider IPackageProvider `field:"optional" json:"packageProvider" yaml:"packageProvider"`
}

