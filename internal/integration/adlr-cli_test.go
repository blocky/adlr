package integration_test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/blocky/adlr/adlr-cli/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	BadBuildListFile = "./testdata/adlr-cli/bad-buildlist.json"
	BuildListFile    = "./testdata/adlr-cli/buildlist.json"
	LocatedFile      = "./testdata/adlr-cli/located-licenses.json"
	IdentifiedFile   = "./testdata/adlr-cli/identified-licenses.json"
)

//go:embed testdata/adlr-cli/located-licenses.json
var located []byte

//go:embed testdata/adlr-cli/identified-licenses.json
var identified []byte

//go:embed testdata/adlr-cli/verified-licenses.json
var verified []byte

//go:embed testdata/adlr-cli/unlocated-licenses.json
var unlocated string

//go:embed testdata/adlr-cli/unidentified-licenses.json
var unidentified string

//go:embed testdata/adlr-cli/unverified-licenses.json
var unverified string

func TestADLRCLI_Locate(t *testing.T) {
	gotLocatedFile := "./got-located-licenses.json"
	defer func() {
		err := os.Remove(gotLocatedFile)
		require.NoError(t, err)
	}()

	t.Run("happy path - located licenses and listed missing", func(t *testing.T) {
		wantLocated := located
		err := cmd.Locate(BuildListFile, gotLocatedFile)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), unlocated)

		gotLocated, err := os.ReadFile(gotLocatedFile)
		require.NoError(t, err)
		assert.Equal(t, gotLocated, wantLocated)
	})
	t.Run("error opening buildlist file", func(t *testing.T) {
		err := cmd.Locate("nonexistent-buildlist.json", gotLocatedFile)
		assert.ErrorContains(t, err, "opening buildlist file")
	})
	t.Run("error parsing module list", func(t *testing.T) {
		err := cmd.Locate(BadBuildListFile, gotLocatedFile)
		assert.ErrorContains(t, err, "parsing module list")
	})
	t.Run("error writing located file", func(t *testing.T) {
		err := cmd.Locate(BuildListFile, "nonexistent/dir/located.json")
		assert.ErrorContains(t, err, "writing located file")
	})
}

func TestADLRCLI_Identify(t *testing.T) {
	gotIdentifiedFile := "./got-identified-licenses.json"
	defer func() {
		err := os.Remove(gotIdentifiedFile)
		require.NoError(t, err)
	}()

	t.Run("happy path - identified licenses and listed unidentified", func(t *testing.T) {
		wantIdentified := identified
		err := cmd.Identify(LocatedFile, gotIdentifiedFile)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), unidentified)

		gotIdentified, err := os.ReadFile(gotIdentifiedFile)
		require.NoError(t, err)
		assert.Equal(t, gotIdentified, wantIdentified)
	})
	t.Run("error opening located file", func(t *testing.T) {
		err := cmd.Identify("nonexistent-located.json", gotIdentifiedFile)
		assert.ErrorContains(t, err, "opening located file")
	})
	t.Run("error decoding located list", func(t *testing.T) {
		err := cmd.Identify(BuildListFile, gotIdentifiedFile)
		assert.ErrorContains(t, err, "decoding located list")
	})
	t.Run("error writing identified file", func(t *testing.T) {
		err := cmd.Identify(LocatedFile, "nonexistent/dir/identified.json")
		assert.ErrorContains(t, err, "writing identified file")
	})
}

func TestADLRCLI_Verify(t *testing.T) {
	gotVerifiedFile := "./got-verified-licenses.json"
	defer func() {
		err := os.Remove(gotVerifiedFile)
		require.NoError(t, err)
	}()

	t.Run("happy path - verified licenses and listed unverified", func(t *testing.T) {
		wantVerified := verified
		err := cmd.Verify(IdentifiedFile, gotVerifiedFile)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), unverified)

		gotVerified, err := os.ReadFile(gotVerifiedFile)
		require.NoError(t, err)
		assert.Equal(t, gotVerified, wantVerified)
	})
	t.Run("error opening identified file", func(t *testing.T) {
		err := cmd.Verify("nonexistent-identified.json", gotVerifiedFile)
		assert.ErrorContains(t, err, "opening identified file")
	})
	t.Run("error decoding identified list", func(t *testing.T) {
		err := cmd.Verify(BuildListFile, gotVerifiedFile)
		assert.ErrorContains(t, err, "decoding identified list")
	})
	t.Run("error writing verified file", func(t *testing.T) {
		err := cmd.Verify(IdentifiedFile, "nonexistent/dir/verified.json")
		assert.ErrorContains(t, err, "writing verified file")
	})
}
