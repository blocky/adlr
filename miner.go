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
}

func MakeLicenseMiner() LicenseMiner {
	mins := Minimums{
		Confidence: Confidence,
		Lead:       Lead,
	}
	return MakeLicenseMinerFromRaw(mins)
}

func MakeLicenseMinerFromRaw(
	minimums Minimums,
) LicenseMiner {
	return LicenseMiner{minimums}
}

func (lm LicenseMiner) Mine(
	prospects ...Dependency,
) ([]Dependency, error) {
	var mineErrs []Dependency

	for i := 0; i < len(prospects); i++ {
		license, err := lm.MineLicense(prospects[i])

		if err != nil {
			prospects[i].AddError(err)
			mineErrs = append(mineErrs, prospects[i])
			continue
		} else {
			prospects[i].AddLicense(license)
		}
	}
	if len(mineErrs) != 0 {
		// print this err. license lock will attempt to handle these
		return prospects, &LicenseMineError{mineErrs}
	}
	return prospects, nil
}

func (lm LicenseMiner) MineLicense(
	dep Dependency,
) (License, error) {
	match, err := lm.DetermineMatch(dep.Result.Matches...)
	if err != nil {
		return License{}, err
	}
	text, err := lm.DetermineLicenseText(dep.Module.Dir + "/" + match.File)
	if err != nil {
		return License{}, err
	}
	return MakeLicense(match.License, text), nil
}

func (lm LicenseMiner) DetermineLicenseText(
	path string,
) (string, error) {
	reader := reader.NewLimitedReader()
	bytes, err := reader.ReadFileFromPath(path)
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
