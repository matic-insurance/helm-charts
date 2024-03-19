package golden

import (
	"github.com/matic-insurance/helm-charts/ops/golden_testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"testing"
)

func TestMeshGoldenTemplates(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	testCases := []golden_testing.Suite{
		{
			GoldenFileName: "defaults/defaults.golden.yaml",
			ValuesFiles:    []string{"defaults/defaults.values.yaml"},
		},
		{
			GoldenFileName: "components/destination-rule.golden.yaml",
			ValuesFiles:    []string{"components/destination-rule.values.yaml"},
			Templates:      []string{"templates/destination-rule.yaml"},
		},
		{
			GoldenFileName: "components/virtual-service.golden.yaml",
			ValuesFiles:    []string{"components/virtual-service.values.yaml"},
			Templates:      []string{"templates/virtual-service.mesh.yaml"},
		},
		{
			GoldenFileName: "components/ingress-allow-locations.golden.yaml",
			ValuesFiles:    []string{"components/ingress-allow-locations.values.yaml"},
			Templates:      []string{"templates/virtual-service.mesh.yaml", "templates/virtual-service.gateway.yaml"},
		},
		{
			GoldenFileName: "components/ingress-deny-locations.golden.yaml",
			ValuesFiles:    []string{"components/ingress-deny-locations.values.yaml"},
			Templates:      []string{"templates/virtual-service.mesh.yaml", "templates/virtual-service.gateway.yaml"},
		},
		{
			GoldenFileName: "components/ingress-deny-and-allow.golden.yaml",
			ValuesFiles:    []string{"components/ingress-deny-and-allow.values.yaml"},
			Templates:      []string{"templates/virtual-service.mesh.yaml", "templates/virtual-service.gateway.yaml"},
		},
		{
			GoldenFileName: "components/service-entry.golden.yaml",
			ValuesFiles:    []string{"components/service-entry.values.yaml"},
			Templates:      []string{"templates/service-entry.yaml"},
		},
		{
			GoldenFileName: "components/envoy-filter-max-body-size.golden.yaml",
			ValuesFiles:    []string{"components/envoy-filter-max-body-size.values.yaml"},
			Templates:      []string{"templates/envoy-filter.yaml"},
		},
		{
			GoldenFileName: "components/external-gateway.golden.yaml",
			ValuesFiles:    []string{"components/external-gateway.values.yaml"},
			Templates:      []string{"templates/certificate.yaml", "templates/gateway.yaml"},
		},
		{
			GoldenFileName: "features/no-trace-proxy.golden.yaml",
			ValuesFiles:    []string{"features/no-trace-proxy.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/disabled.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/disabled.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/single-webserver.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/single-webserver.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/webserver-and-websockets.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/webserver-and-websockets.values.yaml"},
		},
	}

	for _, testCase := range testCases {
		suite.Run(t, &golden_testing.Suite{
			ChartPath:      chartPath,
			Release:        "app-component-test",
			GoldenFileName: testCase.GoldenFileName,
			Templates:      testCase.Templates,
			ValuesFiles:    testCase.ValuesFiles,
		})
	}
}
