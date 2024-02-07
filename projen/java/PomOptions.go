package java


// Options for `Pom`.
// Experimental.
type PomOptions struct {
	// The artifactId is generally the name that the project is known by.
	//
	// Although
	// the groupId is important, people within the group will rarely mention the
	// groupId in discussion (they are often all be the same ID, such as the
	// MojoHaus project groupId: org.codehaus.mojo). It, along with the groupId,
	// creates a key that separates this project from every other project in the
	// world (at least, it should :) ). Along with the groupId, the artifactId
	// fully defines the artifact's living quarters within the repository. In the
	// case of the above project, my-project lives in
	// $M2_REPO/org/codehaus/mojo/my-project.
	// Default: "my-app".
	//
	// Experimental.
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// This is generally unique amongst an organization or a project.
	//
	// For example,
	// all core Maven artifacts do (well, should) live under the groupId
	// org.apache.maven. Group ID's do not necessarily use the dot notation, for
	// example, the junit project. Note that the dot-notated groupId does not have
	// to correspond to the package structure that the project contains. It is,
	// however, a good practice to follow. When stored within a repository, the
	// group acts much like the Java packaging structure does in an operating
	// system. The dots are replaced by OS specific directory separators (such as
	// '/' in Unix) which becomes a relative directory structure from the base
	// repository. In the example given, the org.codehaus.mojo group lives within
	// the directory $M2_REPO/org/codehaus/mojo.
	// Default: "org.acme"
	//
	// Experimental.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// This is the last piece of the naming puzzle.
	//
	// groupId:artifactId denotes a
	// single project but they cannot delineate which incarnation of that project
	// we are talking about. Do we want the junit:junit of 2018 (version 4.12), or
	// of 2007 (version 3.8.2)? In short: code changes, those changes should be
	// versioned, and this element keeps those versions in line. It is also used
	// within an artifact's repository to separate versions from each other.
	// my-project version 1.0 files live in the directory structure
	// $M2_REPO/org/codehaus/mojo/my-project/1.0.
	// Default: "0.1.0"
	//
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Default: undefined.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Project packaging format.
	// Default: "jar".
	//
	// Experimental.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// A Parent Pom can be used to have a child project inherit properties/plugins/ect in order to reduce duplication and keep standards across a large amount of repos.
	// Default: undefined.
	//
	// Experimental.
	ParentPom *ParentPom `field:"optional" json:"parentPom" yaml:"parentPom"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Default: undefined.
	//
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

