package biomeconfig


// Controls whether Biome keeps, adds, or removes line breaks in Markdown paragraphs.
//
// Manual line breaks are always kept. In Markdown, a manual line break is created by ending a
// line with two spaces or a backslash.
// Experimental.
type ProseWrap string

const (
	// Keep line breaks as written in the source file.
	//
	// This is the default. (preserve)
	// Experimental.
	ProseWrap_PRESERVE ProseWrap = "PRESERVE"
	// Wrap paragraphs to fit the configured `lineWidth`.
	//
	// (always).
	// Experimental.
	ProseWrap_ALWAYS ProseWrap = "ALWAYS"
	// Remove line breaks from paragraphs so that each paragraph is on one line.
	//
	// (never).
	// Experimental.
	ProseWrap_NEVER ProseWrap = "NEVER"
)

