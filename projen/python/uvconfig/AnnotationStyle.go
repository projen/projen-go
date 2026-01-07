package uvconfig


// Indicate the style of annotation comments, used to indicate the dependencies that requested each package.
// Experimental.
type AnnotationStyle string

const (
	// Render the annotations on a single, comma-separated line.
	//
	// (line).
	// Experimental.
	AnnotationStyle_LINE AnnotationStyle = "LINE"
	// Render each annotation on its own line.
	//
	// (split).
	// Experimental.
	AnnotationStyle_SPLIT AnnotationStyle = "SPLIT"
)

