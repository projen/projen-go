package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type IA11y interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Enforce that the accessKey attribute is not used on any HTML element.
	// Experimental.
	NoAccessKey() interface{}
	// Experimental.
	SetNoAccessKey(n interface{})
	// Enforce that aria-hidden="true" is not set on focusable elements.
	// Experimental.
	NoAriaHiddenOnFocusable() interface{}
	// Experimental.
	SetNoAriaHiddenOnFocusable(n interface{})
	// Enforce that elements that do not support ARIA roles, states, and properties do not have those attributes.
	// Experimental.
	NoAriaUnsupportedElements() interface{}
	// Experimental.
	SetNoAriaUnsupportedElements(n interface{})
	// Enforce that autoFocus prop is not used on elements.
	// Experimental.
	NoAutofocus() interface{}
	// Experimental.
	SetNoAutofocus(n interface{})
	// Disallow target="_blank" attribute without rel="noreferrer".
	// Experimental.
	NoBlankTarget() interface{}
	// Experimental.
	SetNoBlankTarget(n interface{})
	// Enforces that no distracting elements are used.
	// Experimental.
	NoDistractingElements() interface{}
	// Experimental.
	SetNoDistractingElements(n interface{})
	// The scope prop should be used only on \<th> elements.
	// Experimental.
	NoHeaderScope() interface{}
	// Experimental.
	SetNoHeaderScope(n interface{})
	// Enforce that non-interactive ARIA roles are not assigned to interactive HTML elements.
	// Experimental.
	NoInteractiveElementToNoninteractiveRole() interface{}
	// Experimental.
	SetNoInteractiveElementToNoninteractiveRole(n interface{})
	// Enforce that a label element or component has a text label and an associated input.
	// Experimental.
	NoLabelWithoutControl() interface{}
	// Experimental.
	SetNoLabelWithoutControl(n interface{})
	// Enforce that interactive ARIA roles are not assigned to non-interactive HTML elements.
	// Experimental.
	NoNoninteractiveElementToInteractiveRole() interface{}
	// Experimental.
	SetNoNoninteractiveElementToInteractiveRole(n interface{})
	// Enforce that tabIndex is not assigned to non-interactive HTML elements.
	// Experimental.
	NoNoninteractiveTabindex() interface{}
	// Experimental.
	SetNoNoninteractiveTabindex(n interface{})
	// Prevent the usage of positive integers on tabIndex property.
	// Experimental.
	NoPositiveTabindex() interface{}
	// Experimental.
	SetNoPositiveTabindex(n interface{})
	// Enforce img alt prop does not contain the word "image", "picture", or "photo".
	// Experimental.
	NoRedundantAlt() interface{}
	// Experimental.
	SetNoRedundantAlt(n interface{})
	// Enforce explicit role property is not the same as implicit/default role property on an element.
	// Experimental.
	NoRedundantRoles() interface{}
	// Experimental.
	SetNoRedundantRoles(n interface{})
	// Enforces the usage of the title element for the svg element.
	// Experimental.
	NoSvgWithoutTitle() interface{}
	// Experimental.
	SetNoSvgWithoutTitle(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Enforce that all elements that require alternative text have meaningful information to relay back to the end user.
	// Experimental.
	UseAltText() interface{}
	// Experimental.
	SetUseAltText(u interface{})
	// Enforce that anchors have content and that the content is accessible to screen readers.
	// Experimental.
	UseAnchorContent() interface{}
	// Experimental.
	SetUseAnchorContent(u interface{})
	// Enforce that tabIndex is assigned to non-interactive HTML elements with aria-activedescendant.
	// Experimental.
	UseAriaActivedescendantWithTabindex() interface{}
	// Experimental.
	SetUseAriaActivedescendantWithTabindex(u interface{})
	// Enforce that elements with ARIA roles must have all required ARIA attributes for that role.
	// Experimental.
	UseAriaPropsForRole() interface{}
	// Experimental.
	SetUseAriaPropsForRole(u interface{})
	// Enforces the usage of the attribute type for the element button.
	// Experimental.
	UseButtonType() interface{}
	// Experimental.
	SetUseButtonType(u interface{})
	// Elements with an interactive role and interaction handlers must be focusable.
	// Experimental.
	UseFocusableInteractive() interface{}
	// Experimental.
	SetUseFocusableInteractive(u interface{})
	// Disallow a missing generic family keyword within font families.
	// Experimental.
	UseGenericFontNames() interface{}
	// Experimental.
	SetUseGenericFontNames(u interface{})
	// Enforce that heading elements (h1, h2, etc.) have content and that the content is accessible to screen readers. Accessible means that it is not hidden using the aria-hidden prop.
	// Experimental.
	UseHeadingContent() interface{}
	// Experimental.
	SetUseHeadingContent(u interface{})
	// Enforce that html element has lang attribute.
	// Experimental.
	UseHtmlLang() interface{}
	// Experimental.
	SetUseHtmlLang(u interface{})
	// Enforces the usage of the attribute title for the element iframe.
	// Experimental.
	UseIframeTitle() interface{}
	// Experimental.
	SetUseIframeTitle(u interface{})
	// Enforce onClick is accompanied by at least one of the following: onKeyUp, onKeyDown, onKeyPress.
	// Experimental.
	UseKeyWithClickEvents() interface{}
	// Experimental.
	SetUseKeyWithClickEvents(u interface{})
	// Enforce onMouseOver / onMouseOut are accompanied by onFocus / onBlur.
	// Experimental.
	UseKeyWithMouseEvents() interface{}
	// Experimental.
	SetUseKeyWithMouseEvents(u interface{})
	// Enforces that audio and video elements must have a track for captions.
	// Experimental.
	UseMediaCaption() interface{}
	// Experimental.
	SetUseMediaCaption(u interface{})
	// It detects the use of role attributes in JSX elements and suggests using semantic elements instead.
	// Experimental.
	UseSemanticElements() interface{}
	// Experimental.
	SetUseSemanticElements(u interface{})
	// Enforce that all anchors are valid, and they are navigable elements.
	// Experimental.
	UseValidAnchor() interface{}
	// Experimental.
	SetUseValidAnchor(u interface{})
	// Ensures that ARIA properties aria-* are all valid.
	// Experimental.
	UseValidAriaProps() interface{}
	// Experimental.
	SetUseValidAriaProps(u interface{})
	// Elements with ARIA roles must use a valid, non-abstract ARIA role.
	// Experimental.
	UseValidAriaRole() interface{}
	// Experimental.
	SetUseValidAriaRole(u interface{})
	// Enforce that ARIA state and property values are valid.
	// Experimental.
	UseValidAriaValues() interface{}
	// Experimental.
	SetUseValidAriaValues(u interface{})
	// Ensure that the attribute passed to the lang attribute is a correct ISO language and/or country.
	// Experimental.
	UseValidLang() interface{}
	// Experimental.
	SetUseValidLang(u interface{})
}

