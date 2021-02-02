package jsii

import (
	rt "github.com/aws-cdk/jsii/jsii-experimental"
	"sync"
)

var once sync.Once

// Initialize performs the necessary work for the enclosing
// module to be loaded in the jsii kernel.
func Initialize() {
	once.Do(func(){
		// Load this library into the kernel
		rt.Load("projen", "0.15.15", tarball)
	})
}
