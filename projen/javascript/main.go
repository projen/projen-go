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
		"projen.javascript.BuildWorkflowOptions",
		reflect.TypeOf((*BuildWorkflowOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.Bundle",
		reflect.TypeOf((*Bundle)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.BundleLogLevel",
		reflect.TypeOf((*BundleLogLevel)(nil)).Elem(),
		map[string]interface{}{
			"VERBOSE": BundleLogLevel_VERBOSE,
			"DEBUG": BundleLogLevel_DEBUG,
			"INFO": BundleLogLevel_INFO,
			"WARNING": BundleLogLevel_WARNING,
			"ERROR": BundleLogLevel_ERROR,
			"SILENT": BundleLogLevel_SILENT,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.Bundler",
		reflect.TypeOf((*Bundler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBundle", GoMethod: "AddBundle"},
			_jsii_.MemberProperty{JsiiProperty: "bundledir", GoGetter: "Bundledir"},
			_jsii_.MemberProperty{JsiiProperty: "bundleTask", GoGetter: "BundleTask"},
			_jsii_.MemberProperty{JsiiProperty: "esbuildVersion", GoGetter: "EsbuildVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
	_jsii_.RegisterEnum(
		"projen.javascript.Charset",
		reflect.TypeOf((*Charset)(nil)).Elem(),
		map[string]interface{}{
			"ASCII": Charset_ASCII,
			"UTF8": Charset_UTF8,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.CodeArtifactAuthProvider",
		reflect.TypeOf((*CodeArtifactAuthProvider)(nil)).Elem(),
		map[string]interface{}{
			"ACCESS_AND_SECRET_KEY_PAIR": CodeArtifactAuthProvider_ACCESS_AND_SECRET_KEY_PAIR,
			"GITHUB_OIDC": CodeArtifactAuthProvider_GITHUB_OIDC,
		},
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
			_jsii_.MemberMethod{JsiiMethod: "addLintPattern", GoMethod: "AddLintPattern"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPlugins", GoMethod: "AddPlugins"},
			_jsii_.MemberMethod{JsiiMethod: "addRules", GoMethod: "AddRules"},
			_jsii_.MemberMethod{JsiiMethod: "allowDevDeps", GoMethod: "AllowDevDeps"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "eslintTask", GoGetter: "EslintTask"},
			_jsii_.MemberProperty{JsiiProperty: "ignorePatterns", GoGetter: "IgnorePatterns"},
			_jsii_.MemberProperty{JsiiProperty: "lintPatterns", GoGetter: "LintPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberMethod{JsiiMethod: "addModuleNameMappers", GoMethod: "AddModuleNameMappers"},
			_jsii_.MemberMethod{JsiiMethod: "addModulePaths", GoMethod: "AddModulePaths"},
			_jsii_.MemberMethod{JsiiMethod: "addReporter", GoMethod: "AddReporter"},
			_jsii_.MemberMethod{JsiiMethod: "addRoots", GoMethod: "AddRoots"},
			_jsii_.MemberMethod{JsiiMethod: "addSetupFile", GoMethod: "AddSetupFile"},
			_jsii_.MemberMethod{JsiiMethod: "addSetupFileAfterEnv", GoMethod: "AddSetupFileAfterEnv"},
			_jsii_.MemberMethod{JsiiMethod: "addSnapshotResolver", GoMethod: "AddSnapshotResolver"},
			_jsii_.MemberMethod{JsiiMethod: "addTestMatch", GoMethod: "AddTestMatch"},
			_jsii_.MemberMethod{JsiiMethod: "addWatchIgnorePattern", GoMethod: "AddWatchIgnorePattern"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "jestVersion", GoGetter: "JestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Jest{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
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
		"projen.javascript.JestReporter",
		reflect.TypeOf((*JestReporter)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JestReporter{}
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.LicenseChecker",
		reflect.TypeOf((*LicenseChecker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LicenseChecker{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.LicenseCheckerOptions",
		reflect.TypeOf((*LicenseCheckerOptions)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberMethod{JsiiMethod: "hasScript", GoMethod: "HasScript"},
			_jsii_.MemberProperty{JsiiProperty: "installAndUpdateLockfileCommand", GoGetter: "InstallAndUpdateLockfileCommand"},
			_jsii_.MemberProperty{JsiiProperty: "installCiTask", GoGetter: "InstallCiTask"},
			_jsii_.MemberProperty{JsiiProperty: "installCommand", GoGetter: "InstallCommand"},
			_jsii_.MemberProperty{JsiiProperty: "installTask", GoGetter: "InstallTask"},
			_jsii_.MemberProperty{JsiiProperty: "license", GoGetter: "License"},
			_jsii_.MemberProperty{JsiiProperty: "lockFile", GoGetter: "LockFile"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "maxNodeVersion", GoGetter: "MaxNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minNodeVersion", GoGetter: "MinNodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "npmAccess", GoGetter: "NpmAccess"},
			_jsii_.MemberProperty{JsiiProperty: "npmProvenance", GoGetter: "NpmProvenance"},
			_jsii_.MemberProperty{JsiiProperty: "npmRegistry", GoGetter: "NpmRegistry"},
			_jsii_.MemberProperty{JsiiProperty: "npmRegistryUrl", GoGetter: "NpmRegistryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "npmTokenSecret", GoGetter: "NpmTokenSecret"},
			_jsii_.MemberProperty{JsiiProperty: "packageManager", GoGetter: "PackageManager"},
			_jsii_.MemberProperty{JsiiProperty: "packageName", GoGetter: "PackageName"},
			_jsii_.MemberProperty{JsiiProperty: "pnpmVersion", GoGetter: "PnpmVersion"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projenCommand", GoGetter: "ProjenCommand"},
			_jsii_.MemberMethod{JsiiMethod: "removeScript", GoMethod: "RemoveScript"},
			_jsii_.MemberProperty{JsiiProperty: "scopedPackagesOptions", GoGetter: "ScopedPackagesOptions"},
			_jsii_.MemberMethod{JsiiMethod: "setScript", GoMethod: "SetScript"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryResolveDependencyVersion", GoMethod: "TryResolveDependencyVersion"},
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
			"YARN2": NodePackageManager_YARN2,
			"YARN_CLASSIC": NodePackageManager_YARN_CLASSIC,
			"YARN_BERRY": NodePackageManager_YARN_BERRY,
			"NPM": NodePackageManager_NPM,
			"PNPM": NodePackageManager_PNPM,
			"BUN": NodePackageManager_BUN,
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
			_jsii_.MemberMethod{JsiiMethod: "addScripts", GoMethod: "AddScripts"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeVersion", GoGetter: "NodeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "npmignore", GoGetter: "Npmignore"},
			_jsii_.MemberProperty{JsiiProperty: "npmrc", GoGetter: "Npmrc"},
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
			_jsii_.MemberProperty{JsiiProperty: "subprojects", GoGetter: "Subprojects"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberProperty{JsiiProperty: "testTask", GoGetter: "TestTask"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindFile", GoMethod: "TryFindFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindJsonFile", GoMethod: "TryFindJsonFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryFindObjectFile", GoMethod: "TryFindObjectFile"},
			_jsii_.MemberMethod{JsiiMethod: "tryRemoveFile", GoMethod: "TryRemoveFile"},
			_jsii_.MemberProperty{JsiiProperty: "upgradeWorkflow", GoGetter: "UpgradeWorkflow"},
			_jsii_.MemberProperty{JsiiProperty: "vscode", GoGetter: "Vscode"},
			_jsii_.MemberProperty{JsiiProperty: "workflowBootstrapSteps", GoGetter: "WorkflowBootstrapSteps"},
			_jsii_.MemberProperty{JsiiProperty: "workflowPackageCache", GoGetter: "WorkflowPackageCache"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "overrides", GoGetter: "Overrides"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Projenrc{}
			_jsii_.InitJsiiProxy(&j.Type__projenProjenrcFile)
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
	_jsii_.RegisterEnum(
		"projen.javascript.RunBundleTask",
		reflect.TypeOf((*RunBundleTask)(nil)).Elem(),
		map[string]interface{}{
			"MANUAL": RunBundleTask_MANUAL,
			"PRE_COMPILE": RunBundleTask_PRE_COMPILE,
			"POST_COMPILE": RunBundleTask_POST_COMPILE,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.ScopedPackagesOptions",
		reflect.TypeOf((*ScopedPackagesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.SourceMapMode",
		reflect.TypeOf((*SourceMapMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": SourceMapMode_DEFAULT,
			"EXTERNAL": SourceMapMode_EXTERNAL,
			"INLINE": SourceMapMode_INLINE,
			"BOTH": SourceMapMode_BOTH,
		},
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
	_jsii_.RegisterClass(
		"projen.javascript.Transform",
		reflect.TypeOf((*Transform)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Transform{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.TypeScriptCompilerOptions",
		reflect.TypeOf((*TypeScriptCompilerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.TypeScriptImportsNotUsedAsValues",
		reflect.TypeOf((*TypeScriptImportsNotUsedAsValues)(nil)).Elem(),
		map[string]interface{}{
			"REMOVE": TypeScriptImportsNotUsedAsValues_REMOVE,
			"PRESERVE": TypeScriptImportsNotUsedAsValues_PRESERVE,
			"ERROR": TypeScriptImportsNotUsedAsValues_ERROR,
		},
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
		"projen.javascript.TypeScriptModuleDetection",
		reflect.TypeOf((*TypeScriptModuleDetection)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": TypeScriptModuleDetection_AUTO,
			"LEGACY": TypeScriptModuleDetection_LEGACY,
			"FORCE": TypeScriptModuleDetection_FORCE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.TypeScriptModuleResolution",
		reflect.TypeOf((*TypeScriptModuleResolution)(nil)).Elem(),
		map[string]interface{}{
			"CLASSIC": TypeScriptModuleResolution_CLASSIC,
			"NODE": TypeScriptModuleResolution_NODE,
			"NODE16": TypeScriptModuleResolution_NODE16,
			"NODE_NEXT": TypeScriptModuleResolution_NODE_NEXT,
			"BUNDLER": TypeScriptModuleResolution_BUNDLER,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.TypescriptConfig",
		reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addExclude", GoMethod: "AddExclude"},
			_jsii_.MemberMethod{JsiiMethod: "addExtends", GoMethod: "AddExtends"},
			_jsii_.MemberMethod{JsiiMethod: "addInclude", GoMethod: "AddInclude"},
			_jsii_.MemberProperty{JsiiProperty: "compilerOptions", GoGetter: "CompilerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "extends", GoGetter: "Extends"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "removeExclude", GoMethod: "RemoveExclude"},
			_jsii_.MemberMethod{JsiiMethod: "removeInclude", GoMethod: "RemoveInclude"},
			_jsii_.MemberMethod{JsiiMethod: "resolveExtendsPath", GoMethod: "ResolveExtendsPath"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TypescriptConfig{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.TypescriptConfigExtends",
		reflect.TypeOf((*TypescriptConfigExtends)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJSON", GoMethod: "ToJSON"},
		},
		func() interface{} {
			return &jsiiProxy_TypescriptConfigExtends{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.TypescriptConfigOptions",
		reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.UpdateSnapshot",
		reflect.TypeOf((*UpdateSnapshot)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": UpdateSnapshot_ALWAYS,
			"NEVER": UpdateSnapshot_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.UpgradeDependencies",
		reflect.TypeOf((*UpgradeDependencies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPostBuildSteps", GoMethod: "AddPostBuildSteps"},
			_jsii_.MemberProperty{JsiiProperty: "containerOptions", GoGetter: "ContainerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "postUpgradeTask", GoGetter: "PostUpgradeTask"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
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
	_jsii_.RegisterClass(
		"projen.javascript.WatchPlugin",
		reflect.TypeOf((*WatchPlugin)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WatchPlugin{}
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnBerryOptions",
		reflect.TypeOf((*YarnBerryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnCacheMigrationMode",
		reflect.TypeOf((*YarnCacheMigrationMode)(nil)).Elem(),
		map[string]interface{}{
			"REQUIRED_ONLY": YarnCacheMigrationMode_REQUIRED_ONLY,
			"MATCH_SPEC": YarnCacheMigrationMode_MATCH_SPEC,
			"ALWAYS": YarnCacheMigrationMode_ALWAYS,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnChecksumBehavior",
		reflect.TypeOf((*YarnChecksumBehavior)(nil)).Elem(),
		map[string]interface{}{
			"THROW": YarnChecksumBehavior_THROW,
			"UPDATE": YarnChecksumBehavior_UPDATE,
			"RESET": YarnChecksumBehavior_RESET,
			"IGNORE": YarnChecksumBehavior_IGNORE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnDefaultSemverRangePrefix",
		reflect.TypeOf((*YarnDefaultSemverRangePrefix)(nil)).Elem(),
		map[string]interface{}{
			"CARET": YarnDefaultSemverRangePrefix_CARET,
			"TILDE": YarnDefaultSemverRangePrefix_TILDE,
			"EMPTY_STRING": YarnDefaultSemverRangePrefix_EMPTY_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnLogFilter",
		reflect.TypeOf((*YarnLogFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnLogFilterLevel",
		reflect.TypeOf((*YarnLogFilterLevel)(nil)).Elem(),
		map[string]interface{}{
			"INFO": YarnLogFilterLevel_INFO,
			"WARNING": YarnLogFilterLevel_WARNING,
			"ERROR": YarnLogFilterLevel_ERROR,
			"DISCARD": YarnLogFilterLevel_DISCARD,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnNetworkSetting",
		reflect.TypeOf((*YarnNetworkSetting)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnNmHoistingLimit",
		reflect.TypeOf((*YarnNmHoistingLimit)(nil)).Elem(),
		map[string]interface{}{
			"DEPENDENCIES": YarnNmHoistingLimit_DEPENDENCIES,
			"NONE": YarnNmHoistingLimit_NONE,
			"WORKSPACES": YarnNmHoistingLimit_WORKSPACES,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnNmMode",
		reflect.TypeOf((*YarnNmMode)(nil)).Elem(),
		map[string]interface{}{
			"CLASSIC": YarnNmMode_CLASSIC,
			"HARDLINKS_LOCAL": YarnNmMode_HARDLINKS_LOCAL,
			"HARDLINKS_GLOBAL": YarnNmMode_HARDLINKS_GLOBAL,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnNodeLinker",
		reflect.TypeOf((*YarnNodeLinker)(nil)).Elem(),
		map[string]interface{}{
			"PNP": YarnNodeLinker_PNP,
			"PNPM": YarnNodeLinker_PNPM,
			"NODE_MODULES": YarnNodeLinker_NODE_MODULES,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnNpmPublishAccess",
		reflect.TypeOf((*YarnNpmPublishAccess)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": YarnNpmPublishAccess_PUBLIC,
			"RESTRICTED": YarnNpmPublishAccess_RESTRICTED,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnNpmRegistry",
		reflect.TypeOf((*YarnNpmRegistry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnNpmScope",
		reflect.TypeOf((*YarnNpmScope)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnPackageExtension",
		reflect.TypeOf((*YarnPackageExtension)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnPeerDependencyMeta",
		reflect.TypeOf((*YarnPeerDependencyMeta)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnPnpFallbackMode",
		reflect.TypeOf((*YarnPnpFallbackMode)(nil)).Elem(),
		map[string]interface{}{
			"NONE": YarnPnpFallbackMode_NONE,
			"DEPENDENCIES_ONLY": YarnPnpFallbackMode_DEPENDENCIES_ONLY,
			"ALL": YarnPnpFallbackMode_ALL,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnPnpMode",
		reflect.TypeOf((*YarnPnpMode)(nil)).Elem(),
		map[string]interface{}{
			"STRICT": YarnPnpMode_STRICT,
			"LOOSE": YarnPnpMode_LOOSE,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnProgressBarStyle",
		reflect.TypeOf((*YarnProgressBarStyle)(nil)).Elem(),
		map[string]interface{}{
			"PATRICK": YarnProgressBarStyle_PATRICK,
			"SIMBA": YarnProgressBarStyle_SIMBA,
			"JACK": YarnProgressBarStyle_JACK,
			"HOGSFATHER": YarnProgressBarStyle_HOGSFATHER,
			"DEFAULT": YarnProgressBarStyle_DEFAULT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnSupportedArchitectures",
		reflect.TypeOf((*YarnSupportedArchitectures)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnWinLinkType",
		reflect.TypeOf((*YarnWinLinkType)(nil)).Elem(),
		map[string]interface{}{
			"JUNCTIONS": YarnWinLinkType_JUNCTIONS,
			"SYMLINKS": YarnWinLinkType_SYMLINKS,
		},
	)
	_jsii_.RegisterEnum(
		"projen.javascript.YarnWorkerPoolMode",
		reflect.TypeOf((*YarnWorkerPoolMode)(nil)).Elem(),
		map[string]interface{}{
			"ASYNC": YarnWorkerPoolMode_ASYNC,
			"WORKERS": YarnWorkerPoolMode_WORKERS,
		},
	)
	_jsii_.RegisterClass(
		"projen.javascript.Yarnrc",
		reflect.TypeOf((*Yarnrc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "postSynthesize", GoMethod: "PostSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "preSynthesize", GoMethod: "PreSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Yarnrc{}
			_jsii_.InitJsiiProxy(&j.Type__projenComponent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen.javascript.YarnrcOptions",
		reflect.TypeOf((*YarnrcOptions)(nil)).Elem(),
	)
}
