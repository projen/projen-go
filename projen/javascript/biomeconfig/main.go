package biomeconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IA11y",
		reflect.TypeOf((*IA11y)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noAccessKey", GoGetter: "NoAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "noAriaHiddenOnFocusable", GoGetter: "NoAriaHiddenOnFocusable"},
			_jsii_.MemberProperty{JsiiProperty: "noAriaUnsupportedElements", GoGetter: "NoAriaUnsupportedElements"},
			_jsii_.MemberProperty{JsiiProperty: "noAutofocus", GoGetter: "NoAutofocus"},
			_jsii_.MemberProperty{JsiiProperty: "noBlankTarget", GoGetter: "NoBlankTarget"},
			_jsii_.MemberProperty{JsiiProperty: "noDistractingElements", GoGetter: "NoDistractingElements"},
			_jsii_.MemberProperty{JsiiProperty: "noHeaderScope", GoGetter: "NoHeaderScope"},
			_jsii_.MemberProperty{JsiiProperty: "noInteractiveElementToNoninteractiveRole", GoGetter: "NoInteractiveElementToNoninteractiveRole"},
			_jsii_.MemberProperty{JsiiProperty: "noLabelWithoutControl", GoGetter: "NoLabelWithoutControl"},
			_jsii_.MemberProperty{JsiiProperty: "noNoninteractiveElementToInteractiveRole", GoGetter: "NoNoninteractiveElementToInteractiveRole"},
			_jsii_.MemberProperty{JsiiProperty: "noNoninteractiveTabindex", GoGetter: "NoNoninteractiveTabindex"},
			_jsii_.MemberProperty{JsiiProperty: "noPositiveTabindex", GoGetter: "NoPositiveTabindex"},
			_jsii_.MemberProperty{JsiiProperty: "noRedundantAlt", GoGetter: "NoRedundantAlt"},
			_jsii_.MemberProperty{JsiiProperty: "noRedundantRoles", GoGetter: "NoRedundantRoles"},
			_jsii_.MemberProperty{JsiiProperty: "noSvgWithoutTitle", GoGetter: "NoSvgWithoutTitle"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useAltText", GoGetter: "UseAltText"},
			_jsii_.MemberProperty{JsiiProperty: "useAnchorContent", GoGetter: "UseAnchorContent"},
			_jsii_.MemberProperty{JsiiProperty: "useAriaActivedescendantWithTabindex", GoGetter: "UseAriaActivedescendantWithTabindex"},
			_jsii_.MemberProperty{JsiiProperty: "useAriaPropsForRole", GoGetter: "UseAriaPropsForRole"},
			_jsii_.MemberProperty{JsiiProperty: "useButtonType", GoGetter: "UseButtonType"},
			_jsii_.MemberProperty{JsiiProperty: "useFocusableInteractive", GoGetter: "UseFocusableInteractive"},
			_jsii_.MemberProperty{JsiiProperty: "useGenericFontNames", GoGetter: "UseGenericFontNames"},
			_jsii_.MemberProperty{JsiiProperty: "useHeadingContent", GoGetter: "UseHeadingContent"},
			_jsii_.MemberProperty{JsiiProperty: "useHtmlLang", GoGetter: "UseHtmlLang"},
			_jsii_.MemberProperty{JsiiProperty: "useIframeTitle", GoGetter: "UseIframeTitle"},
			_jsii_.MemberProperty{JsiiProperty: "useKeyWithClickEvents", GoGetter: "UseKeyWithClickEvents"},
			_jsii_.MemberProperty{JsiiProperty: "useKeyWithMouseEvents", GoGetter: "UseKeyWithMouseEvents"},
			_jsii_.MemberProperty{JsiiProperty: "useMediaCaption", GoGetter: "UseMediaCaption"},
			_jsii_.MemberProperty{JsiiProperty: "useSemanticElements", GoGetter: "UseSemanticElements"},
			_jsii_.MemberProperty{JsiiProperty: "useValidAnchor", GoGetter: "UseValidAnchor"},
			_jsii_.MemberProperty{JsiiProperty: "useValidAriaProps", GoGetter: "UseValidAriaProps"},
			_jsii_.MemberProperty{JsiiProperty: "useValidAriaRole", GoGetter: "UseValidAriaRole"},
			_jsii_.MemberProperty{JsiiProperty: "useValidAriaValues", GoGetter: "UseValidAriaValues"},
			_jsii_.MemberProperty{JsiiProperty: "useValidLang", GoGetter: "UseValidLang"},
		},
		func() interface{} {
			return &jsiiProxy_IA11y{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IActions",
		reflect.TypeOf((*IActions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
		},
		func() interface{} {
			return &jsiiProxy_IActions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IAllowDomainOptions",
		reflect.TypeOf((*IAllowDomainOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowDomains", GoGetter: "AllowDomains"},
		},
		func() interface{} {
			return &jsiiProxy_IAllowDomainOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IAssistsConfiguration",
		reflect.TypeOf((*IAssistsConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
		},
		func() interface{} {
			return &jsiiProxy_IAssistsConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IComplexity",
		reflect.TypeOf((*IComplexity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noBannedTypes", GoGetter: "NoBannedTypes"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyTypeParameters", GoGetter: "NoEmptyTypeParameters"},
			_jsii_.MemberProperty{JsiiProperty: "noExcessiveCognitiveComplexity", GoGetter: "NoExcessiveCognitiveComplexity"},
			_jsii_.MemberProperty{JsiiProperty: "noExcessiveNestedTestSuites", GoGetter: "NoExcessiveNestedTestSuites"},
			_jsii_.MemberProperty{JsiiProperty: "noExtraBooleanCast", GoGetter: "NoExtraBooleanCast"},
			_jsii_.MemberProperty{JsiiProperty: "noForEach", GoGetter: "NoForEach"},
			_jsii_.MemberProperty{JsiiProperty: "noMultipleSpacesInRegularExpressionLiterals", GoGetter: "NoMultipleSpacesInRegularExpressionLiterals"},
			_jsii_.MemberProperty{JsiiProperty: "noStaticOnlyClass", GoGetter: "NoStaticOnlyClass"},
			_jsii_.MemberProperty{JsiiProperty: "noThisInStatic", GoGetter: "NoThisInStatic"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessCatch", GoGetter: "NoUselessCatch"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessConstructor", GoGetter: "NoUselessConstructor"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessEmptyExport", GoGetter: "NoUselessEmptyExport"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessFragments", GoGetter: "NoUselessFragments"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessLabel", GoGetter: "NoUselessLabel"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessLoneBlockStatements", GoGetter: "NoUselessLoneBlockStatements"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessRename", GoGetter: "NoUselessRename"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessStringConcat", GoGetter: "NoUselessStringConcat"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessSwitchCase", GoGetter: "NoUselessSwitchCase"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessTernary", GoGetter: "NoUselessTernary"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessThisAlias", GoGetter: "NoUselessThisAlias"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessTypeConstraint", GoGetter: "NoUselessTypeConstraint"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessUndefinedInitialization", GoGetter: "NoUselessUndefinedInitialization"},
			_jsii_.MemberProperty{JsiiProperty: "noVoid", GoGetter: "NoVoid"},
			_jsii_.MemberProperty{JsiiProperty: "noWith", GoGetter: "NoWith"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useArrowFunction", GoGetter: "UseArrowFunction"},
			_jsii_.MemberProperty{JsiiProperty: "useDateNow", GoGetter: "UseDateNow"},
			_jsii_.MemberProperty{JsiiProperty: "useFlatMap", GoGetter: "UseFlatMap"},
			_jsii_.MemberProperty{JsiiProperty: "useLiteralKeys", GoGetter: "UseLiteralKeys"},
			_jsii_.MemberProperty{JsiiProperty: "useOptionalChain", GoGetter: "UseOptionalChain"},
			_jsii_.MemberProperty{JsiiProperty: "useRegexLiterals", GoGetter: "UseRegexLiterals"},
			_jsii_.MemberProperty{JsiiProperty: "useSimpleNumberKeys", GoGetter: "UseSimpleNumberKeys"},
			_jsii_.MemberProperty{JsiiProperty: "useSimplifiedLogicExpression", GoGetter: "UseSimplifiedLogicExpression"},
		},
		func() interface{} {
			return &jsiiProxy_IComplexity{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IComplexityOptions",
		reflect.TypeOf((*IComplexityOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "maxAllowedComplexity", GoGetter: "MaxAllowedComplexity"},
		},
		func() interface{} {
			return &jsiiProxy_IComplexityOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IConfiguration",
		reflect.TypeOf((*IConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assists", GoGetter: "Assists"},
			_jsii_.MemberProperty{JsiiProperty: "css", GoGetter: "Css"},
			_jsii_.MemberProperty{JsiiProperty: "extends", GoGetter: "Extends"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "graphql", GoGetter: "Graphql"},
			_jsii_.MemberProperty{JsiiProperty: "javascript", GoGetter: "Javascript"},
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
			_jsii_.MemberProperty{JsiiProperty: "organizeImports", GoGetter: "OrganizeImports"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberProperty{JsiiProperty: "vcs", GoGetter: "Vcs"},
		},
		func() interface{} {
			return &jsiiProxy_IConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IConsistentArrayTypeOptions",
		reflect.TypeOf((*IConsistentArrayTypeOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "syntax", GoGetter: "Syntax"},
		},
		func() interface{} {
			return &jsiiProxy_IConsistentArrayTypeOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IConsistentMemberAccessibilityOptions",
		reflect.TypeOf((*IConsistentMemberAccessibilityOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessibility", GoGetter: "Accessibility"},
		},
		func() interface{} {
			return &jsiiProxy_IConsistentMemberAccessibilityOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IConvention",
		reflect.TypeOf((*IConvention)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "formats", GoGetter: "Formats"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "selector", GoGetter: "Selector"},
		},
		func() interface{} {
			return &jsiiProxy_IConvention{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICorrectness",
		reflect.TypeOf((*ICorrectness)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noChildrenProp", GoGetter: "NoChildrenProp"},
			_jsii_.MemberProperty{JsiiProperty: "noConstantCondition", GoGetter: "NoConstantCondition"},
			_jsii_.MemberProperty{JsiiProperty: "noConstantMathMinMaxClamp", GoGetter: "NoConstantMathMinMaxClamp"},
			_jsii_.MemberProperty{JsiiProperty: "noConstAssign", GoGetter: "NoConstAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noConstructorReturn", GoGetter: "NoConstructorReturn"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyCharacterClassInRegex", GoGetter: "NoEmptyCharacterClassInRegex"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyPattern", GoGetter: "NoEmptyPattern"},
			_jsii_.MemberProperty{JsiiProperty: "noFlatMapIdentity", GoGetter: "NoFlatMapIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "noGlobalObjectCalls", GoGetter: "NoGlobalObjectCalls"},
			_jsii_.MemberProperty{JsiiProperty: "noInnerDeclarations", GoGetter: "NoInnerDeclarations"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidBuiltinInstantiation", GoGetter: "NoInvalidBuiltinInstantiation"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidConstructorSuper", GoGetter: "NoInvalidConstructorSuper"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidDirectionInLinearGradient", GoGetter: "NoInvalidDirectionInLinearGradient"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidGridAreas", GoGetter: "NoInvalidGridAreas"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidNewBuiltin", GoGetter: "NoInvalidNewBuiltin"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidPositionAtImportRule", GoGetter: "NoInvalidPositionAtImportRule"},
			_jsii_.MemberProperty{JsiiProperty: "noInvalidUseBeforeDeclaration", GoGetter: "NoInvalidUseBeforeDeclaration"},
			_jsii_.MemberProperty{JsiiProperty: "noNewSymbol", GoGetter: "NoNewSymbol"},
			_jsii_.MemberProperty{JsiiProperty: "noNodejsModules", GoGetter: "NoNodejsModules"},
			_jsii_.MemberProperty{JsiiProperty: "noNonoctalDecimalEscape", GoGetter: "NoNonoctalDecimalEscape"},
			_jsii_.MemberProperty{JsiiProperty: "noPrecisionLoss", GoGetter: "NoPrecisionLoss"},
			_jsii_.MemberProperty{JsiiProperty: "noRenderReturnValue", GoGetter: "NoRenderReturnValue"},
			_jsii_.MemberProperty{JsiiProperty: "noSelfAssign", GoGetter: "NoSelfAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noSetterReturn", GoGetter: "NoSetterReturn"},
			_jsii_.MemberProperty{JsiiProperty: "noStringCaseMismatch", GoGetter: "NoStringCaseMismatch"},
			_jsii_.MemberProperty{JsiiProperty: "noSwitchDeclarations", GoGetter: "NoSwitchDeclarations"},
			_jsii_.MemberProperty{JsiiProperty: "noUndeclaredDependencies", GoGetter: "NoUndeclaredDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "noUndeclaredVariables", GoGetter: "NoUndeclaredVariables"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownFunction", GoGetter: "NoUnknownFunction"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownMediaFeatureName", GoGetter: "NoUnknownMediaFeatureName"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownProperty", GoGetter: "NoUnknownProperty"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownUnit", GoGetter: "NoUnknownUnit"},
			_jsii_.MemberProperty{JsiiProperty: "noUnmatchableAnbSelector", GoGetter: "NoUnmatchableAnbSelector"},
			_jsii_.MemberProperty{JsiiProperty: "noUnnecessaryContinue", GoGetter: "NoUnnecessaryContinue"},
			_jsii_.MemberProperty{JsiiProperty: "noUnreachable", GoGetter: "NoUnreachable"},
			_jsii_.MemberProperty{JsiiProperty: "noUnreachableSuper", GoGetter: "NoUnreachableSuper"},
			_jsii_.MemberProperty{JsiiProperty: "noUnsafeFinally", GoGetter: "NoUnsafeFinally"},
			_jsii_.MemberProperty{JsiiProperty: "noUnsafeOptionalChaining", GoGetter: "NoUnsafeOptionalChaining"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedFunctionParameters", GoGetter: "NoUnusedFunctionParameters"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedImports", GoGetter: "NoUnusedImports"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedLabels", GoGetter: "NoUnusedLabels"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedPrivateClassMembers", GoGetter: "NoUnusedPrivateClassMembers"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedVariables", GoGetter: "NoUnusedVariables"},
			_jsii_.MemberProperty{JsiiProperty: "noVoidElementsWithChildren", GoGetter: "NoVoidElementsWithChildren"},
			_jsii_.MemberProperty{JsiiProperty: "noVoidTypeReturn", GoGetter: "NoVoidTypeReturn"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useArrayLiterals", GoGetter: "UseArrayLiterals"},
			_jsii_.MemberProperty{JsiiProperty: "useExhaustiveDependencies", GoGetter: "UseExhaustiveDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "useHookAtTopLevel", GoGetter: "UseHookAtTopLevel"},
			_jsii_.MemberProperty{JsiiProperty: "useImportExtensions", GoGetter: "UseImportExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "useIsNan", GoGetter: "UseIsNan"},
			_jsii_.MemberProperty{JsiiProperty: "useJsxKeyInIterable", GoGetter: "UseJsxKeyInIterable"},
			_jsii_.MemberProperty{JsiiProperty: "useValidForDirection", GoGetter: "UseValidForDirection"},
			_jsii_.MemberProperty{JsiiProperty: "useYield", GoGetter: "UseYield"},
		},
		func() interface{} {
			return &jsiiProxy_ICorrectness{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICssAssists",
		reflect.TypeOf((*ICssAssists)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_ICssAssists{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICssConfiguration",
		reflect.TypeOf((*ICssConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assists", GoGetter: "Assists"},
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
		},
		func() interface{} {
			return &jsiiProxy_ICssConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICssFormatter",
		reflect.TypeOf((*ICssFormatter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "quoteStyle", GoGetter: "QuoteStyle"},
		},
		func() interface{} {
			return &jsiiProxy_ICssFormatter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICssLinter",
		reflect.TypeOf((*ICssLinter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_ICssLinter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICssParser",
		reflect.TypeOf((*ICssParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowWrongLineComments", GoGetter: "AllowWrongLineComments"},
			_jsii_.MemberProperty{JsiiProperty: "cssModules", GoGetter: "CssModules"},
		},
		func() interface{} {
			return &jsiiProxy_ICssParser{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ICustomRestrictedTypeOptions",
		reflect.TypeOf((*ICustomRestrictedTypeOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "use", GoGetter: "Use"},
		},
		func() interface{} {
			return &jsiiProxy_ICustomRestrictedTypeOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IDeprecatedHooksOptions",
		reflect.TypeOf((*IDeprecatedHooksOptions)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IDeprecatedHooksOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IFilenamingConventionOptions",
		reflect.TypeOf((*IFilenamingConventionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filenameCases", GoGetter: "FilenameCases"},
			_jsii_.MemberProperty{JsiiProperty: "requireAscii", GoGetter: "RequireAscii"},
			_jsii_.MemberProperty{JsiiProperty: "strictCase", GoGetter: "StrictCase"},
		},
		func() interface{} {
			return &jsiiProxy_IFilenamingConventionOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IFilesConfiguration",
		reflect.TypeOf((*IFilesConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreUnknown", GoGetter: "IgnoreUnknown"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "maxSize", GoGetter: "MaxSize"},
		},
		func() interface{} {
			return &jsiiProxy_IFilesConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IFormatterConfiguration",
		reflect.TypeOf((*IFormatterConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributePosition", GoGetter: "AttributePosition"},
			_jsii_.MemberProperty{JsiiProperty: "bracketSpacing", GoGetter: "BracketSpacing"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "formatWithErrors", GoGetter: "FormatWithErrors"},
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "indentSize", GoGetter: "IndentSize"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "useEditorconfig", GoGetter: "UseEditorconfig"},
		},
		func() interface{} {
			return &jsiiProxy_IFormatterConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IGraphqlConfiguration",
		reflect.TypeOf((*IGraphqlConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
		},
		func() interface{} {
			return &jsiiProxy_IGraphqlConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IGraphqlFormatter",
		reflect.TypeOf((*IGraphqlFormatter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bracketSpacing", GoGetter: "BracketSpacing"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "quoteStyle", GoGetter: "QuoteStyle"},
		},
		func() interface{} {
			return &jsiiProxy_IGraphqlFormatter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IGraphqlLinter",
		reflect.TypeOf((*IGraphqlLinter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IGraphqlLinter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IHook",
		reflect.TypeOf((*IHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "closureIndex", GoGetter: "ClosureIndex"},
			_jsii_.MemberProperty{JsiiProperty: "dependenciesIndex", GoGetter: "DependenciesIndex"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "stableResult", GoGetter: "StableResult"},
		},
		func() interface{} {
			return &jsiiProxy_IHook{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptAssists",
		reflect.TypeOf((*IJavascriptAssists)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IJavascriptAssists{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptConfiguration",
		reflect.TypeOf((*IJavascriptConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assists", GoGetter: "Assists"},
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "globals", GoGetter: "Globals"},
			_jsii_.MemberProperty{JsiiProperty: "jsxRuntime", GoGetter: "JsxRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
			_jsii_.MemberProperty{JsiiProperty: "organizeImports", GoGetter: "OrganizeImports"},
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
		},
		func() interface{} {
			return &jsiiProxy_IJavascriptConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptFormatter",
		reflect.TypeOf((*IJavascriptFormatter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arrowParentheses", GoGetter: "ArrowParentheses"},
			_jsii_.MemberProperty{JsiiProperty: "attributePosition", GoGetter: "AttributePosition"},
			_jsii_.MemberProperty{JsiiProperty: "bracketSameLine", GoGetter: "BracketSameLine"},
			_jsii_.MemberProperty{JsiiProperty: "bracketSpacing", GoGetter: "BracketSpacing"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "indentSize", GoGetter: "IndentSize"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "jsxQuoteStyle", GoGetter: "JsxQuoteStyle"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "quoteProperties", GoGetter: "QuoteProperties"},
			_jsii_.MemberProperty{JsiiProperty: "quoteStyle", GoGetter: "QuoteStyle"},
			_jsii_.MemberProperty{JsiiProperty: "semicolons", GoGetter: "Semicolons"},
			_jsii_.MemberProperty{JsiiProperty: "trailingComma", GoGetter: "TrailingComma"},
			_jsii_.MemberProperty{JsiiProperty: "trailingCommas", GoGetter: "TrailingCommas"},
		},
		func() interface{} {
			return &jsiiProxy_IJavascriptFormatter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptLinter",
		reflect.TypeOf((*IJavascriptLinter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IJavascriptLinter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptOrganizeImports",
		reflect.TypeOf((*IJavascriptOrganizeImports)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IJavascriptOrganizeImports{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJavascriptParser",
		reflect.TypeOf((*IJavascriptParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "unsafeParameterDecoratorsEnabled", GoGetter: "UnsafeParameterDecoratorsEnabled"},
		},
		func() interface{} {
			return &jsiiProxy_IJavascriptParser{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJsonAssists",
		reflect.TypeOf((*IJsonAssists)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonAssists{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJsonConfiguration",
		reflect.TypeOf((*IJsonConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assists", GoGetter: "Assists"},
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJsonFormatter",
		reflect.TypeOf((*IJsonFormatter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "indentSize", GoGetter: "IndentSize"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
			_jsii_.MemberProperty{JsiiProperty: "trailingCommas", GoGetter: "TrailingCommas"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonFormatter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJsonLinter",
		reflect.TypeOf((*IJsonLinter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonLinter{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IJsonParser",
		reflect.TypeOf((*IJsonParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowComments", GoGetter: "AllowComments"},
			_jsii_.MemberProperty{JsiiProperty: "allowTrailingCommas", GoGetter: "AllowTrailingCommas"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonParser{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ILinterConfiguration",
		reflect.TypeOf((*ILinterConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			return &jsiiProxy_ILinterConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INamingConventionOptions",
		reflect.TypeOf((*INamingConventionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "conventions", GoGetter: "Conventions"},
			_jsii_.MemberProperty{JsiiProperty: "enumMemberCase", GoGetter: "EnumMemberCase"},
			_jsii_.MemberProperty{JsiiProperty: "requireAscii", GoGetter: "RequireAscii"},
			_jsii_.MemberProperty{JsiiProperty: "strictCase", GoGetter: "StrictCase"},
		},
		func() interface{} {
			return &jsiiProxy_INamingConventionOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INoConsoleOptions",
		reflect.TypeOf((*INoConsoleOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allow", GoGetter: "Allow"},
		},
		func() interface{} {
			return &jsiiProxy_INoConsoleOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INoDoubleEqualsOptions",
		reflect.TypeOf((*INoDoubleEqualsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ignoreNull", GoGetter: "IgnoreNull"},
		},
		func() interface{} {
			return &jsiiProxy_INoDoubleEqualsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INoLabelWithoutControlOptions",
		reflect.TypeOf((*INoLabelWithoutControlOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inputComponents", GoGetter: "InputComponents"},
			_jsii_.MemberProperty{JsiiProperty: "labelAttributes", GoGetter: "LabelAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "labelComponents", GoGetter: "LabelComponents"},
		},
		func() interface{} {
			return &jsiiProxy_INoLabelWithoutControlOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INoRestrictedTypesOptions",
		reflect.TypeOf((*INoRestrictedTypesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
		},
		func() interface{} {
			return &jsiiProxy_INoRestrictedTypesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INoSecretsOptions",
		reflect.TypeOf((*INoSecretsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "entropyThreshold", GoGetter: "EntropyThreshold"},
		},
		func() interface{} {
			return &jsiiProxy_INoSecretsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.INursery",
		reflect.TypeOf((*INursery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noCommonJs", GoGetter: "NoCommonJs"},
			_jsii_.MemberProperty{JsiiProperty: "noDescendingSpecificity", GoGetter: "NoDescendingSpecificity"},
			_jsii_.MemberProperty{JsiiProperty: "noDocumentCookie", GoGetter: "NoDocumentCookie"},
			_jsii_.MemberProperty{JsiiProperty: "noDocumentImportInPage", GoGetter: "NoDocumentImportInPage"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateCustomProperties", GoGetter: "NoDuplicateCustomProperties"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicatedFields", GoGetter: "NoDuplicatedFields"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateElseIf", GoGetter: "NoDuplicateElseIf"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateProperties", GoGetter: "NoDuplicateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "noDynamicNamespaceImportAccess", GoGetter: "NoDynamicNamespaceImportAccess"},
			_jsii_.MemberProperty{JsiiProperty: "noEnum", GoGetter: "NoEnum"},
			_jsii_.MemberProperty{JsiiProperty: "noExportedImports", GoGetter: "NoExportedImports"},
			_jsii_.MemberProperty{JsiiProperty: "noHeadElement", GoGetter: "NoHeadElement"},
			_jsii_.MemberProperty{JsiiProperty: "noHeadImportInDocument", GoGetter: "NoHeadImportInDocument"},
			_jsii_.MemberProperty{JsiiProperty: "noImgElement", GoGetter: "NoImgElement"},
			_jsii_.MemberProperty{JsiiProperty: "noIrregularWhitespace", GoGetter: "NoIrregularWhitespace"},
			_jsii_.MemberProperty{JsiiProperty: "noMissingVarFunction", GoGetter: "NoMissingVarFunction"},
			_jsii_.MemberProperty{JsiiProperty: "noNestedTernary", GoGetter: "NoNestedTernary"},
			_jsii_.MemberProperty{JsiiProperty: "noOctalEscape", GoGetter: "NoOctalEscape"},
			_jsii_.MemberProperty{JsiiProperty: "noProcessEnv", GoGetter: "NoProcessEnv"},
			_jsii_.MemberProperty{JsiiProperty: "noRestrictedImports", GoGetter: "NoRestrictedImports"},
			_jsii_.MemberProperty{JsiiProperty: "noRestrictedTypes", GoGetter: "NoRestrictedTypes"},
			_jsii_.MemberProperty{JsiiProperty: "noSecrets", GoGetter: "NoSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "noStaticElementInteractions", GoGetter: "NoStaticElementInteractions"},
			_jsii_.MemberProperty{JsiiProperty: "noSubstr", GoGetter: "NoSubstr"},
			_jsii_.MemberProperty{JsiiProperty: "noTemplateCurlyInString", GoGetter: "NoTemplateCurlyInString"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownPseudoClass", GoGetter: "NoUnknownPseudoClass"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownPseudoElement", GoGetter: "NoUnknownPseudoElement"},
			_jsii_.MemberProperty{JsiiProperty: "noUnknownTypeSelector", GoGetter: "NoUnknownTypeSelector"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessEscapeInRegex", GoGetter: "NoUselessEscapeInRegex"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessStringRaw", GoGetter: "NoUselessStringRaw"},
			_jsii_.MemberProperty{JsiiProperty: "noValueAtRule", GoGetter: "NoValueAtRule"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useAdjacentOverloadSignatures", GoGetter: "UseAdjacentOverloadSignatures"},
			_jsii_.MemberProperty{JsiiProperty: "useAriaPropsSupportedByRole", GoGetter: "UseAriaPropsSupportedByRole"},
			_jsii_.MemberProperty{JsiiProperty: "useAtIndex", GoGetter: "UseAtIndex"},
			_jsii_.MemberProperty{JsiiProperty: "useCollapsedIf", GoGetter: "UseCollapsedIf"},
			_jsii_.MemberProperty{JsiiProperty: "useComponentExportOnlyModules", GoGetter: "UseComponentExportOnlyModules"},
			_jsii_.MemberProperty{JsiiProperty: "useConsistentCurlyBraces", GoGetter: "UseConsistentCurlyBraces"},
			_jsii_.MemberProperty{JsiiProperty: "useConsistentMemberAccessibility", GoGetter: "UseConsistentMemberAccessibility"},
			_jsii_.MemberProperty{JsiiProperty: "useDeprecatedReason", GoGetter: "UseDeprecatedReason"},
			_jsii_.MemberProperty{JsiiProperty: "useExplicitType", GoGetter: "UseExplicitType"},
			_jsii_.MemberProperty{JsiiProperty: "useGoogleFontDisplay", GoGetter: "UseGoogleFontDisplay"},
			_jsii_.MemberProperty{JsiiProperty: "useGuardForIn", GoGetter: "UseGuardForIn"},
			_jsii_.MemberProperty{JsiiProperty: "useImportRestrictions", GoGetter: "UseImportRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "useSortedClasses", GoGetter: "UseSortedClasses"},
			_jsii_.MemberProperty{JsiiProperty: "useStrictMode", GoGetter: "UseStrictMode"},
			_jsii_.MemberProperty{JsiiProperty: "useTrimStartEnd", GoGetter: "UseTrimStartEnd"},
			_jsii_.MemberProperty{JsiiProperty: "useValidAutocomplete", GoGetter: "UseValidAutocomplete"},
		},
		func() interface{} {
			return &jsiiProxy_INursery{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IOrganizeImports",
		reflect.TypeOf((*IOrganizeImports)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
		},
		func() interface{} {
			return &jsiiProxy_IOrganizeImports{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IOverrideFormatterConfiguration",
		reflect.TypeOf((*IOverrideFormatterConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributePosition", GoGetter: "AttributePosition"},
			_jsii_.MemberProperty{JsiiProperty: "bracketSpacing", GoGetter: "BracketSpacing"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "formatWithErrors", GoGetter: "FormatWithErrors"},
			_jsii_.MemberProperty{JsiiProperty: "indentSize", GoGetter: "IndentSize"},
			_jsii_.MemberProperty{JsiiProperty: "indentStyle", GoGetter: "IndentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "indentWidth", GoGetter: "IndentWidth"},
			_jsii_.MemberProperty{JsiiProperty: "lineEnding", GoGetter: "LineEnding"},
			_jsii_.MemberProperty{JsiiProperty: "lineWidth", GoGetter: "LineWidth"},
		},
		func() interface{} {
			return &jsiiProxy_IOverrideFormatterConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IOverrideLinterConfiguration",
		reflect.TypeOf((*IOverrideLinterConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			return &jsiiProxy_IOverrideLinterConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IOverrideOrganizeImportsConfiguration",
		reflect.TypeOf((*IOverrideOrganizeImportsConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IOverrideOrganizeImportsConfiguration{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IOverridePattern",
		reflect.TypeOf((*IOverridePattern)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "css", GoGetter: "Css"},
			_jsii_.MemberProperty{JsiiProperty: "formatter", GoGetter: "Formatter"},
			_jsii_.MemberProperty{JsiiProperty: "graphql", GoGetter: "Graphql"},
			_jsii_.MemberProperty{JsiiProperty: "ignore", GoGetter: "Ignore"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "javascript", GoGetter: "Javascript"},
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "linter", GoGetter: "Linter"},
			_jsii_.MemberProperty{JsiiProperty: "organizeImports", GoGetter: "OrganizeImports"},
		},
		func() interface{} {
			return &jsiiProxy_IOverridePattern{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IPerformance",
		reflect.TypeOf((*IPerformance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noAccumulatingSpread", GoGetter: "NoAccumulatingSpread"},
			_jsii_.MemberProperty{JsiiProperty: "noBarrelFile", GoGetter: "NoBarrelFile"},
			_jsii_.MemberProperty{JsiiProperty: "noDelete", GoGetter: "NoDelete"},
			_jsii_.MemberProperty{JsiiProperty: "noReExportAll", GoGetter: "NoReExportAll"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useTopLevelRegex", GoGetter: "UseTopLevelRegex"},
		},
		func() interface{} {
			return &jsiiProxy_IPerformance{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRestrictedGlobalsOptions",
		reflect.TypeOf((*IRestrictedGlobalsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deniedGlobals", GoGetter: "DeniedGlobals"},
		},
		func() interface{} {
			return &jsiiProxy_IRestrictedGlobalsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRestrictedImportsOptions",
		reflect.TypeOf((*IRestrictedImportsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "paths", GoGetter: "Paths"},
		},
		func() interface{} {
			return &jsiiProxy_IRestrictedImportsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithAllowDomainOptions",
		reflect.TypeOf((*IRuleWithAllowDomainOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithAllowDomainOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithComplexityOptions",
		reflect.TypeOf((*IRuleWithComplexityOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithComplexityOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithConsistentArrayTypeOptions",
		reflect.TypeOf((*IRuleWithConsistentArrayTypeOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithConsistentArrayTypeOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithConsistentMemberAccessibilityOptions",
		reflect.TypeOf((*IRuleWithConsistentMemberAccessibilityOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithDeprecatedHooksOptions",
		reflect.TypeOf((*IRuleWithDeprecatedHooksOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithDeprecatedHooksOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithFilenamingConventionOptions",
		reflect.TypeOf((*IRuleWithFilenamingConventionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithFilenamingConventionOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithFixNoOptions",
		reflect.TypeOf((*IRuleWithFixNoOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithFixNoOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNamingConventionOptions",
		reflect.TypeOf((*IRuleWithNamingConventionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNamingConventionOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoConsoleOptions",
		reflect.TypeOf((*IRuleWithNoConsoleOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoConsoleOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoDoubleEqualsOptions",
		reflect.TypeOf((*IRuleWithNoDoubleEqualsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoDoubleEqualsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoLabelWithoutControlOptions",
		reflect.TypeOf((*IRuleWithNoLabelWithoutControlOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoLabelWithoutControlOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoOptions",
		reflect.TypeOf((*IRuleWithNoOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoRestrictedTypesOptions",
		reflect.TypeOf((*IRuleWithNoRestrictedTypesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoRestrictedTypesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithNoSecretsOptions",
		reflect.TypeOf((*IRuleWithNoSecretsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithNoSecretsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithRestrictedGlobalsOptions",
		reflect.TypeOf((*IRuleWithRestrictedGlobalsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithRestrictedGlobalsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithRestrictedImportsOptions",
		reflect.TypeOf((*IRuleWithRestrictedImportsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithRestrictedImportsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithUseComponentExportOnlyModulesOptions",
		reflect.TypeOf((*IRuleWithUseComponentExportOnlyModulesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithUseComponentExportOnlyModulesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithUseExhaustiveDependenciesOptions",
		reflect.TypeOf((*IRuleWithUseExhaustiveDependenciesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithUseImportExtensionsOptions",
		reflect.TypeOf((*IRuleWithUseImportExtensionsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithUseImportExtensionsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithUseValidAutocompleteOptions",
		reflect.TypeOf((*IRuleWithUseValidAutocompleteOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithUseValidAutocompleteOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithUtilityClassSortingOptions",
		reflect.TypeOf((*IRuleWithUtilityClassSortingOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithUtilityClassSortingOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRuleWithValidAriaRoleOptions",
		reflect.TypeOf((*IRuleWithValidAriaRoleOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fix", GoGetter: "Fix"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleWithValidAriaRoleOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IRules",
		reflect.TypeOf((*IRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "a11y", GoGetter: "A11y"},
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "complexity", GoGetter: "Complexity"},
			_jsii_.MemberProperty{JsiiProperty: "correctness", GoGetter: "Correctness"},
			_jsii_.MemberProperty{JsiiProperty: "nursery", GoGetter: "Nursery"},
			_jsii_.MemberProperty{JsiiProperty: "performance", GoGetter: "Performance"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "security", GoGetter: "Security"},
			_jsii_.MemberProperty{JsiiProperty: "style", GoGetter: "Style"},
			_jsii_.MemberProperty{JsiiProperty: "suspicious", GoGetter: "Suspicious"},
		},
		func() interface{} {
			return &jsiiProxy_IRules{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ISecurity",
		reflect.TypeOf((*ISecurity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noDangerouslySetInnerHtml", GoGetter: "NoDangerouslySetInnerHtml"},
			_jsii_.MemberProperty{JsiiProperty: "noDangerouslySetInnerHtmlWithChildren", GoGetter: "NoDangerouslySetInnerHtmlWithChildren"},
			_jsii_.MemberProperty{JsiiProperty: "noGlobalEval", GoGetter: "NoGlobalEval"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
		},
		func() interface{} {
			return &jsiiProxy_ISecurity{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ISelector",
		reflect.TypeOf((*ISelector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "modifiers", GoGetter: "Modifiers"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
		},
		func() interface{} {
			return &jsiiProxy_ISelector{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ISource",
		reflect.TypeOf((*ISource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "sortJsxProps", GoGetter: "SortJsxProps"},
			_jsii_.MemberProperty{JsiiProperty: "useSortedKeys", GoGetter: "UseSortedKeys"},
		},
		func() interface{} {
			return &jsiiProxy_ISource{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IStyle",
		reflect.TypeOf((*IStyle)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noArguments", GoGetter: "NoArguments"},
			_jsii_.MemberProperty{JsiiProperty: "noCommaOperator", GoGetter: "NoCommaOperator"},
			_jsii_.MemberProperty{JsiiProperty: "noDefaultExport", GoGetter: "NoDefaultExport"},
			_jsii_.MemberProperty{JsiiProperty: "noDoneCallback", GoGetter: "NoDoneCallback"},
			_jsii_.MemberProperty{JsiiProperty: "noImplicitBoolean", GoGetter: "NoImplicitBoolean"},
			_jsii_.MemberProperty{JsiiProperty: "noInferrableTypes", GoGetter: "NoInferrableTypes"},
			_jsii_.MemberProperty{JsiiProperty: "noNamespace", GoGetter: "NoNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "noNamespaceImport", GoGetter: "NoNamespaceImport"},
			_jsii_.MemberProperty{JsiiProperty: "noNegationElse", GoGetter: "NoNegationElse"},
			_jsii_.MemberProperty{JsiiProperty: "noNonNullAssertion", GoGetter: "NoNonNullAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "noParameterAssign", GoGetter: "NoParameterAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noParameterProperties", GoGetter: "NoParameterProperties"},
			_jsii_.MemberProperty{JsiiProperty: "noRestrictedGlobals", GoGetter: "NoRestrictedGlobals"},
			_jsii_.MemberProperty{JsiiProperty: "noShoutyConstants", GoGetter: "NoShoutyConstants"},
			_jsii_.MemberProperty{JsiiProperty: "noUnusedTemplateLiteral", GoGetter: "NoUnusedTemplateLiteral"},
			_jsii_.MemberProperty{JsiiProperty: "noUselessElse", GoGetter: "NoUselessElse"},
			_jsii_.MemberProperty{JsiiProperty: "noVar", GoGetter: "NoVar"},
			_jsii_.MemberProperty{JsiiProperty: "noYodaExpression", GoGetter: "NoYodaExpression"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useAsConstAssertion", GoGetter: "UseAsConstAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "useBlockStatements", GoGetter: "UseBlockStatements"},
			_jsii_.MemberProperty{JsiiProperty: "useCollapsedElseIf", GoGetter: "UseCollapsedElseIf"},
			_jsii_.MemberProperty{JsiiProperty: "useConsistentArrayType", GoGetter: "UseConsistentArrayType"},
			_jsii_.MemberProperty{JsiiProperty: "useConsistentBuiltinInstantiation", GoGetter: "UseConsistentBuiltinInstantiation"},
			_jsii_.MemberProperty{JsiiProperty: "useConst", GoGetter: "UseConst"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultParameterLast", GoGetter: "UseDefaultParameterLast"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultSwitchClause", GoGetter: "UseDefaultSwitchClause"},
			_jsii_.MemberProperty{JsiiProperty: "useEnumInitializers", GoGetter: "UseEnumInitializers"},
			_jsii_.MemberProperty{JsiiProperty: "useExplicitLengthCheck", GoGetter: "UseExplicitLengthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "useExponentiationOperator", GoGetter: "UseExponentiationOperator"},
			_jsii_.MemberProperty{JsiiProperty: "useExportType", GoGetter: "UseExportType"},
			_jsii_.MemberProperty{JsiiProperty: "useFilenamingConvention", GoGetter: "UseFilenamingConvention"},
			_jsii_.MemberProperty{JsiiProperty: "useForOf", GoGetter: "UseForOf"},
			_jsii_.MemberProperty{JsiiProperty: "useFragmentSyntax", GoGetter: "UseFragmentSyntax"},
			_jsii_.MemberProperty{JsiiProperty: "useImportType", GoGetter: "UseImportType"},
			_jsii_.MemberProperty{JsiiProperty: "useLiteralEnumMembers", GoGetter: "UseLiteralEnumMembers"},
			_jsii_.MemberProperty{JsiiProperty: "useNamingConvention", GoGetter: "UseNamingConvention"},
			_jsii_.MemberProperty{JsiiProperty: "useNodeAssertStrict", GoGetter: "UseNodeAssertStrict"},
			_jsii_.MemberProperty{JsiiProperty: "useNodejsImportProtocol", GoGetter: "UseNodejsImportProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "useNumberNamespace", GoGetter: "UseNumberNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "useNumericLiterals", GoGetter: "UseNumericLiterals"},
			_jsii_.MemberProperty{JsiiProperty: "useSelfClosingElements", GoGetter: "UseSelfClosingElements"},
			_jsii_.MemberProperty{JsiiProperty: "useShorthandArrayType", GoGetter: "UseShorthandArrayType"},
			_jsii_.MemberProperty{JsiiProperty: "useShorthandAssign", GoGetter: "UseShorthandAssign"},
			_jsii_.MemberProperty{JsiiProperty: "useShorthandFunctionType", GoGetter: "UseShorthandFunctionType"},
			_jsii_.MemberProperty{JsiiProperty: "useSingleCaseStatement", GoGetter: "UseSingleCaseStatement"},
			_jsii_.MemberProperty{JsiiProperty: "useSingleVarDeclarator", GoGetter: "UseSingleVarDeclarator"},
			_jsii_.MemberProperty{JsiiProperty: "useTemplate", GoGetter: "UseTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "useThrowNewError", GoGetter: "UseThrowNewError"},
			_jsii_.MemberProperty{JsiiProperty: "useThrowOnlyError", GoGetter: "UseThrowOnlyError"},
			_jsii_.MemberProperty{JsiiProperty: "useWhile", GoGetter: "UseWhile"},
		},
		func() interface{} {
			return &jsiiProxy_IStyle{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ISuggestedExtensionMapping",
		reflect.TypeOf((*ISuggestedExtensionMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "component", GoGetter: "Component"},
			_jsii_.MemberProperty{JsiiProperty: "module", GoGetter: "Module"},
		},
		func() interface{} {
			return &jsiiProxy_ISuggestedExtensionMapping{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.ISuspicious",
		reflect.TypeOf((*ISuspicious)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
			_jsii_.MemberProperty{JsiiProperty: "noApproximativeNumericConstant", GoGetter: "NoApproximativeNumericConstant"},
			_jsii_.MemberProperty{JsiiProperty: "noArrayIndexKey", GoGetter: "NoArrayIndexKey"},
			_jsii_.MemberProperty{JsiiProperty: "noAssignInExpressions", GoGetter: "NoAssignInExpressions"},
			_jsii_.MemberProperty{JsiiProperty: "noAsyncPromiseExecutor", GoGetter: "NoAsyncPromiseExecutor"},
			_jsii_.MemberProperty{JsiiProperty: "noCatchAssign", GoGetter: "NoCatchAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noClassAssign", GoGetter: "NoClassAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noCommentText", GoGetter: "NoCommentText"},
			_jsii_.MemberProperty{JsiiProperty: "noCompareNegZero", GoGetter: "NoCompareNegZero"},
			_jsii_.MemberProperty{JsiiProperty: "noConfusingLabels", GoGetter: "NoConfusingLabels"},
			_jsii_.MemberProperty{JsiiProperty: "noConfusingVoidType", GoGetter: "NoConfusingVoidType"},
			_jsii_.MemberProperty{JsiiProperty: "noConsole", GoGetter: "NoConsole"},
			_jsii_.MemberProperty{JsiiProperty: "noConsoleLog", GoGetter: "NoConsoleLog"},
			_jsii_.MemberProperty{JsiiProperty: "noConstEnum", GoGetter: "NoConstEnum"},
			_jsii_.MemberProperty{JsiiProperty: "noControlCharactersInRegex", GoGetter: "NoControlCharactersInRegex"},
			_jsii_.MemberProperty{JsiiProperty: "noDebugger", GoGetter: "NoDebugger"},
			_jsii_.MemberProperty{JsiiProperty: "noDoubleEquals", GoGetter: "NoDoubleEquals"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateAtImportRules", GoGetter: "NoDuplicateAtImportRules"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateCase", GoGetter: "NoDuplicateCase"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateClassMembers", GoGetter: "NoDuplicateClassMembers"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateFontNames", GoGetter: "NoDuplicateFontNames"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateJsxProps", GoGetter: "NoDuplicateJsxProps"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateObjectKeys", GoGetter: "NoDuplicateObjectKeys"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateParameters", GoGetter: "NoDuplicateParameters"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateSelectorsKeyframeBlock", GoGetter: "NoDuplicateSelectorsKeyframeBlock"},
			_jsii_.MemberProperty{JsiiProperty: "noDuplicateTestHooks", GoGetter: "NoDuplicateTestHooks"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyBlock", GoGetter: "NoEmptyBlock"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyBlockStatements", GoGetter: "NoEmptyBlockStatements"},
			_jsii_.MemberProperty{JsiiProperty: "noEmptyInterface", GoGetter: "NoEmptyInterface"},
			_jsii_.MemberProperty{JsiiProperty: "noEvolvingTypes", GoGetter: "NoEvolvingTypes"},
			_jsii_.MemberProperty{JsiiProperty: "noExplicitAny", GoGetter: "NoExplicitAny"},
			_jsii_.MemberProperty{JsiiProperty: "noExportsInTest", GoGetter: "NoExportsInTest"},
			_jsii_.MemberProperty{JsiiProperty: "noExtraNonNullAssertion", GoGetter: "NoExtraNonNullAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "noFallthroughSwitchClause", GoGetter: "NoFallthroughSwitchClause"},
			_jsii_.MemberProperty{JsiiProperty: "noFocusedTests", GoGetter: "NoFocusedTests"},
			_jsii_.MemberProperty{JsiiProperty: "noFunctionAssign", GoGetter: "NoFunctionAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noGlobalAssign", GoGetter: "NoGlobalAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noGlobalIsFinite", GoGetter: "NoGlobalIsFinite"},
			_jsii_.MemberProperty{JsiiProperty: "noGlobalIsNan", GoGetter: "NoGlobalIsNan"},
			_jsii_.MemberProperty{JsiiProperty: "noImplicitAnyLet", GoGetter: "NoImplicitAnyLet"},
			_jsii_.MemberProperty{JsiiProperty: "noImportantInKeyframe", GoGetter: "NoImportantInKeyframe"},
			_jsii_.MemberProperty{JsiiProperty: "noImportAssign", GoGetter: "NoImportAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noLabelVar", GoGetter: "NoLabelVar"},
			_jsii_.MemberProperty{JsiiProperty: "noMisleadingCharacterClass", GoGetter: "NoMisleadingCharacterClass"},
			_jsii_.MemberProperty{JsiiProperty: "noMisleadingInstantiator", GoGetter: "NoMisleadingInstantiator"},
			_jsii_.MemberProperty{JsiiProperty: "noMisplacedAssertion", GoGetter: "NoMisplacedAssertion"},
			_jsii_.MemberProperty{JsiiProperty: "noMisrefactoredShorthandAssign", GoGetter: "NoMisrefactoredShorthandAssign"},
			_jsii_.MemberProperty{JsiiProperty: "noPrototypeBuiltins", GoGetter: "NoPrototypeBuiltins"},
			_jsii_.MemberProperty{JsiiProperty: "noReactSpecificProps", GoGetter: "NoReactSpecificProps"},
			_jsii_.MemberProperty{JsiiProperty: "noRedeclare", GoGetter: "NoRedeclare"},
			_jsii_.MemberProperty{JsiiProperty: "noRedundantUseStrict", GoGetter: "NoRedundantUseStrict"},
			_jsii_.MemberProperty{JsiiProperty: "noSelfCompare", GoGetter: "NoSelfCompare"},
			_jsii_.MemberProperty{JsiiProperty: "noShadowRestrictedNames", GoGetter: "NoShadowRestrictedNames"},
			_jsii_.MemberProperty{JsiiProperty: "noShorthandPropertyOverrides", GoGetter: "NoShorthandPropertyOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "noSkippedTests", GoGetter: "NoSkippedTests"},
			_jsii_.MemberProperty{JsiiProperty: "noSparseArray", GoGetter: "NoSparseArray"},
			_jsii_.MemberProperty{JsiiProperty: "noSuspiciousSemicolonInJsx", GoGetter: "NoSuspiciousSemicolonInJsx"},
			_jsii_.MemberProperty{JsiiProperty: "noThenProperty", GoGetter: "NoThenProperty"},
			_jsii_.MemberProperty{JsiiProperty: "noUnsafeDeclarationMerging", GoGetter: "NoUnsafeDeclarationMerging"},
			_jsii_.MemberProperty{JsiiProperty: "noUnsafeNegation", GoGetter: "NoUnsafeNegation"},
			_jsii_.MemberProperty{JsiiProperty: "recommended", GoGetter: "Recommended"},
			_jsii_.MemberProperty{JsiiProperty: "useAwait", GoGetter: "UseAwait"},
			_jsii_.MemberProperty{JsiiProperty: "useDefaultSwitchClauseLast", GoGetter: "UseDefaultSwitchClauseLast"},
			_jsii_.MemberProperty{JsiiProperty: "useErrorMessage", GoGetter: "UseErrorMessage"},
			_jsii_.MemberProperty{JsiiProperty: "useGetterReturn", GoGetter: "UseGetterReturn"},
			_jsii_.MemberProperty{JsiiProperty: "useIsArray", GoGetter: "UseIsArray"},
			_jsii_.MemberProperty{JsiiProperty: "useNamespaceKeyword", GoGetter: "UseNamespaceKeyword"},
			_jsii_.MemberProperty{JsiiProperty: "useNumberToFixedDigitsArgument", GoGetter: "UseNumberToFixedDigitsArgument"},
			_jsii_.MemberProperty{JsiiProperty: "useValidTypeof", GoGetter: "UseValidTypeof"},
		},
		func() interface{} {
			return &jsiiProxy_ISuspicious{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IUseComponentExportOnlyModulesOptions",
		reflect.TypeOf((*IUseComponentExportOnlyModulesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowConstantExport", GoGetter: "AllowConstantExport"},
			_jsii_.MemberProperty{JsiiProperty: "allowExportNames", GoGetter: "AllowExportNames"},
		},
		func() interface{} {
			return &jsiiProxy_IUseComponentExportOnlyModulesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IUseExhaustiveDependenciesOptions",
		reflect.TypeOf((*IUseExhaustiveDependenciesOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hooks", GoGetter: "Hooks"},
			_jsii_.MemberProperty{JsiiProperty: "reportMissingDependenciesArray", GoGetter: "ReportMissingDependenciesArray"},
			_jsii_.MemberProperty{JsiiProperty: "reportUnnecessaryDependencies", GoGetter: "ReportUnnecessaryDependencies"},
		},
		func() interface{} {
			return &jsiiProxy_IUseExhaustiveDependenciesOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IUseImportExtensionsOptions",
		reflect.TypeOf((*IUseImportExtensionsOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "suggestedExtensions", GoGetter: "SuggestedExtensions"},
		},
		func() interface{} {
			return &jsiiProxy_IUseImportExtensionsOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IUseValidAutocompleteOptions",
		reflect.TypeOf((*IUseValidAutocompleteOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inputComponents", GoGetter: "InputComponents"},
		},
		func() interface{} {
			return &jsiiProxy_IUseValidAutocompleteOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IUtilityClassSortingOptions",
		reflect.TypeOf((*IUtilityClassSortingOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "functions", GoGetter: "Functions"},
		},
		func() interface{} {
			return &jsiiProxy_IUtilityClassSortingOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IValidAriaRoleOptions",
		reflect.TypeOf((*IValidAriaRoleOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowInvalidRoles", GoGetter: "AllowInvalidRoles"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreNonDom", GoGetter: "IgnoreNonDom"},
		},
		func() interface{} {
			return &jsiiProxy_IValidAriaRoleOptions{}
		},
	)
	_jsii_.RegisterInterface(
		"projen.javascript.biome_config.IVcsConfiguration",
		reflect.TypeOf((*IVcsConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientKind", GoGetter: "ClientKind"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranch", GoGetter: "DefaultBranch"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "useIgnoreFile", GoGetter: "UseIgnoreFile"},
		},
		func() interface{} {
			return &jsiiProxy_IVcsConfiguration{}
		},
	)
}
