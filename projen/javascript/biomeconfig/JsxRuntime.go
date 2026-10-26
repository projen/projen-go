package biomeconfig


// How Biome's analyzer accounts for imports used by JSX.
// Experimental.
type JsxRuntime string

const (
	// An automatic or native JSX environment that doesn't require an in-scope factory import.
	//
	// (transparent).
	// Experimental.
	JsxRuntime_TRANSPARENT JsxRuntime = "TRANSPARENT"
	// A classic React environment that requires the React import, or custom JSX factory imports configured in `tsconfig.json`.
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

