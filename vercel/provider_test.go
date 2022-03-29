package vercel_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/vercel/terraform-provider-vercel/client"
	"github.com/vercel/terraform-provider-vercel/vercel"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"vercel": func() (tfprotov6.ProviderServer, error) {
		return tfsdk.NewProtocol6Server(vercel.New()), nil
	},
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("VERCEL_API_TOKEN"); v == "" {
		t.Fatal("VERCEL_API_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("VERCEL_TERRAFORM_TESTING_TEAM"); v == "" {
		t.Fatal("VERCEL_TERRAFORM_TESTING_TEAM must be set for acceptance tests against a specific team")
	}
	if v := os.Getenv("VERCEL_TERRAFORM_TESTING_DOMAIN"); v == "" {
		t.Fatal("VERCEL_TERRAFORM_TESTING_DOMAIN must be set for acceptance tests for DNS records")
	}
}

var tc *client.Client

func testClient() *client.Client {
	if tc == nil {
		tc = client.New(os.Getenv("VERCEL_API_TOKEN"))
	}

	return tc
}

func testDomain() string {
	return os.Getenv("VERCEL_TERRAFORM_TESTING_DOMAIN")
}

func testTeam() string {
	return os.Getenv("VERCEL_TERRAFORM_TESTING_TEAM")
}
