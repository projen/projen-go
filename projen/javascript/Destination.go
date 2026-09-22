package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Where a reporter's output is written.
// See: https://nodejs.org/api/test.html#test-reporters
//
// Experimental.
type Destination interface {
	// The underlying value: `"stdout"`, `"stderr"`, or a file path.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Destination
type jsiiProxy_Destination struct {
	_ byte // padding
}

func (j *jsiiProxy_Destination) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Write to a file at the given path.
// Experimental.
func Destination_File(path *string) Destination {
	_init_.Initialize()

	if err := validateDestination_FileParameters(path); err != nil {
		panic(err)
	}
	var returns Destination

	_jsii_.StaticInvoke(
		"projen.javascript.Destination",
		"file",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func Destination_STDERR() Destination {
	_init_.Initialize()
	var returns Destination
	_jsii_.StaticGet(
		"projen.javascript.Destination",
		"STDERR",
		&returns,
	)
	return returns
}

func Destination_STDOUT() Destination {
	_init_.Initialize()
	var returns Destination
	_jsii_.StaticGet(
		"projen.javascript.Destination",
		"STDOUT",
		&returns,
	)
	return returns
}

