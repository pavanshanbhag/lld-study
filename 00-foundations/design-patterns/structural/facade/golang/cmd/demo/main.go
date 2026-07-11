package main

import (
	"fmt"

	"facade"
)

func main() {
	deploymentFacade := facade.NewDeploymentFacade()

	fmt.Println("=== Performing Full Deployment ===")
	success := deploymentFacade.DeployApplication("main", "production-server")
	if !success {
		fmt.Println("Full deployment failed!")
	}

	fmt.Println("\n=== Performing Hotfix Deployment ===")
	success = deploymentFacade.DeployHotfix("hotfix-123", "production-server")
	if !success {
		fmt.Println("Hotfix deployment failed!")
	}
}
