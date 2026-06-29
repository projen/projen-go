package javascript


// The default configuration of fake timers for all tests.
// See: https://jestjs.io/docs/configuration#faketimers-object
//
// Experimental.
type FakeTimers struct {
	// If set to `true` all timers will be advanced automatically by 20 milliseconds every 20 milliseconds.
	//
	// A custom time delta may be provided by passing a number.
	// Default: - false.
	//
	// Experimental.
	AdvanceTimers interface{} `field:"optional" json:"advanceTimers" yaml:"advanceTimers"`
	// List of names of APIs (e.g. `Date`, `nextTick`, `setTimeout`) that should not be faked.
	// Default: - [] (all APIs are faked).
	//
	// Experimental.
	DoNotFake *[]*string `field:"optional" json:"doNotFake" yaml:"doNotFake"`
	// Whether fake timers should be enabled for all test files.
	// Default: - false.
	//
	// Experimental.
	EnableGlobally *bool `field:"optional" json:"enableGlobally" yaml:"enableGlobally"`
	// Use the old fake timers implementation instead of one backed by `@sinonjs/fake-timers`.
	// Default: - false.
	//
	// Experimental.
	LegacyFakeTimers *bool `field:"optional" json:"legacyFakeTimers" yaml:"legacyFakeTimers"`
	// Sets current system time to be used by fake timers, in milliseconds.
	// Default: - Date.now()
	//
	// Experimental.
	Now *float64 `field:"optional" json:"now" yaml:"now"`
	// Maximum number of recursive timers that will be run.
	// Default: - 100000.
	//
	// Experimental.
	TimerLimit *float64 `field:"optional" json:"timerLimit" yaml:"timerLimit"`
}

