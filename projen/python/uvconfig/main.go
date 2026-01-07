package uvconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.AddBoundsKind",
		reflect.TypeOf((*AddBoundsKind)(nil)).Elem(),
		map[string]interface{}{
			"LOWER": AddBoundsKind_LOWER,
			"MAJOR": AddBoundsKind_MAJOR,
			"MINOR": AddBoundsKind_MINOR,
			"EXACT": AddBoundsKind_EXACT,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.AnnotationStyle",
		reflect.TypeOf((*AnnotationStyle)(nil)).Elem(),
		map[string]interface{}{
			"LINE": AnnotationStyle_LINE,
			"SPLIT": AnnotationStyle_SPLIT,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.AuthPolicy",
		reflect.TypeOf((*AuthPolicy)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": AuthPolicy_AUTO,
			"ALWAYS": AuthPolicy_ALWAYS,
			"NEVER": AuthPolicy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.BuildBackendSettings",
		reflect.TypeOf((*BuildBackendSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.DependencyGroupSettings",
		reflect.TypeOf((*DependencyGroupSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.ForkStrategy",
		reflect.TypeOf((*ForkStrategy)(nil)).Elem(),
		map[string]interface{}{
			"FEWEST": ForkStrategy_FEWEST,
			"REQUIRES_HYPHEN_PYTHON": ForkStrategy_REQUIRES_HYPHEN_PYTHON,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.Index",
		reflect.TypeOf((*Index)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.IndexCacheControl",
		reflect.TypeOf((*IndexCacheControl)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.IndexFormat",
		reflect.TypeOf((*IndexFormat)(nil)).Elem(),
		map[string]interface{}{
			"SIMPLE": IndexFormat_SIMPLE,
			"FLAT": IndexFormat_FLAT,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.IndexStrategy",
		reflect.TypeOf((*IndexStrategy)(nil)).Elem(),
		map[string]interface{}{
			"FIRST_HYPHEN_INDEX": IndexStrategy_FIRST_HYPHEN_INDEX,
			"UNSAFE_HYPHEN_FIRST_HYPHEN_MATCH": IndexStrategy_UNSAFE_HYPHEN_FIRST_HYPHEN_MATCH,
			"UNSAFE_HYPHEN_BEST_HYPHEN_MATCH": IndexStrategy_UNSAFE_HYPHEN_BEST_HYPHEN_MATCH,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.KeyringProviderType",
		reflect.TypeOf((*KeyringProviderType)(nil)).Elem(),
		map[string]interface{}{
			"DISABLED": KeyringProviderType_DISABLED,
			"SUBPROCESS": KeyringProviderType_SUBPROCESS,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.LinkMode",
		reflect.TypeOf((*LinkMode)(nil)).Elem(),
		map[string]interface{}{
			"CLONE": LinkMode_CLONE,
			"COPY": LinkMode_COPY,
			"HARDLINK": LinkMode_HARDLINK,
			"SYMLINK": LinkMode_SYMLINK,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.PipGroupName",
		reflect.TypeOf((*PipGroupName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.PipOptions",
		reflect.TypeOf((*PipOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.PrereleaseMode",
		reflect.TypeOf((*PrereleaseMode)(nil)).Elem(),
		map[string]interface{}{
			"DISALLOW": PrereleaseMode_DISALLOW,
			"ALLOW": PrereleaseMode_ALLOW,
			"IF_HYPHEN_NECESSARY": PrereleaseMode_IF_HYPHEN_NECESSARY,
			"EXPLICIT": PrereleaseMode_EXPLICIT,
			"IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT": PrereleaseMode_IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.PythonDownloads",
		reflect.TypeOf((*PythonDownloads)(nil)).Elem(),
		map[string]interface{}{
			"AUTOMATIC": PythonDownloads_AUTOMATIC,
			"MANUAL": PythonDownloads_MANUAL,
			"NEVER": PythonDownloads_NEVER,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.PythonPreference",
		reflect.TypeOf((*PythonPreference)(nil)).Elem(),
		map[string]interface{}{
			"ONLY_HYPHEN_MANAGED": PythonPreference_ONLY_HYPHEN_MANAGED,
			"MANAGED": PythonPreference_MANAGED,
			"SYSTEM": PythonPreference_SYSTEM,
			"ONLY_HYPHEN_SYSTEM": PythonPreference_ONLY_HYPHEN_SYSTEM,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.ResolutionMode",
		reflect.TypeOf((*ResolutionMode)(nil)).Elem(),
		map[string]interface{}{
			"HIGHEST": ResolutionMode_HIGHEST,
			"LOWEST": ResolutionMode_LOWEST,
			"LOWEST_HYPHEN_DIRECT": ResolutionMode_LOWEST_HYPHEN_DIRECT,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.SchemaConflictItem",
		reflect.TypeOf((*SchemaConflictItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.StaticMetadata",
		reflect.TypeOf((*StaticMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.TargetTriple",
		reflect.TypeOf((*TargetTriple)(nil)).Elem(),
		map[string]interface{}{
			"WINDOWS": TargetTriple_WINDOWS,
			"LINUX": TargetTriple_LINUX,
			"MACOS": TargetTriple_MACOS,
			"X86_UNDERSCORE_64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC": TargetTriple_X86_UNDERSCORE_64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC,
			"AARCH64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC": TargetTriple_AARCH64_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC,
			"I686_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC": TargetTriple_I686_HYPHEN_PC_HYPHEN_WINDOWS_HYPHEN_MSVC,
			"X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU": TargetTriple_X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU,
			"AARCH64_HYPHEN_APPLE_HYPHEN_DARWIN": TargetTriple_AARCH64_HYPHEN_APPLE_HYPHEN_DARWIN,
			"X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_DARWIN": TargetTriple_X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_DARWIN,
			"AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU": TargetTriple_AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_GNU,
			"AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL": TargetTriple_AARCH64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL,
			"X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL": TargetTriple_X86_UNDERSCORE_64_HYPHEN_UNKNOWN_HYPHEN_LINUX_HYPHEN_MUSL,
			"RISCV64_HYPHEN_UNKNOWN_HYPHEN_LINUX": TargetTriple_RISCV64_HYPHEN_UNKNOWN_HYPHEN_LINUX,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX2014": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX2014,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39,
			"X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40": TargetTriple_X86_UNDERSCORE_64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40,
			"AARCH64_HYPHEN_MANYLINUX2014": TargetTriple_AARCH64_HYPHEN_MANYLINUX2014,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_17,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_28,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_31,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_32,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_33,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_34,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_35,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_36,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_37,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_38,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_39,
			"AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40": TargetTriple_AARCH64_HYPHEN_MANYLINUX_UNDERSCORE_2_UNDERSCORE_40,
			"AARCH64_HYPHEN_LINUX_HYPHEN_ANDROID": TargetTriple_AARCH64_HYPHEN_LINUX_HYPHEN_ANDROID,
			"X86_UNDERSCORE_64_HYPHEN_LINUX_HYPHEN_ANDROID": TargetTriple_X86_UNDERSCORE_64_HYPHEN_LINUX_HYPHEN_ANDROID,
			"WASM32_HYPHEN_PYODIDE2024": TargetTriple_WASM32_HYPHEN_PYODIDE2024,
			"ARM64_HYPHEN_APPLE_HYPHEN_IOS": TargetTriple_ARM64_HYPHEN_APPLE_HYPHEN_IOS,
			"ARM64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR": TargetTriple_ARM64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR,
			"X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR": TargetTriple_X86_UNDERSCORE_64_HYPHEN_APPLE_HYPHEN_IOS_HYPHEN_SIMULATOR,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.ToolUvWorkspace",
		reflect.TypeOf((*ToolUvWorkspace)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.TorchMode",
		reflect.TypeOf((*TorchMode)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": TorchMode_AUTO,
			"CPU": TorchMode_CPU,
			"CU130": TorchMode_CU130,
			"CU129": TorchMode_CU129,
			"CU128": TorchMode_CU128,
			"CU126": TorchMode_CU126,
			"CU125": TorchMode_CU125,
			"CU124": TorchMode_CU124,
			"CU123": TorchMode_CU123,
			"CU122": TorchMode_CU122,
			"CU121": TorchMode_CU121,
			"CU120": TorchMode_CU120,
			"CU118": TorchMode_CU118,
			"CU117": TorchMode_CU117,
			"CU116": TorchMode_CU116,
			"CU115": TorchMode_CU115,
			"CU114": TorchMode_CU114,
			"CU113": TorchMode_CU113,
			"CU112": TorchMode_CU112,
			"CU111": TorchMode_CU111,
			"CU110": TorchMode_CU110,
			"CU102": TorchMode_CU102,
			"CU101": TorchMode_CU101,
			"CU100": TorchMode_CU100,
			"CU92": TorchMode_CU92,
			"CU91": TorchMode_CU91,
			"CU90": TorchMode_CU90,
			"CU80": TorchMode_CU80,
			"ROCM6_3": TorchMode_ROCM6_3,
			"ROCM6_2_4": TorchMode_ROCM6_2_4,
			"ROCM6_2": TorchMode_ROCM6_2,
			"ROCM6_1": TorchMode_ROCM6_1,
			"ROCM6_0": TorchMode_ROCM6_0,
			"ROCM5_7": TorchMode_ROCM5_7,
			"ROCM5_6": TorchMode_ROCM5_6,
			"ROCM5_5": TorchMode_ROCM5_5,
			"ROCM5_4_2": TorchMode_ROCM5_4_2,
			"ROCM5_4": TorchMode_ROCM5_4,
			"ROCM5_3": TorchMode_ROCM5_3,
			"ROCM5_2": TorchMode_ROCM5_2,
			"ROCM5_1_1": TorchMode_ROCM5_1_1,
			"ROCM4_2": TorchMode_ROCM4_2,
			"ROCM4_1": TorchMode_ROCM4_1,
			"ROCM4_0_1": TorchMode_ROCM4_0_1,
			"XPU": TorchMode_XPU,
		},
	)
	_jsii_.RegisterEnum(
		"projen.python.uvConfig.TrustedPublishing",
		reflect.TypeOf((*TrustedPublishing)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": TrustedPublishing_ALWAYS,
			"NEVER": TrustedPublishing_NEVER,
			"AUTOMATIC": TrustedPublishing_AUTOMATIC,
		},
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.UvConfiguration",
		reflect.TypeOf((*UvConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"projen.python.uvConfig.WheelDataIncludes",
		reflect.TypeOf((*WheelDataIncludes)(nil)).Elem(),
	)
}
