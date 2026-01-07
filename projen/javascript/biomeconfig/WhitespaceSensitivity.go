package biomeconfig


// Whitespace sensitivity for HTML formatting.
//
// The following two cases won't produce the same output:
//
// |                |      html      |    output    |
// | -------------- | :------------: | :----------: |
// | with spaces    | `1<b> 2 </b>3` | 1<b> 2 </b>3 |
// | without spaces |  `1<b>2</b>3`  |  1<b>2</b>3  |
//
// This happens because whitespace is significant in inline elements.
//
// As a consequence of this, the formatter must format blocks that look like this (assume a small line width, <20):
// ```html
// <span>really long content</span>
// ```
// as this, where the content hugs the tags:
// ```html
// <span
// >really long content</span
// >
// ```
//
// Note that this is only necessary for inline elements. Block elements do not have this restriction.
// Experimental.
type WhitespaceSensitivity string

const (
	// The formatter considers whitespace significant for elements that have an "inline" display style by default in browser's user agent style sheets.
	//
	// (css).
	// Experimental.
	WhitespaceSensitivity_CSS WhitespaceSensitivity = "CSS"
	// Leading and trailing whitespace in content is considered significant for all elements.
	//
	// The formatter should leave at least one whitespace character if whitespace is present.
	// Otherwise, if there is no whitespace, it should not add any after `>` or before `<`. In other words, if there's no whitespace, the text content should hug the tags.
	//
	// Example of text hugging the tags:
	// ```html
	// <b
	// >content</b
	// >
	// ``` (strict).
	// Experimental.
	WhitespaceSensitivity_STRICT WhitespaceSensitivity = "STRICT"
	// Whitespace is considered insignificant.
	//
	// The formatter is free to remove or add whitespace as it sees fit. (ignore)
	// Experimental.
	WhitespaceSensitivity_IGNORE WhitespaceSensitivity = "IGNORE"
)

