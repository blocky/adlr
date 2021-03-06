package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
)

const (
	ProspectorHappyPathDir      = "./testdata/prospector/happypath"
	ProspectorMissingDir        = "./unicorn"
	ProspectorMissingLicenseDir = "./testdata/prospector/missinglicense"
)

func (suite *IntegrationTestSuite) TestProspector() {
	suite.T().Run("happy path", func(t *testing.T) {
		p := ascertain.MakeLicenseProspector()
		r := p.ProspectLicenses(ProspectorHappyPathDir)[0]

		assert.Equal(t, "", r.ErrStr)
		assert.Equal(t, "MIT", r.Matches[0].License)
		assert.InDelta(t, 0.92, r.Matches[0].Confidence, 0.2)
	})
	suite.T().Run("error on missing dir", func(t *testing.T) {
		p := ascertain.MakeLicenseProspector()
		r := p.ProspectLicenses(ProspectorMissingDir)[0]
		// error remains same regardless of internet connection
		expected := "could not clone repo from " +
			ProspectorMissingDir +
			": repository not found"

		assert.Equal(t, expected, r.ErrStr)
	})
	suite.T().Run("error on missing license file", func(t *testing.T) {
		p := ascertain.MakeLicenseProspector()
		r := p.ProspectLicenses(ProspectorMissingLicenseDir)[0]

		assert.Equal(t, "no license file was found", r.ErrStr)
	})
}
