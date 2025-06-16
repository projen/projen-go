package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type IStyle interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Disallow the use of arguments.
	// Experimental.
	NoArguments() interface{}
	// Experimental.
	SetNoArguments(n interface{})
	// Disallow comma operator.
	// Experimental.
	NoCommaOperator() interface{}
	// Experimental.
	SetNoCommaOperator(n interface{})
	// Disallow default exports.
	// Experimental.
	NoDefaultExport() interface{}
	// Experimental.
	SetNoDefaultExport(n interface{})
	// Disallow using a callback in asynchronous tests and hooks.
	// Experimental.
	NoDoneCallback() interface{}
	// Experimental.
	SetNoDoneCallback(n interface{})
	// Disallow implicit true values on JSX boolean attributes.
	// Experimental.
	NoImplicitBoolean() interface{}
	// Experimental.
	SetNoImplicitBoolean(n interface{})
	// Disallow type annotations for variables, parameters, and class properties initialized with a literal expression.
	// Experimental.
	NoInferrableTypes() interface{}
	// Experimental.
	SetNoInferrableTypes(n interface{})
	// Disallow the use of TypeScript's namespaces.
	// Experimental.
	NoNamespace() interface{}
	// Experimental.
	SetNoNamespace(n interface{})
	// Disallow the use of namespace imports.
	// Experimental.
	NoNamespaceImport() interface{}
	// Experimental.
	SetNoNamespaceImport(n interface{})
	// Disallow negation in the condition of an if statement if it has an else clause.
	// Experimental.
	NoNegationElse() interface{}
	// Experimental.
	SetNoNegationElse(n interface{})
	// Disallow non-null assertions using the !
	//
	// postfix operator.
	// Experimental.
	NoNonNullAssertion() interface{}
	// Experimental.
	SetNoNonNullAssertion(n interface{})
	// Disallow reassigning function parameters.
	// Experimental.
	NoParameterAssign() interface{}
	// Experimental.
	SetNoParameterAssign(n interface{})
	// Disallow the use of parameter properties in class constructors.
	// Experimental.
	NoParameterProperties() interface{}
	// Experimental.
	SetNoParameterProperties(n interface{})
	// This rule allows you to specify global variable names that you don’t want to use in your application.
	// Experimental.
	NoRestrictedGlobals() interface{}
	// Experimental.
	SetNoRestrictedGlobals(n interface{})
	// Disallow the use of constants which its value is the upper-case version of its name.
	// Experimental.
	NoShoutyConstants() interface{}
	// Experimental.
	SetNoShoutyConstants(n interface{})
	// Disallow template literals if interpolation and special-character handling are not needed.
	// Experimental.
	NoUnusedTemplateLiteral() interface{}
	// Experimental.
	SetNoUnusedTemplateLiteral(n interface{})
	// Disallow else block when the if block breaks early.
	// Experimental.
	NoUselessElse() interface{}
	// Experimental.
	SetNoUselessElse(n interface{})
	// Disallow the use of var.
	// Experimental.
	NoVar() interface{}
	// Experimental.
	SetNoVar(n interface{})
	// Disallow the use of yoda expressions.
	// Experimental.
	NoYodaExpression() interface{}
	// Experimental.
	SetNoYodaExpression(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Enforce the use of as const over literal type and type annotation.
	// Experimental.
	UseAsConstAssertion() interface{}
	// Experimental.
	SetUseAsConstAssertion(u interface{})
	// Requires following curly brace conventions.
	// Experimental.
	UseBlockStatements() interface{}
	// Experimental.
	SetUseBlockStatements(u interface{})
	// Enforce using else if instead of nested if in else clauses.
	// Experimental.
	UseCollapsedElseIf() interface{}
	// Experimental.
	SetUseCollapsedElseIf(u interface{})
	// Require consistently using either T\[] or Array\<T>.
	// Experimental.
	UseConsistentArrayType() interface{}
	// Experimental.
	SetUseConsistentArrayType(u interface{})
	// Enforce the use of new for all builtins, except String, Number and Boolean.
	// Experimental.
	UseConsistentBuiltinInstantiation() interface{}
	// Experimental.
	SetUseConsistentBuiltinInstantiation(u interface{})
	// Require const declarations for variables that are only assigned once.
	// Experimental.
	UseConst() interface{}
	// Experimental.
	SetUseConst(u interface{})
	// Enforce default function parameters and optional function parameters to be last.
	// Experimental.
	UseDefaultParameterLast() interface{}
	// Experimental.
	SetUseDefaultParameterLast(u interface{})
	// Require the default clause in switch statements.
	// Experimental.
	UseDefaultSwitchClause() interface{}
	// Experimental.
	SetUseDefaultSwitchClause(u interface{})
	// Require that each enum member value be explicitly initialized.
	// Experimental.
	UseEnumInitializers() interface{}
	// Experimental.
	SetUseEnumInitializers(u interface{})
	// Enforce explicitly comparing the length, size, byteLength or byteOffset property of a value.
	// Experimental.
	UseExplicitLengthCheck() interface{}
	// Experimental.
	SetUseExplicitLengthCheck(u interface{})
	// Disallow the use of Math.pow in favor of the ** operator.
	// Experimental.
	UseExponentiationOperator() interface{}
	// Experimental.
	SetUseExponentiationOperator(u interface{})
	// Promotes the use of export type for types.
	// Experimental.
	UseExportType() interface{}
	// Experimental.
	SetUseExportType(u interface{})
	// Enforce naming conventions for JavaScript and TypeScript filenames.
	// Experimental.
	UseFilenamingConvention() interface{}
	// Experimental.
	SetUseFilenamingConvention(u interface{})
	// This rule recommends a for-of loop when in a for loop, the index used to extract an item from the iterated array.
	// Experimental.
	UseForOf() interface{}
	// Experimental.
	SetUseForOf(u interface{})
	// This rule enforces the use of \<>...\</> over \<Fragment>...\</Fragment>.
	// Experimental.
	UseFragmentSyntax() interface{}
	// Experimental.
	SetUseFragmentSyntax(u interface{})
	// Promotes the use of import type for types.
	// Experimental.
	UseImportType() interface{}
	// Experimental.
	SetUseImportType(u interface{})
	// Require all enum members to be literal values.
	// Experimental.
	UseLiteralEnumMembers() interface{}
	// Experimental.
	SetUseLiteralEnumMembers(u interface{})
	// Enforce naming conventions for everything across a codebase.
	// Experimental.
	UseNamingConvention() interface{}
	// Experimental.
	SetUseNamingConvention(u interface{})
	// Promotes the usage of node:assert/strict over node:assert.
	// Experimental.
	UseNodeAssertStrict() interface{}
	// Experimental.
	SetUseNodeAssertStrict(u interface{})
	// Enforces using the node: protocol for Node.js builtin modules.
	// Experimental.
	UseNodejsImportProtocol() interface{}
	// Experimental.
	SetUseNodejsImportProtocol(u interface{})
	// Use the Number properties instead of global ones.
	// Experimental.
	UseNumberNamespace() interface{}
	// Experimental.
	SetUseNumberNamespace(u interface{})
	// Disallow parseInt() and Number.parseInt() in favor of binary, octal, and hexadecimal literals.
	// Experimental.
	UseNumericLiterals() interface{}
	// Experimental.
	SetUseNumericLiterals(u interface{})
	// Prevent extra closing tags for components without children.
	// Experimental.
	UseSelfClosingElements() interface{}
	// Experimental.
	SetUseSelfClosingElements(u interface{})
	// When expressing array types, this rule promotes the usage of T\[] shorthand instead of Array\<T>.
	// Experimental.
	UseShorthandArrayType() interface{}
	// Experimental.
	SetUseShorthandArrayType(u interface{})
	// Require assignment operator shorthand where possible.
	// Experimental.
	UseShorthandAssign() interface{}
	// Experimental.
	SetUseShorthandAssign(u interface{})
	// Enforce using function types instead of object type with call signatures.
	// Experimental.
	UseShorthandFunctionType() interface{}
	// Experimental.
	SetUseShorthandFunctionType(u interface{})
	// Enforces switch clauses have a single statement, emits a quick fix wrapping the statements in a block.
	// Experimental.
	UseSingleCaseStatement() interface{}
	// Experimental.
	SetUseSingleCaseStatement(u interface{})
	// Disallow multiple variable declarations in the same variable statement.
	// Experimental.
	UseSingleVarDeclarator() interface{}
	// Experimental.
	SetUseSingleVarDeclarator(u interface{})
	// Prefer template literals over string concatenation.
	// Experimental.
	UseTemplate() interface{}
	// Experimental.
	SetUseTemplate(u interface{})
	// Require new when throwing an error.
	// Experimental.
	UseThrowNewError() interface{}
	// Experimental.
	SetUseThrowNewError(u interface{})
	// Disallow throwing non-Error values.
	// Experimental.
	UseThrowOnlyError() interface{}
	// Experimental.
	SetUseThrowOnlyError(u interface{})
	// Enforce the use of while loops instead of for loops when the initializer and update expressions are not needed.
	// Experimental.
	UseWhile() interface{}
	// Experimental.
	SetUseWhile(u interface{})
}

