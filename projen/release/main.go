package release

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.release.BranchOptions",
		reflect.TypeOf((*BranchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.release.CodeArtifactAuthProvider",
		reflect.TypeOf((*CodeArtifactAuthProvider)(nil)).Elem(),
		map[string]interface{}{
			"ACCESS_AND_SECRET_KEY_PAIR": CodeArtifactAuthProvider_ACCESS_AND_SECRET_KEY_PAIR,
			"GITHUB_OIDC": CodeArtifactAuthProvider_GITHUB_OIDC,
		},
	)
	_jsii_.RegisterStruct(
		"projen.release.CodeArtifactOptions",
		reflect.TypeOf((*CodeArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.CommonPublishOptions",
		reflect.TypeOf((*CommonPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.ContinuousReleaseOptions",
		reflect.TypeOf((*ContinuousReleaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.GitHubReleasesPublishOptions",
		reflect.TypeOf((*GitHubReleasesPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.GitPublishOptions",
		reflect.TypeOf((*GitPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.GoPublishOptions",
		reflect.TypeOf((*GoPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.JsiiReleaseGo",
		reflect.TypeOf((*JsiiReleaseGo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.JsiiReleaseMaven",
		reflect.TypeOf((*JsiiReleaseMaven)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.JsiiReleaseNpm",
		reflect.TypeOf((*JsiiReleaseNpm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.JsiiReleaseNuget",
		reflect.TypeOf((*JsiiReleaseNuget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.JsiiReleasePyPi",
		reflect.TypeOf((*JsiiReleasePyPi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.ManualReleaseOptions",
		reflect.TypeOf((*ManualReleaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.MavenPublishOptions",
		reflect.TypeOf((*MavenPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.NpmPublishOptions",
		reflect.TypeOf((*NpmPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.NugetPublishOptions",
		reflect.TypeOf((*NugetPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.release.Publisher",
		reflect.TypeOf((*Publisher)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGitHubPostPublishingSteps", GoMethod: "AddGitHubPostPublishingSteps"},
			_jsii_.MemberMethod{JsiiMethod: "addGitHubPrePublishingSteps", GoMethod: "AddGitHubPrePublishingSteps"},
			_jsii_.MemberProperty{JsiiProperty: "artifactName", GoGetter: "ArtifactName"},
			_jsii_.MemberProperty{JsiiProperty: "buildJobId", GoGetter: "BuildJobId"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "jsiiReleaseVersion", GoGetter: "JsiiReleaseVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "publibVersion", GoGetter: "PublibVersion"},
			_jsii_.MemberMethod{JsiiMethod: "publishToGit", GoMethod: "PublishToGit"},
			_jsii_.MemberMethod{JsiiMethod: "publishToGitHubReleases", GoMethod: "PublishToGitHubReleases"},
			_jsii_.MemberMethod{JsiiMethod: "publishToGo", GoMethod: "PublishToGo"},
			_jsii_.MemberMethod{JsiiMethod: "publishToMaven", GoMethod: "PublishToMaven"},
			_jsii_.MemberMethod{JsiiMethod: "publishToNpm", GoMethod: "PublishToNpm"},
			_jsii_.MemberMethod{JsiiMethod: "publishToNuget", GoMethod: "PublishToNuget"},
			_jsii_.MemberMethod{JsiiMethod: "publishToPyPi", GoMethod: "PublishToPyPi"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Publisher{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.release.PublisherOptions",
		reflect.TypeOf((*PublisherOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.PyPiPublishOptions",
		reflect.TypeOf((*PyPiPublishOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.release.Release",
		reflect.TypeOf((*Release)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addJobs", GoMethod: "AddJobs"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Release{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.release.ReleaseOptions",
		reflect.TypeOf((*ReleaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.ReleaseProjectOptions",
		reflect.TypeOf((*ReleaseProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.release.ReleaseTrigger",
		reflect.TypeOf((*ReleaseTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changelogPath", GoGetter: "ChangelogPath"},
			_jsii_.MemberProperty{JsiiProperty: "gitPushCommand", GoGetter: "GitPushCommand"},
			_jsii_.MemberProperty{JsiiProperty: "isContinuous", GoGetter: "IsContinuous"},
			_jsii_.MemberProperty{JsiiProperty: "isManual", GoGetter: "IsManual"},
			_jsii_.MemberProperty{JsiiProperty: "paths", GoGetter: "Paths"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
		},
		func() interface{} {
			return &jsiiProxy_ReleaseTrigger{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.release.ScheduledReleaseOptions",
		reflect.TypeOf((*ScheduledReleaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.release.TagReleaseOptions",
		reflect.TypeOf((*TagReleaseOptions)(nil)).Elem(),
	)
}
