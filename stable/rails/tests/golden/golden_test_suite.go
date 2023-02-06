package golden

import (
	"flag"
	"io/ioutil"

	"regexp"
	"strings"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/suite"
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