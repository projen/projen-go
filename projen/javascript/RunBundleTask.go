package javascript


// Options for BundlerOptions.runBundleTask.
// Experimental.
type RunBundleTask string

const (
	// Don't bundle automatically as part of the build.
	// Experimental.
	RunBundleTask_MANUAL RunBundleTask = "MANUAL"
	// Bundle automatically before compilation.
	// Experimental.
	RunBundleTask_PRE_COMPILE RunBundleTask = "PRE_COMPILE"
	// Bundle automatically after compilation. This is useful if you want to bundle the compiled results.
	//
	// Thus will run compilation tasks (using tsc, etc.) before running file
	// through bundling step.
	//
	// This is only required unless you are using new experimental features that
	// are not supported by `esbuild` but are supported by typescript's `tsc`
	// compiler. One example of such feature is `emitDecoratorMetadata`.
	//
	// ```typescript
	// // In a TypeScript project with output configured
	// // to go to the "lib" directory:
	// const project = new TypeScriptProject({
	//   name: "test",
	//   defaultReleaseBranch: "main",
	//   tsconfig: {
	//     compilerOptions: {
	//       outDir: "lib",
	//     },
	//   },
	//   bundlerOptions: {
	//     // ensure we compile with `tsc` before bundling
	//     runBundleTask: RunBundleTask.POST_COMPILE,
	//   },
	// });
	//
	// // Tell the bundler to bundle the compiled results (from the "lib" directory)
	// project.bundler.addBundle("./lib/index.js", {
	//   platform: "node",
	//   target: "node18",
	//   sourcemap: false,
	//   format: "esm",
	// });
	// ```.
	// Experimental.
	RunBundleTask_POST_COMPILE RunBundleTask = "POST_COMPILE"
)

