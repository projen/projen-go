package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// The runtime for the AWS Lambda function.
// Experimental.
type LambdaRuntime interface {
	// Experimental.
	EsbuildPlatform() *string
	// The esbuild setting to use.
	// Experimental.
	EsbuildTarget() *string
	// The Node.js runtime to use.
	// Experimental.
	FunctionRuntime() *string
}

// The jsii proxy struct for LambdaRuntime
type jsiiProxy_LambdaRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaRuntime) EsbuildPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRuntime) EsbuildTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRuntime) FunctionRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionRuntime",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaRuntime(functionRuntime *string, esbuildTarget *string) LambdaRuntime {
	_init_.Initialize()

	if err := validateNewLambdaRuntimeParameters(functionRuntime, esbuildTarget); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaRuntime{}

	_jsii_.Create(
		"projen.awscdk.LambdaRuntime",
		[]interface{}{functionRuntime, esbuildTarget},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaRuntime_Override(l LambdaRuntime, functionRuntime *string, esbuildTarget *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaRuntime",
		[]interface{}{functionRuntime, esbuildTarget},
		l,
	)
}

func LambdaRuntime_NODEJS_10_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_10_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_12_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_12_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_14_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_14_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_16_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_16_X",
		&returns,
	)
	return returns
}

