package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for Maven releases.
// Experimental.
type MavenPublishOptions struct {
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Experimental.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if needed.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Experimental.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Experimental.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// URL of Nexus repository.
	//
	// if not set, defaults to https://oss.sonatype.org
	// Default: - "https://oss.sonatype.org" or none when publishing to Maven Central
	//
	// Experimental.
	MavenEndpoint *string `field:"optional" json:"mavenEndpoint" yaml:"mavenEndpoint"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven packages. See instructions.
	// See: https://github.com/aws/publib#maven
	//
	// Default: "MAVEN_GPG_PRIVATE_KEY_PASSPHRASE" or not set when using GitHub Packages.
	//
	// Experimental.
	MavenGpgPrivateKeyPassphrase *string `field:"optional" json:"mavenGpgPrivateKeyPassphrase" yaml:"mavenGpgPrivateKeyPassphrase"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven
	// packages. See instructions.
	// See: https://github.com/aws/publib#maven
	//
	// Default: "MAVEN_GPG_PRIVATE_KEY" or not set when using GitHub Packages.
	//
	// Experimental.
	MavenGpgPrivateKeySecret *string `field:"optional" json:"mavenGpgPrivateKeySecret" yaml:"mavenGpgPrivateKeySecret"`
	// GitHub secret name which contains the Password for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Default: "MAVEN_PASSWORD" or "GITHUB_TOKEN" when using GitHub Packages.
	//
	// Experimental.
	MavenPassword *string `field:"optional" json:"mavenPassword" yaml:"mavenPassword"`
	// Deployment repository when not deploying to Maven Central.
	// Default: - not set.
	//
	// Experimental.
	MavenRepositoryUrl *string `field:"optional" json:"mavenRepositoryUrl" yaml:"mavenRepositoryUrl"`
	// Used in maven settings for credential lookup (e.g. use github when publishing to GitHub).
	//
	// Set to `central-ossrh` to publish to Maven Central.
	// Default: "ossrh" (Maven Central) or "github" when using GitHub Packages.
	//
	// Experimental.
	MavenServerId *string `field:"optional" json:"mavenServerId" yaml:"mavenServerId"`
	// GitHub secret name which contains the Maven Central (sonatype) staging profile ID (e.g. 68a05363083174). Staging profile ID can be found in the URL of the "Releases" staging profile under "Staging Profiles" in https://oss.sonatype.org (e.g. https://oss.sonatype.org/#stagingProfiles;11a33451234521).
	// Default: "MAVEN_STAGING_PROFILE_ID" or not set when using GitHub Packages.
	//
	// Experimental.
	MavenStagingProfileId *string `field:"optional" json:"mavenStagingProfileId" yaml:"mavenStagingProfileId"`
	// GitHub secret name which contains the Username for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Default: "MAVEN_USERNAME" or the GitHub Actor when using GitHub Packages.
	//
	// Experimental.
	MavenUsername *string `field:"optional" json:"mavenUsername" yaml:"mavenUsername"`
}

