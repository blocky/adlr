package adlr

import (
	"encoding/json"
	"fmt"

	"github.com/go-enry/go-license-detector/v4/licensedb"
)

const (
	MinConfErr = "does not meet minimum confidence: got %f lt %f"
	MinLeadErr = "does not meet minimum lead: %f - %f lt %f"
	LicProsErr = "error license prospecting: %v"
	LicMineErr = "error license mining: %v"
)

type MinConfidenceError struct {
	Got, MinConf float32
}

func (mce MinConfidenceError) Error() string {
	return fmt.Sprintf(MinConfErr, mce.Got, mce.MinConf)
}

type MinLeadError struct {
	Got1, Got2, MinLead float32
}

func (mle MinLeadError) Error() string {
	return fmt.Sprintf(MinLeadErr, mle.Got1, mle.Got2, mle.MinLead)
}

type LicenseProspectingError struct {
	Results []licensedb.Result
}

func (lpe LicenseProspectingError) Error() string {
	bytes, err := json.MarshalIndent(lpe.Results, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(LicProsErr, string(bytes))
}

type LicenseMineError struct {
	Deps []Dependency
}

func (lme LicenseMineError) Error() string {
	bytes, err := json.MarshalIndent(lme.Deps, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(LicMineErr, string(bytes))
}