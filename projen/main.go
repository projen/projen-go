// CDK for software projects
package projen

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen.Component",
		reflect.TypeOf((*Component)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Component{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.CreateProjectOptions",
		reflect.TypeOf((*CreateProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Dependencies",
		reflect.TypeOf((*Dependencies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberMethod{JsiiMethod: "getDependency", GoMethod: "GetDependency"},
			_jsii_.MemberMethod{JsiiMethod: "isDependencySatisfied", GoMethod: "IsDependencySatisfied"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryGetDependency", GoMethod: "TryGetDependency"},
		},
		func() interface{} {
			j := jsiiProxy_Dependencies{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.Dependency",
		reflect.TypeOf((*Dependency)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DependencyCoordinates",
		reflect.TypeOf((*DependencyCoordinates)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.DependencyType",
		reflect.TypeOf((*DependencyType)(nil)).Elem(),
		map[string]interface{}{
			"RUNTIME": DependencyType_RUNTIME,
			"PEER": DependencyType_PEER,
			"BUNDLED": DependencyType_BUNDLED,
			"BUILD": DependencyType_BUILD,
			"TEST": DependencyType_TEST,
			"DEVENV": DependencyType_DEVENV,
			"OVERRIDE": DependencyType_OVERRIDE,
			"OPTIONAL": DependencyType_OPTIONAL,
		},
	)
	_jsii_.RegisterStruct(
		"projen.DepsManifest",
		reflect.TypeOf((*DepsManifest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.DevEnvironmentDockerImage",
		reflect.TypeOf((*DevEnvironmentDockerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dockerFile", GoGetter: "DockerFile"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
		},
		func() interface{} {
			return &jsiiProxy_DevEnvironmentDockerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.DevEnvironmentOptions",
		reflect.TypeOf((*DevEnvironmentOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.DockerCompose",
		reflect.TypeOf((*DockerCompose)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addService", GoMethod: "AddService"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DockerCompose{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeBuild",
		reflect.TypeOf((*DockerComposeBuild)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeNetworkConfig",
		reflect.TypeOf((*DockerComposeNetworkConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeNetworkIpamConfig",
		reflect.TypeOf((*DockerComposeNetworkIpamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeNetworkIpamSubnetConfig",
		reflect.TypeOf((*DockerComposeNetworkIpamSubnetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposePortMappingOptions",
		reflect.TypeOf((*DockerComposePortMappingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeProps",
		reflect.TypeOf((*DockerComposeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.DockerComposeProtocol",
		reflect.TypeOf((*DockerComposeProtocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": DockerComposeProtocol_TCP,
			"UDP": DockerComposeProtocol_UDP,
		},
	)
	_jsii_.RegisterClass(
		"projen.DockerComposeService",
		reflect.TypeOf((*DockerComposeService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addLabel", GoMethod: "AddLabel"},
			_jsii_.MemberMethod{JsiiMethod: "addNetwork", GoMethod: "AddNetwork"},
			_jsii_.MemberMethod{JsiiMethod: "addPort", GoMethod: "AddPort"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageBuild", GoGetter: "ImageBuild"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "networks", GoGetter: "Networks"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_DockerComposeService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDockerComposeServiceName)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeServiceDescription",
		reflect.TypeOf((*DockerComposeServiceDescription)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeServicePort",
		reflect.TypeOf((*DockerComposeServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeVolumeConfig",
		reflect.TypeOf((*DockerComposeVolumeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.DockerComposeVolumeMount",
		reflect.TypeOf((*DockerComposeVolumeMount)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.EndOfLine",
		reflect.TypeOf((*EndOfLine)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": EndOfLine_AUTO,
			"CRLF": EndOfLine_CRLF,
			"LF": EndOfLine_LF,
			"NONE": EndOfLine_NONE,
		},
	)
	_jsii_.RegisterClass(
		"projen.FileBase",
		reflect.TypeOf((*FileBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FileBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.FileBaseOptions",
		reflect.TypeOf((*FileBaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.GitAttributesFile",
		reflect.TypeOf((*GitAttributesFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addAttributes", GoMethod: "AddAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "addLfsPattern", GoMethod: "AddLfsPattern"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "endOfLine", GoGetter: "EndOfLine"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "hasLfsPatterns", GoGetter: "HasLfsPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GitAttributesFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.GitAttributesFileOptions",
		reflect.TypeOf((*GitAttributesFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.GitOptions",
		reflect.TypeOf((*GitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Gitpod",
		reflect.TypeOf((*Gitpod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCustomTask", GoMethod: "AddCustomTask"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addPorts", GoMethod: "AddPorts"},
			_jsii_.MemberMethod{JsiiMethod: "addPrebuilds", GoMethod: "AddPrebuilds"},
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
			j := jsiiProxy_Gitpod{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDevEnvironment)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"projen.GitpodOnOpen",
		reflect.TypeOf((*GitpodOnOpen)(nil)).Elem(),
		map[string]interface{}{
			"OPEN_BROWSER": GitpodOnOpen_OPEN_BROWSER,
			"OPEN_PREVIEW": GitpodOnOpen_OPEN_PREVIEW,
			"NOTIFY": GitpodOnOpen_NOTIFY,
			"IGNORE": GitpodOnOpen_IGNORE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.GitpodOpenIn",
		reflect.TypeOf((*GitpodOpenIn)(nil)).Elem(),
		map[string]interface{}{
			"BOTTOM": GitpodOpenIn_BOTTOM,
			"LEFT": GitpodOpenIn_LEFT,
			"RIGHT": GitpodOpenIn_RIGHT,
			"MAIN": GitpodOpenIn_MAIN,
		},
	)
	_jsii_.RegisterEnum(
		"projen.GitpodOpenMode",
		reflect.TypeOf((*GitpodOpenMode)(nil)).Elem(),
		map[string]interface{}{
			"TAB_AFTER": GitpodOpenMode_TAB_AFTER,
			"TAB_BEFORE": GitpodOpenMode_TAB_BEFORE,
			"SPLIT_RIGHT": GitpodOpenMode_SPLIT_RIGHT,
			"SPLIT_LEFT": GitpodOpenMode_SPLIT_LEFT,
			"SPLIT_TOP": GitpodOpenMode_SPLIT_TOP,
			"SPLIT_BOTTOM": GitpodOpenMode_SPLIT_BOTTOM,
		},
	)
	_jsii_.RegisterStruct(
		"projen.GitpodOptions",
		reflect.TypeOf((*GitpodOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.GitpodPort",
		reflect.TypeOf((*GitpodPort)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.GitpodPortVisibility",
		reflect.TypeOf((*GitpodPortVisibility)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": GitpodPortVisibility_PUBLIC,
			"PRIVATE": GitpodPortVisibility_PRIVATE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.GitpodPrebuilds",
		reflect.TypeOf((*GitpodPrebuilds)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.GitpodTask",
		reflect.TypeOf((*GitpodTask)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.GroupRunnerOptions",
		reflect.TypeOf((*GroupRunnerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"projen.ICompareString",
		reflect.TypeOf((*ICompareString)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "compare", GoMethod: "Compare"},
		},
		func() interface{} {
			return &jsiiProxy_ICompareString{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDevEnvironment",
		reflect.TypeOf((*IDevEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addPorts", GoMethod: "AddPorts"},
			_jsii_.MemberMethod{JsiiMethod: "addTasks", GoMethod: "AddTasks"},
			_jsii_.MemberMethod{JsiiMethod: "addVscodeExtensions", GoMethod: "AddVscodeExtensions"},
		},
		func() interface{} {
			return &jsiiProxy_IDevEnvironment{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDockerComposeNetworkBinding",
		reflect.TypeOf((*IDockerComposeNetworkBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDockerComposeNetworkBinding{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDockerComposeNetworkConfig",
		reflect.TypeOf((*IDockerComposeNetworkConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addNetworkConfiguration", GoMethod: "AddNetworkConfiguration"},
		},
		func() interface{} {
			return &jsiiProxy_IDockerComposeNetworkConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDockerComposeServiceName",
		reflect.TypeOf((*IDockerComposeServiceName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
		},
		func() interface{} {
			return &jsiiProxy_IDockerComposeServiceName{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDockerComposeVolumeBinding",
		reflect.TypeOf((*IDockerComposeVolumeBinding)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDockerComposeVolumeBinding{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IDockerComposeVolumeConfig",
		reflect.TypeOf((*IDockerComposeVolumeConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolumeConfiguration", GoMethod: "AddVolumeConfiguration"},
		},
		func() interface{} {
			return &jsiiProxy_IDockerComposeVolumeConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IResolvable",
		reflect.TypeOf((*IResolvable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
		},
		func() interface{} {
			return &jsiiProxy_IResolvable{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.IResolver",
		reflect.TypeOf((*IResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
		},
		func() interface{} {
			return &jsiiProxy_IResolver{}
		},
	)
	_jsii_.RegisterClass(
		"projen.IgnoreFile",
		reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addPatterns", GoMethod: "AddPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberMethod{JsiiMethod: "exclude", GoMethod: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "filterCommentLines", GoGetter: "FilterCommentLines"},
			_jsii_.MemberProperty{JsiiProperty: "filterEmptyLines", GoGetter: "FilterEmptyLines"},
			_jsii_.MemberMethod{JsiiMethod: "include", GoMethod: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "removePatterns", GoMethod: "RemovePatterns"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IgnoreFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.IgnoreFileOptions",
		reflect.TypeOf((*IgnoreFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.IniFile",
		reflect.TypeOf((*IniFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IniFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ObjectFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.IniFileOptions",
		reflect.TypeOf((*IniFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.InitProject",
		reflect.TypeOf((*InitProject)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.InitProjectOptionHints",
		reflect.TypeOf((*InitProjectOptionHints)(nil)).Elem(),
		map[string]interface{}{
			"ALL": InitProjectOptionHints_ALL,
			"FEATURED": InitProjectOptionHints_FEATURED,
			"NONE": InitProjectOptionHints_NONE,
		},
	)
	_jsii_.RegisterClass(
		"projen.JsonFile",
		reflect.TypeOf((*JsonFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberProperty{JsiiProperty: "supportsComments", GoGetter: "SupportsComments"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JsonFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ObjectFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.JsonFileOptions",
		reflect.TypeOf((*JsonFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.JsonPatch",
		reflect.TypeOf((*JsonPatch)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JsonPatch{}
		},
	)
	_jsii_.RegisterClass(
		"projen.License",
		reflect.TypeOf((*License)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_License{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.LicenseOptions",
		reflect.TypeOf((*LicenseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"OFF": LogLevel_OFF,
			"ERROR": LogLevel_ERROR,
			"WARN": LogLevel_WARN,
			"INFO": LogLevel_INFO,
			"DEBUG": LogLevel_DEBUG,
			"VERBOSE": LogLevel_VERBOSE,
		},
	)
	_jsii_.RegisterClass(
		"projen.Logger",
		reflect.TypeOf((*Logger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "debug", GoMethod: "Debug"},
			_jsii_.MemberMethod{JsiiMethod: "error", GoMethod: "Error"},
			_jsii_.MemberMethod{JsiiMethod: "info", GoMethod: "Info"},
			_jsii_.MemberMethod{JsiiMethod: "log", GoMethod: "Log"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "verbose", GoMethod: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "warn", GoMethod: "Warn"},
		},
		func() interface{} {
			j := jsiiProxy_Logger{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.LoggerOptions",
		reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Makefile",
		reflect.TypeOf((*Makefile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addAll", GoMethod: "AddAll"},
			_jsii_.MemberMethod{JsiiMethod: "addAlls", GoMethod: "AddAlls"},
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRules", GoMethod: "AddRules"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Makefile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.MakefileOptions",
		reflect.TypeOf((*MakefileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.ObjectFile",
		reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ObjectFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.ObjectFileOptions",
		reflect.TypeOf((*ObjectFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Project",
		reflect.TypeOf((*Project)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberProperty{JsiiProperty: "subprojects", GoGetter: "Subprojects"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
		},
		func() interface{} {
			j := jsiiProxy_Project{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.ProjectBuild",
		reflect.TypeOf((*ProjectBuild)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectBuild{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.ProjectOptions",
		reflect.TypeOf((*ProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.ProjectTree",
		reflect.TypeOf((*ProjectTree)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectTree{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"projen.ProjectType",
		reflect.TypeOf((*ProjectType)(nil)).Elem(),
		map[string]interface{}{
			"UNKNOWN": ProjectType_UNKNOWN,
			"LIB": ProjectType_LIB,
			"APP": ProjectType_APP,
		},
	)
	_jsii_.RegisterClass(
		"projen.Projects",
		reflect.TypeOf((*Projects)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Projects{}
		},
	)
	_jsii_.RegisterClass(
		"projen.Projenrc",
		reflect.TypeOf((*Projenrc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Projenrc{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ProjenrcJson)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.ProjenrcFile",
		reflect.TypeOf((*ProjenrcFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProjenrcFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.ProjenrcJson",
		reflect.TypeOf((*ProjenrcJson)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProjenrcJson{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ProjenrcFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.ProjenrcJsonOptions",
		reflect.TypeOf((*ProjenrcJsonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.ProjenrcOptions",
		reflect.TypeOf((*ProjenrcOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.ReleasableCommits",
		reflect.TypeOf((*ReleasableCommits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cmd", GoGetter: "Cmd"},
		},
		func() interface{} {
			return &jsiiProxy_ReleasableCommits{}
		},
	)
	_jsii_.RegisterClass(
		"projen.Renovatebot",
		reflect.TypeOf((*Renovatebot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Renovatebot{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.RenovatebotOptions",
		reflect.TypeOf((*RenovatebotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.RenovatebotScheduleInterval",
		reflect.TypeOf((*RenovatebotScheduleInterval)(nil)).Elem(),
		map[string]interface{}{
			"ANY_TIME": RenovatebotScheduleInterval_ANY_TIME,
			"EARLY_MONDAYS": RenovatebotScheduleInterval_EARLY_MONDAYS,
			"DAILY": RenovatebotScheduleInterval_DAILY,
			"MONTHLY": RenovatebotScheduleInterval_MONTHLY,
			"QUARTERLY": RenovatebotScheduleInterval_QUARTERLY,
			"WEEKENDS": RenovatebotScheduleInterval_WEEKENDS,
			"WEEKDAYS": RenovatebotScheduleInterval_WEEKDAYS,
		},
	)
	_jsii_.RegisterStruct(
		"projen.ResolveOptions",
		reflect.TypeOf((*ResolveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.Rule",
		reflect.TypeOf((*Rule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.SampleDir",
		reflect.TypeOf((*SampleDir)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SampleDir{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.SampleDirOptions",
		reflect.TypeOf((*SampleDirOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.SampleFile",
		reflect.TypeOf((*SampleFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SampleFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.SampleFileOptions",
		reflect.TypeOf((*SampleFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.SampleReadme",
		reflect.TypeOf((*SampleReadme)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SampleReadme{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SampleFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.SampleReadmeProps",
		reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Semver",
		reflect.TypeOf((*Semver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_Semver{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.SnapshotOptions",
		reflect.TypeOf((*SnapshotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.SourceCode",
		reflect.TypeOf((*SourceCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "close", GoMethod: "Close"},
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberMethod{JsiiMethod: "line", GoMethod: "Line"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "open", GoMethod: "Open"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SourceCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.SourceCodeOptions",
		reflect.TypeOf((*SourceCodeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Task",
		reflect.TypeOf((*Task)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCondition", GoMethod: "AddCondition"},
			_jsii_.MemberMethod{JsiiMethod: "builtin", GoMethod: "Builtin"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "cwd", GoGetter: "Cwd"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "env", GoMethod: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "envVars", GoGetter: "EnvVars"},
			_jsii_.MemberMethod{JsiiMethod: "exec", GoMethod: "Exec"},
			_jsii_.MemberMethod{JsiiMethod: "insertStep", GoMethod: "InsertStep"},
			_jsii_.MemberMethod{JsiiMethod: "lock", GoMethod: "Lock"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "prepend", GoMethod: "Prepend"},
			_jsii_.MemberMethod{JsiiMethod: "prependExec", GoMethod: "PrependExec"},
			_jsii_.MemberMethod{JsiiMethod: "prependSay", GoMethod: "PrependSay"},
			_jsii_.MemberMethod{JsiiMethod: "prependSpawn", GoMethod: "PrependSpawn"},
			_jsii_.MemberMethod{JsiiMethod: "removeStep", GoMethod: "RemoveStep"},
			_jsii_.MemberMethod{JsiiMethod: "reset", GoMethod: "Reset"},
			_jsii_.MemberMethod{JsiiMethod: "say", GoMethod: "Say"},
			_jsii_.MemberMethod{JsiiMethod: "spawn", GoMethod: "Spawn"},
			_jsii_.MemberProperty{JsiiProperty: "steps", GoGetter: "Steps"},
			_jsii_.MemberMethod{JsiiMethod: "updateStep", GoMethod: "UpdateStep"},
		},
		func() interface{} {
			return &jsiiProxy_Task{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.TaskCommonOptions",
		reflect.TypeOf((*TaskCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.TaskOptions",
		reflect.TypeOf((*TaskOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.TaskRuntime",
		reflect.TypeOf((*TaskRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberMethod{JsiiMethod: "runTask", GoMethod: "RunTask"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindTask", GoMethod: "TryFindTask"},
			_jsii_.MemberProperty{JsiiProperty: "workdir", GoGetter: "Workdir"},
		},
		func() interface{} {
			return &jsiiProxy_TaskRuntime{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.TaskSpec",
		reflect.TypeOf((*TaskSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.TaskStep",
		reflect.TypeOf((*TaskStep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.TaskStepOptions",
		reflect.TypeOf((*TaskStepOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Tasks",
		reflect.TypeOf((*Tasks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryFind", GoMethod: "TryFind"},
		},
		func() interface{} {
			j := jsiiProxy_Tasks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.TasksManifest",
		reflect.TypeOf((*TasksManifest)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.TestFailureBehavior",
		reflect.TypeOf((*TestFailureBehavior)(nil)).Elem(),
		map[string]interface{}{
			"SKIP": TestFailureBehavior_SKIP,
			"FAIL_SYNTHESIS": TestFailureBehavior_FAIL_SYNTHESIS,
		},
	)
	_jsii_.RegisterClass(
		"projen.Testing",
		reflect.TypeOf((*Testing)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Testing{}
		},
	)
	_jsii_.RegisterClass(
		"projen.TextFile",
		reflect.TypeOf((*TextFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addLine", GoMethod: "AddLine"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TextFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.TextFileOptions",
		reflect.TypeOf((*TextFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.TomlFile",
		reflect.TypeOf((*TomlFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TomlFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ObjectFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.TomlFileOptions",
		reflect.TypeOf((*TomlFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.Version",
		reflect.TypeOf((*Version)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bumpPackage", GoGetter: "BumpPackage"},
			_jsii_.MemberProperty{JsiiProperty: "bumpTask", GoGetter: "BumpTask"},
			_jsii_.MemberProperty{JsiiProperty: "changelogFileName", GoGetter: "ChangelogFileName"},
			_jsii_.MemberMethod{JsiiMethod: "envForBranch", GoMethod: "EnvForBranch"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "releaseTagFileName", GoGetter: "ReleaseTagFileName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unbumpTask", GoGetter: "UnbumpTask"},
			_jsii_.MemberProperty{JsiiProperty: "versionFileName", GoGetter: "VersionFileName"},
		},
		func() interface{} {
			j := jsiiProxy_Version{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.VersionBranchOptions",
		reflect.TypeOf((*VersionBranchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.VersionOptions",
		reflect.TypeOf((*VersionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.XmlFile",
		reflect.TypeOf((*XmlFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_XmlFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ObjectFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.XmlFileOptions",
		reflect.TypeOf((*XmlFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.YamlFile",
		reflect.TypeOf((*YamlFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addToArray", GoMethod: "AddToArray"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "omitEmpty", GoGetter: "OmitEmpty"},
			_jsii_.MemberMethod{JsiiMethod: "patch", GoMethod: "Patch"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_YamlFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ObjectFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.YamlFileOptions",
		reflect.TypeOf((*YamlFileOptions)(nil)).Elem(),
	)
}
