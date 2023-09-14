package golden

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
	"regexp"
	"strings"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update-golden", false, "update golden test output files")

type GoldenTestSuite struct {
	suite.Suite
	ChartPath      string
	Release        string
	Namespace      string
	GoldenFileName string
	Templates      []string
	ValuesFiles    []string
}

func TestGoldenTemplates(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	testCases := []GoldenTestSuite {
	    {
	        GoldenFileName: "defaults/defaults.golden.yaml",
	        ValuesFiles: []string{"defaults/defaults.values.yaml"},
	    },
	    {
	        GoldenFileName: "components/monitoring.golden.yaml",
	        ValuesFiles: []string{"components/monitoring.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/serviceaccount.golden.yaml",
	        ValuesFiles: []string{"components/serviceaccount.values.yaml"},
	        Templates: []string{"templates/serviceaccount.yaml", "templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/image-global.golden.yaml",
	        ValuesFiles: []string{"components/image-global.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/image-custom.golden.yaml",
	        ValuesFiles: []string{"components/image-custom.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/deployment-scheduling.golden.yaml",
	        ValuesFiles: []string{"components/deployment-scheduling.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/deployment-runtime.golden.yaml",
	        ValuesFiles: []string{"components/deployment-runtime.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/deployment-advanced.golden.yaml",
	        ValuesFiles: []string{"components/deployment-advanced.values.yaml"},
	        Templates: []string{"templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "components/service.golden.yaml",
	        ValuesFiles: []string{"components/service.values.yaml"},
	        Templates: []string{"templates/deployment.yaml", "templates/service.yaml"},
	    },
	    {
	        GoldenFileName: "components/ingress.golden.yaml",
	        ValuesFiles: []string{"components/ingress.values.yaml"},
	        Templates: []string{"templates/ingress.yaml", "templates/service.yaml"},
	    },
	    {
	        GoldenFileName: "components/hpa.golden.yaml",
	        ValuesFiles: []string{"components/hpa.values.yaml"},
	        Templates: []string{"templates/hpa.yaml", "templates/datadog-metric.yaml", "templates/deployment.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/critical-mixed.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/critical-mixed.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/critical-on_demand.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/critical-on_demand.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/critical-spot.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/critical-spot.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/high-mixed.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/high-mixed.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/high-on_demand.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/high-on_demand.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/high-spot.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/high-spot.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/normal-mixed.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/normal-mixed.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/normal-on_demand.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/normal-mixed.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/normal-spot.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/normal-spot.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/irrelevant-mixed.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/irrelevant-mixed.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/irrelevant-on_demand.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/irrelevant-on_demand.values.yaml"},
	    },
	    {
	        GoldenFileName: "deployment-schemes/irrelevant-spot.golden.yaml",
	        ValuesFiles: []string{"deployment-schemes/irrelevant-spot.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/regular-webserver.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/regular-webserver.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/no-injection.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/no-injection.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/no-mesh.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/no-mesh.values.yaml"},
	    },
	}

	for _, testCase := range testCases {
		suite.Run(t, &GoldenTestSuite{
			ChartPath:      chartPath,
			Release:        "app-component-test",
			GoldenFileName: testCase.GoldenFileName,
			Templates:      testCase.Templates,
			ValuesFiles:    testCase.ValuesFiles,
		})
	}
}

func (s *GoldenTestSuite) TestTemplateMatchesGoldenFile() {
	actual := s.RenderTemplates()
	expected := s.ReadGoldenFile()

    if diff := cmp.Diff(strings.Split(string(expected), "\n"), strings.Split(actual, "\n")); diff != "" {
        s.T().Errorf("%s: mismatch (-want +got):\n%s", s.GoldenFilePath(), diff)
    }

// 	regex := regexp.MustCompile(`\s+helm.sh/chart:\s+.*`)
// 	bytes := regex.ReplaceAll([]byte(actual), []byte(""))
// 	actual = string(bytes)
}

func (s *GoldenTestSuite) RenderTemplates() string {
    namespace := s.Release + strings.ToLower(random.UniqueId())
    options := &helm.Options{
        KubectlOptions: k8s.NewKubectlOptions("test", "", namespace),
        ValuesFiles:    s.ValuesFiles,
    }
    template := helm.RenderTemplate(s.T(), options, s.ChartPath, s.Release, s.Templates)
    template = stripRandomData(template)

    if *update {
        err := ioutil.WriteFile(s.GoldenFilePath(), []byte(template), 0644)
        s.Require().NoError(err, "Golden file was not writable")
    }

    return template
}

func (s *GoldenTestSuite) GoldenFilePath() string {
	return s.GoldenFileName
}

func (s *GoldenTestSuite) ReadGoldenFile() string {
	expected, err := ioutil.ReadFile(s.GoldenFilePath())

	s.Require().NoError(err, "Golden file doesn't exist or was not readable")

    return string(expected)
}

func stripRandomData(template string) string {
    template = stripRegexp(template, `\s+helm.sh/chart:\s+.*`, "")
    template = stripRegexp(template, `rollme:\s+.*`, "rollme: \"123abc\"")
    template = stripRegexp(template, `app-component-test[a-z0-9]{6}`, "app-component-test00000")

    return template
}

func stripRegexp(template, pattern, replacement string) string {
	regex := regexp.MustCompile(pattern)
	bytes := regex.ReplaceAll([]byte(template), []byte(replacement))

	return string(bytes)
}