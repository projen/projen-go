package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `MavenPublishOptions` instead.
type JsiiReleaseMaven struct {
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if neede.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `MavenPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Deprecated: Use `MavenPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// URL of Nexus repository.
	//
	// if not set, defaults to https://oss.sonatype.org
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenEndpoint *string `field:"optional" json:"mavenEndpoint" yaml:"mavenEndpoint"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven packages. See instructions.
	// See: https://github.com/aws/publib#maven
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenGpgPrivateKeyPassphrase *string `field:"optional" json:"mavenGpgPrivateKeyPassphrase" yaml:"mavenGpgPrivateKeyPassphrase"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven
	// packages. See instructions.
	// See: https://github.com/aws/publib#maven
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenGpgPrivateKeySecret *string `field:"optional" json:"mavenGpgPrivateKeySecret" yaml:"mavenGpgPrivateKeySecret"`
	// GitHub secret name which contains the Password for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenPassword *string `field:"optional" json:"mavenPassword" yaml:"mavenPassword"`
	// Deployment repository when not deploying to Maven Central.
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenRepositoryUrl *string `field:"optional" json:"mavenRepositoryUrl" yaml:"mavenRepositoryUrl"`
	// Used in maven settings for credential lookup (e.g. use github when publishing to GitHub).
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenServerId *string `field:"optional" json:"mavenServerId" yaml:"mavenServerId"`
	// GitHub secret name which contains the Maven Central (sonatype) staging profile ID (e.g. 68a05363083174). Staging profile ID can be found in the URL of the "Releases" staging profile under "Staging Profiles" in https://oss.sonatype.org (e.g. https://oss.sonatype.org/#stagingProfiles;11a33451234521).
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenStagingProfileId *string `field:"optional" json:"mavenStagingProfileId" yaml:"mavenStagingProfileId"`
	// GitHub secret name which contains the Username for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenUsername *string `field:"optional" json:"mavenUsername" yaml:"mavenUsername"`
}

