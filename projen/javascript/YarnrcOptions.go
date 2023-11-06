package javascript


// Configuration for .yarnrc.yml in Yarn Berry v4.
// Experimental.
type YarnrcOptions struct {
	// https://yarnpkg.com/configuration/yarnrc#cacheFolder.
	// Experimental.
	CacheFolder *string `field:"optional" json:"cacheFolder" yaml:"cacheFolder"`
	// https://yarnpkg.com/configuration/yarnrc#cacheMigrationMode.
	// Experimental.
	CacheMigrationMode YarnCacheMigrationMode `field:"optional" json:"cacheMigrationMode" yaml:"cacheMigrationMode"`
	// https://yarnpkg.com/configuration/yarnrc#changesetBaseRefs.
	// Experimental.
	ChangesetBaseRefs *[]*string `field:"optional" json:"changesetBaseRefs" yaml:"changesetBaseRefs"`
	// https://yarnpkg.com/configuration/yarnrc#changesetIgnorePatterns.
	// Experimental.
	ChangesetIgnorePatterns *[]*string `field:"optional" json:"changesetIgnorePatterns" yaml:"changesetIgnorePatterns"`
	// https://yarnpkg.com/configuration/yarnrc#checksumBehavior.
	// Experimental.
	ChecksumBehavior YarnChecksumBehavior `field:"optional" json:"checksumBehavior" yaml:"checksumBehavior"`
	// https://yarnpkg.com/configuration/yarnrc#cloneConcurrency.
	// Experimental.
	CloneConcurrency *float64 `field:"optional" json:"cloneConcurrency" yaml:"cloneConcurrency"`
	// https://yarnpkg.com/configuration/yarnrc#compressionLevel.
	// Experimental.
	CompressionLevel interface{} `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// https://yarnpkg.com/configuration/yarnrc#constraintsPath.
	// Experimental.
	ConstraintsPath *string `field:"optional" json:"constraintsPath" yaml:"constraintsPath"`
	// https://yarnpkg.com/configuration/yarnrc#defaultLanguageName.
	// Experimental.
	DefaultLanguageName *string `field:"optional" json:"defaultLanguageName" yaml:"defaultLanguageName"`
	// https://yarnpkg.com/configuration/yarnrc#defaultProtocol.
	// Experimental.
	DefaultProtocol *string `field:"optional" json:"defaultProtocol" yaml:"defaultProtocol"`
	// https://yarnpkg.com/configuration/yarnrc#defaultSemverRangePrefix.
	// Experimental.
	DefaultSemverRangePrefix YarnDefaultSemverRangePrefix `field:"optional" json:"defaultSemverRangePrefix" yaml:"defaultSemverRangePrefix"`
	// https://yarnpkg.com/configuration/yarnrc#deferredVersionFolder.
	// Experimental.
	DeferredVersionFolder *string `field:"optional" json:"deferredVersionFolder" yaml:"deferredVersionFolder"`
	// https://yarnpkg.com/configuration/yarnrc#enableColors.
	// Experimental.
	EnableColors *bool `field:"optional" json:"enableColors" yaml:"enableColors"`
	// https://yarnpkg.com/configuration/yarnrc#enableConstraintsCheck.
	// Experimental.
	EnableConstraintsCheck *bool `field:"optional" json:"enableConstraintsCheck" yaml:"enableConstraintsCheck"`
	// https://yarnpkg.com/configuration/yarnrc#enableGlobalCache.
	// Experimental.
	EnableGlobalCache *bool `field:"optional" json:"enableGlobalCache" yaml:"enableGlobalCache"`
	// https://yarnpkg.com/configuration/yarnrc#enableHardenedMode.
	// Experimental.
	EnableHardenedMode *bool `field:"optional" json:"enableHardenedMode" yaml:"enableHardenedMode"`
	// https://yarnpkg.com/configuration/yarnrc#enableHyperlinks.
	// Experimental.
	EnableHyperlinks *bool `field:"optional" json:"enableHyperlinks" yaml:"enableHyperlinks"`
	// https://yarnpkg.com/configuration/yarnrc#enableImmutableCache.
	// Experimental.
	EnableImmutableCache *bool `field:"optional" json:"enableImmutableCache" yaml:"enableImmutableCache"`
	// https://yarnpkg.com/configuration/yarnrc#enableImmutableInstalls.
	// Experimental.
	EnableImmutableInstalls *bool `field:"optional" json:"enableImmutableInstalls" yaml:"enableImmutableInstalls"`
	// https://yarnpkg.com/configuration/yarnrc#enableInlineBuilds.
	// Experimental.
	EnableInlineBuilds *bool `field:"optional" json:"enableInlineBuilds" yaml:"enableInlineBuilds"`
	// https://yarnpkg.com/configuration/yarnrc#enableInlineHunks.
	// Experimental.
	EnableInlineHunks *bool `field:"optional" json:"enableInlineHunks" yaml:"enableInlineHunks"`
	// https://yarnpkg.com/configuration/yarnrc#enableMessageNames.
	// Experimental.
	EnableMessageNames *bool `field:"optional" json:"enableMessageNames" yaml:"enableMessageNames"`
	// https://yarnpkg.com/configuration/yarnrc#enableMirror.
	// Experimental.
	EnableMirror *bool `field:"optional" json:"enableMirror" yaml:"enableMirror"`
	// https://yarnpkg.com/configuration/yarnrc#enableNetwork.
	// Experimental.
	EnableNetwork *bool `field:"optional" json:"enableNetwork" yaml:"enableNetwork"`
	// https://yarnpkg.com/configuration/yarnrc#enableOfflineMode.
	// Experimental.
	EnableOfflineMode *bool `field:"optional" json:"enableOfflineMode" yaml:"enableOfflineMode"`
	// https://yarnpkg.com/configuration/yarnrc#enableProgressBars.
	// Experimental.
	EnableProgressBars *bool `field:"optional" json:"enableProgressBars" yaml:"enableProgressBars"`
	// https://yarnpkg.com/configuration/yarnrc#enableScripts.
	// Experimental.
	EnableScripts *bool `field:"optional" json:"enableScripts" yaml:"enableScripts"`
	// https://yarnpkg.com/configuration/yarnrc#enableStrictSsl.
	// Experimental.
	EnableStrictSsl *bool `field:"optional" json:"enableStrictSsl" yaml:"enableStrictSsl"`
	// https://yarnpkg.com/configuration/yarnrc#enableTelemetry.
	// Experimental.
	EnableTelemetry *bool `field:"optional" json:"enableTelemetry" yaml:"enableTelemetry"`
	// https://yarnpkg.com/configuration/yarnrc#enableTimers.
	// Experimental.
	EnableTimers *bool `field:"optional" json:"enableTimers" yaml:"enableTimers"`
	// https://yarnpkg.com/configuration/yarnrc#enableTransparentWorkspaces.
	// Experimental.
	EnableTransparentWorkspaces *bool `field:"optional" json:"enableTransparentWorkspaces" yaml:"enableTransparentWorkspaces"`
	// https://yarnpkg.com/configuration/yarnrc#globalFolder.
	// Experimental.
	GlobalFolder *string `field:"optional" json:"globalFolder" yaml:"globalFolder"`
	// https://yarnpkg.com/configuration/yarnrc#httpProxy.
	// Experimental.
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// https://yarnpkg.com/configuration/yarnrc#httpRetry.
	// Experimental.
	HttpRetry *float64 `field:"optional" json:"httpRetry" yaml:"httpRetry"`
	// https://yarnpkg.com/configuration/yarnrc#httpsCaFilePath.
	// Experimental.
	HttpsCaFilePath *string `field:"optional" json:"httpsCaFilePath" yaml:"httpsCaFilePath"`
	// https://yarnpkg.com/configuration/yarnrc#httpsCertFilePath.
	// Experimental.
	HttpsCertFilePath *string `field:"optional" json:"httpsCertFilePath" yaml:"httpsCertFilePath"`
	// https://yarnpkg.com/configuration/yarnrc#httpsKeyFilePath.
	// Experimental.
	HttpsKeyFilePath *string `field:"optional" json:"httpsKeyFilePath" yaml:"httpsKeyFilePath"`
	// https://yarnpkg.com/configuration/yarnrc#httpsProxy.
	// Experimental.
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// https://yarnpkg.com/configuration/yarnrc#httpTimeout.
	// Experimental.
	HttpTimeout *float64 `field:"optional" json:"httpTimeout" yaml:"httpTimeout"`
	// https://v3.yarnpkg.com/configuration/yarnrc#ignoreCwd.
	// Deprecated: - removed in Yarn v4 and newer.
	IgnoreCwd *bool `field:"optional" json:"ignoreCwd" yaml:"ignoreCwd"`
	// https://yarnpkg.com/configuration/yarnrc#ignorePath.
	// Experimental.
	IgnorePath *bool `field:"optional" json:"ignorePath" yaml:"ignorePath"`
	// https://yarnpkg.com/configuration/yarnrc#immutablePatterns.
	// Experimental.
	ImmutablePatterns *[]*string `field:"optional" json:"immutablePatterns" yaml:"immutablePatterns"`
	// https://yarnpkg.com/configuration/yarnrc#initFields.
	// Experimental.
	InitFields *map[string]interface{} `field:"optional" json:"initFields" yaml:"initFields"`
	// https://yarnpkg.com/configuration/yarnrc#initScope.
	// Experimental.
	InitScope *string `field:"optional" json:"initScope" yaml:"initScope"`
	// https://yarnpkg.com/configuration/yarnrc#injectEnvironmentFiles.
	// Experimental.
	InjectEnvironmentFiles *[]*string `field:"optional" json:"injectEnvironmentFiles" yaml:"injectEnvironmentFiles"`
	// https://yarnpkg.com/configuration/yarnrc#installStatePath.
	// Experimental.
	InstallStatePath *string `field:"optional" json:"installStatePath" yaml:"installStatePath"`
	// https://v3.yarnpkg.com/configuration/yarnrc#lockfileFilename.
	// Deprecated: - removed in Yarn v4 and newer.
	LockfileFilename *string `field:"optional" json:"lockfileFilename" yaml:"lockfileFilename"`
	// https://yarnpkg.com/configuration/yarnrc#logFilters.
	// Experimental.
	LogFilters *[]*YarnLogFilter `field:"optional" json:"logFilters" yaml:"logFilters"`
	// https://yarnpkg.com/configuration/yarnrc#networkConcurrency.
	// Experimental.
	NetworkConcurrency *float64 `field:"optional" json:"networkConcurrency" yaml:"networkConcurrency"`
	// https://yarnpkg.com/configuration/yarnrc#networkSettings.
	// Experimental.
	NetworkSettings *map[string]*YarnNetworkSetting `field:"optional" json:"networkSettings" yaml:"networkSettings"`
	// https://yarnpkg.com/configuration/yarnrc#nmHoistingLimits.
	// Experimental.
	NmHoistingLimits YarnNmHoistingLimit `field:"optional" json:"nmHoistingLimits" yaml:"nmHoistingLimits"`
	// https://yarnpkg.com/configuration/yarnrc#nmMode.
	// Experimental.
	NmMode YarnNmMode `field:"optional" json:"nmMode" yaml:"nmMode"`
	// https://yarnpkg.com/configuration/yarnrc#nmSelfReferences.
	// Experimental.
	NmSelfReferences *bool `field:"optional" json:"nmSelfReferences" yaml:"nmSelfReferences"`
	// https://yarnpkg.com/configuration/yarnrc#nodeLinker.
	// Experimental.
	NodeLinker YarnNodeLinker `field:"optional" json:"nodeLinker" yaml:"nodeLinker"`
	// https://yarnpkg.com/configuration/yarnrc#npmAlwaysAuth.
	// Experimental.
	NpmAlwaysAuth *bool `field:"optional" json:"npmAlwaysAuth" yaml:"npmAlwaysAuth"`
	// https://yarnpkg.com/configuration/yarnrc#npmAuditExcludePackages.
	// Experimental.
	NpmAuditExcludePackages *[]*string `field:"optional" json:"npmAuditExcludePackages" yaml:"npmAuditExcludePackages"`
	// https://yarnpkg.com/configuration/yarnrc#npmAuditIgnoreAdvisories.
	// Experimental.
	NpmAuditIgnoreAdvisories *[]*string `field:"optional" json:"npmAuditIgnoreAdvisories" yaml:"npmAuditIgnoreAdvisories"`
	// https://yarnpkg.com/configuration/yarnrc#npmAuditRegistry.
	// Experimental.
	NpmAuditRegistry *string `field:"optional" json:"npmAuditRegistry" yaml:"npmAuditRegistry"`
	// https://yarnpkg.com/configuration/yarnrc#npmAuthIdent.
	// Experimental.
	NpmAuthIdent *string `field:"optional" json:"npmAuthIdent" yaml:"npmAuthIdent"`
	// https://yarnpkg.com/configuration/yarnrc#npmAuthToken.
	// Experimental.
	NpmAuthToken *string `field:"optional" json:"npmAuthToken" yaml:"npmAuthToken"`
	// https://yarnpkg.com/configuration/yarnrc#npmPublishAccess.
	// Experimental.
	NpmPublishAccess YarnNpmPublishAccess `field:"optional" json:"npmPublishAccess" yaml:"npmPublishAccess"`
	// https://yarnpkg.com/configuration/yarnrc#npmPublishRegistry.
	// Experimental.
	NpmPublishRegistry *string `field:"optional" json:"npmPublishRegistry" yaml:"npmPublishRegistry"`
	// https://yarnpkg.com/configuration/yarnrc#npmRegistries.
	// Experimental.
	NpmRegistries *map[string]*YarnNpmRegistry `field:"optional" json:"npmRegistries" yaml:"npmRegistries"`
	// https://yarnpkg.com/configuration/yarnrc#npmRegistryServer.
	// Experimental.
	NpmRegistryServer *string `field:"optional" json:"npmRegistryServer" yaml:"npmRegistryServer"`
	// https://yarnpkg.com/configuration/yarnrc#npmScopes.
	// Experimental.
	NpmScopes *map[string]*YarnNpmScope `field:"optional" json:"npmScopes" yaml:"npmScopes"`
	// https://yarnpkg.com/configuration/yarnrc#packageExtensions.
	// Experimental.
	PackageExtensions *map[string]*YarnPackageExtension `field:"optional" json:"packageExtensions" yaml:"packageExtensions"`
	// https://yarnpkg.com/configuration/yarnrc#patchFolder.
	// Experimental.
	PatchFolder *string `field:"optional" json:"patchFolder" yaml:"patchFolder"`
	// https://v3.yarnpkg.com/configuration/yarnrc#pnpDataPath.
	// Deprecated: - removed in Yarn v4 and newer.
	PnpDataPath *string `field:"optional" json:"pnpDataPath" yaml:"pnpDataPath"`
	// https://yarnpkg.com/configuration/yarnrc#pnpEnableEsmLoader.
	// Experimental.
	PnpEnableEsmLoader *bool `field:"optional" json:"pnpEnableEsmLoader" yaml:"pnpEnableEsmLoader"`
	// https://yarnpkg.com/configuration/yarnrc#pnpEnableInlining.
	// Experimental.
	PnpEnableInlining *bool `field:"optional" json:"pnpEnableInlining" yaml:"pnpEnableInlining"`
	// https://yarnpkg.com/configuration/yarnrc#pnpFallbackMode.
	// Experimental.
	PnpFallbackMode YarnPnpFallbackMode `field:"optional" json:"pnpFallbackMode" yaml:"pnpFallbackMode"`
	// https://yarnpkg.com/configuration/yarnrc#pnpIgnorePatterns.
	// Experimental.
	PnpIgnorePatterns *[]*string `field:"optional" json:"pnpIgnorePatterns" yaml:"pnpIgnorePatterns"`
	// https://yarnpkg.com/configuration/yarnrc#pnpMode.
	// Experimental.
	PnpMode YarnPnpMode `field:"optional" json:"pnpMode" yaml:"pnpMode"`
	// https://yarnpkg.com/configuration/yarnrc#pnpShebang.
	// Experimental.
	PnpShebang *string `field:"optional" json:"pnpShebang" yaml:"pnpShebang"`
	// https://yarnpkg.com/configuration/yarnrc#pnpUnpluggedFolder.
	// Experimental.
	PnpUnpluggedFolder *string `field:"optional" json:"pnpUnpluggedFolder" yaml:"pnpUnpluggedFolder"`
	// https://v3.yarnpkg.com/configuration/yarnrc#preferAggregateCacheInfo.
	// Deprecated: - removed in Yarn v4 and newer.
	PreferAggregateCacheInfo *bool `field:"optional" json:"preferAggregateCacheInfo" yaml:"preferAggregateCacheInfo"`
	// https://yarnpkg.com/configuration/yarnrc#preferDeferredVersions.
	// Experimental.
	PreferDeferredVersions *bool `field:"optional" json:"preferDeferredVersions" yaml:"preferDeferredVersions"`
	// https://yarnpkg.com/configuration/yarnrc#preferInteractive.
	// Experimental.
	PreferInteractive *bool `field:"optional" json:"preferInteractive" yaml:"preferInteractive"`
	// https://yarnpkg.com/configuration/yarnrc#preferReuse.
	// Experimental.
	PreferReuse *bool `field:"optional" json:"preferReuse" yaml:"preferReuse"`
	// https://yarnpkg.com/configuration/yarnrc#preferTruncatedLines.
	// Experimental.
	PreferTruncatedLines *bool `field:"optional" json:"preferTruncatedLines" yaml:"preferTruncatedLines"`
	// https://yarnpkg.com/configuration/yarnrc#progressBarStyle.
	// Experimental.
	ProgressBarStyle YarnProgressBarStyle `field:"optional" json:"progressBarStyle" yaml:"progressBarStyle"`
	// https://yarnpkg.com/configuration/yarnrc#rcFilename.
	// Experimental.
	RcFilename *string `field:"optional" json:"rcFilename" yaml:"rcFilename"`
	// https://yarnpkg.com/configuration/yarnrc#supportedArchitectures.
	// Experimental.
	SupportedArchitectures *YarnSupportedArchitectures `field:"optional" json:"supportedArchitectures" yaml:"supportedArchitectures"`
	// https://yarnpkg.com/configuration/yarnrc#taskPoolConcurrency.
	// Experimental.
	TaskPoolConcurrency *string `field:"optional" json:"taskPoolConcurrency" yaml:"taskPoolConcurrency"`
	// https://yarnpkg.com/configuration/yarnrc#telemetryInterval.
	// Experimental.
	TelemetryInterval *float64 `field:"optional" json:"telemetryInterval" yaml:"telemetryInterval"`
	// https://yarnpkg.com/configuration/yarnrc#telemetryUserId.
	// Experimental.
	TelemetryUserId *string `field:"optional" json:"telemetryUserId" yaml:"telemetryUserId"`
	// https://yarnpkg.com/configuration/yarnrc#tsEnableAutoTypes.
	// Experimental.
	TsEnableAutoTypes *bool `field:"optional" json:"tsEnableAutoTypes" yaml:"tsEnableAutoTypes"`
	// https://yarnpkg.com/configuration/yarnrc#unsafeHttpWhitelist.
	// Experimental.
	UnsafeHttpWhitelist *[]*string `field:"optional" json:"unsafeHttpWhitelist" yaml:"unsafeHttpWhitelist"`
	// https://yarnpkg.com/configuration/yarnrc#virtualFolder.
	// Experimental.
	VirtualFolder *string `field:"optional" json:"virtualFolder" yaml:"virtualFolder"`
	// https://yarnpkg.com/configuration/yarnrc#winLinkType.
	// Experimental.
	WinLinkType YarnWinLinkType `field:"optional" json:"winLinkType" yaml:"winLinkType"`
	// https://yarnpkg.com/configuration/yarnrc#workerPoolMode.
	// Experimental.
	WorkerPoolMode YarnWorkerPoolMode `field:"optional" json:"workerPoolMode" yaml:"workerPoolMode"`
	// https://yarnpkg.com/configuration/yarnrc#yarnPath.
	// Experimental.
	YarnPath *string `field:"optional" json:"yarnPath" yaml:"yarnPath"`
}

