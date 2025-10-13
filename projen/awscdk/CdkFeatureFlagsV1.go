package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// CDK V1 feature flags configuration.
// Deprecated: CDK V1 is EOS. Upgrade to CDK V2.
type CdkFeatureFlagsV1 interface {
	ICdkFeatureFlags
	// Deprecated: CDK V1 is EOS. Upgrade to CDK V2.
	Flags() *map[string]interface{}
}

// The jsii proxy struct for CdkFeatureFlagsV1
type jsiiProxy_CdkFeatureFlagsV1 struct {
	jsiiProxy_ICdkFeatureFlags
}

func (j *jsiiProxy_CdkFeatureFlagsV1) Flags() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}


func CdkFeatureFlagsV1_ALL() CdkFeatureFlagsV1 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV1
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlagsV1",
		"ALL",
		&returns,
	)
	return returns
}

func CdkFeatureFlagsV1_NONE() CdkFeatureFlagsV1 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV1
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlagsV1",
		"NONE",
		&returns,
	)
	return returns
}

