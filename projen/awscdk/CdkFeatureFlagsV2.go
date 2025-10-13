package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// CDK V2 feature flags configuration.
// Experimental.
type CdkFeatureFlagsV2 interface {
	ICdkFeatureFlags
	// Experimental.
	Flags() *map[string]interface{}
}

// The jsii proxy struct for CdkFeatureFlagsV2
type jsiiProxy_CdkFeatureFlagsV2 struct {
	jsiiProxy_ICdkFeatureFlags
}

func (j *jsiiProxy_CdkFeatureFlagsV2) Flags() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}


// Attempt to load the feature flags from the `aws-cdk-lib/recommended-feature-flags.json` in a locally available npm package. This file is typically only present in AWS CDK TypeScript projects, but can yield more accurate results.
//
// Falls back to all known feature flags if not found.
// Experimental.
func CdkFeatureFlagsV2_FromLocalAwsCdkLib() CdkFeatureFlagsV2 {
	_init_.Initialize()

	var returns CdkFeatureFlagsV2

	_jsii_.StaticInvoke(
		"projen.awscdk.CdkFeatureFlagsV2",
		"fromLocalAwsCdkLib",
		nil, // no parameters
		&returns,
	)

	return returns
}

func CdkFeatureFlagsV2_ALL() CdkFeatureFlagsV2 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV2
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlagsV2",
		"ALL",
		&returns,
	)
	return returns
}

func CdkFeatureFlagsV2_NONE() CdkFeatureFlagsV2 {
	_init_.Initialize()
	var returns CdkFeatureFlagsV2
	_jsii_.StaticGet(
		"projen.awscdk.CdkFeatureFlagsV2",
		"NONE",
		&returns,
	)
	return returns
}

