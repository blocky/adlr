package integration_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blocky/adlr/gotool"
	"github.com/blocky/adlr/internal"
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
	prospects := internal.MakeProspects(direct...)

	// using the paths of the modules, find their licenses
	prospector := internal.MakeLicenseProspector()
	mines, err := prospector.Prospect(prospects...)
	suite.Nil(err)

	// determine (best guess) module license specifics
	miner := internal.MakeLicenseMiner()
	locks, err := miner.Mine(mines...)
	suite.Nil(err)

	// create a license.lock with dependency licenses
	licenselock := internal.MakeLicenseLock("./")
	err = licenselock.Lock(locks...)
	defer os.Remove("./" + internal.LicenseLockName)
	suite.Nil(err)

	// vet license types with whitelist to ensure license
	// requirement fulfillment
	locks, err = licenselock.Read()
	suite.Nil(err)
	auditor := internal.MakeLicenseAuditor()
	err = auditor.Audit(locks...)
	suite.Nil(err)
}
