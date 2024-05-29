package javascript


// This setting controls how TypeScript determines whether a file is a script or a module.
// See: https://www.typescriptlang.org/docs/handbook/modules/theory.html#scripts-and-modules-in-javascript
//
// Experimental.
type TypeScriptModuleDetection string

const (
	// TypeScript will not only look for import and export statements, but it will also check whether the "type" field in a package.json is set to "module" when running with module: nodenext or node16, and check whether the current file is a JSX file when running under jsx: react-jsx.
	// See: https://www.typescriptlang.org/tsconfig/#moduleDetection
	//
	// Experimental.
	TypeScriptModuleDetection_AUTO TypeScriptModuleDetection = "AUTO"
	// The same behavior as 4.6 and prior, usings import and export statements to determine whether a file is a module.
	// See: https://www.typescriptlang.org/tsconfig/#moduleDetection
	//
	// Experimental.
	TypeScriptModuleDetection_LEGACY TypeScriptModuleDetection = "LEGACY"
	// Ensures that every non-declaration file is treated as a module.
	// See: https://www.typescriptlang.org/tsconfig/#moduleDetection
	//
	// Experimental.
	TypeScriptModuleDetection_FORCE TypeScriptModuleDetection = "FORCE"
)

