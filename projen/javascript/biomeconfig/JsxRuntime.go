package biomeconfig


// Indicates the type of runtime or transformation used for interpreting JSX.
// Experimental.
type JsxRuntime string

const (
	// transparent.
	// Experimental.
	JsxRuntime_TRANSPARENT JsxRuntime = "TRANSPARENT"
	// reactClassic.
	// Experimental.
	JsxRuntime_REACT_CLASSIC JsxRuntime = "REACT_CLASSIC"
)

