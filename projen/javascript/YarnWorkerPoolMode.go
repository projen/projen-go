package javascript


// https://yarnpkg.com/configuration/yarnrc#taskPoolMode.
// Deprecated: use {@link YarnTaskPoolMode } instead. Yarn calls this setting
// `taskPoolMode`; there is no `workerPoolMode` setting.
type YarnWorkerPoolMode string

const (
	// Deprecated: use {@link YarnTaskPoolMode } instead. Yarn calls this setting
	// `taskPoolMode`; there is no `workerPoolMode` setting.
	YarnWorkerPoolMode_ASYNC YarnWorkerPoolMode = "ASYNC"
	// Deprecated: use {@link YarnTaskPoolMode } instead. Yarn calls this setting
	// `taskPoolMode`; there is no `workerPoolMode` setting.
	YarnWorkerPoolMode_WORKERS YarnWorkerPoolMode = "WORKERS"
)

