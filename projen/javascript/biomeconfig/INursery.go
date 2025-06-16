package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type INursery interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Disallow use of CommonJs module system in favor of ESM style imports.
	// Experimental.
	NoCommonJs() interface{}
	// Experimental.
	SetNoCommonJs(n interface{})
	// Disallow a lower specificity selector from coming after a higher specificity selector.
	// Experimental.
	NoDescendingSpecificity() interface{}
	// Experimental.
	SetNoDescendingSpecificity(n interface{})
	// Disallow direct assignments to document.cookie.
	// Experimental.
	NoDocumentCookie() interface{}
	// Experimental.
	SetNoDocumentCookie(n interface{})
	// Prevents importing next/document outside of pages/_document.jsx in Next.js projects.
	// Experimental.
	NoDocumentImportInPage() interface{}
	// Experimental.
	SetNoDocumentImportInPage(n interface{})
	// Disallow duplicate custom properties within declaration blocks.
	// Experimental.
	NoDuplicateCustomProperties() interface{}
	// Experimental.
	SetNoDuplicateCustomProperties(n interface{})
	// No duplicated fields in GraphQL operations.
	// Experimental.
	NoDuplicatedFields() interface{}
	// Experimental.
	SetNoDuplicatedFields(n interface{})
	// Disallow duplicate conditions in if-else-if chains.
	// Experimental.
	NoDuplicateElseIf() interface{}
	// Experimental.
	SetNoDuplicateElseIf(n interface{})
	// Disallow duplicate properties within declaration blocks.
	// Experimental.
	NoDuplicateProperties() interface{}
	// Experimental.
	SetNoDuplicateProperties(n interface{})
	// Disallow accessing namespace imports dynamically.
	// Experimental.
	NoDynamicNamespaceImportAccess() interface{}
	// Experimental.
	SetNoDynamicNamespaceImportAccess(n interface{})
	// Disallow TypeScript enum.
	// Experimental.
	NoEnum() interface{}
	// Experimental.
	SetNoEnum(n interface{})
	// Disallow exporting an imported variable.
	// Experimental.
	NoExportedImports() interface{}
	// Experimental.
	SetNoExportedImports(n interface{})
	// Prevent usage of \<head> element in a Next.js project.
	// Experimental.
	NoHeadElement() interface{}
	// Experimental.
	SetNoHeadElement(n interface{})
	// Prevent using the next/head module in pages/_document.js on Next.js projects.
	// Experimental.
	NoHeadImportInDocument() interface{}
	// Experimental.
	SetNoHeadImportInDocument(n interface{})
	// Prevent usage of \<img> element in a Next.js project.
	// Experimental.
	NoImgElement() interface{}
	// Experimental.
	SetNoImgElement(n interface{})
	// Disallows the use of irregular whitespace characters.
	// Experimental.
	NoIrregularWhitespace() interface{}
	// Experimental.
	SetNoIrregularWhitespace(n interface{})
	// Disallow missing var function for css variables.
	// Experimental.
	NoMissingVarFunction() interface{}
	// Experimental.
	SetNoMissingVarFunction(n interface{})
	// Disallow nested ternary expressions.
	// Experimental.
	NoNestedTernary() interface{}
	// Experimental.
	SetNoNestedTernary(n interface{})
	// Disallow octal escape sequences in string literals.
	// Experimental.
	NoOctalEscape() interface{}
	// Experimental.
	SetNoOctalEscape(n interface{})
	// Disallow the use of process.env.
	// Experimental.
	NoProcessEnv() interface{}
	// Experimental.
	SetNoProcessEnv(n interface{})
	// Disallow specified modules when loaded by import or require.
	// Experimental.
	NoRestrictedImports() interface{}
	// Experimental.
	SetNoRestrictedImports(n interface{})
	// Disallow user defined types.
	// Experimental.
	NoRestrictedTypes() interface{}
	// Experimental.
	SetNoRestrictedTypes(n interface{})
	// Disallow usage of sensitive data such as API keys and tokens.
	// Experimental.
	NoSecrets() interface{}
	// Experimental.
	SetNoSecrets(n interface{})
	// Enforce that static, visible elements (such as \<div>) that have click handlers use the valid role attribute.
	// Experimental.
	NoStaticElementInteractions() interface{}
	// Experimental.
	SetNoStaticElementInteractions(n interface{})
	// Enforce the use of String.slice() over String.substr() and String.substring().
	// Experimental.
	NoSubstr() interface{}
	// Experimental.
	SetNoSubstr(n interface{})
	// Disallow template literal placeholder syntax in regular strings.
	// Experimental.
	NoTemplateCurlyInString() interface{}
	// Experimental.
	SetNoTemplateCurlyInString(n interface{})
	// Disallow unknown pseudo-class selectors.
	// Experimental.
	NoUnknownPseudoClass() interface{}
	// Experimental.
	SetNoUnknownPseudoClass(n interface{})
	// Disallow unknown pseudo-element selectors.
	// Experimental.
	NoUnknownPseudoElement() interface{}
	// Experimental.
	SetNoUnknownPseudoElement(n interface{})
	// Disallow unknown type selectors.
	// Experimental.
	NoUnknownTypeSelector() interface{}
	// Experimental.
	SetNoUnknownTypeSelector(n interface{})
	// Disallow unnecessary escape sequence in regular expression literals.
	// Experimental.
	NoUselessEscapeInRegex() interface{}
	// Experimental.
	SetNoUselessEscapeInRegex(n interface{})
	// Disallow unnecessary String.raw function in template string literals without any escape sequence.
	// Experimental.
	NoUselessStringRaw() interface{}
	// Experimental.
	SetNoUselessStringRaw(n interface{})
	// Disallow use of.
	// Experimental.
	NoValueAtRule() interface{}
	// Experimental.
	SetNoValueAtRule(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Disallow the use of overload signatures that are not next to each other.
	// Experimental.
	UseAdjacentOverloadSignatures() interface{}
	// Experimental.
	SetUseAdjacentOverloadSignatures(u interface{})
	// Enforce that ARIA properties are valid for the roles that are supported by the element.
	// Experimental.
	UseAriaPropsSupportedByRole() interface{}
	// Experimental.
	SetUseAriaPropsSupportedByRole(u interface{})
	// Use at() instead of integer index access.
	// Experimental.
	UseAtIndex() interface{}
	// Experimental.
	SetUseAtIndex(u interface{})
	// Enforce using single if instead of nested if clauses.
	// Experimental.
	UseCollapsedIf() interface{}
	// Experimental.
	SetUseCollapsedIf(u interface{})
	// Enforce declaring components only within modules that export React Components exclusively.
	// Experimental.
	UseComponentExportOnlyModules() interface{}
	// Experimental.
	SetUseComponentExportOnlyModules(u interface{})
	// This rule enforces consistent use of curly braces inside JSX attributes and JSX children.
	// Experimental.
	UseConsistentCurlyBraces() interface{}
	// Experimental.
	SetUseConsistentCurlyBraces(u interface{})
	// Require consistent accessibility modifiers on class properties and methods.
	// Experimental.
	UseConsistentMemberAccessibility() interface{}
	// Experimental.
	SetUseConsistentMemberAccessibility(u interface{})
	// Require specifying the reason argument when using.
	// Deprecated: directive.
	UseDeprecatedReason() interface{}
	// Deprecated: directive.
	SetUseDeprecatedReason(u interface{})
	// Require explicit return types on functions and class methods.
	// Experimental.
	UseExplicitType() interface{}
	// Experimental.
	SetUseExplicitType(u interface{})
	// Enforces the use of a recommended display strategy with Google Fonts.
	// Experimental.
	UseGoogleFontDisplay() interface{}
	// Experimental.
	SetUseGoogleFontDisplay(u interface{})
	// Require for-in loops to include an if statement.
	// Experimental.
	UseGuardForIn() interface{}
	// Experimental.
	SetUseGuardForIn(u interface{})
	// Disallows package private imports.
	// Experimental.
	UseImportRestrictions() interface{}
	// Experimental.
	SetUseImportRestrictions(u interface{})
	// Enforce the sorting of CSS utility classes.
	// Experimental.
	UseSortedClasses() interface{}
	// Experimental.
	SetUseSortedClasses(u interface{})
	// Enforce the use of the directive "use strict" in script files.
	// Experimental.
	UseStrictMode() interface{}
	// Experimental.
	SetUseStrictMode(u interface{})
	// Enforce the use of String.trimStart() and String.trimEnd() over String.trimLeft() and String.trimRight().
	// Experimental.
	UseTrimStartEnd() interface{}
	// Experimental.
	SetUseTrimStartEnd(u interface{})
	// Use valid values for the autocomplete attribute on input elements.
	// Experimental.
	UseValidAutocomplete() interface{}
	// Experimental.
	SetUseValidAutocomplete(u interface{})
}

