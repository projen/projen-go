package cdk8s

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen.cdk8s.AutoDiscover",
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
		"projen.cdk8s.AutoDiscoverOptions",
		reflect.TypeOf((*AutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.Cdk8sDeps",
		reflect.TypeOf((*Cdk8sDeps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cdk8sMajorVersion", GoGetter: "Cdk8sMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdk8sMinimumVersion", GoGetter: "Cdk8sMinimumVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdk8sVersion", GoGetter: "Cdk8sVersion"},
			_jsii_.MemberMethod{JsiiMethod: "packageNames", GoMethod: "PackageNames"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Cdk8sDeps{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.Cdk8sDepsCommonOptions",
		reflect.TypeOf((*Cdk8sDepsCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.Cdk8sDepsOptions",
		reflect.TypeOf((*Cdk8sDepsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.Cdk8sDepsPy",
		reflect.TypeOf((*Cdk8sDepsPy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cdk8sMajorVersion", GoGetter: "Cdk8sMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdk8sMinimumVersion", GoGetter: "Cdk8sMinimumVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdk8sVersion", GoGetter: "Cdk8sVersion"},
			_jsii_.MemberMethod{JsiiMethod: "packageNames", GoMethod: "PackageNames"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Cdk8sDepsPy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Cdk8sDeps)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.Cdk8sPackageNames",
		reflect.TypeOf((*Cdk8sPackageNames)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.Cdk8sPythonApp",
		reflect.TypeOf((*Cdk8sPythonApp)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cdk8sDeps", GoGetter: "Cdk8sDeps"},
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
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
			_jsii_.MemberProperty{JsiiProperty: "subprojects", GoGetter: "Subprojects"},
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
			j := jsiiProxy_Cdk8sPythonApp{}
			_jsii_.InitJsiiProxy(&j.Type__pythonPythonProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.Cdk8sPythonOptions",
		reflect.TypeOf((*Cdk8sPythonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		reflect.TypeOf((*Cdk8sTypeScriptApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addScripts", GoMethod: "AddScripts"},
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
			_jsii_.MemberProperty{JsiiProperty: "cdk8sDeps", GoGetter: "Cdk8sDeps"},
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
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
			_jsii_.MemberProperty{JsiiProperty: "nodeVersion", GoGetter: "NodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "npmrc", GoGetter: "Npmrc"},
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
			_jsii_.MemberProperty{JsiiProperty: "subprojects", GoGetter: "Subprojects"},
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
			_jsii_.MemberProperty{JsiiProperty: "workflowBootstrapSteps", GoGetter: "WorkflowBootstrapSteps"},
		},
		func() interface{} {
			j := jsiiProxy_Cdk8sTypeScriptApp{}
			_jsii_.InitJsiiProxy(&j.Type__typescriptTypeScriptAppProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.Cdk8sTypeScriptAppOptions",
		reflect.TypeOf((*Cdk8sTypeScriptAppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.ConstructLibraryCdk8s",
		reflect.TypeOf((*ConstructLibraryCdk8s)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addScripts", GoMethod: "AddScripts"},
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
			_jsii_.MemberProperty{JsiiProperty: "cdk8sVersion", GoGetter: "Cdk8sVersion"},
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "constructsVersion", GoGetter: "ConstructsVersion"},
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
			_jsii_.MemberProperty{JsiiProperty: "nodeVersion", GoGetter: "NodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "npmrc", GoGetter: "Npmrc"},
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
			_jsii_.MemberProperty{JsiiProperty: "subprojects", GoGetter: "Subprojects"},
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
			_jsii_.MemberProperty{JsiiProperty: "workflowBootstrapSteps", GoGetter: "WorkflowBootstrapSteps"},
		},
		func() interface{} {
			j := jsiiProxy_ConstructLibraryCdk8s{}
			_jsii_.InitJsiiProxy(&j.Type__cdkConstructLibrary)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.ConstructLibraryCdk8sOptions",
		reflect.TypeOf((*ConstructLibraryCdk8sOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.IntegrationTest",
		reflect.TypeOf((*IntegrationTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertTask", GoGetter: "AssertTask"},
			_jsii_.MemberProperty{JsiiProperty: "deployTask", GoGetter: "DeployTask"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotDir", GoGetter: "SnapshotDir"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotTask", GoGetter: "SnapshotTask"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tmpDir", GoGetter: "TmpDir"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationTest{}
			_jsii_.InitJsiiProxy(&j.Type__cdkIntegrationTestBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.cdk8s.IntegrationTestAutoDiscover",
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
		"projen.cdk8s.IntegrationTestAutoDiscoverOptions",
		reflect.TypeOf((*IntegrationTestAutoDiscoverOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.cdk8s.IntegrationTestOptions",
		reflect.TypeOf((*IntegrationTestOptions)(nil)).Elem(),
	)
}
