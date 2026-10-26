package biomeconfig


// Options that change how the Markdown parser behaves.
// Experimental.
type MarkdownParserConfiguration struct {
	// Recognizes bold, italic, and strikethrough markers placed directly next to Chinese, Japanese, or Korean text. Defaults to `false`.
	//
	// Standard Markdown doesn't treat `**テスト。**テスト` as bold text, because the closing `**`
	// is between a punctuation mark and a letter with no space around it. Languages that don't
	// put spaces between words run into this often. When enabled, Biome follows the
	// [CommonMark CJK-friendly amendments](https://github.com/tats-u/markdown-cjk-friendly/blob/main/specification.md)
	// and treats such text as bold or italic. The option applies to `*`, `_`, and `~~` markers.
	// After a variation selector (an invisible character that picks a glyph style for the
	// character before it), Biome classifies a marker by the character before the selector.
	//
	// The option can also change markers that aren't next to Chinese, Japanese, or Korean text,
	// because a marker that can now open or close formatting can take another marker's partner
	// in the same paragraph, heading, or table cell. For example, standard Markdown shows
	// `**テスト。**テスト a**` with `テスト a` in bold. With the option enabled, `テスト。` is bold
	// and the last `**` is shown as written.
	//
	// The formatter and the linter both use this reading of the document. Enable the option only
	// when the tool that renders your Markdown also supports these amendments. Otherwise, Biome
	// treats as bold some text that your renderer shows with literal `**`, and formatting can
	// change what your renderer shows. For example, the option makes `__` after a variation
	// selector a bold marker, and the formatter rewrites it as `**`, which standard Markdown
	// shows as bold even where it showed the original `__` as written.
	// Default: false`.
	//
	// Experimental.
	CjkFriendlyEmphasis *bool `field:"optional" json:"cjkFriendlyEmphasis" yaml:"cjkFriendlyEmphasis"`
	// Enables parsing frontmatter at the start of the file.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	Frontmatter *bool `field:"optional" json:"frontmatter" yaml:"frontmatter"`
	// Enables GitHub Flavored Markdown extensions.
	//
	// Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	Gfm *bool `field:"optional" json:"gfm" yaml:"gfm"`
}

