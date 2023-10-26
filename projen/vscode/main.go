package vscode

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"projen.vscode.Console",
		reflect.TypeOf((*Console)(nil)).Elem(),
		map[string]interface{}{
			"INTERNAL_CONSOLE": Console_INTERNAL_CONSOLE,
			"INTEGRATED_TERMINAL": Console_INTEGRATED_TERMINAL,
			"EXTERNAL_TERMINAL": Console_EXTERNAL_TERMINAL,
		},
	)
	_jsii_.RegisterClass(
		"projen.vscode.DevContainer",
		reflect.TypeOf((*DevContainer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addFeatures", GoMethod: "AddFeatures"},
			_jsii_.MemberMethod{JsiiMethod: "addPorts", GoMethod: "AddPorts"},
			_jsii_.MemberMethod{JsiiMethod: "addTasks", GoMethod: "AddTasks"},
			_jsii_.MemberMethod{JsiiMethod: "addVscodeExtensions", GoMethod: "AddVscodeExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DevContainer{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDevContainerEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.vscode.DevContainerFeature",
		reflect.TypeOf((*DevContainerFeature)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.vscode.DevContainerOptions",
		reflect.TypeOf((*DevContainerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"projen.vscode.IDevContainerEnvironment",
		reflect.TypeOf((*IDevContainerEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addFeatures", GoMethod: "AddFeatures"},
			_jsii_.MemberMethod{JsiiMethod: "addPorts", GoMethod: "AddPorts"},
			_jsii_.MemberMethod{JsiiMethod: "addTasks", GoMethod: "AddTasks"},
			_jsii_.MemberMethod{JsiiMethod: "addVscodeExtensions", GoMethod: "AddVscodeExtensions"},
		},
		func() interface{} {
			j := jsiiProxy_IDevContainerEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__projenIDevEnvironment)
			return &j
		},
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
			_jsii_.MemberProperty{JsiiProperty: "extensions", GoGetter: "Extensions"},
			_jsii_.MemberProperty{JsiiProperty: "launchConfiguration", GoGetter: "LaunchConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
	_jsii_.RegisterClass(
		"projen.vscode.VsCodeRecommendedExtensions",
		reflect.TypeOf((*VsCodeRecommendedExtensions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRecommendations", GoMethod: "AddRecommendations"},
			_jsii_.MemberMethod{JsiiMethod: "addUnwantedRecommendations", GoMethod: "AddUnwantedRecommendations"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VsCodeRecommendedExtensions{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.vscode.VsCodeSettings",
		reflect.TypeOf((*VsCodeSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSetting", GoMethod: "AddSetting"},
			_jsii_.MemberMethod{JsiiMethod: "addSettings", GoMethod: "AddSettings"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VsCodeSettings{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
}
