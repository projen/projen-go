package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type ISuspicious interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Use standard constants instead of approximated literals.
	// Experimental.
	NoApproximativeNumericConstant() interface{}
	// Experimental.
	SetNoApproximativeNumericConstant(n interface{})
	// Discourage the usage of Array index in keys.
	// Experimental.
	NoArrayIndexKey() interface{}
	// Experimental.
	SetNoArrayIndexKey(n interface{})
	// Disallow assignments in expressions.
	// Experimental.
	NoAssignInExpressions() interface{}
	// Experimental.
	SetNoAssignInExpressions(n interface{})
	// Disallows using an async function as a Promise executor.
	// Experimental.
	NoAsyncPromiseExecutor() interface{}
	// Experimental.
	SetNoAsyncPromiseExecutor(n interface{})
	// Disallow reassigning exceptions in catch clauses.
	// Experimental.
	NoCatchAssign() interface{}
	// Experimental.
	SetNoCatchAssign(n interface{})
	// Disallow reassigning class members.
	// Experimental.
	NoClassAssign() interface{}
	// Experimental.
	SetNoClassAssign(n interface{})
	// Prevent comments from being inserted as text nodes.
	// Experimental.
	NoCommentText() interface{}
	// Experimental.
	SetNoCommentText(n interface{})
	// Disallow comparing against -0.
	// Experimental.
	NoCompareNegZero() interface{}
	// Experimental.
	SetNoCompareNegZero(n interface{})
	// Disallow labeled statements that are not loops.
	// Experimental.
	NoConfusingLabels() interface{}
	// Experimental.
	SetNoConfusingLabels(n interface{})
	// Disallow void type outside of generic or return types.
	// Experimental.
	NoConfusingVoidType() interface{}
	// Experimental.
	SetNoConfusingVoidType(n interface{})
	// Disallow the use of console.
	// Experimental.
	NoConsole() interface{}
	// Experimental.
	SetNoConsole(n interface{})
	// Disallow the use of console.log.
	// Experimental.
	NoConsoleLog() interface{}
	// Experimental.
	SetNoConsoleLog(n interface{})
	// Disallow TypeScript const enum.
	// Experimental.
	NoConstEnum() interface{}
	// Experimental.
	SetNoConstEnum(n interface{})
	// Prevents from having control characters and some escape sequences that match control characters in regular expressions.
	// Experimental.
	NoControlCharactersInRegex() interface{}
	// Experimental.
	SetNoControlCharactersInRegex(n interface{})
	// Disallow the use of debugger.
	// Experimental.
	NoDebugger() interface{}
	// Experimental.
	SetNoDebugger(n interface{})
	// Require the use of === and !==.
	// Experimental.
	NoDoubleEquals() interface{}
	// Experimental.
	SetNoDoubleEquals(n interface{})
	// Disallow duplicate.
	// Experimental.
	NoDuplicateAtImportRules() interface{}
	// Experimental.
	SetNoDuplicateAtImportRules(n interface{})
	// Disallow duplicate case labels.
	// Experimental.
	NoDuplicateCase() interface{}
	// Experimental.
	SetNoDuplicateCase(n interface{})
	// Disallow duplicate class members.
	// Experimental.
	NoDuplicateClassMembers() interface{}
	// Experimental.
	SetNoDuplicateClassMembers(n interface{})
	// Disallow duplicate names within font families.
	// Experimental.
	NoDuplicateFontNames() interface{}
	// Experimental.
	SetNoDuplicateFontNames(n interface{})
	// Prevents JSX properties to be assigned multiple times.
	// Experimental.
	NoDuplicateJsxProps() interface{}
	// Experimental.
	SetNoDuplicateJsxProps(n interface{})
	// Disallow two keys with the same name inside objects.
	// Experimental.
	NoDuplicateObjectKeys() interface{}
	// Experimental.
	SetNoDuplicateObjectKeys(n interface{})
	// Disallow duplicate function parameter name.
	// Experimental.
	NoDuplicateParameters() interface{}
	// Experimental.
	SetNoDuplicateParameters(n interface{})
	// Disallow duplicate selectors within keyframe blocks.
	// Experimental.
	NoDuplicateSelectorsKeyframeBlock() interface{}
	// Experimental.
	SetNoDuplicateSelectorsKeyframeBlock(n interface{})
	// A describe block should not contain duplicate hooks.
	// Experimental.
	NoDuplicateTestHooks() interface{}
	// Experimental.
	SetNoDuplicateTestHooks(n interface{})
	// Disallow CSS empty blocks.
	// Experimental.
	NoEmptyBlock() interface{}
	// Experimental.
	SetNoEmptyBlock(n interface{})
	// Disallow empty block statements and static blocks.
	// Experimental.
	NoEmptyBlockStatements() interface{}
	// Experimental.
	SetNoEmptyBlockStatements(n interface{})
	// Disallow the declaration of empty interfaces.
	// Experimental.
	NoEmptyInterface() interface{}
	// Experimental.
	SetNoEmptyInterface(n interface{})
	// Disallow variables from evolving into any type through reassignments.
	// Experimental.
	NoEvolvingTypes() interface{}
	// Experimental.
	SetNoEvolvingTypes(n interface{})
	// Disallow the any type usage.
	// Experimental.
	NoExplicitAny() interface{}
	// Experimental.
	SetNoExplicitAny(n interface{})
	// Disallow using export or module.exports in files containing tests.
	// Experimental.
	NoExportsInTest() interface{}
	// Experimental.
	SetNoExportsInTest(n interface{})
	// Prevents the wrong usage of the non-null assertion operator (!) in TypeScript files.
	// Experimental.
	NoExtraNonNullAssertion() interface{}
	// Experimental.
	SetNoExtraNonNullAssertion(n interface{})
	// Disallow fallthrough of switch clauses.
	// Experimental.
	NoFallthroughSwitchClause() interface{}
	// Experimental.
	SetNoFallthroughSwitchClause(n interface{})
	// Disallow focused tests.
	// Experimental.
	NoFocusedTests() interface{}
	// Experimental.
	SetNoFocusedTests(n interface{})
	// Disallow reassigning function declarations.
	// Experimental.
	NoFunctionAssign() interface{}
	// Experimental.
	SetNoFunctionAssign(n interface{})
	// Disallow assignments to native objects and read-only global variables.
	// Experimental.
	NoGlobalAssign() interface{}
	// Experimental.
	SetNoGlobalAssign(n interface{})
	// Use Number.isFinite instead of global isFinite.
	// Experimental.
	NoGlobalIsFinite() interface{}
	// Experimental.
	SetNoGlobalIsFinite(n interface{})
	// Use Number.isNaN instead of global isNaN.
	// Experimental.
	NoGlobalIsNan() interface{}
	// Experimental.
	SetNoGlobalIsNan(n interface{})
	// Disallow use of implicit any type on variable declarations.
	// Experimental.
	NoImplicitAnyLet() interface{}
	// Experimental.
	SetNoImplicitAnyLet(n interface{})
	// Disallow invalid !important within keyframe declarations.
	// Experimental.
	NoImportantInKeyframe() interface{}
	// Experimental.
	SetNoImportantInKeyframe(n interface{})
	// Disallow assigning to imported bindings.
	// Experimental.
	NoImportAssign() interface{}
	// Experimental.
	SetNoImportAssign(n interface{})
	// Disallow labels that share a name with a variable.
	// Experimental.
	NoLabelVar() interface{}
	// Experimental.
	SetNoLabelVar(n interface{})
	// Disallow characters made with multiple code points in character class syntax.
	// Experimental.
	NoMisleadingCharacterClass() interface{}
	// Experimental.
	SetNoMisleadingCharacterClass(n interface{})
	// Enforce proper usage of new and constructor.
	// Experimental.
	NoMisleadingInstantiator() interface{}
	// Experimental.
	SetNoMisleadingInstantiator(n interface{})
	// Checks that the assertion function, for example expect, is placed inside an it() function call.
	// Experimental.
	NoMisplacedAssertion() interface{}
	// Experimental.
	SetNoMisplacedAssertion(n interface{})
	// Disallow shorthand assign when variable appears on both sides.
	// Experimental.
	NoMisrefactoredShorthandAssign() interface{}
	// Experimental.
	SetNoMisrefactoredShorthandAssign(n interface{})
	// Disallow direct use of Object.prototype builtins.
	// Experimental.
	NoPrototypeBuiltins() interface{}
	// Experimental.
	SetNoPrototypeBuiltins(n interface{})
	// Prevents React-specific JSX properties from being used.
	// Experimental.
	NoReactSpecificProps() interface{}
	// Experimental.
	SetNoReactSpecificProps(n interface{})
	// Disallow variable, function, class, and type redeclarations in the same scope.
	// Experimental.
	NoRedeclare() interface{}
	// Experimental.
	SetNoRedeclare(n interface{})
	// Prevents from having redundant "use strict".
	// Experimental.
	NoRedundantUseStrict() interface{}
	// Experimental.
	SetNoRedundantUseStrict(n interface{})
	// Disallow comparisons where both sides are exactly the same.
	// Experimental.
	NoSelfCompare() interface{}
	// Experimental.
	SetNoSelfCompare(n interface{})
	// Disallow identifiers from shadowing restricted names.
	// Experimental.
	NoShadowRestrictedNames() interface{}
	// Experimental.
	SetNoShadowRestrictedNames(n interface{})
	// Disallow shorthand properties that override related longhand properties.
	// Experimental.
	NoShorthandPropertyOverrides() interface{}
	// Experimental.
	SetNoShorthandPropertyOverrides(n interface{})
	// Disallow disabled tests.
	// Experimental.
	NoSkippedTests() interface{}
	// Experimental.
	SetNoSkippedTests(n interface{})
	// Disallow sparse arrays.
	// Experimental.
	NoSparseArray() interface{}
	// Experimental.
	SetNoSparseArray(n interface{})
	// It detects possible "wrong" semicolons inside JSX elements.
	// Experimental.
	NoSuspiciousSemicolonInJsx() interface{}
	// Experimental.
	SetNoSuspiciousSemicolonInJsx(n interface{})
	// Disallow then property.
	// Experimental.
	NoThenProperty() interface{}
	// Experimental.
	SetNoThenProperty(n interface{})
	// Disallow unsafe declaration merging between interfaces and classes.
	// Experimental.
	NoUnsafeDeclarationMerging() interface{}
	// Experimental.
	SetNoUnsafeDeclarationMerging(n interface{})
	// Disallow using unsafe negation.
	// Experimental.
	NoUnsafeNegation() interface{}
	// Experimental.
	SetNoUnsafeNegation(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Ensure async functions utilize await.
	// Experimental.
	UseAwait() interface{}
	// Experimental.
	SetUseAwait(u interface{})
	// Enforce default clauses in switch statements to be last.
	// Experimental.
	UseDefaultSwitchClauseLast() interface{}
	// Experimental.
	SetUseDefaultSwitchClauseLast(u interface{})
	// Enforce passing a message value when creating a built-in error.
	// Experimental.
	UseErrorMessage() interface{}
	// Experimental.
	SetUseErrorMessage(u interface{})
	// Enforce get methods to always return a value.
	// Experimental.
	UseGetterReturn() interface{}
	// Experimental.
	SetUseGetterReturn(u interface{})
	// Use Array.isArray() instead of instanceof Array.
	// Experimental.
	UseIsArray() interface{}
	// Experimental.
	SetUseIsArray(u interface{})
	// Require using the namespace keyword over the module keyword to declare TypeScript namespaces.
	// Experimental.
	UseNamespaceKeyword() interface{}
	// Experimental.
	SetUseNamespaceKeyword(u interface{})
	// Enforce using the digits argument with Number#toFixed().
	// Experimental.
	UseNumberToFixedDigitsArgument() interface{}
	// Experimental.
	SetUseNumberToFixedDigitsArgument(u interface{})
	// This rule verifies the result of typeof $expr unary expressions is being compared to valid values, either string literals containing valid type names or other typeof expressions.
	// Experimental.
	UseValidTypeof() interface{}
	// Experimental.
	SetUseValidTypeof(u interface{})
}

