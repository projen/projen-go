package biomeconfig


// Indicates the type of runtime or transformation used for interpreting JSX.
// Experimental.
type JsxRuntime string

const (
	// Indicates a modern or native JSX environment, that doesn't require special handling by Biome.
	//
	// (transparent).
	// Experimental.
	JsxRuntime_TRANSPARENT JsxRuntime = "TRANSPARENT"
	// Indicates a classic React environment that requires the `React` import.
	//
	// Corresponds to the `react` value for the `jsx` option in TypeScript's
	// `tsconfig.json`.
	//
	// This option should only be necessary if you cannot upgrade to a React
	// version that supports the new JSX runtime. For more information about
	// the old vs. new JSX runtime, please see:
	// <https://legacy.reactjs.org/blog/2020/09/22/introducing-the-new-jsx-transform.html> (reactClassic)
	// Experimental.
	JsxRuntime_REACT_CLASSIC JsxRuntime = "REACT_CLASSIC"
)

