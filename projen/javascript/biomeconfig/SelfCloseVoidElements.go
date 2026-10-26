package biomeconfig


// Controls whether HTML void elements such as `<img>` and `<input>` include a slash before `>`.
// Experimental.
type SelfCloseVoidElements string

const (
	// Removes the slash, for example `<img src="image.png">`. (never).
	// Experimental.
	SelfCloseVoidElements_NEVER SelfCloseVoidElements = "NEVER"
	// Adds the slash, for example `<img src="image.png" />`. (always).
	// Experimental.
	SelfCloseVoidElements_ALWAYS SelfCloseVoidElements = "ALWAYS"
)

