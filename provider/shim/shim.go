package shim

import (
	"github.com/SAP/terraform-provider-btp/internal/provider"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

// fix provider.Provider here to match whats in internal/provider
// create shim example https://github.com/pulumi/pulumi-random/blob/master/provider/shim/shim.go#L8
func Provider() tfpf.Provider {
	return provider.New()
}
