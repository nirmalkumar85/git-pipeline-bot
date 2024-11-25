package template

import (
	"bytes"
	"text/template"
	"fmt"
)

// GenerateProbeCRDSpec generates the CRD spec content using a Go template
func GenerateProbeCRDSpec() (string, error) {
	// Define a simple template for the Probe CRD spec
	tmpl := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: probe.example.com
spec:
  group: example.com
  names:
    kind: Probe
    plural: probes
    singular: probe
  scope: Namespaced
  versions:
    - name: v1
      served: true
      storage: true
`

	// Parse and execute the template
	t, err := template.New("probe").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %v", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, nil); err != nil {
		return "", fmt.Errorf("error executing template: %v", err)
	}

	return buf.String(), nil
}
