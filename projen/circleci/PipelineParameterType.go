package circleci


// Pipeline parameter types.
// See: https://circleci.com/docs/2.0/reusing-config#parameter-syntax
//
// Experimental.
type PipelineParameterType string

const (
	// Experimental.
	PipelineParameterType_STRING PipelineParameterType = "STRING"
	// Experimental.
	PipelineParameterType_BOOLEAN PipelineParameterType = "BOOLEAN"
	// Experimental.
	PipelineParameterType_INTEGER PipelineParameterType = "INTEGER"
	// Experimental.
	PipelineParameterType_ENUM PipelineParameterType = "ENUM"
)

