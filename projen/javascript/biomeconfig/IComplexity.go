package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type IComplexity interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Disallow primitive type aliases and misleading types.
	// Experimental.
	NoBannedTypes() interface{}
	// Experimental.
	SetNoBannedTypes(n interface{})
	// Disallow empty type parameters in type aliases and interfaces.
	// Experimental.
	NoEmptyTypeParameters() interface{}
	// Experimental.
	SetNoEmptyTypeParameters(n interface{})
	// Disallow functions that exceed a given Cognitive Complexity score.
	// Experimental.
	NoExcessiveCognitiveComplexity() interface{}
	// Experimental.
	SetNoExcessiveCognitiveComplexity(n interface{})
	// This rule enforces a maximum depth to nested describe() in test files.
	// Experimental.
	NoExcessiveNestedTestSuites() interface{}
	// Experimental.
	SetNoExcessiveNestedTestSuites(n interface{})
	// Disallow unnecessary boolean casts.
	// Experimental.
	NoExtraBooleanCast() interface{}
	// Experimental.
	SetNoExtraBooleanCast(n interface{})
	// Prefer for...of statement instead of Array.forEach.
	// Experimental.
	NoForEach() interface{}
	// Experimental.
	SetNoForEach(n interface{})
	// Disallow unclear usage of consecutive space characters in regular expression literals.
	// Experimental.
	NoMultipleSpacesInRegularExpressionLiterals() interface{}
	// Experimental.
	SetNoMultipleSpacesInRegularExpressionLiterals(n interface{})
	// This rule reports when a class has no non-static members, such as for a class used exclusively as a static namespace.
	// Experimental.
	NoStaticOnlyClass() interface{}
	// Experimental.
	SetNoStaticOnlyClass(n interface{})
	// Disallow this and super in static contexts.
	// Experimental.
	NoThisInStatic() interface{}
	// Experimental.
	SetNoThisInStatic(n interface{})
	// Disallow unnecessary catch clauses.
	// Experimental.
	NoUselessCatch() interface{}
	// Experimental.
	SetNoUselessCatch(n interface{})
	// Disallow unnecessary constructors.
	// Experimental.
	NoUselessConstructor() interface{}
	// Experimental.
	SetNoUselessConstructor(n interface{})
	// Disallow empty exports that don't change anything in a module file.
	// Experimental.
	NoUselessEmptyExport() interface{}
	// Experimental.
	SetNoUselessEmptyExport(n interface{})
	// Disallow unnecessary fragments.
	// Experimental.
	NoUselessFragments() interface{}
	// Experimental.
	SetNoUselessFragments(n interface{})
	// Disallow unnecessary labels.
	// Experimental.
	NoUselessLabel() interface{}
	// Experimental.
	SetNoUselessLabel(n interface{})
	// Disallow unnecessary nested block statements.
	// Experimental.
	NoUselessLoneBlockStatements() interface{}
	// Experimental.
	SetNoUselessLoneBlockStatements(n interface{})
	// Disallow renaming import, export, and destructured assignments to the same name.
	// Experimental.
	NoUselessRename() interface{}
	// Experimental.
	SetNoUselessRename(n interface{})
	// Disallow unnecessary concatenation of string or template literals.
	// Experimental.
	NoUselessStringConcat() interface{}
	// Experimental.
	SetNoUselessStringConcat(n interface{})
	// Disallow useless case in switch statements.
	// Experimental.
	NoUselessSwitchCase() interface{}
	// Experimental.
	SetNoUselessSwitchCase(n interface{})
	// Disallow ternary operators when simpler alternatives exist.
	// Experimental.
	NoUselessTernary() interface{}
	// Experimental.
	SetNoUselessTernary(n interface{})
	// Disallow useless this aliasing.
	// Experimental.
	NoUselessThisAlias() interface{}
	// Experimental.
	SetNoUselessThisAlias(n interface{})
	// Disallow using any or unknown as type constraint.
	// Experimental.
	NoUselessTypeConstraint() interface{}
	// Experimental.
	SetNoUselessTypeConstraint(n interface{})
	// Disallow initializing variables to undefined.
	// Experimental.
	NoUselessUndefinedInitialization() interface{}
	// Experimental.
	SetNoUselessUndefinedInitialization(n interface{})
	// Disallow the use of void operators, which is not a familiar operator.
	// Experimental.
	NoVoid() interface{}
	// Experimental.
	SetNoVoid(n interface{})
	// Disallow with statements in non-strict contexts.
	// Experimental.
	NoWith() interface{}
	// Experimental.
	SetNoWith(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Use arrow functions over function expressions.
	// Experimental.
	UseArrowFunction() interface{}
	// Experimental.
	SetUseArrowFunction(u interface{})
	// Use Date.now() to get the number of milliseconds since the Unix Epoch.
	// Experimental.
	UseDateNow() interface{}
	// Experimental.
	SetUseDateNow(u interface{})
	// Promotes the use of .flatMap() when map().flat() are used together.
	// Experimental.
	UseFlatMap() interface{}
	// Experimental.
	SetUseFlatMap(u interface{})
	// Enforce the usage of a literal access to properties over computed property access.
	// Experimental.
	UseLiteralKeys() interface{}
	// Experimental.
	SetUseLiteralKeys(u interface{})
	// Enforce using concise optional chain instead of chained logical expressions.
	// Experimental.
	UseOptionalChain() interface{}
	// Experimental.
	SetUseOptionalChain(u interface{})
	// Enforce the use of the regular expression literals instead of the RegExp constructor if possible.
	// Experimental.
	UseRegexLiterals() interface{}
	// Experimental.
	SetUseRegexLiterals(u interface{})
	// Disallow number literal object member names which are not base10 or uses underscore as separator.
	// Experimental.
	UseSimpleNumberKeys() interface{}
	// Experimental.
	SetUseSimpleNumberKeys(u interface{})
	// Discard redundant terms from logical expressions.
	// Experimental.
	UseSimplifiedLogicExpression() interface{}
	// Experimental.
	SetUseSimplifiedLogicExpression(u interface{})
}

