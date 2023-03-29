package python


// Fields to pass in the setup() function of setup.py.
// See: https://docs.python.org/3/distutils/setupscript.html
//
// Experimental.
type SetupPyOptions struct {
	// Escape hatch to allow any value.
	// Experimental.
	AdditionalOptions *map[string]interface{} `field:"optional" json:"additionalOptions" yaml:"additionalOptions"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// A short project description.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// The project license.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Name of the package.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of submodules to be packaged.
	// Experimental.
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
	// Manually specify package version.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

