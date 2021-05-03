package integration_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/gotool"
)

type IntegrationTestSuite struct {
	suite.Suite
}

func TestIntegrationTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration test")
	}
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite IntegrationTestSuite) TestADLR() {
	buildlist, err := os.Open("../buildlist.json")
	if err != nil {
		panic("no buildlist found")
	}
	defer buildlist.Close()

	// unmarshal adlr's golang buildlist into modules
	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(buildlist)
	suite.Nil(err)

	// filter out non-direct modules from the buildlist
	direct := gotool.FilterDirectImportModules(mods)
	deps := adlr.MakeDependencies(direct...)

	// using the paths of the modules, find their licenses
	prospector := adlr.MakeLicenseProspector()
	prospects, err := prospector.Prospect(deps...)
	suite.Nil(err)

	// determine (best guess) module license specifics
	miner := adlr.MakeLicenseMiner()
	mined, err := miner.Mine(prospects...)
	suite.Nil(err)

	// create a license.lock with dependency licenses
	licenselock := adlr.MakeLicenseLock("./")
	err = licenselock.Lock(mined)
	defer os.Remove("./" + adlr.LicenseLockName)
	suite.Nil(err)

	// vet license types with whitelist to ensure license
	// requirement fulfillment
	locks, err := licenselock.Read()
	suite.Nil(err)
	auditor := adlr.MakeLicenseAuditor()
	suite.Nil(err)
	err = auditor.AuditLocks(locks)
	suite.Nil(err)
}
