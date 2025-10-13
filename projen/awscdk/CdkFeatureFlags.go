package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// CDK feature flags configuration.
// Experimental.
type CdkFeatureFlags interface {
	ICdkFeatureFlags
	// Experimental.
	Flags() *map[string]interface{}
}

// The jsii proxy struct for CdkFeatureFlags
type jsiiProxy_CdkFeatureFlags struct {
	jsiiProxy_ICdkFeatureFlags
}

func (j *jsiiProxy_CdkFeatureFlags) Flags() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}


func CdkFeatureFlags_V1() CdkFeatureFlagsV1 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV1
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlags",
		"V1",
		&returns,
	)
	return returns
}

func CdkFeatureFlags_V2() CdkFeatureFlagsV2 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV2
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlags",
		"V2",
		&returns,
	)
	return returns
}

