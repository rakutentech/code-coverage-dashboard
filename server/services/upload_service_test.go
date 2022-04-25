package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	Setup()
}
func TestArchiveDirs(t *testing.T) {
	tests := []struct {
		name       string
		assetsDir  string
		orgName    string
		repoName   string
		branchName string
		errWant    error
	}{
		{
			name:       "Remove, make, then remove archive dir",
			assetsDir:  "../assets_testing/",
			orgName:    "test",
			repoName:   "test",
			branchName: "test",
			errWant:    nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewUploadService()
			err := s.RemoveArchiveDirs(test.assetsDir, test.orgName, test.repoName, test.branchName)
			assert.Equal(t, test.errWant, err)

			dstDir, err := s.MakeArchiveDirs(test.assetsDir, test.orgName, test.repoName, test.branchName)
			assert.Equal(t, test.errWant, err)
			assert.Equal(t, "../assets_testing/test/test/test", dstDir)

			err = s.RemoveArchiveDirs(test.assetsDir, test.orgName, test.repoName, test.branchName)
			assert.Equal(t, test.errWant, err)
		})
	}
}
