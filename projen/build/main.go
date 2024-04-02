package build

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.build.AddPostBuildJobCommandsOptions",
		reflect.TypeOf((*AddPostBuildJobCommandsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.build.AddPostBuildJobTaskOptions",
		reflect.TypeOf((*AddPostBuildJobTaskOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.build.BuildWorkflow",
		reflect.TypeOf((*BuildWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildJob", GoMethod: "AddPostBuildJob"},
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildJobCommands", GoMethod: "AddPostBuildJobCommands"},
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildJobTask", GoMethod: "AddPostBuildJobTask"},
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildSteps", GoMethod: "AddPostBuildSteps"},
			_jsii_.MemberProperty{JsiiProperty: "buildJobIds", GoGetter: "BuildJobIds"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BuildWorkflow{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.build.BuildWorkflowCommonOptions",
		reflect.TypeOf((*BuildWorkflowCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.build.BuildWorkflowOptions",
		reflect.TypeOf((*BuildWorkflowOptions)(nil)).Elem(),
	)
}
