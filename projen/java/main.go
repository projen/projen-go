package java

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen.java.JavaProject",
		reflect.TypeOf((*JavaProject)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
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
			j := jsiiProxy_JavaProject{}
			_jsii_.InitJsiiProxy(&j.Type__githubGitHubProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.JavaProjectCommonOptions",
		reflect.TypeOf((*JavaProjectCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.java.JavaProjectOptions",
		reflect.TypeOf((*JavaProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.Junit",
		reflect.TypeOf((*Junit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Junit{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.JunitOptions",
		reflect.TypeOf((*JunitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.MavenCompile",
		reflect.TypeOf((*MavenCompile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_MavenCompile{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.MavenCompileOptions",
		reflect.TypeOf((*MavenCompileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.MavenPackaging",
		reflect.TypeOf((*MavenPackaging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_MavenPackaging{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.MavenPackagingOptions",
		reflect.TypeOf((*MavenPackagingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.java.MavenRepository",
		reflect.TypeOf((*MavenRepository)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.MavenSample",
		reflect.TypeOf((*MavenSample)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_MavenSample{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.MavenSampleOptions",
		reflect.TypeOf((*MavenSampleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.java.PluginExecution",
		reflect.TypeOf((*PluginExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.java.PluginOptions",
		reflect.TypeOf((*PluginOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.Pom",
		reflect.TypeOf((*Pom)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addPlugin", GoMethod: "AddPlugin"},
			_jsii_.MemberMethod{JsiiMethod: "addProperty", GoMethod: "AddProperty"},
			_jsii_.MemberMethod{JsiiMethod: "addRepository", GoMethod: "AddRepository"},
			_jsii_.MemberMethod{JsiiMethod: "addTestDependency", GoMethod: "AddTestDependency"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Pom{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.java.PomOptions",
		reflect.TypeOf((*PomOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.java.Projenrc",
		reflect.TypeOf((*Projenrc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "className", GoGetter: "ClassName"},
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
		"projen.java.ProjenrcOptions",
		reflect.TypeOf((*ProjenrcOptions)(nil)).Elem(),
	)
}
