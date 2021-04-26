package integration_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
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
