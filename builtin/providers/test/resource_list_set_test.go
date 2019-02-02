package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestResourceListSet_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: strings.TrimSpace(`
resource "test_resource_list_set" "foo" {
  list {
    set {
      elem = "A"
    }
    set {
      elem = "B"
    }
  }
}
				`),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						fmt.Println("STATE:", s)
						return nil
					},
				),
			},
			resource.TestStep{
				Config: strings.TrimSpace(`
resource "test_resource_list_set" "foo" {
  list {
    set {
      elem = "B"
    }
    set {
      elem = "C"
    }
  }
}
				`),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						fmt.Println("STATE:", s)
						return nil
					},
				),
			},
		},
	})
}
