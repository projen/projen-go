package uvconfig


// Experimental.
type IndexStrategy string

const (
	// Only use results from the first index that returns a match for a given package name.
	//
	// While this differs from pip's behavior, it's the default index strategy as it's the most
	// secure. (first-index)
	// Experimental.
	IndexStrategy_FIRST_HYPHEN_INDEX IndexStrategy = "FIRST_HYPHEN_INDEX"
	// Search for every package name across all indexes, exhausting the versions from the first index before moving on to the next.
	//
	// In this strategy, we look for every package across all indexes. When resolving, we attempt
	// to use versions from the indexes in order, such that we exhaust all available versions from
	// the first index before moving on to the next. Further, if a version is found to be
	// incompatible in the first index, we do not reconsider that version in subsequent indexes,
	// even if the secondary index might contain compatible versions (e.g., variants of the same
	// versions with different ABI tags or Python version constraints).
	//
	// See: <https://peps.python.org/pep-0708/> (unsafe-first-match)
	// Experimental.
	IndexStrategy_UNSAFE_HYPHEN_FIRST_HYPHEN_MATCH IndexStrategy = "UNSAFE_HYPHEN_FIRST_HYPHEN_MATCH"
	// Search for every package name across all indexes, preferring the "best" version found.
	//
	// If a
	// package version is in multiple indexes, only look at the entry for the first index.
	//
	// In this strategy, we look for every package across all indexes. When resolving, we consider
	// all versions from all indexes, choosing the "best" version found (typically, the highest
	// compatible version).
	//
	// This most closely matches pip's behavior, but exposes the resolver to "dependency confusion"
	// attacks whereby malicious actors can publish packages to public indexes with the same name
	// as internal packages, causing the resolver to install the malicious package in lieu of
	// the intended internal package.
	//
	// See: <https://peps.python.org/pep-0708/> (unsafe-best-match)
	// Experimental.
	IndexStrategy_UNSAFE_HYPHEN_BEST_HYPHEN_MATCH IndexStrategy = "UNSAFE_HYPHEN_BEST_HYPHEN_MATCH"
)

