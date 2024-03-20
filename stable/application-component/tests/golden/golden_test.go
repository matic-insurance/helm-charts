package golden

import (
	"github.com/matic-insurance/helm-charts/ops/golden_testing"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenTemplates(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	testCases := []golden_testing.Suite{
		{
			GoldenFileName: "defaults/defaults.golden.yaml",
			ValuesFiles:    []string{"defaults/defaults.values.yaml"},
		},
		{
			GoldenFileName: "components/monitoring.golden.yaml",
			ValuesFiles:    []string{"components/monitoring.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/serviceaccount.golden.yaml",
			ValuesFiles:    []string{"components/serviceaccount.values.yaml"},
			Templates:      []string{"templates/serviceaccount.yaml", "templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/image-global.golden.yaml",
			ValuesFiles:    []string{"components/image-global.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/image-custom.golden.yaml",
			ValuesFiles:    []string{"components/image-custom.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/deployment-scheduling.golden.yaml",
			ValuesFiles:    []string{"components/deployment-scheduling.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/deployment-runtime.golden.yaml",
			ValuesFiles:    []string{"components/deployment-runtime.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/deployment-advanced.golden.yaml",
			ValuesFiles:    []string{"components/deployment-advanced.values.yaml"},
			Templates:      []string{"templates/deployment.yaml"},
		},
		{
			GoldenFileName: "components/service.golden.yaml",
			ValuesFiles:    []string{"components/service.values.yaml"},
			Templates:      []string{"templates/deployment.yaml", "templates/service.yaml"},
		},
		{
			GoldenFileName: "components/ingress.golden.yaml",
			ValuesFiles:    []string{"components/ingress.values.yaml"},
			Templates:      []string{"templates/ingress.yaml", "templates/service.yaml"},
		},
		{
			GoldenFileName: "components/hpa.golden.yaml",
			ValuesFiles:    []string{"components/hpa.values.yaml"},
			Templates:      []string{"templates/hpa.yaml", "templates/datadog-metric.yaml", "templates/deployment.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/critical-mixed.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/critical-mixed.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/critical-on_demand.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/critical-on_demand.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/critical-spot.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/critical-spot.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/high-mixed.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/high-mixed.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/high-on_demand.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/high-on_demand.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/high-spot.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/high-spot.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/normal-mixed.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/normal-mixed.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/normal-on_demand.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/normal-mixed.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/normal-spot.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/normal-spot.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/irrelevant-mixed.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/irrelevant-mixed.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/irrelevant-on_demand.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/irrelevant-on_demand.values.yaml"},
		},
		{
			GoldenFileName: "deployment-schemes/irrelevant-spot.golden.yaml",
			ValuesFiles:    []string{"deployment-schemes/irrelevant-spot.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/regular-webserver.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/regular-webserver.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/no-injection.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/no-injection.values.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/no-mesh.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/no-mesh.values.yaml"},
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
