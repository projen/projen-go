package gitlab

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"projen.gitlab.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		map[string]interface{}{
			"PREPARE": Action_PREPARE,
			"START": Action_START,
			"STOP": Action_STOP,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.AllowFailure",
		reflect.TypeOf((*AllowFailure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Artifacts",
		reflect.TypeOf((*Artifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Assets",
		reflect.TypeOf((*Assets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Cache",
		reflect.TypeOf((*Cache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.CacheKeyFiles",
		reflect.TypeOf((*CacheKeyFiles)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.CachePolicy",
		reflect.TypeOf((*CachePolicy)(nil)).Elem(),
		map[string]interface{}{
			"PULL": CachePolicy_PULL,
			"PUSH": CachePolicy_PUSH,
			"PULL_PUSH": CachePolicy_PULL_PUSH,
		},
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.CacheWhen",
		reflect.TypeOf((*CacheWhen)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": CacheWhen_ALWAYS,
			"ON_FAILURE": CacheWhen_ON_FAILURE,
			"ON_SUCCESS": CacheWhen_ON_SUCCESS,
		},
	)
	_jsii_.RegisterClass(
		"projen.gitlab.CiConfiguration",
		reflect.TypeOf((*CiConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCaches", GoMethod: "AddDefaultCaches"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultHooks", GoMethod: "AddDefaultHooks"},
			_jsii_.MemberMethod{JsiiMethod: "addGlobalVariables", GoMethod: "AddGlobalVariables"},
			_jsii_.MemberMethod{JsiiMethod: "addIncludes", GoMethod: "AddIncludes"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberMethod{JsiiMethod: "addServices", GoMethod: "AddServices"},
			_jsii_.MemberMethod{JsiiMethod: "addStages", GoMethod: "AddStages"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAfterScript", GoGetter: "DefaultAfterScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultArtifacts", GoGetter: "DefaultArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBeforeScript", GoGetter: "DefaultBeforeScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCache", GoGetter: "DefaultCache"},
			_jsii_.MemberProperty{JsiiProperty: "defaultIdTokens", GoGetter: "DefaultIdTokens"},
			_jsii_.MemberProperty{JsiiProperty: "defaultImage", GoGetter: "DefaultImage"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInterruptible", GoGetter: "DefaultInterruptible"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRetry", GoGetter: "DefaultRetry"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTags", GoGetter: "DefaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTimeout", GoGetter: "DefaultTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pages", GoGetter: "Pages"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
			_jsii_.MemberProperty{JsiiProperty: "workflow", GoGetter: "Workflow"},
		},
		func() interface{} {
			j := jsiiProxy_CiConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.CiConfigurationOptions",
		reflect.TypeOf((*CiConfigurationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.CoverageReport",
		reflect.TypeOf((*CoverageReport)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Default",
		reflect.TypeOf((*Default)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.DefaultElement",
		reflect.TypeOf((*DefaultElement)(nil)).Elem(),
		map[string]interface{}{
			"AFTER_SCRIPT": DefaultElement_AFTER_SCRIPT,
			"ARTIFACTS": DefaultElement_ARTIFACTS,
			"BEFORE_SCRIPT": DefaultElement_BEFORE_SCRIPT,
			"CACHE": DefaultElement_CACHE,
			"IMAGE": DefaultElement_IMAGE,
			"INTERRUPTIBLE": DefaultElement_INTERRUPTIBLE,
			"RETRY": DefaultElement_RETRY,
			"SERVICES": DefaultElement_SERVICES,
			"TAGS": DefaultElement_TAGS,
			"TIMEOUT": DefaultElement_TIMEOUT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.DefaultHooks",
		reflect.TypeOf((*DefaultHooks)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.DeploymentTier",
		reflect.TypeOf((*DeploymentTier)(nil)).Elem(),
		map[string]interface{}{
			"DEVELOPMENT": DeploymentTier_DEVELOPMENT,
			"OTHER": DeploymentTier_OTHER,
			"PRODUCTION": DeploymentTier_PRODUCTION,
			"STAGING": DeploymentTier_STAGING,
			"TESTING": DeploymentTier_TESTING,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Engine",
		reflect.TypeOf((*Engine)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Filter",
		reflect.TypeOf((*Filter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.gitlab.GitlabConfiguration",
		reflect.TypeOf((*GitlabConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCaches", GoMethod: "AddDefaultCaches"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultHooks", GoMethod: "AddDefaultHooks"},
			_jsii_.MemberMethod{JsiiMethod: "addGlobalVariables", GoMethod: "AddGlobalVariables"},
			_jsii_.MemberMethod{JsiiMethod: "addIncludes", GoMethod: "AddIncludes"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberMethod{JsiiMethod: "addServices", GoMethod: "AddServices"},
			_jsii_.MemberMethod{JsiiMethod: "addStages", GoMethod: "AddStages"},
			_jsii_.MemberMethod{JsiiMethod: "createNestedTemplates", GoMethod: "CreateNestedTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAfterScript", GoGetter: "DefaultAfterScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultArtifacts", GoGetter: "DefaultArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBeforeScript", GoGetter: "DefaultBeforeScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCache", GoGetter: "DefaultCache"},
			_jsii_.MemberProperty{JsiiProperty: "defaultIdTokens", GoGetter: "DefaultIdTokens"},
			_jsii_.MemberProperty{JsiiProperty: "defaultImage", GoGetter: "DefaultImage"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInterruptible", GoGetter: "DefaultInterruptible"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRetry", GoGetter: "DefaultRetry"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTags", GoGetter: "DefaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTimeout", GoGetter: "DefaultTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nestedTemplates", GoGetter: "NestedTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pages", GoGetter: "Pages"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
			_jsii_.MemberProperty{JsiiProperty: "workflow", GoGetter: "Workflow"},
		},
		func() interface{} {
			j := jsiiProxy_GitlabConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CiConfiguration)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"projen.gitlab.IDToken",
		reflect.TypeOf((*IDToken)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aud", GoGetter: "Aud"},
		},
		func() interface{} {
			return &jsiiProxy_IDToken{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Image",
		reflect.TypeOf((*Image)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Include",
		reflect.TypeOf((*Include)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.IncludeRule",
		reflect.TypeOf((*IncludeRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Inherit",
		reflect.TypeOf((*Inherit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.JobWhen",
		reflect.TypeOf((*JobWhen)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": JobWhen_ALWAYS,
			"DELAYED": JobWhen_DELAYED,
			"MANUAL": JobWhen_MANUAL,
			"NEVER": JobWhen_NEVER,
			"ON_FAILURE": JobWhen_ON_FAILURE,
			"ON_SUCCESS": JobWhen_ON_SUCCESS,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.KubernetesConfig",
		reflect.TypeOf((*KubernetesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.KubernetesEnum",
		reflect.TypeOf((*KubernetesEnum)(nil)).Elem(),
		map[string]interface{}{
			"ACTIVE": KubernetesEnum_ACTIVE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Link",
		reflect.TypeOf((*Link)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.LinkType",
		reflect.TypeOf((*LinkType)(nil)).Elem(),
		map[string]interface{}{
			"IMAGE": LinkType_IMAGE,
			"OTHER": LinkType_OTHER,
			"PACKAGE": LinkType_PACKAGE,
			"RUNBOOK": LinkType_RUNBOOK,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Need",
		reflect.TypeOf((*Need)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.gitlab.NestedConfiguration",
		reflect.TypeOf((*NestedConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCaches", GoMethod: "AddDefaultCaches"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultHooks", GoMethod: "AddDefaultHooks"},
			_jsii_.MemberMethod{JsiiMethod: "addGlobalVariables", GoMethod: "AddGlobalVariables"},
			_jsii_.MemberMethod{JsiiMethod: "addIncludes", GoMethod: "AddIncludes"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberMethod{JsiiMethod: "addServices", GoMethod: "AddServices"},
			_jsii_.MemberMethod{JsiiMethod: "addStages", GoMethod: "AddStages"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAfterScript", GoGetter: "DefaultAfterScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultArtifacts", GoGetter: "DefaultArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBeforeScript", GoGetter: "DefaultBeforeScript"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCache", GoGetter: "DefaultCache"},
			_jsii_.MemberProperty{JsiiProperty: "defaultIdTokens", GoGetter: "DefaultIdTokens"},
			_jsii_.MemberProperty{JsiiProperty: "defaultImage", GoGetter: "DefaultImage"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInterruptible", GoGetter: "DefaultInterruptible"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRetry", GoGetter: "DefaultRetry"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTags", GoGetter: "DefaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTimeout", GoGetter: "DefaultTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "pages", GoGetter: "Pages"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "stages", GoGetter: "Stages"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
			_jsii_.MemberProperty{JsiiProperty: "workflow", GoGetter: "Workflow"},
		},
		func() interface{} {
			j := jsiiProxy_NestedConfiguration{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_CiConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Parallel",
		reflect.TypeOf((*Parallel)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.PullPolicy",
		reflect.TypeOf((*PullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": PullPolicy_ALWAYS,
			"NEVER": PullPolicy_NEVER,
			"IF_NOT_PRESENT": PullPolicy_IF_NOT_PRESENT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Release",
		reflect.TypeOf((*Release)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Reports",
		reflect.TypeOf((*Reports)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Retry",
		reflect.TypeOf((*Retry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.Strategy",
		reflect.TypeOf((*Strategy)(nil)).Elem(),
		map[string]interface{}{
			"DEPEND": Strategy_DEPEND,
			"MIRROR": Strategy_MIRROR,
		},
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Trigger",
		reflect.TypeOf((*Trigger)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.TriggerInclude",
		reflect.TypeOf((*TriggerInclude)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.VariableConfig",
		reflect.TypeOf((*VariableConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.VaultConfig",
		reflect.TypeOf((*VaultConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.Workflow",
		reflect.TypeOf((*Workflow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.gitlab.WorkflowRule",
		reflect.TypeOf((*WorkflowRule)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.gitlab.WorkflowWhen",
		reflect.TypeOf((*WorkflowWhen)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": WorkflowWhen_ALWAYS,
			"NEVER": WorkflowWhen_NEVER,
		},
	)
}
