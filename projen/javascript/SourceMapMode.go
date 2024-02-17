package javascript


// SourceMap mode for esbuild.
// See: https://esbuild.github.io/api/#sourcemap
//
// Experimental.
type SourceMapMode string

const (
	// Default sourceMap mode - will generate a .js.map file alongside any generated .js file and add a special //# sourceMappingURL= comment to the bottom of the .js file pointing to the .js.map file.
	// Experimental.
	SourceMapMode_DEFAULT SourceMapMode = "DEFAULT"
	// External sourceMap mode - If you want to omit the special //# sourceMappingURL= comment from the generated .js file but you still want to generate the .js.map files.
	// Experimental.
	SourceMapMode_EXTERNAL SourceMapMode = "EXTERNAL"
	// Inline sourceMap mode - If you want to insert the entire source map into the .js file instead of generating a separate .js.map file.
	// Experimental.
	SourceMapMode_INLINE SourceMapMode = "INLINE"
	// Both sourceMap mode - If you want to have the effect of both inline and external simultaneously.
	// Experimental.
	SourceMapMode_BOTH SourceMapMode = "BOTH"
)