// The jsii proxy for IA11y
type jsiiProxy_IA11y struct {
	_ byte // padding
}

func (j *jsiiProxy_IA11y) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoAccessKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoAccessKey(val interface{}) {
	if err := j.validateSetNoAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAccessKey",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoAriaHiddenOnFocusable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAriaHiddenOnFocusable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoAriaHiddenOnFocusable(val interface{}) {
	if err := j.validateSetNoAriaHiddenOnFocusableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAriaHiddenOnFocusable",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoAriaUnsupportedElements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAriaUnsupportedElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoAriaUnsupportedElements(val interface{}) {
	if err := j.validateSetNoAriaUnsupportedElementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAriaUnsupportedElements",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoAutofocus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAutofocus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoAutofocus(val interface{}) {
	if err := j.validateSetNoAutofocusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAutofocus",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoBlankTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noBlankTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoBlankTarget(val interface{}) {
	if err := j.validateSetNoBlankTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noBlankTarget",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoDistractingElements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDistractingElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoDistractingElements(val interface{}) {
	if err := j.validateSetNoDistractingElementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDistractingElements",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoHeaderScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHeaderScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoHeaderScope(val interface{}) {
	if err := j.validateSetNoHeaderScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHeaderScope",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoInteractiveElementToNoninteractiveRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noInteractiveElementToNoninteractiveRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoInteractiveElementToNoninteractiveRole(val interface{}) {
	if err := j.validateSetNoInteractiveElementToNoninteractiveRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noInteractiveElementToNoninteractiveRole",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoLabelWithoutControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noLabelWithoutControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoLabelWithoutControl(val interface{}) {
	if err := j.validateSetNoLabelWithoutControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noLabelWithoutControl",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoNoninteractiveElementToInteractiveRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNoninteractiveElementToInteractiveRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoNoninteractiveElementToInteractiveRole(val interface{}) {
	if err := j.validateSetNoNoninteractiveElementToInteractiveRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNoninteractiveElementToInteractiveRole",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoNoninteractiveTabindex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNoninteractiveTabindex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoNoninteractiveTabindex(val interface{}) {
	if err := j.validateSetNoNoninteractiveTabindexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNoninteractiveTabindex",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoPositiveTabindex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPositiveTabindex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoPositiveTabindex(val interface{}) {
	if err := j.validateSetNoPositiveTabindexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPositiveTabindex",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoRedundantAlt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRedundantAlt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoRedundantAlt(val interface{}) {
	if err := j.validateSetNoRedundantAltParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRedundantAlt",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoRedundantRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRedundantRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoRedundantRoles(val interface{}) {
	if err := j.validateSetNoRedundantRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRedundantRoles",
		val,
	)
}