// The jsii proxy for INursery
type jsiiProxy_INursery struct {
	_ byte // padding
}

func (j *jsiiProxy_INursery) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_INursery) NoCommonJs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCommonJs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoCommonJs(val interface{}) {
	if err := j.validateSetNoCommonJsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCommonJs",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDescendingSpecificity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDescendingSpecificity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDescendingSpecificity(val interface{}) {
	if err := j.validateSetNoDescendingSpecificityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDescendingSpecificity",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDocumentCookie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDocumentCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDocumentCookie(val interface{}) {
	if err := j.validateSetNoDocumentCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDocumentCookie",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDocumentImportInPage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDocumentImportInPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDocumentImportInPage(val interface{}) {
	if err := j.validateSetNoDocumentImportInPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDocumentImportInPage",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDuplicateCustomProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateCustomProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDuplicateCustomProperties(val interface{}) {
	if err := j.validateSetNoDuplicateCustomPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateCustomProperties",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDuplicatedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicatedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDuplicatedFields(val interface{}) {
	if err := j.validateSetNoDuplicatedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicatedFields",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDuplicateElseIf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateElseIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDuplicateElseIf(val interface{}) {
	if err := j.validateSetNoDuplicateElseIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateElseIf",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDuplicateProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDuplicateProperties(val interface{}) {
	if err := j.validateSetNoDuplicatePropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateProperties",
		val,
	)
}

