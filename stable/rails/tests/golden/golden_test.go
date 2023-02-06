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
	        GoldenFileName: "defaults/full.golden.yaml",
	        ValuesFiles: []string{"defaults/values.golden.yaml"},
	    },
	    {
	        GoldenFileName: "regular-app/full.golden.yaml",
	        ValuesFiles: []string{"regular-app/values.golden.yaml"},
	    },
	    {
	        GoldenFileName: "components/hpa.golden.yaml",
	        ValuesFiles: []string{"components/hpa.values.yaml"},
	        Templates: []string{"templates/hpa-worker.yaml", "templates/hpa-webserver.yaml", "templates/datadog-metric.yaml"},
	    },
	    {
	        GoldenFileName: "components/ingress.golden.yaml",
	        ValuesFiles: []string{"components/ingress.values.yaml"},
	        Templates: []string{"templates/ingress.yaml"},
	    },
	    {
	        GoldenFileName: "components/migrations.golden.yaml",
	        ValuesFiles: []string{"components/migrations.values.yaml"},
	        Templates: []string{"templates/migrations.yaml"},
	    },
	    {
	        GoldenFileName: "components/webserver.golden.yaml",
	        ValuesFiles: []string{"components/webserver.values.yaml"},
	        Templates: []string{"templates/deployment-webserver.yaml", "templates/service-webserver.yaml"},
	    },
	    {
	        GoldenFileName: "components/worker.golden.yaml",
	        ValuesFiles: []string{"components/worker.values.yaml"},
	        Templates: []string{"templates/deployment-worker.yaml"},
	    },
	    {
	        GoldenFileName: "components/websocket.golden.yaml",
	        ValuesFiles: []string{"components/websocket.values.yaml"},
	        Templates: []string{"templates/deployment-websocket.yaml", "templates/service-websocket.yaml"},
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