package circleci

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.circleci.CircleCiProps",
		reflect.TypeOf((*CircleCiProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.circleci.Circleci",
		reflect.TypeOf((*Circleci)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOrb", GoMethod: "AddOrb"},
			_jsii_.MemberMethod{JsiiMethod: "addWorkflow", GoMethod: "AddWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Circleci{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Docker",
		reflect.TypeOf((*Docker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Filter",
		reflect.TypeOf((*Filter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.FilterConfig",
		reflect.TypeOf((*FilterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.circleci.JobType",
		reflect.TypeOf((*JobType)(nil)).Elem(),
		map[string]interface{}{
			"APPROVAL": JobType_APPROVAL,
		},
	)
	_jsii_.RegisterEnum(
		"projen.circleci.JobWhen",
		reflect.TypeOf((*JobWhen)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": JobWhen_ALWAYS,
			"ON_SUCCESS": JobWhen_ON_SUCCESS,
			"ON_FAIL": JobWhen_ON_FAIL,
		},
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Machine",
		reflect.TypeOf((*Machine)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Macos",
		reflect.TypeOf((*Macos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Matrix",
		reflect.TypeOf((*Matrix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.PipelineParameter",
		reflect.TypeOf((*PipelineParameter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.circleci.PipelineParameterType",
		reflect.TypeOf((*PipelineParameterType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": PipelineParameterType_STRING,
			"BOOLEAN": PipelineParameterType_BOOLEAN,
			"INTEGER": PipelineParameterType_INTEGER,
			"ENUM": PipelineParameterType_ENUM,
		},
	)
	_jsii_.RegisterEnum(
		"projen.circleci.ResourceClass",
		reflect.TypeOf((*ResourceClass)(nil)).Elem(),
		map[string]interface{}{
			"SMALL": ResourceClass_SMALL,
			"MEDIUM": ResourceClass_MEDIUM,
			"MEDIUM_PLUS": ResourceClass_MEDIUM_PLUS,
			"LARGE_X": ResourceClass_LARGE_X,
			"LARGE_2X": ResourceClass_LARGE_2X,
			"LARGE_2X_PLUS": ResourceClass_LARGE_2X_PLUS,
		},
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Run",
		reflect.TypeOf((*Run)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.StepRun",
		reflect.TypeOf((*StepRun)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Triggers",
		reflect.TypeOf((*Triggers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.Workflow",
		reflect.TypeOf((*Workflow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.circleci.WorkflowJob",
		reflect.TypeOf((*WorkflowJob)(nil)).Elem(),
	)
}