func (j *jsiiProxy_INursery) NoDynamicNamespaceImportAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDynamicNamespaceImportAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoDynamicNamespaceImportAccess(val interface{}) {
	if err := j.validateSetNoDynamicNamespaceImportAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDynamicNamespaceImportAccess",
		val,
	)
}

func (j *jsiiProxy_INursery) NoEnum() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoEnum(val interface{}) {
	if err := j.validateSetNoEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEnum",
		val,
	)
}

func (j *jsiiProxy_INursery) NoExportedImports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExportedImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoExportedImports(val interface{}) {
	if err := j.validateSetNoExportedImportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExportedImports",
		val,
	)
}

func (j *jsiiProxy_INursery) NoHeadElement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHeadElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoHeadElement(val interface{}) {
	if err := j.validateSetNoHeadElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHeadElement",
		val,
	)
}

func (j *jsiiProxy_INursery) NoHeadImportInDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHeadImportInDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoHeadImportInDocument(val interface{}) {
	if err := j.validateSetNoHeadImportInDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHeadImportInDocument",
		val,
	)
}

func (j *jsiiProxy_INursery) NoImgElement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noImgElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoImgElement(val interface{}) {
	if err := j.validateSetNoImgElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noImgElement",
		val,
	)
}

func (j *jsiiProxy_INursery) NoIrregularWhitespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noIrregularWhitespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoIrregularWhitespace(val interface{}) {
	if err := j.validateSetNoIrregularWhitespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noIrregularWhitespace",
		val,
	)
}

func (j *jsiiProxy_INursery) NoMissingVarFunction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMissingVarFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoMissingVarFunction(val interface{}) {
	if err := j.validateSetNoMissingVarFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMissingVarFunction",
		val,
	)
}

func (j *jsiiProxy_INursery) NoNestedTernary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNestedTernary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoNestedTernary(val interface{}) {
	if err := j.validateSetNoNestedTernaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNestedTernary",
		val,
	)
}

func (j *jsiiProxy_INursery) NoOctalEscape() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noOctalEscape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoOctalEscape(val interface{}) {
	if err := j.validateSetNoOctalEscapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noOctalEscape",
		val,
	)
}

func (j *jsiiProxy_INursery) NoProcessEnv() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noProcessEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoProcessEnv(val interface{}) {
	if err := j.validateSetNoProcessEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noProcessEnv",
		val,
	)
}

func (j *jsiiProxy_INursery) NoRestrictedImports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictedImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoRestrictedImports(val interface{}) {
	if err := j.validateSetNoRestrictedImportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRestrictedImports",
		val,
	)
}

func (j *jsiiProxy_INursery) NoRestrictedTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictedTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoRestrictedTypes(val interface{}) {
	if err := j.validateSetNoRestrictedTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRestrictedTypes",
		val,
	)
}

func (j *jsiiProxy_INursery) NoSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoSecrets(val interface{}) {
	if err := j.validateSetNoSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSecrets",
		val,
	)
}

func (j *jsiiProxy_INursery) NoStaticElementInteractions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStaticElementInteractions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoStaticElementInteractions(val interface{}) {
	if err := j.validateSetNoStaticElementInteractionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStaticElementInteractions",
		val,
	)
}

func (j *jsiiProxy_INursery) NoSubstr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSubstr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoSubstr(val interface{}) {
	if err := j.validateSetNoSubstrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSubstr",
		val,
	)
}

func (j *jsiiProxy_INursery) NoTemplateCurlyInString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTemplateCurlyInString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoTemplateCurlyInString(val interface{}) {
	if err := j.validateSetNoTemplateCurlyInStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTemplateCurlyInString",
		val,
	)
}

func (j *jsiiProxy_INursery) NoUnknownPseudoClass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownPseudoClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoUnknownPseudoClass(val interface{}) {
	if err := j.validateSetNoUnknownPseudoClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownPseudoClass",
		val,
	)
}

func (j *jsiiProxy_INursery) NoUnknownPseudoElement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownPseudoElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoUnknownPseudoElement(val interface{}) {
	if err := j.validateSetNoUnknownPseudoElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownPseudoElement",
		val,
	)
}

