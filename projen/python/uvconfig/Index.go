package uvconfig


// Experimental.
type Index struct {
	// The URL of the index.
	//
	// Expects to receive a URL (e.g., `https://pypi.org/simple`) or a local path.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// When uv should use authentication for requests to the index.
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "my-index"
	// url = "https://<omitted>/simple"
	// authenticate = "always"
	// ```.
	// Experimental.
	Authenticate *string `field:"optional" json:"authenticate" yaml:"authenticate"`
	// Cache control configuration for this index.
	//
	// When set, these headers will override the server's cache control headers
	// for both package metadata requests and artifact downloads.
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "my-index"
	// url = "https://<omitted>/simple"
	// cache-control = { api = "max-age=600", files = "max-age=3600" }
	// ```.
	// Experimental.
	CacheControl *IndexCacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Mark the index as the default index.
	//
	// By default, uv uses PyPI as the default index, such that even if additional indexes are
	// defined via `[[tool.uv.index]]`, PyPI will still be used as a fallback for packages that
	// aren't found elsewhere. To disable the PyPI default, set `default = true` on at least one
	// other index.
	//
	// Marking an index as default will move it to the front of the list of indexes, such that it
	// is given the highest priority when resolving packages.
	// Experimental.
	Default *bool `field:"optional" json:"default" yaml:"default"`
	// Mark the index as explicit.
	//
	// Explicit indexes will _only_ be used when explicitly requested via a `[tool.uv.sources]`
	// definition, as in:
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "pytorch"
	// url = "https://download.pytorch.org/whl/cu121"
	// explicit = true
	//
	// [tool.uv.sources]
	// torch = { index = "pytorch" }
	// ```.
	// Experimental.
	Explicit *bool `field:"optional" json:"explicit" yaml:"explicit"`
	// The format used by the index.
	//
	// Indexes can either be PEP 503-compliant (i.e., a PyPI-style registry implementing the Simple
	// API) or structured as a flat list of distributions (e.g., `--find-links`). In both cases,
	// indexes can point to either local or remote resources.
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Status codes that uv should ignore when deciding whether to continue searching in the next index after a failure.
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "my-index"
	// url = "https://<omitted>/simple"
	// ignore-error-codes = [401, 403]
	// ```.
	// Experimental.
	IgnoreErrorCodes *[]*float64 `field:"optional" json:"ignoreErrorCodes" yaml:"ignoreErrorCodes"`
	// The name of the index.
	//
	// Index names can be used to reference indexes elsewhere in the configuration. For example,
	// you can pin a package to a specific index by name:
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "pytorch"
	// url = "https://download.pytorch.org/whl/cu121"
	//
	// [tool.uv.sources]
	// torch = { index = "pytorch" }
	// ```.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL of the upload endpoint.
	//
	// When using `uv publish --index <name>`, this URL is used for publishing.
	//
	// A configuration for the default index PyPI would look as follows:
	//
	// ```toml
	// [[tool.uv.index]]
	// name = "pypi"
	// url = "https://pypi.org/simple"
	// publish-url = "https://upload.pypi.org/legacy/"
	// ```.
	// Experimental.
	PublishUrl *string `field:"optional" json:"publishUrl" yaml:"publishUrl"`
}

