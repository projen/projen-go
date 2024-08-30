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
	// `--moduleResolution node` was renamed to `node10` (keeping `node` as an alias for backward compatibility) in TypeScript 5.0. It reflects the CommonJS module resolution algorithm as it existed in Node.js versions earlier than v12. It should no longer be used.
	// See: https://www.typescriptlang.org/docs/handbook/modules/reference.html#node10-formerly-known-as-node
	//
	// Experimental.
	TypeScriptModuleResolution_NODE10 TypeScriptModuleResolution = "NODE10"
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
	// Resolution strategy which attempts to mimic resolution patterns of modern bundlers;
	//
	// from TypeScript 5.0 onwards.
	// See: https://www.typescriptlang.org/tsconfig#moduleResolution
	//
	// Experimental.
	TypeScriptModuleResolution_BUNDLER TypeScriptModuleResolution = "BUNDLER"
)