// The jsii proxy for IComplexity
type jsiiProxy_IComplexity struct {
	_ byte // padding
}

func (j *jsiiProxy_IComplexity) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoBannedTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noBannedTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoBannedTypes(val interface{}) {
	if err := j.validateSetNoBannedTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noBannedTypes",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoEmptyTypeParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyTypeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoEmptyTypeParameters(val interface{}) {
	if err := j.validateSetNoEmptyTypeParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyTypeParameters",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoExcessiveCognitiveComplexity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExcessiveCognitiveComplexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoExcessiveCognitiveComplexity(val interface{}) {
	if err := j.validateSetNoExcessiveCognitiveComplexityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExcessiveCognitiveComplexity",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoExcessiveNestedTestSuites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExcessiveNestedTestSuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoExcessiveNestedTestSuites(val interface{}) {
	if err := j.validateSetNoExcessiveNestedTestSuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExcessiveNestedTestSuites",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoExtraBooleanCast() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExtraBooleanCast",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoExtraBooleanCast(val interface{}) {
	if err := j.validateSetNoExtraBooleanCastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExtraBooleanCast",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoForEach() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noForEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoForEach(val interface{}) {
	if err := j.validateSetNoForEachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noForEach",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoMultipleSpacesInRegularExpressionLiterals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noMultipleSpacesInRegularExpressionLiterals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoMultipleSpacesInRegularExpressionLiterals(val interface{}) {
	if err := j.validateSetNoMultipleSpacesInRegularExpressionLiteralsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noMultipleSpacesInRegularExpressionLiterals",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoStaticOnlyClass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStaticOnlyClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoStaticOnlyClass(val interface{}) {
	if err := j.validateSetNoStaticOnlyClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStaticOnlyClass",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoThisInStatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noThisInStatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoThisInStatic(val interface{}) {
	if err := j.validateSetNoThisInStaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noThisInStatic",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessCatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessCatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessCatch(val interface{}) {
	if err := j.validateSetNoUselessCatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessCatch",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessConstructor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessConstructor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessConstructor(val interface{}) {
	if err := j.validateSetNoUselessConstructorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessConstructor",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessEmptyExport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessEmptyExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessEmptyExport(val interface{}) {
	if err := j.validateSetNoUselessEmptyExportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessEmptyExport",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessFragments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessFragments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessFragments(val interface{}) {
	if err := j.validateSetNoUselessFragmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessFragments",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessLabel(val interface{}) {
	if err := j.validateSetNoUselessLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessLabel",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessLoneBlockStatements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessLoneBlockStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessLoneBlockStatements(val interface{}) {
	if err := j.validateSetNoUselessLoneBlockStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessLoneBlockStatements",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessRename() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessRename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessRename(val interface{}) {
	if err := j.validateSetNoUselessRenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessRename",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessStringConcat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessStringConcat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessStringConcat(val interface{}) {
	if err := j.validateSetNoUselessStringConcatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessStringConcat",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessSwitchCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessSwitchCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessSwitchCase(val interface{}) {
	if err := j.validateSetNoUselessSwitchCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessSwitchCase",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessTernary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessTernary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessTernary(val interface{}) {
	if err := j.validateSetNoUselessTernaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessTernary",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessThisAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessThisAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessThisAlias(val interface{}) {
	if err := j.validateSetNoUselessThisAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessThisAlias",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessTypeConstraint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessTypeConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessTypeConstraint(val interface{}) {
	if err := j.validateSetNoUselessTypeConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessTypeConstraint",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoUselessUndefinedInitialization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUselessUndefinedInitialization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoUselessUndefinedInitialization(val interface{}) {
	if err := j.validateSetNoUselessUndefinedInitializationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUselessUndefinedInitialization",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoVoid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noVoid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoVoid(val interface{}) {
	if err := j.validateSetNoVoidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noVoid",
		val,
	)
}

func (j *jsiiProxy_IComplexity) NoWith() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetNoWith(val interface{}) {
	if err := j.validateSetNoWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noWith",
		val,
	)
}

func (j *jsiiProxy_IComplexity) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseArrowFunction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useArrowFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseArrowFunction(val interface{}) {
	if err := j.validateSetUseArrowFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useArrowFunction",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseDateNow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDateNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseDateNow(val interface{}) {
	if err := j.validateSetUseDateNowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDateNow",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseFlatMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFlatMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseFlatMap(val interface{}) {
	if err := j.validateSetUseFlatMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFlatMap",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseLiteralKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLiteralKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseLiteralKeys(val interface{}) {
	if err := j.validateSetUseLiteralKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLiteralKeys",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseOptionalChain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOptionalChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseOptionalChain(val interface{}) {
	if err := j.validateSetUseOptionalChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOptionalChain",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseRegexLiterals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRegexLiterals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseRegexLiterals(val interface{}) {
	if err := j.validateSetUseRegexLiteralsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRegexLiterals",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseSimpleNumberKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSimpleNumberKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseSimpleNumberKeys(val interface{}) {
	if err := j.validateSetUseSimpleNumberKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSimpleNumberKeys",
		val,
	)
}

func (j *jsiiProxy_IComplexity) UseSimplifiedLogicExpression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSimplifiedLogicExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexity)SetUseSimplifiedLogicExpression(val interface{}) {
	if err := j.validateSetUseSimplifiedLogicExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSimplifiedLogicExpression",
		val,
	)
}

