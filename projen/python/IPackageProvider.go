package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

// Experimental.
type IPackageProvider interface {
	// An array of packages (may be dynamically generated).
	// Experimental.
	Packages() *[]*projen.Dependency
}

// The jsii proxy for IPackageProvider
type jsiiProxy_IPackageProvider struct {
	_ byte // padding
}

func (j *jsiiProxy_IPackageProvider) Packages() *[]*projen.Dependency {
	var returns *[]*projen.Dependency
	_jsii_.Get(
		j,
		"packages",
		&returns,
	)
	return returns
}

