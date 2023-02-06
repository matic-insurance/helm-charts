package golden

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GoldenTestCase struct {
    GoldenFileName string
    Templates []string
    ValuesFiles []string
}

func TestGoldenTemplates(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	testCases := []GoldenTestCase {
	    {
	        GoldenFileName: "defaults/deployment-webserver.golden.yaml",
	        Templates: []string{"templates/deployment-webserver.yaml"},
	    },
	    {
	        GoldenFileName: "defaults/full.golden.yaml",
	        ValuesFiles: []string{"defaults/values.golden.yaml"},
	    },
	    {
	        GoldenFileName: "regular-app/full.golden.yaml",
	        ValuesFiles: []string{"regular-app/values.golden.yaml"},
	    },
	    {
	        GoldenFileName: "templates/hpa.golden.yaml",
	        Templates: []string{"templates/hpa-worker.yaml", "templates/hpa-webserver.yaml", "templates/datadog-metric.yaml"},
	        ValuesFiles: []string{"templates/hpa.values.yaml"},
	    },
	}

	for _, testCase := range testCases {
		suite.Run(t, &GoldenTestSuite{
			ChartPath:      chartPath,
			Release:        "rails-test",
			GoldenFileName: testCase.GoldenFileName,
			Templates:      testCase.Templates,
			ValuesFiles:    testCase.ValuesFiles,
		})
	}
}