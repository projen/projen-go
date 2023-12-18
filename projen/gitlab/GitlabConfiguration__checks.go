//go:build !no_runtime_type_checking

package gitlab

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (g *jsiiProxy_GitlabConfiguration) validateAddDefaultCachesParameters(caches *[]*Cache) error {
	if caches == nil {
		return fmt.Errorf("parameter caches is required, but nil was provided")
	}
	for idx_efbe9d, v := range *caches {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter caches[%#v]", idx_efbe9d) }); err != nil {
			return err
		}
	}

	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddGlobalVariablesParameters(variables *map[string]interface{}) error {
	if variables == nil {
		return fmt.Errorf("parameter variables is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddIncludesParameters(includes *[]*Include) error {
	for idx_9472f2, v := range *includes {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter includes[%#v]", idx_9472f2) }); err != nil {
			return err
		}
	}

	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddJobsParameters(jobs *map[string]*Job) error {
	if jobs == nil {
		return fmt.Errorf("parameter jobs is required, but nil was provided")
	}
	for idx_5d9a17, v := range *jobs {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter jobs[%#v]", idx_5d9a17) }); err != nil {
			return err
		}
	}

	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddServicesParameters(services *[]*Service) error {
	for idx_ef1c4b, v := range *services {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter services[%#v]", idx_ef1c4b) }); err != nil {
			return err
		}
	}

	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateCreateNestedTemplatesParameters(config *map[string]*CiConfigurationOptions) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	for idx_b79606, v := range *config {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter config[%#v]", idx_b79606) }); err != nil {
			return err
		}
	}

	return nil
}

func validateGitlabConfiguration_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGitlabConfiguration_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewGitlabConfigurationParameters(project projen.Project, options *CiConfigurationOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