// The jsii proxy for IStyle
type jsiiProxy_IStyle struct {
	_ byte // padding
}

func (j *jsiiProxy_IStyle) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoArguments(val interface{}) {
	if err := j.validateSetNoArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noArguments",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoCommaOperator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCommaOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoCommaOperator(val interface{}) {
	if err := j.validateSetNoCommaOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCommaOperator",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoDefaultExport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoDefaultExport(val interface{}) {
	if err := j.validateSetNoDefaultExportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDefaultExport",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoDoneCallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDoneCallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoDoneCallback(val interface{}) {
	if err := j.validateSetNoDoneCallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDoneCallback",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoImplicitBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noImplicitBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoImplicitBoolean(val interface{}) {
	if err := j.validateSetNoImplicitBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noImplicitBoolean",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoInferrableTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInferrableTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoInferrableTypes(val interface{}) {
	if err := j.validateSetNoInferrableTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInferrableTypes",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoNamespace(val interface{}) {
	if err := j.validateSetNoNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNamespace",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoNamespaceImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespaceImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoNamespaceImport(val interface{}) {
	if err := j.validateSetNoNamespaceImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNamespaceImport",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoNegationElse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNegationElse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoNegationElse(val interface{}) {
	if err := j.validateSetNoNegationElseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNegationElse",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoNonNullAssertion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNonNullAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoNonNullAssertion(val interface{}) {
	if err := j.validateSetNoNonNullAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNonNullAssertion",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoParameterAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParameterAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoParameterAssign(val interface{}) {
	if err := j.validateSetNoParameterAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noParameterAssign",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoParameterProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParameterProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoParameterProperties(val interface{}) {
	if err := j.validateSetNoParameterPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noParameterProperties",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoRestrictedGlobals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictedGlobals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoRestrictedGlobals(val interface{}) {
	if err := j.validateSetNoRestrictedGlobalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRestrictedGlobals",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoShoutyConstants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noShoutyConstants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoShoutyConstants(val interface{}) {
	if err := j.validateSetNoShoutyConstantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noShoutyConstants",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoUnusedTemplateLiteral() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedTemplateLiteral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoUnusedTemplateLiteral(val interface{}) {
	if err := j.validateSetNoUnusedTemplateLiteralParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedTemplateLiteral",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoUselessElse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessElse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoUselessElse(val interface{}) {
	if err := j.validateSetNoUselessElseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessElse",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoVar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noVar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoVar(val interface{}) {
	if err := j.validateSetNoVarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noVar",
		val,
	)
}

