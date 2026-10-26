package biomeconfig


// Controls how the formatter treats whitespace around text and child elements in HTML, Vue, Svelte, and Astro markup.
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
	// Treats whitespace as significant for elements that browsers display inline by default.
	//
	// (css).
	// Experimental.
	WhitespaceSensitivity_CSS WhitespaceSensitivity = "CSS"
	// Treats leading and trailing whitespace as significant for every element.
	//
	// Existing whitespace
	// is preserved as at least one character and absent whitespace is not added.
	//
	// Example of text hugging the tags:
	// ```html
	// <b
	// >content</b
	// >
	// ``` (strict).
	// Experimental.
	WhitespaceSensitivity_STRICT WhitespaceSensitivity = "STRICT"
	// Treats whitespace as insignificant, allowing the formatter to add or remove it.
	//
	// Use this only
	// when whitespace cannot affect rendered output. (ignore)
	// Experimental.
	WhitespaceSensitivity_IGNORE WhitespaceSensitivity = "IGNORE"
)

