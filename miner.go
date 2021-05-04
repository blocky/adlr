package adlr

import (
	"errors"
	"math"

	"github.com/go-enry/go-license-detector/v4/licensedb"

	"github.com/blocky/adlr/reader"
)

const Confidence float32 = 0.85
const Lead float32 = 0.07

type Minimums struct {
	Confidence float32
	Lead       float32
}

type LicenseMiner struct {
	minimum Minimums
	reader  *reader.LimitedReader
}

func MakeLicenseMiner() LicenseMiner {
	mins := Minimums{
		Confidence: Confidence,
		Lead:       Lead,
	}
	reader := reader.NewLimitedReaderFromRaw(reader.Kilobyte * 100)
	return MakeLicenseMinerFromRaw(mins, reader)
}

func MakeLicenseMinerFromRaw(
	minimums Minimums,
	reader *reader.LimitedReader,
) LicenseMiner {
	return LicenseMiner{minimums, reader}
}

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

func (lm LicenseMiner) DetermineLicenseText(
	path string,
) (string, error) {
	bytes, err := lm.reader.ReadFileFromPath(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

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
	// licensedb matches are sorted by decreasing confidence
	m1, m2 := matches[0], matches[1]
	err := lm.DetermineMultipleMatch(m1, m2)
	if err != nil {
		return licensedb.Match{}, err
	}
	return m1, nil
}

func (lm LicenseMiner) DetermineSingleMatch(
	m licensedb.Match,
) error {
	err := lm.MeetsMinimumConfidence(m.Confidence)
	if err != nil {
		return err
	}
	return nil
}

func (lm LicenseMiner) DetermineMultipleMatch(
	m1, m2 licensedb.Match,
) error {
	err := lm.MeetsMinimumConfidence(m1.Confidence)
	if err != nil {
		return err
	}
	err = lm.MeetsMinimumLead(m1.Confidence, m2.Confidence)
	if err != nil {
		return err
	}
	return nil
}

func (lm LicenseMiner) MeetsMinimumConfidence(
	a float32,
) error {
	var b float32 = lm.minimum.Confidence
	if greaterThan(a, b) {
		return nil
	}
	return &MinConfidenceError{a, b}
}

func (lm LicenseMiner) MeetsMinimumLead(
	a, b float32,
) error {
	var c float32 = lm.minimum.Lead
	if greaterThan(a-b, c) {
		return nil
	}
	return &MinLeadError{a, b, c}
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
