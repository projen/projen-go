package javascript


// Options for `addBundle()`.
// Experimental.
type AddBundleOptions struct {
	// You can mark a file or a package as external to exclude it from your build.
	//
	// Instead of being bundled, the import will be preserved (using require for
	// the iife and cjs formats and using import for the esm format) and will be
	// evaluated at run time instead.
	//
	// This has several uses. First of all, it can be used to trim unnecessary
	// code from your bundle for a code path that you know will never be executed.
	// For example, a package may contain code that only runs in node but you will
	// only be using that package in the browser. It can also be used to import
	// code in node at run time from a package that cannot be bundled. For
	// example, the fsevents package contains a native extension, which esbuild
	// doesn't support.
	// Default: [].
	//
	// Experimental.
	Externals *[]*string `field:"optional" json:"externals" yaml:"externals"`
	// Include a source map in the bundle.
	// Default: false.
	//
	// Experimental.
	Sourcemap *bool `field:"optional" json:"sourcemap" yaml:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Default: true.
	//
	// Experimental.
	WatchTask *bool `field:"optional" json:"watchTask" yaml:"watchTask"`
	// esbuild platform.
	//
	// Example:
	//   "node"
	//
	// Experimental.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// esbuild target.
	//
	// Example:
	//   "node12"
	//
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Use this to insert an arbitrary string at the beginning of generated JavaScript files.
	//
	// This is similar to footer which inserts at the end instead of the beginning.
	//
	// This is commonly used to insert comments:.
	// Default: - no comments are passed.
	//
	// Experimental.
	Banner *string `field:"optional" json:"banner" yaml:"banner"`
	// The charset to use for esbuild's output.
	//
	// By default esbuild's output is ASCII-only. Any non-ASCII characters are escaped
	// using backslash escape sequences. Using escape sequences makes the generated output
	// slightly bigger, and also makes it harder to read. If you would like for esbuild to print
	// the original characters without using escape sequences, use `Charset.UTF8`.
	// See: https://esbuild.github.io/api/#charset
	//
	// Default: Charset.ASCII
	//
	// Experimental.
	Charset Charset `field:"optional" json:"charset" yaml:"charset"`
	// Replace global identifiers with constant expressions.
	//
	// For example, `{ 'process.env.DEBUG': 'true' }`.
	//
	// Another example, `{ 'process.env.API_KEY': JSON.stringify('xxx-xxxx-xxx') }`.
	// Default: - no replacements are made.
	//
	// Experimental.
	Define *map[string]*string `field:"optional" json:"define" yaml:"define"`
	// Build arguments to pass into esbuild.
	//
	// For example, to add the [--log-limit](https://esbuild.github.io/api/#log-limit) flag:
	//
	// ```text
	// project.bundler.addBundle("./src/hello.ts", {
	//   platform: "node",
	//   target: "node18",
	//   sourcemap: true,
	//   format: "esm",
	//   esbuildArgs: {
	//     "--log-limit": "0",
	//   },
	// });
	// ```.
	// Default: - no additional esbuild arguments are passed.
	//
	// Experimental.
	EsbuildArgs *map[string]interface{} `field:"optional" json:"esbuildArgs" yaml:"esbuildArgs"`
	// Mark the output file as executable.
	// Default: false.
	//
	// Experimental.
	Executable *bool `field:"optional" json:"executable" yaml:"executable"`
	// Use this to insert an arbitrary string at the end of generated JavaScript files.
	//
	// This is similar to banner which inserts at the beginning instead of the end.
	//
	// This is commonly used to insert comments.
	// Default: - no comments are passed.
	//
	// Experimental.
	Footer *string `field:"optional" json:"footer" yaml:"footer"`
	// Output format for the generated JavaScript files.
	//
	// There are currently three possible values that can be configured: `"iife"`, `"cjs"`, and `"esm"`.
	//
	// If not set (`undefined`), esbuild picks an output format for you based on `platform`:
	// - `"cjs"` if `platform` is `"node"`
	// - `"iife"` if `platform` is `"browser"`
	// - `"esm"` if `platform` is `"neutral"`
	//
	// Note: If making a bundle to run under node with ESM, set `format` to `"esm"` instead of setting `platform` to `"neutral"`.
	// See: https://esbuild.github.io/api/#format
	//
	// Default: undefined.
	//
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// This option allows you to automatically replace a global variable with an import from another file.
	// See: https://esbuild.github.io/api/#inject
	//
	// Default: - no code is injected.
	//
	// Experimental.
	Inject *[]*string `field:"optional" json:"inject" yaml:"inject"`
	// Whether to preserve the original `name` values even in minified code.
	//
	// In JavaScript the `name` property on functions and classes defaults to a
	// nearby identifier in the source code.
	//
	// However, minification renames symbols to reduce code size and bundling
	// sometimes need to rename symbols to avoid collisions. That changes value of
	// the `name` property for many of these cases. This is usually fine because
	// the `name` property is normally only used for debugging. However, some
	// frameworks rely on the `name` property for registration and binding purposes.
	// If this is the case, you can enable this option to preserve the original
	// `name` values even in minified code.
	// Default: false.
	//
	// Experimental.
	KeepNames *bool `field:"optional" json:"keepNames" yaml:"keepNames"`
	// Map of file extensions (without dot) and loaders to use for this file type.
	//
	// Loaders are appended to the esbuild command by `--loader:.extension=loader`
	// Experimental.
	Loaders *map[string]*string `field:"optional" json:"loaders" yaml:"loaders"`
	// Log level for esbuild.
	//
	// This is also propagated to the package manager and
	// applies to its specific install command.
	// Default: LogLevel.WARNING
	//
	// Experimental.
	LogLevel BundleLogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// How to determine the entry point for modules.
	//
	// Try ['module', 'main'] to default to ES module versions.
	// Default: [].
	//
	// Experimental.
	MainFields *[]*string `field:"optional" json:"mainFields" yaml:"mainFields"`
	// This option tells esbuild to write out a JSON file relative to output directory with metadata about the build.
	//
	// The metadata in this JSON file follows this schema (specified using TypeScript syntax):
	//
	// ```text
	// {
	//   outputs: {
	//     [path: string]: {
	//       bytes: number
	//       inputs: {
	//         [path: string]: { bytesInOutput: number }
	//       }
	//       imports: { path: string }[]
	//       exports: string[]
	//     }
	//   }
	// }
	// ```
	// This data can then be analyzed by other tools. For example,
	// bundle buddy can consume esbuild's metadata format and generates a treemap visualization
	// of the modules in your bundle and how much space each one takes up.
	// See: https://esbuild.github.io/api/#metafile
	//
	// Default: false.
	//
	// Experimental.
	Metafile *bool `field:"optional" json:"metafile" yaml:"metafile"`
	// Whether to minify files when bundling.
	// Default: false.
	//
	// Experimental.
	Minify *bool `field:"optional" json:"minify" yaml:"minify"`
	// Bundler output path relative to the asset's output directory.
	// Default: "index.js"
	//
	// Experimental.
	Outfile *string `field:"optional" json:"outfile" yaml:"outfile"`
	// Source map mode to be used when bundling.
	// See: https://esbuild.github.io/api/#sourcemap
	//
	// Default: SourceMapMode.DEFAULT
	//
	// Experimental.
	SourceMapMode SourceMapMode `field:"optional" json:"sourceMapMode" yaml:"sourceMapMode"`
	// Whether to include original source code in source maps when bundling.
	// See: https://esbuild.github.io/api/#sources-content
	//
	// Default: true.
	//
	// Experimental.
	SourcesContent *bool `field:"optional" json:"sourcesContent" yaml:"sourcesContent"`
	// The path of the tsconfig.json file to use for bundling.
	// Default: "tsconfig.json"
	//
	// Experimental.
	TsconfigPath *string `field:"optional" json:"tsconfigPath" yaml:"tsconfigPath"`
}

