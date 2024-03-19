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
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "components/serviceaccount.golden.yaml",
			ValuesFiles:    []string{"components/serviceaccount.values.yaml"},
			Templates:      []string{"templates/serviceaccount.yaml", "templates/job.yaml"},
		},
		{
			GoldenFileName: "components/image-custom.golden.yaml",
			ValuesFiles:    []string{"components/image-custom.values.yaml"},
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "components/image-global.golden.yaml",
			ValuesFiles:    []string{"components/image-global.values.yaml"},
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "components/job-scheduling.golden.yaml",
			ValuesFiles:    []string{"components/job-scheduling.values.yaml"},
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "components/job-runtime.golden.yaml",
			ValuesFiles:    []string{"components/job-runtime.values.yaml"},
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "components/job-advanced.golden.yaml",
			ValuesFiles:    []string{"components/job-advanced.values.yaml"},
			Templates:      []string{"templates/job.yaml"},
		},
		{
			GoldenFileName: "standard-configurations/regular-migration.golden.yaml",
			ValuesFiles:    []string{"standard-configurations/regular-migration.values.yaml"},
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
