package adlr

import (
	"encoding/json"
	"fmt"

	"github.com/go-enry/go-license-detector/v4/licensedb"
)

const (
	LicProsErr = "error license prospecting: %v"
)

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
