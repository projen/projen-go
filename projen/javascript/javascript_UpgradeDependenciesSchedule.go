package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// How often to check for new versions and raise pull requests for version upgrades.
// Experimental.
type UpgradeDependenciesSchedule interface {
	// Experimental.
	Cron() *[]*string
}

// The jsii proxy struct for UpgradeDependenciesSchedule
type jsiiProxy_UpgradeDependenciesSchedule struct {
	_ byte // padding
}

func (j *jsiiProxy_UpgradeDependenciesSchedule) Cron() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cron",
		&returns,
	)
	return returns
}


// Create a schedule from a raw cron expression.
// Experimental.
func UpgradeDependenciesSchedule_Expressions(cron *[]*string) UpgradeDependenciesSchedule {
	_init_.Initialize()

	var returns UpgradeDependenciesSchedule

	_jsii_.StaticInvoke(
		"projen.javascript.UpgradeDependenciesSchedule",
		"expressions",
		[]interface{}{cron},
		&returns,
	)

	return returns
}

func UpgradeDependenciesSchedule_DAILY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"DAILY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_MONTHLY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"MONTHLY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_NEVER() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"NEVER",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_WEEKDAY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"WEEKDAY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_WEEKLY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"WEEKLY",
		&returns,
	)
	return returns
}