func (j *jsiiProxy_IStyle) NoYodaExpression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noYodaExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetNoYodaExpression(val interface{}) {
	if err := j.validateSetNoYodaExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noYodaExpression",
		val,
	)
}

func (j *jsiiProxy_IStyle) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseAsConstAssertion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAsConstAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseAsConstAssertion(val interface{}) {
	if err := j.validateSetUseAsConstAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAsConstAssertion",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseBlockStatements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlockStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseBlockStatements(val interface{}) {
	if err := j.validateSetUseBlockStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBlockStatements",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseCollapsedElseIf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCollapsedElseIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseCollapsedElseIf(val interface{}) {
	if err := j.validateSetUseCollapsedElseIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCollapsedElseIf",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseConsistentArrayType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useConsistentArrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseConsistentArrayType(val interface{}) {
	if err := j.validateSetUseConsistentArrayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConsistentArrayType",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseConsistentBuiltinInstantiation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useConsistentBuiltinInstantiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseConsistentBuiltinInstantiation(val interface{}) {
	if err := j.validateSetUseConsistentBuiltinInstantiationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConsistentBuiltinInstantiation",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseConst() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useConst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseConst(val interface{}) {
	if err := j.validateSetUseConstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConst",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseDefaultParameterLast() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultParameterLast",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseDefaultParameterLast(val interface{}) {
	if err := j.validateSetUseDefaultParameterLastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultParameterLast",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseDefaultSwitchClause() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSwitchClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseDefaultSwitchClause(val interface{}) {
	if err := j.validateSetUseDefaultSwitchClauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultSwitchClause",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseEnumInitializers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEnumInitializers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseEnumInitializers(val interface{}) {
	if err := j.validateSetUseEnumInitializersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useEnumInitializers",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseExplicitLengthCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExplicitLengthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseExplicitLengthCheck(val interface{}) {
	if err := j.validateSetUseExplicitLengthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExplicitLengthCheck",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseExponentiationOperator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExponentiationOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseExponentiationOperator(val interface{}) {
	if err := j.validateSetUseExponentiationOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExponentiationOperator",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseExportType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseExportType(val interface{}) {
	if err := j.validateSetUseExportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExportType",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseFilenamingConvention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFilenamingConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseFilenamingConvention(val interface{}) {
	if err := j.validateSetUseFilenamingConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFilenamingConvention",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseForOf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useForOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseForOf(val interface{}) {
	if err := j.validateSetUseForOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useForOf",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseFragmentSyntax() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFragmentSyntax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseFragmentSyntax(val interface{}) {
	if err := j.validateSetUseFragmentSyntaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFragmentSyntax",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseImportType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useImportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseImportType(val interface{}) {
	if err := j.validateSetUseImportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useImportType",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseLiteralEnumMembers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLiteralEnumMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseLiteralEnumMembers(val interface{}) {
	if err := j.validateSetUseLiteralEnumMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLiteralEnumMembers",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseNamingConvention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNamingConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseNamingConvention(val interface{}) {
	if err := j.validateSetUseNamingConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNamingConvention",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseNodeAssertStrict() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNodeAssertStrict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseNodeAssertStrict(val interface{}) {
	if err := j.validateSetUseNodeAssertStrictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNodeAssertStrict",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseNodejsImportProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNodejsImportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseNodejsImportProtocol(val interface{}) {
	if err := j.validateSetUseNodejsImportProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNodejsImportProtocol",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseNumberNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNumberNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseNumberNamespace(val interface{}) {
	if err := j.validateSetUseNumberNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNumberNamespace",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseNumericLiterals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNumericLiterals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseNumericLiterals(val interface{}) {
	if err := j.validateSetUseNumericLiteralsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNumericLiterals",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseSelfClosingElements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSelfClosingElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseSelfClosingElements(val interface{}) {
	if err := j.validateSetUseSelfClosingElementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSelfClosingElements",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseShorthandArrayType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useShorthandArrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseShorthandArrayType(val interface{}) {
	if err := j.validateSetUseShorthandArrayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useShorthandArrayType",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseShorthandAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useShorthandAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseShorthandAssign(val interface{}) {
	if err := j.validateSetUseShorthandAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useShorthandAssign",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseShorthandFunctionType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useShorthandFunctionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseShorthandFunctionType(val interface{}) {
	if err := j.validateSetUseShorthandFunctionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useShorthandFunctionType",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseSingleCaseStatement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSingleCaseStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseSingleCaseStatement(val interface{}) {
	if err := j.validateSetUseSingleCaseStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSingleCaseStatement",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseSingleVarDeclarator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSingleVarDeclarator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseSingleVarDeclarator(val interface{}) {
	if err := j.validateSetUseSingleVarDeclaratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSingleVarDeclarator",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseTemplate(val interface{}) {
	if err := j.validateSetUseTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTemplate",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseThrowNewError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useThrowNewError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseThrowNewError(val interface{}) {
	if err := j.validateSetUseThrowNewErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useThrowNewError",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseThrowOnlyError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useThrowOnlyError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseThrowOnlyError(val interface{}) {
	if err := j.validateSetUseThrowOnlyErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useThrowOnlyError",
		val,
	)
}

func (j *jsiiProxy_IStyle) UseWhile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useWhile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStyle)SetUseWhile(val interface{}) {
	if err := j.validateSetUseWhileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useWhile",
		val,
	)
}

