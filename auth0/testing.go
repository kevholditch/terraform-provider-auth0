package auth0

import "github.com/hashicorp/terraform/terraform"

func getResourcesByType(resourceType string, state *terraform.State) []*terraform.ResourceState {

	var result []*terraform.ResourceState

	for _, rs := range state.RootModule().Resources {
		if rs.Type == resourceType {
			result = append(result, rs)
		}

	}

	return result
}
