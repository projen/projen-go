package javascript


// Determines how modules get resolved.
// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html
//
// Experimental.
type TypeScriptModuleResolution string

const (
	// TypeScript's former default resolution strategy.
	// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html#classic
	//
	// Experimental.
	TypeScriptModuleResolution_CLASSIC TypeScriptModuleResolution = "CLASSIC"
	// Resolution strategy which attempts to mimic the Node.js module resolution strategy at runtime.
	// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html#node
	//
	// Experimental.
	TypeScriptModuleResolution_NODE TypeScriptModuleResolution = "NODE"
	// Node.js’ ECMAScript Module Support from TypeScript 4.7 onwards.
	// See: https://www.typescriptlang.org/tsconfig#moduleResolution
	//
	// Experimental.
	TypeScriptModuleResolution_NODE16 TypeScriptModuleResolution = "NODE16"
	// Node.js’ ECMAScript Module Support from TypeScript 4.7 onwards.
	// See: https://www.typescriptlang.org/tsconfig#moduleResolution
	//
	// Experimental.
	TypeScriptModuleResolution_NODE_NEXT TypeScriptModuleResolution = "NODE_NEXT"
)

