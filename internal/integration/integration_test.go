package integration_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blocky/adlr/api"
	"github.com/blocky/adlr/pkg/gotool"
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
	buildlist, err := os.Open("../../buildlist.json")
	if err != nil {
		panic("no buildlist found")
	}
	defer buildlist.Close()

	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(buildlist)
	suite.Nil(err)

	direct := gotool.FilterDirectImportModules(mods)
	prospects := api.MakeProspects(direct...)

	prospector := api.MakeProspector()
	mines, err := prospector.Prospect(prospects...)
	suite.Nil(err)

	miner := api.MakeMiner()
	locks, err := miner.Mine(mines...)

	licenselock := api.MakeLicenseLockManager("./")
	err = licenselock.Lock(locks...)
	defer os.Remove("./" + "license.lock")
	suite.Nil(err)

	locks, err = licenselock.Read()
	suite.Nil(err)

	whitelist := api.MakeWhitelist(api.DefaultWhitelist)
	auditor := api.MakeAuditor(whitelist)
	err = auditor.Audit(locks...)
	suite.Nil(err)
}
