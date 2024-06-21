package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ICompareString interface {
	// Returns: It is expected to return a negative value if the first argument is less than the second argument, zero if they're equal, and a positive value otherwise.
	// Experimental.
	Compare(a *string, b *string) *float64
}

// The jsii proxy for ICompareString
type jsiiProxy_ICompareString struct {
	_ byte // padding
}

func (i *jsiiProxy_ICompareString) Compare(a *string, b *string) *float64 {
	if err := i.validateCompareParameters(a, b); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"compare",
		[]interface{}{a, b},
		&returns,
	)

	return returns
}

