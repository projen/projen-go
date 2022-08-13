package vscode

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen.vscode.DevContainer",
		reflect.TypeOf((*DevContainer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addPorts", GoMethod: "AddPorts"},
			_jsii_.MemberMethod{JsiiMethod: "addTasks", GoMethod: "AddTasks"},
			_jsii_.MemberMethod{JsiiMethod: "addVscodeExtensions", GoMethod: "AddVscodeExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_DevContainer{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.Type__projenIDevEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.vscode.DevContainerOptions",
		reflect.TypeOf((*DevContainerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.vscode.InternalConsoleOptions",
		reflect.TypeOf((*InternalConsoleOptions)(nil)).Elem(),
		map[string]interface{}{
			"NEVER_OPEN": InternalConsoleOptions_NEVER_OPEN,
			"OPEN_ON_FIRST_SESSION_START": InternalConsoleOptions_OPEN_ON_FIRST_SESSION_START,
			"OPEN_ON_SESSION_START": InternalConsoleOptions_OPEN_ON_SESSION_START,
		},
	)
	_jsii_.RegisterStruct(
		"projen.vscode.Presentation",
		reflect.TypeOf((*Presentation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.vscode.ServerReadyAction",
		reflect.TypeOf((*ServerReadyAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.vscode.VsCode",
		reflect.TypeOf((*VsCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "launchConfiguration", GoGetter: "LaunchConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_VsCode{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.vscode.VsCodeLaunchConfig",
		reflect.TypeOf((*VsCodeLaunchConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfiguration", GoMethod: "AddConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_VsCodeLaunchConfig{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.vscode.VsCodeLaunchConfigurationEntry",
		reflect.TypeOf((*VsCodeLaunchConfigurationEntry)(nil)).Elem(),
	)
}
