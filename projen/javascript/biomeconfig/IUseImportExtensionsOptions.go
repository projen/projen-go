package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IUseImportExtensionsOptions interface {
	// A map of custom import extension mappings, where the key is the inspected file extension, and the value is a pair of `module` extension and `component` import extension.
	// Experimental.
	SuggestedExtensions() *map[string]ISuggestedExtensionMapping
	// Experimental.
	SetSuggestedExtensions(s *map[string]ISuggestedExtensionMapping)
}

// The jsii proxy for IUseImportExtensionsOptions
type jsiiProxy_IUseImportExtensionsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IUseImportExtensionsOptions) SuggestedExtensions() *map[string]ISuggestedExtensionMapping {
	var returns *map[string]ISuggestedExtensionMapping
	_jsii_.Get(
		j,
		"suggestedExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseImportExtensionsOptions)SetSuggestedExtensions(val *map[string]ISuggestedExtensionMapping) {
	_jsii_.Set(
		j,
		"suggestedExtensions",
		val,
	)
}

