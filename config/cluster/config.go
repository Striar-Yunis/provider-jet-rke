package cluster

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure maps the name of the rke_cluster to the cluster_name terraform field
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rke_cluster", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
			base["cluster_name"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"cluster_name",
		}
		if s, ok := r.TerraformResource.Schema["certificates"]; ok {
			s.Sensitive = false
		}
		// r.Sensitive.AdditionalConnectionDetailsFn
	})
}
