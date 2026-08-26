package sonarqube

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeCoverageOptions",
		reflect.TypeOf((*SonarqubeCoverageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeCpdOptions",
		reflect.TypeOf((*SonarqubeCpdOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeFileOptions",
		reflect.TypeOf((*SonarqubeFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeJavascriptOptions",
		reflect.TypeOf((*SonarqubeJavascriptOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.sonarqube.SonarqubeJavascriptProperties",
		reflect.TypeOf((*SonarqubeJavascriptProperties)(nil)).Elem(),
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
			j := jsiiProxy_SonarqubeJavascriptProperties{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SonarqubeProperties)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeJavascriptPropertiesOptions",
		reflect.TypeOf((*SonarqubeJavascriptPropertiesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeLcovOptions",
		reflect.TypeOf((*SonarqubeLcovOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.sonarqube.SonarqubeLogLevel",
		reflect.TypeOf((*SonarqubeLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"INFO": SonarqubeLogLevel_INFO,
			"DEBUG": SonarqubeLogLevel_DEBUG,
			"TRACE": SonarqubeLogLevel_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeLogOptions",
		reflect.TypeOf((*SonarqubeLogOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.sonarqube.SonarqubeProperties",
		reflect.TypeOf((*SonarqubeProperties)(nil)).Elem(),
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
			j := jsiiProxy_SonarqubeProperties{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubePropertiesOptions",
		reflect.TypeOf((*SonarqubePropertiesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeQualityGateOptions",
		reflect.TypeOf((*SonarqubeQualityGateOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.sonarqube.SonarqubeRegion",
		reflect.TypeOf((*SonarqubeRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": SonarqubeRegion_EU,
			"US": SonarqubeRegion_US,
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeRustClippyOptions",
		reflect.TypeOf((*SonarqubeRustClippyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeRustClippyReportOptions",
		reflect.TypeOf((*SonarqubeRustClippyReportOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeRustOptions",
		reflect.TypeOf((*SonarqubeRustOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.sonarqube.SonarqubeRustProperties",
		reflect.TypeOf((*SonarqubeRustProperties)(nil)).Elem(),
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
			j := jsiiProxy_SonarqubeRustProperties{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SonarqubeProperties)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeRustPropertiesOptions",
		reflect.TypeOf((*SonarqubeRustPropertiesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeScmExclusionsOptions",
		reflect.TypeOf((*SonarqubeScmExclusionsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeScmOptions",
		reflect.TypeOf((*SonarqubeScmOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeTypescriptOptions",
		reflect.TypeOf((*SonarqubeTypescriptOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.sonarqube.SonarqubeTypescriptProperties",
		reflect.TypeOf((*SonarqubeTypescriptProperties)(nil)).Elem(),
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
			j := jsiiProxy_SonarqubeTypescriptProperties{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SonarqubeProperties)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.sonarqube.SonarqubeTypescriptPropertiesOptions",
		reflect.TypeOf((*SonarqubeTypescriptPropertiesOptions)(nil)).Elem(),
	)
}
