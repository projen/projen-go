package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type ICorrectness interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Prevent passing of children as props.
	// Experimental.
	NoChildrenProp() interface{}
	// Experimental.
	SetNoChildrenProp(n interface{})
	// Disallow constant expressions in conditions.
	// Experimental.
	NoConstantCondition() interface{}
	// Experimental.
	SetNoConstantCondition(n interface{})
	// Disallow the use of Math.min and Math.max to clamp a value where the result itself is constant.
	// Experimental.
	NoConstantMathMinMaxClamp() interface{}
	// Experimental.
	SetNoConstantMathMinMaxClamp(n interface{})
	// Prevents from having const variables being re-assigned.
	// Experimental.
	NoConstAssign() interface{}
	// Experimental.
	SetNoConstAssign(n interface{})
	// Disallow returning a value from a constructor.
	// Experimental.
	NoConstructorReturn() interface{}
	// Experimental.
	SetNoConstructorReturn(n interface{})
	// Disallow empty character classes in regular expression literals.
	// Experimental.
	NoEmptyCharacterClassInRegex() interface{}
	// Experimental.
	SetNoEmptyCharacterClassInRegex(n interface{})
	// Disallows empty destructuring patterns.
	// Experimental.
	NoEmptyPattern() interface{}
	// Experimental.
	SetNoEmptyPattern(n interface{})
	// Disallow to use unnecessary callback on flatMap.
	// Experimental.
	NoFlatMapIdentity() interface{}
	// Experimental.
	SetNoFlatMapIdentity(n interface{})
	// Disallow calling global object properties as functions.
	// Experimental.
	NoGlobalObjectCalls() interface{}
	// Experimental.
	SetNoGlobalObjectCalls(n interface{})
	// Disallow function and var declarations that are accessible outside their block.
	// Experimental.
	NoInnerDeclarations() interface{}
	// Experimental.
	SetNoInnerDeclarations(n interface{})
	// Ensure that builtins are correctly instantiated.
	// Experimental.
	NoInvalidBuiltinInstantiation() interface{}
	// Experimental.
	SetNoInvalidBuiltinInstantiation(n interface{})
	// Prevents the incorrect use of super() inside classes.
	//
	// It also checks whether a call super() is missing from classes that extends other constructors.
	// Experimental.
	NoInvalidConstructorSuper() interface{}
	// Experimental.
	SetNoInvalidConstructorSuper(n interface{})
	// Disallow non-standard direction values for linear gradient functions.
	// Experimental.
	NoInvalidDirectionInLinearGradient() interface{}
	// Experimental.
	SetNoInvalidDirectionInLinearGradient(n interface{})
	// Disallows invalid named grid areas in CSS Grid Layouts.
	// Experimental.
	NoInvalidGridAreas() interface{}
	// Experimental.
	SetNoInvalidGridAreas(n interface{})
	// Disallow new operators with global non-constructor functions.
	// Experimental.
	NoInvalidNewBuiltin() interface{}
	// Experimental.
	SetNoInvalidNewBuiltin(n interface{})
	// Disallow the use of.
	// Experimental.
	NoInvalidPositionAtImportRule() interface{}
	// Experimental.
	SetNoInvalidPositionAtImportRule(n interface{})
	// Disallow the use of variables and function parameters before their declaration.
	// Experimental.
	NoInvalidUseBeforeDeclaration() interface{}
	// Experimental.
	SetNoInvalidUseBeforeDeclaration(n interface{})
	// Disallow new operators with the Symbol object.
	// Experimental.
	NoNewSymbol() interface{}
	// Experimental.
	SetNoNewSymbol(n interface{})
	// Forbid the use of Node.js builtin modules.
	// Experimental.
	NoNodejsModules() interface{}
	// Experimental.
	SetNoNodejsModules(n interface{})
	// Disallow \8 and \9 escape sequences in string literals.
	// Experimental.
	NoNonoctalDecimalEscape() interface{}
	// Experimental.
	SetNoNonoctalDecimalEscape(n interface{})
	// Disallow literal numbers that lose precision.
	// Experimental.
	NoPrecisionLoss() interface{}
	// Experimental.
	SetNoPrecisionLoss(n interface{})
	// Prevent the usage of the return value of React.render.
	// Experimental.
	NoRenderReturnValue() interface{}
	// Experimental.
	SetNoRenderReturnValue(n interface{})
	// Disallow assignments where both sides are exactly the same.
	// Experimental.
	NoSelfAssign() interface{}
	// Experimental.
	SetNoSelfAssign(n interface{})
	// Disallow returning a value from a setter.
	// Experimental.
	NoSetterReturn() interface{}
	// Experimental.
	SetNoSetterReturn(n interface{})
	// Disallow comparison of expressions modifying the string case with non-compliant value.
	// Experimental.
	NoStringCaseMismatch() interface{}
	// Experimental.
	SetNoStringCaseMismatch(n interface{})
	// Disallow lexical declarations in switch clauses.
	// Experimental.
	NoSwitchDeclarations() interface{}
	// Experimental.
	SetNoSwitchDeclarations(n interface{})
	// Disallow the use of dependencies that aren't specified in the package.json.
	// Experimental.
	NoUndeclaredDependencies() interface{}
	// Experimental.
	SetNoUndeclaredDependencies(n interface{})
	// Prevents the usage of variables that haven't been declared inside the document.
	// Experimental.
	NoUndeclaredVariables() interface{}
	// Experimental.
	SetNoUndeclaredVariables(n interface{})
	// Disallow unknown CSS value functions.
	// Experimental.
	NoUnknownFunction() interface{}
	// Experimental.
	SetNoUnknownFunction(n interface{})
	// Disallow unknown media feature names.
	// Experimental.
	NoUnknownMediaFeatureName() interface{}
	// Experimental.
	SetNoUnknownMediaFeatureName(n interface{})
	// Disallow unknown properties.
	// Experimental.
	NoUnknownProperty() interface{}
	// Experimental.
	SetNoUnknownProperty(n interface{})
	// Disallow unknown CSS units.
	// Experimental.
	NoUnknownUnit() interface{}
	// Experimental.
	SetNoUnknownUnit(n interface{})
	// Disallow unmatchable An+B selectors.
	// Experimental.
	NoUnmatchableAnbSelector() interface{}
	// Experimental.
	SetNoUnmatchableAnbSelector(n interface{})
	// Avoid using unnecessary continue.
	// Experimental.
	NoUnnecessaryContinue() interface{}
	// Experimental.
	SetNoUnnecessaryContinue(n interface{})
	// Disallow unreachable code.
	// Experimental.
	NoUnreachable() interface{}
	// Experimental.
	SetNoUnreachable(n interface{})
	// Ensures the super() constructor is called exactly once on every code  path in a class constructor before this is accessed if the class has a superclass.
	// Experimental.
	NoUnreachableSuper() interface{}
	// Experimental.
	SetNoUnreachableSuper(n interface{})
	// Disallow control flow statements in finally blocks.
	// Experimental.
	NoUnsafeFinally() interface{}
	// Experimental.
	SetNoUnsafeFinally(n interface{})
	// Disallow the use of optional chaining in contexts where the undefined value is not allowed.
	// Experimental.
	NoUnsafeOptionalChaining() interface{}
	// Experimental.
	SetNoUnsafeOptionalChaining(n interface{})
	// Disallow unused function parameters.
	// Experimental.
	NoUnusedFunctionParameters() interface{}
	// Experimental.
	SetNoUnusedFunctionParameters(n interface{})
	// Disallow unused imports.
	// Experimental.
	NoUnusedImports() interface{}
	// Experimental.
	SetNoUnusedImports(n interface{})
	// Disallow unused labels.
	// Experimental.
	NoUnusedLabels() interface{}
	// Experimental.
	SetNoUnusedLabels(n interface{})
	// Disallow unused private class members.
	// Experimental.
	NoUnusedPrivateClassMembers() interface{}
	// Experimental.
	SetNoUnusedPrivateClassMembers(n interface{})
	// Disallow unused variables.
	// Experimental.
	NoUnusedVariables() interface{}
	// Experimental.
	SetNoUnusedVariables(n interface{})
	// This rules prevents void elements (AKA self-closing elements) from having children.
	// Experimental.
	NoVoidElementsWithChildren() interface{}
	// Experimental.
	SetNoVoidElementsWithChildren(n interface{})
	// Disallow returning a value from a function with the return type 'void'.
	// Experimental.
	NoVoidTypeReturn() interface{}
	// Experimental.
	SetNoVoidTypeReturn(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Disallow Array constructors.
	// Experimental.
	UseArrayLiterals() interface{}
	// Experimental.
	SetUseArrayLiterals(u interface{})
	// Enforce all dependencies are correctly specified in a React hook.
	// Experimental.
	UseExhaustiveDependencies() interface{}
	// Experimental.
	SetUseExhaustiveDependencies(u interface{})
	// Enforce that all React hooks are being called from the Top Level component functions.
	// Experimental.
	UseHookAtTopLevel() interface{}
	// Experimental.
	SetUseHookAtTopLevel(u interface{})
	// Enforce file extensions for relative imports.
	// Experimental.
	UseImportExtensions() interface{}
	// Experimental.
	SetUseImportExtensions(u interface{})
	// Require calls to isNaN() when checking for NaN.
	// Experimental.
	UseIsNan() interface{}
	// Experimental.
	SetUseIsNan(u interface{})
	// Disallow missing key props in iterators/collection literals.
	// Experimental.
	UseJsxKeyInIterable() interface{}
	// Experimental.
	SetUseJsxKeyInIterable(u interface{})
	// Enforce "for" loop update clause moving the counter in the right direction.
	// Experimental.
	UseValidForDirection() interface{}
	// Experimental.
	SetUseValidForDirection(u interface{})
	// Require generator functions to contain yield.
	// Experimental.
	UseYield() interface{}
	// Experimental.
	SetUseYield(u interface{})
}

