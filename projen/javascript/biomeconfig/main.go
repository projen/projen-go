package biomeconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.Actions",
		reflect.TypeOf((*Actions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.ArrowParentheses",
		reflect.TypeOf((*ArrowParentheses)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ArrowParentheses_ALWAYS,
			"AS_NEEDED": ArrowParentheses_AS_NEEDED,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.AssistConfiguration",
		reflect.TypeOf((*AssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.AttributePosition",
		reflect.TypeOf((*AttributePosition)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": AttributePosition_AUTO,
			"MULTILINE": AttributePosition_MULTILINE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.BiomeConfiguration",
		reflect.TypeOf((*BiomeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.CssAssistConfiguration",
		reflect.TypeOf((*CssAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.CssConfiguration",
		reflect.TypeOf((*CssConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.CssFormatterConfiguration",
		reflect.TypeOf((*CssFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.CssLinterConfiguration",
		reflect.TypeOf((*CssLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.CssParserConfiguration",
		reflect.TypeOf((*CssParserConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.Expand",
		reflect.TypeOf((*Expand)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": Expand_AUTO,
			"ALWAYS": Expand_ALWAYS,
			"NEVER": Expand_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.FilesConfiguration",
		reflect.TypeOf((*FilesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.FormatterConfiguration",
		reflect.TypeOf((*FormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GraphqlAssistConfiguration",
		reflect.TypeOf((*GraphqlAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GraphqlConfiguration",
		reflect.TypeOf((*GraphqlConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GraphqlFormatterConfiguration",
		reflect.TypeOf((*GraphqlFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GraphqlLinterConfiguration",
		reflect.TypeOf((*GraphqlLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GritAssistConfiguration",
		reflect.TypeOf((*GritAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GritConfiguration",
		reflect.TypeOf((*GritConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GritFormatterConfiguration",
		reflect.TypeOf((*GritFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.GritLinterConfiguration",
		reflect.TypeOf((*GritLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.HtmlAssistConfiguration",
		reflect.TypeOf((*HtmlAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.HtmlConfiguration",
		reflect.TypeOf((*HtmlConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.HtmlFormatterConfiguration",
		reflect.TypeOf((*HtmlFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.HtmlLinterConfiguration",
		reflect.TypeOf((*HtmlLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.HtmlParserConfiguration",
		reflect.TypeOf((*HtmlParserConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.IndentStyle",
		reflect.TypeOf((*IndentStyle)(nil)).Elem(),
		map[string]interface{}{
			"TAB": IndentStyle_TAB,
			"SPACE": IndentStyle_SPACE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsAssistConfiguration",
		reflect.TypeOf((*JsAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsConfiguration",
		reflect.TypeOf((*JsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsFormatterConfiguration",
		reflect.TypeOf((*JsFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsLinterConfiguration",
		reflect.TypeOf((*JsLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsParserConfiguration",
		reflect.TypeOf((*JsParserConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.JsTrailingCommas",
		reflect.TypeOf((*JsTrailingCommas)(nil)).Elem(),
		map[string]interface{}{
			"ALL": JsTrailingCommas_ALL,
			"ES5": JsTrailingCommas_ES5,
			"NONE": JsTrailingCommas_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsonAssistConfiguration",
		reflect.TypeOf((*JsonAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsonConfiguration",
		reflect.TypeOf((*JsonConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsonFormatterConfiguration",
		reflect.TypeOf((*JsonFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsonLinterConfiguration",
		reflect.TypeOf((*JsonLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.JsonParserConfiguration",
		reflect.TypeOf((*JsonParserConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.JsonTrailingCommas",
		reflect.TypeOf((*JsonTrailingCommas)(nil)).Elem(),
		map[string]interface{}{
			"NONE": JsonTrailingCommas_NONE,
			"ALL": JsonTrailingCommas_ALL,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.JsxRuntime",
		reflect.TypeOf((*JsxRuntime)(nil)).Elem(),
		map[string]interface{}{
			"TRANSPARENT": JsxRuntime_TRANSPARENT,
			"REACT_CLASSIC": JsxRuntime_REACT_CLASSIC,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.LineEnding",
		reflect.TypeOf((*LineEnding)(nil)).Elem(),
		map[string]interface{}{
			"LF": LineEnding_LF,
			"CRLF": LineEnding_CRLF,
			"CR": LineEnding_CR,
			"AUTO": LineEnding_AUTO,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.LinterConfiguration",
		reflect.TypeOf((*LinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.OperatorLinebreak",
		reflect.TypeOf((*OperatorLinebreak)(nil)).Elem(),
		map[string]interface{}{
			"AFTER": OperatorLinebreak_AFTER,
			"BEFORE": OperatorLinebreak_BEFORE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.OverrideAssistConfiguration",
		reflect.TypeOf((*OverrideAssistConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.OverrideFilesConfiguration",
		reflect.TypeOf((*OverrideFilesConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.OverrideFormatterConfiguration",
		reflect.TypeOf((*OverrideFormatterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.OverrideLinterConfiguration",
		reflect.TypeOf((*OverrideLinterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.OverridePattern",
		reflect.TypeOf((*OverridePattern)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.QuoteProperties",
		reflect.TypeOf((*QuoteProperties)(nil)).Elem(),
		map[string]interface{}{
			"AS_NEEDED": QuoteProperties_AS_NEEDED,
			"PRESERVE": QuoteProperties_PRESERVE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.QuoteStyle",
		reflect.TypeOf((*QuoteStyle)(nil)).Elem(),
		map[string]interface{}{
			"DOUBLE": QuoteStyle_DOUBLE,
			"SINGLE": QuoteStyle_SINGLE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.RuleDomainValue",
		reflect.TypeOf((*RuleDomainValue)(nil)).Elem(),
		map[string]interface{}{
			"ALL": RuleDomainValue_ALL,
			"NONE": RuleDomainValue_NONE,
			"RECOMMENDED": RuleDomainValue_RECOMMENDED,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.Rules",
		reflect.TypeOf((*Rules)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.SelfCloseVoidElements",
		reflect.TypeOf((*SelfCloseVoidElements)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": SelfCloseVoidElements_NEVER,
			"ALWAYS": SelfCloseVoidElements_ALWAYS,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.Semicolons",
		reflect.TypeOf((*Semicolons)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": Semicolons_ALWAYS,
			"AS_NEEDED": Semicolons_AS_NEEDED,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.Source",
		reflect.TypeOf((*Source)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.VcsClientKind",
		reflect.TypeOf((*VcsClientKind)(nil)).Elem(),
		map[string]interface{}{
			"GIT": VcsClientKind_GIT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.biome_config.VcsConfiguration",
		reflect.TypeOf((*VcsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.biome_config.WhitespaceSensitivity",
		reflect.TypeOf((*WhitespaceSensitivity)(nil)).Elem(),
		map[string]interface{}{
			"CSS": WhitespaceSensitivity_CSS,
			"STRICT": WhitespaceSensitivity_STRICT,
			"IGNORE": WhitespaceSensitivity_IGNORE,
		},
	)
}
