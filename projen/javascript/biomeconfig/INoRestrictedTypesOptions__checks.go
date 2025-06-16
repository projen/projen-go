//go:build !no_runtime_type_checking

package biomeconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_INoRestrictedTypesOptions) validateSetTypesParameters(val *map[string]interface{}) error {
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case *string:
			// ok
		case string:
			// ok
		case ICustomRestrictedTypeOptions:
			// ok
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *string, ICustomRestrictedTypeOptions; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

