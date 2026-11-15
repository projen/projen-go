package uvconfig


// A hash algorithm that can be required for distributions resolved from an index.
// Experimental.
type IndexHashAlgorithm string

const (
	// md5.
	// Experimental.
	IndexHashAlgorithm_MD5 IndexHashAlgorithm = "MD5"
	// sha256.
	// Experimental.
	IndexHashAlgorithm_SHA256 IndexHashAlgorithm = "SHA256"
	// sha384.
	// Experimental.
	IndexHashAlgorithm_SHA384 IndexHashAlgorithm = "SHA384"
	// sha512.
	// Experimental.
	IndexHashAlgorithm_SHA512 IndexHashAlgorithm = "SHA512"
	// blake2b.
	// Experimental.
	IndexHashAlgorithm_BLAKE2B IndexHashAlgorithm = "BLAKE2B"
)

