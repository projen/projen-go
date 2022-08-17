// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Utility for applying RFC-6902 JSON-Patch to a document.
//
// Use the the `JsonPatch.apply(doc, ...ops)` function to apply a set of
// operations to a JSON document and return the result.
//
// Operations can be created using the factory methods `JsonPatch.add()`,
// `JsonPatch.remove()`, etc.
//
// Example:
//   const output = JsonPatch.apply(input,
//    JsonPatch.replace('/world/hi/there', 'goodbye'),
//    JsonPatch.add('/world/foo/', 'boom'),
//    JsonPatch.remove('/hello'));
//
// Experimental.
type JsonPatch interface {
}

// The jsii proxy struct for JsonPatch
type jsiiProxy_JsonPatch struct {
	_ byte // padding
}

// Adds a value to an object or inserts it into an array.
//
// In the case of an
// array, the value is inserted before the given index. The - character can be
// used instead of an index to insert at the end of an array.
//
// Example:
//   JsonPatch.add('/biscuits/1', { "name": "Ginger Nut" })
//
// Experimental.
func JsonPatch_Add(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"add",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

// Applies a set of JSON-Patch (RFC-6902) operations to `document` and returns the result.
//
// Returns: The result document.
// Experimental.
func JsonPatch_Apply(document interface{}, ops ...JsonPatch) interface{} {
	_init_.Initialize()

	args := []interface{}{document}
	for _, a := range ops {
		args = append(args, a)
	}

	var returns interface{}

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"apply",
		args,
		&returns,
	)

	return returns
}

// Copies a value from one location to another within the JSON document.
//
// Both
// from and path are JSON Pointers.
//
// Example:
//   JsonPatch.copy('/biscuits/0', '/best_biscuit')
//
// Experimental.
func JsonPatch_Copy(from *string, path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"copy",
		[]interface{}{from, path},
		&returns,
	)

	return returns
}

// Moves a value from one location to the other.
//
// Both from and path are JSON Pointers.
//
// Example:
//   JsonPatch.move('/biscuits', '/cookies')
//
// Experimental.
func JsonPatch_Move(from *string, path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"move",
		[]interface{}{from, path},
		&returns,
	)

	return returns
}

// Removes a value from an object or array.
//
// Example:
//   JsonPatch.remove('/biscuits/0')
//
// Experimental.
func JsonPatch_Remove(path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"remove",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Replaces a value.
//
// Equivalent to a “remove” followed by an “add”.
//
// Example:
//   JsonPatch.replace('/biscuits/0/name', 'Chocolate Digestive')
//
// Experimental.
func JsonPatch_Replace(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"replace",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

// Tests that the specified value is set in the document.
//
// If the test fails,
// then the patch as a whole should not apply.
//
// Example:
//   JsonPatch.test('/best_biscuit/name', 'Choco Leibniz')
//
// Experimental.
func JsonPatch_Test(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"projen.JsonPatch",
		"test",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

