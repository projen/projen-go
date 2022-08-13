package javascript


// Determines how JSX should get transformed into valid JavaScript.
// See: https://www.typescriptlang.org/docs/handbook/jsx.html
//
// Experimental.
type TypeScriptJsxMode string

const (
	// Keeps the JSX as part of the output to be further consumed by another transform step (e.g. Babel).
	// Experimental.
	TypeScriptJsxMode_PRESERVE TypeScriptJsxMode = "PRESERVE"
	// Converts JSX syntax into React.createElement, does not need to go through a JSX transformation before use, and the output will have a .js file extension.
	// Experimental.
	TypeScriptJsxMode_REACT TypeScriptJsxMode = "REACT"
	// Keeps all JSX like 'preserve' mode, but output will have a .js extension.
	// Experimental.
	TypeScriptJsxMode_REACT_NATIVE TypeScriptJsxMode = "REACT_NATIVE"
	// Passes `key` separately from props and always passes `children` as props (since React 17).
	// See: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-1.html#react-17-jsx-factories
	//
	// Experimental.
	TypeScriptJsxMode_REACT_JSX TypeScriptJsxMode = "REACT_JSX"
	// Same as `REACT_JSX` with additional debug data.
	// Experimental.
	TypeScriptJsxMode_REACT_JSXDEV TypeScriptJsxMode = "REACT_JSXDEV"
)

