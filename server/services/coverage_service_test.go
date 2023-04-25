package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	Setup()
}

func TestFindCoverageXMLPath(t *testing.T) {
	tests := []struct {
		language string
		folder   string
		wantPath string
		errWant  error
	}{
		{
			language: "php",
			folder:   "../test_data/php_coverage/",
			wantPath: "../test_data/php_coverage/coverage.xml",
			errWant:  nil,
		},
		{
			language: "go",
			folder:   "../test_data/go_coverage/",
			wantPath: "../test_data/go_coverage/coverage.xml",
			errWant:  nil,
		},
		{
			language: "clover",
			folder:   "../test_data/clover_coverage/",
			wantPath: "../test_data/clover_coverage/coverage.xml",
			errWant:  nil,
		},
		{
			language: "java",
			folder:   "../test_data/java_coverage/",
			wantPath: "../test_data/java_coverage/coverage.xml",
			errWant:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.language, func(t *testing.T) {
			s := NewCoverageService()
			got, err := s.FindCoverageXMLPath(test.folder, "coverage.xml")
			assert.Equal(t, test.errWant, err)
			assert.Equal(t, test.wantPath, got)

		})
	}
}
func TestParseCoveragePercentage(t *testing.T) {
	tests := []struct {
		language        string
		coverageXMLPath string
		wantPercentage  float64
		errWant         error
	}{
		{
			language:        "php",
			coverageXMLPath: "../test_data/php_coverage/coverage.xml",
			wantPercentage:  76.85,
			errWant:         nil,
		},
		{
			language:        "go",
			coverageXMLPath: "../test_data/go_coverage/coverage.xml",
			wantPercentage:  33.03,
			errWant:         nil,
		},
		{
			language:        "js",
			coverageXMLPath: "../test_data/clover_coverage/coverage.xml",
			wantPercentage:  97.91,
			errWant:         nil,
		},
		{
			language:        "java",
			coverageXMLPath: "../test_data/java_coverage/coverage.xml",
			wantPercentage:  63.04,
			errWant:         nil,
		},
	}
	for _, test := range tests {
		t.Run(test.language, func(t *testing.T) {
			s := NewCoverageService()
			got, err := s.ParseCoveragePercentage(test.coverageXMLPath)
			assert.Equal(t, test.errWant, err)
			assert.Equal(t, test.wantPercentage, got)

		})
	}
}
