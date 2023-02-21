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
	        Templates: []string{"templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "components/serviceaccount.golden.yaml",
	        ValuesFiles: []string{"components/serviceaccount.values.yaml"},
	        Templates: []string{"templates/serviceaccount.yaml", "templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "components/image.golden.yaml",
	        ValuesFiles: []string{"components/image.values.yaml"},
	        Templates: []string{"templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "components/job-scheduling.golden.yaml",
	        ValuesFiles: []string{"components/job-scheduling.values.yaml"},
	        Templates: []string{"templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "components/job-runtime.golden.yaml",
	        ValuesFiles: []string{"components/job-runtime.values.yaml"},
	        Templates: []string{"templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "components/job-advanced.golden.yaml",
	        ValuesFiles: []string{"components/job-advanced.values.yaml"},
	        Templates: []string{"templates/job.yaml"},
	    },
	    {
	        GoldenFileName: "regular-migration/full.golden.yaml",
	        ValuesFiles: []string{"regular-migration/full.values.yaml"},
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

    return template
}

func stripRegexp(template, pattern, replacement string) string {
	regex := regexp.MustCompile(pattern)
	bytes := regex.ReplaceAll([]byte(template), []byte(replacement))

	return string(bytes)
}