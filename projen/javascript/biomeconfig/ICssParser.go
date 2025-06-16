package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the CSS parser behaves.
// Experimental.
type ICssParser interface {
	// Allow comments to appear on incorrect lines in `.css` files.
	// Experimental.
	AllowWrongLineComments() *bool
	// Experimental.
	SetAllowWrongLineComments(a *bool)
	// Enables parsing of CSS Modules specific features.
	// Experimental.
	CssModules() *bool
	// Experimental.
	SetCssModules(c *bool)
}

// The jsii proxy for ICssParser
type jsiiProxy_ICssParser struct {
	_ byte // padding
}

func (j *jsiiProxy_ICssParser) AllowWrongLineComments() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowWrongLineComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssParser)SetAllowWrongLineComments(val *bool) {
	_jsii_.Set(
		j,
		"allowWrongLineComments",
		val,
	)
}

func (j *jsiiProxy_ICssParser) CssModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cssModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssParser)SetCssModules(val *bool) {
	_jsii_.Set(
		j,
		"cssModules",
		val,
	)
}

