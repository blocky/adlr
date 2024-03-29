package ascertain

import (
	"errors"
	"math"

	"github.com/go-enry/go-license-detector/v4/licensedb"

	"github.com/blocky/adlr/pkg/reader"
)

// Minimum required confidence for a probable license from text mining
const Confidence float32 = 0.85

// Minimum confidence difference of primary from secondary license matches
const Lead float32 = 0.07

// Holds confidence values for a LicenseMiner
type Minimums struct {
	Confidence float32
	Lead       float32
}

// LicenseMiner attempts to automatically determine Licenses for Mines
type LicenseMiner struct {
	minimum Minimums
	reader  *reader.LimitedReader
}

// Create a LicenseMiner with default values
func MakeLicenseMiner() LicenseMiner {
	mins := Minimums{
		Confidence: Confidence,
		Lead:       Lead,
	}
	reader := reader.NewLimitedReaderFromRaw(reader.Kilobyte * 100)
	return MakeLicenseMinerFromRaw(mins, reader)
}

// Create a LicenseMiner from specified values
func MakeLicenseMinerFromRaw(
	minimums Minimums,
	reader *reader.LimitedReader,
) LicenseMiner {
	return LicenseMiner{minimums, reader}
}

// Attempt to automatically derive a License for a Mine from
// from the licensedb.Match(s) of a licensedb.Result from text mining
// a golang module directory. Return a list of DependencyLock. Returned
// error list can be printed for debugging, and are potentially recoverable
// by LicenseLock
func (lm LicenseMiner) Mine(
	mines ...Mine,
) ([]DependencyLock, error) {
	var mineErrs []Mine

	var locks = make([]DependencyLock, len(mines))
	for i, mine := range mines {
		license, err := lm.MineLicense(mine)

		if err != nil {
			mine.AddError(err)
			mineErrs = append(mineErrs, mine)
		}
		locks[i] = MakeDependencyLock(
			mine.Name,
			mine.Version,
			license,
		)
	}
	if len(mineErrs) != 0 {
		return locks, &LicenseMineError{mineErrs}
	}
	return locks, nil
}

// Attempt to automatically determine a License from for a Mine using
// its licensedb.Match(s)
func (lm LicenseMiner) MineLicense(
	mine Mine,
) (License, error) {
	match, err := lm.DetermineMatch(mine.Matches...)
	if err != nil {
		return License{}, err
	}
	text, err := lm.DetermineLicenseText(mine.Dir + "/" + match.File)
	if err != nil {
		return License{}, err
	}
	return MakeLicense(match.License, text), nil
}

// Fetch license text from a golang module license file
func (lm LicenseMiner) DetermineLicenseText(
	path string,
) (string, error) {
	bytes, err := lm.reader.ReadFileFromPath(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Attempt to automatically determine the correct licensedb.Match
// from a list
func (lm LicenseMiner) DetermineMatch(
	matches ...licensedb.Match,
) (licensedb.Match, error) {
	switch len(matches) {
	case 0:
		return licensedb.Match{}, errors.New("no matches")
	case 1:
		m := matches[0]
		err := lm.DetermineSingleMatch(m)
		if err != nil {
			return licensedb.Match{}, err
		}
		return m, nil
	}
	match, err := lm.DetermineMultipleMatch(matches[0], matches[1])
	if err != nil {
		return licensedb.Match{}, err
	}
	return match, nil
}

// Determine if the singular licensedb.Match meets confidence
// requirement
func (lm LicenseMiner) DetermineSingleMatch(
	m licensedb.Match,
) error {
	err := lm.MeetsMinimumConfidence(m.Confidence)
	if err != nil {
		return err
	}
	return nil
}

// Determine if the primary licensedb.Match meets confidence
// requirements to beat the secondary licensedb.Match
func (lm LicenseMiner) DetermineMultipleMatch(
	m1, m2 licensedb.Match,
) (licensedb.Match, error) {

	var primary, secondary licensedb.Match
	switch compare(m1.Confidence, m2.Confidence) {
	case 1:
		primary, secondary = m1, m2
	case -1:
		primary, secondary = m2, m1
	default:
		primary = lm.DetermineMatchFromEqualConfidences(m1, m2)
		err := lm.MeetsMinimumConfidence(primary.Confidence)
		if err != nil {
			return licensedb.Match{}, err
		}
		return primary, nil
	}

	err := lm.MeetsMinimumConfidence(primary.Confidence)
	if err != nil {
		return licensedb.Match{}, err
	}
	err = lm.MeetsMinimumLead(primary.Confidence, secondary.Confidence)
	if err != nil {
		return licensedb.Match{}, err
	}
	return primary, nil
}

func (lm LicenseMiner) DetermineMatchFromEqualConfidences(
	m1, m2 licensedb.Match,
) (primary licensedb.Match) {
	// return shorter license - longer licenses are usually variants
	diff := len(m1.License) - len(m2.License)
	switch {
	case diff > 0:
		return m2
	case diff < 0:
		return m1
	}
	// licenses are equal length - return first license lexographically
	if m1.License < m2.License {
		return m1
	} else if m1.License == m2.License {
		return m1
	}
	return m2
}

// Determine whether a probable license confidence meets minimum
func (lm LicenseMiner) MeetsMinimumConfidence(
	a float32,
) error {
	var b float32 = lm.minimum.Confidence
	if greaterThan(a, b) {
		return nil
	}
	return &MinConfidenceError{a, b}
}

// Deteremine if the difference of the primary and secondary license
// confidences meets minimum
func (lm LicenseMiner) MeetsMinimumLead(
	a, b float32,
) error {
	var c float32 = lm.minimum.Lead
	if greaterThan(a-b, c) {
		return nil
	}
	return &MinLeadError{a, b, c}
}

func compare(a, b float32) int {
	if a > b {
		return 1
	} else if equalTo(a, b) {
		return 0
	}
	return -1
}

func greaterThan(a, b float32) bool {
	if equalTo(a, b) {
		return false
	}
	return a > b
}

func equalTo(a, b float32) bool {
	return withinTolerance(a - b)
}

func withinTolerance(diff float32) bool {
	var tolerance float64 = 0.00001
	return math.Abs(float64(diff)) < tolerance
}
