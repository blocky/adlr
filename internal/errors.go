package internal

import (
	"encoding/json"
	"fmt"

	"github.com/blocky/prettyprinter"
)

const (
	MinConfErr  = "does not meet minimum confidence: got %f lt %f"
	MinLeadErr  = "does not meet minimum lead: %f - %f lt %f"
	LicProsErr  = "error license prospecting: %v"
	LicMineErr  = "error license mining: %v"
	LicAuditErr = "detected non-whitelisted licenses. Remove or Whitelist: %v"
	DepLockErr  = "error locking dependencies %v:"
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
	Prospects []Prospect
}

func (lpe LicenseProspectingError) Error() string {
	bytes, err := json.MarshalIndent(lpe.Prospects, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(LicProsErr, string(bytes))
}

type LicenseMineError struct {
	Mines []Mine
}

func (lme LicenseMineError) Error() string {
	bytes, err := json.MarshalIndent(lme.Mines, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(LicMineErr, string(bytes))
}

type DependencyLockerError struct {
	Locks []DependencyLock
}

func (dle DependencyLockerError) Error() string {
	bytes, err := json.MarshalIndent(dle.Locks, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(DepLockErr, string(bytes))
}

type LockerError struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	Err []prettyprinter.FieldError `json:"errors"`
}

func MakeLockerError(
	name, version string,
	errs ...error,
) LockerError {
	return LockerError{
		Name:    name,
		Version: version,
		Err:     makeFieldErrors(errs),
	}
}

func makeFieldErrors(
	errs []error,
) []prettyprinter.FieldError {
	var fieldErrs = make([]prettyprinter.FieldError, len(errs))
	for i, err := range errs {
		fieldErrs[i] = prettyprinter.FieldError{err}
	}
	return fieldErrs
}

type LicenseAuditError struct {
	Locks []DependencyLock
}

func (lae LicenseAuditError) Error() string {
	bytes, err := json.MarshalIndent(lae.Locks, "", " ")
	if err != nil {
		return fmt.Sprintf("could not marshal: %v", err)
	}
	return fmt.Sprintf(LicAuditErr, string(bytes))
}
