package ascertain_test

import (
	"io/ioutil"
	"testing"

	"github.com/go-enry/go-license-detector/v4/licensedb"
	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
	"github.com/blocky/adlr/pkg/reader"
)

const MinerHappyPathPath = "./testdata/miner/happypath/license"
const MinerMissingFilePath = "./testdata/miner/missinglicense/"

func TestLicenseMinerMeetsMinimumConfidence(t *testing.T) {
	var minConf float32 = 0.5
	t.Run("happy path", func(t *testing.T) {
		m := float32(0.9)
		mins := ascertain.Minimums{Confidence: minConf}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumConfidence(m)
		assert.Nil(t, err)
	})
	t.Run("error on equal to", func(t *testing.T) {
		m := minConf
		confErr := &ascertain.MinConfidenceError{m, minConf}
		mins := ascertain.Minimums{Confidence: minConf}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumConfidence(m)
		assert.EqualError(t, err, confErr.Error())
	})
	t.Run("error on less than", func(t *testing.T) {
		m := float32(0.1)
		confErr := &ascertain.MinConfidenceError{m, minConf}
		mins := ascertain.Minimums{Confidence: minConf}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumConfidence(m)
		assert.EqualError(t, err, confErr.Error())
	})
}

func TestLicenseMinerMeetsMinimumLead(t *testing.T) {
	var minLead float32 = 0.2
	t.Run("happy path", func(t *testing.T) {
		m1 := float32(0.9)
		m2 := float32(0.1)
		mins := ascertain.Minimums{Lead: minLead}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumLead(m1, m2)
		assert.Nil(t, err)
	})
	t.Run("error on equal to", func(t *testing.T) {
		m1 := float32(minLead)
		m2 := float32(minLead)
		leadErr := &ascertain.MinLeadError{m1, m2, minLead}
		mins := ascertain.Minimums{Lead: minLead}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumLead(m1, m2)
		assert.EqualError(t, err, leadErr.Error())
	})
	t.Run("error on less than", func(t *testing.T) {
		m1 := float32(0.9)
		m2 := float32(0.8)
		leadErr := &ascertain.MinLeadError{m1, m2, minLead}
		mins := ascertain.Minimums{Lead: minLead}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)

		err := miner.MeetsMinimumLead(m1, m2)
		assert.EqualError(t, err, leadErr.Error())
	})
}

func TestLicenseMinerDetermineLicenseText(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		path := MinerHappyPathPath
		bytes, err := ioutil.ReadFile(path)
		assert.Nil(t, err)
		expected := string(bytes)

		miner := ascertain.MakeLicenseMiner()
		text, err := miner.DetermineLicenseText(path)

		assert.Nil(t, err)
		assert.Equal(t, expected, text)
	})
	t.Run("error on bad path", func(t *testing.T) {
		path := MinerMissingFilePath + "unicorn"
		expected := "open " + path + ": no such file or directory"
		miner := ascertain.MakeLicenseMiner()
		_, err := miner.DetermineLicenseText(path)

		assert.EqualError(t, err, expected)
	})
}

func TestLicenseMinerDetermineMatch(t *testing.T) {
	t.Run("error on 0 matches", func(t *testing.T) {
		matches := []licensedb.Match{}
		miner := ascertain.MakeLicenseMiner()
		_, err := miner.DetermineMatch(matches...)

		assert.EqualError(t, err, "no matches")
	})
	t.Run("happy path single match", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.9, License: "MIT"},
		}
		mins := ascertain.Minimums{Confidence: 0.75, Lead: 0.0}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		m, err := miner.DetermineMatch(matches...)

		assert.Nil(t, err)
		assert.Equal(t, matches[0], m)
	})
	t.Run("conf error on single match", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.5, License: "MIT"},
		}
		confErr := &ascertain.MinConfidenceError{0.5, 0.6}
		mins := ascertain.Minimums{Confidence: 0.6, Lead: 0.0}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		_, err := miner.DetermineMatch(matches...)

		assert.EqualError(t, err, confErr.Error())
	})
	t.Run("happy path multiple match", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.9, License: "MIT1"},
			licensedb.Match{Confidence: 0.5, License: "MIT2"},
		}
		mins := ascertain.Minimums{Confidence: 0.0, Lead: 0.0}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		m, err := miner.DetermineMatch(matches...)

		assert.Nil(t, err)
		assert.Equal(t, matches[0], m)
	})
	t.Run("happy path multiple match same confidence", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.9, License: "MIT"},
			licensedb.Match{Confidence: 0.9, License: "MIT-variant-2"},
		}
		mins := ascertain.Minimums{Confidence: 0.8, Lead: 0.1}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		m, err := miner.DetermineMatch(matches...)

		assert.Nil(t, err)
		assert.Equal(t, matches[0], m)
	})
	t.Run("happy path multiple match same confidence similar SPDX", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.9, License: "MIT-variant-2"},
			licensedb.Match{Confidence: 0.9, License: "MIT-variant-1"},
		}
		mins := ascertain.Minimums{Confidence: 0.8, Lead: 0.1}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		m, err := miner.DetermineMatch(matches...)

		assert.Nil(t, err)
		assert.Equal(t, matches[1], m)
	})
	t.Run("conf error on multiple match", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.1, License: "MIT1"},
			licensedb.Match{Confidence: 0.1, License: "MIT2"},
		}
		confErr := &ascertain.MinConfidenceError{0.1, 0.2}
		mins := ascertain.Minimums{Confidence: 0.2, Lead: 0.0}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		_, err := miner.DetermineMatch(matches...)

		assert.EqualError(t, err, confErr.Error())
	})
	t.Run("lead error on multiple match", func(t *testing.T) {
		matches := []licensedb.Match{
			licensedb.Match{Confidence: 0.9, License: "MIT1"},
			licensedb.Match{Confidence: 0.8, License: "MIT2"},
		}
		leadErr := &ascertain.MinLeadError{0.9, 0.8, 0.3}
		mins := ascertain.Minimums{Confidence: 0.0, Lead: 0.3}
		reader := reader.NewLimitedReader()
		miner := ascertain.MakeLicenseMinerFromRaw(mins, reader)
		_, err := miner.DetermineMatch(matches...)

		assert.EqualError(t, err, leadErr.Error())
	})
}
