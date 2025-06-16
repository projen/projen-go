//go:build !no_runtime_type_checking

package biomeconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_ISecurity) validateSetNoDangerouslySetInnerHtmlParameters(val interface{}) error {
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

func (j *jsiiProxy_ISecurity) validateSetNoDangerouslySetInnerHtmlWithChildrenParameters(val interface{}) error {
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

func (j *jsiiProxy_ISecurity) validateSetNoGlobalEvalParameters(val interface{}) error {
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

