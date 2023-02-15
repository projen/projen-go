package circleci


// The branches key controls whether the current branch should have a schedule trigger created for it, where current branch is the branch containing the config.yml file with the trigger stanza. That is, a push on the main branch will only schedule a workflow for the main branch.
//
// Branches can have the keys only and ignore which either map to a single string naming a branch.
// You may also use regular expressions to match against branches by enclosing them with /’s, or map to a list of such strings.
// Regular expressions must match the entire string.
//
// Any branches that match only will run the job.
// Any branches that match ignore will not run the job.
// If neither only nor ignore are specified then all branches will run the job.
// If both only and ignore are specified the only is considered before ignore.
// See: https://circleci.com/docs/2.0/configuration-reference/#filters
//
// Experimental.
type Filter struct {
	// Experimental.
	Branches *FilterConfig `field:"optional" json:"branches" yaml:"branches"`
	// Experimental.
	Tags *FilterConfig `field:"optional" json:"tags" yaml:"tags"`
}