// The jsii proxy for ISuspicious
type jsiiProxy_ISuspicious struct {
	_ byte // padding
}

func (j *jsiiProxy_ISuspicious) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoApproximativeNumericConstant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noApproximativeNumericConstant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoApproximativeNumericConstant(val interface{}) {
	if err := j.validateSetNoApproximativeNumericConstantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noApproximativeNumericConstant",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoArrayIndexKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noArrayIndexKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoArrayIndexKey(val interface{}) {
	if err := j.validateSetNoArrayIndexKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noArrayIndexKey",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoAssignInExpressions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAssignInExpressions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoAssignInExpressions(val interface{}) {
	if err := j.validateSetNoAssignInExpressionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAssignInExpressions",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoAsyncPromiseExecutor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAsyncPromiseExecutor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoAsyncPromiseExecutor(val interface{}) {
	if err := j.validateSetNoAsyncPromiseExecutorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAsyncPromiseExecutor",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoCatchAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCatchAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoCatchAssign(val interface{}) {
	if err := j.validateSetNoCatchAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCatchAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoClassAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noClassAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoClassAssign(val interface{}) {
	if err := j.validateSetNoClassAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noClassAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoCommentText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCommentText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoCommentText(val interface{}) {
	if err := j.validateSetNoCommentTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCommentText",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoCompareNegZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCompareNegZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoCompareNegZero(val interface{}) {
	if err := j.validateSetNoCompareNegZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCompareNegZero",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoConfusingLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConfusingLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoConfusingLabels(val interface{}) {
	if err := j.validateSetNoConfusingLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConfusingLabels",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoConfusingVoidType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConfusingVoidType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoConfusingVoidType(val interface{}) {
	if err := j.validateSetNoConfusingVoidTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConfusingVoidType",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoConsole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConsole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoConsole(val interface{}) {
	if err := j.validateSetNoConsoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConsole",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoConsoleLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConsoleLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoConsoleLog(val interface{}) {
	if err := j.validateSetNoConsoleLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConsoleLog",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoConstEnum() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConstEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoConstEnum(val interface{}) {
	if err := j.validateSetNoConstEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConstEnum",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoControlCharactersInRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noControlCharactersInRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoControlCharactersInRegex(val interface{}) {
	if err := j.validateSetNoControlCharactersInRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noControlCharactersInRegex",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDebugger() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDebugger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDebugger(val interface{}) {
	if err := j.validateSetNoDebuggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDebugger",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDoubleEquals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDoubleEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDoubleEquals(val interface{}) {
	if err := j.validateSetNoDoubleEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDoubleEquals",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateAtImportRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateAtImportRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateAtImportRules(val interface{}) {
	if err := j.validateSetNoDuplicateAtImportRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateAtImportRules",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateCase(val interface{}) {
	if err := j.validateSetNoDuplicateCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateCase",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateClassMembers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateClassMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateClassMembers(val interface{}) {
	if err := j.validateSetNoDuplicateClassMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateClassMembers",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateFontNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateFontNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateFontNames(val interface{}) {
	if err := j.validateSetNoDuplicateFontNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateFontNames",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateJsxProps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateJsxProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateJsxProps(val interface{}) {
	if err := j.validateSetNoDuplicateJsxPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateJsxProps",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateObjectKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateObjectKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateObjectKeys(val interface{}) {
	if err := j.validateSetNoDuplicateObjectKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateObjectKeys",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateParameters(val interface{}) {
	if err := j.validateSetNoDuplicateParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateParameters",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateSelectorsKeyframeBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateSelectorsKeyframeBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateSelectorsKeyframeBlock(val interface{}) {
	if err := j.validateSetNoDuplicateSelectorsKeyframeBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateSelectorsKeyframeBlock",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoDuplicateTestHooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDuplicateTestHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoDuplicateTestHooks(val interface{}) {
	if err := j.validateSetNoDuplicateTestHooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDuplicateTestHooks",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoEmptyBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoEmptyBlock(val interface{}) {
	if err := j.validateSetNoEmptyBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyBlock",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoEmptyBlockStatements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyBlockStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoEmptyBlockStatements(val interface{}) {
	if err := j.validateSetNoEmptyBlockStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyBlockStatements",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoEmptyInterface() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoEmptyInterface(val interface{}) {
	if err := j.validateSetNoEmptyInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyInterface",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoEvolvingTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEvolvingTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoEvolvingTypes(val interface{}) {
	if err := j.validateSetNoEvolvingTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEvolvingTypes",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoExplicitAny() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExplicitAny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoExplicitAny(val interface{}) {
	if err := j.validateSetNoExplicitAnyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExplicitAny",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoExportsInTest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExportsInTest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoExportsInTest(val interface{}) {
	if err := j.validateSetNoExportsInTestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExportsInTest",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoExtraNonNullAssertion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExtraNonNullAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoExtraNonNullAssertion(val interface{}) {
	if err := j.validateSetNoExtraNonNullAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExtraNonNullAssertion",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoFallthroughSwitchClause() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noFallthroughSwitchClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoFallthroughSwitchClause(val interface{}) {
	if err := j.validateSetNoFallthroughSwitchClauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noFallthroughSwitchClause",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoFocusedTests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noFocusedTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoFocusedTests(val interface{}) {
	if err := j.validateSetNoFocusedTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noFocusedTests",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoFunctionAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noFunctionAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoFunctionAssign(val interface{}) {
	if err := j.validateSetNoFunctionAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noFunctionAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoGlobalAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGlobalAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoGlobalAssign(val interface{}) {
	if err := j.validateSetNoGlobalAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGlobalAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoGlobalIsFinite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGlobalIsFinite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoGlobalIsFinite(val interface{}) {
	if err := j.validateSetNoGlobalIsFiniteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGlobalIsFinite",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoGlobalIsNan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGlobalIsNan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoGlobalIsNan(val interface{}) {
	if err := j.validateSetNoGlobalIsNanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGlobalIsNan",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoImplicitAnyLet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noImplicitAnyLet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoImplicitAnyLet(val interface{}) {
	if err := j.validateSetNoImplicitAnyLetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noImplicitAnyLet",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoImportantInKeyframe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noImportantInKeyframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoImportantInKeyframe(val interface{}) {
	if err := j.validateSetNoImportantInKeyframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noImportantInKeyframe",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoImportAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noImportAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoImportAssign(val interface{}) {
	if err := j.validateSetNoImportAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noImportAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoLabelVar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noLabelVar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoLabelVar(val interface{}) {
	if err := j.validateSetNoLabelVarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noLabelVar",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoMisleadingCharacterClass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMisleadingCharacterClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoMisleadingCharacterClass(val interface{}) {
	if err := j.validateSetNoMisleadingCharacterClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMisleadingCharacterClass",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoMisleadingInstantiator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMisleadingInstantiator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoMisleadingInstantiator(val interface{}) {
	if err := j.validateSetNoMisleadingInstantiatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMisleadingInstantiator",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoMisplacedAssertion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMisplacedAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoMisplacedAssertion(val interface{}) {
	if err := j.validateSetNoMisplacedAssertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMisplacedAssertion",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoMisrefactoredShorthandAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMisrefactoredShorthandAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoMisrefactoredShorthandAssign(val interface{}) {
	if err := j.validateSetNoMisrefactoredShorthandAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMisrefactoredShorthandAssign",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoPrototypeBuiltins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPrototypeBuiltins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoPrototypeBuiltins(val interface{}) {
	if err := j.validateSetNoPrototypeBuiltinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPrototypeBuiltins",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoReactSpecificProps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noReactSpecificProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoReactSpecificProps(val interface{}) {
	if err := j.validateSetNoReactSpecificPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noReactSpecificProps",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoRedeclare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRedeclare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoRedeclare(val interface{}) {
	if err := j.validateSetNoRedeclareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRedeclare",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoRedundantUseStrict() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRedundantUseStrict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoRedundantUseStrict(val interface{}) {
	if err := j.validateSetNoRedundantUseStrictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRedundantUseStrict",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoSelfCompare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSelfCompare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoSelfCompare(val interface{}) {
	if err := j.validateSetNoSelfCompareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSelfCompare",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoShadowRestrictedNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noShadowRestrictedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoShadowRestrictedNames(val interface{}) {
	if err := j.validateSetNoShadowRestrictedNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noShadowRestrictedNames",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoShorthandPropertyOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noShorthandPropertyOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoShorthandPropertyOverrides(val interface{}) {
	if err := j.validateSetNoShorthandPropertyOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noShorthandPropertyOverrides",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoSkippedTests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSkippedTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoSkippedTests(val interface{}) {
	if err := j.validateSetNoSkippedTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSkippedTests",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoSparseArray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSparseArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoSparseArray(val interface{}) {
	if err := j.validateSetNoSparseArrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSparseArray",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoSuspiciousSemicolonInJsx() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSuspiciousSemicolonInJsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoSuspiciousSemicolonInJsx(val interface{}) {
	if err := j.validateSetNoSuspiciousSemicolonInJsxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSuspiciousSemicolonInJsx",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoThenProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noThenProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoThenProperty(val interface{}) {
	if err := j.validateSetNoThenPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noThenProperty",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoUnsafeDeclarationMerging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeDeclarationMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoUnsafeDeclarationMerging(val interface{}) {
	if err := j.validateSetNoUnsafeDeclarationMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnsafeDeclarationMerging",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) NoUnsafeNegation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeNegation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetNoUnsafeNegation(val interface{}) {
	if err := j.validateSetNoUnsafeNegationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnsafeNegation",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseAwait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAwait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseAwait(val interface{}) {
	if err := j.validateSetUseAwaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAwait",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseDefaultSwitchClauseLast() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSwitchClauseLast",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseDefaultSwitchClauseLast(val interface{}) {
	if err := j.validateSetUseDefaultSwitchClauseLastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultSwitchClauseLast",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseErrorMessage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseErrorMessage(val interface{}) {
	if err := j.validateSetUseErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useErrorMessage",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseGetterReturn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGetterReturn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseGetterReturn(val interface{}) {
	if err := j.validateSetUseGetterReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGetterReturn",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseIsArray() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIsArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseIsArray(val interface{}) {
	if err := j.validateSetUseIsArrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIsArray",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseNamespaceKeyword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNamespaceKeyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseNamespaceKeyword(val interface{}) {
	if err := j.validateSetUseNamespaceKeywordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNamespaceKeyword",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseNumberToFixedDigitsArgument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useNumberToFixedDigitsArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseNumberToFixedDigitsArgument(val interface{}) {
	if err := j.validateSetUseNumberToFixedDigitsArgumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNumberToFixedDigitsArgument",
		val,
	)
}

func (j *jsiiProxy_ISuspicious) UseValidTypeof() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidTypeof",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuspicious)SetUseValidTypeof(val interface{}) {
	if err := j.validateSetUseValidTypeofParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidTypeof",
		val,
	)
}

