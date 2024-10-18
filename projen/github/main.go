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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
		"projen.github.AutoQueue",
		reflect.TypeOf((*AutoQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AutoQueue{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.AutoQueueOptions",
		reflect.TypeOf((*AutoQueueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.CheckoutOptions",
		reflect.TypeOf((*CheckoutOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.CheckoutWith",
		reflect.TypeOf((*CheckoutWith)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.CheckoutWithPatchOptions",
		reflect.TypeOf((*CheckoutWithPatchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.ConcurrencyOptions",
		reflect.TypeOf((*ConcurrencyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.ContributorStatementOptions",
		reflect.TypeOf((*ContributorStatementOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.CreatePullRequestOptions",
		reflect.TypeOf((*CreatePullRequestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.Dependabot",
		reflect.TypeOf((*Dependabot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAllow", GoMethod: "AddAllow"},
			_jsii_.MemberMethod{JsiiMethod: "addIgnore", GoMethod: "AddIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "ignoresProjen", GoGetter: "IgnoresProjen"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Dependabot{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.DependabotAllow",
		reflect.TypeOf((*DependabotAllow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.DependabotGroup",
		reflect.TypeOf((*DependabotGroup)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"projen.github.DownloadArtifactOptions",
		reflect.TypeOf((*DownloadArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.DownloadArtifactWith",
		reflect.TypeOf((*DownloadArtifactWith)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.GitHub",
		reflect.TypeOf((*GitHub)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "addDependabot", GoMethod: "AddDependabot"},
			_jsii_.MemberMethod{JsiiMethod: "addPullRequestTemplate", GoMethod: "AddPullRequestTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "addWorkflow", GoMethod: "AddWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "downloadLfs", GoGetter: "DownloadLfs"},
			_jsii_.MemberProperty{JsiiProperty: "mergeQueue", GoGetter: "MergeQueue"},
			_jsii_.MemberProperty{JsiiProperty: "mergify", GoGetter: "Mergify"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenCredentials", GoGetter: "ProjenCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
	_jsii_.RegisterClass(
		"projen.github.GitHubActionsProvider",
		reflect.TypeOf((*GitHubActionsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "set", GoMethod: "Set"},
		},
		func() interface{} {
			return &jsiiProxy_GitHubActionsProvider{}
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
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
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
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
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
		"projen.github.GithubCredentials",
		reflect.TypeOf((*GithubCredentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "setupSteps", GoGetter: "SetupSteps"},
			_jsii_.MemberProperty{JsiiProperty: "tokenRef", GoGetter: "TokenRef"},
		},
		func() interface{} {
			return &jsiiProxy_GithubCredentials{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.GithubCredentialsAppOptions",
		reflect.TypeOf((*GithubCredentialsAppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.GithubCredentialsPersonalAccessTokenOptions",
		reflect.TypeOf((*GithubCredentialsPersonalAccessTokenOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.GithubWorkflow",
		reflect.TypeOf((*GithubWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addJob", GoMethod: "AddJob"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberProperty{JsiiProperty: "concurrency", GoGetter: "Concurrency"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberMethod{JsiiMethod: "getJob", GoMethod: "GetJob"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenCredentials", GoGetter: "ProjenCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "removeJob", GoMethod: "RemoveJob"},
			_jsii_.MemberProperty{JsiiProperty: "runName", GoGetter: "RunName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateJob", GoMethod: "UpdateJob"},
			_jsii_.MemberMethod{JsiiMethod: "updateJobs", GoMethod: "UpdateJobs"},
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
	_jsii_.RegisterEnum(
		"projen.github.MergeMethod",
		reflect.TypeOf((*MergeMethod)(nil)).Elem(),
		map[string]interface{}{
			"SQUASH": MergeMethod_SQUASH,
			"MERGE": MergeMethod_MERGE,
			"REBASE": MergeMethod_REBASE,
		},
	)
	_jsii_.RegisterClass(
		"projen.github.MergeQueue",
		reflect.TypeOf((*MergeQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MergeQueue{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.MergeQueueOptions",
		reflect.TypeOf((*MergeQueueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.Mergify",
		reflect.TypeOf((*Mergify)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addQueue", GoMethod: "AddQueue"},
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
		"projen.github.MergifyQueue",
		reflect.TypeOf((*MergifyQueue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.MergifyRule",
		reflect.TypeOf((*MergifyRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.PullRequestBackport",
		reflect.TypeOf((*PullRequestBackport)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workflow", GoGetter: "Workflow"},
		},
		func() interface{} {
			j := jsiiProxy_PullRequestBackport{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.PullRequestBackportOptions",
		reflect.TypeOf((*PullRequestBackportOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.PullRequestFromPatchOptions",
		reflect.TypeOf((*PullRequestFromPatchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.PullRequestLint",
		reflect.TypeOf((*PullRequestLint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
	_jsii_.RegisterStruct(
		"projen.github.PullRequestPatchSource",
		reflect.TypeOf((*PullRequestPatchSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.PullRequestTemplate",
		reflect.TypeOf((*PullRequestTemplate)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"projen.github.SetupGitIdentityOptions",
		reflect.TypeOf((*SetupGitIdentityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.github.Stale",
		reflect.TypeOf((*Stale)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberMethod{JsiiMethod: "getJob", GoMethod: "GetJob"},
			_jsii_.MemberProperty{JsiiProperty: "jobId", GoGetter: "JobId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "on", GoMethod: "On"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenCredentials", GoGetter: "ProjenCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "removeJob", GoMethod: "RemoveJob"},
			_jsii_.MemberProperty{JsiiProperty: "runName", GoGetter: "RunName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateJob", GoMethod: "UpdateJob"},
			_jsii_.MemberMethod{JsiiMethod: "updateJobs", GoMethod: "UpdateJobs"},
		},
		func() interface{} {
			j := jsiiProxy_TaskWorkflow{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GithubWorkflow)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.github.TaskWorkflowJob",
		reflect.TypeOf((*TaskWorkflowJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "concurrency", GoGetter: "Concurrency"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnError", GoGetter: "ContinueOnError"},
			_jsii_.MemberProperty{JsiiProperty: "defaults", GoGetter: "Defaults"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "if", GoGetter: "If"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "needs", GoGetter: "Needs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputs", GoGetter: "Outputs"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "runsOn", GoGetter: "RunsOn"},
			_jsii_.MemberProperty{JsiiProperty: "runsOnGroup", GoGetter: "RunsOnGroup"},
			_jsii_.MemberProperty{JsiiProperty: "services", GoGetter: "Services"},
			_jsii_.MemberProperty{JsiiProperty: "steps", GoGetter: "Steps"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMinutes", GoGetter: "TimeoutMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "tools", GoGetter: "Tools"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TaskWorkflowJob{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.github.TaskWorkflowJobOptions",
		reflect.TypeOf((*TaskWorkflowJobOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.TaskWorkflowOptions",
		reflect.TypeOf((*TaskWorkflowOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.UploadArtifactOptions",
		reflect.TypeOf((*UploadArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.UploadArtifactWith",
		reflect.TypeOf((*UploadArtifactWith)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.github.UploadGitPatchOptions",
		reflect.TypeOf((*UploadGitPatchOptions)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"projen.github.WorkflowActions",
		reflect.TypeOf((*WorkflowActions)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkflowActions{}
		},
	)
	_jsii_.RegisterClass(
		"projen.github.WorkflowJobs",
		reflect.TypeOf((*WorkflowJobs)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkflowJobs{}
		},
	)
	_jsii_.RegisterClass(
		"projen.github.WorkflowSteps",
		reflect.TypeOf((*WorkflowSteps)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkflowSteps{}
		},
	)
}
