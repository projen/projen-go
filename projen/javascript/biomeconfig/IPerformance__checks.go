//go:build !no_runtime_type_checking

package biomeconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_IPerformance) validateSetNoAccumulatingSpreadParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case IRuleWithNoOptions:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *string, IRuleWithNoOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IPerformance) validateSetNoBarrelFileParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case IRuleWithNoOptions:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *string, IRuleWithNoOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IPerformance) validateSetNoDeleteParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case IRuleWithFixNoOptions:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *string, IRuleWithFixNoOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IPerformance) validateSetNoReExportAllParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case IRuleWithNoOptions:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *string, IRuleWithNoOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_IPerformance) validateSetUseTopLevelRegexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case IRuleWithNoOptions:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *string, IRuleWithNoOptions; received %#v (a %T)", val, val)
		}
	}

	return nil
}

