package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

var cerealList = []string{
	"fruitloops",
	"cheerios",
	"cocoapuffs",
}

var AuditLocks = []adlr.DependencyLock{
	adlr.DependencyLock{
		Name:    "1",
		Version: "v1",
		License: adlr.License{
			Kind: "fruitloops",
		},
	},
	adlr.DependencyLock{
		Name:    "2",
		Version: "v2",
		License: adlr.License{
			Kind: "cheerios",
		},
	},
	adlr.DependencyLock{
		Name:    "3",
		Version: "v3",
		License: adlr.License{
			Kind: "cocoapuffs",
		},
	},
}

func TestLicenseAuditorAuditLock(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(cerealList)
		a := adlr.MakeLicenseAuditorFromRaw(w)

		err := a.AuditLock(AuditLocks[0])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[1])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[2])
		assert.Nil(t, err)
	})
	t.Run("error on non-whitelist license", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw([]string{"unicorn"})
		a := adlr.MakeLicenseAuditorFromRaw(w)

		lock := AuditLocks[0]
		err := a.AuditLock(lock)
		assert.EqualError(t, err, adlr.NonWhitelistedLicenseErr+lock.License.Kind)

		lock = AuditLocks[1]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, adlr.NonWhitelistedLicenseErr+lock.License.Kind)

		lock = AuditLocks[2]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, adlr.NonWhitelistedLicenseErr+lock.License.Kind)
	})
}

func TestLicenseAuditorAudit(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(cerealList)
		a := adlr.MakeLicenseAuditorFromRaw(w)

		err := a.Audit(AuditLocks)
		assert.Nil(t, err)
	})
	t.Run("error on non-whitelisted licenses", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw([]string{"unicorn"})
		a := adlr.MakeLicenseAuditorFromRaw(w)
		auditErr := "detected non-whitelisted licenses. Remove or Whitelist: [\n " +
			"{\n  " +
			"\"name\": \"1\",\n  " +
			"\"version\": \"v1\",\n  " +
			"\"err\": \"non-whitelisted license: fruitloops\",\n  " +
			"\"license\": {\n   " +
			"\"kind\": \"fruitloops\",\n   " +
			"\"text\": \"\"\n  " +
			"}\n },\n " +
			"{\n  " +
			"\"name\": \"2\",\n  " +
			"\"version\": \"v2\",\n  " +
			"\"err\": \"non-whitelisted license: cheerios\",\n  " +
			"\"license\": {\n   " +
			"\"kind\": \"cheerios\",\n   " +
			"\"text\": \"\"\n  " +
			"}\n },\n " +
			"{\n  " +
			"\"name\": \"3\",\n  " +
			"\"version\": \"v3\",\n  " +
			"\"err\": \"non-whitelisted license: cocoapuffs\",\n  " +
			"\"license\": {\n   " +
			"\"kind\": \"cocoapuffs\",\n   " +
			"\"text\": \"\"\n  " +
			"}\n }\n]"

		err := a.Audit(AuditLocks)
		assert.EqualError(t, err, auditErr)
	})
}
