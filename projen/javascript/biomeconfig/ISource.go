package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type ISource interface {
	// Enforce props sorting in JSX elements.
	// Experimental.
	SortJsxProps() *string
	// Experimental.
	SetSortJsxProps(s *string)
	// Sorts the keys of a JSON object in natural order.
	// Experimental.
	UseSortedKeys() *string
	// Experimental.
	SetUseSortedKeys(u *string)
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (j *jsiiProxy_ISource) SortJsxProps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortJsxProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource)SetSortJsxProps(val *string) {
	_jsii_.Set(
		j,
		"sortJsxProps",
		val,
	)
}

func (j *jsiiProxy_ISource) UseSortedKeys() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSortedKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource)SetUseSortedKeys(val *string) {
	_jsii_.Set(
		j,
		"useSortedKeys",
		val,
	)
}

