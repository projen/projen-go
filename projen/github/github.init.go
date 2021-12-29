package github

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen.github.AutoApprove",
		reflect.TypeOf((*AutoApprove)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AutoApprove{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.AutoApproveOptions",
		reflect.TypeOf((*AutoApproveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.AutoMerge",
		reflect.TypeOf((*AutoMerge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConditions", GoMethod: "AddConditions"},
			_jsii_.MemberMethod{JsiiMethod: "addConditionsLater", GoMethod: "AddConditionsLater"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_AutoMerge{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.AutoMergeOptions",
		reflect.TypeOf((*AutoMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.Dependabot",
		reflect.TypeOf((*Dependabot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIgnore", GoMethod: "AddIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "ignoresProjen", GoGetter: "IgnoresProjen"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Dependabot{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.DependabotIgnore",
		reflect.TypeOf((*DependabotIgnore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.DependabotOptions",
		reflect.TypeOf((*DependabotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.DependabotRegistry",
		reflect.TypeOf((*DependabotRegistry)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.github.DependabotRegistryType",
		reflect.TypeOf((*DependabotRegistryType)(nil)).Elem(),
		map[string]interface{}{
			"COMPOSER_REGISTRY": DependabotRegistryType_COMPOSER_REGISTRY,
			"DOCKER_REGISTRY": DependabotRegistryType_DOCKER_REGISTRY,
			"GIT": DependabotRegistryType_GIT,
			"HEX_ORGANIZATION": DependabotRegistryType_HEX_ORGANIZATION,
			"MAVEN_REPOSITORY": DependabotRegistryType_MAVEN_REPOSITORY,
			"NPM_REGISTRY": DependabotRegistryType_NPM_REGISTRY,
			"NUGET_FEED": DependabotRegistryType_NUGET_FEED,
			"PYTHON_INDEX": DependabotRegistryType_PYTHON_INDEX,
			"RUBYGEMS_SERVER": DependabotRegistryType_RUBYGEMS_SERVER,
			"TERRAFORM_REGISTRY": DependabotRegistryType_TERRAFORM_REGISTRY,
		},
	)
	_jsii_.RegisterEnum(
		"projen.github.DependabotScheduleInterval",
		reflect.TypeOf((*DependabotScheduleInterval)(nil)).Elem(),
		map[string]interface{}{
			"DAILY": DependabotScheduleInterval_DAILY,
			"WEEKLY": DependabotScheduleInterval_WEEKLY,
			"MONTHLY": DependabotScheduleInterval_MONTHLY,
		},
	)
	_jsii_.RegisterClass(
		"projen.github.GitHub",
		reflect.TypeOf((*GitHub)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependabot", GoMethod: "AddDependabot"},
			_jsii_.MemberMethod{JsiiMethod: "addPullRequestTemplate", GoMethod: "AddPullRequestTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "addWorkflow", GoMethod: "AddWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "mergify", GoGetter: "Mergify"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenTokenSecret", GoGetter: "ProjenTokenSecret"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindWorkflow", GoMethod: "TryFindWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "workflows", GoGetter: "Workflows"},
			_jsii_.MemberProperty{JsiiProperty: "workflowsEnabled", GoGetter: "WorkflowsEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_GitHub{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.GitHubOptions",
		reflect.TypeOf((*GitHubOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.GitHubProject",
		reflect.TypeOf((*GitHubProject)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubProject{}
			_jsii_.InitJsiiProxy(&j.Type__projenProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.GitHubProjectOptions",
		reflect.TypeOf((*GitHubProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.GitIdentity",
		reflect.TypeOf((*GitIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.GithubWorkflow",
		reflect.TypeOf((*GithubWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addJob", GoMethod: "AddJob"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberProperty{JsiiProperty: "concurrency", GoGetter: "Concurrency"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenTokenSecret", GoGetter: "ProjenTokenSecret"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_GithubWorkflow{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.GithubWorkflowOptions",
		reflect.TypeOf((*GithubWorkflowOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"projen.github.IAddConditionsLater",
		reflect.TypeOf((*IAddConditionsLater)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_IAddConditionsLater{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.github.IJobProvider",
		reflect.TypeOf((*IJobProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderJobs", GoMethod: "RenderJobs"},
		},
		func() interface{} {
			return &jsiiProxy_IJobProvider{}
		},
	)
	_jsii_.RegisterClass(
		"projen.github.Mergify",
		reflect.TypeOf((*Mergify)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Mergify{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.MergifyConditionalOperator",
		reflect.TypeOf((*MergifyConditionalOperator)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.MergifyOptions",
		reflect.TypeOf((*MergifyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.MergifyRule",
		reflect.TypeOf((*MergifyRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.PullRequestLint",
		reflect.TypeOf((*PullRequestLint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_PullRequestLint{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.PullRequestLintOptions",
		reflect.TypeOf((*PullRequestLintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.PullRequestTemplate",
		reflect.TypeOf((*PullRequestTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "absolutePath", GoGetter: "AbsolutePath"},
			_jsii_.MemberMethod{JsiiMethod: "addLine", GoMethod: "AddLine"},
			_jsii_.MemberProperty{JsiiProperty: "executable", GoGetter: "Executable"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeContent", GoMethod: "SynthesizeContent"},
		},
		func() interface{} {
			j := jsiiProxy_PullRequestTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__projenTextFile)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.PullRequestTemplateOptions",
		reflect.TypeOf((*PullRequestTemplateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.SemanticTitleOptions",
		reflect.TypeOf((*SemanticTitleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.Stale",
		reflect.TypeOf((*Stale)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Stale{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.StaleBehavior",
		reflect.TypeOf((*StaleBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.StaleOptions",
		reflect.TypeOf((*StaleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.TaskWorkflow",
		reflect.TypeOf((*TaskWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addJob", GoMethod: "AddJob"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "concurrency", GoGetter: "Concurrency"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "jobId", GoGetter: "JobId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenTokenSecret", GoGetter: "ProjenTokenSecret"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_TaskWorkflow{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GithubWorkflow)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.TaskWorkflowOptions",
		reflect.TypeOf((*TaskWorkflowOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.github.VersioningStrategy",
		reflect.TypeOf((*VersioningStrategy)(nil)).Elem(),
		map[string]interface{}{
			"LOCKFILE_ONLY": VersioningStrategy_LOCKFILE_ONLY,
			"AUTO": VersioningStrategy_AUTO,
			"WIDEN": VersioningStrategy_WIDEN,
			"INCREASE": VersioningStrategy_INCREASE,
			"INCREASE_IF_NECESSARY": VersioningStrategy_INCREASE_IF_NECESSARY,
		},
	)
}
