package integration_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ionrock/procs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testdata "github.com/blocky/adlr/internal/integration/testdata/cli"
)

func deleteLicenseFiles() {
	_ = os.Remove("./located-licenses.json")
	_ = os.Remove("./identified-licenses.json")
	_ = os.Remove("./verified-licenses.json")
}

func runADLRCLI(args string) (string, error) {
	cmd := fmt.Sprintf("go run ../../main.go license %s", args)
	adlrCLI := procs.NewProcess(cmd)
	adlrCLI.OutputHandler = func(line string) string { return line }
	adlrCLI.ErrHandler = func(line string) string { return line }

	err := adlrCLI.Start()
	if err != nil {
		return "", fmt.Errorf("starting adlr-cli: %w", err)
	}

	_ = adlrCLI.Wait()

	errOutBytes, err := adlrCLI.ErrOutput()
	if err != nil {
		lErr := fmt.Errorf("getting adlr-cli err output: %w", err)
		return "", lErr
	}
	return string(errOutBytes), nil
}

func TestADLRCLILicenseLocate(t *testing.T) {
	// given
	wantErrOut := testdata.LocateErrOut
	defer deleteLicenseFiles()

	// when
	args := "locate -p ../../"
	errOut, err := runADLRCLI(args)
	require.NoError(t, err)

	// then
	errOut = strings.ReplaceAll(errOut, "error license prospecting:", "")
	errOut = strings.ReplaceAll(errOut, "exit status 1", "")
	assert.JSONEq(t, wantErrOut, errOut)
}

func TestADLRCLILicenseIdentify(t *testing.T) {
	// given
	wantErrOut := testdata.IdentifyErrOut
	defer deleteLicenseFiles()

	args := "locate -p ../../"
	_, err := runADLRCLI(args)
	require.NoError(t, err)

	// when
	args = "identify ../../"
	errOut, err := runADLRCLI(args)
	require.NoError(t, err)

	// then
	errOut = strings.ReplaceAll(errOut, "error license mining:", "")
	errOut = strings.ReplaceAll(errOut, "exit status 1", "")
	assert.JSONEq(t, wantErrOut, errOut)
}

func TestADLRCLILicenseVerify(t *testing.T) {
	// given
	wantErrOut := testdata.VerifyErrOut
	defer deleteLicenseFiles()

	args := "locate -p ../../"
	_, err := runADLRCLI(args)
	require.NoError(t, err)

	args = "identify ../../"
	_, err = runADLRCLI(args)
	require.NoError(t, err)

	// when
	args = "verify ../../"
	errOut, err := runADLRCLI(args)
	require.NoError(t, err)

	// then
	errOut = strings.ReplaceAll(errOut, "detected non-whitelisted licenses. Remove or Whitelist:", "")
	errOut = strings.ReplaceAll(errOut, "exit status 1", "")
	assert.JSONEq(t, wantErrOut, errOut)
}
