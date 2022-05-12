package awscdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"projen.awscdk.ApprovalLevel",
		reflect.TypeOf((*ApprovalLevel)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": ApprovalLevel_NEVER,
			"ANY_CHANGE": ApprovalLevel_ANY_CHANGE,
			"BROADENING": ApprovalLevel_BROADENING,
		},
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AutoDiscover",
		reflect.TypeOf((*AutoDiscover)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AutoDiscover{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AutoDiscoverCommonOptions",
		reflect.TypeOf((*AutoDiscoverCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AutoDiscoverOptions",
		reflect.TypeOf((*AutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkConstructLibrary",
		reflect.TypeOf((*AwsCdkConstructLibrary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCdkDependencies", GoMethod: "AddCdkDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addCdkTestDependencies", GoMethod: "AddCdkTestDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTestCommand", GoMethod: "AddTestCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberProperty{JsiiProperty: "allowLibraryDependencies", GoGetter: "AllowLibraryDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsJavascriptDirectory", GoGetter: "ArtifactsJavascriptDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "autoMerge", GoGetter: "AutoMerge"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflow", GoGetter: "BuildWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflowJobId", GoGetter: "BuildWorkflowJobId"},
			_jsii_.MemberProperty{JsiiProperty: "bundler", GoGetter: "Bundler"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDeps", GoGetter: "CdkDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "docgen", GoGetter: "Docgen"},
			_jsii_.MemberProperty{JsiiProperty: "docsDirectory", GoGetter: "DocsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "eslint", GoGetter: "Eslint"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "jest", GoGetter: "Jest"},
			_jsii_.MemberProperty{JsiiProperty: "libdir", GoGetter: "Libdir"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "package", GoGetter: "Package"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "prettier", GoGetter: "Prettier"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "release", GoGetter: "Release"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberMethod{JsiiMethod: "renderWorkflowSetup", GoMethod: "RenderWorkflowSetup"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "runScriptCommand", GoGetter: "RunScriptCommand"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberProperty{JsiiProperty: "srcdir", GoGetter: "Srcdir"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testdir", GoGetter: "Testdir"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfig", GoGetter: "Tsconfig"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigDev", GoGetter: "TsconfigDev"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigEslint", GoGetter: "TsconfigEslint"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeWorkflow", GoGetter: "UpgradeWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
			_jsii_.MemberProperty{JsiiProperty: "watchTask", GoGetter: "WatchTask"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkConstructLibrary{}
			_jsii_.InitJsiiProxy(&j.Type__cdkConstructLibrary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkConstructLibraryOptions",
		reflect.TypeOf((*AwsCdkConstructLibraryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkDeps",
		reflect.TypeOf((*AwsCdkDeps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addV1Dependencies", GoMethod: "AddV1Dependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addV1DevDependencies", GoMethod: "AddV1DevDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDependenciesAsDeps", GoGetter: "CdkDependenciesAsDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMajorVersion", GoGetter: "CdkMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMinimumVersion", GoGetter: "CdkMinimumVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberMethod{JsiiMethod: "packageNames", GoMethod: "PackageNames"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkDeps{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkDepsCommonOptions",
		reflect.TypeOf((*AwsCdkDepsCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkDepsJava",
		reflect.TypeOf((*AwsCdkDepsJava)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addV1Dependencies", GoMethod: "AddV1Dependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addV1DevDependencies", GoMethod: "AddV1DevDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDependenciesAsDeps", GoGetter: "CdkDependenciesAsDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMajorVersion", GoGetter: "CdkMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMinimumVersion", GoGetter: "CdkMinimumVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberMethod{JsiiMethod: "packageNames", GoMethod: "PackageNames"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkDepsJava{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsCdkDeps)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkDepsJs",
		reflect.TypeOf((*AwsCdkDepsJs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addV1Dependencies", GoMethod: "AddV1Dependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addV1DevDependencies", GoMethod: "AddV1DevDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDependenciesAsDeps", GoGetter: "CdkDependenciesAsDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMajorVersion", GoGetter: "CdkMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkMinimumVersion", GoGetter: "CdkMinimumVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberMethod{JsiiMethod: "packageNames", GoMethod: "PackageNames"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkDepsJs{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsCdkDeps)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkDepsOptions",
		reflect.TypeOf((*AwsCdkDepsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkJavaApp",
		reflect.TypeOf((*AwsCdkJavaApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCdkDependency", GoMethod: "AddCdkDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPlugin", GoMethod: "AddPlugin"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTestDependency", GoMethod: "AddTestDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "cdkConfig", GoGetter: "CdkConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDeps", GoGetter: "CdkDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTasks", GoGetter: "CdkTasks"},
			_jsii_.MemberProperty{JsiiProperty: "compile", GoGetter: "Compile"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "distdir", GoGetter: "Distdir"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "junit", GoGetter: "Junit"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "mainClass", GoGetter: "MainClass"},
			_jsii_.MemberProperty{JsiiProperty: "mainClassName", GoGetter: "MainClassName"},
			_jsii_.MemberProperty{JsiiProperty: "mainPackage", GoGetter: "MainPackage"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "pom", GoGetter: "Pom"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "projenrc", GoGetter: "Projenrc"},
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
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkJavaApp{}
			_jsii_.InitJsiiProxy(&j.Type__javaJavaProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkJavaAppOptions",
		reflect.TypeOf((*AwsCdkJavaAppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkPackageNames",
		reflect.TypeOf((*AwsCdkPackageNames)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkPythonApp",
		reflect.TypeOf((*AwsCdkPythonApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDependency", GoMethod: "AddDevDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "appEntrypoint", GoGetter: "AppEntrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "cdkConfig", GoGetter: "CdkConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDeps", GoGetter: "CdkDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTasks", GoGetter: "CdkTasks"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
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
			_jsii_.MemberProperty{JsiiProperty: "testdir", GoGetter: "Testdir"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkPythonApp{}
			_jsii_.InitJsiiProxy(&j.Type__pythonPythonProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkPythonAppOptions",
		reflect.TypeOf((*AwsCdkPythonAppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.AwsCdkTypeScriptApp",
		reflect.TypeOf((*AwsCdkTypeScriptApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCdkDependency", GoMethod: "AddCdkDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTestCommand", GoMethod: "AddTestCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberProperty{JsiiProperty: "allowLibraryDependencies", GoGetter: "AllowLibraryDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "appEntrypoint", GoGetter: "AppEntrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsJavascriptDirectory", GoGetter: "ArtifactsJavascriptDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "autoMerge", GoGetter: "AutoMerge"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflow", GoGetter: "BuildWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflowJobId", GoGetter: "BuildWorkflowJobId"},
			_jsii_.MemberProperty{JsiiProperty: "bundler", GoGetter: "Bundler"},
			_jsii_.MemberProperty{JsiiProperty: "cdkConfig", GoGetter: "CdkConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDeps", GoGetter: "CdkDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTasks", GoGetter: "CdkTasks"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "docgen", GoGetter: "Docgen"},
			_jsii_.MemberProperty{JsiiProperty: "docsDirectory", GoGetter: "DocsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "eslint", GoGetter: "Eslint"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "jest", GoGetter: "Jest"},
			_jsii_.MemberProperty{JsiiProperty: "libdir", GoGetter: "Libdir"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "package", GoGetter: "Package"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "prettier", GoGetter: "Prettier"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "release", GoGetter: "Release"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberMethod{JsiiMethod: "renderWorkflowSetup", GoMethod: "RenderWorkflowSetup"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "runScriptCommand", GoGetter: "RunScriptCommand"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberProperty{JsiiProperty: "srcdir", GoGetter: "Srcdir"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testdir", GoGetter: "Testdir"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfig", GoGetter: "Tsconfig"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigDev", GoGetter: "TsconfigDev"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigEslint", GoGetter: "TsconfigEslint"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeWorkflow", GoGetter: "UpgradeWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
			_jsii_.MemberProperty{JsiiProperty: "watchTask", GoGetter: "WatchTask"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkTypeScriptApp{}
			_jsii_.InitJsiiProxy(&j.Type__typescriptTypeScriptAppProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.AwsCdkTypeScriptAppOptions",
		reflect.TypeOf((*AwsCdkTypeScriptAppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.CdkConfig",
		reflect.TypeOf((*CdkConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cdkout", GoGetter: "Cdkout"},
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_CdkConfig{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.CdkConfigCommonOptions",
		reflect.TypeOf((*CdkConfigCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.CdkConfigOptions",
		reflect.TypeOf((*CdkConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.CdkTasks",
		reflect.TypeOf((*CdkTasks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploy", GoGetter: "Deploy"},
			_jsii_.MemberProperty{JsiiProperty: "destroy", GoGetter: "Destroy"},
			_jsii_.MemberProperty{JsiiProperty: "diff", GoGetter: "Diff"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "synth", GoGetter: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "synthSilent", GoGetter: "SynthSilent"},
			_jsii_.MemberProperty{JsiiProperty: "watch", GoGetter: "Watch"},
		},
		func() interface{} {
			j := jsiiProxy_CdkTasks{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.awscdk.ConstructLibraryAws",
		reflect.TypeOf((*ConstructLibraryAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCdkDependencies", GoMethod: "AddCdkDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addCdkTestDependencies", GoMethod: "AddCdkTestDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTestCommand", GoMethod: "AddTestCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberProperty{JsiiProperty: "allowLibraryDependencies", GoGetter: "AllowLibraryDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsJavascriptDirectory", GoGetter: "ArtifactsJavascriptDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "autoMerge", GoGetter: "AutoMerge"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflow", GoGetter: "BuildWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflowJobId", GoGetter: "BuildWorkflowJobId"},
			_jsii_.MemberProperty{JsiiProperty: "bundler", GoGetter: "Bundler"},
			_jsii_.MemberProperty{JsiiProperty: "cdkDeps", GoGetter: "CdkDeps"},
			_jsii_.MemberProperty{JsiiProperty: "cdkVersion", GoGetter: "CdkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "docgen", GoGetter: "Docgen"},
			_jsii_.MemberProperty{JsiiProperty: "docsDirectory", GoGetter: "DocsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "eslint", GoGetter: "Eslint"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "jest", GoGetter: "Jest"},
			_jsii_.MemberProperty{JsiiProperty: "libdir", GoGetter: "Libdir"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "package", GoGetter: "Package"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "prettier", GoGetter: "Prettier"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "release", GoGetter: "Release"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberMethod{JsiiMethod: "renderWorkflowSetup", GoMethod: "RenderWorkflowSetup"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "runScriptCommand", GoGetter: "RunScriptCommand"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberProperty{JsiiProperty: "srcdir", GoGetter: "Srcdir"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testdir", GoGetter: "Testdir"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfig", GoGetter: "Tsconfig"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigDev", GoGetter: "TsconfigDev"},
			_jsii_.MemberProperty{JsiiProperty: "tsconfigEslint", GoGetter: "TsconfigEslint"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeWorkflow", GoGetter: "UpgradeWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
			_jsii_.MemberProperty{JsiiProperty: "watchTask", GoGetter: "WatchTask"},
		},
		func() interface{} {
			j := jsiiProxy_ConstructLibraryAws{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsCdkConstructLibrary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.ConstructLibraryAwsOptions",
		reflect.TypeOf((*ConstructLibraryAwsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		reflect.TypeOf((*EdgeLambdaAutoDiscover)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entrypoints", GoGetter: "Entrypoints"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_EdgeLambdaAutoDiscover{}
			_jsii_.InitJsiiProxy(&j.Type__cdkAutoDiscoverBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.EdgeLambdaAutoDiscoverOptions",
		reflect.TypeOf((*EdgeLambdaAutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.IntegrationTest",
		reflect.TypeOf((*IntegrationTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertTask", GoGetter: "AssertTask"},
			_jsii_.MemberProperty{JsiiProperty: "deployTask", GoGetter: "DeployTask"},
			_jsii_.MemberProperty{JsiiProperty: "destroyTask", GoGetter: "DestroyTask"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotDir", GoGetter: "SnapshotDir"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotTask", GoGetter: "SnapshotTask"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tmpDir", GoGetter: "TmpDir"},
			_jsii_.MemberProperty{JsiiProperty: "watchTask", GoGetter: "WatchTask"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationTest{}
			_jsii_.InitJsiiProxy(&j.Type__cdkIntegrationTestBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.awscdk.IntegrationTestAutoDiscover",
		reflect.TypeOf((*IntegrationTestAutoDiscover)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entrypoints", GoGetter: "Entrypoints"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationTestAutoDiscover{}
			_jsii_.InitJsiiProxy(&j.Type__cdkIntegrationTestAutoDiscoverBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.IntegrationTestAutoDiscoverOptions",
		reflect.TypeOf((*IntegrationTestAutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.IntegrationTestCommonOptions",
		reflect.TypeOf((*IntegrationTestCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.IntegrationTestOptions",
		reflect.TypeOf((*IntegrationTestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.LambdaAutoDiscover",
		reflect.TypeOf((*LambdaAutoDiscover)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entrypoints", GoGetter: "Entrypoints"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaAutoDiscover{}
			_jsii_.InitJsiiProxy(&j.Type__cdkAutoDiscoverBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaAutoDiscoverOptions",
		reflect.TypeOf((*LambdaAutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.LambdaExtension",
		reflect.TypeOf((*LambdaExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaExtension{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		reflect.TypeOf((*LambdaExtensionAutoDiscover)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entrypoints", GoGetter: "Entrypoints"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaExtensionAutoDiscover{}
			_jsii_.InitJsiiProxy(&j.Type__cdkAutoDiscoverBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaExtensionAutoDiscoverOptions",
		reflect.TypeOf((*LambdaExtensionAutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaExtensionCommonOptions",
		reflect.TypeOf((*LambdaExtensionCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaExtensionOptions",
		reflect.TypeOf((*LambdaExtensionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.LambdaFunction",
		reflect.TypeOf((*LambdaFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunction{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaFunctionCommonOptions",
		reflect.TypeOf((*LambdaFunctionCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.awscdk.LambdaFunctionOptions",
		reflect.TypeOf((*LambdaFunctionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.awscdk.LambdaRuntime",
		reflect.TypeOf((*LambdaRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "esbuildPlatform", GoGetter: "EsbuildPlatform"},
			_jsii_.MemberProperty{JsiiProperty: "esbuildTarget", GoGetter: "EsbuildTarget"},
			_jsii_.MemberProperty{JsiiProperty: "functionRuntime", GoGetter: "FunctionRuntime"},
		},
		func() interface{} {
			return &jsiiProxy_LambdaRuntime{}
		},
	)
}
