package javascript


// Experimental.
type EmbeddedLanguageFormatting string

const (
	// Format embedded code if Prettier can automatically identify it.
	// Experimental.
	EmbeddedLanguageFormatting_AUTO EmbeddedLanguageFormatting = "AUTO"
	// Never automatically format embedded code.
	// Experimental.
	EmbeddedLanguageFormatting_OFF EmbeddedLanguageFormatting = "OFF"
)

