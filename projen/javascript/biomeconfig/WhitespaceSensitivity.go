package biomeconfig


// Whitespace sensitivity for HTML formatting.
//
// The following two cases won't produce the same output:
//
// |                |      html      |    output    | | -------------- | :------------: | :----------: | | with spaces    | `1<b> 2 </b>3` | 1<b> 2 </b>3 | | without spaces |  `1<b>2</b>3`  |  1<b>2</b>3  |
//
// This happens because whitespace is significant in inline elements.
//
// As a consequence of this, the formatter must format blocks that look like this (assume a small line width, <20): ```html <span>really long content</span> ``` as this, where the content hugs the tags: ```html <span >really long content</span > ```
//
// Note that this is only necessary for inline elements. Block elements do not have this restriction.
// Experimental.
type WhitespaceSensitivity string

const (
	// css.
	// Experimental.
	WhitespaceSensitivity_CSS WhitespaceSensitivity = "CSS"
	// strict.
	// Experimental.
	WhitespaceSensitivity_STRICT WhitespaceSensitivity = "STRICT"
	// ignore.
	// Experimental.
	WhitespaceSensitivity_IGNORE WhitespaceSensitivity = "IGNORE"
)

