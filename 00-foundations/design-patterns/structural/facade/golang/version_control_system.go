package facade

import "fmt"

// VersionControlSystem pulls source code for deployment
type VersionControlSystem struct{}

// NewVersionControlSystem creates a new VersionControlSystem instance
func NewVersionControlSystem() *VersionControlSystem {
	return &VersionControlSystem{}
}

// PullLatestChanges pulls the latest changes from the given branch
func (v *VersionControlSystem) PullLatestChanges(branch string) {
	fmt.Printf("VCS: Pulling latest changes from '%s'...\n", branch)
	simulateDelay()
	fmt.Println("VCS: Pull complete.")
}
