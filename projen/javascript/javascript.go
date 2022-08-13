package javascript

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"projen.javascript.AddBundleOptions",
		reflect.TypeOf((*AddBundleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.ArrowParens",
		reflect.TypeOf((*ArrowParens)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ArrowParens_ALWAYS,
			"AVOID": ArrowParens_AVOID,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.AutoRelease",
		reflect.TypeOf((*AutoRelease)(nil)).Elem(),
		map[string]interface{}{
			"EVERY_COMMIT": AutoRelease_EVERY_COMMIT,
			"DAILY": AutoRelease_DAILY,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.Bundle",
		reflect.TypeOf((*Bundle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.Bundler",
		reflect.TypeOf((*Bundler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBundle", GoMethod: "AddBundle"},
			_jsii_.MemberProperty{JsiiProperty: "bundledir", GoGetter: "Bundledir"},
			_jsii_.MemberProperty{JsiiProperty: "bundleTask", GoGetter: "BundleTask"},
			_jsii_.MemberProperty{JsiiProperty: "esbuildVersion", GoGetter: "EsbuildVersion"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Bundler{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.BundlerOptions",
		reflect.TypeOf((*BundlerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.BundlingOptions",
		reflect.TypeOf((*BundlingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.CodeArtifactOptions",
		reflect.TypeOf((*CodeArtifactOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.CoverageThreshold",
		reflect.TypeOf((*CoverageThreshold)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.EmbeddedLanguageFormatting",
		reflect.TypeOf((*EmbeddedLanguageFormatting)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": EmbeddedLanguageFormatting_AUTO,
			"OFF": EmbeddedLanguageFormatting_OFF,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.EndOfLine",
		reflect.TypeOf((*EndOfLine)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": EndOfLine_AUTO,
			"CR": EndOfLine_CR,
			"CRLF": EndOfLine_CRLF,
			"LF": EndOfLine_LF,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.Eslint",
		reflect.TypeOf((*Eslint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExtends", GoMethod: "AddExtends"},
			_jsii_.MemberMethod{JsiiMethod: "addIgnorePattern", GoMethod: "AddIgnorePattern"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPlugins", GoMethod: "AddPlugins"},
			_jsii_.MemberMethod{JsiiMethod: "addRules", GoMethod: "AddRules"},
			_jsii_.MemberMethod{JsiiMethod: "allowDevDeps", GoMethod: "AllowDevDeps"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "ignorePatterns", GoGetter: "IgnorePatterns"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Eslint{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.EslintOptions",
		reflect.TypeOf((*EslintOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.EslintOverride",
		reflect.TypeOf((*EslintOverride)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.HTMLWhitespaceSensitivity",
		reflect.TypeOf((*HTMLWhitespaceSensitivity)(nil)).Elem(),
		map[string]interface{}{
			"CSS": HTMLWhitespaceSensitivity_CSS,
			"IGNORE": HTMLWhitespaceSensitivity_IGNORE,
			"STRICT": HTMLWhitespaceSensitivity_STRICT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.HasteConfig",
		reflect.TypeOf((*HasteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.Jest",
		reflect.TypeOf((*Jest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIgnorePattern", GoMethod: "AddIgnorePattern"},
			_jsii_.MemberMethod{JsiiMethod: "addReporter", GoMethod: "AddReporter"},
			_jsii_.MemberMethod{JsiiMethod: "addSnapshotResolver", GoMethod: "AddSnapshotResolver"},
			_jsii_.MemberMethod{JsiiMethod: "addTestMatch", GoMethod: "AddTestMatch"},
			_jsii_.MemberMethod{JsiiMethod: "addWatchIgnorePattern", GoMethod: "AddWatchIgnorePattern"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "jestVersion", GoGetter: "JestVersion"},
		},
		func() interface{} {
			return &jsiiProxy_Jest{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.JestConfigOptions",
		reflect.TypeOf((*JestConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.JestOptions",
		reflect.TypeOf((*JestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.NodePackage",
		reflect.TypeOf((*NodePackage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBin", GoMethod: "AddBin"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addEngine", GoMethod: "AddEngine"},
			_jsii_.MemberMethod{JsiiMethod: "addField", GoMethod: "AddField"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageResolutions", GoMethod: "AddPackageResolutions"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addVersion", GoMethod: "AddVersion"},
			_jsii_.MemberProperty{JsiiProperty: "allowLibraryDependencies", GoGetter: "AllowLibraryDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "codeArtifactOptions", GoGetter: "CodeArtifactOptions"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "installAndUpdateLockfileCommand", GoGetter: "InstallAndUpdateLockfileCommand"},
			_jsii_.MemberProperty{JsiiProperty: "installCommand", GoGetter: "InstallCommand"},
			_jsii_.MemberProperty{JsiiProperty: "license", GoGetter: "License"},
			_jsii_.MemberProperty{JsiiProperty: "lockFile", GoGetter: "LockFile"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "npmAccess", GoGetter: "NpmAccess"},
			_jsii_.MemberProperty{JsiiProperty: "npmRegistry", GoGetter: "NpmRegistry"},
			_jsii_.MemberProperty{JsiiProperty: "npmRegistryUrl", GoGetter: "NpmRegistryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "npmTokenSecret", GoGetter: "NpmTokenSecret"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageName", GoGetter: "PackageName"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberMethod{JsiiMethod: "renderUpgradePackagesCommand", GoMethod: "RenderUpgradePackagesCommand"},
			_jsii_.MemberProperty{JsiiProperty: "scopedPackagesOptions", GoGetter: "ScopedPackagesOptions"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_NodePackage{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.NodePackageManager",
		reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		map[string]interface{}{
			"YARN": NodePackageManager_YARN,
			"NPM": NodePackageManager_NPM,
			"PNPM": NodePackageManager_PNPM,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.NodePackageOptions",
		reflect.TypeOf((*NodePackageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.NodeProject",
		reflect.TypeOf((*NodeProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBins", GoMethod: "AddBins"},
			_jsii_.MemberMethod{JsiiMethod: "addBundledDeps", GoMethod: "AddBundledDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addCompileCommand", GoMethod: "AddCompileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addDeps", GoMethod: "AddDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addDevDeps", GoMethod: "AddDevDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addExcludeFromCleanup", GoMethod: "AddExcludeFromCleanup"},
			_jsii_.MemberMethod{JsiiMethod: "addFields", GoMethod: "AddFields"},
			_jsii_.MemberMethod{JsiiMethod: "addGitIgnore", GoMethod: "AddGitIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addKeywords", GoMethod: "AddKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "addPackageIgnore", GoMethod: "AddPackageIgnore"},
			_jsii_.MemberMethod{JsiiMethod: "addPeerDeps", GoMethod: "AddPeerDeps"},
			_jsii_.MemberMethod{JsiiMethod: "addTask", GoMethod: "AddTask"},
			_jsii_.MemberMethod{JsiiMethod: "addTestCommand", GoMethod: "AddTestCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addTip", GoMethod: "AddTip"},
			_jsii_.MemberProperty{JsiiProperty: "allowLibraryDependencies", GoGetter: "AllowLibraryDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "annotateGenerated", GoMethod: "AnnotateGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsDirectory", GoGetter: "ArtifactsDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsJavascriptDirectory", GoGetter: "ArtifactsJavascriptDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "autoApprove", GoGetter: "AutoApprove"},
			_jsii_.MemberProperty{JsiiProperty: "autoMerge", GoGetter: "AutoMerge"},
			_jsii_.MemberProperty{JsiiProperty: "buildTask", GoGetter: "BuildTask"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflow", GoGetter: "BuildWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "buildWorkflowJobId", GoGetter: "BuildWorkflowJobId"},
			_jsii_.MemberProperty{JsiiProperty: "bundler", GoGetter: "Bundler"},
			_jsii_.MemberProperty{JsiiProperty: "commitGenerated", GoGetter: "CommitGenerated"},
			_jsii_.MemberProperty{JsiiProperty: "compileTask", GoGetter: "CompileTask"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTask", GoGetter: "DefaultTask"},
			_jsii_.MemberProperty{JsiiProperty: "deps", GoGetter: "Deps"},
			_jsii_.MemberProperty{JsiiProperty: "devContainer", GoGetter: "DevContainer"},
			_jsii_.MemberProperty{JsiiProperty: "ejected", GoGetter: "Ejected"},
			_jsii_.MemberProperty{JsiiProperty: "entrypoint", GoGetter: "Entrypoint"},
			_jsii_.MemberProperty{JsiiProperty: "files", GoGetter: "Files"},
			_jsii_.MemberProperty{JsiiProperty: "gitattributes", GoGetter: "Gitattributes"},
			_jsii_.MemberProperty{JsiiProperty: "github", GoGetter: "Github"},
			_jsii_.MemberProperty{JsiiProperty: "gitignore", GoGetter: "Gitignore"},
			_jsii_.MemberProperty{JsiiProperty: "gitpod", GoGetter: "Gitpod"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "initProject", GoGetter: "InitProject"},
			_jsii_.MemberProperty{JsiiProperty: "jest", GoGetter: "Jest"},
			_jsii_.MemberProperty{JsiiProperty: "logger", GoGetter: "Logger"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "package", GoGetter: "Package"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageTask", GoGetter: "PackageTask"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "postCompileTask", GoGetter: "PostCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "preCompileTask", GoGetter: "PreCompileTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "prettier", GoGetter: "Prettier"},
			_jsii_.MemberProperty{JsiiProperty: "projectBuild", GoGetter: "ProjectBuild"},
			_jsii_.MemberProperty{JsiiProperty: "projectType", GoGetter: "ProjectType"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberProperty{JsiiProperty: "publisher", GoGetter: "Publisher"},
			_jsii_.MemberProperty{JsiiProperty: "release", GoGetter: "Release"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberMethod{JsiiMethod: "removeTask", GoMethod: "RemoveTask"},
			_jsii_.MemberMethod{JsiiMethod: "renderWorkflowSetup", GoMethod: "RenderWorkflowSetup"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "runScriptCommand", GoGetter: "RunScriptCommand"},
			_jsii_.MemberMethod{JsiiMethod: "runTaskCommand", GoMethod: "RunTaskCommand"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeWorkflow", GoGetter: "UpgradeWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
		},
		func() interface{} {
			j := jsiiProxy_NodeProject{}
			_jsii_.InitJsiiProxy(&j.Type__githubGitHubProject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.NodeProjectOptions",
		reflect.TypeOf((*NodeProjectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.NpmAccess",
		reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": NpmAccess_PUBLIC,
			"RESTRICTED": NpmAccess_RESTRICTED,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.NpmConfig",
		reflect.TypeOf((*NpmConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfig", GoMethod: "AddConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addRegistry", GoMethod: "AddRegistry"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_NpmConfig{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.NpmConfigOptions",
		reflect.TypeOf((*NpmConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.PeerDependencyOptions",
		reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.Prettier",
		reflect.TypeOf((*Prettier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIgnorePattern", GoMethod: "AddIgnorePattern"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreFile", GoGetter: "IgnoreFile"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Prettier{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.PrettierOptions",
		reflect.TypeOf((*PrettierOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.PrettierOverride",
		reflect.TypeOf((*PrettierOverride)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.PrettierSettings",
		reflect.TypeOf((*PrettierSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.Projenrc",
		reflect.TypeOf((*Projenrc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_Projenrc{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.ProjenrcOptions",
		reflect.TypeOf((*ProjenrcOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.ProseWrap",
		reflect.TypeOf((*ProseWrap)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ProseWrap_ALWAYS,
			"NEVER": ProseWrap_NEVER,
			"PRESERVE": ProseWrap_PRESERVE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.QuoteProps",
		reflect.TypeOf((*QuoteProps)(nil)).Elem(),
		map[string]interface{}{
			"ASNEEDED": QuoteProps_ASNEEDED,
			"CONSISTENT": QuoteProps_CONSISTENT,
			"PRESERVE": QuoteProps_PRESERVE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.RenderWorkflowSetupOptions",
		reflect.TypeOf((*RenderWorkflowSetupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.ScopedPackagesOptions",
		reflect.TypeOf((*ScopedPackagesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.TrailingComma",
		reflect.TypeOf((*TrailingComma)(nil)).Elem(),
		map[string]interface{}{
			"ALL": TrailingComma_ALL,
			"ES5": TrailingComma_ES5,
			"NONE": TrailingComma_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.TypeScriptCompilerOptions",
		reflect.TypeOf((*TypeScriptCompilerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.TypeScriptJsxMode",
		reflect.TypeOf((*TypeScriptJsxMode)(nil)).Elem(),
		map[string]interface{}{
			"PRESERVE": TypeScriptJsxMode_PRESERVE,
			"REACT": TypeScriptJsxMode_REACT,
			"REACT_NATIVE": TypeScriptJsxMode_REACT_NATIVE,
			"REACT_JSX": TypeScriptJsxMode_REACT_JSX,
			"REACT_JSXDEV": TypeScriptJsxMode_REACT_JSXDEV,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.TypeScriptModuleResolution",
		reflect.TypeOf((*TypeScriptModuleResolution)(nil)).Elem(),
		map[string]interface{}{
			"CLASSIC": TypeScriptModuleResolution_CLASSIC,
			"NODE": TypeScriptModuleResolution_NODE,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.TypescriptConfig",
		reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExclude", GoMethod: "AddExclude"},
			_jsii_.MemberMethod{JsiiMethod: "addInclude", GoMethod: "AddInclude"},
			_jsii_.MemberProperty{JsiiProperty: "compilerOptions", GoGetter: "CompilerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
		},
		func() interface{} {
			return &jsiiProxy_TypescriptConfig{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.TypescriptConfigOptions",
		reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.UpgradeDependencies",
		reflect.TypeOf((*UpgradeDependencies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildSteps", GoMethod: "AddPostBuildSteps"},
			_jsii_.MemberProperty{JsiiProperty: "containerOptions", GoGetter: "ContainerOptions"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "postUpgradeTask", GoGetter: "PostUpgradeTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeTask", GoGetter: "UpgradeTask"},
			_jsii_.MemberProperty{JsiiProperty: "workflows", GoGetter: "Workflows"},
		},
		func() interface{} {
			j := jsiiProxy_UpgradeDependencies{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.UpgradeDependenciesOptions",
		reflect.TypeOf((*UpgradeDependenciesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"projen.javascript.UpgradeDependenciesSchedule",
		reflect.TypeOf((*UpgradeDependenciesSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cron", GoGetter: "Cron"},
		},
		func() interface{} {
			return &jsiiProxy_UpgradeDependenciesSchedule{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.UpgradeDependenciesWorkflowOptions",
		reflect.TypeOf((*UpgradeDependenciesWorkflowOptions)(nil)).Elem(),
	)
}
