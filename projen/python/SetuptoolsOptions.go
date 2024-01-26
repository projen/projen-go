package python


// Experimental.
type SetuptoolsOptions struct {
	// Author's e-mail.
	// Default: $GIT_USER_EMAIL.
	//
	// Experimental.
	AuthorEmail *string `field:"required" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Default: $GIT_USER_NAME.
	//
	// Experimental.
	AuthorName *string `field:"required" json:"authorName" yaml:"authorName"`
	// Version of the package.
	// Default: "0.1.0"
	//
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Package name.
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *PoetryPyprojectOptionsWithoutDeps `field:"optional" json:"poetryOptions" yaml:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `field:"optional" json:"setupConfig" yaml:"setupConfig"`
	// Path to the python executable to use.
	// Default: "python".
	//
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
}

