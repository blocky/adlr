package integration_test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/adlr-cli/cmd"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	BadBuildListFile   = "./testdata/adlr-cli/bad-buildlist.json"
	BuildListFile      = "../../buildlist.json"
	GotLocatedFile     = "./got-located-licenses.json"
	GotIdentifiedFile  = "./got-identified-licenses.json"
	GotVerifiedFile    = "./got-verified-licenses.json"
	WantLocatedFile    = "./testdata/adlr-cli/located-licenses.json"
	WantIdentifiedFile = "./testdata/adlr-cli/identified-licenses.json"
	WantVerifiedFile   = "./testdata/adlr-cli/verified-licenses.json"
)

func removeGotFiles() {
	_ = os.Remove(GotLocatedFile)
	_ = os.Remove(GotIdentifiedFile)
	_ = os.Remove(GotVerifiedFile)
}

func TestADLRCLI_Locate(t *testing.T) {
	defer removeGotFiles()

	t.Run("happy path - locate licenses and errors for missing", func(t *testing.T) {
		var wantLocated []adlr.Mine
		err := cmd.ReadJSONFile(WantLocatedFile, &wantLocated)
		require.NoError(t, err)

		err = cmd.Locate(BuildListFile, GotLocatedFile)
		assert.Error(t, err)

		var gotLocated []adlr.Mine
		err = cmd.ReadJSONFile(GotLocatedFile, &gotLocated)
		require.NoError(t, err)

		gotMap := lo.SliceToMap(gotLocated, func(got adlr.Mine) (string, adlr.Mine) {
			return got.Name, got
		})

		for _, want := range wantLocated {
			got, ok := gotMap[want.Name]
			assert.True(t, ok)
			assert.Equal(t, want.Version, got.Version)
			assert.Equal(t, want.ErrStr, got.ErrStr)
		}
	})
	t.Run("error opening buildlist file", func(t *testing.T) {
		err := cmd.Locate("nonexistent-buildlist.json", "")
		assert.ErrorContains(t, err, "opening buildlist file")
	})
	t.Run("error parsing module list", func(t *testing.T) {
		err := cmd.Locate(BadBuildListFile, "")
		assert.ErrorContains(t, err, "parsing module list")
	})
	t.Run("error writing located file", func(t *testing.T) {
		err := cmd.Locate(BuildListFile, "nonexistent/dir/located.json")
		assert.ErrorContains(t, err, "writing located file")
	})
}

func TestADLRCLI_Identify(t *testing.T) {
	defer removeGotFiles()
	err := cmd.Locate(BuildListFile, GotLocatedFile)
	require.Error(t, err)

	t.Run("happy path - identify licenses and error for unidentified", func(t *testing.T) {
		var wantIdentified []adlr.DependencyLock
		err = cmd.ReadJSONFile(WantIdentifiedFile, &wantIdentified)
		require.NoError(t, err)

		err = cmd.Identify(GotLocatedFile, GotIdentifiedFile)
		assert.Error(t, err)

		var gotIdentified []adlr.DependencyLock
		err = cmd.ReadJSONFile(GotIdentifiedFile, &gotIdentified)
		require.NoError(t, err)

		gotMap := lo.SliceToMap(gotIdentified, func(got adlr.DependencyLock) (string, adlr.DependencyLock) {
			return got.Name, got
		})

		for _, want := range wantIdentified {
			got, ok := gotMap[want.Name]
			assert.True(t, ok)
			assert.Equal(t, want, got)
		}
	})
	t.Run("error reading located file", func(t *testing.T) {
		err = cmd.Identify("nonexistent-located.json", "")
		assert.ErrorContains(t, err, "reading bytes")
	})
	t.Run("error unmarshaling located list", func(t *testing.T) {
		err = cmd.Identify(BuildListFile, "")
		assert.ErrorContains(t, err, "unmarshaling bytes")
	})
	t.Run("error writing identified file", func(t *testing.T) {
		err = cmd.Identify(GotLocatedFile, "nonexistent/dir/identified.json")
		assert.ErrorContains(t, err, "writing identified file")
	})
}

func TestADLRCLI_Verify(t *testing.T) {
	defer removeGotFiles()
	err := cmd.Locate(BuildListFile, GotLocatedFile)
	require.Error(t, err)

	err = cmd.Identify(GotLocatedFile, GotIdentifiedFile)
	require.Error(t, err)

	t.Run("happy path - verify licenses and error for unverified", func(t *testing.T) {
		var wantVerified []adlr.DependencyLock
		err = cmd.ReadJSONFile(WantVerifiedFile, &wantVerified)
		require.NoError(t, err)

		err = cmd.Verify(GotIdentifiedFile, GotVerifiedFile)
		assert.Error(t, err)

		var gotVerified []adlr.DependencyLock
		err = cmd.ReadJSONFile(GotVerifiedFile, &gotVerified)
		require.NoError(t, err)

		gotMap := lo.SliceToMap(gotVerified, func(got adlr.DependencyLock) (string, adlr.DependencyLock) {
			return got.Name, got
		})

		for _, want := range wantVerified {
			got, ok := gotMap[want.Name]
			assert.True(t, ok)
			assert.Equal(t, want, got)
		}
	})
	t.Run("error reading identified file", func(t *testing.T) {
		err = cmd.Verify("nonexistent-identified.json", "")
		assert.ErrorContains(t, err, "reading identified file")
	})
	t.Run("error unmarshaling identified list", func(t *testing.T) {
		err = cmd.Verify(BuildListFile, "")
		assert.ErrorContains(t, err, "unmarshaling bytes")
	})
	t.Run("error writing verified file", func(t *testing.T) {
		err = cmd.Verify(GotIdentifiedFile, "nonexistent/dir/verified.json")
		assert.ErrorContains(t, err, "writing verified file")
	})
}
