package integration_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/blocky/adlr/pkg/reader"
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
	suite.Require().NoError(err)

	direct := gotool.FilterDirectImportModules(mods)
	prospects := adlr.MakeProspects(direct...)

	prospector := adlr.MakeProspector()
	mines, err := prospector.Prospect(prospects...)
	suite.Require().NoError(err)

	confidence := float32(0.85)
	lead := float32(0.05)
	reader := reader.NewLimitedReaderFromRaw(500 * reader.Kilobyte)
	miner := adlr.MakeMinerFromRaw(confidence, lead, reader)
	locks, err := miner.Mine(mines...)
	suite.Require().NoError(err)

	licenselock := adlr.MakeLicenseLockManager("./")
	err = licenselock.Lock(locks...)
	defer func() {
		err := os.Remove("./" + "license.lock")
		suite.Require().NoError(err)
	}()
	suite.Require().NoError(err)

	locks, err = licenselock.Read()
	suite.Require().NoError(err)

	whitelist := adlr.MakeWhitelist(adlr.DefaultWhitelist)
	auditor := adlr.MakeAuditor(whitelist)
	err = auditor.Audit(locks...)
	suite.Require().NoError(err)
}
