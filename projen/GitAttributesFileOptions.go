package projen


// Options for `GitAttributesFile`.
// Experimental.
type GitAttributesFileOptions struct {
	// The default end of line character for text files.
	//
	// endOfLine it's useful to keep the same end of line between Windows and Unix operative systems for git checking/checkout operations. Hence, it can avoid simple repository mutations consisting only of changes in the end of line characters. It will be set in the first line of the .gitattributes file to make it the first match with high priority but it can be overriden in a later line. Can be disabled by setting explicitly: `{ endOfLine: EndOfLine.NONE }`.
	// Default: EndOfLine.LF
	//
	// Experimental.
	EndOfLine EndOfLine `field:"optional" json:"endOfLine" yaml:"endOfLine"`
}
