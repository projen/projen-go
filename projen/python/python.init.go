package python

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"projen.python.IPackageProvider",
		reflect.TypeOf((*IPackageProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "packages", GoGetter: "Packages"},
		},
		func() interface{} {
			return &jsiiProxy_IPackageProvider{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.python.IPythonDeps",
		reflect.TypeOf((*IPythonDeps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDependency", GoMethod: "AddDevDependency"},
			_jsii_.MemberMethod{JsiiMethod: "installDependencies", GoMethod: "InstallDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "installTask", GoGetter: "InstallTask"},
		},
		func() interface{} {
			return &jsiiProxy_IPythonDeps{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.python.IPythonEnv",
		reflect.TypeOf((*IPythonEnv)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "setupEnvironment", GoMethod: "SetupEnvironment"},
		},
		func() interface{} {
			return &jsiiProxy_IPythonEnv{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.python.IPythonPackaging",
		reflect.TypeOf((*IPythonPackaging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "publishTask", GoGetter: "PublishTask"},
		},
		func() interface{} {
			return &jsiiProxy_IPythonPackaging{}
		},
	)
	_jsii_.RegisterClass(
		"projen.python.Pip",
		reflect.TypeOf((*Pip)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDependency", GoMethod: "AddDevDependency"},
			_jsii_.MemberMethod{JsiiMethod: "installDependencies", GoMethod: "InstallDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "installTask", GoGetter: "InstallTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Pip{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonDeps)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PipOptions",
		reflect.TypeOf((*PipOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.Poetry",
		reflect.TypeOf((*Poetry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDependency", GoMethod: "AddDevDependency"},
			_jsii_.MemberMethod{JsiiMethod: "installDependencies", GoMethod: "InstallDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "installTask", GoGetter: "InstallTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "publishTask", GoGetter: "PublishTask"},
			_jsii_.MemberProperty{JsiiProperty: "publishTestTask", GoGetter: "PublishTestTask"},
			_jsii_.MemberMethod{JsiiMethod: "setupEnvironment", GoMethod: "SetupEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Poetry{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonDeps)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonEnv)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonPackaging)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.python.PoetryPyproject",
		reflect.TypeOf((*PoetryPyproject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_PoetryPyproject{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PoetryPyprojectOptions",
		reflect.TypeOf((*PoetryPyprojectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.PoetryPyprojectOptionsWithoutDeps",
		reflect.TypeOf((*PoetryPyprojectOptionsWithoutDeps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.Projenrc",
		reflect.TypeOf((*Projenrc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Projenrc{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.ProjenrcOptions",
		reflect.TypeOf((*ProjenrcOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.Pytest",
		reflect.TypeOf((*Pytest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "testdir", GoGetter: "Testdir"},
		},
		func() interface{} {
			j := jsiiProxy_Pytest{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PytestOptions",
		reflect.TypeOf((*PytestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.PytestSample",
		reflect.TypeOf((*PytestSample)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_PytestSample{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PythonPackagingOptions",
		reflect.TypeOf((*PythonPackagingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.PythonProject",
		reflect.TypeOf((*PythonProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDependency", GoMethod: "AddDevDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "depsManager", GoGetter: "DepsManager"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "envManager", GoGetter: "EnvManager"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "moduleName", GoGetter: "ModuleName"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "packagingManager", GoGetter: "PackagingManager"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "pytest", GoGetter: "Pytest"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
		},
		func() interface{} {
			j := jsiiProxy_PythonProject{}
			_jsii_.InitJsiiProxy(&j.Type__githubGitHubProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PythonProjectOptions",
		reflect.TypeOf((*PythonProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.PythonSample",
		reflect.TypeOf((*PythonSample)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_PythonSample{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.PythonSampleOptions",
		reflect.TypeOf((*PythonSampleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.RequirementsFile",
		reflect.TypeOf((*RequirementsFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addPackages", GoMethod: "AddPackages"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
		},
		func() interface{} {
			j := jsiiProxy_RequirementsFile{}
			_jsii_.InitJsiiProxy(&j.Type__projenFileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.RequirementsFileOptions",
		reflect.TypeOf((*RequirementsFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.SetupPy",
		reflect.TypeOf((*SetupPy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberProperty{JsiiProperty: "changed", GoGetter: "Changed"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "marker", GoGetter: "Marker"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
		},
		func() interface{} {
			j := jsiiProxy_SetupPy{}
			_jsii_.InitJsiiProxy(&j.Type__projenFileBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.SetupPyOptions",
		reflect.TypeOf((*SetupPyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.python.Setuptools",
		reflect.TypeOf((*Setuptools)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "publishTask", GoGetter: "PublishTask"},
			_jsii_.MemberProperty{JsiiProperty: "publishTestTask", GoGetter: "PublishTestTask"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Setuptools{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonPackaging)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.python.Venv",
		reflect.TypeOf((*Venv)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "setupEnvironment", GoMethod: "SetupEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Venv{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPythonEnv)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.VenvOptions",
		reflect.TypeOf((*VenvOptions)(nil)).Elem(),
	)
}
