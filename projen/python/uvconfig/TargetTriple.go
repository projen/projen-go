package uvconfig


// The supported target triples. Each triple consists of an architecture, vendor, and operating system.
//
// See: <https://doc.rust-lang.org/nightly/rustc/platform-support.html>
// Experimental.
type TargetTriple string

const (
	// An alias for `x86_64-pc-windows-msvc`, the default target for Windows.
	//
	// (windows).
	// Experimental.
	TargetTriple_WINDOWS TargetTriple = "WINDOWS"
	// An alias for `x86_64-unknown-linux-gnu`, the default target for Linux.
	//
	// (linux).
	// Experimental.
	TargetTriple_LINUX TargetTriple = "LINUX"
	// An alias for `aarch64-apple-darwin`, the default target for macOS.
	//
	// (macos).
	// Experimental.
	TargetTriple_MACOS TargetTriple = "MACOS"
	// A 64-bit x86 Windows target.
	//
	// (x86_64-pc-windows-msvc).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC TargetTriple = "X86_UNDERSCORE_64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC"
	// An ARM64 Windows target.
	//
	// (aarch64-pc-windows-msvc).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC TargetTriple = "AARCH64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC"
	// A 32-bit x86 Windows target.
	//
	// (i686-pc-windows-msvc).
	// Experimental.
	TargetTriple_I686_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC TargetTriple = "I686_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC"
	// An x86 Linux target.
	//
	// Equivalent to `x86_64-manylinux_2_28`. (x86_64-unknown-linux-gnu)
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU TargetTriple = "X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU"
	// An ARM-based macOS target, as seen on Apple Silicon devices.
	//
	// By default, assumes the least-recent, non-EOL macOS version (13.0), but respects
	// the `MACOSX_DEPLOYMENT_TARGET` environment variable if set. (aarch64-apple-darwin)
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_APPLE_HYPHEN_DARWIN TargetTriple = "AARCH64_HYPHEN_APPLE_HYPHEN_DARWIN"
	// An x86 macOS target.
	//
	// By default, assumes the least-recent, non-EOL macOS version (13.0), but respects
	// the `MACOSX_DEPLOYMENT_TARGET` environment variable if set. (x86_64-apple-darwin)
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_DARWIN TargetTriple = "X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_DARWIN"
	// An ARM64 Linux target.
	//
	// Equivalent to `aarch64-manylinux_2_28`. (aarch64-unknown-linux-gnu)
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU TargetTriple = "AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU"
	// An ARM64 Linux target.
	//
	// (aarch64-unknown-linux-musl).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL TargetTriple = "AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL"
	// An `x86_64` Linux target.
	//
	// (x86_64-unknown-linux-musl).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL TargetTriple = "X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL"
	// A RISCV64 Linux target.
	//
	// (riscv64-unknown-linux).
	// Experimental.
	TargetTriple_RISCV64_HYPHEN_UNKNOWN_HYPHEN_LINUX TargetTriple = "RISCV64_HYPHEN_UNKNOWN_HYPHEN_LINUX"
	// An `x86_64` target for the `manylinux2014` platform.
	//
	// Equivalent to `x86_64-manylinux_2_17`. (x86_64-manylinux2014)
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX2014 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX2014"
	// An `x86_64` target for the `manylinux_2_17` platform.
	//
	// (x86_64-manylinux_2_17).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17"
	// An `x86_64` target for the `manylinux_2_28` platform.
	//
	// (x86_64-manylinux_2_28).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28"
	// An `x86_64` target for the `manylinux_2_31` platform.
	//
	// (x86_64-manylinux_2_31).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31"
	// An `x86_64` target for the `manylinux_2_32` platform.
	//
	// (x86_64-manylinux_2_32).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32"
	// An `x86_64` target for the `manylinux_2_33` platform.
	//
	// (x86_64-manylinux_2_33).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33"
	// An `x86_64` target for the `manylinux_2_34` platform.
	//
	// (x86_64-manylinux_2_34).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34"
	// An `x86_64` target for the `manylinux_2_35` platform.
	//
	// (x86_64-manylinux_2_35).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35"
	// An `x86_64` target for the `manylinux_2_36` platform.
	//
	// (x86_64-manylinux_2_36).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36"
	// An `x86_64` target for the `manylinux_2_37` platform.
	//
	// (x86_64-manylinux_2_37).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37"
	// An `x86_64` target for the `manylinux_2_38` platform.
	//
	// (x86_64-manylinux_2_38).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38"
	// An `x86_64` target for the `manylinux_2_39` platform.
	//
	// (x86_64-manylinux_2_39).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39"
	// An `x86_64` target for the `manylinux_2_40` platform.
	//
	// (x86_64-manylinux_2_40).
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40 TargetTriple = "X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40"
	// An ARM64 target for the `manylinux2014` platform.
	//
	// Equivalent to `aarch64-manylinux_2_17`. (aarch64-manylinux2014)
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX2014 TargetTriple = "AARCH64_HYPHEN_MANYLINUX2014"
	// An ARM64 target for the `manylinux_2_17` platform.
	//
	// (aarch64-manylinux_2_17).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17"
	// An ARM64 target for the `manylinux_2_28` platform.
	//
	// (aarch64-manylinux_2_28).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28"
	// An ARM64 target for the `manylinux_2_31` platform.
	//
	// (aarch64-manylinux_2_31).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31"
	// An ARM64 target for the `manylinux_2_32` platform.
	//
	// (aarch64-manylinux_2_32).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32"
	// An ARM64 target for the `manylinux_2_33` platform.
	//
	// (aarch64-manylinux_2_33).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33"
	// An ARM64 target for the `manylinux_2_34` platform.
	//
	// (aarch64-manylinux_2_34).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34"
	// An ARM64 target for the `manylinux_2_35` platform.
	//
	// (aarch64-manylinux_2_35).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35"
	// An ARM64 target for the `manylinux_2_36` platform.
	//
	// (aarch64-manylinux_2_36).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36"
	// An ARM64 target for the `manylinux_2_37` platform.
	//
	// (aarch64-manylinux_2_37).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37"
	// An ARM64 target for the `manylinux_2_38` platform.
	//
	// (aarch64-manylinux_2_38).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38"
	// An ARM64 target for the `manylinux_2_39` platform.
	//
	// (aarch64-manylinux_2_39).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39"
	// An ARM64 target for the `manylinux_2_40` platform.
	//
	// (aarch64-manylinux_2_40).
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40 TargetTriple = "AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40"
	// An ARM64 Android target.
	//
	// By default uses Android API level 24, but respects
	// the `ANDROID_API_LEVEL` environment variable if set. (aarch64-linux-android)
	// Experimental.
	TargetTriple_AARCH64_HYPHEN_LINUX_HYPHEN_ANDROID TargetTriple = "AARCH64_HYPHEN_LINUX_HYPHEN_ANDROID"
	// An `x86_64` Android target.
	//
	// By default uses Android API level 24, but respects
	// the `ANDROID_API_LEVEL` environment variable if set. (x86_64-linux-android)
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_LINUX_HYPHEN_ANDROID TargetTriple = "X86_UNDERSCORE_64_HYPHEN_LINUX_HYPHEN_ANDROID"
	// A wasm32 target using the Pyodide 2024 platform.
	//
	// Meant for use with Python 3.12. (wasm32-pyodide2024)
	// Experimental.
	TargetTriple_WASM32_HYPHEN_PYODIDE2024 TargetTriple = "WASM32_HYPHEN_PYODIDE2024"
	// An ARM64 target for iOS device.
	//
	// By default, iOS 13.0 is used, but respects the `IPHONEOS_DEPLOYMENT_TARGET`
	// environment variable if set. (arm64-apple-ios)
	// Experimental.
	TargetTriple_ARM64_HYPHEN_APPLE_HYPHEN_IOS TargetTriple = "ARM64_HYPHEN_APPLE_HYPHEN_IOS"
	// An ARM64 target for iOS simulator.
	//
	// By default, iOS 13.0 is used, but respects the `IPHONEOS_DEPLOYMENT_TARGET`
	// environment variable if set. (arm64-apple-ios-simulator)
	// Experimental.
	TargetTriple_ARM64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR TargetTriple = "ARM64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR"
	// An `x86_64` target for iOS simulator.
	//
	// By default, iOS 13.0 is used, but respects the `IPHONEOS_DEPLOYMENT_TARGET`
	// environment variable if set. (x86_64-apple-ios-simulator)
	// Experimental.
	TargetTriple_X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR TargetTriple = "X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR"
)

