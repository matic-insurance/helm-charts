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

func TestMeshGoldenTemplates(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	testCases := []GoldenTestSuite {
	    {
	        GoldenFileName: "defaults/defaults.golden.yaml",
	        ValuesFiles: []string{"defaults/defaults.values.yaml"},
	    },
	    {
	        GoldenFileName: "components/destination-rule.golden.yaml",
	        ValuesFiles: []string{"components/destination-rule.values.yaml"},
	        Templates: []string{"templates/destination-rule.yaml"},
	    },
	    {
	        GoldenFileName: "components/virtual-service.golden.yaml",
	        ValuesFiles: []string{"components/virtual-service.values.yaml"},
	        Templates: []string{"templates/virtual-service.yaml"},
	    },
	    {
	        GoldenFileName: "components/ingress-allow-locations.golden.yaml",
	        ValuesFiles: []string{"components/ingress-allow-locations.values.yaml"},
	        Templates: []string{"templates/virtual-service.yaml"},
	    },
	    {
	        GoldenFileName: "components/ingress-deny-locations.golden.yaml",
	        ValuesFiles: []string{"components/ingress-deny-locations.values.yaml"},
	        Templates: []string{"templates/virtual-service.yaml"},
	    },
	    {
	        GoldenFileName: "components/ingress-deny-and-allow.golden.yaml",
	        ValuesFiles: []string{"components/ingress-deny-and-allow.values.yaml"},
	        Templates: []string{"templates/virtual-service.yaml"},
	    },
	    {
	        GoldenFileName: "components/service-entry.golden.yaml",
	        ValuesFiles: []string{"components/service-entry.values.yaml"},
	        Templates: []string{"templates/service-entry.yaml"},
	    },
	    {
	        GoldenFileName: "features/rewrite-port.golden.yaml",
	        ValuesFiles: []string{"features/rewrite-port.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/disabled.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/disabled.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/single-webserver.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/single-webserver.values.yaml"},
	    },
	    {
	        GoldenFileName: "standard-configurations/webserver-and-websockets.golden.yaml",
	        ValuesFiles: []string{"standard-configurations/webserver-and-websockets.values.yaml"},
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