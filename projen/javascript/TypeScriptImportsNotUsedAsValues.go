package javascript


// This flag controls how `import` works, there are 3 different options.
// See: https://www.typescriptlang.org/tsconfig#importsNotUsedAsValues
//
// Experimental.
type TypeScriptImportsNotUsedAsValues string

const (
	// The default behavior of dropping `import` statements which only reference types.
	// Experimental.
	TypeScriptImportsNotUsedAsValues_REMOVE TypeScriptImportsNotUsedAsValues = "REMOVE"
	// Preserves all `import` statements whose values or types are never used.
	//
	// This can cause imports/side-effects to be preserved.
	// Experimental.
	TypeScriptImportsNotUsedAsValues_PRESERVE TypeScriptImportsNotUsedAsValues = "PRESERVE"
	// This preserves all imports (the same as the preserve option), but will error when a value import is only used as a type.
	//
	// This might be useful if you want to ensure no values are being accidentally imported, but still make side-effect imports explicit.
	// Experimental.
	TypeScriptImportsNotUsedAsValues_ERROR TypeScriptImportsNotUsedAsValues = "ERROR"
)

