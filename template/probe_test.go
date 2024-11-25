package template

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateProbeCRDSpec(t *testing.T) {
	expectedOutput := `apiVersion: apiextensions.k8s.io/v1
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

	// Generate the CRD spec
	output, err := GenerateProbeCRDSpec()

	// Assert that there was no error and the output matches the expected output
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
