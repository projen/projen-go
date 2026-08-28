package polaris


// Specifies how new defects should be handled.
// Experimental.
type CommitConfigurationConnectTriage struct {
	// User to whom any new defects will be assigned.
	//
	// The specified user must already exist in the Coverity Connect database. The default is the current user.
	// Experimental.
	NewDefectOwner *string `field:"optional" json:"newDefectOwner" yaml:"newDefectOwner"`
	// Limit on the number of defects to assign to the specified user.
	//
	// If the number of discovered defects is more than the limit, then no assignment is done.
	// Experimental.
	NewDefectOwnerLimit *float64 `field:"optional" json:"newDefectOwnerLimit" yaml:"newDefectOwnerLimit"`
	// If true, the owner for newly detected defects that exist locally is set to the specified user.
	// Experimental.
	SetNewDefectOwner *bool `field:"optional" json:"setNewDefectOwner" yaml:"setNewDefectOwner"`
}