func (j *jsiiProxy_IA11y) NoSvgWithoutTitle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSvgWithoutTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetNoSvgWithoutTitle(val interface{}) {
	if err := j.validateSetNoSvgWithoutTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSvgWithoutTitle",
		val,
	)
}

func (j *jsiiProxy_IA11y) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseAltText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAltText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseAltText(val interface{}) {
	if err := j.validateSetUseAltTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAltText",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseAnchorContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAnchorContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseAnchorContent(val interface{}) {
	if err := j.validateSetUseAnchorContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAnchorContent",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseAriaActivedescendantWithTabindex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAriaActivedescendantWithTabindex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseAriaActivedescendantWithTabindex(val interface{}) {
	if err := j.validateSetUseAriaActivedescendantWithTabindexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAriaActivedescendantWithTabindex",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseAriaPropsForRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAriaPropsForRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseAriaPropsForRole(val interface{}) {
	if err := j.validateSetUseAriaPropsForRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAriaPropsForRole",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseButtonType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useButtonType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseButtonType(val interface{}) {
	if err := j.validateSetUseButtonTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useButtonType",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseFocusableInteractive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFocusableInteractive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseFocusableInteractive(val interface{}) {
	if err := j.validateSetUseFocusableInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFocusableInteractive",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseGenericFontNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useGenericFontNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseGenericFontNames(val interface{}) {
	if err := j.validateSetUseGenericFontNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGenericFontNames",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseHeadingContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useHeadingContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseHeadingContent(val interface{}) {
	if err := j.validateSetUseHeadingContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHeadingContent",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseHtmlLang() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useHtmlLang",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseHtmlLang(val interface{}) {
	if err := j.validateSetUseHtmlLangParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHtmlLang",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseIframeTitle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIframeTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseIframeTitle(val interface{}) {
	if err := j.validateSetUseIframeTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIframeTitle",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseKeyWithClickEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useKeyWithClickEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseKeyWithClickEvents(val interface{}) {
	if err := j.validateSetUseKeyWithClickEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useKeyWithClickEvents",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseKeyWithMouseEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useKeyWithMouseEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseKeyWithMouseEvents(val interface{}) {
	if err := j.validateSetUseKeyWithMouseEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useKeyWithMouseEvents",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseMediaCaption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMediaCaption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseMediaCaption(val interface{}) {
	if err := j.validateSetUseMediaCaptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMediaCaption",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseSemanticElements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSemanticElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseSemanticElements(val interface{}) {
	if err := j.validateSetUseSemanticElementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSemanticElements",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseValidAnchor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidAnchor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseValidAnchor(val interface{}) {
	if err := j.validateSetUseValidAnchorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidAnchor",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseValidAriaProps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidAriaProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseValidAriaProps(val interface{}) {
	if err := j.validateSetUseValidAriaPropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidAriaProps",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseValidAriaRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidAriaRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseValidAriaRole(val interface{}) {
	if err := j.validateSetUseValidAriaRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidAriaRole",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseValidAriaValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidAriaValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseValidAriaValues(val interface{}) {
	if err := j.validateSetUseValidAriaValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidAriaValues",
		val,
	)
}

func (j *jsiiProxy_IA11y) UseValidLang() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useValidLang",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IA11y)SetUseValidLang(val interface{}) {
	if err := j.validateSetUseValidLangParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useValidLang",
		val,
	)
}

