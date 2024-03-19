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
			GoldenFileName: "defaults/full.golden.yaml",
			ValuesFiles:    []string{"defaults/full.values.yaml"},
		},
		{
			GoldenFileName: "regular-app/full.golden.yaml",
			ValuesFiles:    []string{"regular-app/full.values.yaml"},
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
