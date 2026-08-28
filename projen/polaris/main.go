package polaris

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.polaris.AnalysisConfiguration",
		reflect.TypeOf((*AnalysisConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.AnalysisConfigurationAggressivenessLevel",
		reflect.TypeOf((*AnalysisConfigurationAggressivenessLevel)(nil)).Elem(),
		map[string]interface{}{
			"LOW": AnalysisConfigurationAggressivenessLevel_LOW,
			"MEDIUM": AnalysisConfigurationAggressivenessLevel_MEDIUM,
			"HIGH": AnalysisConfigurationAggressivenessLevel_HIGH,
		},
	)
	_jsii_.RegisterEnum(
		"projen.polaris.AnalysisConfigurationLocation",
		reflect.TypeOf((*AnalysisConfigurationLocation)(nil)).Elem(),
		map[string]interface{}{
			"LOCAL": AnalysisConfigurationLocation_LOCAL,
			"CONNECT": AnalysisConfigurationLocation_CONNECT,
			"SRM": AnalysisConfigurationLocation_SRM,
		},
	)
	_jsii_.RegisterEnum(
		"projen.polaris.AnalysisConfigurationMode",
		reflect.TypeOf((*AnalysisConfigurationMode)(nil)).Elem(),
		map[string]interface{}{
			"HFI": AnalysisConfigurationMode_HFI,
			"PFI": AnalysisConfigurationMode_PFI,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.AnalyzeConnectConfiguration",
		reflect.TypeOf((*AnalyzeConnectConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.AnalyzeConnectConfigurationUploadArtifacts",
		reflect.TypeOf((*AnalyzeConnectConfigurationUploadArtifacts)(nil)).Elem(),
		map[string]interface{}{
			"ALL": AnalyzeConnectConfigurationUploadArtifacts_ALL,
			"LOGS_ONLY": AnalyzeConnectConfigurationUploadArtifacts_LOGS_ONLY,
			"NONE": AnalyzeConnectConfigurationUploadArtifacts_NONE,
			"ON_FAILURE": AnalyzeConnectConfigurationUploadArtifacts_ON_FAILURE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.AnalyzeFilesConfiguration",
		reflect.TypeOf((*AnalyzeFilesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.BuildConfiguration",
		reflect.TypeOf((*BuildConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CachingConfiguration",
		reflect.TypeOf((*CachingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CaptureConfiguration",
		reflect.TypeOf((*CaptureConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CheckerConfiguration",
		reflect.TypeOf((*CheckerConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CheckerConfigurationWebappSecurity",
		reflect.TypeOf((*CheckerConfigurationWebappSecurity)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.CheckerConfigurationWebappSecurityAggressivenessLevel",
		reflect.TypeOf((*CheckerConfigurationWebappSecurityAggressivenessLevel)(nil)).Elem(),
		map[string]interface{}{
			"LOW": CheckerConfigurationWebappSecurityAggressivenessLevel_LOW,
			"MEDIUM": CheckerConfigurationWebappSecurityAggressivenessLevel_MEDIUM,
			"HIGH": CheckerConfigurationWebappSecurityAggressivenessLevel_HIGH,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CodingStandardConfiguration",
		reflect.TypeOf((*CodingStandardConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CodingStandardDeviation",
		reflect.TypeOf((*CodingStandardDeviation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CommitConfiguration",
		reflect.TypeOf((*CommitConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CommitConfigurationConnect",
		reflect.TypeOf((*CommitConfigurationConnect)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.CommitConfigurationConnectOnNewCert",
		reflect.TypeOf((*CommitConfigurationConnectOnNewCert)(nil)).Elem(),
		map[string]interface{}{
			"TRUST": CommitConfigurationConnectOnNewCert_TRUST,
			"DISTRUST": CommitConfigurationConnectOnNewCert_DISTRUST,
		},
	)
	_jsii_.RegisterEnum(
		"projen.polaris.CommitConfigurationConnectScm",
		reflect.TypeOf((*CommitConfigurationConnectScm)(nil)).Elem(),
		map[string]interface{}{
			"ADS": CommitConfigurationConnectScm_ADS,
			"CLEARCASE": CommitConfigurationConnectScm_CLEARCASE,
			"CVS": CommitConfigurationConnectScm_CVS,
			"GIT": CommitConfigurationConnectScm_GIT,
			"HG": CommitConfigurationConnectScm_HG,
			"PERFORCE": CommitConfigurationConnectScm_PERFORCE,
			"PLASTIC": CommitConfigurationConnectScm_PLASTIC,
			"PLASTIC_HYPHEN_DISTRIBUTED": CommitConfigurationConnectScm_PLASTIC_HYPHEN_DISTRIBUTED,
			"SVN": CommitConfigurationConnectScm_SVN,
			"TFS": CommitConfigurationConnectScm_TFS,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CommitConfigurationConnectTriage",
		reflect.TypeOf((*CommitConfigurationConnectTriage)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.CommitConfigurationConnectUploadArtifacts",
		reflect.TypeOf((*CommitConfigurationConnectUploadArtifacts)(nil)).Elem(),
		map[string]interface{}{
			"ALL": CommitConfigurationConnectUploadArtifacts_ALL,
			"LOGS_ONLY": CommitConfigurationConnectUploadArtifacts_LOGS_ONLY,
			"NONE": CommitConfigurationConnectUploadArtifacts_NONE,
			"ON_FAILURE": CommitConfigurationConnectUploadArtifacts_ON_FAILURE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CommitConfigurationLocal",
		reflect.TypeOf((*CommitConfigurationLocal)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.CommitConfigurationLocalFormat",
		reflect.TypeOf((*CommitConfigurationLocalFormat)(nil)).Elem(),
		map[string]interface{}{
			"HTML": CommitConfigurationLocalFormat_HTML,
			"JSON": CommitConfigurationLocalFormat_JSON,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CommitConfigurationSrm",
		reflect.TypeOf((*CommitConfigurationSrm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CompilerConfiguration",
		reflect.TypeOf((*CompilerConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.CovTranslateConfiguration",
		reflect.TypeOf((*CovTranslateConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.DirectivesConfiguration",
		reflect.TypeOf((*DirectivesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.DirectivesConfigurationConfig",
		reflect.TypeOf((*DirectivesConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.DirectivesConfigurationConfigType",
		reflect.TypeOf((*DirectivesConfigurationConfigType)(nil)).Elem(),
		map[string]interface{}{
			"COVERITY_ANALYSIS_CONFIGURATION": DirectivesConfigurationConfigType_COVERITY_ANALYSIS_CONFIGURATION,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.FilesConfiguration",
		reflect.TypeOf((*FilesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.ImportScmConfiguration",
		reflect.TypeOf((*ImportScmConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.JobsConfiguration",
		reflect.TypeOf((*JobsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.LanguagesConfiguration",
		reflect.TypeOf((*LanguagesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.LanguagesConfigurationExclude",
		reflect.TypeOf((*LanguagesConfigurationExclude)(nil)).Elem(),
		map[string]interface{}{
			"APEX": LanguagesConfigurationExclude_APEX,
			"C_HYPHEN_FAMILY": LanguagesConfigurationExclude_C_HYPHEN_FAMILY,
			"CSHARP": LanguagesConfigurationExclude_CSHARP,
			"DART": LanguagesConfigurationExclude_DART,
			"GO": LanguagesConfigurationExclude_GO,
			"JAVA": LanguagesConfigurationExclude_JAVA,
			"JAVASCRIPT": LanguagesConfigurationExclude_JAVASCRIPT,
			"KOTLIN": LanguagesConfigurationExclude_KOTLIN,
			"PHP": LanguagesConfigurationExclude_PHP,
			"PYTHON": LanguagesConfigurationExclude_PYTHON,
			"RUBY": LanguagesConfigurationExclude_RUBY,
			"SWIFT": LanguagesConfigurationExclude_SWIFT,
			"VB": LanguagesConfigurationExclude_VB,
			"CONFIGURATION": LanguagesConfigurationExclude_CONFIGURATION,
		},
	)
	_jsii_.RegisterEnum(
		"projen.polaris.LanguagesConfigurationInclude",
		reflect.TypeOf((*LanguagesConfigurationInclude)(nil)).Elem(),
		map[string]interface{}{
			"APEX": LanguagesConfigurationInclude_APEX,
			"C_HYPHEN_FAMILY": LanguagesConfigurationInclude_C_HYPHEN_FAMILY,
			"CSHARP": LanguagesConfigurationInclude_CSHARP,
			"DART": LanguagesConfigurationInclude_DART,
			"GO": LanguagesConfigurationInclude_GO,
			"JAVA": LanguagesConfigurationInclude_JAVA,
			"JAVASCRIPT": LanguagesConfigurationInclude_JAVASCRIPT,
			"KOTLIN": LanguagesConfigurationInclude_KOTLIN,
			"PHP": LanguagesConfigurationInclude_PHP,
			"PYTHON": LanguagesConfigurationInclude_PYTHON,
			"RUBY": LanguagesConfigurationInclude_RUBY,
			"SWIFT": LanguagesConfigurationInclude_SWIFT,
			"VB": LanguagesConfigurationInclude_VB,
			"CONFIGURATION": LanguagesConfigurationInclude_CONFIGURATION,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.ParallelTranslateConfiguration",
		reflect.TypeOf((*ParallelTranslateConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.ParseWarningsConfiguration",
		reflect.TypeOf((*ParseWarningsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.polaris.PolarisCoverity",
		reflect.TypeOf((*PolarisCoverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postProjectCreation", GoMethod: "PostProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "projectCreation", GoMethod: "ProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolarisCoverity{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.PolarisCoverityGoOptions",
		reflect.TypeOf((*PolarisCoverityGoOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.PolarisCoverityJavaOptions",
		reflect.TypeOf((*PolarisCoverityJavaOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.PolarisCoverityJavascriptOptions",
		reflect.TypeOf((*PolarisCoverityJavascriptOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.PolarisCoverityOptions",
		reflect.TypeOf((*PolarisCoverityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.PolarisCoveritySchema",
		reflect.TypeOf((*PolarisCoveritySchema)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.polaris.PolarisGoCoverity",
		reflect.TypeOf((*PolarisGoCoverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postProjectCreation", GoMethod: "PostProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "projectCreation", GoMethod: "ProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolarisGoCoverity{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PolarisCoverity)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.polaris.PolarisJavaCoverity",
		reflect.TypeOf((*PolarisJavaCoverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postProjectCreation", GoMethod: "PostProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "projectCreation", GoMethod: "ProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolarisJavaCoverity{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PolarisCoverity)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.polaris.PolarisJavascriptCoverity",
		reflect.TypeOf((*PolarisJavascriptCoverity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postProjectCreation", GoMethod: "PostProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "projectCreation", GoMethod: "ProjectCreation"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PolarisJavascriptCoverity{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PolarisCoverity)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.ResolvedCodingStandardConfiguration",
		reflect.TypeOf((*ResolvedCodingStandardConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.SigmaConfiguration",
		reflect.TypeOf((*SigmaConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.polaris.SigmaConfigurationEnableCheckSet",
		reflect.TypeOf((*SigmaConfigurationEnableCheckSet)(nil)).Elem(),
		map[string]interface{}{
			"ALL": SigmaConfigurationEnableCheckSet_ALL,
			"CIS": SigmaConfigurationEnableCheckSet_CIS,
			"DEFAULT": SigmaConfigurationEnableCheckSet_DEFAULT,
			"EMPTY": SigmaConfigurationEnableCheckSet_EMPTY,
		},
	)
	_jsii_.RegisterStruct(
		"projen.polaris.SnapshotConfiguration",
		reflect.TypeOf((*SnapshotConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.SpecificCodingStandardConfiguration",
		reflect.TypeOf((*SpecificCodingStandardConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.polaris.WebappArchiveConfiguration",
		reflect.TypeOf((*WebappArchiveConfiguration)(nil)).Elem(),
	)
}
