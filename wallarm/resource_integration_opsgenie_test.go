package wallarm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIntegrationOpsGenieRequiredFields(t *testing.T) {
	rnd := generateRandomResourceName(5)
	name := "wallarm_integration_opsgenie." + rnd
	rndToken := generateRandomUUID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testWallarmIntegrationOpsGenieRequiredOnly(rnd, rndToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "api_token", rndToken),
				),
			},
		},
	})
}

func TestAccIntegrationOpsGenieFullSettings(t *testing.T) {
	rnd := generateRandomResourceName(5)
	name := "wallarm_integration_opsgenie." + rnd
	rndToken := generateRandomUUID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testWallarmIntegrationOpsGenieFullConfig(rnd, "tf-test-"+rnd, rndToken, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "tf-test-"+rnd),
					resource.TestCheckResourceAttr(name, "api_token", rndToken),
					resource.TestCheckResourceAttr(name, "active", "true"),
					resource.TestCheckResourceAttr(name, "event.#", "2"),
				),
			},
		},
	})
}

func TestAccIntegrationOpsGenieCreateThenUpdate(t *testing.T) {
	rnd := generateRandomResourceName(5)
	name := "wallarm_integration_opsgenie." + rnd
	rndToken := generateRandomUUID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testWallarmIntegrationOpsGenieFullConfig(rnd, "tf-test-"+rnd, rndToken, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "tf-test-"+rnd),
					resource.TestCheckResourceAttr(name, "api_token", rndToken),
					resource.TestCheckResourceAttr(name, "active", "true"),
					resource.TestCheckResourceAttr(name, "event.#", "2"),
				),
			},
			{
				Config: testWallarmIntegrationOpsGenieFullConfig(rnd, "tf-updated-"+rnd, rndToken, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "tf-updated-"+rnd),
					resource.TestCheckResourceAttr(name, "api_token", rndToken),
					resource.TestCheckResourceAttr(name, "active", "false"),
					resource.TestCheckResourceAttr(name, "event.#", "2"),
				),
			},
		},
	})
}

func testWallarmIntegrationOpsGenieRequiredOnly(resourceID, token string) string {
	return fmt.Sprintf(`
resource "wallarm_integration_opsgenie" "%[1]s" {
	api_token = "%[2]s"
}`, resourceID, token)
}

func testWallarmIntegrationOpsGenieFullConfig(resourceID, name, token, active string) string {
	return fmt.Sprintf(`
resource "wallarm_integration_opsgenie" "%[1]s" {
	name = "%[2]s"
	api_token = "%[3]s"
	active = "%[4]s"
	
	event {
		event_type = "hit"
		active = true
	}
	event {
		event_type = "vuln"
		active = "%[4]s"
	}
}`, resourceID, name, token, active)
}