// The jsii proxy for ICorrectness
type jsiiProxy_ICorrectness struct {
	_ byte // padding
}

func (j *jsiiProxy_ICorrectness) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoChildrenProp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noChildrenProp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoChildrenProp(val interface{}) {
	if err := j.validateSetNoChildrenPropParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noChildrenProp",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoConstantCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConstantCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoConstantCondition(val interface{}) {
	if err := j.validateSetNoConstantConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConstantCondition",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoConstantMathMinMaxClamp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConstantMathMinMaxClamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoConstantMathMinMaxClamp(val interface{}) {
	if err := j.validateSetNoConstantMathMinMaxClampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConstantMathMinMaxClamp",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoConstAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConstAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoConstAssign(val interface{}) {
	if err := j.validateSetNoConstAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConstAssign",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoConstructorReturn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noConstructorReturn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoConstructorReturn(val interface{}) {
	if err := j.validateSetNoConstructorReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noConstructorReturn",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoEmptyCharacterClassInRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyCharacterClassInRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoEmptyCharacterClassInRegex(val interface{}) {
	if err := j.validateSetNoEmptyCharacterClassInRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyCharacterClassInRegex",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoEmptyPattern() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noEmptyPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoEmptyPattern(val interface{}) {
	if err := j.validateSetNoEmptyPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noEmptyPattern",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoFlatMapIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noFlatMapIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoFlatMapIdentity(val interface{}) {
	if err := j.validateSetNoFlatMapIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noFlatMapIdentity",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoGlobalObjectCalls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGlobalObjectCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoGlobalObjectCalls(val interface{}) {
	if err := j.validateSetNoGlobalObjectCallsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGlobalObjectCalls",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInnerDeclarations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInnerDeclarations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInnerDeclarations(val interface{}) {
	if err := j.validateSetNoInnerDeclarationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInnerDeclarations",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidBuiltinInstantiation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidBuiltinInstantiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidBuiltinInstantiation(val interface{}) {
	if err := j.validateSetNoInvalidBuiltinInstantiationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidBuiltinInstantiation",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidConstructorSuper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidConstructorSuper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidConstructorSuper(val interface{}) {
	if err := j.validateSetNoInvalidConstructorSuperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidConstructorSuper",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidDirectionInLinearGradient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidDirectionInLinearGradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidDirectionInLinearGradient(val interface{}) {
	if err := j.validateSetNoInvalidDirectionInLinearGradientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidDirectionInLinearGradient",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidGridAreas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidGridAreas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidGridAreas(val interface{}) {
	if err := j.validateSetNoInvalidGridAreasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidGridAreas",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidNewBuiltin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidNewBuiltin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidNewBuiltin(val interface{}) {
	if err := j.validateSetNoInvalidNewBuiltinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidNewBuiltin",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidPositionAtImportRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidPositionAtImportRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidPositionAtImportRule(val interface{}) {
	if err := j.validateSetNoInvalidPositionAtImportRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidPositionAtImportRule",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoInvalidUseBeforeDeclaration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInvalidUseBeforeDeclaration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoInvalidUseBeforeDeclaration(val interface{}) {
	if err := j.validateSetNoInvalidUseBeforeDeclarationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInvalidUseBeforeDeclaration",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoNewSymbol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNewSymbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoNewSymbol(val interface{}) {
	if err := j.validateSetNoNewSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNewSymbol",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoNodejsModules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNodejsModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoNodejsModules(val interface{}) {
	if err := j.validateSetNoNodejsModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNodejsModules",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoNonoctalDecimalEscape() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNonoctalDecimalEscape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoNonoctalDecimalEscape(val interface{}) {
	if err := j.validateSetNoNonoctalDecimalEscapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNonoctalDecimalEscape",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoPrecisionLoss() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPrecisionLoss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoPrecisionLoss(val interface{}) {
	if err := j.validateSetNoPrecisionLossParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPrecisionLoss",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoRenderReturnValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRenderReturnValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoRenderReturnValue(val interface{}) {
	if err := j.validateSetNoRenderReturnValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRenderReturnValue",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoSelfAssign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSelfAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoSelfAssign(val interface{}) {
	if err := j.validateSetNoSelfAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSelfAssign",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoSetterReturn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSetterReturn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoSetterReturn(val interface{}) {
	if err := j.validateSetNoSetterReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSetterReturn",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoStringCaseMismatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStringCaseMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoStringCaseMismatch(val interface{}) {
	if err := j.validateSetNoStringCaseMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStringCaseMismatch",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoSwitchDeclarations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSwitchDeclarations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoSwitchDeclarations(val interface{}) {
	if err := j.validateSetNoSwitchDeclarationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSwitchDeclarations",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUndeclaredDependencies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUndeclaredDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUndeclaredDependencies(val interface{}) {
	if err := j.validateSetNoUndeclaredDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUndeclaredDependencies",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUndeclaredVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUndeclaredVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUndeclaredVariables(val interface{}) {
	if err := j.validateSetNoUndeclaredVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUndeclaredVariables",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnknownFunction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnknownFunction(val interface{}) {
	if err := j.validateSetNoUnknownFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownFunction",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnknownMediaFeatureName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownMediaFeatureName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnknownMediaFeatureName(val interface{}) {
	if err := j.validateSetNoUnknownMediaFeatureNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownMediaFeatureName",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnknownProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnknownProperty(val interface{}) {
	if err := j.validateSetNoUnknownPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownProperty",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnknownUnit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnknownUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnknownUnit(val interface{}) {
	if err := j.validateSetNoUnknownUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnknownUnit",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnmatchableAnbSelector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnmatchableAnbSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnmatchableAnbSelector(val interface{}) {
	if err := j.validateSetNoUnmatchableAnbSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnmatchableAnbSelector",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnnecessaryContinue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnnecessaryContinue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnnecessaryContinue(val interface{}) {
	if err := j.validateSetNoUnnecessaryContinueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnnecessaryContinue",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnreachable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnreachable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnreachable(val interface{}) {
	if err := j.validateSetNoUnreachableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnreachable",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnreachableSuper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnreachableSuper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnreachableSuper(val interface{}) {
	if err := j.validateSetNoUnreachableSuperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnreachableSuper",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnsafeFinally() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeFinally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnsafeFinally(val interface{}) {
	if err := j.validateSetNoUnsafeFinallyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnsafeFinally",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnsafeOptionalChaining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeOptionalChaining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnsafeOptionalChaining(val interface{}) {
	if err := j.validateSetNoUnsafeOptionalChainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnsafeOptionalChaining",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnusedFunctionParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnusedFunctionParameters(val interface{}) {
	if err := j.validateSetNoUnusedFunctionParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedFunctionParameters",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnusedImports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnusedImports(val interface{}) {
	if err := j.validateSetNoUnusedImportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedImports",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnusedLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnusedLabels(val interface{}) {
	if err := j.validateSetNoUnusedLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedLabels",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnusedPrivateClassMembers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedPrivateClassMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnusedPrivateClassMembers(val interface{}) {
	if err := j.validateSetNoUnusedPrivateClassMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedPrivateClassMembers",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoUnusedVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnusedVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoUnusedVariables(val interface{}) {
	if err := j.validateSetNoUnusedVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnusedVariables",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoVoidElementsWithChildren() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noVoidElementsWithChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoVoidElementsWithChildren(val interface{}) {
	if err := j.validateSetNoVoidElementsWithChildrenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noVoidElementsWithChildren",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) NoVoidTypeReturn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noVoidTypeReturn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetNoVoidTypeReturn(val interface{}) {
	if err := j.validateSetNoVoidTypeReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noVoidTypeReturn",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseArrayLiterals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useArrayLiterals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseArrayLiterals(val interface{}) {
	if err := j.validateSetUseArrayLiteralsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useArrayLiterals",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseExhaustiveDependencies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExhaustiveDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseExhaustiveDependencies(val interface{}) {
	if err := j.validateSetUseExhaustiveDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExhaustiveDependencies",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseHookAtTopLevel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useHookAtTopLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseHookAtTopLevel(val interface{}) {
	if err := j.validateSetUseHookAtTopLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHookAtTopLevel",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseImportExtensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useImportExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseImportExtensions(val interface{}) {
	if err := j.validateSetUseImportExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useImportExtensions",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseIsNan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIsNan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseIsNan(val interface{}) {
	if err := j.validateSetUseIsNanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIsNan",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseJsxKeyInIterable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useJsxKeyInIterable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseJsxKeyInIterable(val interface{}) {
	if err := j.validateSetUseJsxKeyInIterableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useJsxKeyInIterable",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseValidForDirection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidForDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseValidForDirection(val interface{}) {
	if err := j.validateSetUseValidForDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidForDirection",
		val,
	)
}

func (j *jsiiProxy_ICorrectness) UseYield() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useYield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICorrectness)SetUseYield(val interface{}) {
	if err := j.validateSetUseYieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useYield",
		val,
	)
}

