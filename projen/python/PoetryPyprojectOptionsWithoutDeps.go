package python


// Poetry-specific options.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyprojectOptionsWithoutDeps struct {
	// The authors of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Authors *[]*string `field:"optional" json:"authors" yaml:"authors"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// A short description of the package (required).
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A URL to the documentation of the project.
	// Experimental.
	Documentation *string `field:"optional" json:"documentation" yaml:"documentation"`
	// A list of patterns that will be excluded in the final package.
	//
	// If a VCS is being used for a package, the exclude field will be seeded with
	// the VCS’ ignore settings (.gitignore for git for example).
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Package extras.
	// Experimental.
	Extras *map[string]*[]*string `field:"optional" json:"extras" yaml:"extras"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// A list of patterns that will be included in the final package.
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// A list of keywords (max: 5) that the package is related to.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License of this package as an SPDX identifier.
	//
	// If the project is proprietary and does not use a specific license, you
	// can set this value as "Proprietary".
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// the maintainers of the package.
	//
	// Must be in the form "name <email>".
	// Experimental.
	Maintainers *[]*string `field:"optional" json:"maintainers" yaml:"maintainers"`
	// Name of the package (required).
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Package mode (optional).
	//
	// Example:
	//   false
	//
	// See: https://python-poetry.org/docs/pyproject/#package-mode
	//
	// Default: true.
	//
	// Experimental.
	PackageMode *bool `field:"optional" json:"packageMode" yaml:"packageMode"`
	// A list of packages and modules to include in the final distribution.
	// Experimental.
	Packages *[]interface{} `field:"optional" json:"packages" yaml:"packages"`
	// Plugins.
	//
	// Must be specified as a table.
	// See: https://toml.io/en/v1.0.0#table
	//
	// Experimental.
	Plugins interface{} `field:"optional" json:"plugins" yaml:"plugins"`
	// The name of the readme file of the package.
	// Experimental.
	Readme *string `field:"optional" json:"readme" yaml:"readme"`
	// A URL to the repository of the project.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The scripts or executables that will be installed when installing the package.
	// Experimental.
	Scripts *map[string]interface{} `field:"optional" json:"scripts" yaml:"scripts"`
	// Source registries from which packages are retrieved.
	// Experimental.
	Source *[]interface{} `field:"optional" json:"source" yaml:"source"`
	// Project custom URLs, in addition to homepage, repository and documentation.
	//
	// E.g. "Bug Tracker"
	// Experimental.
	Urls *map[string]*string `field:"optional" json:"urls" yaml:"urls"`
	// Version of the package (required).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

