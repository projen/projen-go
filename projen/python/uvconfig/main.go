package uvconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.BuildBackendSettings",
		reflect.TypeOf((*BuildBackendSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.DependencyGroupSettings",
		reflect.TypeOf((*DependencyGroupSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.Index",
		reflect.TypeOf((*Index)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.IndexCacheControl",
		reflect.TypeOf((*IndexCacheControl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.PipGroupName",
		reflect.TypeOf((*PipGroupName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.PipOptions",
		reflect.TypeOf((*PipOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.SchemaConflictItem",
		reflect.TypeOf((*SchemaConflictItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.StaticMetadata",
		reflect.TypeOf((*StaticMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.ToolUvWorkspace",
		reflect.TypeOf((*ToolUvWorkspace)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.TrustedPublishing",
		reflect.TypeOf((*TrustedPublishing)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": TrustedPublishing_ALWAYS,
			"NEVER": TrustedPublishing_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.UvConfiguration",
		reflect.TypeOf((*UvConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.WheelDataIncludes",
		reflect.TypeOf((*WheelDataIncludes)(nil)).Elem(),
	)
}
