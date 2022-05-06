package cluster

import (
	"encoding/json"
	"fmt"

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
			"certificates",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			certs := attr["certificates"]
			for _, cert := range certs.([]interface{}) {
				testcert := cert.(map[string]interface{})
				certname := "certificates" + "." + testcert["name"].(string)
				if a, err := json.Marshal((testcert)); err != nil {
					fmt.Println("Error converting certificate to JSON", err)
				} else {
					conn[certname] = a
				}

			}
			return conn, nil
		}
	})
}
