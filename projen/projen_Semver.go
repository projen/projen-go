// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
type Semver interface {
	// Deprecated: This class will be removed in upcoming releases. if you wish to
	// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
	// `express@^2.1`.
	Mode() *string
	// Deprecated: This class will be removed in upcoming releases. if you wish to
	// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
	// `express@^2.1`.
	Spec() *string
	// Deprecated: This class will be removed in upcoming releases. if you wish to
	// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
	// `express@^2.1`.
	Version() *string
}

// The jsii proxy struct for Semver
type jsiiProxy_Semver struct {
	_ byte // padding
}

func (j *jsiiProxy_Semver) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semver) Spec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semver) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Accept any minor version.
//
// >= version
// < next major version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Caret(version *string) Semver {
	_init_.Initialize()

	if err := validateSemver_CaretParameters(version); err != nil {
		panic(err)
	}
	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"caret",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Latest version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Latest() Semver {
	_init_.Initialize()

	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"latest",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Of(spec *string) Semver {
	_init_.Initialize()

	if err := validateSemver_OfParameters(spec); err != nil {
		panic(err)
	}
	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"of",
		[]interface{}{spec},
		&returns,
	)

	return returns
}

// Accept only an exact version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Pinned(version *string) Semver {
	_init_.Initialize()

	if err := validateSemver_PinnedParameters(version); err != nil {
		panic(err)
	}
	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"pinned",
		[]interface{}{version},
		&returns,
	)

	return returns
}

// Accept patches.
//
// >= version
// < next minor version.
// Deprecated: This class will be removed in upcoming releases. if you wish to
// specify semver requirements in `deps`, `devDeps`, etc, specify them like so
// `express@^2.1`.
func Semver_Tilde(version *string) Semver {
	_init_.Initialize()

	if err := validateSemver_TildeParameters(version); err != nil {
		panic(err)
	}
	var returns Semver

	_jsii_.StaticInvoke(
		"projen.Semver",
		"tilde",
		[]interface{}{version},
		&returns,
	)

	return returns
}