func (j *jsiiProxy_INursery) NoUnknownTypeSelector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownTypeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoUnknownTypeSelector(val interface{}) {
	if err := j.validateSetNoUnknownTypeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownTypeSelector",
		val,
	)
}

func (j *jsiiProxy_INursery) NoUselessEscapeInRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessEscapeInRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoUselessEscapeInRegex(val interface{}) {
	if err := j.validateSetNoUselessEscapeInRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessEscapeInRegex",
		val,
	)
}

func (j *jsiiProxy_INursery) NoUselessStringRaw() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessStringRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoUselessStringRaw(val interface{}) {
	if err := j.validateSetNoUselessStringRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessStringRaw",
		val,
	)
}

func (j *jsiiProxy_INursery) NoValueAtRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noValueAtRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetNoValueAtRule(val interface{}) {
	if err := j.validateSetNoValueAtRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noValueAtRule",
		val,
	)
}

func (j *jsiiProxy_INursery) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_INursery) UseAdjacentOverloadSignatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAdjacentOverloadSignatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseAdjacentOverloadSignatures(val interface{}) {
	if err := j.validateSetUseAdjacentOverloadSignaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAdjacentOverloadSignatures",
		val,
	)
}

func (j *jsiiProxy_INursery) UseAriaPropsSupportedByRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAriaPropsSupportedByRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseAriaPropsSupportedByRole(val interface{}) {
	if err := j.validateSetUseAriaPropsSupportedByRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAriaPropsSupportedByRole",
		val,
	)
}

func (j *jsiiProxy_INursery) UseAtIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAtIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseAtIndex(val interface{}) {
	if err := j.validateSetUseAtIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAtIndex",
		val,
	)
}

func (j *jsiiProxy_INursery) UseCollapsedIf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCollapsedIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseCollapsedIf(val interface{}) {
	if err := j.validateSetUseCollapsedIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCollapsedIf",
		val,
	)
}

func (j *jsiiProxy_INursery) UseComponentExportOnlyModules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useComponentExportOnlyModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseComponentExportOnlyModules(val interface{}) {
	if err := j.validateSetUseComponentExportOnlyModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useComponentExportOnlyModules",
		val,
	)
}

func (j *jsiiProxy_INursery) UseConsistentCurlyBraces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useConsistentCurlyBraces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseConsistentCurlyBraces(val interface{}) {
	if err := j.validateSetUseConsistentCurlyBracesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConsistentCurlyBraces",
		val,
	)
}

func (j *jsiiProxy_INursery) UseConsistentMemberAccessibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useConsistentMemberAccessibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseConsistentMemberAccessibility(val interface{}) {
	if err := j.validateSetUseConsistentMemberAccessibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConsistentMemberAccessibility",
		val,
	)
}

func (j *jsiiProxy_INursery) UseDeprecatedReason() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDeprecatedReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseDeprecatedReason(val interface{}) {
	if err := j.validateSetUseDeprecatedReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDeprecatedReason",
		val,
	)
}

func (j *jsiiProxy_INursery) UseExplicitType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExplicitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseExplicitType(val interface{}) {
	if err := j.validateSetUseExplicitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExplicitType",
		val,
	)
}

func (j *jsiiProxy_INursery) UseGoogleFontDisplay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGoogleFontDisplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseGoogleFontDisplay(val interface{}) {
	if err := j.validateSetUseGoogleFontDisplayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGoogleFontDisplay",
		val,
	)
}

func (j *jsiiProxy_INursery) UseGuardForIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGuardForIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseGuardForIn(val interface{}) {
	if err := j.validateSetUseGuardForInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGuardForIn",
		val,
	)
}

func (j *jsiiProxy_INursery) UseImportRestrictions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useImportRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseImportRestrictions(val interface{}) {
	if err := j.validateSetUseImportRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useImportRestrictions",
		val,
	)
}

func (j *jsiiProxy_INursery) UseSortedClasses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSortedClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseSortedClasses(val interface{}) {
	if err := j.validateSetUseSortedClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSortedClasses",
		val,
	)
}

func (j *jsiiProxy_INursery) UseStrictMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStrictMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseStrictMode(val interface{}) {
	if err := j.validateSetUseStrictModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStrictMode",
		val,
	)
}

func (j *jsiiProxy_INursery) UseTrimStartEnd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTrimStartEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseTrimStartEnd(val interface{}) {
	if err := j.validateSetUseTrimStartEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTrimStartEnd",
		val,
	)
}

func (j *jsiiProxy_INursery) UseValidAutocomplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidAutocomplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INursery)SetUseValidAutocomplete(val interface{}) {
	if err := j.validateSetUseValidAutocompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidAutocomplete",
		val,
	)
}

