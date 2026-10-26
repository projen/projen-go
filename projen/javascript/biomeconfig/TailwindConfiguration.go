package biomeconfig


// Tells Biome where your code lists Tailwind CSS classes, so the Tailwind lint rules know which strings to check.
// Experimental.
type TailwindConfiguration struct {
	// Attributes whose values are Tailwind classes, such as `class` in `<div class="px-2 py-2">`.
	//
	// Defaults to `class` and `className`. In HTML, they are case-insensitive.
	// Default: class` and `className`. In HTML, they are case-insensitive.
	//
	// Experimental.
	Attributes *[]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Functions that combine Tailwind classes, such as `clsx` and `cn`.
	//
	// Biome checks every string passed to these functions, like both strings
	// in `cn("px-2", "py-2")`. When an object is passed, Biome checks its
	// keys: in `cn({ "px-2": isActive })`, it checks `"px-2"`.
	//
	// Properties of these functions are checked too, so listing `tw` also
	// covers `` tw.div`px-2` `` and `tw.div("px-2")`.
	//
	// Defaults to `clsx`, `tw`, `twMerge`, `twJoin`, `cn`, `cc`, `cnb`, and
	// `ctl`.
	// Default: clsx`, `tw`, `twMerge`, `twJoin`, `cn`, `cc`, `cnb`, and.
	//
	// Experimental.
	MergeFunctions *[]*string `field:"optional" json:"mergeFunctions" yaml:"mergeFunctions"`
	// The CSS file that holds your Tailwind CSS configuration, such as `"./src/app.css"`. A relative path starts from the folder of the Biome configuration file.
	//
	// Biome reads the `@theme`, `@utility`, and `@custom-variant` rules in
	// this file and in the files it imports with `@import`, so that the
	// Tailwind rules know your theme values, utilities, and variants.
	//
	// To read the file, Biome scans your project when a Tailwind rule is
	// enabled. Setting this option also turns on `css.parser.tailwindDirectives`,
	// unless you set that option yourself.
	// Experimental.
	Stylesheet *string `field:"optional" json:"stylesheet" yaml:"stylesheet"`
	// Functions that describe a component's styles with an object, such as `cva` and `tv`.
	//
	// Biome checks every string passed directly to these functions. In an
	// object passed to them, Biome checks the classes under these properties:
	//
	// - `base`, `class`, and `className`: classes that always apply.
	// - `variants`: classes for each version of the component, such as each
	// button size.
	// - `slots`: classes for each part of the component, such as its icon.
	// - `compoundVariants` and `compoundSlots`: only the `class` and
	// `className` of each entry.
	//
	// Other properties, such as `defaultVariants`, hold names rather than
	// classes, so Biome skips them. For example, Biome checks `"rounded"`,
	// `"px-2"`, and `"px-4"` here, but not `"small"`:
	//
	// ```js
	// cva("rounded", {
	// variants: {
	// size: { small: "px-2", large: "px-4" },
	// },
	// defaultVariants: { size: "small" },
	// });
	// ```
	//
	// Defaults to `cva` and `tv`.
	// Default: cva` and `tv`.
	//
	// Experimental.
	VariantFunctions *[]*string `field:"optional" json:"variantFunctions" yaml:"variantFunctions"`
}

